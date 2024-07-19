package pkg

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type (
	// KeyKind represents the fundamental key (C, C#, D, etc.)
	KeyKind uint8

	// ScaleKind represents the type of scale
	ScaleKind uint8

	// ExtensionKind represents the type of chord extension
	ExtensionKind uint8
)

const (
	MajorScale ScaleKind = iota
	NaturalMinorScale
	HarmonicMinorScale
	MelodicMinorScale
	DorianScale
	PhrygianScale
	LydianScale
	MixolydianScale
	LocrianScale
	endofScaleKinds
)

// String returns the string representation of a ScaleKind.
func (st ScaleKind) String() string {
	switch st {
	case MajorScale:
		return "Major"
	case NaturalMinorScale:
		return "Natural Minor"
	case HarmonicMinorScale:
		return "Harmonic Minor"
	case MelodicMinorScale:
		return "Melodic Minor"
	case DorianScale:
		return "Dorian"
	case PhrygianScale:
		return "Phrygian"
	case LydianScale:
		return "Lydian"
	case MixolydianScale:
		return "Mixolydian"
	case LocrianScale:
		return "Locrian"
	default:
		return fmt.Sprintf("Unknown ScaleKind(%d)", uint8(st))
	}
}

func ParseScale(scaleStr string) (ScaleKind, error) {
	scaleMap := GetScaleMap()
	scale, ok := scaleMap[scaleStr]
	if !ok {
		return 0, fmt.Errorf("invalid scale: %s", scaleStr)
	}
	return scale, nil
}

func GetScaleMap() map[string]ScaleKind {
	scaleMap := make(map[string]ScaleKind)
	for i := 0; i < int(endofScaleKinds); i++ {
		scaleMap[ScaleKind(i).String()] = ScaleKind(i)
	}
	return scaleMap
}

func GetRandomScale() ScaleKind {
	return ScaleKind(rand.IntN(int(endofScaleKinds)))
}

const (
	C KeyKind = iota
	CSharp
	D
	DSharp
	E
	F
	FSharp
	G
	GSharp
	A
	ASharp
	B
	endOfKeyKind
)

// String method for KeyKind
func (k KeyKind) String() string {
	return [...]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}[k]
}

// GetOffset returns the offset of the key
func (k KeyKind) GetOffset() int {
	return keyInfo[k].Offset
}

// GetKeyMap returns a map of key strings to KeyKind
func GetKeyMap() map[string]KeyKind {
	keyMap := make(map[string]KeyKind)
	for i := 0; i < int(endOfKeyKind); i++ {
		info := keyInfo[KeyKind(i)]
		keyMap[info.Name] = KeyKind(i)
		keyMap[info.FlatName] = KeyKind(i)
	}
	return keyMap
}

// ParseKey converts a string representation to a KeyKind
func ParseKey(keyStr string) (KeyKind, error) {
	for kind, info := range keyInfo {
		if keyStr == info.Name || keyStr == info.FlatName {
			return kind, nil
		}
	}
	return KeyKind(0), fmt.Errorf("invalid key: %s", keyStr)
}

func GetRandomKey() KeyKind {
	return KeyKind(rand.IntN(int(endOfKeyKind)))
}

const (
	NoExtension ExtensionKind = iota
	MinorExtension
	MajorExtension
	DiminishedExtension
	AugmentedExtension
	Sus2Extension
	Sus4Extension
	Maj7Extension
	Min7Extension
	Dom7Extension
	Min7b5Extension
	Dim7Extension
	Aug7Extension
	Add9Extension
	Add11Extension
	Add13Extension
	Maj9Extension
	Dom9Extension
	Min9Extension
	Min11Extension
	Dom11Extension
	Maj13Extension
	Min13Extension
	Dom13Extension
)

// String returns the string representation of the ExtensionKind
func (e ExtensionKind) String() string {
	return [...]string{
		"",
		"min",
		"maj",
		"dim",
		"aug",
		"sus2",
		"sus4",
		"maj7",
		"min7",
		"7",
		"min7b5",
		"dim7",
		"aug7",
		"add9",
		"add11",
		"add13",
		"maj9",
		"9",
		"min9",
		"min11",
		"11",
		"maj13",
		"min13",
		"13",
	}[e]
}

// Key represents a musical key with its scale
type Key struct {
	Kind      KeyKind
	ScaleType ScaleKind
}

// KeyInfo provides additional information about a key
type KeyInfo struct {
	Name     string
	FlatName string
	Offset   int
}

// keyInfo maps KeyKind to its corresponding information
var keyInfo = map[KeyKind]KeyInfo{
	C:      {"C", "C", 0},
	CSharp: {"C#", "Db", 1},
	D:      {"D", "D", 2},
	DSharp: {"D#", "Eb", 3},
	E:      {"E", "E", 4},
	F:      {"F", "F", 5},
	FSharp: {"F#", "Gb", 6},
	G:      {"G", "G", 7},
	GSharp: {"G#", "Ab", 8},
	A:      {"A", "A", 9},
	ASharp: {"A#", "Bb", 10},
	B:      {"B", "B", 11},
}

// String method for Key
func (k Key) String() string {
	info := keyInfo[k.Kind]
	return fmt.Sprintf("%s %s", info.Name, k.ScaleType)
}

// AllKeys returns a slice of all possible keys with the given scale
func AllKeys(scale ScaleKind) []Key {
	keys := make([]Key, 12)
	for i := 0; i < 12; i++ {
		keys[i] = Key{Kind: KeyKind(i), ScaleType: scale}
	}
	return keys
}

// Chord represents a single chord with potential extensions or alterations
type Chord struct {
	Base       int             // Base chord (0 = I, 1 = ii, etc.)
	Extensions []ExtensionKind // Extensions like "7", "maj7", "min7", "dim", "aug", etc.
}

// String returns the string representation of the Chord
func (c Chord) String() string {
	notes := []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	baseNote := notes[c.Base%12]

	var extensions []string
	for _, ext := range c.Extensions {
		if ext != NoExtension {
			extensions = append(extensions, ext.String())
		}
	}

	if len(extensions) == 0 {
		return baseNote
	}

	return baseNote + strings.Join(extensions, "")
}

// getChordQuality determines the quality (major, minor, diminished) of a chord
// based on its position in the scale and the type of scale being used.
//
// Parameters:
//   - chordNumber: An integer representing the chord's position in the scale (0-based).
//   - scaleType: A ScaleKind value indicating the type of scale being used.
//
// Returns:
//
//	A string representing the chord quality:
//	- "" for major chords
//	- "min" for minor chords
//	- "dim" for diminished chords
//
// The function handles the following scale types:
//   - Major, Lydian, and Mixolydian: Use standard major scale chord qualities
//   - Natural Minor, Dorian, Phrygian, and Locrian: Use natural minor scale chord qualities
//   - Harmonic Minor: Similar to natural minor, but with a major V chord
//   - Melodic Minor: Similar to natural minor, but with major V and VI chords
//
// For any unrecognized scale type, an empty string is returned, defaulting to a major quality.
//
// Note: The function uses modulo 7 (%) to ensure that chord numbers outside
// the 0-6 range are mapped back to the correct position in the scale.
func getChordQuality(chordNumber int, scaleType ScaleKind) ExtensionKind {
	majorChords := []ExtensionKind{NoExtension, MinorExtension, MinorExtension, NoExtension, NoExtension, MinorExtension, DiminishedExtension}
	minorChords := []ExtensionKind{MinorExtension, DiminishedExtension, NoExtension, MinorExtension, MinorExtension, NoExtension, NoExtension}

	chordIndex := chordNumber % 7

	switch scaleType {
	case MajorScale, LydianScale, MixolydianScale:
		return majorChords[chordIndex]
	case NaturalMinorScale, DorianScale, PhrygianScale, LocrianScale:
		return minorChords[chordIndex]
	case HarmonicMinorScale:
		if chordIndex == 4 {
			return NoExtension // V chord in harmonic minor is major
		}
		return minorChords[chordIndex]
	case MelodicMinorScale:
		if chordIndex == 4 || chordIndex == 5 {
			return NoExtension // V and VI chords in melodic minor are major
		}
		return minorChords[chordIndex]
	default:
		return NoExtension
	}
}
