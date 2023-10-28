package main

import (
	embed "github.com/Clinet/discordgo-embed"
	"github.com/bwmarrin/discordgo"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func interactionHasRole(i *discordgo.InteractionCreate, role string) bool {
	return stringInSlice(role, i.Member.Roles)
}

func respondsWithMessageOrAck(s *discordgo.Session, i *discordgo.InteractionCreate, fn func() string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsLoading,
		},
	})

	e := embed.NewGenericEmbed("", fn())
	embeds := make([]*discordgo.MessageEmbed, 0)
	embeds = append(embeds, e)

	s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		// Content: message,
		Embeds: &embeds,
	})
}
