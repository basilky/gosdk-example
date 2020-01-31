package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	providersmsp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func CreateChennel(orgname string, username string, channelid string, channelconfigpath string) error {
	sdk, err := CreateSDKInstance(orgname)
	if err != nil {
		return err
	}
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	if err != nil {
		return err
	}
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	MSPClient, err := msp.New(sdk.Context(), msp.WithOrg(orgname))
	if err != nil {
		return err
	}
	adminIdentity, err := MSPClient.GetSigningIdentity("org1user")
	if err != nil {
		return err
	}
	req := resmgmt.SaveChannelRequest{ChannelID: "mychannel", ChannelConfigPath: "network/channel-artifacts/channel.tx", SigningIdentities: []providersmsp.SigningIdentity{adminIdentity}}
	txID, err := resMgmtClient.SaveChannel(req, resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil || txID.TransactionID == "" {
		return err
	} else {
		return nil
	}
}
