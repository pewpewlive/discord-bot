package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/icza/gog"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "update",
			Description: "Check for updates",
		},
		{
			Name:        "fact",
			Description: "Returns a random fact about the PewPew games",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "index",
					Description: "The index of the fact",
					MinValue:    gog.Ptr(1.0),
				},
			},
		},
		{
			Name:        "player-fact",
			Description: "Returns a random fact about the PewPew playerbase",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "index",
					Description: "The index of the fact",
					MinValue:    gog.Ptr(1.0),
				},
			},
		},
		// Level management commands
		{
			Type: discordgo.MessageApplicationCommand,
			Name: "Approve Level",
		},
		{
			Name:        "levels",
			Description: "Commands concerning the management of levels",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "list",
					Description: "Lists a portion of public experimental (and or in review) levels",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "pending",
					Description: "Lists the levels that require a review",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "approve",
					Description: "Approves a level from the given ID",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "level-id",
							Description: "The ID of the level to approve",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "delist",
					Description: "Delists a level from the given ID",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "level-id",
							Description: "The ID of the level to delist",
							Required:    true,
						},
					},
				},
			},
		},
		// Comment management commands
		{
			Type: discordgo.MessageApplicationCommand,
			Name: "Hide Comment",
		},
		{
			Type: discordgo.MessageApplicationCommand,
			Name: "Unhide Comment",
		},
		{
			Type: discordgo.MessageApplicationCommand,
			Name: "Ban from Comments",
		},
		{
			Type: discordgo.MessageApplicationCommand,
			Name: "Unban from Comments",
		},
		{
			Name:        "comments",
			Description: "Commands concerning the management of comments",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "pending",
					Description: "Lists the comments that have been flagged and are awaiting moderation",
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "hide",
					Description: "Hides a comment from the given ID",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "comment-id",
							Description: "The ID of the comment",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "unhide",
					Description: "Unhides a comment from the given ID",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "comment-id",
							Description: "The ID of the comment",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "ban",
					Description: "Bans a user from further commenting, while also hiding their previous comments",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "account-id",
							Description: "The ID of the account",
							Required:    true,
						},
					},
				},
				{
					Type:        discordgo.ApplicationCommandOptionSubCommand,
					Name:        "unban",
					Description: "Unbans a user from not being able to comment, but keeps their previous comments hidden",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Type:        discordgo.ApplicationCommandOptionString,
							Name:        "account-id",
							Description: "The ID of the account",
							Required:    true,
						},
					},
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"update": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: discordgo.MessageFlagsLoading | discordgo.MessageFlagsEphemeral,
				},
			})

			output, shouldRestart := pullUpdates(s, i)
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &output,
			})

			if shouldRestart {
				WaitForClose <- true
			}
		},
		"fact": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			index := 0

			if len(i.ApplicationCommandData().Options) != 0 {
				index = int(i.ApplicationCommandData().Options[0].IntValue())
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: getFact(index),
				},
			})
		},
		"player-fact": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			index := 0

			if len(i.ApplicationCommandData().Options) != 0 {
				index = int(i.ApplicationCommandData().Options[0].IntValue())
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: getPlayerFact(index),
				},
			})
		},
		// Level management
		"Approve Level": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			resolvedMessage := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
			if len(resolvedMessage.Embeds) == 0 || resolvedMessage.Embeds[0].Footer == nil {
				respondsWithMessageOrAck(s, i, func() string { return "Invalid level review request message" })
				return
			}
			levelUUID := resolvedMessage.Embeds[0].Footer.Text
			respondsWithMessageOrAck(s, i, func() string { return handleLevelApproveCommand(s, i, levelUUID) })
		},
		"levels": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			subcommand := i.ApplicationCommandData().Options[0]

			switch subcommand.Name {
			case "list":
				respondsWithMessageOrAck(s, i, func() string { return handleLevelListCommand(s, i) })
			case "pending":
				respondsWithMessageOrAck(s, i, func() string { return handleLevelPendingCommand(s, i) })
			case "approve":
				levelUUID := subcommand.Options[0].StringValue()
				respondsWithMessageOrAck(s, i, func() string { return handleLevelApproveCommand(s, i, levelUUID) })
			case "delist":
				levelUUID := subcommand.Options[0].StringValue()
				respondsWithMessageOrAck(s, i, func() string { return handleLevelDelistCommand(s, i, levelUUID) })
			}
		},
		// Comment management
		"Hide Comment": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			resolvedMessage := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
			if len(resolvedMessage.Embeds) == 0 || resolvedMessage.Embeds[0].Footer == nil {
				respondsWithMessageOrAck(s, i, func() string { return "Invalid comment report message" })
				return
			}
			commentUUID := resolvedMessage.Embeds[0].Footer.Text
			respondsWithMessageOrAck(s, i, func() string { return handleCommentSetHiddenStatus(s, i, commentUUID, true) })
		},
		"Unhide Comment": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			resolvedMessage := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
			if len(resolvedMessage.Embeds) == 0 || resolvedMessage.Embeds[0].Footer == nil {
				respondsWithMessageOrAck(s, i, func() string { return "Invalid comment report message" })
				return
			}
			commentUUID := resolvedMessage.Embeds[0].Footer.Text
			respondsWithMessageOrAck(s, i, func() string { return handleCommentSetHiddenStatus(s, i, commentUUID, false) })
		},
		"Ban from Comments": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			resolvedMessage := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
			if len(resolvedMessage.Embeds) == 0 || resolvedMessage.Embeds[0].URL == "" {
				respondsWithMessageOrAck(s, i, func() string { return "Invalid comment report message" })
				return
			}
			accountUUID := strings.Split(resolvedMessage.Embeds[0].URL, "=")[1]
			respondsWithMessageOrAck(s, i, func() string { return handleModerationActionOnAccount(s, i, accountUUID, "ban", *i.Member.User) })
		},
		"Unban from Comments": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			resolvedMessage := i.ApplicationCommandData().Resolved.Messages[i.ApplicationCommandData().TargetID]
			if len(resolvedMessage.Embeds) == 0 || resolvedMessage.Embeds[0].URL == "" {
				respondsWithMessageOrAck(s, i, func() string { return "Invalid comment report message" })
				return
			}
			accountUUID := strings.Split(resolvedMessage.Embeds[0].URL, "=")[1]
			respondsWithMessageOrAck(s, i, func() string { return handleModerationActionOnAccount(s, i, accountUUID, "unban", *i.Member.User) })
		},
		"comments": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			subcommand := i.ApplicationCommandData().Options[0]

			switch subcommand.Name {
			case "pending":
				respondsWithMessageOrAck(s, i, func() string { return handleCommentListPending(s, i) })
			case "hide":
				commentUUID := subcommand.Options[0].StringValue()
				respondsWithMessageOrAck(s, i, func() string { return handleCommentSetHiddenStatus(s, i, commentUUID, true) })
			case "unhide":
				commentUUID := subcommand.Options[0].StringValue()
				respondsWithMessageOrAck(s, i, func() string { return handleCommentSetHiddenStatus(s, i, commentUUID, false) })
			case "ban":
				accountUUID := subcommand.Options[0].StringValue()
				respondsWithMessageOrAck(s, i, func() string { return handleModerationActionOnAccount(s, i, accountUUID, "ban", *i.Member.User) })
			case "unban":
				accountUUID := subcommand.Options[0].StringValue()
				respondsWithMessageOrAck(s, i, func() string { return handleModerationActionOnAccount(s, i, accountUUID, "unban", *i.Member.User) })
			}
		},
	}
)

func registerCommands(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	if WipeOldCommands {
		log.Printf("Deleting old commands\n")
		cmd, _ := s.ApplicationCommands(s.State.User.ID, GuildID)
		for _, v := range cmd {
			s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
		}
	}

	createdCommands, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, GuildID, commands)
	if err != nil {
		log.Panicf("Could not create command(s): %v", err)
	}

	for _, v := range createdCommands {
		log.Printf("Created new command(s) %v\n", v.Name)
	}
}
