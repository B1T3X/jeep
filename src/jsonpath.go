package main

import (
	"github.com/PaesslerAG/jsonpath"
	"fmt"
	"encoding/json"
	)

type jiraIssue struct {
	Key string
	Description string
	Type string
	Reporter string
}

func parseIssue(issueJson []byte) (issue jiraIssue, err error) {
	var j interface{}
	json.Unmarshal(issueJson, &j)
	var issueKey interface{}
	var issueDescription interface{}
	var issueType interface{}
	var issueReporter interface{}

	// To simplify error handling
	if issueKey, err = jsonpath.Get("$.issue.key", j); err != nil {
		panic(err)
	}
	if issueDescription, err = jsonpath.Get("$.issue.fields.summary", j); err != nil {
		panic(err)
	}
	if issueType, err = jsonpath.Get("$.issue.fields.issuetype.name", j); err != nil {
		panic(err)
	}
	if issueReporter, err = jsonpath.Get("$.issue.fields.reporter.displayName", j); err != nil {
		panic(err)
	}

	issue = jiraIssue{
		Key: fmt.Sprintf("%v", issueKey),
		Description: fmt.Sprintf("%v", issueDescription),
		Type: fmt.Sprintf("%v", issueType),
		Reporter: fmt.Sprintf("%v", issueReporter),
	}
	return
}

