package armory

import (
	"github.com/hashicorp/go-hclog"
	"github.com/privateerproj/privateer-sdk/config"
	"github.com/privateerproj/privateer-sdk/raidengine"
)

type ArmoryData struct {
	graphql GraphqlData
	rest    RestData
}

var (
	Authenticated bool
	GlobalConfig  *config.Config
	Logger        hclog.Logger
	Data          ArmoryData
	Armory        = raidengine.Armory{
		Tactics: map[string][]raidengine.Strike{
			"dev": {
				DO_01,
				DO_02,
				DO_04,
				DO_05,
				BR_06,
				AC_01,
				QA_01,
				BR_02,
			},
			"maturity_1": {
				AC_01,
				AC_02,
				AC_03,
				AC_04,
				BR_01,
				BR_02,
				BR_03,
				DO_01,
				DO_02,
				LE_01,
				LE_02,
				LE_03,
				QA_01,
				QA_02,
			},
			"maturity_2": {
				AC_05,
				BR_04,
				BR_05,
				BR_06,
				BR_07,
				DO_03,
				DO_04,
				DO_05,
				DO_06,
				DO_07,
				DO_11,
				DO_12,
				LE_04,
				QA_03,
				QA_04,
				QA_05,
				QA_06,
			},
			"maturity_3": {
				AC_06,
				DO_08,
				DO_09,
				DO_10,
				QA_07,
			},
		},
	}
)

func SetupArmory(c *config.Config) {
	GlobalConfig = c
	Logger = c.Logger
	if c.GetString("token") == "" {
		Armory.Tactics = unauthenticatedTactics()
	} else {
		Authenticated = true
	}
}

func unauthenticatedTactics() map[string][]raidengine.Strike {
	return map[string][]raidengine.Strike{
		"dev": {
			QA_01,
			BR_02,
		},
		"maturity_1": {},
		"maturity_2": {},
		"maturity_3": {},
	}
}

func (r *ArmoryData) Rest() RestData {
	if r.rest.repo == "" {
		r.rest.loadData()
	}
	return r.rest
}

func (r *ArmoryData) GraphQL() GraphqlData {
	if r.graphql.Repository.Name == "" {
		r.loadGraphQLData()
	}
	return r.graphql
}
