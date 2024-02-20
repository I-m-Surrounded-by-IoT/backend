package create

import (
	"context"
	"strings"
	"time"

	userApi "github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create user",
	Run:   CreateUser,
}

func CreateUser(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatalf("usage: user create <name> <password>")
	}
	username := args[0]
	password := args[1]
	config := cmd.Context().Value("config").(*conf.UserServer)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	d, err := dbdial.Dial(ctx, config.Database)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	urole, ok := userApi.Role_value[strings.ToUpper(role)]
	if !ok {
		log.Fatalf("invalid role: %s", role)
	}
	ustatus, ok := userApi.Status_value[strings.ToUpper(status)]
	if !ok {
		log.Fatalf("invalid status: %s", status)
	}
	u := &model.User{
		Username: username,
		Role:     userApi.Role(urole),
		Status:   userApi.Status(ustatus),
	}
	err = user.SetUserPassword(u, password)
	if err != nil {
		log.Fatalf("failed to set user password: %v", err)
	}
	err = user.NewDBUtils(d).CreateUser(u)
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	log.Infof("create user success: %s", u.Username)
}

func init() {
	CreateCmd.PersistentFlags().StringVar(&role, "role", userApi.Role_USER.String(), "user role")
	CreateCmd.PersistentFlags().StringVar(&status, "status", userApi.Status_ACTIVE.String(), "user status")
}
