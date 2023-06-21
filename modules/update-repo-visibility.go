// This script will update the visibility of a GHEC repository to the visibility of the source GHES repository.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v53/github"
)

func updateRepoVisibility() {
	// GHEC Repo
	ghecOrgName := "kuhman-labs-fabrikam-org"
	ghecRepoName := "test-repo"

	// GHES Repo
	ghesOrgName := "Engineering"
	ghesRepoName := "platform"

	// Get visibility of GHES repo
	ghesVisibility := getGHESRepoVisibility(ghesOrgName, ghesRepoName)

	// Get visibility of GHEC repo
	ghecVisibility := getGHECRepoVisibility(ghecOrgName, ghecRepoName)

	// Compare visibilities
	if ghecVisibility == ghesVisibility {
		log.Printf("Visibility of %s is already %s", ghecRepoName, ghesVisibility)
		return
	}

	// Update visibility of GHEC repo
	updateGHECRepoVisibility(ghecOrgName, ghecRepoName, ghesVisibility)
}

func getGHESRepoVisibility(orgName, repoName string) (visibility string) {
	log.Printf("Getting visibility of %s", repoName)

	client := NewEnterpriseServerClient()
	ctx := context.Background()

	repo, _, err := client.Repositories.Get(ctx, orgName, repoName)
	if err != nil {
		log.Fatal(err)
	}

	if repo.GetVisibility() == "public" {
		return "internal"
	}

	return repo.GetVisibility()
}

func getGHECRepoVisibility(orgName, repoName string) (visibility string) {
	log.Printf("Getting visibility of %s", repoName)

	client := NewClient()
	ctx := context.Background()

	repo, _, err := client.Repositories.Get(ctx, orgName, repoName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(repo.GetVisibility())

	return repo.GetVisibility()
}

func updateGHECRepoVisibility(orgName, repoName, visibility string) {
	log.Printf("Updating visibility of %s", repoName)

	client := NewClient()
	ctx := context.Background()

	repo := &github.Repository{
		Visibility: &visibility,
	}

	_, _, err := client.Repositories.Edit(ctx, orgName, repoName, repo)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Visibility of %s updated to %s", repoName, visibility)
}
