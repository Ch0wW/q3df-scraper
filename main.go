package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	latest_map_downloaded = ""
	folderpath            = ""
	maxdownloads          = 0

	downloadamount = 0
)

func main() {

	// Get the arguments
	flag.StringVar(&folderpath, "output-directory", ".", "Directory to save .pk3 files")
	flag.IntVar(&maxdownloads, "max-downloads", 0, "Maximum downloads before exiting the program")
	flag.Parse()

	// Reading file
	data, err := ioutil.ReadFile("latest_map_downloaded")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Creating file...")

		_, err := os.Create("latest_map_downloaded")

		if err != nil {
			fmt.Println(err)
		}

	}

	latest_map_downloaded = string(data)

	if latest_map_downloaded != "" {
		log.Println("Latest downloaded file:", latest_map_downloaded)
	}

	CheckRSSInfo()
}

func UpdateFile(name string) {

	f, err := os.Create("latest_map_downloaded")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(name)

	if err != nil {
		log.Fatal("Unable to save!!")
	}

	w.Flush()

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `defrag-server-scraper`) // Set useragent to defrag-server-scraper... (just from https://github.com/q3defrag/defrag-server/blob/master/scraper.py)

	// Get the data
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fullpath := fmt.Sprintf("%s/%s", folderpath, filepath)

	// Make sure we have a valid path
	if _, err := os.Stat(folderpath); os.IsNotExist(err) {
		err := os.Mkdir(folderpath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Create the file
	out, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
