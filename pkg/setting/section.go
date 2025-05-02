package setting

type Config struct {
	Server ServerSettings `mapstructure:"server"`
	Logger LoggerSettings `mapstructure:"logger"`
	Mysql  MySQLSettings  `mapstructure:"mysql"`
	Redis  RedisSettings  `mapstructure:"redis"`
	AwsS3  AWSS3Settings  `mapstructure:"awsS3"`
	Gmail  GmailSettings  `mapstructure:"Gmail"`
}

type ServerSettings struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type LoggerSettings struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_Backups   int    `mapstructure:"max_backups"`
	Max_Age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}

type MySQLSettings struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	UserName        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type RedisSettings struct {
}

type AWSS3Settings struct {
	BucketName      string `mapstructure:"bucket_name"`
	Region          string `mapstructure:"region"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"seret_access_key"`
}

type GmailSettings struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}
