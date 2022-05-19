package operate

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/base"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

// 示例 >> 定义驱动器操作

// File 指定 path 返回文件对象,
// 相当于让 Driver 执行一次 /path/to/file 的文件查询操作
func File(driver base.Driver, account *model.DriverAccount, path string) (*model.File, error) {
	return driver.File(path, account)
}
