package main

import (
	"github.com/PaesslerAG/jsonpath"
	"os"
	"fmt"
	"encoding/json"
	//"reflect"
	)

type issue struct {
	Key string
	Description string
	Type string
	Reporter string
}

func main() {
	var j interface{}
	data, err := os.ReadFile("/tmp/json.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &j)
	//var issueKey string
	//var issueDescription string
	//var issueType string
	//var issueReporter string
	issueKey, err := jsonpath.Get("$.issue.key", j)
	issueDescription, err := jsonpath.Get("$.issue.fields.summary", j)
	issueType, err := jsonpath.Get("$.issue.fields.issuetype.name", j)
	issueReporter, err := jsonpath.Get("$.issue.fields.reporter.displayName", j)
	fmt.Println(issueKey)
	fmt.Println(issueDescription)
	fmt.Println(issueType)
	fmt.Println(issueReporter)
	myIssue := issue{
		Key: fmt.Sprintf("%v", issueKey),
		Description: fmt.Sprintf("%v", issueDescription),
		Type: fmt.Sprintf("%v", issueType),
		Reporter: fmt.Sprintf("%v", issueReporter),
	}

	if err != nil {
		panic(err)
	}

	fmt.Println(myIssue.Key)

}
