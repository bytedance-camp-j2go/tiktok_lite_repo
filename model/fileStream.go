package model

import (
	"io"
)

// FileStream 文件上传封装对象, 需要 handler 收到上传后创建对象
// 然后移交此对象给 driver 对象处理具体的存储
type FileStream struct {
	File io.ReadCloser
	Size int64
	// 父级 url, 相当于指定传入的保存文件夹, 比如 /uuid/2006/01/
	ParentPath string
	// 文件名称，要求和上传的一样，最好有 .xxx 后缀
	Name string
	// 文件类型 application/json.... from 对象中有
	MIMEType string
}

func (file FileStream) Read(p []byte) (n int, err error) {
	return file.File.Read(p)
}

func (file FileStream) GetMIMEType() string {
	return file.MIMEType
}

func (file FileStream) GetSize() int64 {
	return file.Size
}

func (file FileStream) Close() error {
	return file.File.Close()
}

func (file FileStream) GetFileName() string {
	return file.Name
}

func (file FileStream) GetParentPath() string {
	return file.ParentPath
}
