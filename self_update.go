package main

import (
	"fmt"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func pullUpdates(s *discordgo.Session, i *discordgo.InteractionCreate) (string, bool) {
	if !(interactionHasRole(i, PPLCORERoleID)) {
		return ErrorMessageNeedRole, false
	}

	fmt.Println("Update command. Fetching...")
	_, err := exec.Command("git", "fetch").Output()
	if err != nil {
		fmt.Println("Error while fetching repository, ", err.Error())
		return err.Error(), true
	}

	fmt.Println("Pulling changes...")
	output, err := exec.Command("git", "pull").Output()
	if err != nil {
		fmt.Println("Error while pulling repository, ", err.Error())
		return err.Error(), true
	}
	fmt.Println(string(output))

	return "Changes pulled, bot is restarting.", true
}
