package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	filename = "chloe___isabel_20150924_0033.csv"
)

func main() {
	fi, err := os.Open("./" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	r := csv.NewReader(fi)
	// CSV contains variable number of fields per record
	r.FieldsPerRecord = -1
	head, err := r.Read()
	if err == io.EOF {
		log.Fatal("reached EOF before reading header")
	}
	if err != nil {
		log.Fatal(err)
	}

	stories := []map[string]interface{}{}
	for {
		// TODO: ReadAll then create an array of stories of specified length
		// Should be more efficient than using append
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Build story
		story := map[string]interface{}{}
		for i := 0; i < len(record); i++ {
			if val, ok := story[head[i]]; ok && record[i] != "" {
				switch t := val.(type) {
				case string:
					story[head[i]] = append([]string{record[i]}, t)
				case []string:
					story[head[i]] = append([]string{record[i]}, t...)
				default:
					log.Fatalf("unexpected type %T for %v\n", t, val)
				}
			} else {
				story[head[i]] = record[i]
			}
		}
		stories = append(stories, story)
	}

	fmt.Println("Pivotal Tracker To GitHub!")
	fmt.Println("OMGWTFBBQ YOU HAZ", len(stories), "STORIES")
}
