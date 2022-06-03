package bootstrap

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"tiktok-lite/global"
	// "gorm.io/gorm"
	"moul.io/zapgorm2"
)

// InitDB 初始化mysql，获取mysql连接
func InitDB() *gorm.DB {

	// 获取全局配置对象
	serverConfig := global.Conf
	// 从serverConfig中获取mysql信息
	mysqlInfo := serverConfig.Mysql
	// 获取dsn
	// dsn := "root:drldrl521521@tcp(localhost:3306)/ssmbuild?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlInfo.Name,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName,
	)

	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 在查询中，可以使查询的表名为单数
		},
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}
	global.DB = db
	return db
}
