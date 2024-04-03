#! /bin/bash

killall reboot-bot

CGO_ENABLED=0 go build -o dist/reboot-bot

dist/reboot-bot &
