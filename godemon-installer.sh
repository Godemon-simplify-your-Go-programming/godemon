#!/bin/sh

printf "What do you want to do? \n1. Install Godemon \n2. Update Godemon \n"
read OPTION

if [ "$OPTION" = "1" ]
then
  wget https://github.com/nProgrammer/godemon/releases/download/2.6.1/godemon
  sudo cp -r godemon /bin/godemon
  sudo chmod 777 /bin/godemon
elif [ "$OPTION" = "2" ]
then
  wget https://github.com/nProgrammer/godemon/releases/download/2.6.1/godemon
  sudo rm -r /bin/godemon
  sudo mv godemon /bin/godemon
  sudo chmod 777 /bin/godemon
fi

printf "\n Everything done \n"