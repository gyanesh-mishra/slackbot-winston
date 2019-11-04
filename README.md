# Winston [![Codacy Badge](https://api.codacy.com/project/badge/Grade/3ff58876c279451eb5ef1367e2d6aa0b)](https://www.codacy.com/manual/gyanesh-mishra/slackbot-winston?utm_source=github.com&utm_medium=referral&utm_content=gyanesh-mishra/slackbot-winston&utm_campaign=Badge_Grade) [![Go Report Card](https://goreportcard.com/badge/github.com/gyanesh-mishra/slackbot-winston)](https://goreportcard.com/report/github.com/gyanesh-mishra/slackbot-winston)

Winston is your trusty slackbot here to service all your questions and needs.

At its very core, it's a simple key-value store backed by a mongo-database.

It's heavily under development, but winston uses Natural Language Processing under the hood to be a bit smarter.

Winston uses [Prove V.2](https://github.com/jdkato/prose) to parse out relevant speech tokens and respond accordingly.
Currently if you teach winston, "Can you tell me X", it's able to respond to "Do you know about X", "Where can I find info about X", "What is X" and a few other combinations with the same answer.

## Language and Libraries

The project uses [Golang](https://golang.org/) for the API and [MongoDB](https://www.mongodb.com/) for the database.

All aspects of the application runs in [Docker containers](https://www.docker.com/).

For local development, it's using [Air](https://github.com/cosmtrek/air) for Live-reloading.

Here's a list of libraries that are being used:

- [HTTPRouter](https://github.com/julienschmidt/httprouter)
- [Slack](https://github.com/nlopes/slack)
- [Viper](https://github.com/spf13/viper)
- [Mongo](https://github.com/mongodb/mongo-go-driver)
- [Prose V.2](https://github.com/jdkato/prose)

## Development

To get started you need Docker and Docker-Compose installed.
Look at Makefile for more information. But you can get started by running `make app`.
It will build the images, run the database and application container.

- The App container runs on Port 8080
- Mongo runs on port 27017
- Mongo Express runs on port 8081

## Currently in Progress

1.  Personality (Add some more human prompts to responses)
2.  Slash commands for slack
3.  Logo and branding
4.  Super easy to teach

## Upcoming

1.  Website (AngularJS)
2.  Tests (Unit/EE)
