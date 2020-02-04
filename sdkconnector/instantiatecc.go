package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func InstantiateCC(orgname string, username string, channelname string, r resmgmt.InstantiateCCRequest) error {

	sdk, err := CreateSDKInstance(orgname)
	if err != nil {
		return err
	}
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	resp, err := resMgmtClient.InstantiateCC(channelname, r)
	if err != nil || resp.TransactionID == "" {
		return err
	}
	sdk.Close()
	return nil
}
