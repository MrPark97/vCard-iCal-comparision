package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/emersion/go-vcard"
	emersion "github.com/emersion/go-vcard"
	mapaiva "github.com/mapaiva/vcard-go"
)

func testEmersionCard(filePath string) time.Duration {
	start := time.Now()
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dec := emersion.NewDecoder(f)
	for {
		card, err := dec.Decode()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		log.Println(card.PreferredValue(vcard.FieldFormattedName))
	}

	elapsed := time.Since(start)
	log.Println("Emersion parse vcard by: ", elapsed)
	return elapsed
}

func testMapaivaCard(filePath string) time.Duration {
	start := time.Now()
	cards, err := mapaiva.GetVCards(filePath)

	if err != nil {
		log.Fatal(err)
	}

	for _, card := range cards {
		log.Println(card.FormattedName)
	}

	elapsed := time.Since(start)
	log.Println("Mapaiva parse vcard by: ", elapsed)
	return elapsed
}
