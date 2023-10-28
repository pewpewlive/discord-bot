package main

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
)

func setRandomStatus(session *discordgo.Session) {
	for {
		statuses := [...]string{
			"Eskiv",
			"Hexagon",
			"Asteroids",
			"Fury",
			"Waves",
			// Unreleased
			"Asteroweird",
			// PP2
			"Pandemonium",
			"Symbiosis",
			"Pacifism",
			"Highway",
			"Amalgam",
			"Dodge this",
			"Assault",
			"in the Sandbox mode",
			// Custom levels
			"CRASBAF",
			"Purple World",
			"zale",
			"Simon Says",
			"Cube Prison",
			"UFOMania",
			"Deadline",
			"Pew-Man",
			"Waves Blue",
			"Pacifism Live",
			"Duel",
			"Rolling Cubes vs Spinning Hexagons",
			"Madreinka",
			"NumFall",
			"Khorne's Box",
			"Sphell",
			"Better Tutorial",
			"Intro",
			"Inertiac World",
			"Timebomb",
			"Bombardment",
			"Warmup",
			"Cage theory",
			"Interdimensional problems",
			"Triskiv",
			"Deadzone",
			"Claustrophobia",
			"Bouncy Recoil",
			"Fury Pro",
			"Heavy Plummet",
			"Restanvi",
			"Bee Hive",
			"Linkage",
			"Waves Pro",
			"Pew pong",
			"Just pong",
			"Field of wormholes",
			"Inertiacs RAGE",
			"Matching cubes v2",
			"Cubes and Squares?",
			"Waves Red",
			"Stripped Hell",
			// Other
			"PewPew 3",
			"PewPew 4",
			"PewPewPew",
			"the sekret level",
			"the other sekret level",
			"random levels",
			"ping...pong",
			"the only game that matters",
			"Geometry Wars",
			"Battle Girl",
			"Robotron",
			"Maelstrom",
			"the intro, as 4 players",
			"the warm up",
			"the tutorial",
			"Waves Blue",
			"with the white ship",
			"with the wall API",
		}
		newStatus := statuses[rand.Intn(len(statuses))]
		session.UpdateGameStatus(0, newStatus)
		time.Sleep(30 * time.Minute)
	}
}
