# go-do
TODO List console application to manage and export hierarchies of tasks 📝

## Installation
* Clone or download repository
* Go to go-do/src
* In console run the following: go build -o go-do.exe
* Run the executable you created

## Getting started
List of arguments you can use:
* add $task-name$ – creates task depending on specified/unspecified parent
* remove $task-position$ – removes task at position depending on specified/unspecified parent
* list $task-name$ – displays task, type root or all to display all tasks
* composite y/n – use with add argument; defines whether created task to be composite or not
* parent $task-name$ – use with add/remove arguments; specifies parent for add and remove commands
* export $file-name$ – creates .txt file on Desktop with exported root task structure

## Requirements
* Windows OS

## Built With
Project was built purely on Go native libraries without using third-party extensions

## Contributing
Contribution is always appreciated, so feel free to do so using Git Flow 