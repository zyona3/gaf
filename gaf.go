package gaf

import (
	"fmt"
	"strings"
	"os"
	"log"
	"io"
	"encoding/csv"
	"strconv"
	"errors"
	"flag"
)

var (
	chunks = 8
)

type item struct {
	title string
	amount   int
}

func divmod(count int, increment int) (int, int) {
	tmp := int(count * chunks / increment)
	return tmp / 8, tmp % chunks
}

func leftjust(s string, n int, fill string) string {
	return strings.Repeat(fill, n-len(s)) + s
}

func validateinput(v int) error {
	if (v < 0) {
		return errors.New("(*>△<)<valueは0以上")
	}
	return nil
}

func handleArgs() *string {
	fn := flag.String("filename", "", "File whose each row contains (title, value)")
	flag.Parse()
	return fn
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Run(argv []string, outStream, errStream io.Writer) error {
	fs := flag.NewFlagSet(
		fmt.Sprintf("gaf (v%s rev:%s", version, revision), flag.ContinueOnError)
	fs.SetOutput(errStream)

	v := fs.Bool("version", false, "display version")
	if err := fs.Parse(argv); err != nil {
		return err
	}

	if *v {
		return printVersion(outStream)
	}

	if len(fs.Args()) < 1 {
		return fmt.Errorf("file required")
	}

	filename := fs.Args()[0]

	return run(filename, outStream, errStream)
}

func printVersion(out io.Writer) error {
	_, err := fmt.Fprintf(out, "tt v%s (rev:%s)\n", version, revision)
	return err
}

func run(filename string, outStream, errStream io.Writer) error {
	f, err := os.Open(filename)
	checkError(err)
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	checkError(err)
	var data []item
	for _, record := range records {
		Val, err := strconv.Atoi(record[1])
		checkError(err)
		if err := validateinput(Val); err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		data = append(data, item{title: record[0], amount: Val})
	}
	max_value := 0
	longest_label_length := 0
	for _, v := range data {
		if max_value < v.amount {
			max_value = v.amount
		}
		if longest_label_length< len(v.title) {
			longest_label_length = len(v.title)
		}
	}
	increment := max_value / 25

	for _, v := range data {
		bar_chunks, remainder := divmod(v.amount, increment)
		bar := strings.Repeat("█", bar_chunks)

		if remainder > 0 {
			bar = bar + string(rune(int('█') + (chunks - remainder)))
		}

		if bar == "" {
			bar = "▏"
		}
		out := leftjust(v.title, longest_label_length, " ")
		fmt.Printf("%s :  %s\n", out, bar)
	}
	return nil
}