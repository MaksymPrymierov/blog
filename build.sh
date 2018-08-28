#!/bin/bash
if [[ $1 == "" || $1 == "help" ]]
then
  printf 'This sctipt launch and build server.\n\n'
  printf 'Examples - [  $ ./build "option" ]\n\n'
  printf 'Options: \n\n'
  printf 'start   - execute "go run" command\n'
  printf 'build   - execute "go build" command\n'
  printf 'install - exeture "go install" command\n'
fi
if [[ $1 == "start" ]]
then
  printf 'Launch server...\n'
  go run blog
  if (( $? != 0 ))
  then
    printf "[ERROR] Lounch server exit which error code $?"
  fi
fi
