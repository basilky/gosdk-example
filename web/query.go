package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//Query handles query chaincode API requests.
func (setups OrgSetupArray) Query(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	orgName := r.FormValue("orgname")
	userName := r.FormValue("username")
	chainCodeName := r.FormValue("chaincodeid")
	channelID := r.FormValue("channelid")
	function := r.FormValue("function")
	args := r.Form["args"]
	currentSetup := sdkconnector.LoadSetup(orgName, setups)
	if currentSetup == nil {
		http.Error(w, "Organization '"+orgName+"' does not exist!", 500)
	}
	response, err := sdkconnector.Query(currentSetup, userName, channelID, chainCodeName, function, args)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintln(w, "Response : '"+string(response.Payload)+"'")
}
