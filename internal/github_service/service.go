package github_service

import (
	"context"
	"strings"

	"github.com/google/go-github/v61/github"
)

type GithubService struct {
	client *github.Client
}

func NewGithubService(token string) GithubService {
	client := github.NewClient(nil).WithAuthToken(token)
	return GithubService{
		client: client,
	}
}

func (gs *GithubService) GetIssues(ctx context.Context) []*github.Issue {
	opts := &github.IssueListOptions{}
	issues, _, err := gs.client.Issues.List(ctx, true, opts)
	if err != nil {
		panic(err)
	}
	return issues
}

func (gs *GithubService) GetPullRequest(ctx context.Context, repo string) []*github.PullRequest {
	opts := &github.PullRequestListOptions{}
	p := strings.Split(repo, "/")
	prs, _, err := gs.client.PullRequests.List(ctx, p[0], p[1], opts)
	if err != nil {
		panic(err)
	}
	return prs
}

func (gs *GithubService) IsPullRequestApproved(ctx context.Context, repo string, id int) bool {
	p := strings.Split(repo, "/")
	reviews, _, err := gs.client.PullRequests.ListReviews(ctx, p[0], p[1], id, &github.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, r := range reviews {
		if *r.State == "APPROVED" {
			return true
		}
	}
	return false
}
