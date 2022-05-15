package slashcmd

import (
	"github.com/bwmarrin/discordgo"
	"guanyu.dev/chatopsbot/pkg/restapi/gitlab"
)

// Add slash command and slash command handler to slash command manager
func RegisterAllSlashCommands(scm *SlashCommandManager) {
	// Slash command list
	slashCommands := map[string]*SlashCommand{
		"disgobot": {
			cmd: &discordgo.ApplicationCommand{
				Name:        "disgobot",
				Description: "用於觸發disgobot專案內特定的Pipeline Job",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "branch",
						Description: "目標Branch",
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  "dev",
								Value: "dev",
							},
							{
								Name:  "main",
								Value: "main",
							},
						},
						Required: true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "job",
						Description: "目標Job",
						Choices: []*discordgo.ApplicationCommandOptionChoice{
							{
								Name:  "deploy",
								Value: "deploy",
							},
						},
						Required: true,
					},
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "version",
						Description: "Image版本",
						Required:    true,
					},
				},
			},
			cmdHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				// Create a map from options that provided by the user
				options := i.ApplicationCommandData().Options
				optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
				for _, opt := range options {
					optionMap[opt.Name] = opt
				}

				// Create a pipeline trigger service
				pts := gitlab.NewPipelineTriggerService()
				response, err := pts.TriggerPipeline(optionMap["branch"].StringValue(), optionMap["job"].StringValue(), map[string]string{
					"IMAGE_VERSION": optionMap["version"].StringValue(),
				})

				// Check the request success or not
				if err != nil || response.Status != "created" {
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "❌ Job觸發失敗",
						},
					})
					return
				}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "✅ 成功觸發Job",
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
