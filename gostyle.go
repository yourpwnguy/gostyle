package gostyle

import "fmt"

// gostyle represents the style library
type gostyle struct{}

// New creates a new instance of the gostyle library
func New() *gostyle {
	return &gostyle{}
}

// ApplyBlink applies the blink effect to the given message
func (gs *gostyle) ApplyBlink(msg string) string {
	return fmt.Sprintf("\x1b[5m%s\x1b[0m", msg) // Start blink, then reset
}

// ApplyReverse applies the reverse effect to the given message
func (gs *gostyle) ApplyReverse(msg string) string {
	return fmt.Sprintf("\x1b[7m%s\x1b[0m", msg) // Start reverse, then reset
}

// ApplyHidden applies the hidden effect to the given message
func (gs *gostyle) ApplyHidden(msg string) string {
	return fmt.Sprintf("\x1b[8m%s\x1b[0m", msg) // Start hidden, then reset
}

// LogInfo prints an informational message with a yellow prefix
func (gs *gostyle) LogInfo(msg string) string {
	return gs.Log("[!] ", "33", msg) // Yellow
}

// LogQue prints a question message with a cyan prefix
func (gs *gostyle) LogQue(msg string) string {
	return gs.Log("[?] ", "36", msg) // Cyan
}

// LogBad prints a bad message with a red prefix
func (gs *gostyle) LogBad(msg string) string {
	return gs.Log("[-] ", "31", msg) // Red
}

// LogGood prints a good message with a green prefix
func (gs *gostyle) LogGood(msg string) string {
	return gs.Log("[+] ", "32", msg) // Green
}

// LogRun prints a running message with a blue prefix
func (gs *gostyle) LogRun(msg string) string {
	return gs.Log("[~] ", "34", msg) // Blue
}

// Log prints a log message with the given prefix, color code, and message
func (gs *gostyle) Log(prefix, colorCode, msg string) string {
	// Construct and return the styled log message
	return fmt.Sprintf("\x1b[%sm%s%s\x1b[0m", colorCode, prefix, msg)
}

// Define color functions
func (gs *gostyle) Black(msg string) string {
	return fmt.Sprintf("\x1b[30m%s\x1b[0m", msg)
}

func (gs *gostyle) Red(msg string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", msg)
}

func (gs *gostyle) Green(msg string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", msg)
}

func (gs *gostyle) Yellow(msg string) string {
	return fmt.Sprintf("\x1b[33m%s\x1b[0m", msg)
}

func (gs *gostyle) Blue(msg string) string {
	return fmt.Sprintf("\x1b[34m%s\x1b[0m", msg)
}

func (gs *gostyle) Magenta(msg string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", msg)
}

func (gs *gostyle) Cyan(msg string) string {
	return fmt.Sprintf("\x1b[36m%s\x1b[0m", msg)
}

func (gs *gostyle) Gray(msg string) string {
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", msg)
}

func (gs *gostyle) White(msg string) string {
	return fmt.Sprintf("\x1b[97m%s\x1b[0m", msg)
}

// Define bright color functions
func (gs *gostyle) BrBlack(msg string) string {
	return fmt.Sprintf("\x1b[90m%s\x1b[0m", msg)
}

func (gs *gostyle) BrRed(msg string) string {
	return fmt.Sprintf("\x1b[91m%s\x1b[0m", msg)
}

func (gs *gostyle) BrGreen(msg string) string {
	return fmt.Sprintf("\x1b[92m%s\x1b[0m", msg)
}

func (gs *gostyle) BrYellow(msg string) string {
	return fmt.Sprintf("\x1b[93m%s\x1b[0m", msg)
}

func (gs *gostyle) BrBlue(msg string) string {
	return fmt.Sprintf("\x1b[94m%s\x1b[0m", msg)
}

func (gs *gostyle) BrMagenta(msg string) string {
	return fmt.Sprintf("\x1b[95m%s\x1b[0m", msg)
}

func (gs *gostyle) BrCyan(msg string) string {
	return fmt.Sprintf("\x1b[96m%s\x1b[0m", msg)
}

func (gs *gostyle) BrGray(msg string) string {
	return fmt.Sprintf("\x1b[97m%s\x1b[0m", msg)
}

// Define text style functions
func (gs *gostyle) Bold(msg string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[0m", msg)
}

func (gs *gostyle) Italic(msg string) string {
	return fmt.Sprintf("\x1b[3m%s\x1b[0m", msg)
}

func (gs *gostyle) Underline(msg string) string {
	return fmt.Sprintf("\x1b[4m%s\x1b[0m", msg)
}

func (gs *gostyle) Dim(msg string) string {
	return fmt.Sprintf("\x1b[2m%s\x1b[0m", msg)
}

func (gs *gostyle) Strike(msg string) string {
	return fmt.Sprintf("\x1b[9m%s\x1b[0m", msg)
}

// BGround applies background color to the given message
// Color parameter should be one of the available color strings
func (gs *gostyle) BGround(msg string, color string) string {
	var colorCode int
	switch color {
	case "black":
		colorCode = 40
	case "red":
		colorCode = 41
	case "green":
		colorCode = 42
	case "yellow":
		colorCode = 43
	case "blue":
		colorCode = 44
	case "magenta":
		colorCode = 45
	case "cyan":
		colorCode = 46
	case "white":
		colorCode = 47
	case "BrBlack":
		colorCode = 90
	case "BrRed":
		colorCode = 91
	case "BrGreen":
		colorCode = 92
	case "BrYellow":
		colorCode = 93
	case "BrBlue":
		colorCode = 94
	case "BrMagenta":
		colorCode = 95
	case "BrCyan":
		colorCode = 96
	case "BrWhite":
		colorCode = 97
	default:
		return msg // Return message as is if color not recognized
	}
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, msg)
}
