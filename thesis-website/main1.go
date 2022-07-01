package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	"os"
	"time"
	"io/ioutil"

	"github.com/Jeffail/gabs"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func main() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("❌ ", err.Error())
	}

	log.SetLevel(log.DebugLevel)
	log.Info("Web service up...")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/vote", Vote)

	http.ListenAndServe(":8081", nil)

}

func Vote(w http.ResponseWriter, req *http.Request) {
  log.Infof("Received Pressure")

	var location Location

	err := json.NewDecoder(req.Body).Decode(&location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	var apiKey = os.Getenv("OWM_API_KEY")

	url := fmt.Sprintf(`https://api.openweathermap.org/data/2.5/onecall?lat=%.2f&lon=%.2f&exclude=&appid=%s`, location.Lat, location.Long, apiKey)

	c := http.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := c.Get(url)
	if err != nil {
		fmt.Printf("Error %s", err)

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)

	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	var value float64 // this is pressure///
	var ok bool

	value, ok = jsonParsed.Path("current.pressure").Data().(float64)
	if ok {
		fmt.Println("✅", value)

		fmt.Fprintf(w, `{"pressure":"%.2f"}`, value,)

	}

}


func ProcessVote(vote string) int {
	log.Warnf("Into ProcessVote with parameter: %v", vote)
	_, err := strconv.Atoi(vote)
	if err != nil {
		log.Errorf("User entered a string!", vote)
		log.Fatal(err)
	}
	return 0
}
