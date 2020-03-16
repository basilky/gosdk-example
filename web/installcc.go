package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//InstallCC handles chaincode install API requests.
func (setups OrgSetupArray) InstallCC(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	orgName := r.FormValue("orgname")
	chainCodePath := r.FormValue("path")
	chainCodeName := r.FormValue("name")
	chainCodeVersion := r.FormValue("version")
	peerURL := r.FormValue("peerurl")
	currentSetup := sdkconnector.LoadSetup(orgName, setups)
	if currentSetup == nil {
		http.Error(w, "Organization '"+orgName+"' does not exist!", 500)
	}
	err := sdkconnector.InstallCC(currentSetup, chainCodePath, chainCodeName, chainCodeVersion, peerURL)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Successfully installed chaincode in peer '%s'", peerURL)
}
