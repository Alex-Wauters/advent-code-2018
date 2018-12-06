#!/usr/bin/env bash
awk '{s+=$0}END{print s}' input.txt