package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env読み込みエラー: ", err)
	} else {
		fmt.Println(".env読み込み成功")
	}
}
