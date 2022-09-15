package bbsui

func Intro() string {
	
	ec := "" + 
	"    Somewhere between Andromeda and the Milky Way lies a cluster of stars\n" +
	"containing humanoid and other life forms. After centuries of rule, the\n" +
	"race of Esterians has mysteriously vanished leaving this small galaxy in a\n" +
	"state of anarchy. Each of the once supressed civilizations, now armed with the\n" +
	"modern technology of the Esterians and old fashioned greed, seeks to control\n" +
	"the cluster of stars and rule all other civilizations.\n" + 
	"    Esterian Conquest is a game for 4 to 25 players. Each player is considered\n" +
	"to be a ruler of a young stellar empire.  The object of the game is to rule a\n" +
	"small galaxy by colonizing and/or conquering other worlds. Initially, each\n" +
	"player has one world and a few small fleets of starships. As ruler of your\n" +
	"empire, you must decide how heavily to tax your population, order your planets\n" +
	"to build starships, and send your fleets onto missions.\n" +
	"    Each game turn represents one year of game time and is the time from the\n" +
	"end of one maintenance to the end of the next. Maintenance is typically run\n" +
	"from one to seven times weekly, depending on how the sysop has the game set.\n" +
	"Between each maintenance, players may give orders to their fleets and planets,\n" +
	"send messages to each other, etc. Players may change standing orders to any\n" +
	"fleet or planet at any time. However, if maintenance has run since the orders\n" +
	"were given, he may find his fleets in new locations as they travel towards\n" +
	"their destinations during maintenance.\n\n" +
	"Copyright (C) 1990-1992 by Bentley C. Griffith.  All rights reserved."
	
	c15 := "\n\n" + 
	"Ver 1.5 CREDITS:\n" +
	"------------------------------------------------------------------------------\n" +
	"Program design, concept & execution:   Bentley Griffith\n" +
    "                                       Joel Cohen\n" +
    "                                       Blake Miller\n" +
    "                                       Todd Simmons\n" +
    "                                       William Smith\n\n" +
    "Software documentation.............:   Joel Cohen\n" +
    "Additional Support:  Ken Browning      William Padilla\n" +
    "                     Tom Egan          Bert Pittman\n" +
    "                     Russel Femyer     Rocky Rawlins\n" +
    "                     Scott Fort        Kelly Rosato\n" +
    "                     Ricky Morgan      Charles Struckel\n" +
    "                     Pete Mohney\n\n" +
    "Special thanks to our brave Alpha\n" + 
    "Testers on the following boards....:   F/X BBS\n" +
    "                                       Joker's Castle BBS\n" +
    "                                       ST BBS\n" +
    "                                       The MATRIX BBS\n" +
	"------------------------------------------------------------------------------" 

	c20 := "\n" +
	"EC2 CREDITS:\n" +
	"------------------------------------------------------------------------------\n" +
	"Game programming...................:   Mason Austin Green\n" +
    "Software documentation.............:   TBD\n" +
    "Additional Support.................:   TBD\n" +
    "Special thanks to our brave Alpha\n" + 
    "Testers on the following boards....:   Fool's Quarter BBS\n" +
    "                                       Constructive Chaos BBS\n\n" +
	"EC2 by Mason Austin Green is licensed under CC BY-NC-SA 4.0\n" + 
	"------------------------------------------------------------------------------" 
	
	return ec+c15+c20
}
