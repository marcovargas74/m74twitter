package main

import "fmt"

/*package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/ChimeraCoder/anaconda"
)

func geTwitter() {

	var CONSUMER_KEY = os.Getenv("CONSUMER_KEY")
	var CONSUMER_SECRET = os.Getenv("CONSUMER_SECRET")
	var ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	var ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")

	var api *anaconda.TwitterApi
	//var testBase string

	// Initialize api so it can be used even when invidual tests are run in isolation
	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	api = anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	//api := anaconda.NewTwitterApiWithCredentials("your-access-token", "your-access-token-secret", "your-consumer-key", "your-consumer-secret")

	searchResult, err := api.GetSearch("golang", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}

}

// Titulo obtem o titulo de uma pagina html
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func getBody(url string) string {
	resp, _ := http.Get(url)
	html, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("Test:" + string(html))
	return string(html)
}

func main() {

	fmt.Println("Conect ao Twiter do Firecast")
	//fmt.Println(getTwitter("https://twitter.com/CBMSC193?ref_src=twsrc%5Egoogle%7Ctwcamp%5Eserp%7Ctwgr%5Eauthor"))
	//getBody("https://twitter.com/CBMSC193?ref_src=twsrc%5Egoogle%7Ctwcamp%5Eserp%7Ctwgr%5Eauthor")

	geTwitter()

}
*/

func main() {
	fmt.Println("twiiter")
}
