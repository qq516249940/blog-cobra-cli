/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new post",
	Long:  `create new post. For example:  学习golang cobra`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a color argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("args[0]: %v\n", args[0])
		fileName := "posts/" + args[0]
		err := os.Mkdir("posts", 644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			file, err := os.Create(fileName)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("create file %s", fileName)
			defer file.Close()
		} else {
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
