package user

import (
	"errors"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type dbUtils struct {
	*gorm.DB
}

func NewDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{DB: db}
}

const Salt = "https://github.com/I-m-Surrounded-by-IoT/"

func GenUserPassword(password string) ([]byte, error) {
	if len(password) < 6 {
		return nil, errors.New("password too short")
	}
	return bcrypt.GenerateFromPassword([]byte(Salt+password), bcrypt.DefaultCost)
}

func CheckPassword(password string, hashedPassword []byte) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(Salt+password)) == nil
}

func CheckUserPassword(u *model.User, password string) bool {
	return CheckPassword(password, u.HashedPassword)
}

func SetUserPassword(u *model.User, password string) error {
	hashed, err := GenUserPassword(password)
	if err != nil {
		return err
	}
	u.HashedPassword = hashed
	return nil
}

func (u *dbUtils) GetUser(id string, fields ...string) (*model.User, error) {
	user := new(model.User)
	err := u.Select(fields).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *dbUtils) GetUserByName(name string, fields ...string) (*model.User, error) {
	user := new(model.User)
	err := u.Select(fields).Where("username = ?", name).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *dbUtils) CreateUser(user *model.User) error {
	return u.Create(user).Error
}

func (u *dbUtils) CheckPassword(id string, password string) bool {
	pwd, err := u.GetUserPassword(id)
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

func (u *dbUtils) GetUserPassword(id string) ([]byte, error) {
	user, err := u.GetUser(id, "hashed_password")
	if err != nil {
		return nil, err
	}
	return user.HashedPassword, nil
}

func (u *dbUtils) SetUserPassword(id string, password string) error {
	b, err := GenUserPassword(password)
	if err != nil {
		return err
	}
	return u.Model(&model.User{}).Where("id = ?", id).Update("hashed_password", b).Error
}

func (u *dbUtils) SetUserName(id, username string) error {
	return u.Model(&model.User{}).Where("id = ?", id).Update("username", username).Error
}
