package operate

import (
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/driver/base"
	"github.com/bytedance-camp-j2go/tiktok_lite_repo/model"
)

// 示例 >> 定义驱动器操作

// File 指定 path 返回文件对象,
// 相当于让 Driver 执行一次 /path/to/file 的文件查询操作
func File(account *model.DriverAccount, path string) (*model.File, error) {
	driver, ok := base.GetDriver(account.Name)
	if !ok {
		return nil, base.ErrNotSupport
	}
	return driver.File(path, account)
}

// Upload 上传文件对象, 上传完成后会返回 url ( 根据驱动器配置, 可以直接使用的那种
// 也可以异步上传, 直接拼接 URL 返回给客户端, 但这么一定需要控制文件名不重复,
// 		且驱动器没做强制断点续传要求, 使用 Host() 获取 URL 的方式需要谨慎
func Upload(account *model.DriverAccount, stream model.FileStream) (string, error) {
	driver, ok := base.GetDriver(account.Name)
	if !ok {
		return "", base.ErrNotSupport
	}
	return driver.Upload(stream, account)
}
