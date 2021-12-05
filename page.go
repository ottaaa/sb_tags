package main

import "strings"

type Page struct {
	Lines []string `json:"lines"`
}

func (p Page) putTag() []string {
	var inCode = false
	var pageTags []string
	for _, line := range p.Lines {
		inCode = isCode(line, inCode)
		if inCode == true {
			continue
		}
		pageTags = append(pageTags, lineTags(line)...)
	}
	return pageTags
}

func lineTags(line string) []string {
	var words = strings.Fields(line)
	var tags []string
	for _, word := range words {
		// タグ
		if strings.HasPrefix(word, "#") && len(word) != 1 {
			tags = append(tags, word)
		}
	}
	return tags
}

func isCode(line string, inCode bool) bool {
	if inCode && strings.HasPrefix(line, " ") {
		return true
	}
	if strings.HasPrefix(line, "code:") {
		return true
	}
	return false
}
