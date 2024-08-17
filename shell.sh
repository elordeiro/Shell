#!/bin/sh
set -e # Exit early if any commands fail
(
  cd "$(dirname "$0")"
  go build -o /tmp/shell-target shell/*.go
)
exec /tmp/shell-target
