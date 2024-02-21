package model

import (
	"errors"
	"fmt"
	"regexp"
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

type UserIDReq struct {
	ID string `json:"id"`
}

func (u *UserIDReq) Validate() error {
	if len(u.ID) != 32 {
		return errors.New("id is required")
	}
	return nil
}
