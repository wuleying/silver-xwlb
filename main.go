package main

import (
	"github.com/wuleying/silver-xwlb/config"
	"github.com/wuleying/silver-xwlb/exceptions"
	"github.com/wuleying/silver-xwlb/llog"
	"github.com/wuleying/silver-xwlb/metrics"
)

func main() {
	// http://tv.cctv.com/lm/xwlb/day/20180830.shtml
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
}
