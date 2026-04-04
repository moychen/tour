package compress

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Zip 打包成zip文件
func Zip(srcDir, dstDir string) (string, error) {
	srcDir = strings.TrimSuffix(srcDir, "/")
	zipFileName := filepath.Join(dstDir, filepath.Base(srcDir) + ".zip")

	log.Println("src_dir: ", srcDir, ", zip_file: ",  zipFileName, ", dst_dir: ", dstDir)

	// 预防：旧文件无法覆盖
	os.RemoveAll(zipFileName)

	// 创建：zip文件
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		log.Println(err.Error())
		return zipFileName, err
	}

	defer zipFile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = filepath.Base(path)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	return zipFileName, err
}