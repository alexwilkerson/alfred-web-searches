package main

import (
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strings"
)

// parceCSV parses CSV for links and arguments
func parseCSV() map[string]string {
	var err error

	// Load file
	f, err := os.Open("websites.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	links := make(map[string]string)

	for _, record := range records {
		links[record[0]] = record[1]
	}

	return links
}

// doSearch searches through the websites and returns results to Alfred
func doSearch() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	links := parseCSV()

	re1 := regexp.MustCompile(`.: `)
	re2 := regexp.MustCompile(`(all)`)

	for key, value := range links {
		if strings.Contains(key, "r: ") {
			wf.NewItem(key).Valid(true).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(redditIcon).Var("RECENT", re2.ReplaceAllString(value, `week`)).Subtitle("⌘ = Search past week")
		} else if strings.Contains(key, "d: ") {
			wf.NewItem(key).Valid(true).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(docIcon)
		} else if strings.Contains(key, "g: ") {
			wf.NewItem(key).Valid(true).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(githubIcon)
		} else if strings.Contains(key, "s: ") {
			wf.NewItem(key).Valid(true).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(stackIcon)
		} else if strings.Contains(key, "f: ") {
			wf.NewItem(key).Valid(true).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(forumsIcon)
		} else if strings.Contains(key, "t: ") {
			wf.NewItem(key).Valid(true).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key).Icon(translateIcon)
		} else {
			wf.NewItem(key).Valid(true).Var("ARG", re1.ReplaceAllString(key, ``)).UID(key)
		}
	}

	query := os.Args[1]

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()

	return nil
}
