#!/bin/sh

buildApp() {
    unzip 21.03.zip
    cd ./godemon-21.03
    go build
}

prepareDirs() {
    mkdir ~/.godemon
    mkdir ~/.godemon/logs/
    mkdir ~/.godemon/bin/
}

removing() {
  cd ..
  sudo rm -r godemon-21.03
  sudo rm -r 21.03.zip
}


printf "**********_________________***********\n*********/*****************\**********\n********/*******************\*********\n*******/_____________________\********\n*********WELCOME TO GODEMON***********\n*************INSTALLER****************\n\n"

printf "What do you want to do? \n1. Install Godemon \n2. Update Godemon \n"
printf "\nAnswer: "
read OPTION
printf "\nDo you want to do this: \n1. Global \n2. Local \n"
printf "\nAnswer: "
read GL

wget https://github.com/nProgrammer/godemon/archive/21.03.zip

if [ "$OPTION" = "1" ]
then
  if [ "$GL" = "2" ]
  then
    buildApp
    prepareDirs
    sudo mv godemon ~/.godemon/bin/
    sudo chmod 777 ~/.godemon/bin/godemon
    removing
  elif [ "$GL" = "1" ]
  then
    buildApp
    sudo mkdir /usr/local/.godemon
    prepareDirs
    sudo mkdir /usr/local/.godemon/bin/
    sudo mv godemon /usr/local/.godemon/bin/
    sudo chmod 777 /usr/local/.godemon/bin/godemon
    removing
  fi
elif [ "$OPTION" = "2" ]
then
  if [ "$GL" = "2" ]
  then
    buildApp
    sudo mv godemon ~/.godemon/bin/
    sudo chmod 777 ~/.godemon/bin/godemon
    removing
  elif [ "$GL" = "1" ]
  then
    buildApp
    sudo mv godemon /usr/local/.godemon/bin/
    sudo chmod 777 /usr/local/.godemon/bin/godemon
    removing
  fi
fi

printf "\n Everything done \n"
if [ "$OPTION" = "1" ]
then
  if [ "$GL" = "2" ]
  then
    printf "\nNow add to .bashrc following line: export PATH=\$PATH:~/.godemon/bin"
  elif [ "$GL" = "1" ]
  then
    printf "\nNow add to .bashrc following line: export PATH=\$PATH:/usr/local/.godemon/bin \nIf it's MacOS - add this line to .zshenv"
  fi
fi
