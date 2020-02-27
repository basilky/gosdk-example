package sdkconnector

import "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"

func (f *urlTargetFilter) Accept(peer fab.Peer) bool {
	return peer.URL() == f.url
}

type urlTargetFilter struct {
	url string
}
