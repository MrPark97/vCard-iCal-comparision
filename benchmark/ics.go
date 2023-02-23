package main

import (
	"io"
	"log"
	"os"
	"time"

	arran4 "github.com/arran4/golang-ical"
	emersion "github.com/emersion/go-ical"
)

func testEmersionCalendar(filePath string) time.Duration {
	start := time.Now()
	var calendarReader io.Reader
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	calendarReader = f

	dec := emersion.NewDecoder(calendarReader)
	for {
		calendar, err := dec.Decode()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for _, event := range calendar.Events() {
			summary, err := event.Props.Text(emersion.PropSummary)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Found event: %v", summary)
		}
	}

	elapsed := time.Since(start)
	log.Println("Emersion parse calendar by: ", elapsed)
	return elapsed
}

func testArran4Calendar(filePath string) time.Duration {
	start := time.Now()
	calFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Cannot read the file: ", err)
	}

	cal, err := arran4.ParseCalendar(calFile)
	if err != nil {
		log.Fatal("Cannot parse calendar: ", err)
	}

	for _, event := range cal.Events() {
		log.Println(event.GetProperty("SUMMARY").Value)
	}

	elapsed := time.Since(start)
	log.Println("Arran4 parse calendar by: ", elapsed)
	return elapsed
}
