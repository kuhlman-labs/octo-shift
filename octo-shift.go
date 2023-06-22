package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	octoshift "octo-shift/pkg"
)

func main() {

	// Subcommand: update-repo-visibility
	updateRepoVisibilityCmd := flag.NewFlagSet("update-repo-visibility", flag.ExitOnError)
	// update repo visibility flags
	updateRepoVisibilitySourceOrg := updateRepoVisibilityCmd.String("source-org", "", "The name of the source organization that owns the repository")
	updateRepoVisibilitySourceRepo := updateRepoVisibilityCmd.String("source-repo", "", "The name of the source repository")
	updateRepoVisibilityTargetOrg := updateRepoVisibilityCmd.String("target-org", "", "The name of the target organization that owns the repository")
	updateRepoVisibilityTargetRepo := updateRepoVisibilityCmd.String("target-repo", "", "The name of the target repository")
	updateRepoVisibilitySourceToken := updateRepoVisibilityCmd.String("source-token", "", "The token for the source GitHub Enterprise Server")
	updateRepoVisibilitySourceURL := updateRepoVisibilityCmd.String("source-url", "", "The URL for the source GitHub Enterprise Server")
	updateRepoVisibilityTargetToken := updateRepoVisibilityCmd.String("target-token", "", "The token for the target GitHub Enterprise Cloud")

	// Subcommand: update-webhooks
	updateWebhooksCmd := flag.NewFlagSet("update-webhooks", flag.ExitOnError)
	// update webhooks flags
	updateWebhooksTargetOrg := updateWebhooksCmd.String("target-org", "", "The name of the target organization.")
	updateWebhooksTargetToken := updateWebhooksCmd.String("target-token", "", "The token for the target GitHub Enterprise Cloud")
	updateWebhooksSecret := updateWebhooksCmd.String("secret", "", "The secret for the webhook")

	// Subcommand: Create-Teams
	createTeamsCmd := flag.NewFlagSet("create-teams", flag.ExitOnError)
	// create teams flags
	createTeamsSourceOrg := createTeamsCmd.String("source-org", "", "The name of the source GitHub Enterprise Server Organization that has the teams")
	createTeamsSourceToken := createTeamsCmd.String("source-token", "", "The token for the source GitHub Enterprise Server")
	createTeamsSourceURL := createTeamsCmd.String("source-url", "", "The URL for the source GitHub Enterprise Server")
	createTeamsTargetOrg := createTeamsCmd.String("target-org", "", "The name of the target GitHub Enterprise Cloud Organization that will have the teams")
	createTeamsTargetToken := createTeamsCmd.String("target-token", "", "The token for the target GitHub Enterprise Cloud")

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
		if *updateRepoVisibilitySourceOrg == "" {
			fmt.Println("Please specify a source organization")
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}
		if *updateRepoVisibilitySourceRepo == "" {
			fmt.Println("Please specify a source repository")
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}

		if *updateRepoVisibilityTargetOrg == "" {
			fmt.Println("Please specify a target organization")
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}

		if *updateRepoVisibilityTargetRepo == "" {
			fmt.Println("Please specify a target repository")
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)

		}

		if *updateRepoVisibilitySourceToken == "" {
			fmt.Println("Please specify a source token")
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}

		if *updateRepoVisibilitySourceURL == "" {
			fmt.Println("Please specify a source URL")
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}

		if *updateRepoVisibilityTargetToken == "" {
			fmt.Println("Please specify a target token")
			updateRepoVisibilityCmd.PrintDefaults()
			os.Exit(1)
		}

		sourceClient := octoshift.NewClient(*updateRepoVisibilitySourceToken, *updateRepoVisibilitySourceURL)
		targetClient := octoshift.NewClient(*updateRepoVisibilityTargetToken, "https://api.github.com/")

		octoshift.UpdateRepoVisibility(sourceClient, targetClient, *updateRepoVisibilitySourceOrg, *updateRepoVisibilityTargetOrg, *updateRepoVisibilitySourceRepo, *updateRepoVisibilityTargetRepo)
	}

	if updateWebhooksCmd.Parsed() {
		if *updateWebhooksTargetOrg == "" {
			fmt.Println("Please specify a target organization")
			updateWebhooksCmd.PrintDefaults()
			os.Exit(1)
		}

		if *updateWebhooksTargetToken == "" {
			fmt.Println("Please specify a target token")
			updateWebhooksCmd.PrintDefaults()
			os.Exit(1)
		}

		if *updateWebhooksSecret == "" {
			fmt.Println("Please specify a secret")
			updateWebhooksCmd.PrintDefaults()
			os.Exit(1)
		}

		token := *updateWebhooksTargetToken
		url := "https://api.github.com/"
		client := octoshift.NewClient(token, url)

		octoshift.UpdateWebhooks(client, *updateWebhooksTargetOrg, *updateWebhooksSecret)
	}

	if createTeamsCmd.Parsed() {
		if *createTeamsSourceOrg == "" {
			fmt.Println("Please specify a source organization")
			createTeamsCmd.PrintDefaults()
			os.Exit(1)
		}

		if *createTeamsSourceToken == "" {
			fmt.Println("Please specify a source token")
			createTeamsCmd.PrintDefaults()
			os.Exit(1)

		}

		if *createTeamsSourceURL == "" {
			fmt.Println("Please specify a source URL")
			createTeamsCmd.PrintDefaults()
			os.Exit(1)
		}

		if *createTeamsTargetOrg == "" {
			fmt.Println("Please specify a target organization")
			createTeamsCmd.PrintDefaults()
			os.Exit(1)
		}

		if *createTeamsTargetToken == "" {
			fmt.Println("Please specify a target token")
			createTeamsCmd.PrintDefaults()
			os.Exit(1)
		}

		token := *createTeamsTargetToken
		url := "https://api.github.com/"
		sourceClient := octoshift.NewClient(token, *createTeamsSourceURL)
		targetClient := octoshift.NewClient(token, url)

		octoshift.CreateTeams(sourceClient, targetClient, *createTeamsSourceOrg, *createTeamsTargetOrg)
	}

}
