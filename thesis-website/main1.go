package main

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("Web service up...")

	http.Handle("/", http.FileServer(http.Dir("./static")))

  log.Fatal(http.ListenAndServe(":8081", nil))
	http.HandleFunc("/vote", Vote)

}

func Vote(w http.ResponseWriter, req *http.Request) {
	log.Infof("Received Post value")

	err := req.ParseForm()
	if err != nil {
		log.Error("Can not parse form!")
	}

	vote := req.PostFormValue("inputVal")
	inf0 := req.PostFormValue("info")
	mail := req.PostFormValue("mail")

	log.Infof("Value is %v", vote)
	log.Infof("Info is %v", inf0)
	log.Infof("Mail is %v", mail)

	ProcessVote(vote)

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

func Homepage(w http.ResponseWriter, req *http.Request) {
	log.Warnf("Connection established...%v", req)
	fmt.Fprint(w, `
<body>
  <div class="container">
    <font face="verdana" size="5" position="absolute" color="#7a005c" <h1>Please enter a value to be forwarded to Kubernetes</h1><br>
      <font face="verdana" size="2" color="#00e9f4" </font>
        <br>

        <form class="" action="/vote" method="post">

          <label for="">Name</label>
          <input type="text" name="info" value="">

          <label for="">email</label>
          <input type="email" name="mail" value=""> <br>

          <label for="">Message</label> <br>
          <textarea name="inputVal" id="ar" rows="10" cols="55"></textarea> <br>
          <a href="#">
            <span></span>
            <span></span>
            <span></span>
            <span></span>
            <input type="submit" id="submit" name="">
          </a>

        </form>
</body>
`)

}
