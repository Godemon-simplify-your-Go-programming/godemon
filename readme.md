# Godemon

##### Compatible with - Windows, Linux and MacOS

### What do you have to install before using `godemon`:
1. Go compiler,
2. Git,
3. Wget ( if you're using linux )
4. G++

### How to install Godemon?
When you're using Linux or MacOS you can choose the option of installing godemon using Godemon-installer,
the second way is dwonloading the source code from releases, and building it using Go compiler.

### How to use godemon-installer
```shell
git clone https://github.com/Godemon-simplify-your-Go-programming/Godemon-installer
cd Godemon-installer
g++ src/godemon-installer.cpp -o godemon-installer
./godemon-installer
```

### How to build source code?

```shell
unzip <ZIP-DIRECTORY-OF-GODEMON>
cd ./godemon
go build
mkdir ~/.godemon
mkdir ~/.godemon/bin
mv godemon ~/.godemon/bin/
```

### Now you need to add path to godemon to system variables

#### Unix system with zsh:
Add this line `export PATH=$PATH:~/.godemon/bin"` to file .zshenv

---

#### Unix system with bash:
Add this line `export PATH=$PATH:~/.godemon/bin"` to file .bashrc or .profile

---

### How to init Godemon project:
```shell
godemon -init -name=project
cd project
```

Now in project's directory godemon created 2 files - `go.mod` and `project.json`

### How to work with configuration json file
```json
{
   "name": "project",
   "arch": "amd64",
   "os": "darwin",
   "dev-vars": [
      {
         "key": "",
         "value": ""
      }
   ],
   "commands": [
      {
         "name": "run",
         "option": "mod",
         "path": ""
      }
   ],
   "files": null
}
```
We need to create new go file and add it to watchlist list in project.json
#### How to do it?
```shell
touch main.go
godemon -addFile -name=main -path=./main.go
```
Now godemon is watching changes of main.go file

### How to add command using CLI?
```shell
godemon -addFile -name=<name> -<mod/file> -path=<pathToFileOrModule>
```