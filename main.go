package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var daysF, weeksF bool

func init() {
	flag.BoolVar(&daysF, "days", false, "Unit: days (24 h)")
	flag.BoolVar(&weeksF, "weeks", false, "Unit: weeks (168 h)")
}

func main() {
	flag.Parse()

	if (! daysF && ! weeksF) || ( daysF && weeksF) {
		fmt.Println("Please specify a single unit arg")
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			if err := calculate(scanner.Text()); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		for _, input := range flag.Args() {
			if err := calculate(input); err != nil {
				log.Fatal(err)
			}
		}
	}

}

func calculate(inputStr string) (err error) {
	input, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		return
	}

	var duration string
	if daysF {
		duration = strconv.FormatFloat(input * 24.0, 'f', -1, 64) + "h"
	}

	if weeksF {
		duration = strconv.FormatFloat(input * 24.0 * 7.0, 'f', -1, 64) + "h"
	}

	answer, err := time.ParseDuration(duration)
	if err != nil {
		return
	}

	fmt.Println(time.Now().Add(answer * -1).UTC().Format("2006-01-02 T15:04 -0700"))

	return nil
}
