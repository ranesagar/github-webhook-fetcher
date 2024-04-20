package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/google/go-github/github"
)

// ListRepositories retrieves all repositories in the specified organization.
func (c *GitHubClient) ListRepositories(org string) ([]*github.Repository, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	var allRepos []*github.Repository
	for {
		repos, r, err := c.client.Repositories.ListByOrg(c.ctx, org, opt)
		if err != nil {
			return nil, fmt.Errorf("error listing repositories: %w", err)
		}
		allRepos = append(allRepos, repos...)
		if r.NextPage == 0 {
			break
		}
		opt.Page = r.NextPage
	}
	return allRepos, nil
}

// ListWebhooks retrieves all webhooks for the specified repository.
func (c *GitHubClient) ListWebhooks(org, repo string, wg *sync.WaitGroup, ch chan<- Webhook) {
	defer wg.Done()

	hooks, _, err := c.client.Repositories.ListHooks(c.ctx, org, repo, nil)
	if err != nil {
		PrintErrors("Error listing webhooks for %s: %v\n", repo, err)
		return
	}

	var webhookURLs []string
	for _, hook := range hooks {
		url, ok := hook.Config["url"].(string)
		if !ok {
			PrintErrors("Invalid webhook URL format")
			continue
		}
		webhookURLs = append(webhookURLs, url)
	}

	webhook := Webhook{
		RepositoryName: repo,
		RepositoryURL:  fmt.Sprintf("https://github.com/%s/%s", org, repo),
		Webhooks:       webhookURLs,
	}

	ch <- webhook
}

// Helper Functions

// PrintErrors prints error messages
func PrintErrors(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

// PrintRateLimits prints the current rate limits for all GitHub API endpoints.
func (c *GitHubClient) PrintRateLimits() {
	rateLimits, _, err := c.client.RateLimits(c.ctx)
	if err != nil {
		PrintErrors("Error getting rate limits: %v\n", err)
		return
	}

	fmt.Println("Rate Limits:")
	fmt.Printf("Core Limit: %d\n", rateLimits.GetCore().Limit)
	fmt.Printf("Core Remaining: %d\n", rateLimits.GetCore().Remaining)
	fmt.Printf("Search Limit: %d\n", rateLimits.GetSearch().Limit)
	fmt.Printf("Search Remaining: %d\n\n", rateLimits.GetSearch().Remaining)
}
