package user

import (
	"encoding"

	"github.com/redis/go-redis/v9"
	"github.com/zijiren233/stream"
)

func (u Role) IsAdmin() bool {
	return u == Role_ADMIN
}

func (u Status) IsActice() bool {
	return u == Status_ACTIVE
}

func (u Status) IsInActive() bool {
	return u == Status_INACTIVE
}

var _ encoding.BinaryMarshaler = (*Role)(nil)
var _ redis.Scanner = (*Role)(nil)

func (u Role) MarshalBinary() ([]byte, error) {
	return stream.StringToBytes(u.String()), nil
}

func (u *Role) ScanRedis(s string) error {
	*u = Role(Role_value[s])
	return nil
}

var _ encoding.BinaryMarshaler = (*Status)(nil)
var _ redis.Scanner = (*Status)(nil)

func (u Status) MarshalBinary() ([]byte, error) {
	return stream.StringToBytes(u.String()), nil
}

func (u *Status) ScanRedis(s string) error {
	*u = Status(Status_value[s])
	return nil
}
