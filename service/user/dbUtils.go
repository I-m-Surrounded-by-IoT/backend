package user

import (
	"context"
	"errors"
	"hash/crc32"
	"regexp"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user/model"
	"github.com/zijiren233/stream"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dbUtils struct {
	*gorm.DB
}

func NewDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{DB: db}
}

var (
	ErrPasswordTooShort       = errors.New("password too short")
	ErrPasswordTooLong        = errors.New("password too long")
	ErrPasswordHasInvalidChar = errors.New("password has invalid char")

	ErrUsernameTooShort       = errors.New("username too short")
	ErrUsernameTooLong        = errors.New("username too long")
	ErrUsernameHasInvalidChar = errors.New("username has invalid char")
)

var (
	alnumReg         = regexp.MustCompile(`^[[:alnum:]]+$`)
	alnumPrintReg    = regexp.MustCompile(`^[[:print:][:alnum:]]+$`)
	alnumPrintHanReg = regexp.MustCompile(`^[[:print:][:alnum:]\p{Han}]+$`)
)

func GenUserPassword(password string) ([]byte, error) {
	if len(password) < 6 {
		return nil, ErrPasswordTooShort
	}
	if len(password) > 32 {
		return nil, ErrPasswordTooLong
	}
	return bcrypt.GenerateFromPassword(stream.StringToBytes(password), bcrypt.DefaultCost)
}

func GenUserPasswordVersion(hashedPassword []byte) uint32 {
	return crc32.ChecksumIEEE(hashedPassword)
}

func CheckPassword(password string, hashedPassword []byte) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, stream.StringToBytes(password)) == nil
}

func CheckUserPassword(u *model.User, password string) bool {
	return CheckPassword(password, u.HashedPassword)
}

func SetUserName(u *model.User, name string) error {
	if len(name) < 6 {
		return ErrUsernameTooShort
	}
	if len(name) > 32 {
		return ErrUsernameTooLong
	}
	if !alnumPrintHanReg.MatchString(name) {
		return ErrUsernameHasInvalidChar
	}
	u.Username = name
	return nil
}

func SetUserPassword(u *model.User, password string) error {
	if !alnumPrintReg.MatchString(password) {
		return ErrPasswordHasInvalidChar
	}
	hashed, err := GenUserPassword(password)
	if err != nil {
		return err
	}
	u.HashedPassword = hashed
	return nil
}

func (u *dbUtils) GetUser(ctx context.Context, id string, fields ...string) (*model.User, error) {
	user := new(model.User)

	err := u.WithContext(ctx).Select(fields).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *dbUtils) GetUserByName(ctx context.Context, name string, fields ...string) (*model.User, error) {
	user := new(model.User)
	err := u.WithContext(ctx).Select(fields).Where("username = ?", name).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *dbUtils) CreateUser(user *model.User) error {
	return u.Create(user).Error
}

func (u *dbUtils) CheckPassword(ctx context.Context, id string, password string) bool {
	pwd, err := u.GetUserPassword(ctx, id)
	if err != nil {
		return false
	}
	return CheckPassword(password, pwd)
}

func (u *dbUtils) UpdateUser(user *model.User) error {
	return u.Save(user).Error
}

func (u *dbUtils) DeleteUser(id string) error {
	return u.Delete(&model.User{}, id).Error
}

func (u *dbUtils) ListUser(scopes ...func(*gorm.DB) *gorm.DB) ([]*model.User, error) {
	var users []*model.User
	err := u.Scopes(scopes...).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *dbUtils) CountUser(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := u.Scopes(scopes...).Model(&model.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func WithPageAndPageSize(page, pageSize int) func(*gorm.DB) *gorm.DB {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize > 100 {
		pageSize = 100
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func (u *dbUtils) ListUserWithPageAndPageSize(page, pageSize int, scopes ...func(*gorm.DB) *gorm.DB) (int64, []*model.User, error) {
	count, err := u.CountUser(scopes...)
	if err != nil {
		return 0, nil, err
	}
	l, err := u.ListUser(append(scopes, WithPageAndPageSize(page, pageSize))...)
	return count, l, err
}

func (u *dbUtils) SetUserStatus(id string, status user.Status) error {
	return u.Model(&model.User{}).Where("id = ?", id).Update("status", status).Error
}

func (u *dbUtils) SetUserRole(id string, role user.Role) error {
	return u.Model(&model.User{}).Where("id = ?", id).Update("role", role).Error
}

func (u *dbUtils) GetUserPassword(ctx context.Context, id string) ([]byte, error) {
	user, err := u.GetUser(ctx, id, "hashed_password")
	if err != nil {
		return nil, err
	}
	return user.HashedPassword, nil
}

func (u *dbUtils) GetUserPasswordVersion(ctx context.Context, id string) (uint32, error) {
	pwd, err := u.GetUserPassword(ctx, id)
	if err != nil {
		return 0, err
	}
	return GenUserPasswordVersion(pwd), nil
}

func (u *dbUtils) SetUserPassword(id string, password string) error {
	b, err := GenUserPassword(password)
	if err != nil {
		return err
	}
	return u.Model(&model.User{}).Where("id = ?", id).Update("hashed_password", b).Error
}

func (u *dbUtils) SetUserName(id, username string) (string, error) {
	user := model.User{}
	err := u.Model(&user).
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "username"},
			},
		}).
		Where("id = ?", id).
		Update("username", username).Error
	if err != nil {
		return "", err
	}
	return user.Username, nil
}
