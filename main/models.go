package main

type SearchResult struct {
    questions []Question
}

type Question struct {
    title string
    link string
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
