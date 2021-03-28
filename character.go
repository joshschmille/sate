package main

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
)

// A character contains data for the player's character.
type character struct {
	name, moxie, smarts, wiggles, friends, pockets, gumption, skill, style string
}

// render renders the character's stats to the stat block and triggers
// a re-render.
func (c *character) render() {
	statBlock.Text = "Name: " + c.name + "\nSkill: " + c.skill + "\nStyle: " + c.style + "\nMoxie: " + c.moxie + "\nSmarts: " + c.smarts + "\nWiggles: " + c.wiggles + "\nFriends: " + c.friends + "\nPockets: " + c.pockets + "\nGumption: " + c.gumption
	ui.Render(statBlock)
}

// save saves the character data to the character log file.
func (c *character) save() {
	f, err := os.Create("./logs/character")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(c.name + "\n" + c.skill + "\n" + c.style + "\n" + c.moxie + "\n" + c.smarts + "\n" + c.wiggles + "\n" + c.friends + "\n" + c.pockets + "\n" + c.gumption)

	if err2 != nil {
		log.Fatal(err2)
	}
}

// load loads the character from the character log file.
func (c *character) load(f string) {
	stats, err := readNameFile(f)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	if len(stats) < 9 {
		start := len(stats) - 1

		for i := start; i < 9; i++ {
			stats = append(stats, "")
		}
	}

	c.name = stats[0]
	c.skill = stats[1]
	c.style = stats[2]
	c.moxie = stats[3]
	c.smarts = stats[4]
	c.wiggles = stats[5]
	c.friends = stats[6]
	c.pockets = stats[7]
	c.gumption = stats[8]
}

// setAttribute is used to modify a given character stat field.
// When updated, it saves and renders as well.
func (c *character) setAttribute(field string, data string) {
	switch field {
	case "name":
		c.name = data
	case "skill":
		c.skill = data
	case "style":
		c.style = data
	case "moxie":
		c.moxie = data
	case "smarts":
		c.smarts = data
	case "wiggles":
		c.wiggles = data
	case "friends":
		c.friends = data
	case "pockets":
		c.pockets = data
	case "gumption":
		c.gumption = data
	}

	c.save()
	//c.load("./logs/character")
	c.render()
}

func generateSkill() string {
	prefixes, err := readNameFile("./data/pc/skillprefix.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	skills, err := readNameFile("./data/pc/skill.names")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	return prefixes[generateNumber(0, len(prefixes)-1)] + " " + skills[generateNumber(0, len(skills)-1)]
}
