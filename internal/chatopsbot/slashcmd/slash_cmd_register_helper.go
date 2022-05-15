package slashcmd

import (
	"github.com/bwmarrin/discordgo"
)

// Add slash command and slash command handler to slash command manager
func RegisterAllSlashCommands(scm *SlashCommandManager) {
	// Slash command list
	slashCommands := map[string]*SlashCommand{
		"disgobot": {
			cmd: &discordgo.ApplicationCommand{
				Name:        "disgobot",
				Description: "123",
			},
			cmdHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Echo Test",
					},
				})
			},
		},
	}

	// Iterate all handlers and add to slash command manager
	for scName, sc := range slashCommands {
		scm.AddSlashCommand(scName, sc.cmd)
		scm.AddSlashCommandHandler(scName, sc.cmdHandler)
	}
}
