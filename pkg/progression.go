package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Progression represents a single chord progression
type Progression struct {
	Chords []Chord
	Name   string // Optional name or description of the progression
}

func (p Progression) String() string {
	chords := make([]string, len(p.Chords))
	for i, chord := range p.Chords {
		extensions := ""
		if len(chord.Extensions) > 0 {
			extStrings := make([]string, len(chord.Extensions))
			for j, ext := range chord.Extensions {
				extStrings[j] = ext.String()
			}
			extensions = strings.Join(extStrings, "")
		}
		chords[i] = fmt.Sprintf("%s%s", noteNames[chord.Base], extensions)
	}
	return fmt.Sprintf("%s: %s", p.Name, strings.Join(chords, " - "))
}

func (p Progression) MarshalJSON() ([]byte, error) {
	type Alias Progression
	return json.Marshal(&struct {
		Alias
		ChordsStr string `json:"chords_str"`
	}{
		Alias:     Alias(p),
		ChordsStr: p.String(),
	})
}

// ProgressionType represents a type of chord progression for a specific genre
type ProgressionType struct {
	Name         string
	Progressions []Progression
}

func (pt ProgressionType) String() string {
	progressions := make([]string, len(pt.Progressions))
	for i, prog := range pt.Progressions {
		progressions[i] = prog.String()
	}
	return fmt.Sprintf("%s:\n%s", pt.Name, strings.Join(progressions, "\n"))
}

func (pt ProgressionType) MarshalJSON() ([]byte, error) {
	type Alias ProgressionType
	return json.Marshal(&struct {
		Alias
		ProgressionsCount int `json:"progressions_count"`
	}{
		Alias:             Alias(pt),
		ProgressionsCount: len(pt.Progressions),
	})
}
