package user

import "encoding"

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
var _ encoding.TextUnmarshaler = (*Role)(nil)

func (u Role) MarshalBinary() ([]byte, error) {
	return []byte{byte(u)}, nil
}

func (u *Role) UnmarshalText(data []byte) error {
	*u = Role(data[0])
	return nil
}

var _ encoding.BinaryMarshaler = (*Status)(nil)
var _ encoding.TextUnmarshaler = (*Status)(nil)

func (u Status) MarshalBinary() ([]byte, error) {
	return []byte{byte(u)}, nil
}

func (u *Status) UnmarshalText(data []byte) error {
	*u = Status(data[0])
	return nil
}
