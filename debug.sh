#!/bin/bash
$SERIAL=$1
stty -F $SERIAL 9600
nc -ltp 4060 > /dev/null &
./deco/deco < $SERIAL | nc -t localhost 4060


function finish {
  # Your cleanup code here
  killall nc
}
trap finish EXIT