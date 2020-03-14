package main

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"gosdk-example/web"
)

func main() {

	org1Setup, err := sdkconnector.Initialize("Org1")
	if err != nil {
		fmt.Println("Error initializing setup for Org1: ", err)
	}
	org2Setup, err := sdkconnector.Initialize("Org2")
	if err != nil {
		fmt.Println("Error initializing setup for Org2: ", err)
	} /*
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
		}*/
	fmt.Println()
	orgSetups := web.OrgSetupArray{}
	orgSetups = append(orgSetups, *org1Setup, *org2Setup)
	web.Serve(orgSetups)
}
