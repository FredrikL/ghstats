package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/fredrikl/ghstatus/internal/github_service"
)

func main() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(2).
		PaddingLeft(4).
		Width(22)
	fmt.Println(style.Render("Github Status"))

	gh := github_service.NewGithubService("")
	gh.GetRepoStatus(context.Background(), "")
}
