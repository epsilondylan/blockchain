#!/bin/bash

# Replace PORT_NUMBER with the actual port number you want to check
PORT_NUMBER=1230


# Check if any processes are using the specified port
if lsof -i :$PORT_NUMBER >/dev/null; then
    echo "Processes using port $PORT_NUMBER found. Terminating them..."
    kill -9 $(lsof -t -i :$PORT_NUMBER)
    echo "Processes terminated."
else
    echo "No processes found using port $PORT_NUMBER."
fi