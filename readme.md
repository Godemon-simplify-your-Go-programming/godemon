# Godemon

### What do you have to install before using `godemon`:
1. Go compiler,
2. Git,
3. Wget ( if you're using linux )

### How to use it?
First of all you need to build your version:
`cd ./godemon/`
`go build`

### Using default `./godemon`
Now you can move file `godemon` to folder with your go project and 
start it using command: `./godemon -cnf=cnf -command=<command>` or `./godemon -cnf=cmd -path=<fileOrDir> -modOrFile=<fileOrModule>`.
In first argument you're giving godemon the path to project, and in second
argument you're choosing is it simple go file or go module - if module pass
argument - `mod` , if file pass argument - `file`

If you need help with clie use command `./godemon -help`

### Installing `godemon` with `installer.sh`
1. `wget https://github.com/nProgrammer/godemon/releases/download/1.1.0/installer.sh`
2. `sudo chmod 777 ./installer.sh`
3. `./installer.sh`

### Installing `godemon` from source code
1. Download source code from releases or use `git clone https://github.com/nProgrammer/godemon`
2. `cd ./godemon`
3. `go build`
4. `sudo cp godemon /bin/godemon`

### Installing `godemon` from `godemon.exe`
1. Download `godemon.exe`
2. Add path to `godemon.exe` to system PATH

### How to work with configuration json file
1. Create in project file `godemon-cnf.json`
2. Create configuration of commands - sample code: 

```
{
  "project": {
    "name": "api2",
    "platformOS": "linux",
    "platformArch": "amd64",
    "path": "/home/nwagner/Desktop/godemon/api/"
  },
  "commands": [
    {
      "name": "run",
      "path": "/home/nwagner/Desktop/godemon/api/",
      "file": "mod"
    },
    {
     "name": "run-single-file",
      "path": "/home/nwagner/Desktop/godemon/api/main.go",
      "option": "file"
    }
  ]
}
```
3. Now use command: `godemon -cnf=cnf -command=<command-name>` or if you want to make full binary
file you need to use command: `godemon -cnf=deploy`. This will create
   binary file with env variables. Deploy is using command - `env GOOS=project.platformOS
   GOARCH=projet.platformArch go build -o project.name`