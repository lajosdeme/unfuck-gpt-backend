package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	marketAnalysisCronString string
	priceUpdateCronString    string
}

func NewScheduler(marketCron string, priceCron string) Scheduler {
	return Scheduler{marketAnalysisCronString: marketCron, priceUpdateCronString: priceCron}
}

// "0 */5 * * *"
func (s *Scheduler) ScheduleMarketSentimentAnalysis() {
	c := cron.New()
	_, err := c.AddFunc(s.marketAnalysisCronString, RunMarketSentimentAnalysis)
	if err != nil {
		fmt.Println("error scheduling market analysis: ", err)
		return
	}

	c.Start()

	fmt.Println("Market sentiment analysis scheduled to run every 5 hours")

	select {}
}

func (s *Scheduler) SchedulePriceUpdate() {
	c := cron.New()
	chainQuery := NewChainQuery()
	_, err := c.AddFunc(s.priceUpdateCronString, chainQuery.RefreshPricesAndCreateReport)
	if err != nil {
		fmt.Println("error scheduling price update: ", err)
		return
	}

	c.Start()

	fmt.Println("Price update scheduled to run every hour")

	select {}
}
