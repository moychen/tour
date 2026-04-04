package transmit

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type CodeSync struct {
	DBEngine 	*gorm.DB
	DBInfo		*DBInfo
}

type PubTsyncconfig struct {
	Module  string `json:module`
	Ip      string `json:ip`
	Port    uint32 `json:port`
	User    string `json:user`
	Password        string `json:password`
	LocalPath       string `json:local_path`
	RemotePath      string `json:remote_path`
	Timeout int64 `json:timeout`
	CompressType    string `json:compress_type`
	ClearFlag       string `json:clear_flag`
	CreateDate      time.Time `json:create_date`
	UpdateDate      time.Time `json:update_date`
}

func NewCodeSync(dbInfo *DBInfo) *CodeSync {
	return &CodeSync{DBEngine: nil, DBInfo: dbInfo}
}

func (sync *CodeSync) InitDB() error {
	if sync.DBInfo == nil {
		log.Fatal("the DBInfo is not init")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/codesync?charset=%s&parseTime=True&loc=Local",
		sync.DBInfo.UserName,
		sync.DBInfo.Password,
		sync.DBInfo.Host,
		sync.DBInfo.Charset,
	)

	var err error
	sync.DBEngine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (sync *CodeSync) ParseConfigFromDB(module string, sshConfig *SSHConfig, transmitConfig *TransmitConfig) {
	var syncConfig PubTsyncconfig

	if sync.DBEngine == nil || sshConfig == nil || transmitConfig == nil {
		log.Fatal("Config is not init!")
	}

	sync.DBEngine.Table("pub_tsyncconfig").Where("module = ?", module).First(&syncConfig)
	sshConfig.Ip = syncConfig.Ip
	sshConfig.Port = syncConfig.Port
	sshConfig.Password = syncConfig.Password
	sshConfig.Timeout = time.Duration(syncConfig.Timeout)
	sshConfig.User = syncConfig.User
	transmitConfig.RemotePath = syncConfig.RemotePath
	transmitConfig.LocalPath = syncConfig.LocalPath

	compressType, err := strconv.Atoi(syncConfig.CompressType)
	if err != nil {
		log.Fatal("Config err, the compress_type is not defined!")
	}

	transmitConfig.Method = compressType

	return
}

func Close(sync *CodeSync) {
	//sync.DBEngine.Close()
	sync.DBEngine.DB()
}
