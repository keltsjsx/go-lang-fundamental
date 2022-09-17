# Go-lang Fundamental !
This repo is based on [Techworld with Nana Tutorial](https://www.youtube.com/watch?v=yyUHQIec83I) but this notes is created by myself for learning only. This readme will updated regularly as long as i learn about go-lang. You can clone this repo if you want 😁

## Initial Go-Lang ⚡
 1. Initial module with command `go mod init <app-name>`
 2. All our code must belong to **package**
 3. The first statement in Go file must be **"package"**
 4. **Main function** is a entry point for application
 5. Running your go-lang application with `go run <entrypointfile.go>`

## Standard Starting Point Go-Lang 🖊
This could be standard format of go-lang application
```go
package  main

import  "fmt"

func  main() {
	fmt.Print("Hello World")
}
```
## Variable and Const in Go-Lang 😶
 1. For declare variable, you can use `var <variable_name> = <value>`
 2. For declare constnya, you can use `const <variable_name> = <value>`
 3. For declare variable, go-lang have *Syntactic Sugar* like this `<variable_name> := <value>`
 4. If you want to really strict, you can declare with data type like this `var <variable_name> <data_type> = <value>`

## Formatted Output 🧩
We can show output with formatted style like this
```go
var name string = "Keltskaya"
fmt.Printf("Hello my name is %v", name)
```
## Getting User Input 💻
We can get input from user with combining method `scan` and pointer to variable like this
```go
var firstName string;

// ask for name
fmt.Print("Enter your first name: ")
fmt.Scan(&firstName)

// show the output
fmt.Printf("Welcome to the show, %v", firstName)
```
