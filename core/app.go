package core

import "log/slog"

type App interface {
	DataDir() string
	EncryptionEnv() string

	IsDev() bool
	IsBootstrapped() bool

	Bootstrap() error
	ResetBootstrapState() error

	Logger() *slog.Logger

	DB()
	Cache()
	Store()
	Broker()

	Restart() error

	// ---------------------------------------------------------------
	// App event hooks
	// ---------------------------------------------------------------

	OnBeforeBootstrap()
	OnAfterBootstrap()

	OnBeforeServe()

	OnBeforeApiError()
	OnAfterApiError()

	OnTerminate()
}
