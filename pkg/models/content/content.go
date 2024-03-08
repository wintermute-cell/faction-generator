package content

import (
	"fmt"
)

// ==========================
//       RACE & REGION
// ==========================
// NOTE: each model should implement the Stringer interface.

type RaceAlignment string

const (
	RaceAlignmentGood    RaceAlignment = "good"
	RaceAlignmentNeutral RaceAlignment = "neutral"
	RaceAlignmentEvil    RaceAlignment = "evil"
)

type Race struct {
	Name              string        // Name of the race, e.g., "Elves", "Dwarves"
	Alignment         RaceAlignment // The alignment of the race. Helps determine which races are likely to be allied.
	PreferredRegions  []string      // A list of preferred region types, for compatibility reasons, these may reference region types from other content packs.
	ProducedResources []string      // A list of resources the race *primarily* tends to produce.
	HostilityFactor   float64       // A factor between 0 and 1 indicating hostility level.
	// 0.0: Bunch of Stoned Hobbits (very peaceful)
	// 0.1: A group of Human Children (mostly harmless)
	// 0.2: A gathering of Peaceful Elves (harmonious and serene)
	// 0.3: A tribe of Nomadic Halflings (avoid conflict, but will defend themselves)
	// 0.4: A community of Farmers (defensive when threatened)
	// 0.5: A band of Highwaymen (opportunistic but not overly aggressive)
	// 0.6: A clan of Vikings (raiders, aggressive when raiding)
	// 0.7: A squad of Mercenaries (trained for combat, aggressive for pay)
	// 0.8: A legion of Imperial Soldiers (disciplined, aggressive on orders)
	// 0.9: A flight of Fire-breathing Dragons (territorial, extremely dangerous)
	// 1.0: An army from Hell (the epitome of hostility, destruction incarnate)

	TechFactor float64 // A factor between 0 and 1 indicating technological advancement.
	// 0.0: A group of Animals (no technology)
	// 0.1: Primitive Tribes (basic tools)
	// 0.2: Bronze Age Society (early metalworking)
	// 0.3: Iron Age Civilization (advanced metalworking)
	// 0.4: Medieval Kingdom (catapults, crossbows)
	// 0.5: Renaissance Era (gunpowder, early machinery)
	// 0.6: Industrial Age (factories, trains)
	// 0.7: Modern Era (computers, automobiles)
	// 0.8: Digital Age (internet, robotics)
	// 0.9: Near Future (advanced AI, space travel)
	// 1.0: Space Empire (interstellar travel, advanced alien technology)

	MagicFactor float64 // A factor between 0 and 1 indicating magical power.
	// 0.0: An Ordinary Person (no magic)
	// 0.1: Novice Mage (simple spells)
	// 0.2: Journeyman Wizard (elemental magic)
	// 0.3: Experienced Sorcerer (enchantment, divination)
	// 0.4: Master Enchanter (powerful curses and blessings)
	// 0.5: High Priest/Priestess (divine magic)
	// 0.6: Archmage (mastery over multiple schools of magic)
	// 0.7: Demigod (minor reality manipulation)
	// 0.8: Lesser Deity (major reality manipulation, creation)
	// 0.9: Major Deity (world-shaping powers)
	// 1.0: A God (omnipotent magical abilities)

	WeirdnessPreference float64 // A factor between 0 and 1 indicating preferred habitat depth.
	// 0.0: The Surface (forests, plains)
	// 0.1: Shallow Caves (near the surface, light penetrates)
	// 0.2: Surface Dungeons (just below the surface, man-made)
	// 0.3: Deep Woods (dense, dark forests)
	// 0.4: Mountainous Regions (high altitude, rugged terrain)
	// 0.5: Underground Caverns (natural, large cave systems)
	// 0.6: Deep Dungeons (far below the surface, ancient constructions)
	// 0.7: Subterranean Cities (built by intelligent beings, well below surface)
	// 0.8: Abyssal Plains (extreme depths, high pressure)
	// 0.9: Near Core Regions (close to the planetary core, extreme conditions)
	// 1.0: The Abyss (the deepest parts of the world, unknown and mysterious)

	PopFactor float64 // A factor between 0 and 1 indicating population density.
	// 0.0: An Individual (solitary)
	// 0.1: A Family (small, close-knit group)
	// 0.2: A Clan (extended family, small community)
	// 0.3: A Village (a small settlement)
	// 0.4: A Town (moderate-sized settlement)
	// 0.5: A City (large settlement, dense population)
	// 0.6: A Metropolis (very large city, very dense population)
	// 0.7: A Region (multiple cities and towns)
	// 0.8: A Country (nation-state, widespread population)
	// 0.9: A Continent (massive, diverse populations)
	// 1.0: A Large Group (covering vast areas, extremely populous)

	DominanceFactor float64 // A factor between 0 and 1 indicating dominance within the ecosystem.
	// 0.0: A Follower (takes orders, does not lead)
	// 0.1: A Minor Influencer (has some influence, but not dominant)
	// 0.2: A Community Leader (leads a small community)
	// 0.3: A Town Mayor (has significant influence over a town)
	// 0.4: A City Mayor (leads a large city)
	// 0.5: A Regional Governor (controls a region)
	// 0.6: A King/Queen (rules a kingdom)
	// 0.7: An Emperor/Empress (rules over multiple kingdoms)
	// 0.8: A Demigod (worshipped by many, has supernatural influence)
	// 0.9: A Lesser God (has followers across the world)
	// 1.0: A Supreme Deity (ultimate authority, controls the fate of worlds)
}

func (race *Race) String() string {
	return fmt.Sprintf("Race: %s\nAlignment: %v\nHostility: %.2f\nTech: %.2f\nMagic: %.2f\nDeepness: %.2f\nPop: %.2f\nDominance: %.2f",
		race.Name,
		race.Alignment,
		race.HostilityFactor,
		race.TechFactor,
		race.MagicFactor,
		race.WeirdnessPreference,
		race.PopFactor,
		race.DominanceFactor)
}

type RegionType struct {
	Name         string   `json:"name"`
	Resources    []string `json:"resources"`     // A list of resources available in the region. For compatibility reasons, these may reference resources from other content packs.
	DangerFactor float64  `json:"danger_factor"` // A factor between 0 and 1 indicating the danger level of the region.
	// 0.0: Peaceful Plains (very safe, e.g., a serene meadow)
	// 0.1: Gentle Forest (minimal danger, e.g., occasional wildlife)
	// 0.2: Quiet Hills (low danger, e.g., rare bandit sightings)
	// 0.3: Lonely Swamp (mild danger, e.g., tricky terrain and some predators)
	// 0.4: Abandoned Ruins (moderate danger, e.g., unstable structures, some monsters)
	// 0.5: Dense Jungle (elevated danger, e.g., venomous creatures, hidden threats)
	// 0.6: Dark Woods (high danger, e.g., aggressive beasts, malevolent spirits)
	// 0.7: Monster-infested Mountains (very high danger, e.g., dragons, giants)
	// 0.8: Cursed Lands (extreme danger, e.g., undead, curses, very hostile environment)
	// 0.9: Demon's Wasteland (near-apocalyptic danger, e.g., demonic presence, environmental hazards)
	// 1.0: The Nine Hells (ultimate danger, e.g., hellish landscape, demons, constant peril)

	MagicFactor float64 `json:"magic_factor"` // A factor between 0 and 1 indicating the prevalence of magic in the region.
	// 0.0: No Magic (mundane, e.g., a completely natural forest)
	// 0.1: Faint Magic (minimal magical influence, e.g., slight enhancements to flora and fauna)
	// 0.2: Mild Magic (low magic, e.g., occasional magical occurrences, minor enchanted areas)
	// 0.3: Enchanted Lands (moderate magic, e.g., magical creatures, enchanted springs)
	// 0.4: Mystic Woods (elevated magic, e.g., frequent spells, strong magical creatures)
	// 0.5: Arcane Ruins (high magic, e.g., remnants of powerful spells, portals)
	// 0.6: Sorcerer's Domain (very high magic, e.g., intense magic fields, constant spell effects)
	// 0.7: Elemental Planes (extreme magic, e.g., lands governed by elemental forces)
	// 0.8: Ley Line Nexus (near-epic magic, e.g., convergence of powerful magical energies)
	// 0.9: Ancient Gods' Battleground (epic magic, e.g., god-level enchantments, ancient powers)
	// 1.0: Epic Magic (the epitome of magical saturation, e.g., reality bending, creation, destruction)

	Weirdness float64 `json:"weirdness"` // A factor between 0 and 1 indicating the strangeness or otherworldliness of the region.
	// 0.0: Completely Mundane (ordinary, e.g., a typical village with no magical influence)
	// 0.1: Slightly Peculiar (minor oddities, e.g., an unusually high number of twins born in the area)
	// 0.2: Mildly Unusual (small anomalies, e.g., animals with peculiar colors or plants that glow faintly at night)
	// 0.3: Quaintly Quirky (charming oddities, e.g., trees that whisper when the wind blows)
	// 0.4: Notably Strange (noticeable anomalies, e.g., sporadic, unexplained weather changes)
	// 0.5: Moderately Bizarre (moderate weirdness, e.g., floating water droplets, areas where shadows move independently)
	// 0.6: Significantly Otherworldly (strong magical or alien characteristics, e.g., areas where time flows differently)
	// 0.7: Highly Anomalous (highly aberrant, e.g., landscapes that defy conventional physics, like inverted mountains)
	// 0.8: Extremely Eldritch (extreme weirdness, e.g., living architecture, sentient landscapes)
	// 0.9: Nearly Incomprehensible (nearing the peak of strangeness, e.g., spaces where thoughts manifest into reality)
	// 1.0: Realm of Pure Chaos (the epitome of weirdness, a realm where the laws of nature and magic cease to be comprehensible)
}

func (regionType *RegionType) String() string {
	return fmt.Sprintf("RegionType: %s\nDanger: %.2f\nMagic: %.2f\nDepth: %.2f\nResources: %v",
		regionType.Name,
		regionType.DangerFactor,
		regionType.MagicFactor,
		regionType.Weirdness,
		regionType.Resources)
}
