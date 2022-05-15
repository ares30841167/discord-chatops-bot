# Discord ChatOps Bot



## Description
This is a discord bot project which will be written in Golang.

This bot will be used to trigger the pipeline of the disgobot project and return the result back to discord.

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

## Co-op Workflow
You shold follow the step below while you are working on this project:
1. Get the tasks you want to contribute from jira
2. Open an issue for it
3. Start working. At the same time, remember to update your status of your work on jira
4. When you are finish, open an merge request for it
5. Notify your teammates that you are already done with the task
6. (Optional) Let the reviewer of the MR review the code
7. Let the assignee of the MR review the code
8. If it all looks good, the assignee can merge the code into the branch

## Authors
The following are the developers who mainly maintain this project.
- [GUAN-YU CHEN](https://gitlab.guanyu.dev/ares30841167)

## Project status
Just started.

## TODO List
- [ ] Refactor pkg/restapi/gitlab/pipeline_trigger.go TriggerPipeline Function
- [ ] Refactor internal/chatopsbot/slashcmd/slash_cmd_register_helper.go RegisterAllSlashCommands Function
- [ ] Add unit test for pkg/restapi/gitlab
- [ ] Add job scripts with discord channel callback to .gitlab-ci.yml in the disgobot project
- [ ] Make documentation