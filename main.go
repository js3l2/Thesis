package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Web service up...")

	http.HandleFunc("/", Homepage)
	http.HandleFunc("/vote", Vote)
	http.ListenAndServe(":8094", nil)

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

	fmt.Fprintf(w, "Thank you for your time!")

}

func Homepage(w http.ResponseWriter, req *http.Request) {
	log.Warnf("Connection established...%v", req)
	fmt.Fprint(w, `<h1>Success..</h1>

	<body>

	<font face="verdana" size="5" color="#7a005c" <h1>Contact Form</h1><br>
	  </body>
	  <font face="verdana" size="2" color="black" </font>
	  <br>
  
	  <form class="" action="/vote" method="post">
  
		<label for="">Name</label>
		<input type="text" name="info" value="">

		<label for="">email</label>
		<input type="email" name="mail" value=""> <br>
		 
		<label for="">Message</label> <br>
		<textarea name="inputVal" rows="8" cols="52"></textarea>
		<input type="submit" name="">

		
	  </form>
    


	
	`)

}
