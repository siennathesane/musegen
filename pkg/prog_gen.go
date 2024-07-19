package pkg

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"text/tabwriter"
)

// ProgressionKind determines the kind of progression
type ProgressionKind uint16

const (
	BluesKind ProgressionKind = iota
	BossaNovaKind
	ClassicalKind
	CountryKind
	DiscoKind
	EDMKind
	FlamencoKind
	FolkKind
	FunkKind
	GospelKind
	JazzKind
	MetalKind
	PopKind
	ProgressiveRockKind
	PunkKind
	ReggaeKind
	RnBKind
	RockKind
	SalsaKind
	SambaKind
	SoulKind
	endOfProgressionKinds
)

// String returns the string representation of the ProgressionKind.
func (pk ProgressionKind) String() string {
	switch pk {
	case BluesKind:
		return "Blues"
	case BossaNovaKind:
		return "Bossa Nova"
	case ClassicalKind:
		return "Classical"
	case CountryKind:
		return "Country"
	case DiscoKind:
		return "Disco"
	case EDMKind:
		return "EDM"
	case FlamencoKind:
		return "Flamenco"
	case FolkKind:
		return "Folk"
	case FunkKind:
		return "Funk"
	case GospelKind:
		return "Gospel"
	case JazzKind:
		return "Jazz"
	case MetalKind:
		return "Metal"
	case PopKind:
		return "Pop"
	case ProgressiveRockKind:
		return "Progressive Rock"
	case PunkKind:
		return "Punk"
	case ReggaeKind:
		return "Reggae"
	case RnBKind:
		return "R&B"
	case RockKind:
		return "Rock"
	case SalsaKind:
		return "Salsa"
	case SambaKind:
		return "Samba"
	case SoulKind:
		return "Soul"
	default:
		return fmt.Sprintf("Unknown ProgressionKind(%d)", int(pk))
	}
}

func ParseProgressionKind(genreStr string) (ProgressionKind, error) {
	genreMap := GetGenreMap()
	genre, ok := genreMap[genreStr]
	if !ok {
		return 0, fmt.Errorf("invalid genre: %s", genreStr)
	}
	return genre, nil
}

func GetGenreMap() map[string]ProgressionKind {
	genreMap := make(map[string]ProgressionKind)
	for i := 0; i < int(endOfProgressionKinds); i++ {
		genreMap[ProgressionKind(i).String()] = ProgressionKind(i)
	}
	return genreMap
}

var UniqueCombinationPotential = (uint(endOfProgressionKinds) - 1) * (uint(endofScaleKinds) - 1) * (uint(endOfKeyKind) - 1)

func GenerateRandomProgressionConfig() ProgressionConfig {
	return ProgressionConfig{
		Key:            KeyKind(rand.IntN(int(endOfKeyKind))),
		Scale:          ScaleKind(rand.IntN(int(endofScaleKinds))),
		Genre:          ProgressionKind(rand.IntN(int(endOfProgressionKinds))),
		AllowWeirdShit: false,
	}
}

func GetRandomProgression() ProgressionKind {
	return ProgressionKind(rand.IntN(int(endOfProgressionKinds)))
}

// ProgressionConfig represents the configuration for generating a chord progression
type ProgressionConfig struct {
	Key            KeyKind
	Scale          ScaleKind
	Genre          ProgressionKind
	AllowWeirdShit bool
}

func GenerateProgression(config ProgressionConfig) (Progression, error) {
	// Validate input
	if err := validateConfig(config); err != nil {
		return Progression{}, err
	}

	// Select base progression based on genre
	baseProgression := selectBaseProgression(config.Genre)

	// Create a new progression
	newProgression := Progression{
		Name: fmt.Sprintf("%s progression in %s %s",
			config.Genre.String(),
			config.Key,
			config.Scale),
		Chords: make([]Chord, len(baseProgression.Chords)),
	}

	// Modify chords based on key and scale
	for i, baseChord := range baseProgression.Chords {
		newChord := Chord{
			Base:       (baseChord.Base + int(config.Key.GetOffset())) % 12,
			Extensions: make([]ExtensionKind, 0),
		}

		// Adjust chord quality based on scale
		quality := getChordQuality(baseChord.Base, config.Scale)

		// Add the quality extension if it's not NoExtension
		if quality != NoExtension {
			newChord.Extensions = append(newChord.Extensions, quality)
		}

		// Add any additional extensions from the base chord
		for _, ext := range baseChord.Extensions {
			if ext != NoExtension && ext != quality {
				newChord.Extensions = append(newChord.Extensions, ext)
			}
		}

		newProgression.Chords[i] = newChord
	}

	return newProgression, nil
}

func validateConfig(config ProgressionConfig) error {
	// Validate Key
	if config.Key < 0 || config.Key > 11 {
		return fmt.Errorf("invalid key: %d. Must be between 0 and 11", config.Key)
	}

	// Validate Genre
	validGenre := false
	for _, progType := range AllProgressionTypes {
		if progType.Name == config.Genre.String() {
			validGenre = true
			break
		}
	}
	if !validGenre {
		return fmt.Errorf("invalid genre: %s", config.Genre)
	}

	// Check for incompatible combinations
	// For example, certain genres might not work well with certain scales
	if (config.Genre == FlamencoKind && config.Scale != PhrygianScale) && !config.AllowWeirdShit {
		return fmt.Errorf("flamenco progressions typically use the Phrygian mode")
	}

	// Add more specific validation rules as needed

	return nil
}

func selectBaseProgression(genre ProgressionKind) Progression {
	// Find the ProgressionType that matches the given genre
	var selectedType ProgressionType
	for _, progType := range AllProgressionTypes {
		if progType.Name == genre.String() {
			selectedType = progType
			break
		}
	}

	// If no matching genre was found, return an empty progression
	// This shouldn't happen if we've properly validated the input
	if selectedType.Name == "" {
		return Progression{}
	}

	// Randomly select one of the progressions from the chosen type
	return selectedType.Progressions[rand.IntN(len(selectedType.Progressions))]
}

// PrintProgressionTypes prints all the built-in progression types
func PrintProgressionTypes() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Genre\tProgression Name\tChord Progression In C Major")
	fmt.Fprintln(w, "-----\t----------------\t----------------------------")

	for _, progType := range AllProgressionTypes {
		for _, prog := range progType.Progressions {
			chordStr := make([]string, len(prog.Chords))
			for i, chord := range prog.Chords {
				extensions := ""
				if len(chord.Extensions) > 0 {
					extStrings := make([]string, len(chord.Extensions))
					for j, ext := range chord.Extensions {
						extStrings[j] = ext.String()
					}
					extensions = strings.Join(extStrings, "")
				}
				chordStr[i] = fmt.Sprintf("%s%s", noteNames[chord.Base], extensions)
			}
			fmt.Fprintf(w, "%s\t%s\t%s\n", progType.Name, prog.Name, strings.Join(chordStr, " - "))
		}
		//fmt.Fprintln(w) // Add an empty line between genres
	}

	w.Flush()
}

func PrintProgressionsJSON() {
	// Create a map to hold our progression types
	progressionMap := make(map[string]ProgressionType)

	// Populate the map
	for _, pt := range AllProgressionTypes {
		progressionMap[pt.Name] = pt
	}

	// Marshal the map to JSON with indentation
	jsonData, err := json.MarshalIndent(progressionMap, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}

// NoteNames maps the base note numbers to their names
var noteNames = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
