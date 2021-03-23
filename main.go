package main

import (
	"bufio"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var gameLog = widgets.NewList()
var missionBlock = widgets.NewParagraph()
var macguffinBlock = widgets.NewImage(nil)
var mgToggle = false

var termWidth = 0
var termHeight = 0

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	termWidth, termHeight = ui.TerminalDimensions()

	ui.StyleParserColorMap = map[string]ui.Color{
		"red":    ui.Color(196),
		"blue":   ui.Color(32),
		"black":  ui.Color(16),
		"cyan":   ui.Color(87),
		"yellow": ui.Color(226),
		"orange": ui.Color(202),
		"white":  ui.ColorWhite,
		"clear":  ui.ColorClear,
		"green":  ui.Color(46),
		"purple": ui.Color(99),
		"pink":   ui.Color(219),
		"8":      ui.Color(8),
	}

	primaryColor := ui.Color(32)
	secondaryColor := ui.Color(87)

	filteredWords := [9]string{
		"<Space>",
		"<Enter>",
		"<Backspace>",
		"<Tab>",
		"<Up>",
		"<Down>",
		"<Left>",
		"<Right>",
		"<C-x>",
	}

	gameLog.Title = "Game Log"
	gameLog.Rows = append(gameLog.Rows, "Welcome to [Space Aces](fg:purple): Terminal Edition!")
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()
	gameLog.SetRect(0, 0, termWidth-40, termHeight-3)
	gameLog.BorderStyle.Fg = primaryColor
	gameLog.TitleStyle.Fg = secondaryColor

	statBlock := widgets.NewParagraph()
	statBlock.Title = "Stats"
	statBlock.Text = "Name: Riya\nMoxie: +1\nSmarts: 0\nWiggles: 0\nFriends: 0\nPockets: 0\nGumption: 10/10"
	statBlock.SetRect(termWidth-40, 0, termWidth, 9)
	statBlock.BorderStyle.Fg = primaryColor
	statBlock.TitleStyle.Fg = secondaryColor

	missionBlock.Title = "Mission"
	missionBlock.Text = ""
	missionBlock.SetRect(termWidth-40, 9, termWidth, 19)
	missionBlock.BorderStyle.Fg = primaryColor
	missionBlock.TitleStyle.Fg = secondaryColor

	maxX, maxY := calculateMacguffinSize()

	macguffinBlock.SetRect(1, 1, maxX, maxY)

	inputBox := widgets.NewParagraph()
	inputBox.Text = ""
	inputBox.SetRect(0, termHeight-3, termWidth, termHeight)
	inputBox.BorderStyle.Fg = primaryColor
	inputBox.TitleStyle.Fg = secondaryColor

	ui.Render(gameLog, statBlock, missionBlock, inputBox)
	if mgToggle {
		ui.Render(macguffinBlock)
	}

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "<C-c>":
			return
		case "<C-x>":
			if mgToggle {
				mgToggle = false
				ui.Render(gameLog, statBlock, missionBlock, inputBox)
			} else {
				mgToggle = true
				ui.Render(macguffinBlock)
			}
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			termWidth = payload.Width
			termHeight = payload.Height
			gameLog.SetRect(0, 0, payload.Width-40, payload.Height-3)
			statBlock.SetRect(payload.Width-40, 0, payload.Width, 9)
			missionBlock.SetRect(payload.Width-40, 9, payload.Width, 19)

			maxX, maxY = calculateMacguffinSize()
			macguffinBlock.SetRect(1, 1, maxX, maxY) //payload.Width-40, ((payload.Width-40)/2)+1)
			inputBox.SetRect(0, payload.Height-3, payload.Width, payload.Height)
			ui.Clear()
			ui.Render(gameLog, statBlock, missionBlock, inputBox)
			if mgToggle {
				ui.Render(macguffinBlock)
			}
		case "<Enter>":

			parseArgs(inputBox.Text)

			inputBox.Text = ""

			ui.Render(inputBox)

		case "<Space>":
			inputBox.Text += " "
		case "<Backspace>":
			length := len(inputBox.Text)
			if length > 0 {
				inputBox.Text = inputBox.Text[:length-1]
				ui.Render(inputBox)
			}
		case "<Up>":
			// TODO: Command History.
		case "<Down>":
			// TODO: Command History.

		}

		switch e.Type {
		case ui.KeyboardEvent:
			filtered := false
			for i := 0; i < len(filteredWords); i++ {
				if e.ID == filteredWords[i] {
					filtered = true
				}
			}

			if !filtered {
				inputBox.Text += e.ID
				ui.Render(inputBox)
			}
		}
	}
}

func readNameFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Chunks(s string, chunkSize int) []string {
	if chunkSize >= len(s) {
		return []string{s}
	}
	var chunks []string
	chunk := make([]rune, chunkSize)
	len := 0
	for _, r := range s {
		chunk[len] = r
		len++
		if len == chunkSize {
			chunks = append(chunks, string(chunk))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(chunk[:len]))
	}
	return chunks
}

func calculateMacguffinSize() (int, int) {
	mgMaxWidth := (termHeight * 2) - 14

	if termWidth-40 > mgMaxWidth {
		return mgMaxWidth, mgMaxWidth / 2
	}

	return termWidth - 41, ((termWidth - 41) / 2)
}

func renderOutput(s string) {
	chunked := Chunks(s, termWidth-42)
	for i := 0; i < len(chunked); i++ {
		gameLog.Rows = append(gameLog.Rows, chunked[i])
	}
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()

	mgToggle = false
	ui.Render(gameLog)
}

func parseArgs(s string) {
	if len(s) > 0 {
		all := strings.Fields(s)
		cmd := all[0]
		args := make([]string, len(all))

		copy(args, all[1:])

		renderOutput("[>](fg:cyan) [" + cmd + " " + strings.Join(args, " ") + "](fg:8)")

		switch cmd {
		case "roll":
			cmdRoll(args)
		case "name":
			cmdName(args)
		case "likely":
			cmdLikely(args)
		case "possibly":
			cmdPossibly(args)
		case "unlikely":
			cmdUnlikely(args)
		case "mission":
			cmdMission(args)
		case "event":
			cmdEvent(args)
		case "ruin":
			cmdRuin(args)
		case "monster":
			cmdMonster(args)
		case "treasure":
			cmdTreasure(args)
		case "hazard":
			cmdHazard(args)
		case "gizmo":
			cmdGizmo(args)
		case "ship":
			cmdShip(args)
		case "explore":
			cmdExplore(args)
		case "planet":
			cmdPlanet(args)
		case "navigate":
			cmdNavigate(args)
		case "sector":
			cmdSector(args)
		case "npc":
			cmdNpc(args)
		case "mech":
			cmdMech(args)
		case "mm":
			cmdMassiveMonster(args)
		case "beasty":
			cmdBeasty(args)
		case "macguffin":
			cmdMacguffin(args)
		default:
			renderOutput("[Invalid Command.](fg:red)")
		}
	}
}
