package main

import (
	"bufio"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Create the widgets for the UI layout.
var gameLog = widgets.NewList()
var missionBlock = widgets.NewParagraph()
var statBlock = widgets.NewParagraph()
var macguffinBlock = widgets.NewImage(nil)

var player = character{}

// Set the macguffin block to be hidden at first run.
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
		"orange":  ui.Color(202),
		"purple":  ui.Color(99),
		"pink":    ui.Color(219),
		"8":       ui.Color(8),
		"red":     ui.ColorRed,
		"blue":    ui.ColorBlue,
		"black":   ui.ColorBlack,
		"cyan":    ui.ColorCyan,
		"yellow":  ui.ColorYellow,
		"white":   ui.ColorWhite,
		"clear":   ui.ColorClear,
		"green":   ui.ColorGreen,
		"magenta": ui.ColorMagenta,
	}

	primaryColor := ui.Color(32)
	secondaryColor := ui.Color(87)

	filteredWords := []string{
		"<Space>",
		"<Enter>",
		"<Backspace>",
		"<Tab>",
		"<Up>",
		"<Down>",
		"<Left>",
		"<Right>",
		"<C-x>",
		"<C-<Backspace>>",
	}

	gameLog.Title = "Game Log"
	gameLog.Rows = append(gameLog.Rows, "Welcome to [Space Aces](fg:purple): Terminal Edition!")
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()
	gameLog.SetRect(0, 0, termWidth-40, termHeight-3)
	gameLog.BorderStyle.Fg = primaryColor
	gameLog.TitleStyle.Fg = secondaryColor

	statBlock.Title = "Stats"
	player.load("./logs/character")
	player.render()
	statBlock.SetRect(termWidth-40, 0, termWidth, 9)
	statBlock.BorderStyle.Fg = primaryColor
	statBlock.TitleStyle.Fg = secondaryColor

	missionBlock.Title = "Mission"
	missionBlock.Text = ""
	missionBlock.SetRect(termWidth-40, 9, termWidth, 19)
	missionBlock.BorderStyle.Fg = primaryColor
	missionBlock.TitleStyle.Fg = secondaryColor

	startX, startY, endX, endY := calculateMacguffinRect()

	macguffinBlock.SetRect(startX, startY, endX, endY)

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

			startX, startY, endX, endY := calculateMacguffinRect()

			macguffinBlock.SetRect(startX, startY, endX, endY) //payload.Width-40, ((payload.Width-40)/2)+1)
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
		case "<C-<Backspace>>":
			length := len(inputBox.Text)
			if length > 0 {
				inputBox.Text = inputBox.Text[:length-1]
				ui.Render(inputBox)
			}
		case "<Up>":
			gameLog.ScrollPageUp()
			ui.Render(gameLog)
		case "<Down>":
			gameLog.ScrollPageDown()
			ui.Render(gameLog)
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

// readNameFile returns a string slice containing the lines read from the provided file path.
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

// splitStringByMax returns a string slice containing the data from the provided string
// split into chunks based on the given maximum size.
func splitStringByMax(s string, chunkSize int) []string {
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

// combineArgsToString returns a string containing each string within the provided string
// slice, separated by " ".
func combineArgsToString(s []string) string {
	output := ""
	for i := 0; i < len(s); i++ {
		if i > 0 {
			output += " "
		}
		output += s[i]
	}
	return output
}

// renderOutput accepts a string of any length and uses splitStringByMax() to split it
// based on the maximum length, which is determined by the width of the terminal window.
// It then outputs those strings to the game log window, scrolls it to the bottom, and
// triggers it to re-render to show the changes.
// It also sends the content of the original string to the writeLog() func.
func renderOutput(s string) {
	chunked := splitStringByMax(s, termWidth-42)
	for i := 0; i < len(chunked); i++ {
		gameLog.Rows = append(gameLog.Rows, chunked[i])
	}
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()

	writeLog(s)

	mgToggle = false
	ui.Render(gameLog)
}

// writeLog accepts a string of any length, and appends the string to the log file.
func writeLog(s string) {
	f, err := os.OpenFile("logs/all", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := s
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// parseArgs accepts a string, splits it using " " as the delimiter, and determines if it
// is a valid command or not, sending everything after the first word as a list of args.
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
		case "log":
			cmdLog(args)
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
		case "massivemonster":
			cmdMassiveMonster(args)
		case "mm":
			cmdMassiveMonster(args)
		case "beasty":
			cmdBeasty(args)
		case "macguffin":
			cmdMacguffin(args)
		case "mg":
			cmdMacguffin(args)
		case "backstory":
			cmdBackstory(args)
		case "character":
			cmdCharacter(args)
		case "help":
			cmdHelp(args)
		default:
			renderOutput("[Invalid Command.](fg:red)")
		}
	}
}
