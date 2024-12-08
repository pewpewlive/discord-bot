package main

import (
	"net/url"

	"github.com/bwmarrin/discordgo"
)

func handleCommentSetHiddenStatus(session *discordgo.Session, i *discordgo.InteractionCreate, UUID string) string {
	if !(interactionHasRole(i, CommentManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	return getResultOfHTTPQuery("api/hide-comment", url.Values{"UUID": {UUID}})
}

func handleCensorComment(session *discordgo.Session, i *discordgo.InteractionCreate, UUID string) string {
	if !(interactionHasRole(i, CommentManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	return getResultOfHTTPQuery("api/censor-comment", url.Values{"UUID": {UUID}})
}

func handleCommentListPending(session *discordgo.Session, i *discordgo.InteractionCreate) string {
	if !(interactionHasRole(i, CommentManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	return getResultOfHTTPQuery("api/pending-comments", url.Values{})
}

func handleModerationActionOnAccount(session *discordgo.Session, i *discordgo.InteractionCreate, accountUUID string, action string, mod discordgo.User) string {
	if !(interactionHasRole(i, CommentManagerRoleID)) {
		return ErrorMessageNeedRole
	}

	params := url.Values{
		"UUID":   {accountUUID},
		"action": {action},
		"mID":    {mod.ID},
		"mName":  {mod.Username}}

	return getResultOfHTTPQuery("api/moderate-account", params)
}
