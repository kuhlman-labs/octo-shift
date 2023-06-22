package main

import (
	"flag"
	"log"
	"os"

	octoshift "octo-shift/pkg"
)

func main() {

	// Subcommand: update-repo-visibility
	updateRepoVisibilityCmd := flag.NewFlagSet("update-repo-visibility", flag.ExitOnError)
	// update repo visibility flags
	updateRepoVisibilityOrg := updateRepoVisibilityCmd.String("org", "", "The name of the organization that owns the repository")
	updateRepoVisibilityRepo := updateRepoVisibilityCmd.String("repo", "", "The name of the repository")

	// Subcommand: update-webhooks
	updateWebhooksCmd := flag.NewFlagSet("update-webhooks", flag.ExitOnError)
	// update webhooks flags
	updateWebhooksOrg := updateWebhooksCmd.String("org", "", "The name of the organization that owns the repository")
	updateWebhooksRepo := updateWebhooksCmd.String("repo", "", "The name of the repository")

	// Subcommand: Create-Teams
	createTeamsCmd := flag.NewFlagSet("create-teams", flag.ExitOnError)
	// create teams flags
	createTeamsOrg := createTeamsCmd.String("org", "", "The name of the organization that owns the repository")
	createTeamsRepo := createTeamsCmd.String("repo", "", "The name of the repository")

	if len(os.Args) < 2 {
		log.Fatal("Please specify a subcommand. Choose from: update-repo-visibility, update-webhooks, create-teams")
	}

	switch os.Args[1] {
	case "update-repo-visibility":
		updateRepoVisibilityCmd.Parse(os.Args[2:])
	case "update-webhooks":
		updateWebhooksCmd.Parse(os.Args[2:])
	case "create-teams":
		createTeamsCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if updateRepoVisibilityCmd.Parsed() {
		if *updateRepoVisibilityOrg == "" {
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}
		if *updateRepoVisibilityRepo == "" {
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}
		octoshift.UpdateRepoVisibility(token, url, *updateRepoVisibilityOrg, *updateRepoVisibilityRepo)
	}

	if updateWebhooksCmd.Parsed() {
		if *updateWebhooksOrg == "" {
			updateWebhooksCmd.PrintDefaults()
			os.Exit(1)
		}
		if *updateWebhooksRepo == "" {
			updateWebhooksCmd.PrintDefaults()
			os.Exit(1)
		}
		octoshift.UpdateWebhooks(token, url, *updateWebhooksOrg, *updateWebhooksRepo)
	}

	if createTeamsCmd.Parsed() {
		if *createTeamsOrg == "" {
			createTeamsCmd.PrintDefaults()
			os.Exit(1)
		}
		if *createTeamsRepo == "" {
			createTeamsCmd.PrintDefaults()
			os.Exit(1)
		}
		octoshift.CreateTeams(token, url, *createTeamsOrg, *createTeamsRepo)
	}

}
