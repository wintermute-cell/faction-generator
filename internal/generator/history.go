package generator

import (
	"fmt"
	"math/rand"
	"project_factions/internal/util"
	m_content "project_factions/pkg/models/content"
	"slices"
)

func makeYearNumber(century int64) int64 {
	century = century*100 + rand.Int63n(100)
	return century
}

type HistoryRunner struct {
	constr           m_content.BaseConstraints
	dungeonStrata    int64
	availRaces       []m_content.Race
	historicalEvents []func(faction *Faction, century int64) []*Faction
}

func NewHistoryRunner(constr m_content.BaseConstraints, dungeonStrata int64, availRaces []m_content.Race) HistoryRunner {
	hr := HistoryRunner{
		constr:        constr,
		dungeonStrata: dungeonStrata,
		availRaces:    availRaces,
	}
	makeNewFaction := func() *Faction {
		return generateBaseFaction(hr.constr, hr.dungeonStrata, hr.availRaces)
	}

	hr.historicalEvents = []func(faction *Faction, century int64) []*Faction{
		// Diplomatic Alliance Event
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			// Randomly decide if a new faction is formed as a result of the alliance
			if rand.Intn(2) == 0 {
				newFaction := makeNewFaction()
				newFaction.History = append(newFaction.History, fmt.Sprintf("Formed from a diplomatic alliance between an unknown faction and %v in %v.", fac.FactionName, year))
				return []*Faction{newFaction}
			} else {
				fac.History = append(fac.History, fmt.Sprintf("A significant diplomatic alliance with an unknown group in %v strengthened the faction.", year))
				fac.Population = int64(float64(fac.Population) * 1.05) // Small population increase
			}

			return []*Faction{}
		},

		// Political Split Event
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			// There's always a chance for a political split
			if rand.Intn(2) == 0 {
				// A new faction is formed as a result of the split
				newFaction := makeNewFaction()
				newFaction.History = append(newFaction.History, fmt.Sprintf("Formed from a political split within %v in %v.", fac.FactionName, year))

				// The new faction might take a portion of the population
				popTransfer := int64(float64(fac.Population) * 0.2) // 20% of the population moves to the new faction
				fac.Population -= popTransfer
				newFaction.Population += popTransfer

				// Reflect the event in the history of the original faction
				fac.History = append(fac.History, fmt.Sprintf("Experienced a significant political split in %v, leading to the formation of a new faction.", year))

				return []*Faction{fac, newFaction}
			} else {
				// The faction undergoes internal restructuring but remains intact
				fac.History = append(fac.History, fmt.Sprintf("Underwent a major internal restructuring due to political tensions in %v but remained united.", year))
				// This restructuring might affect the faction's governance or military type
				fac.GovernmentType = m_content.GovernmentType(util.Clamp(int(fac.GovernmentType)+1, 1, int(m_content.GovernmentTypeMAX)-1))
			}

			return []*Faction{}
		},

		// The first event is an empty event, it does nothing
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Nothing of note happened in %v.", year))
			return []*Faction{}
		},

		// Great War Event
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			outcome := rand.Intn(2) // Randomly decide if the faction wins or loses

			popMultiplier := 1.0 + (rand.Float64()*0.2 - 0.1) // Multiplier varies between 0.9 and 1.1

			if outcome == 0 {
				// The faction wins the war
				fac.History = append(fac.History, fmt.Sprintf("The faction emerged victorious from a great war in %v, expanding its territory and influence.", year))
				newStrata := rand.Int63n(hr.dungeonStrata)
				if !slices.Contains(fac.HomeStrata, newStrata) {
					fac.HomeStrata = append(fac.HomeStrata, newStrata)
				}
				fac.Population = int64(float64(fac.Population) * popMultiplier)
				fac.Aggressiveness = m_content.AggressionLevel(util.Clamp(int(fac.Aggressiveness)+1, 1, int(m_content.AggressionLevelMAX)-1))
				if fac.Population == 0 {
					fac.History = append(fac.History, fmt.Sprintf("The faction was ultimately destroyed in the aftermath of the war in %v.", year))
					destroyFaction(fac)
					return []*Faction{} // Return an empty slice to indicate the faction is removed
				}
			} else {
				// The faction loses the war
				fac.History = append(fac.History, fmt.Sprintf("The faction was severely weakened by a lost war in %v, losing much of its territory.", year))
				fac.Population = int64(float64(fac.Population) * (2 - popMultiplier)) // Invert multiplier for population decrease
				if fac.Population < 0 {
					fac.Population = 0
				}

				// Faction destruction check
				if fac.Population == 0 {
					fac.History = append(fac.History, fmt.Sprintf("The faction was ultimately destroyed in the aftermath of the war in %v.", year))
					destroyFaction(fac)
					return []*Faction{} // Return an empty slice to indicate the faction is removed
				}
			}

			return []*Faction{}
		},

		// Magical Cataclysm Event
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("A magical cataclysm in %v altered the very fabric of the faction's domain.", year))
			fac.MagicLevel = m_content.MagicLevel(util.Clamp(int(fac.MagicLevel)+1, 1, int(m_content.MagicLevelMAX)-1))

			// Random chance to change home strata
			if rand.Intn(2) == 0 {
				fac.HomeStrata = []int64{rand.Int63n(hr.dungeonStrata)}
				fac.History = append(fac.History, fmt.Sprintf("The faction's home was forcibly relocated due to the cataclysm in %v.", year))
			}

			return []*Faction{}
		},

		// Technological Breakthrough Event
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("A technological breakthrough in %v catapulted the faction into a new era of innovation.", year))
			fac.TechLevel = m_content.TechLevel(util.Clamp(int(fac.TechLevel)+1, 1, int(m_content.TechLevelMAX)-1))

			return []*Faction{}
		},

		// A period of economic prosperity or decline
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			economicChange := rand.Float64() // Randomly determines if the event is positive or negative
			if economicChange < 0.5 {
				// Economic decline
				fac.History = append(fac.History, fmt.Sprintf("Faced economic hardships in %v, leading to a slight decrease in population and resources.", year))
				fac.Population = int64(float64(fac.Population) * 0.95) // 5% population decrease due to hardships
				if fac.Population == 0 {
					fac.History = append(fac.History, fmt.Sprintf("The faction was ultimately destroyed in the aftermath of the war in %v.", year))
					destroyFaction(fac)
					return []*Faction{} // Return an empty slice to indicate the faction is removed
				}
			} else {
				// Economic prosperity
				fac.History = append(fac.History, fmt.Sprintf("Enjoyed a period of economic prosperity in %v, boosting population and resources.", year))
				fac.Population = int64(float64(fac.Population) * 1.1) // 10% population increase due to prosperity
			}
			return []*Faction{}
		},

		// Discovery of a new magic source or loss of an existing one
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			magicChange := rand.Float64() // Randomly decides the nature of the event
			if magicChange < 0.5 {
				// Loss of magic source
				fac.History = append(fac.History, fmt.Sprintf("Lost a major source of magic in %v, causing a decline in magical abilities and influence.", year))
				fac.MagicLevel = m_content.MagicLevel(util.Clamp(int(fac.MagicLevel)-1, 1, int(m_content.MagicLevelMAX)-1))
			} else {
				// Discovery of new magic source
				fac.History = append(fac.History, fmt.Sprintf("Discovered a new source of magic in %v, greatly enhancing magical abilities and influence.", year))
				fac.MagicLevel = m_content.MagicLevel(util.Clamp(int(fac.MagicLevel)+1, 1, int(m_content.MagicLevelMAX)-1))
			}
			return []*Faction{}
		},

		// A significant technological invention or regression
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			techChange := rand.Float64() // Randomly determines if the event is positive or negative
			if techChange < 0.5 {
				// Technological regression
				fac.History = append(fac.History, fmt.Sprintf("Experienced a technological setback in %v, leading to a loss of knowledge and capabilities.", year))
				fac.TechLevel = m_content.TechLevel(util.Clamp(int(fac.TechLevel)-1, 1, int(m_content.TechLevelMAX)-1))
			} else {
				// Technological invention
				fac.History = append(fac.History, fmt.Sprintf("Achieved a groundbreaking technological invention in %v, propelling the faction forward.", year))
				fac.TechLevel = m_content.TechLevel(util.Clamp(int(fac.TechLevel)+1, 1, int(m_content.TechLevelMAX)-1))
			}
			return []*Faction{}
		},

		// A change in leadership
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			// Simulate the event of leadership change
			changeType := rand.Float64()
			if changeType < 0.5 {
				// Peaceful transition
				fac.History = append(fac.History, fmt.Sprintf("Underwent a peaceful change in leadership in %v, stabilizing the faction's governance.", year))
				// Potential impact on faction attributes can be added here
				fac.GovernmentType = generateFactionGovernmentType(fac)
			} else {
				// Turbulent succession
				fac.History = append(fac.History, fmt.Sprintf("Experienced a turbulent change in leadership in %v, leading to internal strife and unrest.", year))
				// Reflect this change by possibly decreasing population or changing aggressiveness
				fac.Population = int64(float64(fac.Population) * 0.98) // Assuming unrest leads to a small population decrease
				fac.Aggressiveness = m_content.AggressionLevel(util.Clamp(int(fac.Aggressiveness)+1, 1, int(m_content.AggressionLevelMAX)-1))
				if fac.Population == 0 {
					fac.History = append(fac.History, fmt.Sprintf("The faction was ultimately destroyed in the aftermath of the war in %v.", year))
					destroyFaction(fac)
					return []*Faction{} // Return an empty slice to indicate the faction is removed
				}
			}
			return []*Faction{}
		},

		// A natural disaster
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			disasterType := rand.Intn(3) // Randomly picks a type of natural disaster
			switch disasterType {
			case 0:
				fac.History = append(fac.History, fmt.Sprintf("Survived a devastating earthquake in %v, causing significant loss and damage.", year))
			case 1:
				fac.History = append(fac.History, fmt.Sprintf("Was hit by a catastrophic flood in %v, leading to widespread destruction.", year))
			case 2:
				fac.History = append(fac.History, fmt.Sprintf("Endured a severe drought in %v, severely impacting resources and population.", year))
			}

			// Reflecting the impact of the disaster
			fac.Population = int64(float64(fac.Population) * 0.9) // Assuming a 10% population decrease due to the disaster
			if fac.Population == 0 {
				fac.History = append(fac.History, fmt.Sprintf("The faction was ultimately destroyed in the aftermath of the war in %v.", year))
				destroyFaction(fac)
				return []*Faction{} // Return an empty slice to indicate the faction is removed
			}

			return []*Faction{}
		},

		// Cultural Renaissance or Dark Age
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			culturalShift := rand.Float64() // Randomly decides the nature of the cultural event
			if culturalShift < 0.5 {
				// Dark Age
				fac.History = append(fac.History, fmt.Sprintf("Entered a dark age in %v, leading to a decline in cultural and intellectual pursuits.", year))
				// This might affect MagicLevel or TechLevel negatively due to a decline in innovation and learning
				fac.TechLevel = m_content.TechLevel(util.Clamp(int(fac.TechLevel)-1, 1, int(m_content.TechLevelMAX)-1))
				fac.MagicLevel = m_content.MagicLevel(util.Clamp(int(fac.MagicLevel)-1, 1, int(m_content.MagicLevelMAX)-1))
			} else {
				// Cultural Renaissance
				fac.History = append(fac.History, fmt.Sprintf("Experienced a cultural renaissance in %v, flourishing in arts, science, and magic.", year))
				// This might boost TechLevel or MagicLevel due to increased innovation and magical discovery
				fac.TechLevel = m_content.TechLevel(util.Clamp(int(fac.TechLevel)+1, 1, int(m_content.TechLevelMAX)-1))
				fac.MagicLevel = m_content.MagicLevel(util.Clamp(int(fac.MagicLevel)+1, 1, int(m_content.MagicLevelMAX)-1))
			}
			return []*Faction{}
		},

		// A bountiful harvest
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Enjoyed a bountiful harvest in %v, boosting morale and resources.", year))
			// Reflect this positive event with a slight increase in population
			fac.Population = int64(float64(fac.Population) * 1.02)
			return []*Faction{}
		},

		// A trade agreement
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Formed a beneficial trade agreement with a neighboring faction in %v.", year))
			// This event might slightly increase TechLevel or MagicLevel due to exchange of knowledge
			fac.TechLevel = m_content.TechLevel(util.Clamp(int(fac.TechLevel)+1, 1, int(m_content.TechLevelMAX)-1))
			return []*Faction{}
		},

		// A minor conflict
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Engaged in a minor conflict with a rival faction in %v, resulting in stalemate.", year))
			// No significant changes, but it adds to the narrative
			return []*Faction{}
		},

		// Discovery of a small but significant resource
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Discovered a small but significant new resource vein in %v.", year))
			// Slight increase in population due to economic boost
			fac.Population = int64(float64(fac.Population) * 1.03)
			return []*Faction{}
		},

		// A minor plague
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Suffered from a minor plague in %v, causing some losses.", year))
			// Slight decrease in population due to the disease
			fac.Population = int64(float64(fac.Population) * 0.98)
			return []*Faction{}
		},

		// Founding of a new settlement or outpost
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Founded a new settlement in %v, expanding territorial control.", year))
			// Increase in population due to expansion
			fac.Population = int64(float64(fac.Population) * 1.05)
			return []*Faction{}
		},

		// A festival or significant cultural event
		func(fac *Faction, century int64) []*Faction {
			year := makeYearNumber(century)
			fac.History = append(fac.History, fmt.Sprintf("Celebrated a major festival in %v, strengthening community bonds.", year))
			// Purely narrative, no change to stats but enriches the faction's culture
			return []*Faction{}
		},
	}

	return hr
}

func (hr *HistoryRunner) RunHistory(faction *Faction, century int64) []*Faction {
	// Run a random historical event
	factions := []*Faction{}
	eventIdx := rand.Intn(len(hr.historicalEvents))
	if eventIdx <= 1 { // events 0 and 1 are rare events
		eventIdx = rand.Intn(len(hr.historicalEvents))
	}
	event := hr.historicalEvents[eventIdx]
	factions = append(factions, event(faction, century)...)
	return factions
}
