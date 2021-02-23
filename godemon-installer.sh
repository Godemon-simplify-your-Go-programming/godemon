#!/bin/sh

printf "**********_________________***********\n*********/*****************\**********\n********/*******************\*********\n*******/_____________________\********\n*********WELCOME TO GODEMON***********\n*************INSTALLER****************\n\n"

printf "What do you want to do? \n1. Install Godemon \n2. Update Godemon \n"
printf "\nAnswer: "
read OPTION
printf "\nDo you want to do this: \n1. Global \n2. Local \n"
printf "\nAnswer: "
read GL

wget https://github.com/nProgrammer/godemon/releases/download/2.7.0/godemon

if [ "$OPTION" = "1" ]
then
  if [ "$GL" = "2" ]
  then
    mkdir ~/.godemon
    mkdir ~/.godemon/bin/
    sudo mv godemon ~/.godemon/bin/
    sudo chmod 777 ~/.godemon/bin/godemon
  elif [ "$GL" = "1" ]
  then
    sudo mkdir /usr/local/.godemon
    sudo mkdir /usr/local/.godemon/bin/
    sudo mv godemon /usr/local/.godemon/bin/
    sudo chmod 777 /usr/local/.godemon/bin/godemon
  fi
elif [ "$OPTION" = "2" ]
then
  if [ "$GL" = "2" ]
  then
    sudo mv godemon ~/.godemon/bin/
    sudo chmod 777 ~/.godemon/bin/godemon
  elif [ "$GL" = "1" ]
  then
    sudo mv godemon /usr/local/.godemon/bin/
    sudo chmod 777 /usr/local/.godemon/bin/godemon
  fi
fi

printf "\n Everything done \n"
if [ "$OPTION" = "1" ]
then
  if [ "$GL" = "2" ]
  then
    printf "\n Now add to .bashrc following line: export PATH=\$PATH:~/.godemon/bin"
  elif [ "$GL" = "1" ]
  then
    printf "\n Now add to .bashrc following line: export PATH=\$PATH:/usr/local/.godemon/bin"
  fi
fi
