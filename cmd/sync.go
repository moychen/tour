package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"tour/internal/transmit"
)

const cfgFile = "F:\\Mine\\Bin\\config\\codesync.sql.ini"
var  sshHost string
var  fileType string

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "将生成的代码打包上传至编译环境并解压缩",
	Long:  "将生成的代码打包上传至编译环境并解压缩",
	Run:   func(cmd *cobra.Command, args []string) {
		dbInfo := &transmit.DBInfo{
			DBType:   "mysql",
			Host:     "127.0.0.1",
			UserName: "root",
			Password: "19960503",
			Charset:  "utf8mb4",
		}
		module := fileType
		if len(sshHost) > 0 {
			module += sshHost
		}

		sync := transmit.NewCodeSync(dbInfo)
		if err := sync.InitDB(); err != nil {
			log.Fatal(err.Error())
		}

		sshConfig := &transmit.SSHConfig{}
		transmitConfig := &transmit.TransmitConfig{}

		sync.ParseConfigFromDB(module, sshConfig, transmitConfig)
		transmit.Transmit(sshConfig, transmitConfig)
	},
}

func init() {
	syncCmd.Flags().StringVarP(&sshHost, "host", "o", "", `ssh主机ip最后一段，与配置文件中的ssh配置相对应`)
	syncCmd.Flags().StringVarP(&fileType, "type", "t", "", `传输的文件代码类型，与配置文件中的transmit配置相对应`)
}