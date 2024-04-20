package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GitHubClient encapsulates the GitHub API client and its methods.
type GitHubClient struct {
	client *github.Client
	ctx    context.Context
}

// Webhook represents a GitHub webhook for a repository.
type Webhook struct {
	RepositoryName string   `json:"repository_name"`
	RepositoryURL  string   `json:"repository_url"`
	Webhooks       []string `json:"webhooks"`
}

// NewGitHubClient creates a new GitHubClient with the provided access token.
func NewGitHubClient(accessToken string) *GitHubClient {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return &GitHubClient{
		client: client,
		ctx:    ctx,
	}
}

func main() {
	// Read GitHub access token from environment variable
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("GitHub access token not provided.. Check README")
		os.Exit(1)
	}

	// Read organization name from environment variable
	org := os.Getenv("GITHUB_ORG")
	if org == "" {
		fmt.Println("GitHub organization name not provided.. Check README")
		os.Exit(1)
	}

	// Create a GitHub client
	client := NewGitHubClient(token)

	// Print rate limits
	fmt.Println("Rate limits before listing repositories")
	client.PrintRateLimits()

	// List all repositories in the organization
	repos, err := client.ListRepositories(org)
	if err != nil {
		PrintErrors("Error: %v\n", err)
		os.Exit(1)
	}

	// Print rate limits
	fmt.Println("Rate limits before calling webhooks")
	client.PrintRateLimits()

	var webhooks []Webhook
	var wg sync.WaitGroup
	ch := make(chan Webhook, len(repos))

	// Fetch webhooks for each repository in parallel
	for _, repo := range repos {
		wg.Add(1)
		go client.ListWebhooks(org, repo.GetName(), &wg, ch)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(ch)

	// Collect results from the channel
	for webhook := range ch {
		webhooks = append(webhooks, webhook)
	}

	// Print rate limits
	fmt.Println("Rate limits after calling webhooks")
	client.PrintRateLimits()

	// Write the webhook information to a JSON file
	file, err := os.Create("webhooks.json")
	if err != nil {
		PrintErrors("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(webhooks); err != nil {
		PrintErrors("Error encoding JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Webhook information saved to webhooks.json")
}
