package sdkconnector

import (
	"fmt"
	"strings"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

type OrgSetup struct {
	OrgName   string
	AdminName string
	sdk       *fabsdk.FabricSDK
}

func Initialize(orgName string) (*OrgSetup, error) {
	setup := OrgSetup{}
	setup.OrgName = orgName
	setup.AdminName = orgName + "Admin"
	//Initialize SDK for Org1
	sdk, err := CreateSDKInstance(orgName)
	if err != nil {
		return nil, fmt.Errorf("error creating SDK for %s : %s", orgName, err)
	}
	setup.sdk = sdk

	//Register and enroll admin user on Org1
	admin := &mspclient.RegistrationRequest{
		Name:           setup.AdminName,
		Type:           "admin",
		MaxEnrollments: 10,
		Affiliation:    strings.ToLower(setup.OrgName) + ".department1",
		CAName:         "ca." + strings.ToLower(setup.OrgName) + ".example.com",
	}
	err = RegisterandEnroll(&setup, admin)
	if err != nil {
		return nil, fmt.Errorf("error on registering and enrolling admin user for %s : %s", orgName, err)
	}
	return &setup, nil
}
