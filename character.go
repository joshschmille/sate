package main

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
)

type character struct {
	name, moxie, smarts, wiggles, friends, pockets, gumption string
}

func (c *character) render() {
	statBlock.Text = "Name: " + c.name + "\nMoxie: " + c.moxie + "\nSmarts: " + c.smarts + "\nWiggles: " + c.wiggles + "\nFriends: " + c.friends + "\nPockets: " + c.pockets + "\nGumption: " + c.gumption
	ui.Render(statBlock)
}

func (c *character) save() {
	f, err := os.Create("./logs/character")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(c.name + "\n" + c.moxie + "\n" + c.smarts + "\n" + c.wiggles + "\n" + c.friends + "\n" + c.pockets + "\n" + c.gumption)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func (c *character) load(f string) {
	stats, err := readNameFile(f)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	c.name = stats[0]
	c.moxie = stats[1]
	c.smarts = stats[2]
	c.wiggles = stats[3]
	c.friends = stats[4]
	c.pockets = stats[5]
	c.gumption = stats[6]
}

func (c *character) setAttribute(field string, data string) {
	switch field {
	case "name":
		c.name = data
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
