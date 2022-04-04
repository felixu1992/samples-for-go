package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

const IssuesURL1 = "https://api.github.com/search/issues"

type IssuesSearchResult1 struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue1
}

type Issue1 struct {
	Number    int
	HTMLRUL1  string `json:"html_url"`
	Title     string
	State     string
	User      *User1
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User1 struct {
	Login    string
	HTMLURL1 string `json:"html_url"`
}

func SearchIssues1(terms []string) (*IssuesSearchResult1, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL1 + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult1
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL1}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL1}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL1}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func main() {
	result, err := SearchIssues1(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
