package slashcmd

import "github.com/bwmarrin/discordgo"

type SlashCommand struct {
	cmd        *discordgo.ApplicationCommand
	cmdHandler func(s *discordgo.Session, i *discordgo.InteractionCreate)
}
