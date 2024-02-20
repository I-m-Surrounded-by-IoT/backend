package model

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
)

var (
	ErrPasswordTooLong        = errors.New("password too long")
	ErrPasswordHasInvalidChar = errors.New("password has invalid char")

	ErrEmptyUserId            = errors.New("empty user id")
	ErrEmptyUsername          = errors.New("empty username")
	ErrUsernameTooLong        = errors.New("username too long")
	ErrUsernameHasInvalidChar = errors.New("username has invalid char")
)

var (
	alnumReg         = regexp.MustCompile(`^[[:alnum:]]+$`)
	alnumPrintReg    = regexp.MustCompile(`^[[:print:][:alnum:]]+$`)
	alnumPrintHanReg = regexp.MustCompile(`^[[:print:][:alnum:]\p{Han}]+$`)
)

type FormatEmptyPasswordError string

func (f FormatEmptyPasswordError) Error() string {
	return fmt.Sprintf("%s password empty", string(f))
}

type SetUserPasswordReq struct {
	Password string `json:"password"`
}

func (s *SetUserPasswordReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(s)
}

func (s *SetUserPasswordReq) Validate() error {
	if s.Password == "" {
		return FormatEmptyPasswordError("user")
	} else if len(s.Password) > 32 {
		return ErrPasswordTooLong
	} else if !alnumPrintReg.MatchString(s.Password) {
		return ErrPasswordHasInvalidChar
	}
	return nil
}

type LoginUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginUserReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(l)
}

func (l *LoginUserReq) Validate() error {
	if l.Username == "" {
		return errors.New("username is empty")
	} else if len(l.Username) > 32 {
		return ErrUsernameTooLong
	} else if !alnumPrintHanReg.MatchString(l.Username) {
		return ErrUsernameHasInvalidChar
	}

	if l.Password == "" {
		return FormatEmptyPasswordError("user")
	} else if len(l.Password) > 32 {
		return ErrPasswordTooLong
	} else if !alnumPrintReg.MatchString(l.Password) {
		return ErrPasswordHasInvalidChar
	}
	return nil
}

type UserInfoResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	// Role      dbModel.Role `json:"role"`
	CreatedAt int64 `json:"createdAt"`
}

type SetUsernameReq struct {
	Username string `json:"username"`
}

func (s *SetUsernameReq) Validate() error {
	if s.Username == "" {
		return errors.New("username is empty")
	} else if len(s.Username) > 32 {
		return ErrUsernameTooLong
	} else if !alnumPrintHanReg.MatchString(s.Username) {
		return ErrUsernameHasInvalidChar
	}
	return nil
}

func (s *SetUsernameReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(s)
}

type UserIDReq struct {
	ID string `json:"id"`
}

func (u *UserIDReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(u)
}

func (u *UserIDReq) Validate() error {
	if len(u.ID) != 32 {
		return errors.New("id is required")
	}
	return nil
}
