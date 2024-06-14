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

	return &Bot{
		db:      db,
		session: sess,
		logger:  logger,
	}, nil
}

type Bot struct {
	db       *ent.Client
	session  *discordgo.Session
	logger   *zap.Logger
	commands map[string]Command
}
