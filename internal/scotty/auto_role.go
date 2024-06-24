package scotty

import (
	"context"
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
	}

	for _, role := range autoRoles {
		err = sess.GuildMemberRoleAdd(event.GuildID, event.User.ID, role.RoleID)
		if err != nil {
			logger.Error("failed to add role", zap.String("role_id", role.RoleID), zap.Error(err))
		}
	}
}
