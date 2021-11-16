package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

// album represents data about a record album.
type Album struct {
	ID     string `json:"id" csv:"song.id"`
	Title  string `json:"title" csv:"title"`
	Artist string `json:"artist" csv:"artist.name"`
	Year   string `json:"year" csv:"year"`
}

var albums = []Album{}

func readCsvData() {
	albumfile, err := os.OpenFile("./music.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer albumfile.Close()

	if err := gocsv.UnmarshalFile(albumfile, &albums); err != nil {
		panic(err)
	}
	// for _, a := range albums {
	// 	fmt.Println(a)
	// }
}

func writeCsvData() error{
	albumfile, err := os.OpenFile("./music.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer albumfile.Close()

	if err := gocsv.MarshalFile(&albums, albumfile); err != nil {
		return err
	}
	var e error = nil
	return e
}