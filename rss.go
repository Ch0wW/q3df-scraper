package main

import (
	"fmt"

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

			fmt.Println("Downloaded: " + fileUrl)
		}

		// Set the latest file in the download
		latest_map_downloaded = feed.Items[0].Title
		UpdateFile(latest_map_downloaded)
	}
}
