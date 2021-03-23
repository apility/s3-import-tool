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
		bucket, err := cmd.Flags().GetString("bucket-name")
		if err != nil {
			return
		}
		if bucket == "" {
			err = fmt.Errorf("Bucket name not set")
			return
		}
		delete, err := cmd.Flags().GetBool("delete")
		if err != nil {
			return
		}

		recursive, err := cmd.Flags().GetBool("recursive")
		if err != nil {
			return
		}

		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return
		}

		dryRun, err := cmd.Flags().GetBool("dry-run")
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
			if dryRun != true {
				src.UploadFile(target, config)
			}
		}
		return nil
	},
}

func appendUploadCommand(cmd *cobra.Command) {
	uploadCommand.Flags().StringP("bucket-name", "b", "", "Which bucket to upload files to")
	uploadCommand.Flags().BoolP("recursive", "r", false, "Recursively upload folders")
	uploadCommand.Flags().BoolP("delete", "d", false, "Delete file after upload")
	uploadCommand.Flags().BoolP("verbose", "v", false, "Verbose mode")
	uploadCommand.Flags().BoolP("dry-run", "y", false, "Dry run")
	cmd.AddCommand(uploadCommand)
}
