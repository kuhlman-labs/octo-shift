// This script will update the visibility of a GHEC repository to the visibility of the source GHES repository.

package octoshift

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
)

func UpdateRepoVisibility(client *github.Client, sourceOrg, targetOrg, sourceRepo, targetRepo string) {
	// GHEC Repo
	ghecOrgName := "kuhman-labs-fabrikam-org"
	ghecRepoName := "test-repo"

	// Get visibility of GHES repo
	sourceVisibility := getSourceRepoVisibility(client, sourceOrg, sourceRepo)

	// Get visibility of GHEC repo
	targetVisibility := getTargetRepoVisibility(client, targetOrg, targetRepo)

	// Compare visibilities
	if sourceVisibility == targetVisibility {
		log.Printf("Visibility of %s is already %s", targetRepo, sourceVisibility)
		return
	}

	// Update visibility of Target repo
	updateTargetRepoVisibility(client, ghecOrgName, ghecRepoName, sourceVisibility)
}

func getSourceRepoVisibility(client *github.Client, orgName, repoName string) (visibility string) {

	log.Printf("Getting visibility of %s", repoName)

	repo, _, err := client.Repositories.Get(context.Background(), orgName, repoName)
	if err != nil {
		log.Fatal(err)
	}

	if repo.GetVisibility() == "public" {
		return "internal"
	}

	return repo.GetVisibility()
}

func getTargetRepoVisibility(client *github.Client, orgName, repoName string) (visibility string) {
	log.Printf("Getting visibility of %s", repoName)

	repo, _, err := client.Repositories.Get(context.Background(), orgName, repoName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(repo.GetVisibility())

	return repo.GetVisibility()
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
