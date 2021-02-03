# Godemon
### Now, using Godemon you can use live reload on your apps.

### How to use it?
First of all you need to build your version:
`cd ./godemon/`
`go build`

Now you can move file `godemon` to folder with your go project and 
start it using command: `./godemon <fileOrDir> <fileOrModule>`.
In first argument you're giving godemon the path to project, and in second
argument you're choosing is it simple go file or go module - if module pass
argument - `mod` , if file pass argument - `file`