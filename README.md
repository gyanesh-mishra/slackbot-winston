# Winston

Winston is your trusty slackbot here to service all your questions and needs.

At its very core, it's a simple key-value store backed by a mongo-database.

It's heavily under development, but winston uses Natural Language Processing under the hood to be a bit smarter.

I'm leveraging [Prove V.2](https://github.com/jdkato/prose), to parse out relevant speech tokens and respond accordingly.
Currently if you teach winston, "Can you tell me X", it's able to respond to "Do you know about X", "Where can I find info about X", "What is X" and a few other combinations with the same answer.

## Language and Frameworks

The project uses [Golang](https://golang.org/) for the API and [MongoDB](https://www.mongodb.com/) for the database.

All aspects of the application runs in [Docker containers](https://www.docker.com/)
