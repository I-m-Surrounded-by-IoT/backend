package model

import (
	"errors"
	"regexp"
)

var (
	ErrEmptyPassword          = errors.New("empty password")
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

type SetUserPasswordReq struct {
	Password string `json:"password"`
}

func (s *SetUserPasswordReq) Validate() error {
	if s.Password == "" {
		return ErrEmptyPassword
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
		return ErrEmptyUsername
	} else if len(l.Username) > 32 {
		return ErrUsernameTooLong
	} else if !alnumPrintHanReg.MatchString(l.Username) {
		return ErrUsernameHasInvalidChar
	}

	if l.Password == "" {
		return ErrEmptyPassword
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
		return ErrEmptyUsername
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

type SendBindEmailCaptchaReq struct {
	Email string `json:"email" binding:"required"`
}

type SendBindEmailCaptchaResp struct {
	Captcha string `json:"captcha"`
}

type BindEmailReq struct {
	Email   string `json:"email" binding:"required"`
	Captcha string `json:"captcha" binding:"required"`
}
