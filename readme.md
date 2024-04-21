# Github overview

CLI program that gives an overview of your current github status, displays issues assigned to you and open PRs in repositories you choose to monitor.

## Access token

Assumes that your shell has a [Personal access tokens](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens) in a enviroment variable named `GITHUB_TOKEN`

## Configuration

Place a file named `config.yaml`.

Properties

- `repos` array add all repositories to monitor for pull request status.
- `bots` name of bots that create pull requests
- `orgskip` name of users or organisations to skip when displaying repository name

```yaml
repos:
  - FredrikL/ghstats
bots:
  - dependabot[bot]
orgskip:
  - FredrikL
```

## Run

`go run .`
