package sdkconnector

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	providersmsp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func CreateChennel(orgname string, username string, channelid string, channelconfigpath string) error {
	// The resource management client is responsible for managing channels (create/update channel)
	sdk, err := CreateSDKInstance(orgname)
	if err != nil {
		return err
	}
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	if err != nil {
		fmt.Println("failed to load Admin identity")
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		fmt.Println("failed to create channel management client from Admin identity")
	}
	admin := resMgmtClient
	fmt.Println("Ressource management client created")
	MSPClient, err := msp.New(sdk.Context(), msp.WithOrg(orgname))
	if err != nil {
		return err
	}
	adminIdentity, err := MSPClient.GetSigningIdentity("org1user")
	if err != nil {
		fmt.Println("failed to get admin signing identity")
	}
	req := resmgmt.SaveChannelRequest{ChannelID: "mychannel", ChannelConfigPath: "network/channel-artifacts/channel.tx", SigningIdentities: []providersmsp.SigningIdentity{adminIdentity}}
	txID, err := resMgmtClient.SaveChannel(req, resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil || txID.TransactionID == "" {
		fmt.Println("failed to save channel", err)
	} else {
		fmt.Println("Channel created")
	}
	return nil
}
