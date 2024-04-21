package data

import (
	"context"
	"sync"

	"github.com/fredrikl/ghstatus/internal/config"
	"github.com/fredrikl/ghstatus/internal/github_service"
	"github.com/google/go-github/v61/github"
)

type PR struct {
	Id       int
	Title    string
	By       string
	Approved bool
	ByBot    bool
	Url      string
	Repo     string
}

type Data struct {
	AssignedIssues []*github.Issue
	OpenPrs        []PR
}

func prByBot(login string, botNames []string) bool {
	for _, name := range botNames {
		if login == name {
			return true
		}
	}
	return false
}

func GetData(ctx context.Context, cfg *config.Config, token string) (*Data, error) {
	gh := github_service.NewGithubService(token)

	data := &Data{}
	data.OpenPrs = []PR{}

	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}

	wg.Add(1)
	go func() {
		data.AssignedIssues = gh.GetIssues(ctx)
		wg.Done()
	}()
	for _, repo := range cfg.Repo {
		wg.Add(1)
		go func() {
			var repoPrs []PR
			prs := gh.GetPullRequest(ctx, repo)

			for _, pr := range prs {
				repoPrs = append(repoPrs, PR{
					Id:       *pr.Number,
					Title:    *pr.Title,
					By:       *pr.User.Login,
					Approved: gh.IsPullRequestApproved(ctx, repo, *pr.Number),
					Url:      *pr.HTMLURL,
					Repo:     repo,
					ByBot:    prByBot(*pr.User.Login, cfg.Bots),
				})
			}
			mutex.Lock()
			data.OpenPrs = append(data.OpenPrs, repoPrs...)
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	return data, nil
}
