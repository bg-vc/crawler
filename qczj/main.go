package main

import (
	"crawler/qczj/pkg"
	"crawler/qczj/service"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"sync"
	"time"
)

var (
	StartUrl = "/2sc/%s/a0_0msdgscncgpi1ltocsp1exb4/"
	BaseUrl  = "https://car.autohome.com.cn"
	maxPage  = 99
	urls     = make([]string, 0)
	rw       sync.RWMutex
	wg       sync.WaitGroup
)

func PushUrl(url string) {
	rw.RLock()
	defer rw.RUnlock()
	log.Printf("PushUrl:%v\n", url)
	urls = append(urls, url)
}

func PopUrl() string {
	rw.Lock()
	defer rw.Unlock()
	length := len(urls)
	if length < 1 {
		return ""
	}
	url := urls[length-1]
	urls = urls[:length-1]
	return url
}

func main() {
	service.InitDB()
	Start()
}

func Start() {
	startTime := time.Now()
	cities := service.GetCities()
	log.Printf("cities num:%v\n", len(cities))
	for _, v := range cities {
		url := fmt.Sprintf(BaseUrl+StartUrl, v.Pinyin)
		PushUrl(url)
	}
	log.Printf("urls num:%v\n", len(urls))

	for _, url := range urls {
		wg.Add(1)
		go handler(url)
	}
	wg.Wait()
	fmt.Printf("coustTime:%v\n", time.Since(startTime))

}

func handler(url string) {
	defer wg.Done()
	for {
		body, err := pkg.Get(url)
		if err != nil {
			log.Printf("pkg.Get error:%v\n", err)
			return
		}
		doc, err := goquery.NewDocumentFromReader(body)
		if err != nil {
			log.Printf("goquery.NewDocumentFromReader error:%v\n", err)
			return
		}
		currentPage := service.GetCurrentPage(doc)
		nextPageUrl, _ := service.GetNextPageUrl(doc)

		if currentPage > 0 && currentPage <= maxPage {
			cars := service.GetCars(doc)
			service.BatchAddCar(cars)
		} else {
			log.Println("max page")
			return
		}
		if nextPageUrl != "" {
			url = BaseUrl + nextPageUrl
		} else {
			return
		}
	}
}
