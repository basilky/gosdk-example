package main

import (
	"fmt"
	"gosdk-example/sdkconnector"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func main() {
	//Register and enroll admn user on Org1
	Org1Admin := &mspclient.RegistrationRequest{
		Name:           "org1admin",
		Type:           "admin",
		MaxEnrollments: 10,
		Affiliation:    "org1.department1",
		CAName:         "ca.org1.example.com",
	}
	err := sdkconnector.ResgisterandEnroll("Org1", Org1Admin)
	if err != nil {
		fmt.Println("error on registering and enrolling admin user for Org1 : ", err)
	}
	sdkconnector.CreateChennel("Org1", "org1admin", "mychannel", "network/channel-artifacts/channel.tx")
	if err != nil {
		fmt.Println("error creating channel : ", err)
	}
	err = sdkconnector.JoinChennel("Org1", "org1admin", "mychannel")
	if err != nil {
		fmt.Println("error joining Org1 peers to channel : ", err)
	}
	//Register and enroll admn user on Org2
	Org2Admin := &mspclient.RegistrationRequest{
		Name:           "org2admin",
		Type:           "admin",
		MaxEnrollments: 10,
		Affiliation:    "org2.department2",
		CAName:         "ca.org2.example.com",
	}
	err = sdkconnector.ResgisterandEnroll("Org2", Org2Admin)
	if err != nil {
		fmt.Println("error on registering and enrolling admin user for Org2 : ", err)
	}
	sdkconnector.CreateChennel("Org2", "org2admin", "mychannel", "network/channel-artifacts/channel.tx")
	if err != nil {
		fmt.Println("error creating channel : ", err)
	}
	err = sdkconnector.JoinChennel("Org2", "org2admin", "mychannel")
	if err != nil {
		fmt.Println("error joining Org2 peers to channel : ", err)
	}
	/*
		err = admin.JoinChannel("mychannel", resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com"))
		if err != nil {
			fmt.Println("err", err)
		} else {
			fmt.Println("Channel joined")
		}

		///////////////////////////////////////////////////////////
		sdk2, err := fabsdk.New(config.FromFile("configs/org2config.yaml"))
		if err != nil {
			fmt.Println(err)
		}
		mspClient2, err := mspclient.New(
			sdk2.Context(),
			mspclient.WithOrg("Org2"),
		)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(mspClient)
		a2 := &mspclient.RegistrationRequest{
			// Name is the unique name of the identity
			Name: "org2user",
			// Type of identity being registered (e.g. "peer, app, user")
			Type: "admin",
			// MaxEnrollments is the number of times the secret can  be reused to enroll.
			// if omitted, this defaults to max_enrollments configured on the server
			MaxEnrollments: 10,
			// The identity's affiliation e.g. org1.department1
			Affiliation: "org2.department1",
			// Optional attributes associated with this identity
			Attributes: nil,
			// CAName is the name of the CA to connect to
			CAName: "ca.org2.example.com",
			// Secret is an optional password.  If not specified,
			// a random secret is generated.  In both cases, the secret
			// is returned from registration.
			Secret: "",
		}
		s2, err := mspClient2.Register(a2)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("secret is", s2)
		err = mspClient2.Enroll("admin",
			mspclient.WithSecret("adminpw"),
			mspclient.WithProfile("tls"),
		)
		if err != nil {
			fmt.Println("err", err)
		}
		err = mspClient2.Enroll("org2user",
			mspclient.WithSecret(s2),
			mspclient.WithProfile("tls"),
		)
		if err != nil {
			fmt.Println("err", err)
		}
		// The resource management client is responsible for managing channels (create/update channel)
		resourceManagerClientContext2 := sdk2.Context(fabsdk.WithUser("org2user"), fabsdk.WithOrg("Org2"))
		if err != nil {
			fmt.Println("failed to load Admin identity")
		}
		resMgmtClient2, err := resmgmt.New(resourceManagerClientContext2)
		if err != nil {
			fmt.Println("failed to create channel management client from Admin identity")
		}
		admin2 := resMgmtClient2

		err = admin2.JoinChannel("mychannel", resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com"))
		if err != nil {
			fmt.Println("err", err)
		} else {
			fmt.Println("Channel joined")
		}
		// Create the chaincode package that will be sent to the peers
		ccPkg, err := packager.NewCCPackage("gosdk-example/chaincode/golang", os.Getenv("GOPATH"))
		if err != nil {
			fmt.Println("failed to create chaincode package", err)
		} else {
			fmt.Println("ccpkg works")
		}
		// Install example cc to org peers
		installCCReq := resmgmt.InstallCCRequest{Name: "mycc", Path: "gosdk-example/chaincode/golang", Version: "v0", Package: ccPkg}
		_, err = admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
		if err != nil {
			fmt.Println(err, "failed to install chaincode by admin")
		}
		_, err = admin2.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
		if err != nil {
			fmt.Println(err, "failed to install chaincode by admin2")
		}
		//fmt.Println("Chaincode installed")
		// Set up chaincode policy
		ccPolicy := cauthdsl.SignedByNOutOfGivenRole(2, mspproto.MSPRole_MEMBER, []string{"org1.example.com", "org2.example.com"})

		resp, err := admin2.InstantiateCC("mychannel", resmgmt.InstantiateCCRequest{Name: "mycc", Path: "chaincode/golang", Version: "v0", Args: [][]byte{[]byte("init")}, Policy: ccPolicy})
		if err != nil || resp.TransactionID == "" {
			fmt.Println(err, "failed to instantiate the chaincode")
		}
		fmt.Println("Chaincode instantiated")

		fmt.Println("Chaincode Installation & Instantiation Successful")
		n := &mspclient.RegistrationRequest{
			// Name is the unique name of the identity
			Name: "org2normal",
			// Type of identity being registered (e.g. "peer, app, user")
			Type: "client",
			// MaxEnrollments is the number of times the secret can  be reused to enroll.
			// if omitted, this defaults to max_enrollments configured on the server
			MaxEnrollments: 10,
			// The identity's affiliation e.g. org1.department1
			Affiliation: "org2.department1",
			// Optional attributes associated with this identity
			Attributes: nil,
			// CAName is the name of the CA to connect to
			CAName: "ca.org2.example.com",
			// Secret is an optional password.  If not specified,
			// a random secret is generated.  In both cases, the secret
			// is returned from registration.
			Secret: "",
		}
		s3, err := mspClient2.Register(n)
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Println("secret is", s3)
		err = mspClient2.Enroll("org2normal",
			mspclient.WithSecret(s3),
			mspclient.WithProfile("tls"),
		)
		if err != nil {
			fmt.Println("err", err)
		}
		// Channel client is used to query and execute transactions
		clientContext := sdk2.ChannelContext("mychannel", fabsdk.WithUser("org2normal"))
		client, err := channel.New(clientContext)
		if err != nil {
			fmt.Println(err, "failed to create new channel client")
		}
		fmt.Println("Channel client created")
		_, err = event.New(clientContext)
		if err != nil {
			fmt.Println(err, "failed to create new event client")
		}
		fmt.Println("Event client created")
		transientDataMap := make(map[string][]byte)
		transientDataMap["result"] = []byte("Transient data in hello invoke")
		res, err := client.Execute(channel.Request{ChaincodeID: "mycc", Fcn: "initLedger", Args: nil, TransientMap: transientDataMap}, channel.WithTargetEndpoints("peer0.org1.example.com", "peer0.org2.example.com"))
		fmt.Println(err, res)
		response, err := client.Query(channel.Request{ChaincodeID: "mycc", Fcn: "queryAllCars", Args: [][]byte{}}, channel.WithTargetEndpoints("peer1.org1.example.com"))
		fmt.Println(response, err)
		sdk2.Close()*/
}
