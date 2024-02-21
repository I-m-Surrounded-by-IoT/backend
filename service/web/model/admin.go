package model

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
)

type CreateUserReq user.CreateUserReq

func (r *CreateUserReq) Validate() error {
	return nil
}

type ListUserReq user.ListUserReq

func (r *ListUserReq) Validate() error {
	return nil
}
