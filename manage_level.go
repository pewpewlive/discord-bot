package main

import (
	"log"
	"net/url"

	"github.com/bwmarrin/discordgo"
)

func handleLevelApproveCommand(session *discordgo.Session, i *discordgo.InteractionCreate, levelUUID string) string {
	if !(interactionHasRole(i, LevelManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	return getResultOfHTTPQuery("api/approve-level", url.Values{"level_uuid": {levelUUID}})
}

func handleLevelDelistCommand(session *discordgo.Session, i *discordgo.InteractionCreate, levelUUID string) string {
	if !(interactionHasRole(i, LevelManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	return getResultOfHTTPQuery("api/delist-level", url.Values{"level_uuid": {levelUUID}})
}

func handleLevelPendingCommand(session *discordgo.Session, i *discordgo.InteractionCreate) string {
	if !(interactionHasRole(i, LevelManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	return getResultOfHTTPQuery("api/pending-levels", url.Values{})
}

func handleLevelListCommand(session *discordgo.Session, i *discordgo.InteractionCreate) string {
	if !(interactionHasRole(i, LevelManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	returnValue := getResultOfHTTPQuery("api/list-unreleased-levels", url.Values{})
	log.Printf("%s", returnValue)

	return returnValue
}
