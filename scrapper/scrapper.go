package scrapper

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type extractedJob struct{
	id			string
	title		string
	location	string
	salary		string
	summary		string
}

func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term +"&limit=50"
	var jobs []extractedJob
	ch := make(chan [] extractedJob)
	totalPages := getPageNumber(baseURL)
	for i := 0 ; i < totalPages ; i++{
		go getPage(i, baseURL, ch)
	}
	for i := 0 ; i < totalPages ; i++{
		extractedJobs := <-ch
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(term,jobs)
	fmt.Println("Done.!!")
}

func getPage(page int, url string, mainC chan<- []extractedJob){
	var jobs []extractedJob
	ch := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page * 50)
	fmt.Println("Requesting", pageURL)
	resp, err := http.Get(pageURL)
	checkErr(err)
	checkCode(resp)

	defer  resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection){
		go extractJob(card, ch)
	})
	for i := 0; i < searchCards.Length();i++{
		job := <-ch
		jobs = append(jobs, job)
	}
	mainC <- jobs

}

func extractJob(card *goquery.Selection, ch chan<- extractedJob){
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())
	ch <- extractedJob{
		id: id,
		title: title,
		location: location,
		salary: salary,
		summary: summary,
	}
}

func getPageNumber(url string) int{
	pages := 0
	resp, err := http.Get(url)
	checkErr(err)
	checkCode(resp)
	defer resp.Body.Close()

	doc, err:= goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, sel *goquery.Selection){
		pages = sel.Find("a").Length()
	})
	return pages
}

func writeJobs(term string,jobs [] extractedJob){
	file, err := os.Create(term +"jobs.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Loc", "Sal", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)
	for _, job := range jobs{
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk="+ job.id,job.title,job.location,job.title,job.salary,job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}


func CleanString(str string) string{
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}


func checkErr(err error){
	if err != nil{
		log.Fatalln(err)
	}
}

func checkCode(resp *http.Response){
	if resp.StatusCode != 200{
		log.Fatalln("Request failed with Status:", resp.StatusCode)
	}
}






