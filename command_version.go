package main

import (
	"fmt"

	"github.com/SumDeusVitae/cli-assistant-client/internal/version"
)

func callbackVer(cfg *config, args ...string) error {
	// Checking if update required
	info := check(cfg.Variables.Version)
	fmt.Println("Current version: ", info)

	return nil
}

func check(ver string) string {
	info := version.FetchUpdateInfo(ver)
	defer info.PromptUpdateIfAvailable()
	return info.CurrentVersion
}

func outdated() error {
	fmt.Println("Update required")
	fmt.Println("QS CLI is outdated!")
	fmt.Println("please run: qs update")
	return nil
}
