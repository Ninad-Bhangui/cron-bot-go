package commands

import (
	"database/sql"

	"github.com/bwmarrin/discordgo"
)

func GetMessageHandler(BotPrefix string, db *sql.DB) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == BotPrefix+"ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		}
		if m.Content == BotPrefix+"sqlite version" {
			var version string
			err := db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

			if err != nil {
				_, _ = s.ChannelMessageSend(m.ChannelID, "Could not determine version")
			}
			_, _ = s.ChannelMessageSend(m.ChannelID, version)
		}
	}
}
