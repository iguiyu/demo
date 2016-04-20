#!/bin/sh
set -e

GOPATH="/home/app/gopath"
DEVPATH="/home/app/gopath/src/github.com/iguiyu/demo/deployment"

echo "Killing the old app"
if ! sudo killall -9 demo > /dev/null 2>&1; then
    echo "no app running" >&2
fi

# rm /home/app/server/demo
# mv /home/app/server/tmp/demo /home/app/server/demo

# if [[ -f /home/app/server/demo ]]; then
    echo "Remove old demo file..."
    rm /home/app/server/demo
    mv /home/app/server/tmp/demo /home/app/server/demo
# fi

cd  $DEVPATH/tmp/
tar xvf template.gz
echo "after tar xvf"
ls -al 
echo "to move templates"
rm -rf $DEVPATH/template
mv template $DEVPATH



echo "Run the app"
nohup /home/app/server/demo >> /home/app/log/demo.log 2>&1 &

tail -f -n 10 /home/app/log/demo.log
