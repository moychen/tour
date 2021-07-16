package transmit

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"tour/internal/compress"
)

const (
	METHOD_SFTP = iota + 1
)

type TransmitConfig struct {
	Method		int8	`ini:"method"`
	LocalPath	string	`ini:"local_path"`
	RemotePath	string	`ini:"remote_path"`
}

func Transmit(sshConfig *SSHConfig, transmitConfig *TransmitConfig) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	sshClient, err := NewSSHClient(sshConfig)
	if err != nil {
		return
	}
	defer sshClient.client.Close()

	transmitConfig.LocalPath = filepath.FromSlash(transmitConfig.LocalPath)
	dstDir := strings.TrimSuffix(transmitConfig.LocalPath, "\\")

	compressConfig := &compress.CompressConfig {
							Method: 1,
							SrcDir: transmitConfig.LocalPath,
							DstDir: filepath.Dir(dstDir),
						}

	zipFileName, err := compress.Compress(compressConfig)
	if err != nil {
		return
	}

	transmitConfig.LocalPath = zipFileName
	remotePathTmp := transmitConfig.RemotePath
	transmitConfig.RemotePath = filepath.ToSlash(filepath.Join(remotePathTmp, filepath.Base(zipFileName)))

	switch transmitConfig.Method {
	case METHOD_SFTP:
		err = SftpUpload(sshClient, transmitConfig)
	default:
		log.Fatal("transmit not supported: ", transmitConfig.Method)
	}

	if err != nil {
		return
	}

	cmd := "unzip -o -d " + filepath.ToSlash(remotePathTmp) + " " + transmitConfig.RemotePath
	out, err := sshClient.Exec(cmd)
	if err != nil {
		log.Println(cmd, out.OutputString(), err.Error())
		return
	} else {
		log.Println(cmd, out.OutputString())
	}

	err = os.RemoveAll(dstDir)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Remove path: ", dstDir)

	err = os.Remove(zipFileName)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Remove file: ", zipFileName)

	return
}