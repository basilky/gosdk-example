package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// Instantiate the chaincode on all peers.
func InstantiateCC(sdk *fabsdk.FabricSDK, orgname string, username string, channelname string, r resmgmt.InstantiateCCRequest, p []string) error {

	resourceManagerClientContext := sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	var ropts []resmgmt.RequestOption
	for _, value := range p {
		ropts = append(ropts, resmgmt.WithTargetFilter(&urlTargetFilter{url: value}))
	}
	resp, err := resMgmtClient.InstantiateCC(channelname, r, ropts...)
	if err != nil || resp.TransactionID == "" {
		return err
	}
	return nil
}
