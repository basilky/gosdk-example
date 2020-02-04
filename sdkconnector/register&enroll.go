package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func ResgisterandEnroll(sdk *fabsdk.FabricSDK, orgname string, r *msp.RegistrationRequest) error {
	MSPClient, err := msp.New(sdk.Context(), msp.WithOrg(orgname))
	if err != nil {
		return err
	}
	secret, err := MSPClient.Register(r)
	if err != nil {
		return err
	}
	err = MSPClient.Enroll(r.Name,
		msp.WithSecret(secret),
		msp.WithProfile("tls"),
	)
	if err != nil {
		return err
	}
	sdk.Close()
	return nil
}
