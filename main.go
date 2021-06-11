package main

import (
	"crypto/md5"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const DefaultConcurrentLimit = 10

type Arguments struct {
	urls     []string
	parallel int
}

func main() {
	parallel := flag.Int("parallel", DefaultConcurrentLimit, "limit of parallel requests")
	flag.Parse()

	args := flag.Args()

	a, err := parseArguments(args, *parallel)
	if err != nil {
		fmt.Println(err)
		return
	}
	startWorkers(&a)

}

func parseArguments(args []string, parallel int) (Arguments, error) {
	if len(args) < 2 {
		return Arguments{}, errors.New("A url argument is required")
	}
	if parallel < 1 {
		return Arguments{}, errors.New("The parallel flag cannot be less than 1")
	}
	var parsedUrls []string
	for _, u := range args {
		uu, err := parseRawURL(u)
		if err != nil {
			fmt.Println(err)
			return Arguments{}, err
		}
		parsedUrls = append(parsedUrls, uu.String())

	}
	return Arguments{
		urls:     parsedUrls,
		parallel: parallel,
	}, nil
}

func parseRawURL(rawurl string) (*url.URL, error) {
	var uu *url.URL
	var err error
	uu, err = url.ParseRequestURI(rawurl)
	if err != nil || uu.Scheme == "" {
		uu, err = url.ParseRequestURI("http://" + rawurl)
		if err != nil {
			fmt.Println(err)
			return &url.URL{}, err
		}
	}
	if !strings.Contains(uu.Host, ".") {
		return &url.URL{}, errors.New("no host in URL " + rawurl)
	}

	return uu, nil
}

func startWorkers(a *Arguments) {
	workers := make(chan struct{}, a.parallel)

	for i := 0; i < a.parallel; i++ {
		workers <- struct{}{}
	}

	var wg sync.WaitGroup
	for _, url := range a.urls {
		<-workers
		wg.Add(1)
		go func(url string) {
			md5s, err := getMD5(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(url, md5s)
			wg.Done()
			workers <- struct{}{}
		}(url)

	}
	wg.Wait()

}

func getMD5(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(b)), nil
}
