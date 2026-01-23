package main

import (
	"dictionary/dict"
	"dictionary/entry"
	"fmt"
)

func main() {
	e1 := entry.New("Haus", "House")
	e2 := entry.New("Fahrrad", "bicycle")

	d := dict.New()
	d.Add(e1)
	d.Add(e2)

	fmt.Print("Bitte ein deutsches Wort eingeben: ")
	de := ""
	fmt.Scanln(&de)

	en := d.Lookup(de)
	if en != "" {
		fmt.Printf("Das englische Wort ist: %s\n", en)
	} else {
		fmt.Println("Nicht gefunden")
	}
}
