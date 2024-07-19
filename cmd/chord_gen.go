package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"musgen/pkg"
	"os"
)

func NewProgressionGeneratorCmd() *cobra.Command {
	var (
		key            string
		scale          string
		genre          string
		allowWeirdShit bool
		randomConfig   bool
	)

	var rootCmd = &cobra.Command{
		Use:   "gen-chords",
		Short: "chord progression generator",
		Long: `generate chord progressions based on specified key, scale, and genre.

if a value is left unspecified, a random value will be chosen. run "musegen list-config" to see the available options.`,
		Run: runProgression,
	}

	rootCmd.Flags().StringVarP(&key, "key", "k", "", "key of the progression (e.g., C, F#, Bb)")
	rootCmd.Flags().StringVarP(&scale, "scale", "s", "", "scale type (e.g., Major, NaturalMinor, Dorian)")
	rootCmd.Flags().StringVarP(&genre, "genre", "g", "", "genre of the progression (e.g., Pop, Jazz, Blues)")
	rootCmd.Flags().BoolVarP(&allowWeirdShit, "allow-weird-shit", "a", false, "allow the generator to break some default rules")

	rootCmd.Flags().BoolVarP(&randomConfig, "random", "r", false, fmt.Sprintf("generate a completely random progression (%d possibilities)", pkg.UniqueCombinationPotential))

	return rootCmd
}

func runProgression(cmd *cobra.Command, args []string) {
	key, err := cmd.Flags().GetString("key")
	cobra.CheckErr(err)

	var keyType pkg.KeyKind
	if key == "" {
		keyType = pkg.GetRandomKey()
	} else {
		keyType, err = pkg.ParseKey(key)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	scale, err := cmd.Flags().GetString("scale")
	cobra.CheckErr(err)

	var scaleType pkg.ScaleKind
	if scale == "" {
		scaleType = pkg.GetRandomScale()
	} else {
		scaleType, err = pkg.ParseScale(scale)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	genre, err := cmd.Flags().GetString("genre")
	cobra.CheckErr(err)

	var genreType pkg.ProgressionKind
	if genre == "" {
		genreType = pkg.GetRandomProgression()
	} else {
		genreType, err = pkg.ParseProgressionKind(genre)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	allowWeirdShit, err := cmd.Flags().GetBool("allow-weird-shit")
	cobra.CheckErr(err)

	config := pkg.ProgressionConfig{
		Key:            keyType,
		Scale:          scaleType,
		Genre:          genreType,
		AllowWeirdShit: allowWeirdShit,
	}

	isRandom, err := cmd.Flags().GetBool("random")
	cobra.CheckErr(err)

	if isRandom {
		fmt.Println("--random was set, overriding --key, --genre, & --scale flags, allowing weird shit if set!")
		config = pkg.GenerateRandomProgressionConfig()
		config.AllowWeirdShit = allowWeirdShit
	}

	progression, err := pkg.GenerateProgression(config)
	if err != nil {
		fmt.Println("error generating progression:", err)
		os.Exit(1)
	}

	printProgression(progression)
}
