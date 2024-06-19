package cmd

import (
	"fmt"

	"github.com/SethCurry/scotty/internal/scotty"
	"github.com/SethCurry/scotty/pkg/eleven"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Discord struct {
	RegisterCommands RegisterCommands `cmd:"register-commands" help:"Register slash commands with Discord"`
	Start            StartBot         `cmd:"start" help:"Start the scotty bot"`
}

type RegisterCommands struct{}

func (r RegisterCommands) Run(ctx *Context) error {
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
		return fmt.Errorf("failed to create discord bot: %w", err)
	}

	for _, v := range cmds {
		_, err = bot.Session.ApplicationCommandCreate(ctx.Config.Discord.AppID, ctx.Config.Discord.GuildID, v)
		if err != nil {
			ctx.Logger.Error("failed to register command", zap.String("name", v.Name), zap.Error(err))
		} else {
			ctx.Logger.Info("registered command", zap.String("name", v.Name))
		}
	}

	return nil
}

// StartBot implements a command to start the Discord bot.
type StartBot struct{}

func (s StartBot) Run(ctx *Context) error {
	bot, err := scotty.NewBot(ctx.Config.Discord.Token, ctx.DB, ctx.Logger)
	if err != nil {
		return fmt.Errorf("failed to create a bot: %w", err)
	}

	ctx.Logger.Info("started Discord bot connection")

	elClient := eleven.NewClient(ctx.Config.TTS.APIKey)

	bot.RegisterCommand("scotty", scotty.ScottyCommand(elClient, ctx.Config.TTS.ScottyVoiceID))
	bot.RegisterCommand("leaderboard", scotty.LeaderboardCommand)

	ctx.Logger.Info("waiting for context to be cancelled")

	// wait until the context is done, it typically blocks forever
	<-ctx.Context.Done()

	return nil
}
