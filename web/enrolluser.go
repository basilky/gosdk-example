package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
	"strings"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

//EnrollUser registers and enrolls user to the organization.
//username and orgname are received through the request.
func (setups OrgSetupArray) EnrollUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	username := r.FormValue("username")
	orgname := r.FormValue("orgname")
	currentsetup := sdkconnector.LoadSetup(orgname, setups)
	if currentsetup == nil {
		http.Error(w, "Organization '"+orgname+"' does not exist!", 500)
	}
	User := &mspclient.RegistrationRequest{
		Name:           username,
		Type:           "client",
		MaxEnrollments: 10,
		Affiliation:    strings.ToLower(orgname) + ".department1",
		CAName:         "ca." + strings.ToLower(orgname) + ".example.com",
	}
	err := sdkconnector.RegisterandEnroll(currentsetup, User)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, "Successfully enrolled user '%s'", username)
}
