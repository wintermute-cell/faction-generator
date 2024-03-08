package content

//go:generate stringer -type=DungeonSizeLevel
//go:generate stringer -type=TechLevel
//go:generate stringer -type=MagicLevel
//go:generate stringer -type=AggressionLevel
//go:generate stringer -type=MultiRaceType
//go:generate stringer -type=GovernmentType
//go:generate stringer -type=MilitaryType
//go:generate stringer -type=FactionPopAgeType

// ==========================
//      CONSTRAINTS
// ==========================

type DungeonSizeLevel int32

const NoSelection = 0
const RandomSelection = -1

const (
	DungeonSizeLevelSmall   DungeonSizeLevel = 1 // a small cave, e.g. an animal den
	DungeonSizeLevelMedium  DungeonSizeLevel = 2 // a medium cave, e.g. a bandit hideout
	DungeonSizeLevelLarge   DungeonSizeLevel = 3 // a large structure, e.g. a dragon's lair or a castle
	DungeonSizeLevelGreater DungeonSizeLevel = 4 // a greater structure, e.g. a city or a fortress
	DungeonSizeLevelMega    DungeonSizeLevel = 5 // a mega structure, e.g. a kingdom or a continent
	DungeonSizeLevelMAX     DungeonSizeLevel = 6
)

type TechLevel int32

const (
	TechLevelPrePrimitive TechLevel = 1 // a pre-primitive faction, e.g. a group of animals
	TechLevelPrimitive    TechLevel = 2 // a primitive faction, e.g. a tribe
	TechLevelMedieval     TechLevel = 3 // a medieval faction, e.g. a kingdom
	TechLevelModern       TechLevel = 4 // a modern faction, e.g. a nation
	TechLevelFuturistic   TechLevel = 5 // a futuristic faction, e.g. a space empire
	TechLevelMAX          TechLevel = 6
)

type MagicLevel int32

const (
	MagicLevelNone   MagicLevel = 1 // a faction with no magic
	MagicLevelLow    MagicLevel = 2 // a faction with low magic, for example common religions rituals without immediate effects
	MagicLevelMedium MagicLevel = 3 // a faction with medium magic, for example magic that has a meaningful effect on everyday life
	MagicLevelHigh   MagicLevel = 4 // a faction with high magic, for example magic that can alter the entire lives of individuals
	MagicLevelEpic   MagicLevel = 5 // a faction with epic magic, for example magic that can alter the state of the entire world
	MagicLevelMAX    MagicLevel = 6
)

type AggressionLevel int32

const (
	AggressionLevelPeaceful   AggressionLevel = 1 // peaceful factions, diplomacy is guaranteed
	AggressionLevelNeutral    AggressionLevel = 2 // neutral factions, diplomacy and escalation are balanced
	AggressionLevelAggressive AggressionLevel = 3 // aggressive factions, diplomacy is possible, but escalation is likely
	AggressionLevelHostile    AggressionLevel = 4 // hostile factions, no hesitation to attack, e.g. "For the Horde"
	AggressionLevelWarlike    AggressionLevel = 5 // warlike factions, constant all out war, e.g. "Blood for the Blood God"
	AggressionLevelMAX        AggressionLevel = 6
)

type GovernmentType int32

const (
	GovernmentTypeDemocratic GovernmentType = 1 + iota
	GovernmentTypeOligarchic
	GovernmentTypeDictatorial
	GovernmentTypeTheocratic
	GovernmentTypeAnarchic
	GovernmentTypeImperial
	GovernmentTypeHivemind
	GovernmentTypeMAX
)

type MilitaryType int32

const (
	MilitaryTypeNone MilitaryType = 1 + iota
	MilitaryTypeMilitia
	MilitaryTypeMercenary
	MilitaryTypeProfessional
	MilitaryTypeConscript
	MilitaryTypeElite
	MilitaryTypeMAX
)

type FactionPopAgeType int32

const (
	FactionPopAgeTypeYoungling FactionPopAgeType = 1 + iota
	FactionPopAgeTypeAdult
	FactionPopAgeTypeElder
	FactionPopAgeTypeMAX
)

type BaseConstraints struct {
	EnabledContent  []*ContentPack   // the content pack to use, default is CoreContentPack
	ForbidMultiRace bool             // if multiple races in one faction are forbidden, default is false
	DungeonSize     DungeonSizeLevel // the size of the dungeon, default is DungeonSizeLevelLarge
	MaxTech         TechLevel        // the maximum tech level of the faction, default is TechLevelMedieval
	MaxMagic        MagicLevel       // the maximum magic level of the faction, default is MagicLevelMedium
	AvgAggression   AggressionLevel  // the average aggression level of the faction, default is AggressionLevelNeutral
}

type MultiRaceType int32

const (
	Domination MultiRaceType = 1 + iota
	Coexistence
)
