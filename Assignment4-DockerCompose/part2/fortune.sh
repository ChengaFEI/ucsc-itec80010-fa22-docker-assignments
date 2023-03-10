#!/bin/sh
trap "exit" INT
 
INTERVAL=${INTERVAL:-30}
OUTPUT_FILE=${1:-/var/local/output/quote}
 
echo "Docker Class Fortune Writer"
echo "Configured to write fortune to $OUTPUT_FILE every $INTERVAL seconds"
 
while :
do
  echo "$(date) Writing fortune..."
  echo "fortune" >> $OUTPUT_FILE
  sleep $INTERVAL
done

