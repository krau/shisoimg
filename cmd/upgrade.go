package cmd

import (
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
)

var Version = "dev"

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"up"},
	Short:   "Upgrade shisoimg to the latest version",
	Run: func(cmd *cobra.Command, args []string) {

		v := semver.MustParse(Version)
		latest, err := selfupdate.UpdateSelf(v, "krau/shisoimg")
		if err != nil {
			log.Println("Binary update failed:", err)
			return
		}
		if latest.Version.Equals(v) {
			log.Println("Current binary is the latest version", Version)
		} else {
			log.Println("Successfully updated to version", latest.Version)
			log.Println("Release note:\n", latest.ReleaseNotes)
		}
	},
}
