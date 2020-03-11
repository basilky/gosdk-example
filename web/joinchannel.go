package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

//JoinChannel joins given organization's peers to the channel
func (setups OrgSetupArray) JoinChannel(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	orgName := r.FormValue("orgname")
	channelID := r.FormValue("channelid")
	currentSetup := sdkconnector.LoadSetup(orgName, setups)
	if currentSetup == nil {
		http.Error(w, "Organization '"+orgName+"' does not exist!", 500)
	}
	err := sdkconnector.JoinChannel(currentSetup, channelID)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "Successfully joined peers to channel '%s'", channelID)
}
