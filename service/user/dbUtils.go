package user

import (
	"context"
	"database/sql"
	"errors"
	"hash/crc32"
	"regexp"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
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

func SetUsername(u *model.User, name string) error {
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
	err := u.
		WithContext(ctx).
		Select(fields).
		Where("id = ?", id).
		First(user).
		Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *dbUtils) GetUserByName(ctx context.Context, name string, fields ...string) (*model.User, error) {
	user := new(model.User)
	err := u.
		WithContext(ctx).
		Select(fields).
		Where("username = ?", name).
		First(user).
		Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *dbUtils) CreateUser(ctx context.Context, user *model.User) error {
	return u.
		WithContext(ctx).
		Create(user).
		Error
}

func (u *dbUtils) CheckPassword(ctx context.Context, id string, password string) bool {
	pwd, err := u.GetUserPassword(ctx, id)
	if err != nil {
		return false
	}
	return CheckPassword(password, pwd)
}

func (u *dbUtils) UpdateUser(ctx context.Context, user *model.User) error {
	return u.
		WithContext(ctx).
		Save(user).
		Error
}

func (u *dbUtils) DeleteUser(ctx context.Context, id string) error {
	return u.
		WithContext(ctx).
		Delete(&model.User{}, id).
		Error
}

func (u *dbUtils) ListUser(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.User, error) {
	var users []*model.User
	err := u.
		WithContext(ctx).
		Scopes(scopes...).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *dbUtils) CountUser(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := u.
		WithContext(ctx).
		Scopes(scopes...).
		Model(&model.User{}).
		Count(&count).
		Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *dbUtils) ListUserWithPageAndPageSize(ctx context.Context, page, pageSize int, scopes ...func(*gorm.DB) *gorm.DB) (int64, []*model.User, error) {
	count, err := u.CountUser(ctx, scopes...)
	if err != nil {
		return 0, nil, err
	}
	l, err := u.ListUser(ctx, append(scopes, utils.WithPageAndPageSize(page, pageSize))...)
	return count, l, err
}

func (u *dbUtils) SetUserStatus(ctx context.Context, id string, status user.Status) error {
	return u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Update("status", status).
		Error
}

func (u *dbUtils) SetUserRole(ctx context.Context, id string, role user.Role) error {
	return u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Update("role", role).
		Error
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

func (u *dbUtils) SetUserPassword(ctx context.Context, id string, password string) error {
	user := model.User{}
	err := SetUserPassword(&user, password)
	if err != nil {
		return err
	}
	return u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", id).
		Update("hashed_password", user.HashedPassword).
		Error
}

func (u *dbUtils) SetUsername(ctx context.Context, id, username string) (string, error) {
	user := model.User{}
	err := SetUsername(&user, username)
	if err != nil {
		return "", err
	}
	err = u.
		WithContext(ctx).
		Model(&user).
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "username"},
			},
		}).
		Where("id = ?", id).
		Update("username", user.Username).Error
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func (u *dbUtils) Transaction(fn func(*dbUtils) error) error {
	return u.DB.Transaction(func(db *gorm.DB) error {
		return fn(NewDBUtils(db))
	})
}

func (u *dbUtils) FollowDevice(ctx context.Context, userId string, deviceId uint64) error {
	return u.
		WithContext(ctx).
		Create(&model.FollowDevice{
			UserID:   userId,
			DeviceID: deviceId,
		}).Error
}

func (u *dbUtils) UnfollowDevice(ctx context.Context, userId string, deviceId uint64) error {
	return u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userId).
		Association("FollowDevices").
		Delete(&model.FollowDevice{DeviceID: deviceId})
}

func (u *dbUtils) ListFollowedDeviceIDs(ctx context.Context, userId string, scopes ...utils.Scope) ([]uint64, error) {
	var devices []uint64
	err := u.
		WithContext(ctx).
		Model(&model.FollowDevice{}).
		Where("user_id = ?", userId).
		Scopes(scopes...).
		Pluck("device_id", &devices).
		Error
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (u *dbUtils) ListFollowedUserIDsByDevice(ctx context.Context, deviceId uint64, scopes ...utils.Scope) ([]string, error) {
	// follow_all_device = true or user_id in (select user_id from follow_device where device_id = ?)
	var users []string
	err := u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("follow_all_device = true").
		Pluck("id", &users).
		Error
	if err != nil {
		return nil, err
	}
	var users2 []string
	err = u.
		WithContext(ctx).
		Model(&model.FollowDevice{}).
		Where("device_id = ? AND user_id NOT IN ?", deviceId, users).
		Pluck("user_id", &users2).
		Error
	if err != nil {
		return nil, err
	}
	users = append(users, users2...)
	return users, nil
}

func (u *dbUtils) ListFollowedUserIDAndNotificationMethodByDevice(ctx context.Context, deviceId uint64, scopes ...utils.Scope) (map[string]*user.NotificationMethod, error) {
	var users []*struct {
		ID string
		*user.NotificationMethod
	}
	err := u.
		WithContext(ctx).
		Model(&model.User{}).
		Select("id, email, phone").
		Where("follow_all_device = true OR id IN (SELECT user_id FROM follow_devices WHERE device_id = ?)", deviceId).
		Scan(&users).
		Error
	if err != nil {
		return nil, err
	}
	m := make(map[string]*user.NotificationMethod, len(users))
	for _, u := range users {
		m[u.ID] = u.NotificationMethod
	}
	return m, nil
}

func (u *dbUtils) DelFollowedDevice(ctx context.Context, deviceId uint64) error {
	return u.
		WithContext(ctx).
		Where("device_id = ?", deviceId).
		Delete(&model.FollowDevice{}).Error
}

func (u *dbUtils) FollowAllDevice(ctx context.Context, userId string) error {
	return u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userId).
		Update("follow_all_device", true).
		Error
}

func (u *dbUtils) UnfollowAllDevice(ctx context.Context, userId string) error {
	return u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userId).
		Update("follow_all_device", false).
		Error
}

func (u *dbUtils) HasFollowedAllDevice(ctx context.Context, userId string) (bool, error) {
	var all sql.NullBool
	err := u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userId).
		Select("follow_all_device").
		First(&all).
		Error
	if err != nil {
		return false, err
	}
	return all.Bool, nil
}

func (u *dbUtils) HasFollowedDevice(ctx context.Context, userId string, deviceId uint64) (bool, error) {
	all, err := u.HasFollowedAllDevice(ctx, userId)
	if err != nil {
		return false, err
	}
	if all {
		return true, nil
	}
	var count int64
	err = u.
		WithContext(ctx).
		Model(&model.FollowDevice{}).
		Where("user_id = ? AND device_id = ?", userId, deviceId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (u *dbUtils) BindEmail(ctx context.Context, userId, email string) error {
	return u.
		WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", userId).
		Update("email", email).
		Error
}

func (u *dbUtils) UnbindEmail(ctx context.Context, userId string) error {
	return u.BindEmail(ctx, userId, "")
}
