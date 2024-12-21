package armory

import (
	"fmt"
	"strings"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_02() (string, raidengine.StrikeResult) {
	result := raidengine.StrikeResult{
		Description: "All releases and released software assets MUST be assigned a unique version identifier for each release intended to be used by users.",
		ControlID:   "OSPS-BR-02",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_02_T01)
	if !strings.Contains(result.Movements["BR_02_T01"].Message, ": 0") {
		result.ExecuteMovement(BR_02_T02)
	}
	return "BR_02", result
}

func BR_02_T01() raidengine.MovementResult {
	releases := Data.Rest().Repo.Releases

	moveResult := raidengine.MovementResult{
		Description: "Discover all releases on the repository",
		Function:    utils.CallerPath(0),
		Passed:      true,
		Value:       releases,
		Message:     fmt.Sprintf("Releases found: %v", len(releases)),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return moveResult
}

func BR_02_T02() raidengine.MovementResult {
	releases := Data.Rest().Repo.Releases

	releaseNames := make(map[string]int)
	var errorCount int
	for _, release := range releases {
		if release.Name == "" {
			errorCount++
			GlobalConfig.Logger.Error("Release %v has no name!", release.Name)
		} else if _, ok := releaseNames[release.Name]; ok {
			errorCount++
			GlobalConfig.Logger.Error(fmt.Sprintf(
				"Release id: %v has the same name as another release: %s", release.Id, release.Name))
		} else {
			releaseNames[release.Name] = release.Id
		}
	}

	moveResult := raidengine.MovementResult{
		Description: "Ensure all releases have a unique name",
		Function:    utils.CallerPath(0),
		Passed:      errorCount == 0,
		Message:     fmt.Sprintf("Non-unique release names: %v", errorCount),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return moveResult
}
