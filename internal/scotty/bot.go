package scotty

import (
	"fmt"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// Command defines the function signature that Discord slash commands should
// implement.
type Command func(*discordgo.Session, *ent.Client, *discordgo.InteractionCreate, *zap.Logger) (*discordgo.InteractionResponse, error)

// NewBot creates a new instance of a *Bot and starts the underlying Discord session.
func NewBot(token string, db *ent.Client, logger *zap.Logger) (*Bot, error) {
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord bot: %w", err)
	}

	sess.Identify.Intents = discordgo.IntentsAll

	err = sess.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord session: %w", err)
	}

	b := &Bot{
		db:       db,
		Session:  sess,
		logger:   logger,
		commands: make(map[string]Command),
	}

	b.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := b.commands[i.ApplicationCommandData().Name]; ok {
			var respErr error

			resp, err := h(s, db, i, logger)
			if err != nil {
				respErr = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Error: " + err.Error(),
					},
				})
			} else {
				respErr = s.InteractionRespond(i.Interaction, resp)
			}

			if respErr != nil {
				logger.Error("failed to send interaction response", zap.Error(err))
			}
		}
	})

	return b, nil
}

// Bot is a struct that implements interactions with Discord such as
// slash commands.
type Bot struct {
	db       *ent.Client
	Session  *discordgo.Session
	logger   *zap.Logger
	commands map[string]Command
}

// RegisterCommand registers a new slash command handler with the bot.
// This does not register it with Discord, it only ensures the bot is
// ready to handle that slash command.
func (b *Bot) RegisterCommand(name string, command Command) {
	if b.commands == nil {
		b.commands = make(map[string]Command)
	}

	b.commands[name] = command
}
