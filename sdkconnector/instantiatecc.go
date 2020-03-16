package sdkconnector

import (
	mspproto "github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric/common/cauthdsl"
)

//InstantiateCC instantiates the chaincode on given peers.
func InstantiateCC(setup *OrgSetup, channelName string, chainCodeName string, chainCodePath string, chainCodeVersion string, peers []string) error {
	ccPolicy := cauthdsl.SignedByNOutOfGivenRole(2, mspproto.MSPRole_MEMBER, []string{"Org1MSP", "Org2MSP"})
	instCCRequest := resmgmt.InstantiateCCRequest{Name: chainCodeName, Path: chainCodePath, Version: chainCodeVersion, Args: [][]byte{[]byte("init")}, Policy: ccPolicy}
	resourceManagerClientContext := setup.sdk.Context(fabsdk.WithUser(setup.AdminName), fabsdk.WithOrg(setup.OrgName))
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return err
	}
	var ropts []resmgmt.RequestOption
	for _, value := range peers {
		ropts = append(ropts, resmgmt.WithTargetFilter(&urlTargetFilter{url: value}))
	}
	resp, err := resMgmtClient.InstantiateCC(channelName, instCCRequest, ropts...)
	if err != nil || resp.TransactionID == "" {
		return err
	}
	return nil
}
