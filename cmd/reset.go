/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/LSyaRE/lumines-cli/tools"

	"github.com/spf13/cobra"
)

// resetCmd represents the serve command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Permite actualizarse el proyecto a master",
	Long: `El comando permite la actualizacion de los ultimos cambios diponibles 
		   subidos en el repositorio sin necesidad de escribir tantos comandos, siguiendo los 
		   estandares de la empresa.`,
	Run: func(cmd *cobra.Command, args []string) {
		var branch string

		fmt.Println("Cual es la rama principal")
		fmt.Scanln(&branch)

		res, err := tools.GetAndFilterCommandOutput(branch,"git", "branch")

		if err != nil {
			return
		}

		if res == "" {
			fmt.Println("No se ha encontrado la rama especificada")
			return
		}

		fmt.Println(res)
		fmt.Println("La rama que sera origen sera actualizada segun la rama:", branch)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// resetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	resetCmd.Flags().BoolP("toggle", "t", true, "Help message for toggle")
}
