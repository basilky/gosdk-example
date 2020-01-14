package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func main() {

	sdk, err := fabsdk.New(config.FromFile("config.yaml"))
	if err != nil {
		fmt.Println(err)
	}
	mspClient, err := mspclient.New(
		sdk.Context(),
		mspclient.WithOrg("Org1"),
	)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(mspClient)
	a := &msp.RegistrationRequest{
		// Name is the unique name of the identity
		"Admin5",
		// Type of identity being registered (e.g. "peer, app, user")
		"user",
		// MaxEnrollments is the number of times the secret can  be reused to enroll.
		// if omitted, this defaults to max_enrollments configured on the server
		10,
		// The identity's affiliation e.g. org1.department1
		"org1.department1",
		// Optional attributes associated with this identity
		nil,
		// CAName is the name of the CA to connect to
		"ca.org1.example.com",
		// Secret is an optional password.  If not specified,
		// a random secret is generated.  In both cases, the secret
		// is returned from registration.
		"",
	}
	s, err := mspClient.Register(a)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("secret is", s)
	err = mspClient.Enroll("Admin5",
		msp.WithSecret(s),
		msp.WithProfile("tls"),
	)
	if err != nil {
		fmt.Println("err", err)
	}
}
