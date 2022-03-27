[![Go](https://github.com/Pravardhitha/git-break/actions/workflows/go.yml/badge.svg)](https://github.com/Pravardhitha/git-break/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Pravardhitha/git-break/src)](https://goreportcard.com/report/github.com/Pravardhitha/git-break/src)

# git-break
repo on github designed to break other repos on github


Demo: https://git-break.tech


## :pushpin: Motivation 
Security is typically an after thought, we are trying to make it just a regular thought

## :memo: Description
Git-break is a tool which helps to defend and protect your hard work. 
We offer 2 services: defend and exploit. 
Worried about your code being manipulated? 
We'll look out for you and send you an alert! 


## :hammer: How to Build

### Make
```sh
cd src

# make "auto" builds the program and runs it 
make
```
### Raw Go
```sh
cd src

# default go run program to get main.go to run
go build main.go
```

## :alembic: Usage 
- Check out the demo, build and go to `localhost:9994` on your browser of choice

## :microscope: Technologies
- Language(s): `go`, `python`, `html`, `css`
- API(s): `github`, `auth0`, `twilio`
- Technologies: `nginx`, `systemd unit rules`, `make`, `ufw`

### In-depth
- Github
    - Github Projects (Agile project using an automated Kanban)
    - Github Actions
    - Github Issues
    - Github PRs
    - Github Code
    - Github API (Core of our app)
    - Insights (Forks)
    - Github Branch Protections
    - Github CODEOWNERS
    - Github Citations
- Twilio
    - Send SMS (to alert users)
- Auth0
    - Something
- Domain
    - Our dope domain! :shipit:
- DigitalOcean
    - Droplet
    - Networking
    - 'Implicit' VPC
    - Firewall

## :card_file_box: Directory Explanation
| Directory | Explanation
| :---:     | :---
| [.github](.github) | CI/CD to ensure that go program and all our dependencies are working like they should   
| [confs](confs) | Server configurations so we get that swanky domain
| [slideshow](slideshow) | Slideshow about our dope proj
| [src](src) | Source code!

## :blue_book: Technical Details
Need `.env` file in `/src` with the following fields:
```txt
TWILIO_ACCOUNT_SID="<>"
TWILIO_AUTH_TOKEN="<>"
TWILIO_PHONE_NUMBER="<>"
TO_PHONE_NUMBER="<>"
```

## :books: Resources
- [Twilio SendSMS via Go](https://www.twilio.com/blog/send-sms-30-seconds-golang)
- [Go env](https://zetcode.com/golang/env/)

[![DigitalOcean Referral Badge](https://web-platforms.sfo2.digitaloceanspaces.com/WWW/Badge%203.svg)](https://www.digitalocean.com/?refcode=a2a8208fca48&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge)
