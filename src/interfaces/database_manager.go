package interfaces

import (
	"github.com/riahimedyassin/pg-sqlite-wrapper/src/config"
)

type DatabaseWrapper interface {
	Reconnect(reconnectConfig config.ReconnectConfig) error
}
