# Brice's Email Bot

This is an email bot made for fun.

## Getting Started

```
git clone https://github.com/btald1331/EmailBot.git
```


### Prerequisites

```
Unix
```

### Installing

```
cd emailbot
go build mail.go
```

## Running

```
Usage of /var/folders/zn/mr3lxrws7f9g9dq5sx10l0c00000gn/T/go-build236959573/command-line-arguments/_obj/exe/mail:
  -interval int
    	Send email every X hours. (default 48)
  -logpath string
    	Path to output logs. (default "/var/log/emailbot.log")
  -message string
    	Body of email. (default "Hi. My name is Brice Aldrich. This is a bot, beep bop boop. This message was delivered by a program I wrote in Golang. I can write code. I can dev. I can hang. I know 5 Spanish words. One being contratame. This will execute every 2 days until contact has been made.\n\n Source: https://github.com/btald1331/EmailBot.git")
  -rotate int
    	Rotate file after X megabytes. (default 1000)
  -stop int
    	Stop sending emails after X hours. (default 168)



./mail --logpath=/var/log/mylog.log --interval=72 --stop=168 --message="My email message." --rotate=100

```


## Authors

* **Brice Aldrich**

## License

This project is licensed under the MIT License
