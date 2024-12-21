package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_01() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles.",
		ControlID:   "OSPS-DO-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_01_T01)

	return "DO_01", result
}

// TODO
func DO_01_T01() raidengine.MovementResult {
	repoData := Data.GraphQL().Repository

	var message string
	if repoData.HasDiscussionsEnabled {
		message = "Discussions are enabled. "
	}
	if repoData.HasIssuesEnabled {
		message = message + "Issues are enabled."
	}

	moveResult := raidengine.MovementResult{
		Description: "Discover whether issues or discussions are enabled on the repo.",
		Function:    utils.CallerPath(0),
		Message:     message,
		Passed:      repoData.HasDiscussionsEnabled || repoData.HasIssuesEnabled,
	}
	return moveResult
}
