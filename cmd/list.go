/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all blog in posts",
	Long:  `list all blog in posts `,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("posts")
		if os.IsNotExist(err) {
			log.Fatal("posts dir is not exits")
		}
		dirs, err := ioutil.ReadDir("posts")
		if err != nil {
			log.Fatal("read posts dir fail")
		}
		fmt.Println("------------------")
		for _, dir := range dirs {
			fmt.Printf(" %s\n", dir.Name())
		}
		fmt.Println("------------------")
		fmt.Printf("total: %d blog\n", len(dirs))

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
