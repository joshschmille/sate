package main

import (
	"encoding/base64"
	"image"
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
)

type macguffin struct {
	finish string
	imgId  int
}

func (m *macguffin) generate() macguffin {
	m.finish = generateMacguffinFinish()
	m.imgId = generateMacguffinImageId()

	return *m
}

func (m *macguffin) render(req string) {
	image, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(mgImages[generateNumber(0, len(mgImages)-1)])))
	if err != nil {
		log.Fatalf("failed to decode gopher image: %v", err)
	}

	macguffinBlock.Title = "Wouldn't it " + m.finish + " - (Ctrl-X: show/hide)"
	macguffinBlock.Image = image
	mgToggle = true
	ui.Render(macguffinBlock)
}

func generateMacguffinFinish() string {
	rnd := generateNumber(1, 3)

	switch rnd {
	case 1:
		return "stink if...?"
	case 2:
		return "be interesting if...?"
	case 3:
		return "be great if...?"
	}

	return "Oops."
}

func generateMacguffinImageId() int {
	return generateNumber(0, len(mgImages)-1)
}

func setMacguffinBlockImg(f string) {

}
