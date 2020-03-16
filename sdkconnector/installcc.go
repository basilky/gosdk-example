package sdkconnector

import (
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//InstallCC packages and installs chaincode on the given peer.
func InstallCC(setup *OrgSetup, chaincodePath string, chaincodeName string, chaincodeVersion string, peerURL string) error {
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.AdminName), fabsdk.WithOrg(setup.OrgName))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	ccPkg, err := packager.NewCCPackage(chaincodePath, os.Getenv("GOPATH"))
	if err != nil {
		return err
	}

	installCCReq := resmgmt.InstallCCRequest{Name: chaincodeName, Path: chaincodePath, Version: chaincodeVersion, Package: ccPkg}
	_, err = resMgmtClient.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargetFilter(&urlTargetFilter{url: peerURL}))
	if err != nil {
		return err
	}
	return nil
}
