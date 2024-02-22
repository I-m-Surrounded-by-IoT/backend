package create

import (
	"strings"

	userApi "github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user"
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
	userService := user.NewUserService(config.Database, config.Config, config.Redis)

	urole, ok := userApi.Role_value[strings.ToUpper(role)]
	if !ok {
		log.Fatalf("invalid role: %s", role)
	}
	ustatus, ok := userApi.Status_value[strings.ToUpper(status)]
	if !ok {
		log.Fatalf("invalid status: %s", status)
	}
	userinfo, err := userService.CreateUser(cmd.Context(), &userApi.CreateUserReq{
		Username: username,
		Password: password,
		Role:     userApi.Role(urole),
		Status:   userApi.Status(ustatus),
	})
	if err != nil {
		log.Fatalf("create user failed: %v", err)
	}
	log.Infof("create user success: %+v", userinfo)
}

func init() {
	CreateCmd.PersistentFlags().StringVar(&role, "role", userApi.Role_USER.String(), "user role")
	CreateCmd.PersistentFlags().StringVar(&status, "status", userApi.Status_ACTIVE.String(), "user status")
}
