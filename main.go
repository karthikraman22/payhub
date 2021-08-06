package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"achuala.in/payhub/pbgen/pacs008"
	"achuala.in/payhub/pbgen/pain001"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	flag "github.com/spf13/pflag"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func interruptHandler() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	os.Exit(0)
}

var k = koanf.New(".")

func main() {
	pain := &pain001.CustomerCreditTransferInitiationRq{}
	fmt.Printf("pain: %v\n", pain)
	pacs := &pacs008.FI2FICustomerCreditTransferInitiationRq{}
	fmt.Printf("pacs: %v\n", pacs)
	// Use the POSIX compliant pflag lib instead of Go's flag lib.
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	// Path to one or more config files to load into koanf along with some config params.
	f.StringSlice("conf", []string{"mock/mock.toml"}, "path to one or more .toml config files")
	f.String("time", "2020-01-01", "a time string")
	f.String("type", "xxx", "type of the app")
	f.Parse(os.Args[1:])

	// Load the config files provided in the commandline.
	cFiles, _ := f.GetStringSlice("conf")
	for _, c := range cFiles {
		if err := k.Load(file.Provider(c), toml.Parser()); err != nil {
			log.Fatalf("error loading file: %v", err)
		}
	}

	// "time" and "type" may have been loaded from the config file, but
	// they can still be overridden with the values from the command line.
	// The bundled posflag.Provider takes a flagset from the spf13/pflag lib.
	// Passing the Koanf instance to posflag helps it deal with default command
	// line flag values that are not present in conf maps from previously loaded
	// providers.
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Println("time is = ", k.String("time"))
}
