package sdkconnector

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//Join given organization's peers to channel
func JoinChannel(setup *OrgSetup, orgname string, username string, channelname string) error {
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	err = resMgmtClient.JoinChannel(channelname, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com"))
	if err != nil {
		fmt.Println("err", err)
	}
	return nil
}
