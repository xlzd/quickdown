package quickdown

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/xlzd/quickdown"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\t%s [url] [-wc workerCount] [-f filename]\n", os.Args[0], path.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	workerCount := flag.Int("wc", 5, "(optional) worker count")
	filename := flag.String("f", "", "(optional) filename to save.")

	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Println("\033[0;31m[goquick]Argument error: need url to download.\033[0m")
		flag.Usage()
		os.Exit(1)
	}

	quickdown.DownloadWithWorkersTo(args[0], *workerCount, *filename)
}
