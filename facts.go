package main

import (
	"fmt"
	"math/rand"
)

func getFact(index int) string {
	facts := [...]string{
		//"***Unused string replaced with a survive fact***",
		"Kouglof is a cake and it is angry",
		"There are 512 possible ship configuration combinations",
		"BAF is short for Back And Forth",
		"Eskiv comes from the verbe \"esquiver\" which means \"to dodge\" in French",
		"You can have a negative score in Eskiv",
		"Every bullet that you shoot that touches a mothership earns you 10 points",
		"Bullets from UFOs also destroy enemies",
		"Previously in Hexagon, patience was a virtue, but now speed is",
		"In Asteroids, there is an infinite amount of straight lines passing through two points",
		"It's possible to survive Fury with the bonus score, without shooting.",
		"In Fury you can theoretically get 9605 points (55×160=8800 plus bonuses: 8800+300+505)",
		"Every single letter in PewPew was hand made <:pepe_hands:779074545969004574>",
		"It's written \"PewPew\", *not* \"Pew Pew\"",
		"It took more than a year to figure out the secret of Khorne's box. It was only discovered when the joysticks were shown in replays.",
		"Earlier there was as Easter egg in the game called \"Russian Roulette\", which had to be removed because it was complicating the UI navigation. For more information about it <https://youtu.be/y1IHCkrPEVE>",
		"The White Alpha ship has been requested for over a thousand days!",
		"The Stream ship fangs are shaped like number 1s to symbolize PewPew Live's 1st anniversary.",
		"It takes 1 tick (0.0333.. seconds) for a quick pause to happen",
		"In Bombardment there is a 0.1% chance that the bomb color is yellow instead of pink. Shiny!",
		"The game was supposed to display peace symbols in the background on September 21st, 2023, the International Day of Peace (United Nations), but there was a mistake and it showed error meshes instead",
		"Delaunay triangulation is used for generating the background of Waves",
		"The AK-286 was either created in 3286 or 4286",
		"Mothership hitboxes are smaller than they look",
		"The idea of allowing multiple colors in names came from Quake 3",
		"The name PewPew **Live** was inspired by Quake **Live**",
		"At first, PewPew Live was going to be a browser-only game",
		"The amount of damage you take by an asteroid depends on its size",
		"The development of Angry Kouglof was live streamed",
		"In the level Hexagon you are in a hexagonal arena with a hexagon in the middle, hexagons are in the background and you are collecting hexagons to progress",
		"Eskiv was created 16 years ago by JF as a flash game",
		"Skins in PewPew 2 and 1 have different buffs",
		"The first debuff was introduced in the PewPew Live custom level \"Restanvi\", and it slowed players down",
		"Heavy Plummet is 3894 lines of code",
		"The Waries are hexagonal",
		"To prevent farming, the bullets of waries do not give out any points when killed.",
		"The game includes a language called meme-nglish, you filthy casual",
		"The developer's name is Jean-Fran**ç**ois Geyelin not Jean-Fran**c**ois Geyelin",
		"The game was first translated to Russian and Croatian",
		"You use Lua to make custom levels",
		"The PewPew Live discord bot was created on the 10th of January, 2021",
		"The official PewPew Live Wiki was created on January 26, 2021",
		"The official discord channel for PewPew was created January 15th, 2018",
		"The domain name pewpew.live was registered June 25th, 2017",
		"The first commit in the PewPew Live git repository was June 25th, 2017",
		"The domain name is registered with Google Domains",
		"The server hosting this bot is running on the Google Compute Engine",
		"The website and score system is running on the Google App Engine",
		"The game is written in C++, Lua, Obj-C++, and Java.",
		"The website, the account system, the score system, ppl-utils, the multiplayer server, and this discord bot are written in Go",
		"The Rolling Cubes have had three names: Rolling cubes, Marching Cubes, RollinCubes",
		"BAFs were called Arrows in PewPew 2",
		"The design of one of the ship from PewPew 2 was shamelessly taken from the game \"Battle Girl\"",
		"PewPew was first released on iOS and was later ported to Android. PewPew Live was first released on Android and was later ported to iOS.",
		"The white ship in PewPew Live's icon first appeared in PewPew 1's Asteroids mode.",
		"The Statistics button in the main screen is 3 vertical bars. The order of the 3 bars changes.",
		"The Information button in the main screen moved from the left column to the right column to leave enough space for the achievement button.",
	}

	// surviveFacts := [...]string{
	// 	"you can survive 75 seconds without moving in Eskiv",
	// 	"you can survive 6 seconds without moving in Asteroids",
	// 	"you can survive 40-50 seconds without moving in Hexagon",
	// 	"you can survive 5 seconds without moving in Fury",
	// 	"you can survive 7 seconds without moving in Waves",
	// 	"you can survive 7-8 seconds without moving in Linkage",
	// 	"you can survive 57 seconds without moving in Hexagon: Challenge 1",
	// 	"you can survive ~16.3 seconds without moving in Bee Hive",
	// 	"you can survive 5 seconds without moving in Restavni",
	// 	"you can survive 16 seconds without moving in Angry Kouglof",
	// 	"you can survive 35 seconds without moving in Heavy Plummet",
	// 	"you can survive 5 seconds without moving in Fury Pro",
	// 	"you can survive ~10 seconds in Waves Pro without moving",
	// }

	if index == 0 {
		randomIndex := rand.Intn(len(facts))
		// if index == 0 {
		// 	return surviveFacts[rand.Intn(len(surviveFacts))]
		// }
		return fmt.Sprintf("Fact #%d: %s", randomIndex+1, facts[randomIndex])
	} else {
		if index >= 1 && index <= len(facts) {
			return fmt.Sprintf("Fact #%d: %s", index, facts[index-1])
		} else {
			return fmt.Sprintf("The index %d is invalid, make sure it goes from 1 to %d", index, len(facts))
		}
	}
}

func getPlayerFact(index int) string {
	facts := [...]string{
		"Momos Necto was the first person to survive Fury with the bonus score, without shooting.",
		"Mutox has created the most custom enemies!",
		"The 2022 PPL Level Jam winner is tavitu!",
		"The first custom community-made level is Waves Pro by FLAVOUR",
		"*Someone* once pinged the entire server by uploading a score with the name `@ everyone`",
		"Brandon was the first person to hold the number 1 spot for all 5 modes simultaneously",
	}

	if index == 0 {
		randomIndex := rand.Intn(len(facts))
		return fmt.Sprintf("Player Fact #%d: %s", randomIndex+1, facts[randomIndex])
	} else {
		if index >= 1 && index <= len(facts) {
			return fmt.Sprintf("Player Fact #%d: %s", index, facts[index-1])
		} else {
			return fmt.Sprintf("The index %d is invalid, make sure it goes from 1 to %d", index, len(facts))
		}
	}
}
