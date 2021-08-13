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

	ProcessVote(vote)
	fmt.Fprint(w, `
  
  <h1 class="gradient-text">Thank you!</h1>
  

  <style>
    .gradient-text {

      background-color: red;


      background-image: linear-gradient(45deg, #c21599, #ffc500);

      background-size: 100%;
      background-repeat: repeat;

      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      -moz-background-clip: text;
      -moz-text-fill-color: transparent;
    }



    body {
      background-color: #181123;
    }

    header {
      margin-top: 1em;
      margin-top: calc(50vh - 3em);
    }

    h1 {
      font-family: Verdana;
      font-weight: normal;
      font-size: 3em;
      text-align: center;
      margin-bottom: 0;
      margin-bottom: -0.25em;
    }
  </style>
   `)

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


<style>
  .container {
    position: absolute;
    transform: translate(-50%, -50%);
    top: 50%;
    left: 50%;
  }

  form {
    background: $white;
    padding: 3em;
    height: 320px;
    border-radius: 20px;
    border-left: 1px solid $white;
    border-top: 1px solid $white;
    backdrop-filter: blur(10px);
    box-shadow: 20px 20px 40px -6px rgba(0, 0, 0, 0.2);
    text-align: center;
    position: relative;
    transition: all 0.2s ease-in-out;




  }

  input {
    background: transparent;
    width: 200px;
    padding: 1em;
    margin-bottom: 2em;
    border: none;
    border-left: 1px solid $white;
    border-top: 1px solid $white;
    border-radius: 5000px;
    backdrop-filter: blur(5px);
    box-shadow: 4px 4px 60px rgba(0, 0, 0, 0.2);
    color: #7a005c;
    font-family: Verdana;
    font-weight: 500;
    transition: all 0.2s ease-in-out;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2);


  }
  }
   

  .ar{
    resize: none;
  }

  #submit {
    text-transform: uppercase;
    padding: 5px;
    margin: 3px;
    width: 100px;
  }

  #submit:hover {
    background-color: #7a005c;
    color: #03e9f4;
    cursor: pointer;
  }
  .container form a {
    position: relative;
    display: inline-block;
    padding: 10px 20px;
    color: #03e9f4;
    font-size: 16px;
    text-decoration: none;
    text-transform: uppercase;
    overflow: hidden;
    transition: .5s;
    margin-top: 40px;
    letter-spacing: 4px
  }

  .container a:hover {
    background: #03e9f4;
    color: #fff;
    border-radius: 5px;
    box-shadow: 0 0 5px #03e9f4,
      0 0 25px #03e9f4,
      0 0 50px #03e9f4,
      0 0 100px #03e9f4;
  }

  .container a span {
    position: absolute;
    display: block;
  }

  .container a span:nth-child(1) {
    top: 0;
    left: -100%;
    width: 100%;
    height: 2px;
    background: linear-gradient(90deg, transparent, #03e9f4);
    animation: btn-anim1 1s linear infinite;
  }

  @keyframes btn-anim1 {
    0% {
      left: -100%;
    }

    50%,
    100% {
      left: 100%;
    }
  }

  .container a span:nth-child(2) {
    top: -100%;
    right: 0;
    width: 2px;
    height: 100%;
    background: linear-gradient(180deg, transparent, #03e9f4);
    animation: btn-anim2 1s linear infinite;
    animation-delay: .25s
  }

  @keyframes btn-anim2 {
    0% {
      top: -100%;
    }

    50%,
    100% {
      top: 100%;
    }
  }

  .container a span:nth-child(3) {
    bottom: 0;
    right: -100%;
    width: 100%;
    height: 2px;
    background: linear-gradient(270deg, transparent, #03e9f4);
    animation: btn-anim3 1s linear infinite;
    animation-delay: .5s
  }

  @keyframes btn-anim3 {
    0% {
      right: -100%;
    }

    50%,
    100% {
      right: 100%;
    }
  }

  .container a span:nth-child(4) {
    bottom: -100%;
    left: 0;
    width: 2px;
    height: 100%;
    background: linear-gradient(360deg, transparent, #03e9f4);
    animation: btn-anim4 1s linear infinite;
    animation-delay: .75s
  }

  @keyframes btn-anim4 {
    0% {
      bottom: -100%;
    }

    50%,
    100% {
      bottom: 100%;
    }
  }
</style>

	`)

}
