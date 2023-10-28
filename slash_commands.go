package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
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
		},
		// Level management commands
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
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: getRandomFact(),
				},
			})
		},
		// Level management
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

	// Delete old commands
	cmds, _ := s.ApplicationCommands(s.State.User.ID, GuildID)
	log.Printf("Deleting old commands")
	for _, v := range cmds {
		s.ApplicationCommandDelete(s.State.User.ID, GuildID, v.ID)
	}

	// Create new ones
	for _, v := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command(s): %v", v.Name, err)
		} else {
			log.Printf("Created new command(s) %v\n", v.Name)
		}
	}
}
