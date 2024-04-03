# Reboot Bot

This is a rad little go program that will send a "reboot" signal to Vultr, to quickly and easily reboot your instances via a web hook, cURL request, whatever.

## How to

Run `go get ...`

Copy the `.env.example` to `.env` and update the settings to whatever you want.

Make sure that the API key you get from Vultr is able to run from the IP address where you're running this program.

Run `build.sh` and you're good to go. Put that `dist/reboot-bot` on a linux server somewhere and make your life easier.

Huzzah!
