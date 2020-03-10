package sdkconnector

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type urlTargetFilter struct {
	url string
}

func (f *urlTargetFilter) Accept(peer fab.Peer) bool {
	return peer.URL() == f.url
}

func LoadSetup(orgname string, setups []OrgSetup) *OrgSetup {
	for _, element := range setups {
		if element.OrgName == orgname {
			return &element
		}
	}
	return nil
}
