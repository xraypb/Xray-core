package udp

import (
	"github.com/xraypb/Xray-core/common/buf"
	"github.com/xraypb/Xray-core/common/net"
)

// Packet is a UDP packet together with its source and destination address.
type Packet struct {
	Payload *buf.Buffer
	Source  net.Destination
	Target  net.Destination
}
