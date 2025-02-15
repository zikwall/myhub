#!/bin/bash

SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"

if [ ! -f /etc/systemd/system/myhub.service ]; then
    echo "systemd service file not found, create new service (/etc/systemd/system/myhub.service)"
    cp "$SCRIPT_DIR"/myhub.service /etc/systemd/system/
fi

"$SCRIPT_DIR"/build.sh
"$SCRIPT_DIR"/run.sh