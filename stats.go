package main

import (
	"context"
	"time"

	"github.com/google/go-github/v40/github"
	"golang.org/x/oauth2"
)

type Stats struct {
	Token    string
	User     string
	Repo     string
	SinceTag string
	UntilTag string
	Client   *github.Client
	ctx      context.Context
}

func (s *Stats) Init() {
	s.ctx = context.Background()
	if s.Token == "" {
		s.Client = github.NewClient(nil)
		return
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: s.Token},
	)
	tc := oauth2.NewClient(s.ctx, ts)
	s.Client = github.NewClient(tc)
}

func (s *Stats) getTags() ([]*github.RepositoryTag, error) {
	var allTags []*github.RepositoryTag
	opt := &github.ListOptions{PerPage: 100}
	for {
		tags, resp, err := s.Client.Repositories.ListTags(s.ctx, s.User, s.Repo, opt)
		if err != nil {
			return allTags, err
		}
		allTags = append(allTags, tags...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allTags, nil
}

func (s *Stats) getLatestTag() (string, error) {
	release, _, err := s.Client.Repositories.GetLatestRelease(s.ctx, s.User, s.Repo)
	if err != nil {
		return "", err
	}
	return release.GetTagName(), nil
}

func (s *Stats) getTagDate(tags []*github.RepositoryTag, tag string) (time.Time, error) {
	ts := time.Unix(0, 0)
	for _, t := range tags {
		if tag == t.GetName() {
			commit, _, err := s.Client.Repositories.GetCommit(s.ctx, s.User, s.Repo, t.GetCommit().GetSHA(), nil)
			if err != nil {
				return ts, err
			}
			ts = commit.GetCommit().GetCommitter().GetDate()
			return ts, nil
		}
	}
	return ts, nil
}

func (s *Stats) getCommitDate(sha string) (time.Time, error) {
	ts := time.Unix(0, 0)
	commit, _, err := s.Client.Repositories.GetCommit(s.ctx, s.User, s.Repo, sha, nil)
	if err != nil {
		return ts, err
	}
	ts = commit.GetCommit().GetCommitter().GetDate()
	return ts, nil
}

func (s *Stats) getTagDates() (time.Time, time.Time, error) {
	since := time.Unix(0, 0)
	until := time.Unix(0, 0)

	tags, err := s.getTags()
	if err != nil {
		return since, until, err
	}

	// if we haven't specified a tag, take the last one
	if s.SinceTag == "" {
		if s.SinceTag, err = s.getLatestTag(); err != nil {
			return since, until, err
		}
	}

	since, err = s.getTagDate(tags, s.SinceTag)
	if err != nil {
		return since, until, err
	}
	if since == time.Unix(0, 0) {
		if since, err = s.getCommitDate(s.SinceTag); err != nil {
			return since, until, err
		}
	}
	if s.SinceTag == "" {
		if s.SinceTag, err = s.getLatestTag(); err != nil {
			return since, until, err
		}
	}
	until, err = s.getTagDate(tags, s.UntilTag)
	if err != nil {
		return since, until, err
	}
	if until == time.Unix(0, 0) {
		if until, err = s.getCommitDate(s.UntilTag); err != nil {
			return since, until, err
		}
	}
	return since, until, nil
}

func (s *Stats) GetCommits() (map[string]*github.RepositoryCommit, error) {
	since, until, err := s.getTagDates()
	if err != nil {
		return nil, err
	}
	var commits []*github.RepositoryCommit
	opt := &github.CommitsListOptions{
		Since:       since,
		Until:       until,
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		cc, resp, err := s.Client.Repositories.ListCommits(s.ctx, s.User, s.Repo, opt)
		if err != nil {
			return nil, err
		}
		commits = append(commits, cc...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
		time.Sleep(time.Millisecond * 200)
	}
	m := make(map[string]*github.RepositoryCommit)
	for _, commit := range commits {
		m[commit.GetAuthor().GetLogin()] = commit
	}
	return m, nil
}
