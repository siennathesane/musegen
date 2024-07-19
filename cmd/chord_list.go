package cmd

import (
	"fmt"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/spf13/cobra"
	"musgen/pkg"
	"reflect"
	"sort"
)

func NewChordGeneratorConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list-config",
		Short: "chord progression generator config options",
		Long:  `see all the different types of configurations available for generating chord progressions`,
		Run: func(_ *cobra.Command, _ []string) {

			// get, convert to string, then sort the keys
			keys := reflect.ValueOf(pkg.GetKeyMap()).MapKeys()
			var keysSlice []string
			for idx := range keys {
				keysSlice = append(keysSlice, keys[idx].String())
			}
			sort.Strings(keysSlice)

			// get, convert to string, then sort the scales
			scales := reflect.ValueOf(pkg.GetScaleMap()).MapKeys()
			var scalesSlice []string
			for idx := range scales {
				scalesSlice = append(scalesSlice, scales[idx].String())
			}
			sort.Strings(scalesSlice)

			// get, convert to string, then sort the genres
			genres := reflect.ValueOf(pkg.GetGenreMap()).MapKeys()
			var genresSlice []string
			for idx := range genres {
				genresSlice = append(genresSlice, genres[idx].String())
			}
			sort.Strings(genresSlice)

			l := list.New(
				"Keys", list.New(keysSlice),
				"Scales", list.New(scalesSlice),
				"Genres", list.New(genresSlice),
			)
			fmt.Println(l)
		},
	}
}

func NewListBuiltInProgressionsCmd() *cobra.Command {
	var json bool

	cmd := &cobra.Command{
		Use:   "list-progressions",
		Short: "list built-in progressions",
		Long:  `see all the different types of built-in chord progressions available`,
		Run: func(c *cobra.Command, _ []string) {
			isJson, err := c.Flags().GetBool("json")
			cobra.CheckErr(err)

			if isJson {
				pkg.PrintProgressionsJSON()
			} else {
				pkg.PrintProgressionTypes()
			}
		},
	}

	cmd.Flags().BoolVarP(&json, "json", "j", false, "output in JSON format")

	return cmd
}
