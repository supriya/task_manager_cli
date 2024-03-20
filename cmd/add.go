/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add command to add the task",
	Long:  `Add command adds the task to a list`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := initDB()
		if err != nil {
			log.Fatal("Cannot open the database!")
		}
		defer db.Close()
		taskDescription := generateTaskDescription(args)
		err = db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(tasks)
			if err != nil {
				return err
			}
			id, err := bucket.NextSequence()
			if err != nil {
				return err
			}
			idInInt := int(id)
			err = bucket.Put(itob(idInInt), []byte(taskDescription))

			if err != nil {
				return err
			}
			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
