package sdkconnector

import (
	"fmt"
	"strings"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type OrgSetup struct {
	orgName   string
	adminName string
	sdk       *fabsdk.FabricSDK
}

func Initialize(orgName string) (*OrgSetup, error) {
	setup := OrgSetup{}
	setup.orgName = orgName
	setup.adminName = orgName + "Admin"
	//Initialize SDK for Org1
	sdk, err := CreateSDKInstance(orgName)
	if err != nil {
		return nil, fmt.Errorf("error creating SDK for %s : %s", orgName, err)
	}
	setup.sdk = sdk

	//Register and enroll admin user on Org1
	admin := &mspclient.RegistrationRequest{
		Name:           setup.adminName,
		Type:           "admin",
		MaxEnrollments: 10,
		Affiliation:    strings.ToLower(setup.orgName) + ".department1",
		CAName:         "ca." + strings.ToLower(setup.orgName) + ".example.com",
	}
	err = RegisterandEnroll(setup, admin)
	if err != nil {
		return nil, fmt.Errorf("error on registering and enrolling admin user for %s : %s", orgName, err)
	}
	return &setup, nil
}
