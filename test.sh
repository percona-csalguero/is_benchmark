#!/bin/bash
echo "MySQL 5.6 at port 3306";
echo "MySQL 5.7 at port 3307";
echo "MySQL 8.0.3-rc at port 3308";
for port in 3306 3307 3308; do
    for i in {1..5}; do
        echo "Testing instance at port $port. Run # ${i}"
        PORT=${port} go test -bench=.
    done
done
