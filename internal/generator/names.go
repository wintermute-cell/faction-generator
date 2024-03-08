package generator

import (
	"fmt"
	"math/rand"
)

var prefixes = []string{
	"Black", "White", "Crimson", "Emerald", "Sapphire", "Twilight", "Dawn",
	"Dusk", "Moon", "Sun", "Blood", "Bone", "Ghost", "Iron", "Steel",
	"Thunder", "Storm", "Ice", "Shadow", "Dark", "Radiant", "Mystic", "Arcane",
	"Elder", "Ancient", "Lost", "Forgotten", "Hidden", "Secret", "Silent",
	"Roaring", "Burning", "Frozen", "Shrouded", "Celestial", "Infernal",
}

var coreWords = []string{
	"Guardians", "Wardens", "Sentinels", "Defenders", "Keepers",
	"Warriors", "Mages", "Sorcerers", "Witches", "Priests",
	"Knights", "Paladins", "Rangers", "Druids", "Necromancers",
	"Thieves", "Assassins", "Berserkers", "Clerics", "Monks",
	"Artificers", "Enchanters", "Seers", "Prophets", "Harbingers",
	"Marauders", "Raiders", "Reavers", "Seekers", "Explorers",
	"Adventurers", "Scholars", "Sages", "Oracles", "Acolytes",
	"Cultists", "Zealots", "Mystics", "Shamans", "Elementalists",
}

var suffixes = []string{
	"of the North", "of the South", "of the East", "of the West",
	"of the High Tower", "of the Deep Woods", "of the Sacred Grove", "of the Crystal Cave",
	"of the Endless Desert", "of the Eternal Winter", "of the Scorching Sun", "of the Full Moon",
	"of the New Dawn", "of Twilight's Veil", "of the Ancient Ruins", "of the Forgotten Realms",
	"of the Shadow Realm", "of the Infernal Depths", "of the Celestial Spheres", "of the Mystic River",
	"of the Dragon's Lair", "of the Phoenix Nest", "of the Leviathan's Abyss", "of the Kraken's Maw",
	"of the Cursed Land", "of the Blessed Isles", "of the Skyward Peaks", "of the Starry Void",
	"of the Boundless Sea", "of the Roaring Thunder", "of the Whispering Winds", "of the Shifting Sands",
	"of the Eternal Flame", "of the Frozen Wastes", "of the Verdant Forest", "of the Stony Path",
	"of the Iron Fortress", "of the Golden City", "of the Silver Stream", "of the Obsidian Gate",
}

func factionNameGenerator() string {
	prefix := prefixes[rand.Intn(len(prefixes))]
	coreWord := coreWords[rand.Intn(len(coreWords))]
	suffix := suffixes[rand.Intn(len(suffixes))]
	return fmt.Sprintf("%s %s %s", prefix, coreWord, suffix)
}
