package main

import (
	"fmt"
	"gosdk-example/sdkconnector"

	mspproto "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/common/cauthdsl"
)

func main() {

	Org1SDK, err := sdkconnector.CreateSDKInstance("Org1")
	if err != nil {
		fmt.Println("error creating sdk for Org1 : ", err)
		return
	}

	Org2SDK, err := sdkconnector.CreateSDKInstance("Org2")
	if err != nil {
		fmt.Println("error creating sdk for Org2 : ", err)
	}

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

	//Register and enroll admn user on Org2
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

	sdkconnector.CreateChennel(Org1SDK, "Org1", "org1admin", "mychannel", "network/channel-artifacts/channel.tx")
	if err != nil {
		fmt.Println("error creating channel : ", err)
		return
	}

	err = sdkconnector.JoinChennel(Org1SDK, "Org1", "org1admin", "mychannel")
	if err != nil {
		fmt.Println("error joining Org1 peers to channel : ", err)
		return
	}

	err = sdkconnector.JoinChennel(Org2SDK, "Org2", "org2admin", "mychannel")
	if err != nil {
		fmt.Println("error joining Org2 peers to channel : ", err)
		return
	}

	err = sdkconnector.InstallCC(Org1SDK, "Org1", "org1admin", "gosdk-example/chaincode/golang", "mycc", "v0")
	if err != nil {
		fmt.Println("erro1")
		return
	}
	err = sdkconnector.InstallCC(Org2SDK, "Org2", "org2admin", "gosdk-example/chaincode/golang", "mycc", "v0")
	if err != nil {
		fmt.Println("erro2")
		return
	}

	// Set up chaincode policy
	ccPolicy := cauthdsl.SignedByNOutOfGivenRole(2, mspproto.MSPRole_MEMBER, []string{"org1.example.com", "org2.example.com"})
	instCCrequest := resmgmt.InstantiateCCRequest{Name: "mycc", Path: "chaincode/golang", Version: "v0", Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	err = sdkconnector.InstantiateCC(Org2SDK, "Org2", "org2admin", "mychannel", instCCrequest)
	if err != nil {
		fmt.Println(err, "failed to instantiate the chaincode")
		return
	}
	fmt.Println("Chaincode instantiated")

	fmt.Println("Chaincode Installation & Instantiation Successful")
	//Register and enroll admn user on Org2
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

	// Channel client is used to query and execute transactions
	clientContext := Org2SDK.ChannelContext("mychannel", fabsdk.WithUser("org2user"))
	client, err := channel.New(clientContext)
	if err != nil {
		fmt.Println(err, "failed to create new channel client")
		return
	}

	res, err := client.Execute(channel.Request{ChaincodeID: "mycc", Fcn: "initLedger", Args: nil, TransientMap: nil}, channel.WithTargetEndpoints("peer0.org1.example.com", "peer0.org2.example.com"))
	if err != nil {
		fmt.Println("Transaction success, ID : ", res.TransactionID)
	} else {
		fmt.Println("Error execute transaction : ", err)
		return
	}
	response, err := client.Query(channel.Request{ChaincodeID: "mycc", Fcn: "queryAllCars", Args: [][]byte{}}, channel.WithTargetEndpoints("peer1.org1.example.com"))
	if err != nil {
		fmt.Println("Query Response : ", string(response.Payload))
	} else {
		fmt.Println("Error : ", err)
		return
	}
}
