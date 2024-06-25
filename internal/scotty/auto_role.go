package scotty

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"github.com/SethCurry/scotty/internal/ent"
	"github.com/SethCurry/scotty/internal/ent/autorolerule"
	"github.com/SethCurry/scotty/internal/ent/guild"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// AutoRoleOnUserJoin is a handler that connects to the OnUserJoinGuild handler for a bot.
// When a user joins the guild, it will automatically add any roles stored in scotty's database
// for auto-roles.
func AutoRoleOnUserJoin(sess *discordgo.Session, event *discordgo.GuildMemberAdd, dbClient *ent.Client, logger *zap.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	autoRoles, err := dbClient.AutoRoleRule.Query().Where(autorolerule.HasGuildWith(guild.GuildIDEQ(event.GuildID))).All(ctx)
	if err != nil {
		logger.Error("failed to fetch auto role rules", zap.Error(err))
		return
	}

	for _, role := range autoRoles {
		err = sess.GuildMemberRoleAdd(event.GuildID, event.User.ID, role.RoleID)
		if err != nil {
			logger.Error("failed to add role", zap.String("role_id", role.RoleID), zap.Error(err))
		}
	}
}

type WelcomeTemplateContext struct {
	User *discordgo.User
}

func WelcomeOnUserJoin(sess *discordgo.Session, event *discordgo.GuildMemberAdd, dbClient *ent.Client, logger *zap.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	foundGuild, err := dbClient.Guild.Query().Where(guild.GuildIDEQ(event.GuildID)).Only(ctx)
	if err != nil {
		logger.Error("failed to get guild for welcome message", zap.Error(err))
		return
	}

	if foundGuild.WelcomeTemplate != "" && foundGuild.WelcomeChannel != "" {
		tmpl, err := template.New("welcome").Parse(foundGuild.WelcomeTemplate)
		if err != nil {
			logger.Error("failed to parse welcome message", zap.String("guild_id", event.GuildID), zap.Error(err))
			return
		}

		buf := bytes.NewBuffer([]byte{})

		err = tmpl.Execute(buf, WelcomeTemplateContext{
			User: event.User,
		})
		if err != nil {
			logger.Error("failed to execute welcome message", zap.String("guild_id", event.GuildID), zap.Error(err))
			return
		}

		_, err = sess.ChannelMessageSend(foundGuild.WelcomeChannel, buf.String())
		if err != nil {
			logger.Error("failed to send welcome message", zap.Error(err))
			return
		}
	}
}
