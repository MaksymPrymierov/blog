#!/bin/bash
go build
if (( $? != 0 ))
then
  echo "Сборка завершилась с ошибкой"
  exit $?
else
  echo "Сборка удачно завершалась"
fi
if [[ $1 == "-s" ]]
then
  echo "Запуск сервера..."
  ./blog
fi
if [[ $1 == "--80" ]]
then
  echo "Запуск сервера от имени суперпользователя"
  sudo export PORT=80
  sudo export GOPATH=/home/connor41/go
  ./blog
fi
