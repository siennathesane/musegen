package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"musgen/pkg"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "musegen",
		Short: "Music Generator",
		Long:  `A CLI tool to generate music-related content.`,
	}

	rootCmd.AddCommand(NewProgressionGeneratorCmd())
	rootCmd.AddCommand(NewChordGeneratorConfigCmd())
	rootCmd.AddCommand(NewListBuiltInProgressionsCmd())

	return rootCmd
}

func printProgression(prog pkg.Progression) {
	boldStyle := lipgloss.NewStyle().Bold(true)

	fmt.Println(boldStyle.Render(prog.Name))
	fmt.Print("Chords: ")
	for i, chord := range prog.Chords {
		if i > 0 {
			fmt.Print(" - ")
		}
		fmt.Print(chord)
	}
	fmt.Println()
}
