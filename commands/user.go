package commands

import (
	"fmt"
	"github.com/ruanda/gogs-cli/gogs"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func showUser(u gogs.User) {
	fmt.Printf("%d;%s;%s;%s;%s\n", u.Id, u.Username, u.FullName, u.Email, u.AvatarURL)
}

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

var CmdUserShow = &cobra.Command{
	Use:   "show",
	Short: "show user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("id;username;fullname;email;avatar")
		for _, u := range args {
			user, err := GogsClient.UserShow(u)
			if err != nil {
				os.Exit(1)
			}
			if user != nil {
				showUser(*user)
			}
		}
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
		fmt.Println("id;username;fullname;email;avatar")
		for _, user := range users {
			showUser(user)
		}
	},
}

func init() {
	CmdRoot.AddCommand(CmdUser)
	CmdUser.AddCommand(CmdUserCreate)
	CmdUser.AddCommand(CmdUserShow)
	CmdUser.AddCommand(CmdUserSearch)
}
