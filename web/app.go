package web

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"net/http"
)

type OrgSetupArray []sdkconnector.OrgSetup

func Serve(setups OrgSetupArray) {

	http.HandleFunc("/users", setups.EnrollUser)

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)
}
