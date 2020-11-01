// A very simple and easy to use fancy terminal output library.
package iterminal

import (
	"fmt"
)

var ForegroundColor = foregroundColors()
var BackgroundColor = backgroundColors()
var Font = fonts()

type fColorList struct {
	Default    string
	Red        string
	Green      string
	LightBlue  string
	Black      string
	White      string
	Yellow     string
	Purple     string
	Beige      string
	Grey       string
	Cyan       string
	Pink       string
	LightRed   string
	LightGreen string
	Blue       string
}

type bColorList struct {
	Default     string
	Black       string
	Red         string
	Green       string
	Yellow      string
	Blue        string
	Purple      string
	Beige       string
	White       string
	LightRed    string
	LightGreen  string
	LightYellow string
	LightBlue   string
	Pink        string
	LightBeige  string
	Grey        string
	LightGrey   string
}

type fList struct {
	Default   string
	Bold      string
	Italic    string
	Underline string
	Blinking  string
}

func backgroundColors() *bColorList {
	return &bColorList{
		Default:     "\033[40m",
		Black:       "\033[40m",
		Red:         "\033[41m",
		Green:       "\033[42m",
		Yellow:      "\033[43m",
		Blue:        "\033[44m",
		Purple:      "\033[45m",
		Beige:       "\033[46m",
		LightGrey:   "\033[47m",
		LightBlue:   "\033[104m",
		LightRed:    "\033[101m",
		LightGreen:  "\033[102m",
		LightYellow: "\033[103m",
		Pink:        "\033[105m",
		LightBeige:  "\033[106m",
		Grey:        "\033[107m",
		White:       "\033[107m",
	}
}

func foregroundColors() *fColorList {
	return &fColorList{
		Default:    "\033[0m",
		Red:        "\033[31m",
		Green:      "\033[32m",
		Blue:       "\033[34m",
		Black:      "\033[30m",
		White:      "\033[37m",
		Yellow:     "\033[33m",
		Purple:     "\033[35m",
		Beige:      "\033[96m",
		Grey:       "\033[90m",
		Cyan:       "\033[36m",
		Pink:       "\033[95m",
		LightRed:   "\033[91m",
		LightGreen: "\033[92m",
		LightBlue:  "\033[94m",
	}
}

func fonts() *fList {
	return &fList{
		Default:   "\033[0m",
		Bold:      "\033[1m",
		Italic:    "\033[3m",
		Underline: "\033[4m",
		Blinking:  "\033[5m",
	}
}

func SetColor(fg, bg string) {
	fmt.Print(fg + bg)
}

func MakeColor(fg, bg string) string {
	return fg + bg
}

func SetFont(f string) {
	fmt.Print(f)
}

func MakeFont(f string) string {
	return f
}

func MoveCursor(x, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func ResetFormatting() {
	fmt.Print("\033[0m")
}

func ClearScreen() {
	fmt.Print("\033[2J")
	MoveCursor(0, 0)
}
