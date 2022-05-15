package chatopsbot

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"guanyu.dev/chatopsbot/internal/chatopsbot/slashcmd"
)

type ChatopsBot struct {
	dg  *discordgo.Session
	scm *slashcmd.SlashCommandManager
}

func init() {
	// Check whether discord token exist or not
	if os.Getenv("DISCORD_TOKEN") == "" {
		fmt.Fprintln(os.Stderr, "No token provided.")
		os.Exit(1)
	}
}

func New() (*ChatopsBot, error) {
	// Initial discordgo instance with bot's token
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating Discord session, %s", err))
	}

	// Initial slash command manager instance
	scm := slashcmd.NewSlashCommandManager()

	// Register all slash commands
	slashcmd.RegisterAllSlashCommands(scm)

	// Create ChatopsBot instance
	bot := &ChatopsBot{dg, scm}

	// Register all event handlers
	RegisterHandlers(bot)

	return bot, nil
}

func (b *ChatopsBot) Start() (bool, error) {
	// Check whether discord session has already been initialized
	if b.dg == nil {
		return false, errors.New("Initial discord session first.")
	}

	// Open a websocket session to Discord and begin listening
	err := b.dg.Open()
	if err != nil {
		return false, errors.New(fmt.Sprintf("An error occur when opening a websocket session to Discord, %s", err))
	}

	// Cleanly close down the Discord session
	defer b.dg.Close()

	// Register all slash commands to the target guild in discord
	err = b.scm.RegisterSlashCommands(b.dg)
	if err != nil {
		return false, errors.New(fmt.Sprintf("An error occur when registering the slash command to Discord, %s", err))
	}
	fmt.Println("Register slash commands succeed")

	// Wait until term signal is received
	fmt.Println("ChatOps bot for disgobot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Deregister all slash commands from the target guild in discord
	err = b.scm.DeregisterSlashCommands(b.dg)
	if err != nil {
		return false, errors.New(fmt.Sprintf("An error occur when deregistering the slash command from Discord, %s", err))
	}
	fmt.Println("Deregister slash commands succeed")

	return true, nil
}

// Return discord session
func (b *ChatopsBot) GetSession() *discordgo.Session {
	return b.dg
}
