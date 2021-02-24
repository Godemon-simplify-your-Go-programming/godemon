#!/bin/sh

function buildApp() {
    unzip godemon-2.7.1-beta.zip
    cd ./godemon-2.7.1-beta.zip
    go build
}

printf "**********_________________***********\n*********/*****************\**********\n********/*******************\*********\n*******/_____________________\********\n*********WELCOME TO GODEMON***********\n*************INSTALLER****************\n\n"

printf "What do you want to do? \n1. Install Godemon \n2. Update Godemon \n"
printf "\nAnswer: "
read OPTION
printf "\nDo you want to do this: \n1. Global \n2. Local \n"
printf "\nAnswer: "
read GL

wget https://github.com/nProgrammer/godemon/archive/2.7.1-beta.zip

if [ "$OPTION" = "1" ]
then
  if [ "$GL" = "2" ]
  then
    buildApp
    mkdir ~/.godemon
    mkdir ~/.godemon/logs/
    mkdir ~/.godemon/bin/
    sudo mv godemon ~/.godemon/bin/
    sudo chmod 777 ~/.godemon/bin/godemon
  elif [ "$GL" = "1" ]
  then
    buildApp
    sudo mkdir /usr/local/.godemon
    sudo mkdir ~/.godemon
    sudo mkdir ~/.godemon/logs/
    sudo mkdir /usr/local/.godemon/bin/
    sudo mv godemon /usr/local/.godemon/bin/
    sudo chmod 777 /usr/local/.godemon/bin/godemon
  fi
elif [ "$OPTION" = "2" ]
then
  if [ "$GL" = "2" ]
  then
    buildApp
    sudo mv godemon ~/.godemon/bin/
    sudo chmod 777 ~/.godemon/bin/godemon
  elif [ "$GL" = "1" ]
  then
    buildApp
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
