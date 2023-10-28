package main

import (
	"net/url"

	"github.com/bwmarrin/discordgo"
)

func handleCommentSetHiddenStatus(session *discordgo.Session, i *discordgo.InteractionCreate, UUID string, hide bool) string {
	if !(interactionHasRole(i, CommentManagerRoleID)) {
		return ErrorMessageNeedRole
	}
	var action string
	if hide {
		action = "hide"
	} else {
		action = "unhide"
	}

	return getResultOfHTTPQuery("api/hide-comment", url.Values{"UUID": {UUID}, "action": {action}})
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
