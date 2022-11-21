package cmd

import (
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "syncer",
		Short: "syncer - a CLI to manage cross cloud photos",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	cmd.AddCommand(
		// NewAmazonPhotosCmd(),
		// NewAwsS3Cmd(),
		// NewFlickrCmd(),
		// NewGoogleDriveCmd(),
		NewGooglePhotos(),
	)
	return cmd
}
