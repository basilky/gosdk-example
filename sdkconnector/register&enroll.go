package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func ResgisterandEnroll(org string, r *msp.RegistrationRequest) error {
	sdk, err := CreateSDKInstance(org)
	if err != nil {
		return err
	}
	MSPClient, err := msp.New(sdk.Context(), msp.WithOrg(org))
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
