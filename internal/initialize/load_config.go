package initialize

import (
	"fmt"
	"mtb_web/global"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load .env trước (ưu tiên biến nhạy cảm)
	if err := godotenv.Load("config/.env"); err != nil {
		panic(fmt.Errorf("⚠️  .env file not found, continuing with system environment variables: %v", err))
	}

	v := viper.New()
	v.AddConfigPath("config") // <-- Nên tương đối từ gốc project
	v.SetConfigName("local")  // Tên file YAML (local.yaml)
	v.SetConfigType("yaml")
	v.AutomaticEnv() // Cho phép override từ ENV

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("❌ error reading configuration file: %v", err))
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("❌ unable to decode configuration into struct: %v", err))
	}

	// ⚠️ Gán thủ công biến nhạy cảm từ ENV
	global.Config.AwsS3.AccessKeyID = v.GetString("AWS_ACCESS_KEY_ID")
	global.Config.AwsS3.SecretAccessKey = v.GetString("AWS_SECRET_ACCESS_KEY")
	global.Config.Gmail.User = v.GetString("EMAIL_USERNAME")
	global.Config.Gmail.Pass = v.GetString("EMAIL_PASSWORD")

	fmt.Println("✅ Loaded config successfully")
}
