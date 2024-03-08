package content

import m_content "project_factions/pkg/models/content"

var CoreContentPack = m_content.ContentPack{
	Races: []m_content.Race{
		{
			Name:                "Human",
			Alignment:           m_content.RaceAlignmentNeutral,
			HostilityFactor:     0.3,
			TechFactor:          0.5,
			MagicFactor:         0.3,
			WeirdnessPreference: 0.0,
			PopFactor:           0.5,
			DominanceFactor:     0.5,
			ProducedResources:   []string{"Water", "Grain", "Berries", "Wood", "Meat"},
		},
		{
			Name:                "Elf",
			Alignment:           m_content.RaceAlignmentGood,
			HostilityFactor:     0.2,
			TechFactor:          0.4,
			MagicFactor:         0.6,
			WeirdnessPreference: 0.2,
			PopFactor:           0.3,
			DominanceFactor:     0.8,
			ProducedResources:   []string{"Water", "Berries", "Wood", "Meat", "Artifacts"},
		},
		{
			Name:                "Dwarf",
			Alignment:           m_content.RaceAlignmentGood,
			HostilityFactor:     0.3,
			TechFactor:          0.6,
			MagicFactor:         0.4,
			WeirdnessPreference: 0.1,
			PopFactor:           0.3,
			DominanceFactor:     1.0,
			ProducedResources:   []string{"Stone", "Metal", "Gemstones", "Mushrooms", "Armor"},
		},
		{
			Name:                "Halfling",
			Alignment:           m_content.RaceAlignmentGood,
			HostilityFactor:     0.1,
			TechFactor:          0.3,
			MagicFactor:         0.4,
			WeirdnessPreference: 0.1,
			PopFactor:           0.6,
			DominanceFactor:     0.5,
			ProducedResources:   []string{"Water", "Grain", "Berries", "Meat", "Cloth"},
		},
		{
			Name:                "Goblin",
			Alignment:           m_content.RaceAlignmentEvil,
			HostilityFactor:     0.5,
			TechFactor:          0.25,
			MagicFactor:         0.2,
			WeirdnessPreference: 0.2,
			PopFactor:           0.8,
			DominanceFactor:     0.2,
			ProducedResources:   []string{"Berries", "Meat", "Mushrooms", "Leather"},
		},
		{
			Name:                "Orc",
			Alignment:           m_content.RaceAlignmentEvil,
			HostilityFactor:     0.7,
			TechFactor:          0.3,
			MagicFactor:         0.1,
			WeirdnessPreference: 0.2,
			PopFactor:           0.5,
			DominanceFactor:     0.7,
			ProducedResources:   []string{"Meat", "Leather", "Weapons", "Armor"},
		},
		{
			Name:                "Troll",
			Alignment:           m_content.RaceAlignmentEvil,
			HostilityFactor:     0.8,
			TechFactor:          0.2,
			MagicFactor:         0.1,
			WeirdnessPreference: 0.3,
			PopFactor:           0.3,
			DominanceFactor:     0.3,
			ProducedResources:   []string{"Meat", "Mushrooms"},
		},
		{
			Name:                "Dragon",
			Alignment:           m_content.RaceAlignmentNeutral,
			HostilityFactor:     0.9,
			TechFactor:          0.0,
			MagicFactor:         0.9,
			WeirdnessPreference: 0.6,
			PopFactor:           0.0,
			DominanceFactor:     1.0,
			ProducedResources:   []string{"Gemstones", "Artifacts", "Armor", "Weapons"},
		},
	},
}
