package udp

import (
	"github.com/xraypb/xray-core/common"
	"github.com/xraypb/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
