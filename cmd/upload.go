package main

import (
	"fmt"

	"github.com/apility/s3-import-tool/src"
	"github.com/spf13/cobra"
)

var uploadCommand = &cobra.Command{
	Use:   "upload",
	Short: "Upload files to S3 bucket",
	Long: `
Uploads files to S3.
Requires a bucket name and a glob that matches atleast one file.

Example:
	netflex-import upload -b sync-bucket-1 customers.csv
	
It is also possible to mention a range of files.

Example:
  netflex-import upload -b sync-bucket-1 Ships/*`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		bucket, err := cmd.PersistentFlags().GetString("bucket-name")
		if err != nil {
			return
		}
		if bucket == "" {
			err = fmt.Errorf("Bucket name not set")
			return
		}
		delete, err := cmd.PersistentFlags().GetBool("delete")
		if err != nil {
			return
		}

		recursive, err := cmd.PersistentFlags().GetBool("recursive")
		if err != nil {
			return
		}

		verbose, err := cmd.PersistentFlags().GetBool("verbose")
		if err != nil {
			return
		}

		config := src.Configuration{
			AWS: src.AWSConfig{
				Region: "eu-west-1",
			},
			Paths:              args,
			DeleteWhenUploaded: delete,
			RecursiveSearch:    recursive,
			BucketName:         bucket,
		}

		list, err := config.CreateTargetsList()
		if err != nil {
			return
		}

		for _, target := range list {
			if verbose {
				fmt.Println(target.ToString(config))
			}
			src.UploadFile(target, config)
		}
		return nil
	},
}

func appendUploadCommand(cmd *cobra.Command) {
	uploadCommand.PersistentFlags().StringP("bucket-name", "b", "", "Which bucket to upload files to")
	uploadCommand.PersistentFlags().BoolP("recursive", "r", false, "Recursively upload folders")
	uploadCommand.PersistentFlags().BoolP("delete", "d", false, "Delete file after upload")
	uploadCommand.PersistentFlags().BoolP("verbose", "v", false, "Verbose mode")

	cmd.AddCommand(uploadCommand)
}
