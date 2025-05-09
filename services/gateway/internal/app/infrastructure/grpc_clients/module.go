package grpc_clients

import "go.uber.org/fx"

var Module = fx.Module("grpc-clients",
	fx.Provide(
		NewUserClient,
		NewMessageClient,
		NewAuthClient,
	),
)
