package armory

import (
	"fmt"
	"strings"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_06() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "All releases MUST provide a descriptive log of functional and security modifications.",
		ControlID:   "OSPS-BR-06",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_06_T01)
	if !strings.Contains(result.Movements["BR_06_T01"].Message, ": 0") {
		Logger.Trace("Releases Found, checking for Change Log")
		result.ExecuteMovement(BR_06_T02)
	}
	return "BR_06", result
}

func BR_06_T01() raidengine.MovementResult {
	releaseCount := Data.GraphQL().Repository.Releases.TotalCount

	return raidengine.MovementResult{
		Description: "Checking whether project has releases, passing if no releases are present",
		Function:    utils.CallerPath(0),
		Passed:      true,
		Message:     fmt.Sprintf("Releases Found: %v", releaseCount),
	}
}

func BR_06_T02() raidengine.MovementResult {
	releaseDescription := Data.GraphQL().Repository.LatestRelease.Description
	contains := (strings.Contains(releaseDescription, "Change Log") || strings.Contains(releaseDescription, "Changelog"))

	moveResult := raidengine.MovementResult{
		Description: "Checking whether project has releases, passing if no releases are present",
		Function:    utils.CallerPath(0),
		Passed:      contains,
		Message:     fmt.Sprintf("Change Log Found in Latest Release: %v", contains),
	}
	return moveResult
}
