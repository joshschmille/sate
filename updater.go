package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func selfUpdate(slug string) error {
	selfupdate.EnableLog()

	previous := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(previous, slug)
	if err != nil {
		return err
	}

	if previous.Equals(latest.Version) {
		fmt.Println("Current binary is the latest version", version)
	} else {
		fmt.Println("Data Migration Started.")
		migrate(previous.String())
		fmt.Println("Update successfully done to version", latest.Version)
		fmt.Println("Release note:\n", latest.ReleaseNotes)
	}
	return nil
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: SATE [flags]\n")
	flag.PrintDefaults()
}

func migrate(previous string) {
	switch previous {
	case "0.9.1":
		fallthrough
	case "0.9.2":
		// Fix help.names data
	}
	fmt.Println("Data Migration Complete.")
}
