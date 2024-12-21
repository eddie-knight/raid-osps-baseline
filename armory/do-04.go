package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_04() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "The project documentation MUST include a policy for coordinated vulnerability reporting, with a clear timeframe for response.",
		ControlID:   "OSPS-DO-04",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_04_T01)

	return "DO_04", result
}

func DO_04_T01() raidengine.MovementResult {
	enabled := Data.GraphQL().Repository.IsSecurityPolicyEnabled

	moveResult := raidengine.MovementResult{
		Description: "Discover whether a security policy is enabled for this repo.",
		Function:    utils.CallerPath(0),
		Passed:      enabled,
		Message:     fmt.Sprintf("Security Policy Enabled: %v", enabled),
	}
	return moveResult
}
