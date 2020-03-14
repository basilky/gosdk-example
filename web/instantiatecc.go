package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//InstantiateCC instantiates chaincode on given peers.
func (setups OrgSetupArray) InstantiateCC(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	orgName := r.FormValue("orgname")
	channelName := r.FormValue("channelid")
	chainCodePath := r.FormValue("path")
	chainCodeName := r.FormValue("name")
	chainCodeVersion := r.FormValue("version")
	peers := r.Form["peer"]
	currentSetup := sdkconnector.LoadSetup(orgName, setups)
	if currentSetup == nil {
		http.Error(w, "Organization '"+orgName+"' does not exist!", 500)
	}
	err := sdkconnector.InstantiateCC(currentSetup, channelName, chainCodeName, chainCodePath, chainCodeVersion, peers)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Chaincode instantiation successfull")
}
