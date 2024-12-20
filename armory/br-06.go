package armory

import (
	"strings"

	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_06() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "BR_06"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "All releases MUST provide a descriptive log of functional and security modifications.",
		ControlID:   "OSPS-BR-06",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_06_T01)
	if result.Movements["BR_06_T01"].Message == "Releases Found" {
		result.ExecuteMovement(BR_06_T02)
	}
	return
}

func BR_06_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "Checking whether project has releases, passing if no releases are present",
		Function:    utils.CallerPath(0),
	}

	data := GetData()

	if data.Repository.Releases.TotalCount > 0 {
		moveResult.Message = "Releases Found"
	} else {
		moveResult.Passed = true
		moveResult.Message = "Releases Not Found"
	}
	// TODO: Use this section to write a single step or test that contributes to DO_01
	return
}

func BR_06_T02() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "Checking whether project has releases, passing if no releases are present",
		Function:    utils.CallerPath(0),
	}

	releaseDescription := GetData().Repository.LatestRelease.Description

	if strings.Contains(releaseDescription, "Change Log") || strings.Contains(releaseDescription, "Changelog") {
		moveResult.Passed = true
		moveResult.Message = "Change Log Found in Latest Release"
	}
	return
}
