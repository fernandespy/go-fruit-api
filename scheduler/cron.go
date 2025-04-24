package scheduler

import (
	"go-fruit-api/crawler"
	"log"

	"github.com/robfig/cron/v3"
)

func StartScheduler() {
	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", crawler.RunCrawler)
	if err != nil {
		log.Println("Error to schedule crawler")
		return
	}
	c.Start()
}
