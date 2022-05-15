package slashcmd

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestNewSlashCommandManager(t *testing.T) {
	scm := NewSlashCommandManager()

	if len(scm.commands) != 0 {
		t.Error("Commands are not empty")
	}

	if len(scm.registeredCommands) != 0 {
		t.Error("registeredCommands are not empty")
	}

	if len(scm.cmdHandlers) != 0 {
		t.Error("cmdHandlers are not empty")
	}
}

func TestAddSlashCommand(t *testing.T) {
	scm := &SlashCommandManager{
		commands:           map[string]*discordgo.ApplicationCommand{},
		registeredCommands: map[string]*discordgo.ApplicationCommand{},
		cmdHandlers:        map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}
	scm.AddSlashCommand("TEST", &discordgo.ApplicationCommand{
		Name:        "TEST",
		Description: "TEST",
	})

	if len(scm.commands) != 1 {
		t.Error("The count fo the command in commands not correct")
	}

	if scm.commands["TEST"].Name != "TEST" {
		t.Error("AddSlashCommand failed")
	}
}

func TestDelSlashCommand(t *testing.T) {
	scm := &SlashCommandManager{
		commands: map[string]*discordgo.ApplicationCommand{
			"TEST": {
				Name:        "TEST",
				Description: "TEST",
			},
		},
		registeredCommands: map[string]*discordgo.ApplicationCommand{},
		cmdHandlers:        map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}

	scm.DelSlashCommand("TEST")

	if len(scm.commands) != 0 {
		t.Error("DelSlashCommand failed")
	}
}

func TestAddSlashCommandHandler(t *testing.T) {
	scm := &SlashCommandManager{
		commands:           map[string]*discordgo.ApplicationCommand{},
		registeredCommands: map[string]*discordgo.ApplicationCommand{},
		cmdHandlers:        map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}

	scm.AddSlashCommandHandler("TEST", func(s *discordgo.Session, i *discordgo.InteractionCreate) {})

	if len(scm.cmdHandlers) != 1 {
		t.Error("The count fo the command handler in cmdHandlers not correct")
	}

	if scm.cmdHandlers["TEST"] == nil {
		t.Error("AddSlashCommandHandler failed")
	}
}

func TestDelSlashCommandHandler(t *testing.T) {
	scm := &SlashCommandManager{
		commands:           map[string]*discordgo.ApplicationCommand{},
		registeredCommands: map[string]*discordgo.ApplicationCommand{},
		cmdHandlers: map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
			"TEST": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			},
		},
	}

	scm.DelSlashCommandHandler("TEST")

	if len(scm.cmdHandlers) != 0 {
		t.Error("DelSlashCommandHandler failed")
	}
}

func TestGetSlashCommandHandlers(t *testing.T) {
	scm := &SlashCommandManager{
		commands:           map[string]*discordgo.ApplicationCommand{},
		registeredCommands: map[string]*discordgo.ApplicationCommand{},
		cmdHandlers: map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
			"TEST1": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			},
			"TEST2": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			},
		},
	}

	handlers := scm.GetSlashCommandHandlers()

	if len(handlers) != 2 {
		t.Error("GetSlashCommandHandlers failed")
	}
}
