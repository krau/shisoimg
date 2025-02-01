package cmd

import (
	"github.com/krau/shisoimg/dao"
	"github.com/krau/shisoimg/utils"
	"github.com/spf13/cobra"
)

var imageAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new images",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		count, err := dao.CreateImagesFromDir(path)
		if err != nil {
			utils.L.Fatal(err)
		}
		utils.L.Infof("Image added: %d", count)
	},
}

func init() {
	rootCmd.AddCommand(imageAddCmd)
}
