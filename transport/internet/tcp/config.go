package tcp

import (
	"github.com/xraypb/xray-core/common"
	"github.com/xraypb/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
