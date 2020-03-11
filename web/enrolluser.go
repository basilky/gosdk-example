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
	userName := r.FormValue("username")
	orgName := r.FormValue("orgname")
	currentSetup := sdkconnector.LoadSetup(orgName, setups)
	if currentSetup == nil {
		http.Error(w, "Organization '"+orgName+"' does not exist!", 500)
	}
	user := &mspclient.RegistrationRequest{
		Name:           userName,
		Type:           "client",
		MaxEnrollments: 10,
		Affiliation:    strings.ToLower(orgName) + ".department1",
		CAName:         "ca." + strings.ToLower(orgName) + ".example.com",
	}
	status, err := sdkconnector.RegisterandEnroll(currentSetup, user)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	if status == 1 {
		fmt.Fprintf(w, "User '%s' already exists!", userName)
		return
	}
	fmt.Fprintf(w, "Successfully enrolled user '%s'", userName)
}
