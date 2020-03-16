package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//Execute chaincode function
func (setups OrgSetupArray) Execute(w http.ResponseWriter, r *http.Request) {
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
	response, err := sdkconnector.Execute(currentSetup, userName, channelID, chainCodeName, function, args)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintln(w, "Transaction ID : ", response.TransactionID)
	fmt.Fprintln(w, "Execute response : \""+string(response.Payload)+"\"")
}
