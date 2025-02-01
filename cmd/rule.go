package cmd

import (
	"fmt"

	"github.com/krau/shisoimg/dao"
	"github.com/krau/shisoimg/utils"
	"github.com/spf13/cobra"
)

var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Manage URL rules",
}

var ruleAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new URL rule",
	Long: "Add a new URL rule. The first argument is the prefix, " +
		"the second argument is the path. The prefix is the URL " +
		"prefix that will be replaced with the path.",
	Args:    cobra.ExactArgs(2),
	Example: "shisoimg rule add https://img.example.com /data/images",
	Run: func(cmd *cobra.Command, args []string) {
		prefix, path := args[0], args[1]
		if err := dao.CreateRule(prefix, path); err != nil {
			utils.L.Fatal(err)
		}
		utils.L.Infof("Rule added: %s -> %s", prefix, path)
	},
}

var ruleDelCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a URL rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		if err := dao.DeleteRule(path); err != nil {
			utils.L.Fatal(err)
		}
		utils.L.Infof("Rule deleted: %s", path)
	},
}

var ruleListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all URL rules",
	Run: func(cmd *cobra.Command, args []string) {
		rules, err := dao.GetRules()
		if err != nil {
			utils.L.Fatal(err)
		}
		fmt.Println("Rules:")
		for _, rule := range rules {
			fmt.Printf("%s -> %s\n", rule.Prefix, rule.Path)
		}
	},
}

func init() {
	rootCmd.AddCommand(ruleCmd)
	ruleCmd.AddCommand(ruleAddCmd)
	ruleCmd.AddCommand(ruleDelCmd)
	ruleCmd.AddCommand(ruleListCmd)
}
