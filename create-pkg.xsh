#!/usr/bin/env xonsh

$RAISE_SUBPROC_ERROR = True

from datetime import datetime

date_str = datetime.now().strftime("%d-%m-%Y")
header = "package solution"

mkdir -p @(date_str)
touch @(date_str)/README.md
echo @(header) > @(date_str)/solution.go
echo @(header) > @(date_str)/solution_test.go