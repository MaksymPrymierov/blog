#!/bin/bash
go build main.go
if [[ $1 == "-s" ]]
then
  ./main
fi
