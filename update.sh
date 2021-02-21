#!/bin/sh
wget https://github.com/nProgrammer/godemon/releases/download/2.6.0/godemon
sudo rm -r /bin/godemon
sudo mv godemon /bin/godemon
sudo chmod 777 /bin/godemon
