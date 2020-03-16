package sdkconnector

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//Query the chaincode.
func Query(setup *OrgSetup, userName string, channelName string, chainCodeName string, function string, args []string) (*channel.Response, error) {
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
	response, err := client.Query(channel.Request{ChaincodeID: chainCodeName, Fcn: function, Args: argsbyte})
	if err != nil {
		return nil, err
	}
	return &response, nil
}
