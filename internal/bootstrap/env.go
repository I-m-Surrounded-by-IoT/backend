package bootstrap

import (
	"fmt"

	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/joho/godotenv"
)

func LoadEnvFromFile() error {
	s, err := utils.GetEnvFiles(".")
	if err != nil {
		return fmt.Errorf("failed to get env files: %v", err)
	}
	if len(s) != 0 {
		err = godotenv.Load(s...)
		if err != nil {
			return fmt.Errorf("failed to load env files: %v", err)
		}
	}
	return nil
}
