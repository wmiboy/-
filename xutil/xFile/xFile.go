package xFile

import (
	"os"
)

//文件操作

//写入文件，不存在则创建
func Save_File(name string, b []byte) error {
	//打开文件
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		return err
	}
	defer f.Close()

	i, err := f.Write(b)
	if err != nil {
		return err
	}
	if i != len(b) {
		return err
	}
	return nil
}
