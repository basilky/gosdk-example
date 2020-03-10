package main

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
	"strings"

	mspproto "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/common/cauthdsl"
)

type OrgSetup struct {
	orgName string
	adminID string
	sdk     *fabsdk.FabricSDK
}

func main() {

	org1Setup := OrgSetup{}
	org1Setup.orgName = "Org1"
	//Initialize SDK for Org1
	Org1SDK, err := sdkconnector.CreateSDKInstance("Org1")
	if err != nil {
		fmt.Println("error creating SDK for Org1 : ", err)
		return
	}
	org1Setup.sdk = Org1SDK
	//Initialize SDK for Org2
	Org2SDK, err := sdkconnector.CreateSDKInstance("Org2")
	if err != nil {
		fmt.Println("error creating SDK for Org2 : ", err)
	}
	fmt.Println("SDK created for Org1 and Org2")

	//Register and enroll admin user on Org1
	Org1Admin := &mspclient.RegistrationRequest{
		Name:           "org1admin",
		Type:           "admin",
		MaxEnrollments: 10,
		Affiliation:    "org1.department1",
		CAName:         "ca.org1.example.com",
	}
	err = sdkconnector.RegisterandEnroll(Org1SDK, "Org1", Org1Admin)
	if err != nil {
		fmt.Println("error on registering and enrolling admin user for Org1 : ", err)
		return
	}
	org1Setup.adminID = "Admin"

	//Register and enroll admin user on Org2
	Org2Admin := &mspclient.RegistrationRequest{
		Name:           "org2admin",
		Type:           "admin",
		MaxEnrollments: 10,
		Affiliation:    "org2.department1",
		CAName:         "ca.org2.example.com",
	}
	err = sdkconnector.RegisterandEnroll(Org2SDK, "Org2", Org2Admin)
	if err != nil {
		fmt.Println("error on registering and enrolling admin user for Org2 : ", err)
		return
	}
	fmt.Println("Enrolled admins for Org1 and Org2")

	//Create mychannel using org1admin
	sdkconnector.CreateChennel(Org1SDK, "Org1", "org1admin", "mychannel", "network/channel-artifacts/channel.tx")
	if err != nil {
		fmt.Println("error creating channel : ", err)
		return
	}
	fmt.Println("Created channel mychannel")

	//Join Org1 peers to mychannel
	err = sdkconnector.JoinChennel(Org1SDK, "Org1", "org1admin", "mychannel")
	if err != nil {
		fmt.Println("error joining Org1 peers to channel : ", err)
		return
	}

	//Join Org2 peers to mychannel
	err = sdkconnector.JoinChennel(Org2SDK, "Org2", "org2admin", "mychannel")
	if err != nil {
		fmt.Println("error joining Org2 peers to channel : ", err)
		return
	}
	fmt.Println("Joined Org1 and Org2 peers to mychannel")

	//Install chaincode on Org1 peers
	err = sdkconnector.InstallCC(Org1SDK, "Org1", "org1admin", "gosdk-example/chaincode/golang", "mycc", "v0", "localhost:7051")
	if err != nil {
		fmt.Println("Error installing chaincode on Org1 peers : ", err)
		return
	}

	//Install chaincode on Org2 peers
	err = sdkconnector.InstallCC(Org2SDK, "Org2", "org2admin", "gosdk-example/chaincode/golang", "mycc", "v0", "localhost:9051")
	if err != nil {
		fmt.Println("Error installing chaincode on Org2 peers : ", err)
		return
	}
	fmt.Println("Chaincode installed on Org1 and Org2 peers")

	//Create chaincode policy (this policy requires transactions to be endorsed by member of both Org1 and Org2)
	ccPolicy := cauthdsl.SignedByNOutOfGivenRole(2, mspproto.MSPRole_MEMBER, []string{"Org1MSP", "Org2MSP"})

	//Instantiate chaincode using Org1 admin identity
	fmt.Println("Trying to instantiate chaincode...")
	instCCrequest := resmgmt.InstantiateCCRequest{Name: "mycc", Path: "chaincode/golang", Version: "v0", Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	peers := []string{"localhost:7051", "localhost:9051"}
	err = sdkconnector.InstantiateCC(Org1SDK, "Org1", "org1admin", "mychannel", instCCrequest, peers)
	if err != nil {
		fmt.Println(err, "failed to instantiate the chaincode")
		return
	}
	fmt.Println("Chaincode instantiation successful")

	//Register and enroll normal user on Org2
	Org2User := &mspclient.RegistrationRequest{
		Name:           "org2user",
		Type:           "client",
		MaxEnrollments: 10,
		Affiliation:    "org2.department1",
		CAName:         "ca.org2.example.com",
	}
	err = sdkconnector.RegisterandEnroll(Org2SDK, "Org2", Org2User)
	if err != nil {
		fmt.Println("error on registering and enrolling org2user user for Org2 : ", err)
		return
	}
	fmt.Println("Enrolled normal user on Org2")

	//Channel client is used to query and execute transactions
	clientContext := Org2SDK.ChannelContext("mychannel", fabsdk.WithUser("org2user"))
	client, err := channel.New(clientContext)
	if err != nil {
		fmt.Println(err, "failed to create new channel client")
		return
	}

	//Execute initLedger fabcar transaction
	res, err := client.Execute(channel.Request{ChaincodeID: "mycc", Fcn: "initLedger", Args: nil, TransientMap: nil})
	if err != nil {
		fmt.Println("Error execute transaction : ", err)
		return
	} else {
		fmt.Println("\ninitLedger transaction success, ID : ", res.TransactionID)
	}

	//Chaincode query queryAllCars function
	response, err := client.Query(channel.Request{ChaincodeID: "mycc", Fcn: "queryAllCars", Args: [][]byte{}})
	if err != nil {
		fmt.Println("Error queryAllCars: ", err)
		return
	} else {
		fmt.Println("\nQuery Response : ", string(response.Payload))
	}

	//Execute createCar fabcar transaction
	res, err = client.Execute(channel.Request{ChaincodeID: "mycc", Fcn: "createCar", Args: [][]byte{[]byte("CAR12"), []byte("Honda"), []byte("Accord"), []byte("Black"), []byte("Tom")}, TransientMap: nil})
	if err != nil {
		fmt.Println("Error execute transaction : ", err)
		return
	} else {
		fmt.Println("\ncreateCar transaction success, ID : ", res.TransactionID)
	}

	//Chaincode query queryCar function
	response, err = client.Query(channel.Request{ChaincodeID: "mycc", Fcn: "queryCar", Args: [][]byte{[]byte("CAR12")}})
	if err != nil {
		fmt.Println("Error queryCar: ", err)
		return
	} else {
		fmt.Println("\nQuery 'CAR12' Response : ", string(response.Payload))
	}

	//Execute changeCarOwner fabcar transaction
	res, err = client.Execute(channel.Request{ChaincodeID: "mycc", Fcn: "changeCarOwner", Args: [][]byte{[]byte("CAR12"), []byte("Dave")}, TransientMap: nil})
	if err != nil {
		fmt.Println("Error execute transaction : ", err)
		return
	} else {
		fmt.Println("\nchangeCarOwner transaction success, ID : ", res.TransactionID)
	}

	//Chaincode query queryCar function
	response, err = client.Query(channel.Request{ChaincodeID: "mycc", Fcn: "queryCar", Args: [][]byte{[]byte("CAR12")}})
	if err != nil {
		fmt.Println("Error queryCar: ", err)
		return
	} else {
		fmt.Println("\nQuery 'CAR12' Response : ", string(response.Payload))
	}
	fmt.Println()

	http.HandleFunc("/users", org1Setup.EnrollUser)
	http.ListenAndServe(":3000", nil)
}

func (orgSetup *OrgSetup) EnrollUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	username := r.FormValue("username")
	orgname := r.FormValue("orgname")
	fmt.Fprintf(w, "Name = %s\n", username)
	fmt.Fprintf(w, "Address = %s\n", orgname)
	//Register and enroll normal user on Org2
	Org2User := &mspclient.RegistrationRequest{
		Name:           username,
		Type:           "client",
		MaxEnrollments: 10,
		Affiliation:    strings.ToLower(orgname) + ".department1",
		CAName:         "ca." + strings.ToLower(orgname) + ".example.com",
	}
	err := sdkconnector.RegisterandEnroll(orgSetup.sdk, orgSetup.orgName, Org2User)
	if err != nil {
		fmt.Println("error on registering and enrolling org2user user for Org2 : ", err)
		return
	}
	fmt.Println("Enrolled normal user on Org2")
}
