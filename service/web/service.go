package web

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWebServer)
