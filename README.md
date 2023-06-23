# octo-shift
    
## Description

Octo Shift is a cli tool to assist with migrating from GitHub Enterprise Server to GitHub Enterprise Cloud. The CLI includes three commands to help with the migration process:

* create-teams
    * Creates teams from a source GitHub Enterprise Server organization in the target GitHub Enterprise Cloud organization and connects them to the External Groups in the Identity Provider.
* update-repo-visibility
    * Updates the visibility of repositories in the target GitHub Enterprise Cloud organization to match the visibility of the repositories in the source GitHub Enterprise Server organization.
* update-webhooks
    * Updates all of the Organization webhooks in the target GitHub Enterprise Cloud organization with a secret. Can also update the Repository webhooks in the target GitHub Enterprise Cloud organization with a secret.

## Installation
To install Octo-Shift, make sure you have the Go programming language installed on your system. You can download and install Go from the [official website](https://go.dev/doc/install).

Once you have Go installed, clone this repository and run the following command to install Octo-Reports:

```bash
go build octo-shift.go
```
This command will create an executable file called `octo-shift` in the current directory. You can move this file to a directory in your PATH to make it available system-wide.

## Usage

### create-teams

```bash
octo-shift create-teams -source-org <ghes org> -source-url <ghes url> -source-token <ghes token>  -target-org <ghec org> -target-token <ghec token>  
```

### update-repo-visibility

```bash
octo-shift update-repo-visibility -source-org <ghes org> -source-url <ghes url> -source-token <ghes token>  -target-org <ghec org> -target-token <ghec token>  
```

### update-webhooks

```bash
octo-shift update-webhooks -target-org <ghec org> -taget-token <ghec token> -secret <webhook secret> -include-repo-webhooks <true|false>
```

## Contributing
We welcome and appreciate contributions to Octo-Reports. If you'd like to contribute, please fork the repository and submit a pull request with your changes.

## License
Octo-Reports is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.