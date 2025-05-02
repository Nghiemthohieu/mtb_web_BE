package initialize

import (
	"fmt"
	"mtb_web/global"
	"mtb_web/internal/models"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CheckErrerPanic(err error, errString string) {
	if err != nil {
		// Log chi tiết lỗi
		global.Logger.Error(fmt.Sprintf("%s: %v", errString, err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.Mysql

	// Xây dựng chuỗi kết nối DSN
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Port, m.DbName)

	// Kết nối MySQL
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	CheckErrerPanic(err, "Failed to connect to MySQL")

	// Log thành công
	global.Logger.Info("Initialized MySQL Successfully")
	global.Mdb = db

	// Thiết lập pool kết nối
	SetPool()

	// Thực hiện migration các bảng
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	CheckErrerPanic(err, "Failed to get MySQL DB instance")

	// Thiết lập các thông số kết nối
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

	// Log pool kết nối
	global.Logger.Info("MySQL connection pool set successfully")
}

func migrateTables() {
	// Thực hiện migration các bảng
	err := global.Mdb.AutoMigrate(
		&models.Category{},
		&models.Product{},
		&models.ProductSize{},
		&models.ProductColor{},
		&models.ProductMaterial{},
		&models.ProductImage{},
		&models.ProductStyle{},
		&models.User{},
		&models.Role{},
		&models.ProductVariant{},
		&models.ProductReview{},
		&models.Permission{},
		&models.ProductOrder{},
		&models.Order{},
		&models.PaymentMethod{},
		&models.Staffs{},
		&models.Customer{},
		&models.ProductReviewImg{},
		&models.OrderStatus{},
		&models.Notification{},
		&models.SlideShow{},
	)

	// Kiểm tra lỗi migration
	if err != nil {
		// Log lỗi khi migration không thành công
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
		panic(fmt.Sprintf("Migration failed: %v", err))
	}

	// Log thành công migration
	global.Logger.Info("Tables migrated successfully")
}
