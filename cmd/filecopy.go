package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"tour/internal/filecopy"
)

var configFile string

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "将某些文件从一个目录拷贝至另一个目录",
	Long:  "将某些文件从一个目录拷贝至另一个目录，方便本地同步改动",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := filecopy.ReadConfig(configFile)
		if err != nil {
			log.Fatal(err.Error())
		}

		err = filecopy.FileCopy(config)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	syncCmd.Flags().StringVarP(&configFile, "config_file", "c", "", `拷贝配置文件`)
}
