# dbt-go
Open Source dbt CLI Wrapper that aims to increase DX for Analytics Engineers

## Roadmap

- add styling
- add issue template
- add pr template
- add CI testing
- add CD pipeline
- define release schedule
- figure out contributor agreement, currently using GPL3 license
- add simple `merge resolution` TUI to simplify merging main to a branch before CI
- Add commands to parse and run changed models and their children
- CI commands to simplify and stream line models triggered in CI
- add modules for specific adapters such as BigQuery, Snowflake
- dbt secure (run DCL,IAM,etc seperate from transformation commands)
- many more to come...


## dependencies

### Go
 - https://github.com/spf13/cobra-cli
 - https://github.com/charmbracelet/bubbletea (coming soon)
 - https://github.com/charmbracelet/lipgloss

### Tools
 - https://github.com/git/git
 - https://github.com/cli/cli
 - https://github.com/dbt-labs/dbt-core