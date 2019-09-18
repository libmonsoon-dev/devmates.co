#!/usr/bin/env xonsh

$RAISE_SUBPROC_ERROR = True

from datetime import datetime

date_str = datetime.now().strftime("%d-%m-%Y")

mkdir @(date_str)
touch @(date_str)/README.md
touch @(date_str)/solution.go
touch @(date_str)/solution_test.go