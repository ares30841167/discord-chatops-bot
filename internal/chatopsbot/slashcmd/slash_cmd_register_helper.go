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
				Description: "disgobot專案管理",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Name:        "run",
						Description: "觸發專案內特定的Pipeline Job",
						Options: []*discordgo.ApplicationCommandOption{
							{
								Name:        "deploy",
								Description: "部屬disgobot容器",
								Type:        discordgo.ApplicationCommandOptionSubCommand,
								Options: []*discordgo.ApplicationCommandOption{
									{
										Type:        discordgo.ApplicationCommandOptionString,
										Name:        "env",
										Description: "部屬目標環境",
										Choices: []*discordgo.ApplicationCommandOptionChoice{
											{
												Name:  "dev",
												Value: "dev",
											},
											{
												Name:  "release",
												Value: "release",
											},
										},
										Required: true,
									},
									{
										Type:        discordgo.ApplicationCommandOptionString,
										Name:        "tag",
										Description: "目標Image標籤",
										Required:    true,
									},
								},
							},
						},
						Type: discordgo.ApplicationCommandOptionSubCommandGroup,
					},
				},
			},
			cmdHandler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				// Create a map from options that provided by the user
				options := i.ApplicationCommandData().Options[0].Options[0].Options
				optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
				for _, opt := range options {
					optionMap[opt.Name] = opt
				}

				// Create a pipeline trigger service
				pts := gitlab.NewPipelineTriggerService()
				response, err := pts.TriggerPipeline("main", "deploy", map[string]string{
					"ENV_TARGET": optionMap["env"].StringValue(),
					"IMAGE_TAG":  optionMap["tag"].StringValue(),
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
