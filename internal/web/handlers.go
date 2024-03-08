package web

import (
	"fmt"
	"net/http"
	"project_factions/internal/content"
	"project_factions/internal/generator"
	"project_factions/internal/logging"
	"project_factions/internal/util"
	"project_factions/internal/web/templates"
	m_content "project_factions/pkg/models/content"
	"strconv"
)

var PROJECT_NAME = "Project Factions"

type handler struct {
	// TODO: add handler state if necessary
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func handle404(w http.ResponseWriter, r *http.Request) {
	component := templates.Layout(PROJECT_NAME, templates.Error404())
	component.Render(r.Context(), w)
}

func (h *handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handle404(w, r)
	} else {
		component := templates.Layout(PROJECT_NAME, templates.Index())
		component.Render(r.Context(), w)
	}
}

type GenerateRequest struct {
	Packs           []string `json:"packs"`
	ForbidMultiRace string   `json:"forbidmultirace"`
	DungeonSize     int32    `json:"dungeonsize"`
	TechLevel       int32    `json:"techlevel"`
	MagicLevel      int32    `json:"magiclevel"`
	AggressionLevel int32    `json:"aggressionlevel"`
}

func (h *handler) HandleGenerate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logging.Error("Error parsing form: %v", err)
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	packs := r.Form["packs"]
	forbidMultiRaceStr := r.FormValue("forbidmultirace")
	dungeonSizeStr := r.FormValue("dungeonsize")
	techLevelStr := r.FormValue("techlevel")
	magicLevelStr := r.FormValue("magiclevel")
	aggressionLevelStr := r.FormValue("aggressionlevel")

	// Convert and default form values
	forbidMultiRace := forbidMultiRaceStr == "true"
	dungeonSize, err := strconv.ParseInt(dungeonSizeStr, 10, 32)
	if err != nil || dungeonSize == 0 {
		dungeonSize = 1 // Default to DungeonSizeLevelSmall
	}
	techLevel, err := strconv.ParseInt(techLevelStr, 10, 32)
	if err != nil || techLevel == 0 {
		techLevel = 1 // Default to TechLevelPrimitive
	}
	magicLevel, err := strconv.ParseInt(magicLevelStr, 10, 32)
	if err != nil || magicLevel == 0 {
		magicLevel = 1 // Default to MagicLevelNone
	}
	aggressionLevel, err := strconv.ParseInt(aggressionLevelStr, 10, 32)
	if err != nil || aggressionLevel == 0 {
		aggressionLevel = 1 // Default to AggressionLevelPeaceful
	}

	if len(packs) == 0 {
		packs = []string{"core"} // Default pack
	}

	// Validate packs
	for _, pack := range packs {
		if _, exists := content.ContentPacks[pack]; !exists {
			msg := fmt.Sprintf("Content pack %s does not exist", pack)
			logging.Error(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}
	}

	// Clamping enum values
	dungeonSize = int64(util.Clamp(int(dungeonSize), 1, int(m_content.DungeonSizeLevelMAX)-1))
	techLevel = int64(util.Clamp(int(techLevel), 1, int(m_content.TechLevelMAX)-1))
	magicLevel = int64(util.Clamp(int(magicLevel), 1, int(m_content.MagicLevelMAX)-1))
	aggressionLevel = int64(util.Clamp(int(aggressionLevel), 1, int(m_content.AggressionLevelMAX)-1))

	// Construct constraints struct
	constraints := content.NewBaseConstraints(
		packs,
		forbidMultiRace,
		m_content.DungeonSizeLevel(dungeonSize),
		m_content.TechLevel(techLevel),
		m_content.MagicLevel(magicLevel),
		m_content.AggressionLevel(aggressionLevel),
	)

	// Proceed with faction generation
	factions := generator.GenerateBaseFactions(constraints)
	ret := fmt.Sprintf("Generated %v factions", len(factions))
	for _, faction := range factions {
		ret += fmt.Sprintf("\n================================\n%v", faction.String())
	}
	fmt.Fprintf(w, ret)
}
