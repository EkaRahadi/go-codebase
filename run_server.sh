#!/bin/bash

./httpserver &

# Keep the container running
tail -f /dev/null
