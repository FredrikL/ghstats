package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/fredrikl/ghstatus/internal/config"
	"github.com/fredrikl/ghstatus/internal/data"
	"github.com/google/go-github/v61/github"
	"golang.org/x/term"
)

var width int
var cfg *config.Config

func main() {
	ctx := context.Background()
	w, _, _ := term.GetSize(0)
	width = w

	renderBanner()

	cfg = config.GetConfig()
	d, _ := data.GetData(ctx, cfg, os.Getenv("GITHUB_TOKEN"))
	renderTableHeader("Assigned Issues")
	fmt.Print(renderIssues(d.AssignedIssues))
	renderTableHeader("Open PRs in watched repos")
	fmt.Print(renderPrs(d.OpenPrs))
}

func renderBanner() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Width(22).Padding(1).Width(width).Align(lipgloss.Center)
	fmt.Println(style.Render("Github overview"))
}

func renderTableHeader(title string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#545454")).
		Background(lipgloss.Color("#baed91")).
		Width(22).PaddingLeft(1).Width(width).Align(lipgloss.Center)
	fmt.Println(style.Render(title))
}

func renderLogin(pr data.PR) string {
	if pr.ByBot {
		return "🤖"
	}
	return pr.By
}

func removeOrg(repo string) string {
	for _, org := range cfg.OrgsToSkip {
		prefix := fmt.Sprintf("%s/", org)
		if strings.HasPrefix(repo, prefix) {
			return repo[len(prefix):]
		}
	}
	return repo
}

func renderPrs(prs []data.PR) string {
	var rows [][]string
	var nonBots [][]string
	var approved [][]string

	for _, pr := range prs {
		approveIcon := "🙅"
		if pr.Approved {
			approveIcon = "✅"
		}

		row := []string{removeOrg(pr.Repo), fmt.Sprintf("%s (#%d) %s", approveIcon, pr.Id, pr.Title), renderLogin(pr), pr.Url}

		switch {
		case !pr.ByBot:
			nonBots = append(nonBots, row)
		case pr.Approved:
			approved = append(approved, row)
		default:
			rows = append(rows, row)
		}
	}

	var all [][]string
	if len(nonBots) > 0 {
		all = nonBots
	}
	if len(approved) > 0 {
		if len(all) > 0 {
			all = append(all, []string{})
		}
		all = append(all, approved...)

	}
	if len(rows) > 0 {
		if len(all) > 0 {
			all = append(all, []string{})
		}
		all = append(all, rows...)
	}

	return renderTable([]string{"Repo", "Title", "By", "Url"}, all)
}

func renderIssues(issues []*github.Issue) string {
	var rows [][]string
	for _, issue := range issues {
		rows = append(rows, []string{*issue.Repository.Name, *issue.Title, *issue.User.Login, *issue.HTMLURL})
	}
	return renderTable([]string{"Repo", "Title", "Reporter", "Url"}, rows)
}

func renderTable(headers []string, rows [][]string) string {
	t := table.New().
		Width(width).
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers(headers...).
		Rows(rows...)

	return fmt.Sprintf("%s\n", t.Render())
}
