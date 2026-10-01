#!/bin/sh
set -e

. /usr/share/debconf/confmodule
db_version 2.0

db_input high vigile-agent/key || true
db_go
