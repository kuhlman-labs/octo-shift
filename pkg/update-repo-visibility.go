// This script will update the visibility of a GHEC repository to the visibility of the source GHES repository.

package octoshift

import (
	"context"
	"log"

	"github.com/google/go-github/v53/github"
)

func UpdateRepoVisibility(sourceClient, targetClient *github.Client, sourceOrg, targetOrg string) {
	// Get repos from Source Org
	sourceRepos := getSourceRepos(sourceClient, sourceOrg)

	// Update repos in Target Org
	updateTargetRepos(targetClient, targetOrg, sourceRepos)
}

func updateTargetRepoVisibility(client *github.Client, orgName, repoName, visibility string) {
	log.Printf("Updating visibility of %s", repoName)

	repo := &github.Repository{
		Visibility: &visibility,
	}

	_, _, err := client.Repositories.Edit(context.Background(), orgName, repoName, repo)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Visibility of %s updated to %s", repoName, visibility)
}

func getSourceRepos(client *github.Client, org string) []*github.Repository {
	repos, _, err := client.Repositories.ListByOrg(context.Background(), org, nil)
	if err != nil {
		log.Fatal(err)
	}

	return repos
}

func updateTargetRepos(client *github.Client, targetOrg string, repos []*github.Repository) {
	for _, repo := range repos {
		visibility := repo.GetVisibility()
		if visibility == "public" {
			visibility = "internal"
		}
		updateTargetRepoVisibility(client, targetOrg, repo.GetName(), visibility)
	}
}
