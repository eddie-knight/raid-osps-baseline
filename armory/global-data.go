package armory

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var Authenticated bool

type RepoData struct {
	// Need to update token for this
	// Organization struct {
	// 	Rulesets struct {
	// 		Nodes []struct {
	// 			BypassActors struct {
	// 				TotalCount int
	// 			}
	// 		}
	// 	} `graphql:"rulesets(first: 10)"`
	// } `graphql:"organization(login: $owner)"`

	Repository struct {
		Name                    string
		HasDiscussionsEnabled   bool
		HasIssuesEnabled        bool
		IsSecurityPolicyEnabled bool
		DefaultBranchRef        struct {
			Name          string
			RefUpdateRule struct {
				AllowsDeletions              bool
				AllowsForcePushes            bool
				RequiredApprovingReviewCount bool
			}
		}
		Releases struct {
			TotalCount int
		}
		// BranchProtectionRule	struct{
		// // 	allowsForcePushes			bool
		// // 	requiresApprovingReviews	bool
		// // 	restrictsPushes				bool
		// // 	allowsDeletions				bool
		// 	RequiresStatusChecks		bool
		// }
		LatestRelease struct {
			Description string
		}
		ContributingGuidelines struct {
			Body         string
			ResourcePath githubv4.URI
		}
		BranchProtectionRules struct {
			Nodes []struct {
				AllowsDeletions          bool
				AllowsForcePushes        bool
				RequiresApprovingReviews bool
				RequiresCommitSignatures bool
				RequiresStatusChecks     bool
			}
		} `graphql:"branchProtectionRules(first: 10)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

var GlobalData RepoData

func GetData() RepoData {
	if GlobalData.Repository.Name != "" {
		return GlobalData
	}
	owner := GlobalConfig.GetString("owner")
	repo := GlobalConfig.GetString("repo")
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GlobalConfig.GetString("token")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(repo),
	}

	err := client.Query(context.Background(), &GlobalData, variables)
	if err != nil {
		Logger.Error("Error querying GitHub GraphQL API: ", err)
	}
	return GlobalData
}
