package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func ResgisterandEnroll(orgname string, r *msp.RegistrationRequest) error {
	sdk, err := CreateSDKInstance(orgname)
	if err != nil {
		return err
	}
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
