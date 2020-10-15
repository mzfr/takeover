package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
)

// ProviderData Structure for each provider stored in providers.json file
type ProviderData struct {
	Name     string   `json:"name"`
	Cname    []string `json:"cname"`
	Response []string `json:"response"`
}

// Providers structure
var Providers []ProviderData

// Targets structure
var Targets []string

var (
	hostsList  string
	threads    int
	verbose    bool
	forceHTTPS bool
	timeout    int
	outputFile string
	directory  string
	provider   string
)

// initializeProviders takes the json files of providers.
// Accepting the path externally would be nice.s
func initializeProviders(providerpath string) {
	raw, err := ioutil.ReadFile(providerpath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(raw, &Providers)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}

// Read host files and subdomains
func readFile(file string) (lines []string, err error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return lines, err
	}

	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Get is used to send a request to the server
func Get(url string, timeout int, https bool) (resp gorequest.Response, body string, errs []error) {
	if https == true {
		url = fmt.Sprintf("https://%s/", url)
	} else {
		url = fmt.Sprintf("http://%s/", url)
	}

	resp, body, errs = gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Timeout(time.Duration(timeout)*time.Second).Get(url).
		Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0").
		End()

	return resp, body, errs
}

func parseArguments() {
	flag.IntVar(&threads, "t", 20, "Number of threads to use")
	flag.StringVar(&hostsList, "l", "", "List of hosts to check takeovers on")
	flag.BoolVar(&verbose, "v", false, "Show verbose output")
	flag.BoolVar(&forceHTTPS, "https", false, "Force HTTPS connections")
	flag.IntVar(&timeout, "timeout", 10, "Seconds to wait before timeout")
	flag.StringVar(&outputFile, "o", "", "File to write enumeration output to")
	flag.StringVar(&directory, "d", "", "directory having files of domains")
	flag.StringVar(&provider, "p", "", "Path of the providers file")

	flag.Parse()
}

// check if the CNAME is present in the providers.json
func cnameExists(key string) bool {
	for _, provider := range Providers {
		for _, cname := range provider.Cname {
			if strings.Contains(key, cname) {
				return true
			}
		}
	}
	return false
}

// check if the Takeover is possible
func check(target string, TargetCNAME string) {
	_, body, errs := Get(target, timeout, forceHTTPS)
	if len(errs) <= 0 {
		{
			for _, provider := range Providers {
				for _, cname := range provider.Cname {
					if strings.Contains(TargetCNAME, cname) {
						for _, response := range provider.Response {
							if strings.Contains(body, response) == true {
								if provider.Name == "cloudfront" {
									_, body2, _ := Get(target, 120, true)
									if strings.Contains(body2, response) == true {
										fmt.Printf("\n[\033[31;1;4m%s\033[0m] Takeover Possible At : %s", provider.Name, target)
									}
								} else {
									fmt.Printf("\n[\033[31;1;4m%s\033[0m] Takeover Possible At %s with CNAME %s", provider.Name, target, TargetCNAME)
								}
							}
							return
						}
					}
				}
			}
		}
	} else {
		if verbose == true {
			log.Printf("[ERROR] Get: %s => %v", target, errs)
		}
	}

	return
}

// checker is a controller made for check function
func checker(target string) {
	TargetCNAME, err := net.LookupCNAME(target)
	if err != nil {
		log.Printf("Error")
		os.Exit(1)
	} else {
		if cnameExists(TargetCNAME) == true {
			if verbose == true {
				fmt.Printf("[SELECTED] %s => %s", target, TargetCNAME)
			}
			check(target, TargetCNAME)
		}
	}
}

func startLooking(hostsList string) {
	Hosts, err := readFile(hostsList)
	if err != nil {
		fmt.Printf("\nread: %s\n", err)
		os.Exit(1)
	}

	Targets = append(Targets, Hosts...)

	hosts := make(chan string, threads)
	processGroup := new(sync.WaitGroup)
	processGroup.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			for {
				host := <-hosts
				if host == "" {
					break
				}
				checker(host)
			}
			processGroup.Done()
		}()
	}

	for _, Host := range Targets {
		hosts <- Host
	}

	close(hosts)
	processGroup.Wait()

}

func main() {
	parseArguments()

	if hostsList == "" && directory == "" {
		fmt.Printf("SubOver: No hosts list or directory specified for testing!")
		fmt.Printf("\nUse -h for usage options\n")
		os.Exit(1)
	}

	if provider != "" && fileExists(provider) {
		initializeProviders(provider)
	} else if fileExists("providers.json") {
		initializeProviders("providers.json")
	} else {
		fmt.Println("Can't find the Providers.json")
		os.Exit(1)
	}

	if directory != "" {
		fmt.Println("--> Got a directory of hostlists!!")
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			fmt.Println("Could read the directory")
		}

		for _, f := range files {
			filename := fmt.Sprintf("%s/%s", directory, f.Name())
			startLooking(filename)
		}
	}
	if hostsList != "" {
		fmt.Println("--> Got Single hostlist!!")
		startLooking(hostsList)
	}

}
