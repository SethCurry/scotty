package main

import (
	"context"
	"fmt"

	"github.com/SethCurry/scotty/internal/ent"
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

	bot.RegisterCommand("scotty", scotty.ScottyCommand(elClient, ctx.Config.TTS.ScottyVoiceID))
	bot.RegisterCommand("leaderboard", scotty.LeaderboardCommand)

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
