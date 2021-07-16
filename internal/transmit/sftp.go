package transmit

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

// conn 存放连接的结构体
type conn struct {
	client	*ssh.Client
	sftpClient 	*sftp.Client
}

// TransferInfo 存放上传或下载的信息
type TransferInfo struct {
	Kind         string 	// upload或download
	Local        string   	// 本地路径
	Dst          string		// 目标路径
	TransferByte int64 		// 传输的字节数(byte)
}

func (t *TransferInfo) String()  string {
	return fmt.Sprintf(`TransferInfo(Kind:"%s", Local: "%s", Dst: "%s", TransferByte: %d)`,
		t.Kind, t.Local, t.Dst, t.TransferByte)
}

type SFTPClient struct {
	client *sftp.Client
}

func NewSFTPClient(sshClient *ssh.Client) (*SFTPClient, error) {
	sftpClient := &SFTPClient{}
	var err error

	sftpClient.client, err = sftp.NewClient(sshClient)
	if err != nil {
		log.Println("SFTP client init failed!")
		return sftpClient, err
	}

	return sftpClient, err
}

// Upload 将本地文件上传到远程主机上
func (s *SFTPClient) Upload(localPath string, dstPath string) (*TransferInfo, error) {
	transferInfo := &TransferInfo{Kind: "upload", Local: localPath, Dst: dstPath, TransferByte: 0}

	var err error // 如果sftp客户端没有打开，就打开，为了复用
	localFileObj, err := os.Open(localPath)

	if err != nil {
		return transferInfo, err
	}

	defer localFileObj.Close()

	dstFileObj, err := s.client.Create(dstPath)

	if err != nil {
		return transferInfo, err
	}

	defer dstFileObj.Close()

	written, err := io.Copy(dstFileObj, localFileObj)

	if err != nil {
		return transferInfo, err
	}

	transferInfo.TransferByte = written

	return transferInfo, nil
}

// Download 从远程主机上下载文件到本地
func (s *SFTPClient) Download(dstPath string, localPath string)  (*TransferInfo, error) {
	transferInfo := &TransferInfo{Kind: "download", Local: localPath, Dst: dstPath, TransferByte: 0}

	var err error

	localFileObj, err := os.Create(localPath)

	if err != nil {
		return transferInfo, err
	}

	defer localFileObj.Close()

	dstFileObj, err := s.client.Open(dstPath)

	if err != nil {
		return transferInfo, err
	}

	defer dstFileObj.Close()

	written, err := io.Copy(localFileObj, dstFileObj)

	if err != nil {
		return transferInfo, err
	}

	transferInfo.TransferByte = written
	return transferInfo, nil
}



func SftpUpload(sshClient *SSHClient, transmitConfig *TransmitConfig) error {
	sftpClient, err := NewSFTPClient(sshClient.client)
	if err != nil {
		return err
	}

	defer sftpClient.client.Close()

	// 上传文件
	transInfoUpload, err := sftpClient.Upload(transmitConfig.LocalPath, transmitConfig.RemotePath)
	if err != nil {
		log.Println(transInfoUpload, err.Error())
		return err
	} else {
		log.Println(transInfoUpload)
	}

	return nil
}