package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/xraypb/xray-core/app/dispatcher"
	_ "github.com/xraypb/xray-core/app/proxyman/inbound"
	_ "github.com/xraypb/xray-core/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/xraypb/xray-core/app/commander"
	_ "github.com/xraypb/xray-core/app/log/command"
	_ "github.com/xraypb/xray-core/app/proxyman/command"
	_ "github.com/xraypb/xray-core/app/stats/command"

	// Developer preview services
	_ "github.com/xraypb/xray-core/app/observatory/command"

	// Other optional features.
	_ "github.com/xraypb/xray-core/app/dns"
	_ "github.com/xraypb/xray-core/app/dns/fakedns"
	_ "github.com/xraypb/xray-core/app/log"
	_ "github.com/xraypb/xray-core/app/metrics"
	_ "github.com/xraypb/xray-core/app/policy"
	_ "github.com/xraypb/xray-core/app/reverse"
	_ "github.com/xraypb/xray-core/app/router"
	_ "github.com/xraypb/xray-core/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/xraypb/xray-core/transport/internet/tagged/taggedimpl"

	// Developer preview features
	_ "github.com/xraypb/xray-core/app/observatory"

	// Inbound and outbound proxies.
	_ "github.com/xraypb/xray-core/proxy/blackhole"
	_ "github.com/xraypb/xray-core/proxy/dns"
	_ "github.com/xraypb/xray-core/proxy/dokodemo"
	_ "github.com/xraypb/xray-core/proxy/freedom"
	_ "github.com/xraypb/xray-core/proxy/http"
	_ "github.com/xraypb/xray-core/proxy/loopback"
	_ "github.com/xraypb/xray-core/proxy/mtproto"
	_ "github.com/xraypb/xray-core/proxy/shadowsocks"
	_ "github.com/xraypb/xray-core/proxy/socks"
	_ "github.com/xraypb/xray-core/proxy/trojan"
	_ "github.com/xraypb/xray-core/proxy/vless/inbound"
	_ "github.com/xraypb/xray-core/proxy/vless/outbound"
	_ "github.com/xraypb/xray-core/proxy/vmess/inbound"
	_ "github.com/xraypb/xray-core/proxy/vmess/outbound"
	_ "github.com/xraypb/xray-core/proxy/wireguard"

	// Transports
	_ "github.com/xraypb/xray-core/transport/internet/domainsocket"
	_ "github.com/xraypb/xray-core/transport/internet/grpc"
	_ "github.com/xraypb/xray-core/transport/internet/http"
	_ "github.com/xraypb/xray-core/transport/internet/kcp"
	_ "github.com/xraypb/xray-core/transport/internet/quic"
	_ "github.com/xraypb/xray-core/transport/internet/reality"
	_ "github.com/xraypb/xray-core/transport/internet/tcp"
	_ "github.com/xraypb/xray-core/transport/internet/tls"
	_ "github.com/xraypb/xray-core/transport/internet/udp"
	_ "github.com/xraypb/xray-core/transport/internet/websocket"

	// Transport headers
	_ "github.com/xraypb/xray-core/transport/internet/headers/http"
	_ "github.com/xraypb/xray-core/transport/internet/headers/noop"
	_ "github.com/xraypb/xray-core/transport/internet/headers/srtp"
	_ "github.com/xraypb/xray-core/transport/internet/headers/tls"
	_ "github.com/xraypb/xray-core/transport/internet/headers/utp"
	_ "github.com/xraypb/xray-core/transport/internet/headers/wechat"
	_ "github.com/xraypb/xray-core/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/xraypb/xray-core/main/json"
	_ "github.com/xraypb/xray-core/main/toml"
	_ "github.com/xraypb/xray-core/main/yaml"

	// Load config from file or http(s)
	_ "github.com/xraypb/xray-core/main/confloader/external"

	// Commands
	_ "github.com/xraypb/xray-core/main/commands/all"
)
