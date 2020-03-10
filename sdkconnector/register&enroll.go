package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

//RegisterandEnroll registers and enrolls user using Fabric CA
func RegisterandEnroll(setup *OrgSetup, r *msp.RegistrationRequest) error {
	MSPClient, err := msp.New(setup.sdk.Context(), msp.WithOrg(setup.OrgName))
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
	return nil
}
