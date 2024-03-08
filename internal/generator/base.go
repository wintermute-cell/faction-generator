package generator

import (
	"fmt"
	"math"
	"math/rand"
	"project_factions/internal/logging"
	"project_factions/internal/util"
	m_content "project_factions/pkg/models/content"
	"project_factions/pkg/models/ethics"
	"slices"
	"strings"
)

type Faction struct {
	FactionName    string
	Races          []m_content.Race
	MultiRaceType  m_content.MultiRaceType
	TechLevel      m_content.TechLevel
	MagicLevel     m_content.MagicLevel
	Aggressiveness m_content.AggressionLevel
	GovernmentType m_content.GovernmentType
	MilitaryType   m_content.MilitaryType
	PopAgeType     m_content.FactionPopAgeType
	Population     int64
	HomeStrata     []int64
	Ethics         []ethics.Ethic
	History        []string
	IsFallen       bool
}

func (fac *Faction) String() string {
	repr := fmt.Sprintf("Name: %s\nMultiRaceType: %v\nRaces: [", fac.FactionName, fac.MultiRaceType)
	for _, race := range fac.Races {
		repr += fmt.Sprintf("%s, ", race.Name)
	}
	repr = strings.TrimSuffix(repr, ", ")
	repr += fmt.Sprintf("]\nTech: %v\nMagic: %v\nAggression: %v\nPopulation: %v\nHomeStrata: %v\nEthics: %v\nGovernment: %v\nMilitary: %v\nPopAge: %v\n", fac.TechLevel, fac.MagicLevel, fac.Aggressiveness, fac.Population, fac.HomeStrata, fac.Ethics, fac.GovernmentType, fac.MilitaryType, fac.PopAgeType)
	repr += fmt.Sprintf("Fallen: %v\n", fac.IsFallen)
	repr += "History: [\n"
	for _, event := range fac.History {
		repr += fmt.Sprintf("\t- %s\n", event)
	}
	repr += "]"
	return repr
}

func destroyFaction(faction *Faction) {
	faction.IsFallen = true
	faction.History = append(faction.History, "The faction has fallen. Ruins were left behind on stratum "+fmt.Sprintf("%v", faction.HomeStrata[0]))
	faction.HomeStrata = []int64{}
	faction.Population = 0
}

func generateFactionName(constr m_content.BaseConstraints, faction Faction, deepness float64) string {
	if faction.Races == nil || len(faction.Races) == 0 {
		return "Missing Races! Can't generate a name."
	}

	return factionNameGenerator()
}

func generateFactionRaces(constr m_content.BaseConstraints, faction *Faction, allRaces []m_content.Race) []m_content.Race {
	// Shuffle all races, select the first three, sort by deepness preference, select a random amount of them

	races := []m_content.Race{}
	rand.Shuffle(len(allRaces), func(i, j int) { allRaces[i], allRaces[j] = allRaces[j], allRaces[i] })
	topThree := allRaces[:(util.Min(3, len(allRaces)))]
	if constr.ForbidMultiRace {
		races = append(races, topThree[0])
	} else {
		amountOfRaces := rand.Intn(3) + 1
		mixedRaces := topThree[:amountOfRaces]
		races = append(races, mixedRaces...)
	}

	// sort the races by their dominance factor
	if rand.Float64() > 0.05 { // Most of the time, the races will adhere to the dominance factor
		slices.SortFunc(races, func(a, b m_content.Race) int {
			return int(a.DominanceFactor - b.DominanceFactor)
		})
	}

	// if there are mixed "good" and "evil" races, set faction MultiRaceType to Domination
	hasGood := false
	hasEvil := false
	for _, race := range races {
		if race.Alignment == m_content.RaceAlignmentGood {
			hasGood = true
		} else if race.Alignment == m_content.RaceAlignmentEvil {
			hasEvil = true
		}
	}
	if hasGood && hasEvil {
		faction.MultiRaceType = m_content.Domination
	} else {
		faction.MultiRaceType = m_content.Coexistence
	}

	return races
}

func generateFactionEthics() []ethics.Ethic {
	var ret []ethics.Ethic
	ethicPoints := 3
	usedEthics := make(map[ethics.EthicKind]bool)

	for ethicPoints > 0 && len(usedEthics) < int(ethics.EthicKindMAX)-1 {
		eKind := ethics.EthicKind(rand.Intn(int(ethics.EthicKindMAX)-1) + 1) // Randomly select an ethic kind
		if eKind == ethics.EthicGestaltConsciousness {                       // Gestalt Consciousness should be very rare
			eKind = ethics.EthicKind(rand.Intn(int(ethics.EthicKindMAX)-1) + 1)
		}

		// Check for mutual exclusivity
		if !isExclusive(eKind, usedEthics) {
			continue // Skip if exclusive ethic is already selected
		}

		// Check if ethic is already selected
		if usedEthics[eKind] {
			continue
		}

		isFanatic := rand.Intn(2) == 0     // Randomly decide if this ethic will be fanatic
		if isFanatic && ethicPoints >= 2 { // Ensure enough points for fanatic
			ret = append(ret, ethics.Ethic{EthicKind: eKind, IsFanatic: true})
			ethicPoints -= 2 // Deduct points for fanatic ethic
		} else if ethicPoints >= 1 {
			ret = append(ret, ethics.Ethic{EthicKind: eKind, IsFanatic: false})
			ethicPoints-- // Deduct point for normal ethic
		}

		usedEthics[eKind] = true
	}

	return ret
}

// isExclusive checks if the chosen ethic is not mutually exclusive with already chosen ethics
func isExclusive(eKind ethics.EthicKind, usedEthics map[ethics.EthicKind]bool) bool {
	switch eKind {
	case ethics.EthicAuthoritarian:
		return !usedEthics[ethics.EthicEgalitarian]
	case ethics.EthicEgalitarian:
		return !usedEthics[ethics.EthicAuthoritarian]
	case ethics.EthicXenophobic:
		return !usedEthics[ethics.EthicXenophilic]
	case ethics.EthicXenophilic:
		return !usedEthics[ethics.EthicXenophobic]
	case ethics.EthicMilitarist:
		return !usedEthics[ethics.EthicPacifist]
	case ethics.EthicPacifist:
		return !usedEthics[ethics.EthicMilitarist]
	case ethics.EthicSpiritualist:
		return !usedEthics[ethics.EthicMaterialist]
	case ethics.EthicMaterialist:
		return !usedEthics[ethics.EthicSpiritualist]
	default:
		return true
	}
}

func generateFactionGovernmentType(faction *Faction) m_content.GovernmentType {
	return m_content.GovernmentType(rand.Intn(int(m_content.GovernmentTypeMAX-2)) + 1)
}

func generateFactionMilitaryType(faction *Faction) m_content.MilitaryType {
	for _, ethic := range faction.Ethics {
		if ethic.EthicKind == ethics.EthicMilitarist {
			return m_content.MilitaryTypeElite
		} else if ethic.EthicKind == ethics.EthicPacifist {
			return m_content.MilitaryTypeNone
		}
	}
	mult := ((faction.Races[0].HostilityFactor*4 + faction.Races[0].TechFactor + faction.Races[0].MagicFactor) / 6) + ((rand.Float64() + 0.5) * 0.3)
	mult = util.Clamp(mult, 0, 1)
	return m_content.MilitaryType(math.Round(mult * float64(m_content.MilitaryTypeMAX-1)))
}

func generatePopAgeType(fac Faction) m_content.FactionPopAgeType {
	ageType := m_content.FactionPopAgeType(rand.Intn(int(m_content.FactionPopAgeTypeMAX-1)) + 1)
	if ageType != m_content.FactionPopAgeTypeAdult { // other types should be rare
		ageType = m_content.FactionPopAgeType(rand.Intn(int(m_content.FactionPopAgeTypeMAX-1)) + 1)
	}
	return ageType
}

func generateBaseFaction(constr m_content.BaseConstraints, dungeonStrata int64, availableRaces []m_content.Race) *Faction {
	homeStratum := rand.Int63n(dungeonStrata) + 1

	fac := Faction{}
	fac.Races = generateFactionRaces(constr, &fac, availableRaces)
	fac.FactionName = generateFactionName(constr, fac, float64(homeStratum))
	fac.TechLevel = m_content.TechLevel(1 + math.Round(util.Clamp( // 1 + round(techFactor * maxTech)
		float64(fac.Races[0].TechFactor)*float64(constr.MaxTech-1),
		0, float64(constr.MaxTech)-1)))
	fac.MagicLevel = m_content.MagicLevel(1 + math.Round(util.Clamp( // 1 + round(magicFactor * maxMagic)
		float64(fac.Races[0].MagicFactor)*float64(constr.MaxMagic-1),
		0, float64(constr.MaxMagic)-1)))
	fac.Aggressiveness = m_content.AggressionLevel(math.Round(util.Clamp( // round((hostilityFactor-0.5)*4 + avgAggression)
		((float64(fac.Races[0].HostilityFactor)-0.5)*4.0)+ // put hostilityFactor in range [-2, 2]
			float64(constr.AvgAggression), // add it to avgAggression
		1, float64(m_content.AggressionLevelMAX)-1)))
	for _, race := range fac.Races {
		fac.Population += 1 + // 1 + round(rand(1000) * popFactor * (1 + len(races)/6))
			int64(math.Round(float64(rand.Intn(1000))*race.PopFactor))* // pop factor * 1000
				int64(constr.DungeonSize) // larger dungeons means more population
	}
	fac.HomeStrata = []int64{homeStratum}
	fac.Ethics = generateFactionEthics()
	fac.GovernmentType = generateFactionGovernmentType(&fac)
	fac.MilitaryType = generateFactionMilitaryType(&fac)
	fac.PopAgeType = generatePopAgeType(fac)

	return &fac
}

func GenerateBaseFactions(constr m_content.BaseConstraints) []*Faction {
	logging.Info("Generating base factions with constraints: %v", constr)

	availableRaces := []m_content.Race{}
	for _, pack := range constr.EnabledContent {
		availableRaces = append(availableRaces, pack.Races...)
	}

	if len(availableRaces) == 0 {
		logging.Error("No availableRaces, returning empty list")
		return []*Faction{}
	}

	// AMOUNT OF FACTIONS
	amountOfFactions := util.Clamp(
		(math.Pow(2.0, float64(constr.DungeonSize)))* // A larger dungeon size means more factions
			float64(rand.Intn(3)+1), // Random factor (1-3)
		1, math.MaxFloat64)
	amountOfFactions = math.Round(amountOfFactions)

	dungeonStrata := int64(math.Round(math.Pow(2.0, float64(constr.DungeonSize)))) // the amount of floors in the dungeon

	// GENERATE FACTIONS
	factions := make([]*Faction, int(amountOfFactions))
	for i := 0; i < int(amountOfFactions); i++ {
		factions[i] = generateBaseFaction(constr, dungeonStrata, availableRaces)
	}

	historyRunner := NewHistoryRunner(constr, dungeonStrata, availableRaces)
	firstStageHistoryFactions := []*Faction{}
	for _, faction := range factions {
		// 10 centuries of history
		for i := 0; i < 10; i++ {
			if faction.IsFallen {
				break
			}
			if rand.Float64() > 0.3 { // random chance to skip history
				firstStageHistoryFactions = append(firstStageHistoryFactions, historyRunner.RunHistory(faction, int64(i))...)
			}
		}
		firstStageHistoryFactions = append(firstStageHistoryFactions, faction)
	}

	// another 10 centuries of history
	secondStageHistoryFactions := []*Faction{}
	for _, faction := range firstStageHistoryFactions {
		for i := 10; i < 20; i++ {
			if faction.IsFallen {
				break
			}
			if rand.Float64() > 0.3 { // random chance to skip history
				secondStageHistoryFactions = append(secondStageHistoryFactions, historyRunner.RunHistory(faction, int64(i))...)
			}
		}
		secondStageHistoryFactions = append(secondStageHistoryFactions, faction)
	}

	return secondStageHistoryFactions
}
