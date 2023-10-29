package main

import (
	"encoding/json"
	"os"
	"time"
)

type Saver struct {
	RSaver  *ReportSaver
	SSaver  *SentimentSaver
	HSaver  *HeadlineSaver
	PASaver *PriceAnalysisSaver
	PSaver  *PriceSaver
}

func NewSaver() Saver {
	return Saver{
		RSaver:  NewReportSaver(),
		SSaver:  NewSentimenttSaver(),
		HSaver:  NewHeadlineSaver(),
		PASaver: NewPriceAnalysisSaver(),
		PSaver:  NewPriceSaver(),
	}
}

type ReportSaver struct {
	fname string
}

func NewReportSaver(fname ...string) *ReportSaver {
	if len(fname) == 0 {
		return &ReportSaver{fname: "reports.json"}
	}
	return &ReportSaver{fname: fname[0]}
}

func (s *ReportSaver) Save(report string) error {
	tstamp := time.Now().Unix()
	reportInfo := ReportInfo{
		Tstamp: tstamp,
		Report: report,
	}

	var reportInfoOld []ReportInfo
	if fileExists(s.fname) {
		data, err := os.ReadFile(s.fname)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, &reportInfoOld); err != nil {
			return err
		}
	}

	reportInfoOld = append(reportInfoOld, reportInfo)

	reportInfoData, err := json.Marshal(reportInfoOld)
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.fname, reportInfoData, os.ModePerm); err != nil {
		return err
	}

	return nil
}

type SentimentSaver struct {
	fname string
}

func NewSentimenttSaver(fname ...string) *SentimentSaver {
	if len(fname) == 0 {
		return &SentimentSaver{fname: "sentiments.json"}
	}
	return &SentimentSaver{fname: fname[0]}
}

func (s *SentimentSaver) Save(sentiment string) error {
	tstamp := time.Now().Unix()
	sentimentInfo := Sentiment{
		Tstamp:    tstamp,
		Sentiment: sentiment,
	}

	var sentimentsOld []Sentiment
	if fileExists(s.fname) {
		data, err := os.ReadFile(s.fname)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, &sentimentsOld); err != nil {
			return err
		}
	}

	sentimentsOld = append(sentimentsOld, sentimentInfo)
	sentimentsData, err := json.Marshal(sentimentsOld)
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.fname, sentimentsData, os.ModePerm); err != nil {
		return err
	}

	return nil
}

type HeadlineSaver struct {
	fname string
}

func NewHeadlineSaver(fname ...string) *HeadlineSaver {
	if len(fname) == 0 {
		return &HeadlineSaver{fname: "headlines.json"}
	}
	return &HeadlineSaver{fname: fname[0]}
}

func (s *HeadlineSaver) Save(headlineInfo HeadlineInfo) error {
	var headlineInfoOld []HeadlineInfo
	if fileExists(s.fname) {
		data, err := os.ReadFile(s.fname)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, &headlineInfoOld); err != nil {
			return err
		}
	}

	headlineInfoOld = append(headlineInfoOld, headlineInfo)
	headlineInfoBytes, err := json.Marshal(headlineInfoOld)
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.fname, headlineInfoBytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}

type PriceSaver struct {
	fname string
}

func NewPriceSaver(fname ...string) *PriceSaver {
	if len(fname) == 0 {
		return &PriceSaver{fname: "prices.json"}
	}
	return &PriceSaver{fname: fname[0]}
}

func (s *PriceSaver) Save(prices []Price, tstamp int64) error {
	var pricesOld []PriceInfo
	if fileExists(s.fname) {
		data, err := os.ReadFile(s.fname)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, &pricesOld); err != nil {
			return err
		}
	}

	pricesOld = append(pricesOld, PriceInfo{Tstamp: tstamp, Prices: prices})
	priceBytes, err := json.Marshal(pricesOld)
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.fname, priceBytes, os.ModePerm); err != nil {
		return err
	}
	return nil
}

type PriceAnalysisSaver struct {
	fname string
}

func NewPriceAnalysisSaver(fname ...string) *PriceAnalysisSaver {
	if len(fname) == 0 {
		return &PriceAnalysisSaver{fname: "price_analysis.json"}
	}
	return &PriceAnalysisSaver{fname: fname[0]}
}

func (s *PriceAnalysisSaver) Save(priceAnalysis string, tstamp int64) error {
	reportInfo := ReportInfo{
		Tstamp: tstamp,
		Report: priceAnalysis,
	}

	var reportInfoOld []ReportInfo
	if fileExists(s.fname) {
		data, err := os.ReadFile(s.fname)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, &reportInfoOld); err != nil {
			return err
		}
	}

	reportInfoOld = append(reportInfoOld, reportInfo)

	reportInfoData, err := json.Marshal(reportInfoOld)
	if err != nil {
		return err
	}

	if err := os.WriteFile(s.fname, reportInfoData, os.ModePerm); err != nil {
		return err
	}

	return nil
}
