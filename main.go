package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/app"
	"github.com/umekikazuya/gh-notify/internal/github"
)

func main() {
	isAll := flag.Bool("all", false, "All repo")
	flag.Parse()
	p := tea.NewProgram(
		app.NewModel(
			100,
			*isAll,
			github.FindAll,
			github.MarkReadNotification,
			github.GetTreadHTMLURL,
		),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
