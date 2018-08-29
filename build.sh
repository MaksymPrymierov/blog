#!/bin/bash
if [[ $1 == "" || $1 == "help" ]]
then
  printf 'This sctipt launch and build server.\n\n'
  printf 'Examples - [  $ ./build "option" ]\n\n'
  printf 'Options: \n\n'
  printf 'start   - execute server\n'
  printf 'build   - execute "go build" command\n'
  printf 'install - exetute "go install" command\n'
  printf 'run     - exetute "go run" command\n'
fi
if [[ $1 == "start" ]]
then
  printf 'Launch server...\n'
  ./main
  if (( $? != 0 ))
  then
    printf "[ERROR] Launch server exit which error code $?"
    exit $?
  fi
fi
if [[ $1 == "build" ]]
then
  printf 'Start compile server...'
  go build main.go
  if (( $? != 0 ))
  then
    printf "[ERROR] Compile server exit which error code $?"
    exit $?
  fi
fi
if [[ $1 == "run" ]]
then
  printf 'Run server...'
  go run main.go
  if (( $? != 0 ))
  then
    printf "[ERROR] Run server exit which error code $?"
    exit $?
  fi
fi
if [[ $1 == "install" ]]
then
  printf 'Install server...'
  go install main.go
  if (( $? != 0 ))
  then
    printf "[ERROR] Install server exit which error code $?"
    exit $?
  fi
fi
exit 0
