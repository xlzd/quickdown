package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/xlzd/quickdown"
)

func getVideoDownloadUrl(viewkey string, quality int) (url string) {
	resp, err := http.Get("https://www.pornhub.com/view_video.php?viewkey=" + viewkey)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	contentBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	content := string(contentBytes)

	pattern, _ := regexp.Compile("\"quality\":\"(\\d+)\",\"videoUrl\":\"(http.+?)\"")
	for _, item := range pattern.FindAllStringSubmatch(content, -1) {
		qStr, url := item[1], item[2]
		url = strings.Replace(url, "\\", "", -1)
		q, _ := strconv.Atoi(qStr)
		if q == quality {
			return url
		}
	}
	return ""
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\t%s [viewkey] [-q quality]\n", os.Args[0], path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	quality := flag.Int("q", 480, "quality")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Println("\033[0;31m[goquick]Argument error: need *ornhub viewkey download.\033[0m")
		flag.Usage()
		os.Exit(1)
	}

	if !strings.HasPrefix(args[0], "ph") {
		log.Printf("\033[0;31m[goquick]Argument error: got invalid *ornhub viewkey: %s.\033[0m", args[0])
		os.Exit(1)
	}

	url := getVideoDownloadUrl(args[0], *quality)

	if url == "" {
		log.Printf("\033[0;31m[gornhub] Get *ornhub viewkey=%s %dP video download url faild.\033[0m", args[0], quality)
	}

	log.Printf("\033[0;32mDownloading %s %dP video.\033[0m", args[0], quality)
	quickdown.DownloadTo(url, fmt.Sprintf("%s.mp4", args[0]))
}
