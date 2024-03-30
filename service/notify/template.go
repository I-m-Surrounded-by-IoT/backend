package notify

import (
	"fmt"
	"time"
)

func formatDeviceOnlineBody(deviceID uint64, timestamp int64) string {
	return fmt.Sprintf("device %d %s at %s", deviceID, "online", time.UnixMilli(timestamp).Format(time.RFC3339))
}

func formatDeviceOfflineBody(deviceID uint64, timestamp int64) string {
	return fmt.Sprintf("device %d %s at %s", deviceID, "offline", time.UnixMilli(timestamp).Format(time.RFC3339))
}
