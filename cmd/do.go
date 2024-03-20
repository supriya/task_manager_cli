/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do command to do the task",
	Long:  `Do command marks the task done in the task list`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := initDB()
		if err != nil {
			log.Fatal("Cannot open the database!")
		}
		defer db.Close()
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(tasks)
			intVar, err := strconv.Atoi(args[0])
			err = b.Delete(itob(intVar))
			if err != nil {
				tx.Rollback()
				return err
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
