package sdkconnector

import (
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//CreateSDKInstance creates SDK instance for given organization.
func CreateSDKInstance(orgName string) (*fabsdk.FabricSDK, error) {
	configpath := filepath.Join("configs", strings.ToLower(orgName)+"config.yaml")
	config := config.FromFile(configpath)
	sdk, err := fabsdk.New(config)
	return sdk, err
}
