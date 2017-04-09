package main

import "encoding/json"

func makeQuestions(jsonInput string) *[]Question {

    var questions []Question

    err := json.Unmarshal([]byte(jsonInput), &questions)

    if err != nil {
        panic(err)
    }

    return &questions
}
