package main

import "fmt"

func start() {
    result := Search(
        "query",
        []string{},
        []string{})

    fmt.Println(result.questions[0].title)

}

func main(){
    start()
}

