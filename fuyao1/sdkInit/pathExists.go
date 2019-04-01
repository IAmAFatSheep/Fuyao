package sdkInit

import "os"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil //表示文件或文件夹存在
	}
	if os.IsNotExist(err) {
		return false, nil //表示文件或文件夹不存在
	}
	return false, err
}
