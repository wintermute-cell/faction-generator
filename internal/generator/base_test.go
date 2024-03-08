package generator

import (
	"project_factions/internal/content"
	"project_factions/internal/logging"
	m_content "project_factions/pkg/models/content"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBaseFactions(t *testing.T) {
	type testCase struct {
		name        string
		constraints m_content.BaseConstraints
	}

	var tests = []testCase{
		{
			name: "Large Dungeon High Tech",
			constraints: content.NewBaseConstraints(
				[]string{"core"},
				false,
				m_content.DungeonSizeLevelLarge,
				m_content.TechLevelFuturistic,
				m_content.MagicLevelHigh,
				m_content.AggressionLevelWarlike,
			),
		},
		{
			name: "Small Dungeon Primitive Tech",
			constraints: content.NewBaseConstraints(
				[]string{"core"},
				true,
				m_content.DungeonSizeLevelSmall,
				m_content.TechLevelPrimitive,
				m_content.MagicLevelNone,
				m_content.AggressionLevelPeaceful,
			),
		},
		// add more test cases as needed...
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logging.Info("Running test case %s", tc.name)
			result := GenerateBaseFactions(tc.constraints)

			for _, faction := range result {
				assert.NotEmpty(t, faction.FactionName, "")

				assert.LessOrEqual(t, faction.MultiRaceType, m_content.Coexistence)
				assert.GreaterOrEqual(t, faction.MultiRaceType, m_content.MultiRaceType(1))

				assert.LessOrEqual(t, faction.TechLevel, m_content.TechLevelMAX-1)
				assert.GreaterOrEqual(t, faction.TechLevel, m_content.TechLevel(1))

				assert.LessOrEqual(t, faction.MagicLevel, m_content.MagicLevelMAX-1)
				assert.GreaterOrEqual(t, faction.MagicLevel, m_content.MagicLevel(1))

				assert.LessOrEqual(t, faction.Aggressiveness, m_content.AggressionLevelMAX-1)
				assert.GreaterOrEqual(t, faction.Aggressiveness, m_content.AggressionLevel(1))

				assert.LessOrEqual(t, faction.GovernmentType, m_content.GovernmentTypeMAX-1)
				assert.GreaterOrEqual(t, faction.GovernmentType, m_content.GovernmentType(1))

				assert.LessOrEqual(t, faction.MilitaryType, m_content.MilitaryTypeMAX-1)
				assert.GreaterOrEqual(t, faction.MilitaryType, m_content.MilitaryType(1))

				assert.LessOrEqual(t, faction.PopAgeType, m_content.FactionPopAgeTypeMAX-1)
				assert.GreaterOrEqual(t, faction.PopAgeType, m_content.FactionPopAgeType(1))

				if faction.IsFallen {
					assert.Equal(t, faction.Population, int64(0))
					assert.Equal(t, len(faction.HomeStrata), 0)
				} else {
					assert.GreaterOrEqual(t, faction.Population, int64(1))
					assert.GreaterOrEqual(t, len(faction.HomeStrata), 1)
				}

				assert.GreaterOrEqual(t, len(faction.Ethics), 1)
			}
			//logging.Info("Got %v results", len(result))
			//for _, faction := range result {
			//	logging.Info("========\n%v", faction.String())
			//}
		})
	}
}
