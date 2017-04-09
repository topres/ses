package main

func Search (query string, tags []string, sites []string) *SearchResult {

    question := Question {
        score: 5,
        title: "question title",
        excerpt: "short question excerpt",
        body: "a bit longer question body",
        tags: []string {"java", "generics", "jvm"},
        comments: []Comment{},
        answers: []Answer{} }

    questions := []Question { question }

    result := SearchResult {
        questions: questions }

    return &result;
}
