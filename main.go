package main

import (
	"cron-bot/commands"
	"cron-bot/config"
	"database/sql"
	"fmt"
	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := config.ReadConfig("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	goBot, err := discordgo.New("Bot " + conf.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	goBot.AddHandler(commands.GetMessageHandler(conf.BotPrefix, db))
	// goBot.ChannelMessageSend()

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	goBot.Close()
}
