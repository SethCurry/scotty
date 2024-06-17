package main

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/SethCurry/scotty/internal/finals"
	"github.com/SethCurry/scotty/internal/scotty"
	"github.com/SethCurry/scotty/pkg/eleven"
	"github.com/alecthomas/kong"
	"github.com/bwmarrin/discordgo"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

type MigrateDBCommand struct{}

func (m MigrateDBCommand) Run(ctx *Context) error {
	return ctx.DB.Schema.Create(ctx.Context)
}

type DBCommands struct {
	Migrate MigrateDBCommand `cmd:"migrate" help:"Migrate the database"`
}

type RegisterCommandsCommand struct{}

func (r RegisterCommandsCommand) Run(ctx *Context) error {
	cmds := []*discordgo.ApplicationCommand{
		{
			Name:        "scotty",
			Description: "Generate an audio message in Scotty's voice.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "text",
					Description: "The text to be spoken by Scotty.",
					Required:    true,
				},
			},
		},
		{
			Name:        "leaderboard",
			Description: "Check a player's leaderboard position.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "player",
					Description: "The player to check the leaderboard for.",
					Required:    true,
				},
			},
		},
	}

	bot, err := scotty.NewBot(ctx.Config.Discord.Token, ctx.DB, ctx.Logger)
	if err != nil {
		return err
	}

	for _, v := range cmds {
		_, err = bot.Session.ApplicationCommandCreate(ctx.Config.Discord.AppID, ctx.Config.Discord.GuildID, v)
		if err != nil {
			return fmt.Errorf("failed to register command: %w", err)
		}
	}

	return nil
}

type StartCommand struct{}

func (s StartCommand) Run(ctx *Context) error {
	bot, err := scotty.NewBot(ctx.Config.Discord.Token, ctx.DB, ctx.Logger)
	if err != nil {
		return err
	}

	elClient := eleven.NewClient(ctx.Config.TTS.APIKey)

	bot.RegisterCommand("scotty", func(sess *discordgo.Session, db *ent.Client, inter *discordgo.InteractionCreate, logger *zap.Logger) {
		buf := bytes.NewBuffer([]byte{})

		now := time.Now().Unix()

		err := elClient.TTS(inter.ApplicationCommandData().Options[0].StringValue(), ctx.Config.TTS.ScottyVoiceID, buf, eleven.VoiceSettings{
			Stability:       0.8,
			SimilarityBoost: 0.6,
			UseSpeakerBoost: true,
		})
		if err != nil {
			logger.Error("failed to create sample", zap.Error(err))
		}

		err = sess.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "It's not safe to go alone, take this",
				Files: []*discordgo.File{
					{
						Name:   fmt.Sprintf("scotty-%d.mp3", now),
						Reader: buf,
					},
				},
			},
		})
		if err != nil {
			logger.Error("failed to respond", zap.Error(err))
		}
	})

	bot.RegisterCommand("leaderboard", func(sess *discordgo.Session, db *ent.Client, inter *discordgo.InteractionCreate, logger *zap.Logger) {
		username := inter.ApplicationCommandData().Options[0].StringValue()
		player, err := finals.CheckLeaderboard(username)
		if err != nil {
			respErr := sess.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: err.Error(),
				},
			})
			if respErr != nil {
				logger.Error("failed to respond", zap.Error(respErr))
				return
			}
		}

		err = sess.InteractionRespond(inter.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("%s: #%d, %s", player.Name, player.Rank, player.League),
			},
		})
		if err != nil {
			logger.Error("failed to respond", zap.Error(err))
		}
	})

	<-ctx.Context.Done()

	return nil
}

type CLI struct {
	DB               DBCommands              `cmd:"db" help:"Interact with Scotty's database"`
	Start            StartCommand            `cmd:"start" help:"Start the bot"`
	RegisterCommands RegisterCommandsCommand `cmd:"register-commands" help:"Register Discord commands for the bot"`
}

type Context struct {
	Logger  *zap.Logger
	Config  *scotty.Config
	DB      *ent.Client
	Context context.Context
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	cfg, err := scotty.LoadDefaultConfig()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	dbConn, err := ent.Open(cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		logger.Fatal("failed to open database", zap.Error(err))
	}

	cli := &CLI{}

	ctx := kong.Parse(cli)

	err = ctx.Run(&Context{Logger: logger, Config: cfg, DB: dbConn, Context: context.Background()})
	if err != nil {
		logger.Fatal("failed to run", zap.Error(err))
	}
}
