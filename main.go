package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"gopkg.in/natefinch/lumberjack.v2"
)

const version = "0.9.0"

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
	help := flag.Bool("help", false, "Show this help")
	ver := flag.Bool("version", false, "Show your current version of SATE.")
	update := flag.Bool("update", false, "Update SATE automatically")
	slug := flag.String("slug", "joshschmille/sate", "Repository of this command")

	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
		os.Exit(0)
	}

	if *ver {
		fmt.Println(version)
		os.Exit(0)
	}

	if *update {
		if err := selfUpdate(*slug); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Setup lumberjack to be used with Go logger.
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/all.md",
		MaxSize:    5, // megabytes
		MaxBackups: 3,
	})

	// Hide everything extra when logging (timestamps, etc.)
	log.SetFlags(0)
	log.Println("# Session Start - " + time.Now().Format("2006.01.02 15:04:05"))

	termWidth, termHeight = ui.TerminalDimensions()

	// Setup color options. This needs to be revisited with an eye on
	// cross-platform compatibility.
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

	// Define the most used colors. Can be used later to create themes potentially.
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

	// Setup the game log block.
	gameLog.Title = "Game Log"
	gameLog.Rows = append(gameLog.Rows, "Welcome to [Space Aces](fg:purple): Terminal Edition!")
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()
	gameLog.SetRect(0, 0, termWidth-40, termHeight-3)
	gameLog.BorderStyle.Fg = primaryColor
	gameLog.TitleStyle.Fg = secondaryColor

	// Setup the Character block.
	statBlock.Title = "Character"
	player.load("./logs/character")
	player.render()
	statBlock.SetRect(termWidth-40, 0, termWidth, 11)
	statBlock.BorderStyle.Fg = primaryColor
	statBlock.TitleStyle.Fg = secondaryColor

	// Setup the Mission block.
	missionBlock.Title = "Mission"
	missionBlock.Text = ""
	missionBlock.SetRect(termWidth-40, 11, termWidth, 21)
	missionBlock.BorderStyle.Fg = primaryColor
	missionBlock.TitleStyle.Fg = secondaryColor

	// Setup the Macguffin block.
	startX, startY, endX, endY := calculateMacguffinRect()
	macguffinBlock.SetRect(startX, startY, endX, endY)

	// Setup the Input block.
	inputBox := widgets.NewParagraph()
	inputBox.Text = ""
	inputBox.SetRect(0, termHeight-3, termWidth, termHeight)
	inputBox.BorderStyle.Fg = primaryColor
	inputBox.TitleStyle.Fg = secondaryColor

	// TODO: Possibly add a scratchpad block that can be used like a notepad.

	// Render the blocks to the terminal.
	ui.Render(gameLog, statBlock, missionBlock, inputBox)
	if mgToggle {
		ui.Render(macguffinBlock)
	}

	// Setup event listeners for various types of input.
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
			statBlock.SetRect(payload.Width-40, 0, payload.Width, 11)
			missionBlock.SetRect(payload.Width-40, 11, payload.Width, 21)

			startX, startY, endX, endY := calculateMacguffinRect()
			macguffinBlock.SetRect(startX, startY, endX, endY)

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
func renderOutput(s string, format string, color string) {
	gameLogPre := ""
	gameLogSuf := ""
	colorPre := ""
	colorSuf := ""

	switch format {
	case "error":
		gameLogPre = "(ERROR): "
		gameLogSuf = ""
	case "h1":
		gameLogPre = "--- "
		gameLogSuf = " ---"
	case "h2":
		gameLogPre = "-- "
		gameLogSuf = " --"
	case "h3":
		gameLogPre = "- "
		gameLogSuf = " -"
	case "input":
		gameLogPre = "> "
		gameLogSuf = ""
	}

	// Replace the colorPre & colorSuf values based on color.
	switch color {
	case "orange":
		colorPre = "["
		colorSuf = "](fg:orange)"
	case "purple":
		colorPre = "["
		colorSuf = "](fg:purple)"
	case "pink":
		colorPre = "["
		colorSuf = "](fg:pink)"
	case "8":
		colorPre = "["
		colorSuf = "](fg:8)"
	case "red":
		colorPre = "["
		colorSuf = "](fg:red)"
	case "blue":
		colorPre = "["
		colorSuf = "](fg:blue)"
	case "black":
		colorPre = "["
		colorSuf = "](fg:black)"
	case "cyan":
		colorPre = "["
		colorSuf = "](fg:cyan)"
	case "yellow":
		colorPre = "["
		colorSuf = "](fg:yellow)"
	case "white":
		colorPre = "["
		colorSuf = "](fg:white)"
	case "clear":
		colorPre = "["
		colorSuf = "](fg:clear)"
	case "green":
		colorPre = "["
		colorSuf = "](fg:green)"
	case "magenta":
		colorPre = "["
		colorSuf = "](fg:magenta)"
	}

	chunked := splitStringByMax(gameLogPre+s+gameLogSuf, termWidth-42)

	for i := 0; i < len(chunked); i++ {
		gameLog.Rows = append(gameLog.Rows, colorPre+chunked[i]+colorSuf)
	}
	gameLog.Rows = append(gameLog.Rows, "")
	gameLog.ScrollBottom()

	writeLogMarkdown(s, format)

	mgToggle = false
	ui.Render(gameLog)
}

// parseArgs accepts a string, splits it using " " as the delimiter, and determines if it
// is a valid command or not, sending everything after the first word as a list of args.
func parseArgs(s string) {
	if len(s) > 0 {
		all := strings.Fields(s)
		cmd := all[0]
		args := make([]string, len(all))

		copy(args, all[1:])

		if cmd == "log" {
			renderOutput("Log:", "input", "clear")
		} else {
			renderOutput(cmd+" "+strings.Join(args, " "), "input", "8")
		}

		switch strings.ToLower(cmd) {
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
		case "skill":
			cmdSkill(args)
		case "help":
			cmdHelp(args)
		case "lipsum":
			cmdLipsum(args)
		default:
			renderOutput("Invalid Command.", "error", "red")
		}
	}
}
