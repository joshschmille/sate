package main

import "log"

type character struct {
	name, moxie, smarts, wiggles, friends, pockets, gumption string
}

func (c *character) render() {
	statBlock.Text = "Name: " + c.name + "\nMoxie: " + c.moxie + "\nSmarts: " + c.smarts + "\nWiggles: " + c.wiggles + "\nFriends: " + c.friends + "\nPockets: " + c.pockets + "\nGumption: " + c.gumption
}

func (c *character) save() {

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

	//
}
