#!/usr/bin/env bash

set -eu

MYSQL_USER=root
MYSQL_PASSWORD=password
MYSQL_DATABASE=playground
MYSQL_HOST=localhost
MYSQL_PORT=23306
MYSQL_SHOW_ALL_LOG=true

export MYSQL_USER
export MYSQL_PASSWORD
export MYSQL_DATABASE
export MYSQL_HOST
export MYSQL_PORT
export MYSQL_SHOW_ALL_LOG

exec "$@"
