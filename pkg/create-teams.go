// This script will query a given GHES organization for teams and create those teams in a given GHEC organization.

package octoshift

import (
	"context"
	"log"

	"github.com/google/go-github/v53/github"
)

func CreateTeams(sourceClient, targetClient *github.Client, sourceOrg, targetOrg string) {
	// Get teams from Source Org
	sourceTeams := getSourceOrgTeams(sourceClient, sourceOrg)

	// Create teams in Target Org
	createTargetTeams(targetClient, targetOrg, sourceTeams)
}

func getSourceOrgTeams(client *github.Client, org string) []*github.Team {
	teams, _, err := client.Teams.ListTeams(context.Background(), org, nil)
	if err != nil {
		log.Fatal(err)
	}

	return teams
}

func createTargetTeams(client *github.Client, orgName string, teams []*github.Team) {

	for _, team := range teams {
		teamName := team.GetName()
		teamLDAPMapping := team.GetLDAPDN()

		// Create team in GHEC Org
		newTeam := github.NewTeam{
			Name: teamName,
		}

		createdTeam, _, err := client.Teams.CreateTeam(context.Background(), orgName, newTeam)
		if err != nil {
			log.Fatal(err)
		}

		// Get External Group based on LDAP DN
		externalGroupList, _, err := client.Teams.ListExternalGroupsForTeamBySlug(context.Background(), orgName, teamLDAPMapping)
		if err != nil {
			log.Fatal(err)
		}

		// Get Group ID
		externalGroupID := externalGroupList.Groups[0].GetGroupID()

		// Add LDAP mapping to team
		teamLDAPMappingOptions := &github.ExternalGroup{
			GroupID: &externalGroupID,
		}
		_, _, err = client.Teams.UpdateConnectedExternalGroup(context.Background(), orgName, createdTeam.GetName(), teamLDAPMappingOptions)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Created team %s in %s\n", createdTeam.GetName(), orgName)
	}
}
