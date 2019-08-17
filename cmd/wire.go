//+build wireinject

package cmd

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"github.com/google/wire"
	"github.com/rerost/issue-creator/domain/issue"
	"github.com/rerost/issue-creator/domain/schedule"
	"github.com/rerost/issue-creator/repo"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func CurrentTime(cfg Config) time.Time {
	return time.Now()
}

func NewGithubClient(ctx context.Context, cfg Config) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GithubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	c := github.NewClient(tc)
	return c
}

func NewK8sCommand(cfg Config) []string {
	return cfg.K8sCommands
}

func NewTemplateFile(cfg Config) string {
	return cfg.ManifestTemplateFile
}

func InitializeCmd(ctx context.Context, cfg Config) (*cobra.Command, error) {
	wire.Build(
		NewCmdRoot,
		issue.NewIssueService,
		repo.NewIssueRepository,
		CurrentTime,
		NewGithubClient,
		schedule.NewScheduleService,
		repo.NewScheduleRepository,
		NewK8sCommand,
		NewTemplateFile,
	)
	return nil, nil
}
