package armory

import (
	"fmt"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func DO_02() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "DO_02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project documentation MUST include an explanation of the contribution process.",
		ControlID:   "OSPS-DO-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(DO_02_T01)

	return
}

func DO_02_T01() raidengine.MovementResult {

	body := GetData().Repository.ContributingGuidelines.Body
	containsGuidelines := len(body) > 100

	return raidengine.MovementResult{
		Description: "Discover whether the GitHub repo's recommended contributing guidelines has content.",
		Function:    utils.CallerPath(0),
		Passed:      containsGuidelines,
		Message:     fmt.Sprintf("Contributing Guidelines Found: %v", containsGuidelines),
	}
}
