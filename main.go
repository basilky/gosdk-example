package main

import (
	"fmt"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {

	sdk, err := fabsdk.New(config.FromFile("config.yaml"))
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
		"org1user",
		// Type of identity being registered (e.g. "peer, app, user")
		"user",
		// MaxEnrollments is the number of times the secret can  be reused to enroll.
		// if omitted, this defaults to max_enrollments configured on the server
		10,
		// The identity's affiliation e.g. org1.department1
		"org1.department1",
		// Optional attributes associated with this identity
		nil,
		// CAName is the name of the CA to connect to
		"ca.org1.example.com",
		// Secret is an optional password.  If not specified,
		// a random secret is generated.  In both cases, the secret
		// is returned from registration.
		"",
	}
	s, err := mspClient.Register(a)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("secret is", s)
	err = mspClient.Enroll("org1user",
		mspclient.WithSecret(s),
		mspclient.WithProfile("tls"),
	)
	if err != nil {
		fmt.Println("err", err)
	}
	// The resource management client is responsible for managing channels (create/update channel)
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser("admin"), fabsdk.WithOrg("Org1"))
	if err != nil {
		fmt.Println("failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		fmt.Println("failed to create channel management client from Admin identity")
	}
	admin := resMgmtClient
	fmt.Println("Ressource management client created")
	adminIdentity, err := mspClient.GetSigningIdentity("admin")
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
}
