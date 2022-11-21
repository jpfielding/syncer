package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGooglePhotos() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "google-photos",
		Short: "google-photos - a CLI to manage google photos",
		Long:  `google-photos is a photos syncing CLI`,
	}
	cmd.AddCommand(NewGooglePhotosAuth())
	cmd.AddCommand(NewGooglePhotosList())
	return cmd
}

func NewGooglePhotosAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "auth - ",
		Long:  `auth`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("google photos auth\n")
			// ctx := context.Background()
			// oc := oauth2Config{
			// 	// ClientID:     "... your application Client ID ...",
			// 	// ClientSecret: "... your application Client Secret ...",
			// 	// // ...
			// }
			// tc := oc.Client(ctx, "... your user Oauth Token ...")
			// client := googlepho.NewClient(tc)
		},
	}
	cmd.AddCommand(NewGooglePhotosList())
	return cmd
}

func NewGooglePhotosList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list - ",
		Long:  `list `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("google photos list\n")
		},
	}
	return cmd
}

func NewGooglePhotosSync() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "sync - ",
		Long:  `sync `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("google photos sync\n")
		},
	}
	return cmd
}
