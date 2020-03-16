package sdkconnector

import (
	"fmt"
	"strings"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//OrgSetup contains organization's settings to interact with the network.
type OrgSetup struct {
	OrgName   string
	AdminName string
	sdk       *fabsdk.FabricSDK
}

//Initialize the setup for given organization.
func Initialize(orgName string) (*OrgSetup, error) {
	setup := OrgSetup{}
	setup.OrgName = orgName
	setup.AdminName = orgName + "Admin"
	sdk, err := CreateSDKInstance(orgName)
	if err != nil {
		return nil, fmt.Errorf("error creating SDK for %s : %s", orgName, err)
	}
	setup.sdk = sdk
	admin := &mspclient.RegistrationRequest{
		Name:           setup.AdminName,
		Type:           "admin",
		MaxEnrollments: 10,
		Affiliation:    strings.ToLower(setup.OrgName) + ".department1",
		CAName:         "ca." + strings.ToLower(setup.OrgName) + ".example.com",
	}
	_, err = RegisterandEnroll(&setup, admin)
	if err != nil {
		return nil, fmt.Errorf("error on registering and enrolling admin user for %s : %s", orgName, err)
	}
	return &setup, nil
}
