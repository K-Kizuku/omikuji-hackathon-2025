package github

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/K-Kizuku/pymon-graphql/pkg/config"
	"github.com/google/go-github/v68/github"
	"golang.org/x/oauth2"
)

type IGithubService interface {
	ListRepositories(ctx context.Context, owner string) ([]*github.Repository, error)
	GetPythonFileChanges(ctx context.Context, owner, repo string) (additions, deletions int, err error)
	ExistRepository(ctx context.Context, owner, repo string) (bool, error)
}

type GithubService struct {
	client *github.Client
}

var _ IGithubService = (*GithubService)(nil)

func NewGithubClient(cfg *config.AppConfig) *GithubService {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GithubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &GithubService{github.NewClient(tc)}
}

func (g *GithubService) GetPythonFileChanges(ctx context.Context, owner string, repo string) (additions int, deletions int, err error) {
	commitOpts := &github.CommitsListOptions{
		Since: time.Now().Add(-24 * time.Hour),
		Until: time.Now(),
	}

	commits, _, err := g.client.Repositories.ListCommits(ctx, owner, repo, commitOpts)
	if err != nil {
		return 0, 0, err
	}

	for _, c := range commits {
		if c.SHA == nil {
			continue
		}
		commitSHA := *c.SHA

		commitDetail, _, err := g.client.Repositories.GetCommit(ctx, owner, repo, commitSHA, nil)
		if err != nil {
			log.Println("Failed to get commit detail:", err)
			continue
		}
		if commitDetail.Files == nil {
			continue
		}

		for _, f := range commitDetail.Files {
			filename := f.GetFilename()
			if strings.HasSuffix(filename, ".py") {
				add := f.GetAdditions()
				delete := f.GetDeletions()

				additions += add
				deletions += delete
			}
		}
	}
	return
}

func (g *GithubService) ListRepositories(ctx context.Context, owner string) ([]*github.Repository, error) {
	repos, _, err := g.client.Repositories.ListByUser(ctx, owner, nil)
	if err != nil {
		return []*github.Repository{}, err
	}
	return repos, nil
}

func (g *GithubService) ExistRepository(ctx context.Context, owner string, repo string) (bool, error) {
	_, _, err := g.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return false, err
	}
	return true, nil
}
