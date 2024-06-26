// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/SethCurry/scotty/internal/ent/guild"
	"github.com/SethCurry/scotty/internal/ent/schema"
	"github.com/SethCurry/scotty/internal/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	guildFields := schema.Guild{}.Fields()
	_ = guildFields
	// guildDescWelcomeTemplate is the schema descriptor for welcome_template field.
	guildDescWelcomeTemplate := guildFields[2].Descriptor()
	// guild.DefaultWelcomeTemplate holds the default value on creation for the welcome_template field.
	guild.DefaultWelcomeTemplate = guildDescWelcomeTemplate.Default.(string)
	// guildDescWelcomeChannel is the schema descriptor for welcome_channel field.
	guildDescWelcomeChannel := guildFields[3].Descriptor()
	// guild.DefaultWelcomeChannel holds the default value on creation for the welcome_channel field.
	guild.DefaultWelcomeChannel = guildDescWelcomeChannel.Default.(string)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescRankedScore is the schema descriptor for ranked_score field.
	userDescRankedScore := userFields[1].Descriptor()
	// user.DefaultRankedScore holds the default value on creation for the ranked_score field.
	user.DefaultRankedScore = userDescRankedScore.Default.(int)
}
