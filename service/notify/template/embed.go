package notify_template

import _ "embed"

var (
	//go:embed device_online.mjml
	DeviceOnline []byte
	//go:embed device_offline.mjml
	DeviceOffline []byte
)
