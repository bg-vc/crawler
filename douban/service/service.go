package service

import (
	"crawler/douban/dao"
	"crawler/douban/model"
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	BaseUtl = "https://movie.douban.com/top250"
)

func InitDB() {
	dao.InitDB()
}

func Start() {
	movies := make([]*model.DouBanMovie, 0)
	pages := GetPages(BaseUtl)
	for _, page := range pages {
		//time.Sleep(10 * time.Millisecond)
		url := strings.Join([]string{BaseUtl, page.Url}, "")
		doc, err := goquery.NewDocument(url)
		if err != nil {
			log.Printf("goquery.NewDocumen error:%v\n", err)
		}

		movies = append(movies, ParseMovies(doc)...)
	}

	AddData(movies)
}

func AddData(movies []*model.DouBanMovie) {
	for index, movie := range movies {
		if err := dao.AddMovie(movie); err != nil {
			log.Printf("dao.AddMovie index:%v, error:%v\n", index, err)
		}
	}
}

type Page struct {
	Page int
	Url  string
}

func GetPages(url string) []*Page {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalf("GetPages error:%v\n", err)
	}

	return ParsePages(doc)
}

func ParsePages(doc *goquery.Document) []*Page {
	pages := make([]*Page, 0)
	pages = append(pages, &Page{1, ""})
	doc.Find("#content > div > div.article > div.paginator > a").Each(func(i int, s *goquery.Selection) {
		page, _ := strconv.Atoi(s.Text())
		url, _ := s.Attr("href")
		item := &Page{page, url}
		pages = append(pages, item)
	})
	return pages
}

func ParseMovies(doc *goquery.Document) []*model.DouBanMovie {
	movies := make([]*model.DouBanMovie, 0)
	doc.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".hd a span").Eq(0).Text()
		subtitle := s.Find(".hd a span").Eq(1).Text()
		subtitle = strings.TrimLeft(subtitle, " / ")
		other := s.Find(".hd a span").Eq(2).Text()
		other = strings.TrimLeft(subtitle, " / ")

		descInfo := s.Find(".bd p").Eq(0).Text()
		descInfo = strings.TrimSpace(descInfo)
		desc := strings.Split(descInfo, "\n")
		descOne := desc[0]
		descTwo := desc[1]

		scoreStr := s.Find(".bd .star .rating_num").Text()
		score, _ := strconv.ParseFloat(scoreStr, 64)
		log.Printf("scoreStr:%v, score:%v\n", scoreStr, score)


		compile := regexp.MustCompile("[0-9]")
		commentStr := s.Find(".bd .star span").Eq(3).Text()
		commentStr = strings.TrimSpace(commentStr)
		commentStr = strings.Join(compile.FindAllString(commentStr, -1), "")
		comment, _ := strconv.ParseInt(commentStr, 10, 64)
		quote := s.Find(".quote .inq").Text()

		movie := &model.DouBanMovie{
			Title:    strings.TrimSpace(title),
			Subtitle: strings.TrimSpace(subtitle),
			Other:    strings.TrimSpace(other),
			DescOne:  strings.TrimSpace(descOne),
			DescTwo:  strings.TrimSpace(descTwo),
			Score:    score,
			Comment:  comment,
			Quote:    strings.TrimSpace(quote),
		}
		log.Printf("i:%v, movie:%v", i, movie)
		movies = append(movies, movie)
	})

	return movies
}
