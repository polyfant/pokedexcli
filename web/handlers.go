package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"pokedex/internal/pokeapi"
	"strings"
	"time"
)

type PageData struct {
	Title         string
	PokemonList   []PokemonListItem
	PokemonDetail map[string]interface{}
	Error         string
}

type PokemonListItem struct {
	Name string
	URL  string
}

func NewHandler() *Handler {
	funcMap := template.FuncMap{
		"percentage": func(value, max int) int {
			return int(float64(value) / float64(max) * 100)
		},
	}

	tmpl := template.New("").Funcs(funcMap)
	tmpl = template.Must(tmpl.ParseGlob(filepath.Join("web", "templates", "*.html")))

	return &Handler{
		templates: tmpl,
		client:    pokeapi.NewClient(time.Hour, time.Hour),
	}
}

type Handler struct {
	templates *template.Template
	client    pokeapi.Client
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Pokédex",
	}
	h.templates.ExecuteTemplate(w, "base.html", data)
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		data := PageData{
			Title: "Pokédex - Search",
		}
		h.templates.ExecuteTemplate(w, "pokemon-list.html", data)
		return
	}

	// Use our client to fetch Pokemon
	pokemonResp, err := h.client.ListPokemon(nil)
	if err != nil {
		data := PageData{
			Title: "Pokédex - Error",
			Error: "Failed to fetch Pokemon",
		}
		h.templates.ExecuteTemplate(w, "pokemon-list.html", data)
		return
	}

	// Filter pokemon by query
	var filtered []PokemonListItem
	for _, p := range pokemonResp.Results {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(query)) {
			filtered = append(filtered, PokemonListItem{
				Name: p.Name,
				URL:  p.URL,
			})
		}
	}

	data := PageData{
		Title:       "Pokédex - Results",
		PokemonList: filtered,
	}
	h.templates.ExecuteTemplate(w, "pokemon-list.html", data)
}

func (h *Handler) PokemonDetails(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Pokemon name is required", http.StatusBadRequest)
		return
	}

	pokemon, err := h.client.GetPokemon(name)
	if err != nil {
		http.Error(w, "Failed to fetch Pokemon details", http.StatusInternalServerError)
		return
	}

	// Convert Pokemon struct to map for template
	pokemonMap := make(map[string]interface{})
	jsonData, err := json.Marshal(pokemon)
	if err != nil {
		http.Error(w, "Failed to process Pokemon details", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(jsonData, &pokemonMap)

	data := PageData{
		PokemonDetail: pokemonMap,
	}
	h.templates.ExecuteTemplate(w, "pokemon-details.html", data)
}
