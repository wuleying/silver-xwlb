package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-clog/clog"
	"github.com/huichen/sego"
	"github.com/wuleying/silver-xwlb/config"
	"github.com/wuleying/silver-xwlb/exceptions"
	"github.com/wuleying/silver-xwlb/globals"
	"github.com/wuleying/silver-xwlb/llog"
	"github.com/wuleying/silver-xwlb/metrics"
)

func main() {
	// Log init
	llog.Init()
	defer llog.Shutdown()

	cfg, err := config.Init()
	exceptions.CheckError(err)

	metric := metrics.Metric{
		Host:     cfg["metrics"]["host"],
		Port:     cfg["metrics"]["port"],
		Database: cfg["metrics"]["database"],
		Username: cfg["metrics"]["username"],
		Password: cfg["metrics"]["password"],
	}
	metric.Init()

	// 抓取目标页
	targetURL := fmt.Sprintf(
		cfg["urls"]["xwlbURL"],
		globals.CurrentTime.AddDate(0, 0, -1).Format("20060102"))
	clog.Info("targetURL = %s", targetURL)

	doc, err := goquery.NewDocument(targetURL)
	exceptions.CheckError(err)

	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(fmt.Sprintf("%s%s", globals.RootDir, "/data/dictionary.txt"))

	doc.Find("ul li").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title").Text()
		href, _ := contentSelection.Find("a").Attr("href")

		clog.Info("%d. %s, %s", i+1, title, href)

		// 中文分词
		segments := segmenter.Segment([]byte(title))
		clog.Info(sego.SegmentsToString(segments, false))
	})
}
