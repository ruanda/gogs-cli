package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var CmdUser = &cobra.Command{
	Use:   "user",
	Short: "user functions",
}

var CmdUserCreate = &cobra.Command{
	Use:   "create",
	Short: "create user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lol")
	},
}

var CmdUserSearch = &cobra.Command{
	Use:   "search",
	Short: "search user",
	Run: func(cmd *cobra.Command, args []string) {

		users, err := GogsClient.UserSearch(strings.Join(args, " "), 10)
		if err != nil {
			os.Exit(1)
		}
		for _, u := range users {
			fmt.Println(u.Username)
		}
	},
}

func init() {
	CmdRoot.AddCommand(CmdUser)
	CmdUser.AddCommand(CmdUserCreate)
	CmdUser.AddCommand(CmdUserSearch)
}
