# Discord ChatOps Bot



## Description
This is a discord bot project which will be written in Golang.

This bot will be used to trigger the pipeline of the disgobot project in GitLab and return the result back to Discord.

2023/06/05 Update: This project is not a generic tool yet. It is just tools I wrote for a certain project. So this repo is more like a showcase, not a repo for developing a production tool.

## Installation
- Clone this project first
- Excute `go mod download` command under the root folder of this project

## Usage
This bot will fetch settings from environment variable.

You need to create a .env file in project root for environment variable settings.

The .env file should look like as below:
```
DISCORD_TOKEN=<Fill in the bot's token>
GUILD_ID=<Fill in the id of the target guild>
GITLAB_BASE_URL=<Fill in the base url of the target gitlab>
CI_TRIGGER_TOKEN=<Fill in the token of the pipeline trigger>
PROJECT_ID=<Fill in the ID of the project belonged the pipeline trigger>
```

Then, you can use `go build cmd\chatopsbot\main.go` to build the executable and run it.

Make sure the .env file and the assets folder are in the same directory as the executable you just build.

## TODO List
- [ ] Refactor pkg/restapi/gitlab/pipeline_trigger.go TriggerPipeline Function
- [ ] Refactor internal/chatopsbot/slashcmd/slash_cmd_register_helper.go RegisterAllSlashCommands Function
- [ ] Add unit test for pkg/restapi/gitlab
- [X] Add job scripts with discord channel callback to .gitlab-ci.yml in the disgobot project
- [ ] Make documentation
