package main

import (
	"fmt"
	"gosdk-example/sdkconnector"
	"gosdk-example/web"
)

func main() {
	//Initialize setup for Org1
	org1Setup, err := sdkconnector.Initialize("Org1")
	if err != nil {
		fmt.Println("Error initializing setup for Org1: ", err)
	}
	//Initialize setup for Org2
	org2Setup, err := sdkconnector.Initialize("Org2")
	if err != nil {
		fmt.Println("Error initializing setup for Org2: ", err)
	}
	orgSetups := web.OrgSetupArray{}
	orgSetups = append(orgSetups, *org1Setup, *org2Setup)
	web.Serve(orgSetups)
}
