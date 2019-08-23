package pkg

import (
	"github.com/axgle/mahonia"
	"io"
	"log"
	"net/http"
)

func Get(url string) (io.Reader, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("http.NewRequest error:%v\n", err)
		return nil, err
	}
	req.Header.Add("User-Agent", GetUserAgent())
	req.Header.Add("Referer", "https://car.autohome.com.cn")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do error:%v\n", err)
		return nil, err
	}
	mash := mahonia.NewDecoder("gbk")
	return mash.NewReader(resp.Body), nil
}