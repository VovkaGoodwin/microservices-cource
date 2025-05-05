package grpc_server

import (
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/get_message_by_user_id"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/ping"
	"github.com/VovkaGoodwin/microservices-cource/services/message/internal/app/application/usecase/send_message"
	"go.uber.org/fx"
)

var Module = fx.Module("grpc-server", fx.Provide(
	func(
		gmuc get_message_by_user_id.Usecase,
		puc ping.Usecase,
		smuc send_message.Usecase,
	) *Deps {
		return &Deps{
			gmuc: gmuc,
			puc:  puc,
			smuc: smuc,
		}
	},
	NewGrpcHandler,
	NewGrpcServer,
))
