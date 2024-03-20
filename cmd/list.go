/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List command lists the tasks.",
	Long:  `List command lists the non finished tasks from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := initDB()
		if err != nil {
			log.Fatal("Cannot open the database!")
		}
		defer db.Close()
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket(tasks)
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("%v. %s\n", btoi(k), v)
			}

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
