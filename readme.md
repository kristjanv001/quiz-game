# Quiz Game

![](https://res.cloudinary.com/du1qfmeoz/image/upload/v1711011697/Various/screely-1711011431015_aupdow.png)

A CLI app written in Go. It reads a given csv file containing quiz questions and answers and presents these questions to the user. In the end the program outputs how many questions the user got right and how many questions there were in total.

There's a time limit. If the user does not finish the quiz in time, the program will terminate, letting the user know his/her result.


## Build and run

1. `go build .`
2. `./quizgame`


## Flags

The program accepts the following flags:
1. `-csv string` to specify the csv file containing the problems
2. `-limit int` to set a time limit (in seconds) for the quiz

## Sample usage
- `./quizgame -csv=problems.csv -limit=40`
- `./quizgame`
