package main

import "testing"

func TestMakeQuestion(t *testing.T) {
    json := `[{ "score": 1 }]`
    question := makeQuestions(json)

    if question == nil {
        t.Error("The parsed question is nil")
    }
}
