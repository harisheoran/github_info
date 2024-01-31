package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v58/github"
	"github.com/spf13/cobra"
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Get list of Collaborators of Github Repo",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := getEnv()
		if err != nil {
			fmt.Println(err)
		} else {
			username, err := cmd.Flags().GetString("u")
			repo, err1 := cmd.Flags().GetString("r")
			if err != nil {
				log.Fatal("Username flag error", err)
			}
			if err1 != nil {
				log.Fatal("Username flag error", err)
			}

			if username == "" {
				fmt.Println("No username. Provide Username: --u=harish")
			} else if repo == "" {
				fmt.Println("No username. Provide Repository: --r=kubernetes")
			} else {
				getCollaborators(username, repo)
			}
		}
	},
}

// Prints the list of all collaborators
func getCollaborators(username, repo string) {
	token, _ := getEnv()
	list := getCollaboratorsList(token, username, repo)
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d. Name: %s Access: %s\n", i, *list[i].Login, *list[i].RoleName)
	}
}

// fetch collaborators list from GitHub API
func getCollaboratorsList(token, username, repo string) []*github.User {

	client := github.NewClient(nil).WithAuthToken(token)

	ctx := context.Background()
	response, _, err := client.Repositories.ListCollaborators(ctx, username, repo, &github.ListCollaboratorsOptions{})

	if err != nil {
		log.Print("Please provide correct username and repository name ", err)
	}

	return response
}

// Load the .env file and return the variable
func getEnv() (string, error) {
	token := os.Getenv("TOKEN")

	if len(token) == 0 {
		return "", errors.New("Need Github Token. Provide using: export TOKEN=\"<your-token>\"")
	}

	return token, nil
}

func init() {
	rootCmd.AddCommand(teamCmd)
	teamCmd.PersistentFlags().String("u", "", "GitHub Username")
	teamCmd.PersistentFlags().String("r", "", "Repository Name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// teamCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teamCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
