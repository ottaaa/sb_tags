package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Project struct {
	Pages []*Page `json:"pages"`
}

func NewProject(arg string) Project {
	bytes, err := ioutil.ReadFile(arg)
	if err != nil {
		log.Fatal(err)
	}

	var project Project
	if err := json.Unmarshal(bytes, &project); err != nil {
		log.Fatal(err)
	}
	return project
}

func (p Project) tags() []string {
	var tags []string
	for _, p := range p.Pages {
		tags = append(tags, p.putTag()...)
	}

	return uniq(tags)
}

func uniq(arr []string) []string {
	m := make(map[string]struct{})
	for _, ele := range arr {
		m[ele] = struct{}{}
	}

	var uniq []string
	for i := range m {
		uniq = append(uniq, i)
	}
	return uniq
}
