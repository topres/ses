package main

type SearchResult struct {
    questions []Question
}

type Question struct {
    score int
    title string
    excerpt string
    body string
    tags []string
    link string
    answers []Answer
    comments []Comment
}

type Comment struct {
    score int
    message string
}

type Answer struct {
    score int
    message string
    comments []Comment
}
