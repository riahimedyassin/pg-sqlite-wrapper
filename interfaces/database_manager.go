package interfaces

import (
	"github.com/riahimedyassin/pg-sqlite-wrapper/config"
)

type DatabaseWrapper interface {
	Reconnect(reconnectConfig config.ReconnectConfig) error
}
