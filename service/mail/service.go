package mail

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewMailService,
	NewMailConsumer,
)
