# github-webhook-fetcher

# GitHub Webhook Fetcher
Fetches a list of webhooks in a given organization

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub issues](https://img.shields.io/github/issues/ranesagar/github-webhook-fetcher.svg)](https://github.com/ranesagar/github-webhook-fetcher/issues)
[![GitHub forks](https://img.shields.io/github/forks/ranesagar/github-webhook-fetcher.svg)](https://github.com/ranesagar/github-webhook-fetcher/network)
[![GitHub stars](https://img.shields.io/github/stars/ranesagar/github-webhook-fetcher.svg)](https://github.com/ranesagar/github-webhook-fetcher/stargazers)

## Installation

### Prerequisites

Before you begin, ensure you have met the following requirements:
- go installed
- Organization name
- Personal access token for your organization

### Installing

To install GitHub Webhook Fetcher, follow these steps:

1.  Clone the repository.
```sh
git clone https://github.com/ranesagar/github-webhook-fetcher.git
```
2. Navigate to the project directory.

```sh
cd github-webhook-fetcher
```
3. Install dependencies.
```sh
go mod tidy
```

4. Build the executable.
```sh
go build
```


## Usage

To use GitHub Webhook Fetcher, follow these steps:

1. Set up your GitHub access token and organization name as environment variables.

```sh
export GITHUB_TOKEN="your_github_access_token"
export GITHUB_ORG="your_organization_name"
```




2. Step 2: Run the executable.


```sh
./github-webhook-fetcher
```


## Contributing

Contributions are what make the open source community such an amazing place to be, learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Contact

Project Link: [https://github.com/ranesagar/github-webhook-fetcher](https://github.com/ranesagar/github-webhook-fetcher)

## Acknowledgements

- [GitHub API Documentation](https://developer.github.com/v3/)
- [Shields.io](https://shields.io/)
- [Choose an Open Source License](https://choosealicense.com)
