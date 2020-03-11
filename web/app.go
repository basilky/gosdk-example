package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//OrgSetupArray is an array of setups of organizations.
type OrgSetupArray []sdkconnector.OrgSetup

//Serve opens the API for http requests.
func Serve(setups OrgSetupArray) {
	http.HandleFunc("/users", setups.EnrollUser)
	http.HandleFunc("/channel", setups.CreateChannel)
	http.HandleFunc("/join", setups.JoinChannel)
	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)
}
