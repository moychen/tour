package cmd

import (
	"github.com/spf13/cobra"
	"tour/internal/config"
	"tour/internal/transmit"
)

const cfgFile = "config/sync.ini"

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "将生成的代码打包上传至编译环境并解压缩",
	Long:  "将生成的代码打包上传至编译环境并解压缩",
	Run:   func(cmd *cobra.Command, args []string) {
		cfg, err := config.Init(cfgFile)
		if err != nil {
			return
		}

		sshConfig := &transmit.SSHConfig{Port: 22}
		transmitConfig := &transmit.TransmitConfig{}

		config.ParseSection(cfg, "ssh", sshConfig)
		config.ParseSection(cfg, "transmit", transmitConfig)

		transmit.Transmit(sshConfig, transmitConfig)
	},
}

func init() {
	// syncCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", `需要计算的时间, 有效单位为时间戳或已格式化后的时间`)
}