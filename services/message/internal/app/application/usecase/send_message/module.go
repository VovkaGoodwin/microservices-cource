package send_message

import "go.uber.org/fx"

var Module = fx.Module("send_message", fx.Provide(New))
