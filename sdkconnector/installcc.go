package sdkconnector

import (
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//InstallCC installs chaincode on given organization's peers.
func InstallCC(setup *OrgSetup, chaincodepath string, chaincodename string, chaincodeversion string, purl string) error {
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.AdminName), fabsdk.WithOrg(setup.OrgName))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	ccPkg, err := packager.NewCCPackage(chaincodepath, os.Getenv("GOPATH"))
	if err != nil {
		return err
	}

	installCCReq := resmgmt.InstallCCRequest{Name: chaincodename, Path: chaincodepath, Version: chaincodeversion, Package: ccPkg}
	_, err = resMgmtClient.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithTargetFilter(&urlTargetFilter{url: purl}))
	if err != nil {
		return err
	}
	return nil
}
