// issuesrport
package main

import (
	"dang/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

const temp1 = `{{.TotalCount}} issues:
{{range .Items}}-------------------------------------------------------------
Numbers: {{.Number}}
User:    {{.User.Login}}
Title:   {{.Title | printf "%.64s"}}
Age:     {{.CreateAt | daysAgo }} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp1))

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
