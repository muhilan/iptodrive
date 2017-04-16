package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	drive "google.golang.org/api/drive/v3"
)

type Fetcher struct {
	url string
}

func generateIPFile(srv *drive.Service) {

	f := Fetcher{url: "http://ipecho.net/plain"}
	contents, err := f.doGet()
	if err != nil {
		fmt.Println(err)
		return
	}
	// file.ModifiedTime = time.Now().String()
	filename := "./ip.txt"
	ioutil.WriteFile(filename, contents, 0777)

	ipFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening %q: %v", filename, err)
	}

	_, err = srv.Files.Create(&drive.File{Name: "ip.txt"}).Media(ipFile).Do()
	if err != nil {
		fmt.Println(err)
	}

}

func (f *Fetcher) doGet() ([]byte, error) {
	client := &http.Client{}

	response, err := client.Get(f.url)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Unable to ping url")
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("   ", string(contents))
		return contents, nil
	}
}

//
// func driveMain(client *http.Client, argv []string) {
// 	if len(argv) != 1 {
// 		fmt.Fprintln(os.Stderr, "Usage: drive filename (to upload a file)")
// 		return
// 	}
//
// 	service, err := drive.New(client)
// 	if err != nil {
// 		log.Fatalf("Unable to create Drive service: %v", err)
// 	}
//
// 	filename := argv[0]
//
// 	goFile, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatalf("error opening %q: %v", filename, err)
// 	}
// 	driveFile, err := service.Files.Insert(&drive.File{Title: filename}).Media(goFile).Do()
// 	log.Printf("Got drive.File, err: %#v, %v", driveFile, err)
// }
