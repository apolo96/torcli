/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var surveyCmd = &cobra.Command{
	Use:   "survey",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("survey called")
		qs := []*survey.Question{
			{
				Name: "firstname",
				Prompt: &survey.Input{
					Message: "What is your first name?",
				},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name: "favoritecolor",
				Prompt: &survey.Select{
					Message: "What is your favorite color?",
					Options: []string{"red", "orange", "yellow", "green", "blue", "black", "white"},
					Default: "white",
				},
			},
			{
				Name: "story",
				Prompt: &survey.Multiline{
					Message: "Tell me a story",
				},
			},
			{
				Name: "secret",
				Prompt: &survey.Password{
					Message: "Write a secret",
				},
			},
			{
				Name: "good",
				Prompt: &survey.Confirm{
					Message: "Are you having a good day?",
				},
			},
			{
				Name: "favoritepies",
				Prompt: &survey.MultiSelect{
					Message: "What pies do you like",
					Options: []string{"Pumpkin", "Lemon Meringue", "Apple", "Key Lime", "Pecan"},
				},
			},
		}
		answers := struct {
			FirstName     string
			FavoriteColor string
			Story         string
			Secret        string
			Good          bool
			FavoritePies  []string
		}{}
		if err := survey.Ask(qs, &answers); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(answers)
	},
}

func init() {
	rootCmd.AddCommand(surveyCmd)
}
