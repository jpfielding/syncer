package cmd

import (
	"github.com/spf13/cobra"
)

func NewGooglePhotos() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "google-photoos",
		Short: "google-photoos - a CLI to manage google photos",
		Long:  `google-photoos is a photos syncing CLI`,
		Run: func(cmd *cobra.Command, args []string) {
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
	return cmd
}
