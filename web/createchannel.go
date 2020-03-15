package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//CreateChannel creates channel based on the channel ID and
//channel config path received through the request.
func (setups OrgSetupArray) CreateChannel(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	orgName := r.FormValue("orgname")
	channelID := r.FormValue("channelid")
	channelConfigPath := r.FormValue("configpath")
	currentSetup := sdkconnector.LoadSetup(orgName, setups)
	if currentSetup == nil {
		http.Error(w, "Organization '"+orgName+"' does not exist!", 500)
	}
	err := sdkconnector.CreateChannel(currentSetup, channelID, channelConfigPath)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Successfully created channel '%s'", channelID)
}
