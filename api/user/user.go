package user

func (u Role) IsAdmin() bool {
	return u == Role_ADMIN
}

func (u Status) IsActice() bool {
	return u == Status_ACTIVE
}

func (u Status) IsInActive() bool {
	return u == Status_INACTIVE
}
