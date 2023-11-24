package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var UseLocalServer bool
var WipeOldCommands bool
var BotRequestToken string
var WaitForClose chan bool

func runAndHandle(token string) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	dg.ShouldReconnectOnError = true
	dg.StateEnabled = true
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	registerCommands(dg)
	go setRandomStatus(dg)

	fmt.Println("The Bot is now running. Press CTRL-C to exit.")

	<-WaitForClose

	fmt.Println("Restarting...")

	dg.Close()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	flag.BoolVar(&UseLocalServer, "local", false, "if the bot should run in localhost mode or not")
	flag.BoolVar(&WipeOldCommands, "reset", false, "if the bot should wipe existing commands, and make new ones")
	flag.Parse()

	WaitForClose = make(chan bool, 1)

	if UseLocalServer {
		fmt.Println("THE BOT IS RUNNING IN LOCAL SERVER MODE!")
	}

	requestToken, err := os.LookupEnv("BotRequestToken")
	if !err {
		fmt.Println("Error looking up bot verification token")
		return
	}
	BotRequestToken = requestToken

	// Create a new Discord session using the provided bot token.
	if token, found := os.LookupEnv("BotToken"); found {
		runAndHandle(token)
	} else {
		fmt.Println("Error looking up bot token")
	}
}
