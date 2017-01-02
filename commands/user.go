package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/ruanda/gogs-cli/gogs"
	"github.com/spf13/cobra"
)

var userNew struct {
	sourceId   int
	loginName  string
	username   string
	email      string
	password   string
	sendNotify bool
}

func showUser(u gogs.User) {
	fmt.Printf(`User
    ID:         %d
    Username:   %s
    Full name:  %s
    Email:      %s
    Avatar:     %s
`, u.Id, u.Username, u.FullName, u.Email, u.AvatarURL)
}

var cmdUser = &cobra.Command{
	Use:   "user",
	Short: "user functions",
}

var cmdUserCreate = &cobra.Command{
	Use:   "create",
	Short: "create user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Name: %s, Email: %s\n", userNew.username, userNew.email)
	},
}

var cmdUserShow = &cobra.Command{
	Use:   "show",
	Short: "show user",
	Run: func(cmd *cobra.Command, args []string) {
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

var cmdUserSearch = &cobra.Command{
	Use:   "search",
	Short: "search user",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := GogsClient.UserSearch(strings.Join(args, " "), 10)
		if err != nil {
			os.Exit(1)
		}
		for _, user := range users {
			showUser(user)
		}
	},
}

func init() {
	cmdRoot.AddCommand(cmdUser)
	cmdUser.AddCommand(cmdUserCreate)
	cmdUser.AddCommand(cmdUserShow)
	cmdUser.AddCommand(cmdUserSearch)
	cmdUserCreate.Flags().IntVar(&userNew.sourceId, "sourceid", 0, "Authentication source ID")
	cmdUserCreate.Flags().StringVar(&userNew.loginName, "login", "", "Authentication source login name")
	cmdUserCreate.Flags().StringVar(&userNew.username, "username", "", "Unique user name")
	cmdUserCreate.Flags().StringVar(&userNew.email, "email", "", "Unique email address of user")
	cmdUserCreate.Flags().StringVar(&userNew.password, "password", "", "Default password for user")
	cmdUserCreate.Flags().BoolVar(&userNew.sendNotify, "notify", false, "Send a notification email for this creation")

}
