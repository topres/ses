package main

import (
    //"bufio"
    "fmt"
    //"os"
    "flag"
)

func start() {

    queryPtr := flag.String("q", "", "Query string to search with")

    flag.Parse()

    if *queryPtr == "" {
        fmt.Println("You need to provide a query string.")
        return
    }

    result := Search(
        *queryPtr,
        []string{},
        []string{})

    if len(result.questions) > 0 {

        for _, question := range result.questions {
            fmt.Println(question.title)
            fmt.Println(question.link)
            fmt.Println("")
        }

    } else {
        fmt.Println("no results")
    }
}

func main(){
    start()
}

