package main

import (
    "net/url"
    "net/http"
    "io/ioutil"
    "fmt"
    "errors"
    "strings"
)

var apiUri string = "http://api.stackexchange.com/2.2"

func Search (query string, tags []string, sites []string) SearchResult {

    urlEncodedQuery := url.PathEscape(query)

    uri := apiUri +
    "/search?order=desc&sort=relevance&site=stackoverflow&intitle=" +
    urlEncodedQuery +
    "&tagged=" +
    strings.Join(tags, ";") +
    "&page=1&pagesize=10"

    fmt.Println(uri)

    body := httpGet(uri);

    return SearchResult {
        questions: makeQuestions(body)}
}

func httpGet (uri string) string {
    resp, err := http.Get(uri)

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        panic(err)
    }

    stringResponse := string(body[:])

    if resp.StatusCode > 200 {
        fmt.Println(uri)
        panic(errors.New(stringResponse))
    }

    return stringResponse
}
