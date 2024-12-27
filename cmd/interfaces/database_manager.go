package interfaces

import (
	"astor-bd-wrapper/cmd/config"
)

type DatabaseWrapper interface {
	Reconnect(reconnectConfig config.ReconnectConfig) error
}
