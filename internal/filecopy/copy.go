package filecopy

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type CopyConfig struct {
	SrcPath  string   `ini:"src_path"`
	DestPath string   `ini:"dest_path"`
	Files    []string `ini:"files"`
}

func ReadConfig(configfile string) (CopyConfig, error) {
	var config = CopyConfig{
		SrcPath:  "E:\\1. Codebase",
		DestPath: "E:\\1. Codebase",
	}

	// config.Files = append(config.Files, "")

	return config, nil
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

// CreateDir 创建目录
func CreateDir(dest string, mode os.FileMode) error {
	//如果目录存在就直接返回
	if _, err := os.Stat(dest); err != nil {
		mkErr := os.MkdirAll(dest, mode)
		if mkErr != nil {
			return mkErr
		}
	}
	return nil
}

func FileCopy(config CopyConfig) error {
	var srcfilename, destfilename string

	log.Println("source dir: ", config.SrcPath)
	log.Println("destination dir: ", config.DestPath)

	config.SrcPath = filepath.FromSlash(config.SrcPath + "\\")
	config.DestPath = filepath.FromSlash(config.DestPath + "\\")
	// dstDir := strings.TrimSuffix(transmitConfig.LocalPath, "\\")

	for idx, name := range config.Files {
		srcfilename = config.SrcPath + name
		destfilename = config.DestPath + name

		srcfilename = filepath.ToSlash(srcfilename)
		destfilename = filepath.ToSlash(destfilename)

		log.Println("copy file: ", idx, " : ", name)

		if !Exists(srcfilename) || (!IsFile(srcfilename) && !IsDir(srcfilename)) {
			log.Println("file not exists: ", srcfilename)
			continue
		}

		if IsDir(srcfilename) {
			err := CreateDir(destfilename, 755)
			if err != nil {
				log.Fatal(err.Error())
			}

			continue
		}

		source, err := os.Open(srcfilename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer source.Close()

		destination, err := os.Create(destfilename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			log.Fatal(err.Error())
		}
		// destination.Seek(0, 0)
		// content, _ := io.ReadAll(destination)
		// log.Println(string(content))
	}

	return nil
}
