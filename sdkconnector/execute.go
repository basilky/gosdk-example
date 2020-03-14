package sdkconnector

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//Execute the chaincode
func Execute(setup *OrgSetup, userName string, channelName string, chainCodeName string, function string, args []string) (*channel.Response, error) {
	clientContext := setup.sdk.ChannelContext(channelName, fabsdk.WithUser(userName))
	client, err := channel.New(clientContext)
	if err != nil {
		fmt.Println(err, "failed to create new channel client")
		return nil, err
	}
	fmt.Println(args)
	argsbyte := make([][]byte, len(args))
	for i, v := range args {
		argsbyte[i] = []byte(v)
	}
	fmt.Println(argsbyte)
	response, err := client.Execute(channel.Request{ChaincodeID: chainCodeName, Fcn: function, Args: argsbyte, TransientMap: nil})
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	return &response, nil
}
