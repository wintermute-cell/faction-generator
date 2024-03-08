package content

import (
	"project_factions/internal/logging"
	m_content "project_factions/pkg/models/content"
)

func makeDefaultConstraints() m_content.BaseConstraints {
	return m_content.BaseConstraints{
		EnabledContent:  []*m_content.ContentPack{&CoreContentPack},
		ForbidMultiRace: false,
		DungeonSize:     m_content.DungeonSizeLevelLarge,
		MaxTech:         m_content.TechLevelMedieval,
		MaxMagic:        m_content.MagicLevelMedium,
		AvgAggression:   m_content.AggressionLevelNeutral,
	}
}

func NewBaseConstraints(
	contentPacks []string,
	forbitMultiRace bool,
	dungeonSize m_content.DungeonSizeLevel,
	maxTech m_content.TechLevel,
	maxMagic m_content.MagicLevel,
	avgAggression m_content.AggressionLevel,
) m_content.BaseConstraints {
	if len(contentPacks) == 0 {
		logging.Error("No content packs provided, using default content pack")
		return makeDefaultConstraints()
	}

	packs := make([]*m_content.ContentPack, len(contentPacks))
	for i, pack := range contentPacks {
		if val, ok := ContentPacks[pack]; ok {
			packs[i] = val
		} else {
			logging.Error("Content pack %s not found, using default content pack", pack)
			return makeDefaultConstraints()
		}
	}
	return m_content.BaseConstraints{
		EnabledContent:  packs,
		ForbidMultiRace: forbitMultiRace,
		DungeonSize:     dungeonSize,
		MaxTech:         maxTech,
		MaxMagic:        maxMagic,
		AvgAggression:   avgAggression,
	}
}
