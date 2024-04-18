package github_service

import (
	"context"

	"github.com/google/go-github/v59/github"
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

func (gs GithubService) GetRepoStatus(ctx context.Context, name string) {
	// gs.client.Issues.ListAssignees()
}
