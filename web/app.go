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
	http.HandleFunc("/install", setups.InstallCC)
	http.HandleFunc("/instantiate", setups.InstantiateCC)
	http.HandleFunc("/upgrade", setups.UpgradeCC)
	http.HandleFunc("/query", setups.Query)
	http.HandleFunc("/execute", setups.Execute)
	fmt.Println("Listening (http://localhost:3000/)...")
	fmt.Println("Now open another terminal and run testAPIs.sh")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
