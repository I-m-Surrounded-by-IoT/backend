package user

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	redsync "github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct {
	urcache *UserRcache
	db      *dbUtils
	user.UnimplementedUserServer
}

func NewUserService(dc *conf.DatabaseServerConfig, uc *conf.UserConfig, rc *conf.RedisConfig) *UserService {
	d, err := dbdial.Dial(context.Background(), dc)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if dc.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.User),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Username: rc.Username,
		Password: rc.Password,
		DB:       int(rc.Db),
	})
	db := NewDBUtils(d)
	rsync := redsync.New(goredis.NewPool(rdb))
	return &UserService{
		urcache: NewUserRcache(rcache.NewRcacheWithRsync(rdb, rsync), db),
		db:      db,
	}
}

func user2Proto(u *model.User) *user.UserInfo {
	return &user.UserInfo{
		Id:        u.ID,
		CreatedAt: u.CreatedAt.UnixMicro(),
		UpdatedAt: u.UpdatedAt.UnixMicro(),
		Name:      u.Username,
		Role:      u.Role,
		Status:    u.Status,
	}
}

func users2Proto(us []*model.User) []*user.UserInfo {
	res := make([]*user.UserInfo, 0, len(us))
	for _, u := range us {
		res = append(res, user2Proto(u))
	}
	return res
}

func (us *UserService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.UserInfo, error) {
	u := &model.User{
		Role:   req.Role,
		Status: req.Status,
	}
	err := SetUserName(u, req.Name)
	if err != nil {
		return nil, err
	}
	err = SetUserPassword(u, req.Password)
	if err != nil {
		return nil, err
	}
	err = us.db.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return user2Proto(u), nil
}

func (us *UserService) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (*user.UserInfo, error) {
	return us.urcache.GetUserInfo(ctx, req.Id, req.Fields...)
}

func (us *UserService) GetUserInfoByName(ctx context.Context, req *user.GetUserInfoByNameReq) (*user.UserInfo, error) {
	return us.urcache.GetUserInfoByName(ctx, req.Name, req.Fields...)
}

func (us *UserService) ListUser(ctx context.Context, req *user.ListUserReq) (*user.ListUserResp, error) {
	opts := []func(*gorm.DB) *gorm.DB{}
	if req.Id != nil {
		opts = append(opts, model.WithIDEq(*req.Id))
	}
	if req.Name != nil {
		opts = append(opts, model.WithNameLike(*req.Name))
	}
	if req.Role != nil {
		opts = append(opts, model.WithRoleEq(*req.Role))
	}
	if req.Status != nil {
		opts = append(opts, model.WithStatusEq(*req.Status))
	}
	count, err := us.db.CountUser(opts...)
	if err != nil {
		return nil, err
	}
	opts = append(opts, WithPageAndPageSize(int(req.Page), int(req.Size)))
	u, err := us.db.ListUser(opts...)
	if err != nil {
		return nil, err
	}
	return &user.ListUserResp{
		Users: users2Proto(u),
		Total: int32(count),
	}, nil
}

func (us *UserService) GetUserId(ctx context.Context, req *user.GetUserIdReq) (*user.GetUserIdResp, error) {
	id, err := us.urcache.GetUserID(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &user.GetUserIdResp{
		Id: id,
	}, nil
}

func (us *UserService) SetUserName(ctx context.Context, req *user.SetUserNameReq) (*user.SetUserNameResp, error) {
	s, err := us.urcache.SetUserName(ctx, req.Id, req.Name)
	if err != nil {
		return nil, err
	}
	return &user.SetUserNameResp{
		Name: s,
	}, nil
}

func (us *UserService) SetUserPassword(ctx context.Context, req *user.SetUserPasswordReq) (*user.Empty, error) {
	return &user.Empty{}, us.urcache.SetUserPassword(ctx, req.Id, req.Password)
}

func (us *UserService) GetUserPasswordVersion(ctx context.Context, req *user.GetUserPasswordVersionReq) (*user.GetUserPasswordVersionResp, error) {
	u, err := us.urcache.GetUserPasswordVersion(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.GetUserPasswordVersionResp{
		Version: u,
	}, nil
}

func (us *UserService) SetUserRole(ctx context.Context, req *user.SetUserRoleReq) (*user.Empty, error) {
	return &user.Empty{}, us.urcache.SetUserRole(ctx, req.Id, req.Role)
}

func (us *UserService) SetUserStatus(ctx context.Context, req *user.SetUserStatusReq) (*user.Empty, error) {
	return &user.Empty{}, us.urcache.SetUserStatus(ctx, req.Id, req.Status)
}

func (us *UserService) ValidateUserPassword(ctx context.Context, req *user.ValidateUserPasswordReq) (*user.ValidateUserPasswordResp, error) {
	return &user.ValidateUserPasswordResp{
		Valid: us.db.CheckPassword(ctx, req.Id, req.Password),
	}, nil
}
