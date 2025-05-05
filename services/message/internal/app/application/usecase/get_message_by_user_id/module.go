package get_message_by_user_id

import "go.uber.org/fx"

var Module = fx.Module("get_message_by_user_id",
	fx.Provide(
		New,
	))
