package main

import (
	"fmt"
	"github.com/go-clog/clog"
	"github.com/wuleying/silver-xwlb/config"
	"github.com/wuleying/silver-xwlb/exceptions"
	"github.com/wuleying/silver-xwlb/globals"
	"github.com/wuleying/silver-xwlb/llog"
	"github.com/wuleying/silver-xwlb/metrics"
	"io/ioutil"
	"net/http"
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
	targetUrl := fmt.Sprintf(cfg["urls"]["xwlb_url"], globals.CurrentTime.AddDate(0, 0, -1).Format("20060102"))
	clog.Info("targetUrl = %s", targetUrl)

	resp, err := http.Get(targetUrl)

	if err != nil {
		clog.Fatal(globals.ClogDisplayInfo, "Get target url context failed: %s", err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		clog.Fatal(globals.ClogDisplayInfo, "Read context failed: %s", err.Error())
	}

	clog.Info(string(body))
}
