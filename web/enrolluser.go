package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
	"strings"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func (setups OrgSetupArray) EnrollUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	username := r.FormValue("username")
	orgname := r.FormValue("orgname")
	fmt.Fprintf(w, "Name = %s\n", username)
	fmt.Fprintf(w, "Address = %s\n", orgname)
	currentsetup := sdkconnector.LoadSetup(orgname, setups)
	if currentsetup == nil {
		return
	}
	//Register and enroll normal user on Org2
	Org2User := &mspclient.RegistrationRequest{
		Name:           username,
		Type:           "client",
		MaxEnrollments: 10,
		Affiliation:    strings.ToLower(orgname) + ".department1",
		CAName:         "ca." + strings.ToLower(orgname) + ".example.com",
	}
	err := sdkconnector.RegisterandEnroll(currentsetup, Org2User)
	if err != nil {
		fmt.Println("error on registering and enrolling org2user user for Org2 : ", err)
		return
	}
	fmt.Println("Enrolled normal user on Org2")
}
