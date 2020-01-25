package main

import (
	"fmt"
	"os"

	mspproto "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/common/cauthdsl"
)

func main() {

	sdk, err := fabsdk.New(config.FromFile("org1config.yaml"))
	if err != nil {
		fmt.Println(err)
	}
	mspClient, err := mspclient.New(
		sdk.Context(),
		mspclient.WithOrg("Org1"),
	)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(mspClient)
	a := &mspclient.RegistrationRequest{
		// Name is the unique name of the identity
		Name: "org1user",
		// Type of identity being registered (e.g. "peer, app, user")
		Type: "admin",
		// MaxEnrollments is the number of times the secret can  be reused to enroll.
		// if omitted, this defaults to max_enrollments configured on the server
		MaxEnrollments: 10,
		// The identity's affiliation e.g. org1.department1
		Affiliation: "org1.department1",
		// Optional attributes associated with this identity
		Attributes: nil,
		// CAName is the name of the CA to connect to
		CAName: "ca.org1.example.com",
		// Secret is an optional password.  If not specified,
		// a random secret is generated.  In both cases, the secret
		// is returned from registration.
		Secret: "",
	}
	s, err := mspClient.Register(a)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("secret is", s)
	err = mspClient.Enroll("admin",
		mspclient.WithSecret("adminpw"),
		mspclient.WithProfile("tls"),
	)
	if err != nil {
		fmt.Println("err", err)
	}
	err = mspClient.Enroll("org1user",
		mspclient.WithSecret(s),
		mspclient.WithProfile("tls"),
	)
	if err != nil {
		fmt.Println("err", err)
	}
	// The resource management client is responsible for managing channels (create/update channel)
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser("org1user"), fabsdk.WithOrg("Org1"))
	if err != nil {
		fmt.Println("failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		fmt.Println("failed to create channel management client from Admin identity")
	}
	admin := resMgmtClient
	fmt.Println("Ressource management client created")
	adminIdentity, err := mspClient.GetSigningIdentity("org1user")
	if err != nil {
		fmt.Println("failed to get admin signing identity")
	}
	req := resmgmt.SaveChannelRequest{ChannelID: "mychannel", ChannelConfigPath: "first-network/channel-artifacts/channel.tx", SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	txID, err := admin.SaveChannel(req, resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil || txID.TransactionID == "" {
		fmt.Println("failed to save channel", err)
	} else {
		fmt.Println("Channel created")
	}
	err = admin.JoinChannel("mychannel", resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("Channel joined")
	}
	sdk.Close()

	///////////////////////////////////////////////////////////
	sdk2, err := fabsdk.New(config.FromFile("org2config.yaml"))
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
	fmt.Println("Ressource management client created")
	adminIdentity2, err := mspClient.GetSigningIdentity("org2user")
	if err != nil {
		fmt.Println("failed to get admin signing identity")
	}
	req2 := resmgmt.SaveChannelRequest{ChannelID: "mychannel", ChannelConfigPath: "first-network/channel-artifacts/channel.tx", SigningIdentities: []msp.SigningIdentity{adminIdentity2}}
	txID2, err := admin2.SaveChannel(req2, resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil || txID2.TransactionID == "" {
		fmt.Println("failed to save channel", err)
	} else {
		fmt.Println("Channel created")
	}
	err = admin2.JoinChannel("mychannel", resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("Channel joined")
	}
	// Create the chaincode package that will be sent to the peers
	ccPkg, err := packager.NewCCPackage("go-sdk-demo/chaincode/golang", os.Getenv("GOPATH"))
	if err != nil {
		fmt.Println("failed to create chaincode package", err)
	} else {
		fmt.Println("ccpkg works")
	}
	// Install example cc to org peers
	installCCReq := resmgmt.InstallCCRequest{Name: "mycc", Path: "go-sdk-demo/chaincode/golang", Version: "v0", Package: ccPkg}
	_, err = admin.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		fmt.Println(err, "failed to install chaincode")
	}
	installCCReq = resmgmt.InstallCCRequest{Name: "mycc", Path: "go-sdk-demo/chaincode/golang", Version: "v0", Package: ccPkg}
	_, err = admin2.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		fmt.Println(err, "failed to install chaincode")
	}
	fmt.Println("Chaincode installed")
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
	_, err = channel.New(clientContext)
	if err != nil {
		fmt.Println(err, "failed to create new channel client")
	}
	fmt.Println("Channel client created")
	sdk2.Close()
}
