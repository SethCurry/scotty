package scotty

import (
	"fmt"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Command func(*discordgo.Session, *ent.Client, *discordgo.InteractionCreate, *zap.Logger)

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
			h(s, db, i, logger)
		}
	})

	return b, nil
}

type Bot struct {
	db       *ent.Client
	Session  *discordgo.Session
	logger   *zap.Logger
	commands map[string]Command
}

func (b *Bot) RegisterCommand(name string, command Command) {
	if b.commands == nil {
		b.commands = make(map[string]Command)
	}

	b.commands[name] = command
}
