package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func AC_01() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "AC_01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project's version control system MUST require multi-factor authentication for collaborators modifying the project repository settings or accessing sensitive data.",
		ControlID:   "OSPS-AC-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_01_T01)

	return
}

// TODO
func AC_01_T01() raidengine.MovementResult {
	required := GetData().Organization.RequiresTwoFactorAuthentication

	return raidengine.MovementResult{
		Description: "Inspect the repo's parent to ensure that all members are required to use MFA",
		Function:    utils.CallerPath(0),
		Passed:      required,
		Message:     fmt.Sprintf("MFA Required: %v", required),
	}
}
