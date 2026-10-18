package main

import (
	"fmt"
	"os"
)

type NepaliBhasiKoLagiAankaEkThichnuHos struct{}

func (n *NepaliBhasiKoLagiAankaEkThichnuHos) StartingMessage() string {
	return "Namaste ! K samosaa parro holla hou!!"
}

type ForEnglishLanguagePressTwo struct{}

func (e *ForEnglishLanguagePressTwo) StartingMessage() string {
	return "Hellow !, How cum I help you ??"
}

type BhasaSelectors interface {
	StartingMessage() string
}

func Speek(L BhasaSelectors) {
	fmt.Println(L.StartingMessage())
}

func main() {
	switch os.Args[1] {
	case "1":
		n := &NepaliBhasiKoLagiAankaEkThichnuHos{}
		Speek(n)
	case "2":
		e := &ForEnglishLanguagePressTwo{}
		Speek(e)
	}
}
