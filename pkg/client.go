package octoshift

import (
	"context"
	"log"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func NewClient(token, url string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	if url != "https://api.github.com/" {
		log.Printf("Using Enterprise Server URL: %s", url)

		client, err := github.NewEnterpriseClient(url, url, tc)
		if err != nil {
			log.Fatal(err)
		}

		return client
	}

	client := github.NewClient(tc)

	return client
}
