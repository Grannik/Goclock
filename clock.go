package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"syscall"
)

var font = map[rune][3]string{
	'0': {"┌─┐", "│ │", "└─┘"},
	'1': {" ┌┐", "  │", "  ┘"},
	'2': {"┌─┐", "┌─┘", "└─┘"},
	'3': {"┌─┐", " ─┤", "└─┘"},
	'4': {"┌ ┐", "└─┤", "  ┘"},
	'5': {"┌─┐", "└─┐", "└─┘"},
	'6': {"┌─┐", "├─┐", "└─┘"},
	'7': {"┌─┐", "  ┤", "  ┘"},
	'8': {"┌─┐", "├─┤", "└─┘"},
	'9': {"┌─┐", "└─┤", "└─┘"},
	':': {" ┌┐ ", " ├┤ ", " └┘ "},
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

func printPseudographicTime(t time.Time) {
	timeStr := t.Format("15:04:05")
	lines := [3]string{}

	for _, ch := range timeStr {
		charFont, ok := font[ch]
		if !ok {
			charFont = [3]string{"   ", "   ", "   "}
		}
		for i := 0; i < 3; i++ {
			if ch == ':' {
				lines[i] += "\033[90m\033[1m" + charFont[i] + "\033[0m"
			} else {
				lines[i] += "\033[37m\033[1m" + charFont[i] + "\033[0m"
			}
		}
	}

	fmt.Print("\033[H")
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	hideCursor()
	defer showCursor()
	clearScreen()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-sigs:
			clearScreen()
			return
		case t := <-ticker.C:
			printPseudographicTime(t)
		}
	}
}
