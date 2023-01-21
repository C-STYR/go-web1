#!/bin/bash

# when building, go will not run test files, so *.go works fine here
go build -o bookings cmd/web/*.go
./bookings

# to make this script executable, in the command line enter: chmod +x run.sh
# then run it with: ./run.sh