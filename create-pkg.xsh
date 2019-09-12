#!/usr/bin/env xonsh

from datetime import datetime

mkdir @(datetime.now().strftime("%d-%m-%Y")) && touch @(datetime.now().strftime("%d-%m-%Y"))/README.md