package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	daysF, weeksF, versionF bool
	version string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of deployed: [VALUE] [UNIT days weeks]\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&versionF, "v", false, "Display version information")
}

func main() {
	flag.Parse()

	if versionF {
		fmt.Println(version)
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			if err := calculate(scanner.Text()); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		if err := calculate(flag.Arg(0) + " " + flag.Arg(1)); err != nil {
			log.Fatal(err)
		}
	}

}

func calculate(inputStr string) (err error) {
	input, err := parseInput(inputStr)
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

func parseInput(input string) (value float64, err error) {
	comp := strings.Split(input, " ")
	if len(comp) == 0 {
		return 0.0, fmt.Errorf("no input\n")
	}

	value, err = strconv.ParseFloat(comp[0], 64)
	if err != nil {
		return
	}

	switch comp[1] {
	case "day", "days":
		daysF = true
	case "week", "weeks":
		weeksF = true
	}

	return
}
