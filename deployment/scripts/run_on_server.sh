#!/bin/sh
set -e

echo "Killing the old app"
if ! sudo killall -9 demo > /dev/null 2>&1; then
    echo "no app running" >&2
fi

if [[ -f /home/app/server/demo ]]; then
	rm /home/app/server/demo
	mv /home/app/server/tmp/demo /home/app/server/demo
fi


echo "Run the app"
nohup /home/app/server/demo >> /home/app/log/demo.log 2>&1 &

tail -f -n 10 /home/app/log/demo.log
