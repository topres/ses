package main

import (
    "encoding/json"
    "fmt"
)

func makeQuestions(jsonInput string) []Question {
    type Request struct {
        Items []struct {
            Title string `json:"title"`
            Link string `json:"link"`
        } `json:"items"`
        HasMore bool `json:"has_more"`
    }

    var request Request

    err := json.Unmarshal([]byte(jsonInput), &request)

    if err != nil {
        fmt.Println(jsonInput)
        panic(err)
    }

    questions := make([]Question, len(request.Items))

    for i, v := range request.Items {
        questions[i] = Question {
            title: v.Title,
            link: v.Link}
    }

    return questions
}
