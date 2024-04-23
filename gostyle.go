package main

import (
	"fmt"
)

// GoStyle represents the style library
type GoStyle struct{}

// New creates a new instance of the GoStyle library
func New() *GoStyle {
	return &GoStyle{}
}

// Define maps for colors, styles, effects, and log types
var (
	colorMap = map[string]int{
		"black":     30,
		"red":       31,
		"green":     32,
		"yellow":    33,
		"blue":      34,
		"magenta":   35,
		"cyan":      36,
		"gray":      37,
		"white":     97,
		"BrBlack":   90,
		"BrRed":     91,
		"BrGreen":   92,
		"BrYellow":  93,
		"BrBlue":    94,
		"BrMagenta": 95,
		"BrCyan":    96,
		"BrGray":    97,
	}

	styleMap = map[string]int{
		"bold":      1,
		"dim":       2,
		"italic":    3,
		"underline": 4,
		"strike":    9,
	}

	effectMap = map[string]string{
		"blink":   "\x1b[5m",
		"reverse": "\x1b[7m",
		"hidden":  "\x1b[8m",
		"reset":   "\x1b[0m",
	}

	logTypeMap = map[string]struct {
		Prefix string
		Color  func(string) string
	}{
		"info": {Prefix: "[!] ", Color: New().Yellow},
		"que":  {Prefix: "[?] ", Color: New().Cyan},
		"bad":  {Prefix: "[-] ", Color: New().Red},
		"good": {Prefix: "[+] ", Color: New().Green},
		"run":  {Prefix: "[~] ", Color: New().Blue},
	}
)

// ApplyEffect applies the specified text effect to the given message
// Currently it includes effects such as blink, reversed, hidden.
func (gs *GoStyle) ApplyEffect(effect, msg string) string {
	code, ok := effectMap[effect]
	if !ok {
		fmt.Println("Invalid effect")
		return msg
	}
	return fmt.Sprintf("%s%s%s", code, msg, effectMap["reset"])
}

// Log prints a log message with the given type and message
func (gs *GoStyle) Log(logType, msg string) string {
	log, ok := logTypeMap[logType]
	if !ok {
		fmt.Println("Invalid log type")
		return ""
	}
	return fmt.Sprintf("%s%s", log.Color(log.Prefix), msg)
}

// Info prints an informational message
func (gs *GoStyle) Info(msg string) string {
	return gs.Log("info", msg)
}

// Que prints a question message
func (gs *GoStyle) Que(msg string) string {
	return gs.Log("que", msg)
}

// Bad prints a bad message
func (gs *GoStyle) Bad(msg string) string {
	return gs.Log("bad", msg)
}

// Good prints a good message
func (gs *GoStyle) Good(msg string) string {
	return gs.Log("good", msg)
}

// Run prints a running message
func (gs *GoStyle) Run(msg string) string {
	return gs.Log("run", msg)
}

// Red applies red color to the given message
func (gs *GoStyle) Red(msg string) string {
	return gs.color(msg, "red")
}

// Green applies green color to the given message
func (gs *GoStyle) Green(msg string) string {
	return gs.color(msg, "green")
}

// Blue applies blue color to the given message
func (gs *GoStyle) Blue(msg string) string {
	return gs.color(msg, "blue")
}

// Yellow applies yellow color to the given message
func (gs *GoStyle) Yellow(msg string) string {
	return gs.color(msg, "yellow")
}

// Magenta applies magenta color to the given message
func (gs *GoStyle) Magenta(msg string) string {
	return gs.color(msg, "magenta")
}

// Cyan applies cyan color to the given message
func (gs *GoStyle) Cyan(msg string) string {
	return gs.color(msg, "cyan")
}

// Gray applies gray color to the given message
func (gs *GoStyle) Gray(msg string) string {
	return gs.color(msg, "gray")
}

// White applies white color to the given message
func (gs *GoStyle) White(msg string) string {
	return gs.color(msg, "white")
}

// BrBlue applies bright blue color to the given message
func (gs *GoStyle) BrBlue(msg string) string {
	return gs.color(msg, "BrBlue")
}

// BrRed applies bright red color to the given message
func (gs *GoStyle) BrRed(msg string) string {
	return gs.color(msg, "BrRed")
}

// BrGreen applies bright green color to the given message
func (gs *GoStyle) BrGreen(msg string) string {
	return gs.color(msg, "BrGreen")
}

// BrYellow applies bright yellow color to the given message
func (gs *GoStyle) BrYellow(msg string) string {
	return gs.color(msg, "BrYellow")
}

// BrBlack applies bright black color to the given message
func (gs *GoStyle) BrBlack(msg string) string {
	return gs.color(msg, "BrBlack")
}

// BrMagenta applies bright magenta color to the given message
func (gs *GoStyle) BrMagenta(msg string) string {
	return gs.color(msg, "BrMagenta")
}

// BrCyan applies bright cyan color to the given message
func (gs *GoStyle) BrCyan(msg string) string {
	return gs.color(msg, "BrCyan")
}

// BrGray applies bright gray color to the given message
func (gs *GoStyle) BrGray(msg string) string {
	return gs.color(msg, "BrGray")
}

// Bold applies bold style to the given message
func (gs *GoStyle) Bold(msg string) string {
	return gs.style(msg, "bold")
}

// Italic applies italic style to the given message
func (gs *GoStyle) Italic(msg string) string {
	return gs.style(msg, "italic")
}

// Underline applies underline style to the given message
func (gs *GoStyle) Underline(msg string) string {
	return gs.style(msg, "underline")
}

// Dim applies dim style to the given message
func (gs *GoStyle) Dim(msg string) string {
	return gs.style(msg, "dim")
}

// Strike applies strike style to the given message
func (gs *GoStyle) Strike(msg string) string {
	return gs.style(msg, "strike")
}

// BGround applies bright background color to the given message
func (gs *GoStyle) BGround(msg string, color string) string {
	return gs.bg(msg, color)
}

// color applies color to the given message
func (gs *GoStyle) color(msg string, color string) string {
	colorCode := colorMap[color]
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, msg)
}

// bg applies background color to the given message
func (gs *GoStyle) bg(msg string, color string) string {
	colorCode := colorMap[color] + 10 // Background color codes are offset by 10
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, msg)
}

// style applies style to the given message
func (gs *GoStyle) style(msg string, styleStr string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", styleMap[styleStr], msg)
}

func main() {
	gs := New()

	// Example usage
	fmt.Println(gs.Info("This is an informational message"))
	fmt.Println(gs.Que("This is a question"))
	fmt.Println(gs.Bad("This is a bad message"))
	fmt.Println(gs.Good("This is a good message"))
	fmt.Println(gs.Run("This is a running message"))

	// Example usage of text effects
	fmt.Println(gs.ApplyEffect("blink", "Blinking Text"))
	fmt.Println(gs.ApplyEffect("reverse", "Reversed Text"))
	fmt.Println(gs.ApplyEffect("hidden", "Hidden Text"))
}
