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
1. `wget https://github.com/nProgrammer/godemon/releases/download/x.x.x/installer.sh`
2. `sudo chmod 777 ./installer.sh`
3. `./installer.sh`

### Installing `godemon` from source code
1. Download source code from releases or use `git clone https://github.com/nProgrammer/godemon`
2. `cd ./godemon`
3. `go build`
   
If Windows:
4. Move `godemon.exe` to `System32`
   
On Linux:
4. `sudo cp godemon /bin/godemon`

### Installing `godemon` from `godemon.exe`
1. Download `godemon.exe`
2. Add path to `godemon.exe` to system PATH

### How to work with configuration json file
1. Create json file `projects.json`
2. Create configuration of project - sample code:
```json
{
   "name": "test2",
   "arch": "amd64",
   "os": "linux",
   "path": "/home/nwagner/Desktop/godemon/test2",
   "dev-vars": [
      {
         "key": "PORT",
         "value": "8800"
      }
   ],
   "commands": [
      {
         "name": "run",
         "path": "/home/nwagner/Desktop/godemon/api/",
         "file": "mod"
      }
   ]
}
```
3. Now you can use 2 types of commands:
    
    a) `godemon -cnf=cnf <command-name>`
    
    b) `godemon -cnf=deploy`

### Automation of configs json files:
If you want to skip process of creating config file you can use command:

`godemon -init -name=<project-name> -arch=<sys-arch> -os=<os>`

#### Types of `-cnf`:

1. cnf - using `godemon-conf.json` to hot live reload
2. deploy - using `projects.json` to compile program for many platforms
3. cmd - to use this you need to pass parameters like - `-path -modOrFile ` allows you to easy hot live reloading

#### What is `dev-vars`?
`dev-vars` are OS variables that are using when godemon is running this app.

`"key"` - name of var

`"value"` - value of var
##### Remember! it doesn't works when you deploy your app