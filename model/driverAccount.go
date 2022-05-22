package model

import (
	"fmt"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

func init() {
	// err := global.DB.AutoMigrate(&DriverAccount{})
	// if err != nil {
	// 	global.Logf.Errorf("driver account init error! | %v\n", err)
	// }
}

// DriverAccount 通过接口向数据库存入驱动器配置
type DriverAccount struct {
	gorm.Model
	Name string `json:"name" gorm:"unique" binding:"required"` // 存储账户的唯一名称
	//Type  string `json:"type" gorm:"type"`                      // 类型，即 driver 名
	Type  string `json:"type"`               // 类型，即 driver 名
	Index int    `json:"index" gorm:"index"` // 序列号, 由于数据中
	// for QiNiu
	AccessKey string `json:"access_key" gorm:"column:qn_ak"`
	SecretKey string `json:"secret_key" gorm:"column:qn_sk"`
	Bucket    string `json:"bucket"`
	Host      string `json:"host"`
	UseHTTPS  bool   `json:"use_https"`
	// 机房位置
	Zone string `json:"zone"`
}

var (
	defDriver = DriverAccount{
		Index:     0,
		Name:      "qbox-01",
		Type:      "qbox",
		AccessKey: "O_qaFxMC9xvkbMURCs5Dhcxf1EDZnUJKIlry72rh",
		SecretKey: "v7YJbHOh-zzKQPw9o-7D60uxSpHdpSObCL7ZOcMF",
		Bucket:    "j2go",
		Host:      "dy-resource.evlic.cn",
		UseHTTPS:  true,
		Zone:      "Huanan",
	}
)

// GetHost 使用后需要判断 host != ""
// 可以用于自行拼接 url 提前返回
func (d DriverAccount) GetHost() string {
	// 参考 qn-sdk/storage/region.go endpoint()
	host := strings.TrimSpace(d.Host)
	host = strings.TrimLeft(host, "http://")
	host = strings.TrimLeft(host, "https://")
	if host == "" {
		return ""
	}
	scheme := "http://"
	if d.UseHTTPS {
		scheme = "https://"
	}
	return fmt.Sprintf("%s%s", scheme, host)
}

// GetDriverAccount todo-feature 通过 path 匹配存储账户
// 现行逻辑直接返回七牛的存储账户
// 	上传时调用
func GetDriverAccount(path string) DriverAccount {
	var res DriverAccount
	if err := global.DB.First(&res).Error; err != nil {
		global.Logf.Errorf("query driver account error >> %v\n will use default driver(%#q)", err, defDriver.Name)
		// 执行 save
		go SaveDriverAccount(&defDriver)
		return defDriver
	}
	return res
}

func SaveDriverAccount(account *DriverAccount) {
	if err := global.DB.Save(account).Error; err != nil {
		zap.L().Error("save account err!!",
			zap.String("AccountName", account.Name),
			zap.Error(err),
		)
	}
}

// GetDriverAccounts 批量查找数据库中的 driver 账户
func GetDriverAccounts() ([]DriverAccount, error) {
	var res []DriverAccount
	if err := global.DB.Order(columnName("driver_no")).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
