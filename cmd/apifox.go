package cmd

import (
	"tour/internal/jsonconvert"

	"github.com/spf13/cobra"
)

var filename string

var jsoncCmd = &cobra.Command{
	Use:   "jsonc",
	Short: "json格式转换",
	Long:  "json格式转换",
	Run: func(cmd *cobra.Command, args []string) {
		jsonconvert.Convert(filename)
	},
}

func init() {
	jsoncCmd.Flags().StringVarP(&filename, "file", "f", "", `需要转换的json文件名`)
}
