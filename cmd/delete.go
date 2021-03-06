/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a post",
	Long:  `delete a post`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a color argument")
		}
		if strings.Contains(args[0], "/") || strings.Contains(args[0], "..") {
			return errors.New("posts name should not contain / or .. ")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fileName := "./posts/" + args[0]
		stat, err := os.Stat(fileName)
		if os.IsNotExist(err) {
			log.Fatalf("post %s is not exist", fileName)
		}
		if stat.IsDir() {
			log.Fatalf("%s is dir ,can not be deleted", fileName)
		}
		err = os.Remove(fileName)
		if err != nil {
			log.Fatalf("delete %s fail, err %v", fileName, err)
		} else {
			log.Printf("delete post %s success", fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
