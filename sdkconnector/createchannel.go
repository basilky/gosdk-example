package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	providersmsp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//Create channel. Need organization name,admin username, channel id and channel config path.
func CreateChennel(sdk *fabsdk.FabricSDK, orgname string, username string, channelid string, channelconfigpath string) error {
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	MSPClient, err := msp.New(sdk.Context(), msp.WithOrg(orgname))
	if err != nil {
		return err
	}
	adminIdentity, err := MSPClient.GetSigningIdentity(username)
	if err != nil {
		return err
	}
	req := resmgmt.SaveChannelRequest{ChannelID: channelid, ChannelConfigPath: channelconfigpath, SigningIdentities: []providersmsp.SigningIdentity{adminIdentity}}
	txID, err := resMgmtClient.SaveChannel(req, resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil || txID.TransactionID == "" {
		return err
	} else {
		return nil
	}
}
