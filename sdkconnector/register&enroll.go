package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

//RegisterandEnroll registers and enrolls user using Fabric CA.
func RegisterandEnroll(setup *OrgSetup, r *msp.RegistrationRequest) (int, error) {
	MSPClient, err := msp.New(setup.sdk.Context(), msp.WithOrg(setup.OrgName))
	if err != nil {
		return 0, err
	}
	_, err = MSPClient.GetSigningIdentity(r.Name)
	if err == nil {
		//Already enrolled user
		return 1, nil
	}
	secret, err := MSPClient.Register(r)
	if err != nil {
		return 0, err
	}
	err = MSPClient.Enroll(r.Name,
		msp.WithSecret(secret),
		msp.WithProfile("tls"),
	)
	if err != nil {
		return 0, err
	}
	return 2, nil
}
