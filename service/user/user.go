package user

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	db *dbUtils
	user.UnimplementedUserServer
}

func NewUserService(c *conf.DatabaseConfig) *UserService {
	d, err := dbdial.NewDatabase(context.Background(), c)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if c.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.User),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	db := &UserService{
		db: newDBUtils(d),
	}
	return db
}

func user2GetUserResp(u *model.User) *user.GetUserResp {
	return &user.GetUserResp{
		Id:        u.ID,
		CreatedAt: u.CreatedAt.UnixMicro(),
		UpdatedAt: u.UpdatedAt.UnixMicro(),
		Name:      u.Username,
		Role:      u.Role,
		Status:    u.Status,
	}

}

func (us *UserService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.GetUserResp, error) {
	u := &model.User{
		Username: req.Name,
		Role:     req.Role,
		Status:   req.Status,
	}
	err := SetUserPassword(u, req.Password)
	if err != nil {
		return nil, err
	}
	err = us.db.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return user2GetUserResp(u), nil
}

func (us *UserService) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	u, err := us.db.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return user2GetUserResp(u), nil
}

func (us *UserService) SetUserName(ctx context.Context, req *user.SetUserNameReq) (*user.Empty, error) {
	return &user.Empty{}, us.db.SetUserName(req.Id, req.Name)
}

func (us *UserService) SetUserPassword(ctx context.Context, req *user.SetUserPasswordReq) (*user.Empty, error) {
	return &user.Empty{}, us.db.SetUserPassword(req.Id, req.Password)
}

func (us *UserService) SetUserRole(ctx context.Context, req *user.SetUserRoleReq) (*user.Empty, error) {
	return &user.Empty{}, us.db.SetUserRole(req.Id, req.Role)
}

func (us *UserService) SetUserStatus(ctx context.Context, req *user.SetUserStatusReq) (*user.Empty, error) {
	return &user.Empty{}, us.db.SetUserStatus(req.Id, req.Status)
}

func (us *UserService) ValidateUserPassword(ctx context.Context, req *user.ValidateUserPasswordReq) (*user.ValidateUserPasswordResp, error) {
	return &user.ValidateUserPasswordResp{
		Valid: us.db.CheckPassword(req.Id, req.Password),
	}, nil
}
