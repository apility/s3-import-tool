package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootName string

func main() {
	rootCmd := &cobra.Command{
		Use:   os.Args[0],
		Short: "Upload files to S3 bucket",
		Long:  `A fast way to upload files to your s3 import netflex-bucket`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	appendUploadCommand(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\r\n", err)
		os.Exit(1)
	}

}
