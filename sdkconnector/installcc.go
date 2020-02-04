package sdkconnector

import (
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func InstallCC(sdk *fabsdk.FabricSDK, orgname string, username string, chaincodepath string, chaincodename string, chaincodeversion string) error {
	resourceManagerClientContext := sdk.Context(fabsdk.WithUser(username), fabsdk.WithOrg(orgname))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	ccPkg, err := packager.NewCCPackage(chaincodepath, os.Getenv("GOPATH"))
	if err != nil {
		return err
	}
	installCCReq := resmgmt.InstallCCRequest{Name: chaincodename, Path: chaincodepath, Version: chaincodeversion, Package: ccPkg}
	_, err = resMgmtClient.InstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return err
	}
	sdk.Close()
	return nil
}
