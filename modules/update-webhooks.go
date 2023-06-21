// This script will update the webhooks of a GHEC Organization to add a secret.

package main

import (
	"context"
	"log"

	"github.com/google/go-github/v53/github"
)

type repoHook struct {
	repoName string
	hooks    []*github.Hook
}

func updateWebhooks(secret string) {
	// GHEC Org
	ghecOrgName := "kuhman-labs-fabrikam-org"

	// Get webhooks from GHEC Org
	ghecOrgWebhooks := getGHECOrgWebhooks(ghecOrgName)

	// Get webhooks from GHEC Repo
	ghecRepoWebhooks := getGHECRepoWebhooks(ghecOrgName)

	// Add secret to Org webhooks
	updateGHECOrgWebhooks(ghecOrgName, secret, ghecOrgWebhooks)

	// Add secret to Repo webhooks
	updateGHECRepoWebhooks(ghecOrgName, secret, ghecRepoWebhooks)

}

func getGHECOrgWebhooks(orgName string) []*github.Hook {
	client := NewClient()

	webhooks, _, err := client.Organizations.ListHooks(context.Background(), orgName, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found %d webhooks in the %s organization.", len(webhooks), orgName)

	return webhooks
}

func getGHECRepoWebhooks(orgName string) []repoHook {
	client := NewClient()

	// Get Repos from Org
	repos, _, err := client.Repositories.ListByOrg(context.Background(), orgName, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d repos in the %s organization.", len(repos), orgName)

	var repoHooks []repoHook

	for _, repo := range repos {

		webhooks, _, err := client.Repositories.ListHooks(context.Background(), orgName, *repo.Name, nil)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Found %d webhooks in the %s/%s repository.", len(webhooks), orgName, *repo.Name)

		if len(webhooks) == 0 {
			continue
		}

		repoHook := repoHook{
			repoName: *repo.Name,
			hooks:    webhooks,
		}

		repoHooks = append(repoHooks, repoHook)
	}

	return repoHooks
}

func updateGHECOrgWebhooks(orgName, secret string, webhooks []*github.Hook) {
	client := NewClient()

	for _, webhook := range webhooks {

		hookConfig := map[string]interface{}{
			"url":    *webhook.URL,
			"secret": secret,
		}

		log.Printf("Updating webhook %s", *webhook.URL)
		_, _, err := client.Organizations.EditHook(context.Background(), orgName, *webhook.ID, &github.Hook{
			Config: hookConfig,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Updated webhook %s in the %s org", *webhook.URL, orgName)
	}
}

func updateGHECRepoWebhooks(orgName, secret string, webhooks []repoHook) {
	client := NewClient()

	for _, webhook := range webhooks {
		for _, hook := range webhook.hooks {
			hookConfig := map[string]interface{}{
				"url":    *hook.URL,
				"secret": secret,
			}
			log.Printf("Updating webhook %s", *hook.URL)
			_, _, err := client.Repositories.EditHook(context.Background(), orgName, webhook.repoName, *hook.ID, &github.Hook{
				Config: hookConfig,
			})
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Updated webhook %s in the %s repository", *hook.URL, webhook.repoName)
		}
	}
}
