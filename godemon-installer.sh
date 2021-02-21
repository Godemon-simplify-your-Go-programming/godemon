#!/bin/sh

wget https://github.com/nProgrammer/godemon/releases/download/2.6.1/godemon

printf "What do you want to do? \n1. Install Godemon \n2. Update Godemon \n"
read OPTION

if [ "$OPTION" = "1" ]
then
  sudo mv -r godemon /bin/godemon
  sudo chmod 777 /bin/godemon
elif [ "$OPTION" = "2" ]
then
  sudo rm -r /bin/godemon
  sudo mv godemon /bin/godemon
  sudo chmod 777 /bin/godemon
fi

printf "\n Everything done \n"