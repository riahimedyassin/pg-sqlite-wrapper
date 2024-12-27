package pglitewrapper

type IDatabaseWrapper interface {
	Reconnect(reconnectConfig ReconnectConfig) error
}
