package compress

import (
	"log"
	"os"
)

const (
	METHOD_ZIP = iota + 1
)

type CompressConfig struct {
	Method int8   `ini:"method"`
	SrcDir string `ini:"src_dir"`
	DstDir string `ini:"dst_dir"`
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息

	if err != nil {
		if os.IsExist(err) {

			return true

		}
		return false
	}

	return true
}

// 判断所给路径是否为文件夹

func IsDir(path string) bool {
	s, err := os.Stat(path)

	if err != nil {

		return false

	}
	return s.IsDir()
}

// 判断所给路径是否为文件

func IsFile(path string) bool {
	return !IsDir(path)
}

func Compress(config *CompressConfig) (string, error) {
	var zipFileName string
	var err error

	if Exists(config.SrcDir) == false {
		log.Fatal("check file path not exist: ", config.SrcDir)
	}

	switch config.Method {
	case METHOD_ZIP:
		zipFileName, err = Zip(config.SrcDir, config.DstDir)
	default:
		log.Fatal("compress not supported method: ", config.Method)
	}

	if err != nil {
		return zipFileName, nil
	}

	return zipFileName, err
}
