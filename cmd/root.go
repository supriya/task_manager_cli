/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/binary"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var tasks = []byte("tasks")

type Task struct {
	ID          int
	Description string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task_manager_cli",
	Short: "Task Management CLI tool",
	Long: `Task Management is a CLI tool to add, list,
	and do the tasks.`,
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("tasks.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return db, err
	}
	return db, nil
}

func generateTaskDescription(args []string) string {
	taskDescription := ""
	for i, val := range args {
		if i == 0 {
			taskDescription = val
		} else {
			taskDescription += " " + val
		}
	}
	return taskDescription
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) uint64 {
	data := binary.BigEndian.Uint64(v)
	return data
}
