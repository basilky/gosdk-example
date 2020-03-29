package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//InstantiateCC handles chaincode instantiate API requests.
func (setups OrgSetupArray) InstantiateCC(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	orgName := r.FormValue("orgname")
	channelName := r.FormValue("channelid")
	chainCodePath := r.FormValue("path")
	chainCodeName := r.FormValue("name")
	chainCodeVersion := r.FormValue("version")
	peers := r.Form["peer"]
	policyString := r.FormValue("policy")
	currentSetup := sdkconnector.LoadSetup(orgName, setups)
	if currentSetup == nil {
		http.Error(w, "Organization '"+orgName+"' does not exist!", 500)
	}
	err := sdkconnector.InstantiateCC(currentSetup, channelName, chainCodeName, chainCodePath, chainCodeVersion, peers, policyString)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Chaincode instantiation successfull")
}
