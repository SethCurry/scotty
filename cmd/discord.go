package cmd

import (
	"github.com/SethCurry/scotty/internal/scotty"
	"github.com/SethCurry/scotty/pkg/eleven"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

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
			ctx.Logger.Error("failed to register command", zap.String("name", v.Name), zap.Error(err))
		} else {
			ctx.Logger.Info("registered command", zap.String("name", v.Name))
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

	// wait until the context is done, it typically blocks forever
	<-ctx.Context.Done()

	return nil
}
