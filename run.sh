#!/bin/bash

go test -bench=. -count 10 -timeout 30m >results.txt

benchstat results.txt >pretty-results.txt