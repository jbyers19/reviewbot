# ReviewBot

ReviewBot is a Telegram bot that helps you to get reviews from your customers.

## Components
- `reviewbot`: server that handles the bot logic. It is written in go.
- `reviewbot_mgr`: REST API server + gRPC client that handles the bot management. It is written in python.

The `reviewbot` and `reviewbot_mgr` communicate via gRPC.
Users (businesses) can use the `reviewbot_mgr` API to define message templates and triggers for controlling how and when the bot responds to customers.

## Table of Contents
- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Starting the services](#starting-the-services)
  - [Using the services](#using-the-services)
  - [Creating message templates](#creating-message-templates)
  - [Updating message templates](#updating-message-templates)
  - [Deleting message templates](#deleting-message-templates)
  - [Message users](#message-users)

## Getting Started

Before you start, make sure you have the following installed and configured on your system. Also note that some of the commands may not work for you as I have only tested this on Mac.

### Requirements

- Docker
- Go v1.21
- Python v3.11 (and pip)
- A Telegram account

### Starting the services

1. Clone this repo: `git clone https://github.com:jbyers19/reviewbot.git && cd reviewbot`
1. Get an API token for the Telegram bot. Check out the [Telegram docs](https://core.telegram.org/bots/features#creating-a-new-bot) for instructions on how to do this.
1. Once you have the token, save it as an environment variable in your `rc` file (`.bashrc`, `.zshrc`, etc.)
   or you can save it just for this session by running `export TELEGRAM_API_KEY="your-key-here"`
1. Start the `reviewbot` and `reviewbot_mgr` Docker containers by running `make run`
1. Run `curl localhost:5000/` to test that the services started successfully. You should get a `200 OK` response with `{}` as the body.
1. When you are ready to stop the services, run `make stop`. This will kill and remove the `reviewbot` and `reviewbot_mgr` containers.

### Using the services

When a customer joins a chat with the bot, they are added to the bot's database.
In your Telegram app, join the bot's chat using the link provided when you generated the API key for the bot.

### Creating message templates

Message templates are used to compose messages sent by the bot to the customer.
Here is what you will send to the `reviewbot_mgr` to create or update a template:
```json
{
  // Template name
  "name": "my-template",
  // Template content. The only available arguments are {{.FirstName}} and {{.LastName}}.
  "content": "Hi, {{.FirstName}} {{.LastName}}!"
  // (Optional) If a customer sends a message with any of these words it will trigger the bot to send a message to them using this template.
  "triggers": ["hi", "hello", "/start"]
}
```

Create a new template by running:
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"name": "review-prompt", "content": "Hi, {{.FirstName}} {{.LastName}}! Please review your recent purchase on our website.", "triggers": ["bought", "purchase"]}' \
  localhost:5000/template
```

You should receive a successful response and with the template sent in the request.
Try it out by sending "I'm happy with my purchase." to the bot on Telegram.
You should get a message back saying: `Hi, <your name>! Please review your recent purchase on our website.`

### Updating message templates

You can update an existing template by sending basically the same request as above except with a `PUT` instead of a `POST` and modifying the content or triggers.
If you change the name though, it will create a new template instead of modifying the existing one.

You can see the current templates by running:
```bash
curl localhost:5000/template
```

### Deleting message templates

Delete a message template by running:
```bash
curl -X DELETE localhost:5000/template\?name=review-prompt
```

### Message users

You can also trigger message to users manually.
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"first_name": "<your first name>", "last_name": "<your last name>", "template_name": "review-prompt"}' \
  localhost:5000/message
```

If the user isn't in the database or the template doesn't exist, the message will fail.
