# I'm very sorry, I've got a probelm with Windows support

---

# Godemon

##### Compatible with - Windows, Linux and MacOS

---

### URL to guide:

https://drive.google.com/file/d/18jIB-47JwSrOb5404U9mkt3pHhrAiack/view?usp=sharing

---

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

### Installation using `godemon-installer.sh`

1. Download the `installer.sh` from the newest release
2. Use command in terminal - `sudo chmod 777 ./godemon-installer.sh`
3. After this use command - `sudo ./godemon-installer.sh` and choose 1st option


---
### Installation from source code

1. Use command - `wget https://github.com/nProgrammer/godemon/archive/XX.XX.zip`
2. Unzip the .zip directory
3. Go to directory with source code
4. Use command - `go build`
5. After this use command - `sudo chmod 777 ./godemon; mkdir ~/.godemon; mkdir ~/.godemon/bin; mv ./godemon ~/.godemon/bin`
6. Now add `export PATH=$PATH:~/.godemon/bin` to file `.bashrc`

---
### Updating the `godemon`

1. Download the `godemon-installer.sh` from the newest release
2. Use command in terminal - `sudo chmod 777 ./godemon-installer.sh`
3. After this use command - `sudo ./godemon-installer.sh` and choose 2nd option

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
   "dev-vars": [
      {
         "key": "PORT",
         "value": "8800"
      }
   ],
   "commands": [
      {
         "name": "run",
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