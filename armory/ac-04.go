package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_04() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project’s version control system MUST prevent unintentional deletion of the primary branch.",
		ControlID:   "OSPS-AC-04",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_04_T01)

	return "AC_04", result
}

func AC_04_T01() raidengine.MovementResult {
	allowed := Data.GraphQL().Repository.DefaultBranchRef.RefUpdateRule.AllowsDeletions
	branchName := Data.GraphQL().Repository.DefaultBranchRef.Name

	message := fmt.Sprintf("Branch Protection Prevents Deletion: %v", !allowed)
	// TODO: check rules as well

	moveResult := raidengine.MovementResult{
		Description: "This movement is still under construction",
		Passed:      !allowed,
		Value:       fmt.Sprintf("Branch name: %s", branchName),
		Message:     message,
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return moveResult
}
