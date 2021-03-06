package main

import (
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func CheckRSSInfo() {

	amount := 0

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://en.ws.q3df.org/rss/released_pk3_files/")

	for i := 0; i < len(feed.Items); i++ {

		// Do we have at least downloaded a file before?
		if latest_map_downloaded != "" {
			if feed.Items[amount].Title == latest_map_downloaded {
				break
			}
		}

		// Increase the amount of downloaded files...
		amount += 1
	}

	if feed.Items[0].Title != latest_map_downloaded {

		for i := 0; i < amount; i++ {
			filename := feed.Items[i].Title
			fileUrl := feed.Items[i].Link

			err := DownloadFile(filename, fileUrl)
			if err != nil {
				panic(err)
			}

			log.Println("Downloaded: " + fileUrl)

			if maxdownloads > 0 {
				downloadamount += 1
			}
		}

		// Set the latest file in the download
		latest_map_downloaded = feed.Items[0].Title
		UpdateFile(latest_map_downloaded)

		// Verify if we exceeded our max downloads...
		if maxdownloads > 0 && downloadamount < maxdownloads {
			log.Println("Exceeded the maximum amount of downloads allowed! Quitting...")
			os.Exit(0)
		}
	} else {
		log.Println("No new map found.")
	}
}
