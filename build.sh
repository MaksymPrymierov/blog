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
