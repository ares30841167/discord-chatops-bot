package slashcmd

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	// Check whether guild id exist or not
	if os.Getenv("GUILD_ID") == "" {
		fmt.Fprintln(os.Stderr, "No guild id provided.")
		os.Exit(1)
	}
}

type SlashCommandManager struct {
	commands           map[string]*discordgo.ApplicationCommand
	registeredCommands map[string]*discordgo.ApplicationCommand
	cmdHandlers        map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

// Initial a new slash command manager instance
func NewSlashCommandManager() *SlashCommandManager {
	return &SlashCommandManager{
		commands:           map[string]*discordgo.ApplicationCommand{},
		registeredCommands: map[string]*discordgo.ApplicationCommand{},
		cmdHandlers:        map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}
}

// Add a new slash command setting to commands
func (sc *SlashCommandManager) AddSlashCommand(cmdName string, command *discordgo.ApplicationCommand) {
	sc.commands[cmdName] = command
}

// Delete a slash command setting from commands
func (sc *SlashCommandManager) DelSlashCommand(cmdName string) {
	delete(sc.commands, cmdName)
}

// Add a new slash command handlers to cmdHandlers
func (sc *SlashCommandManager) AddSlashCommandHandler(cmdName string, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	sc.cmdHandlers[cmdName] = handler
}

// Delete a slash command handlers from cmdHandlers
func (sc *SlashCommandManager) DelSlashCommandHandler(cmdName string) {
	delete(sc.cmdHandlers, cmdName)
}

// Get all slash command handlers from cmdHandlers
func (sc *SlashCommandManager) GetSlashCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return sc.cmdHandlers
}

// Register all slash commands to the target guild in discord
func (sc *SlashCommandManager) RegisterSlashCommands(dg *discordgo.Session) error {
	for i, v := range sc.commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, os.Getenv("GUILD_ID"), v)
		if err != nil {
			return err
		}
		sc.registeredCommands[i] = cmd
	}
	return nil
}

// Deregister all slash commands from the target guild in discord
func (sc *SlashCommandManager) DeregisterSlashCommands(dg *discordgo.Session) error {
	for _, v := range sc.registeredCommands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, os.Getenv("GUILD_ID"), v.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
