package sdkconnector

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//Execute the chaincode function. Require username, chaincode name, function name and function arguments.
func Execute(setup *OrgSetup, userName string, channelName string, chainCodeName string, function string, args []string) (*channel.Response, error) {
	clientContext := setup.sdk.ChannelContext(channelName, fabsdk.WithUser(userName))
	client, err := channel.New(clientContext)
	if err != nil {
		fmt.Println(err, "failed to create new channel client")
		return nil, err
	}
	argsbyte := make([][]byte, len(args))
	for i, v := range args {
		argsbyte[i] = []byte(v)
	}
	response, err := client.Execute(channel.Request{ChaincodeID: chainCodeName, Fcn: function, Args: argsbyte, TransientMap: nil})
	if err != nil {
		return nil, err
	}
	return &response, nil
}
