<img src=".img/logo.png" alt="dg" width="200">
<br><br>
An open source CLI wrapper for dbt/git that aims to increase DX for Analytics Engineers.

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Features](#features)
- [Installation on Unix Based Systems](#installation-on-unix-based-systems)
- [Usage](#usage)
  - [Commands](#commands)
- [Development](#development)
    - [Code Structure](#code-structure)
    - [Dependencies](#dependencies)
  - [Tools](#tools)
  - [Development Setup](#development-setup)
- [Roadmap](#roadmap)
    - [Project](#project)
    - [CLI](#cli)
- [License](#license)

## Features

- Improved Developer Experience (DX) for Analytics Engineers using dbt
- Simple interface to interact with version control
- Slick Interactive TUI (Terminal User Interface)

## Installation on Unix Based Systems

To install `dg`, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/cognite-analytics/dbt-go.git
    cd dbt-go
    ```

2. Build the project:
    ```sh
    make
    ```

## Usage

To use `dbt-go`, run the following command from the root of your dbt project:
```sh
dg
```

### Commands
`info`
Show additional developer information about dbt-go.

`ls`
List changed files on the current branch.

`vc`
Interactive Git Experience.

## Development

#### Code Structure
```
cmd/ - Core CLI Commands
git/ - Reusable Git Functions
style/ - Color & Styling Definitions
```

#### Dependencies

Go modules:

 - github.com/spf13/cobra
 - github.com/charmbracelet/bubbletea
 - github.com/charmbracelet/lipgloss
 - github.com/TheZoraiz/ascii-image-converter

### Tools
 - [Git](https://github.com/git/git)
 - [dbt](https://github.com/dbt-labs/dbt-core)

### Development Setup
- Install Go (version 1.23 or later).
- Clone the repository and navigate to the project directory.
- Build and install the project using the provided Makefile.

## Roadmap

#### Project

- add issue template
- add pr template
- add CI testing
- add CD pipeline
- define release schedule
- figure out contributor agreement, currently using GPL3 license


#### CLI
- `git`
  - add `merge resolution` TUI to simplify merging main in to a feature branch before pushing to remote
- `dbt`
    - Add commands to parse and run changed models instead of relying on dbt's graph operations
    - CI commands to simplify and stream line models triggered in CI and the CI process
    - add modules for specific adapters such as BigQuery, Snowflake
    - dbt secure (run DCL,IAM,etc seperate from transformation commands)
- `gh`
  - integrate `gh` cli to take advantage of it's auth/features

## License
This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.
