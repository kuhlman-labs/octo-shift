package main

import (
	"context"
	"log"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func NewEnterpriseServerClient(ghesToken string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghesToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client, err := github.NewEnterpriseClient(ghesBaseURL, ghesUploadURL, tc)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func NewClient(ghecToken string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghecToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}
