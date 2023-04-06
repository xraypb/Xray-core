package udp

import (
	"github.com/xraypb/Xray-core/common"
	"github.com/xraypb/Xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
