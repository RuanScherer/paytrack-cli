# Paytrack CLI

A command line interface for Paytrack devs. The goal of this project is to make it easier to do common tasks that spend some time or are error prone.

## Installation

1. Download the binary/executable for your OS from the [releases page](https://github.com/RuanScherer/paytrack-cli/releases)
2. Put the binary/executable in a folder
3. Add the folder to your PATH if it is not already there

Then, you're ready to go!

# Usage

Paytrack CLI can be invoked by running `paytrack` in your terminal.

```bash
$ paytrack
```

## Commands

### `paytrack` or `paytrack help`

Displays the help menu.

```bash
$ paytrack help
```

### `paytrack ui`

Displays the help menu for `paytrak ui` command.

```bash
$ paytrack ui
```

#### `paytrack ui rewriteLocal`

Copies the local build of paytrack-ui-library to `node_modules/paytrack-ui-library/dist` folder of specified frontend project.

Required params:
- `--frontend-project` or `-f`: Name of the folder of frontend project

This command uses the environment variable `PAYTRACK_CLI_SOURCE_PATH` to get the base path of your projects. If this variable is not set, it will use the following default values for each OS:
- Windows: `C:/git/paytrack/fontes/`
- OSX (Mac): `~/git/paytrack/fontes/`

```bash
$ paytrack ui rewriteLocal -f paytrack-frontend
```
