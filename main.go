package main

import (
	"fmt"
	"time"
)

/*
todo: schedule a cron job to pull sentiment 5 hours
todo: get price from smart contract and analyse it
todo: open up API for FE to query this
*/

func main() {
	LoadConfig()
	// 0 0 */5 * * *
	scheduler := NewScheduler("0 */5 * * *", "1 */1 * * *")
	go scheduler.ScheduleMarketSentimentAnalysis()
	go scheduler.SchedulePriceUpdate()

	r := ConfigRouter()
	RunRouter(r)
}

func RunMarketSentimentAnalysis() {
	fmt.Println("Market analysis is running at: ", time.Now())
	scraper := NewScraper("https://cointelegraph.com/")
	headlineInfos, err := scraper.Scrape()
	if err != nil {
		fmt.Println("error getting headline infos: ", err)
		return
	}
	headlineSaver := NewHeadlineSaver()
	if err := headlineSaver.Save(headlineInfos); err != nil {
		fmt.Println("error saving headline infos: ", err)
		return
	}

	prompt := G.GenerateHeadlinePrompt(headlineInfos)
	resp := G.GetResponse(prompt)
	if err := S.RSaver.Save(resp); err != nil {
		fmt.Println("error saving report: ", err)
		return
	}

	prompt = G.GenerateSentimentPrompt(headlineInfos)
	resp = G.GetResponse(prompt)
	if err := S.SSaver.Save(resp); err != nil {
		fmt.Println("error saving sentiment: ", err)
		return
	}
}
