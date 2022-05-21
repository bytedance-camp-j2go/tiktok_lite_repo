package bootstrap

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	// "gorm.io/gorm"
)

// 初始化mysql，获取mysql连接
// InitDB 初始化mysql，获取mysql连接
func InitDB() *gorm.DB {

	// 获取全局配置对象
	serverConfig := global.Conf
	// 从serverConfig中获取mysql信息
	mysqlInfo := serverConfig.Mysql
	// 获取dsn
	// dsn := "root:drldrl521521@tcp(localhost:3306)/ssmbuild?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		mysqlInfo.Name,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 在查询中，可以使查询的表名为单数
		},
	})
	if err != nil {
		panic(err)
	}
	global.DB = db
	return db
}