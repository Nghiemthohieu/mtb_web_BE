package global

import (
	"mtb_web/pkg/logger"
	"mtb_web/pkg/setting"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config      setting.Config
	Logger      *logger.LoggerZap = logger.NewLogger(Config.Logger)
	Mdb         *gorm.DB
	AwsConfig   aws.Config
	AwsS3Client *s3.Client
	Rdb         *redis.Client
)
