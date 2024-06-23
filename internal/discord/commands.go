package discord

import (
	"bytes"
	"fmt"
	"time"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/SethCurry/scotty/internal/finals"
	"github.com/SethCurry/scotty/pkg/eleven"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// LeaderboardCommand implements the /leaderboard Discord slash command.
// It allows players to look up another player's leaderboard position and rank
// using their Embark ID.
func LeaderboardCommand(sess *discordgo.Session, db *ent.Client, inter *discordgo.InteractionCreate, logger *zap.Logger) (*discordgo.InteractionResponse, error) {
	username := inter.ApplicationCommandData().Options[0].StringValue()
	player, err := finals.Leaderboard(username)
	if err != nil {
		return nil, err
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s: #%d, %s", player.Name, player.Rank, player.League),
		},
	}, nil
}

// ScottyCommand implements the /scotty Discord slash command. It allows players to
// create soundboard clips using the Eleven Labs API.
func ScottyCommand(elClient *eleven.Client, scottyVoiceID string) func(sess *discordgo.Session, db *ent.Client, inter *discordgo.InteractionCreate, logger *zap.Logger) (*discordgo.InteractionResponse, error) {
	return func(sess *discordgo.Session, db *ent.Client, inter *discordgo.InteractionCreate, logger *zap.Logger) (*discordgo.InteractionResponse, error) {
		buf := bytes.NewBuffer([]byte{})

		now := time.Now().Unix()

		err := elClient.TTS(inter.ApplicationCommandData().Options[0].StringValue(), scottyVoiceID, buf, eleven.VoiceSettings{
			Stability:       0.8,
			SimilarityBoost: 0.6,
			UseSpeakerBoost: true,
		})
		if err != nil {
			logger.Error("failed to create sample", zap.Error(err))
			return nil, err
		}

		return &discordgo.InteractionResponse{
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
		}, nil
	}
}
