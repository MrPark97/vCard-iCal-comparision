package main

import (
	"log"
	"time"
)

type benchmarkResult struct {
	PackageName string
	LastUpdated string
	Stars       int
	Performance time.Duration
}

func main() {
	// vcf tests
	emersionCard := benchmarkResult{
		PackageName: "github.com/emersion/go-vcard",
		LastUpdated: "May 7, 2022",
		Stars:       87,
	}
	emersionCard.Performance = testEmersionCard("../static/sample.vcf")

	mapaivaCard := benchmarkResult{
		PackageName: "github.com/mapaiva/vcard-go",
		LastUpdated: "May 23, 2022",
		Stars:       5,
	}
	mapaivaCard.Performance = testMapaivaCard("../static/sample.vcf")

	log.Println("vCard Performance:")
	log.Println(emersionCard)
	log.Println(mapaivaCard)

	// ics tests
	emersionCalendar := benchmarkResult{
		PackageName: "github.com/emersion/go-ical",
		LastUpdated: "Aug 3, 2022",
		Stars:       38,
	}
	emersionCalendar.Performance = testEmersionCalendar("../static/sample.ics")

	arran4Calendar := benchmarkResult{
		PackageName: "github.com/arran4/golang-ical",
		LastUpdated: "Last week",
		Stars:       170,
	}
	arran4Calendar.Performance = testArran4Calendar("../static/sample.ics")
	log.Println(emersionCalendar)
	log.Println(arran4Calendar)
}
