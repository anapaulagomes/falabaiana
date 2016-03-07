package main

import (
	"net/http"
	"time"
	"math/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Expression struct {
    ID    int    `json:"id"`
    Phrase string `json:"phrase"`
    Meaning   string `json:"meaning"`
}

func (e Expression) toString() string {
    return toJson(e)
}

func toJson(p interface{}) string {
    bytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}


func main(){

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/expression", ChooseExpression)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":" + port, nil)
}

func ChooseExpression(response http.ResponseWriter, request *http.Request){
	expression := GetExpression()
	response.Write([]byte(expression))
}

func GetExpression() string {

	file, err := ioutil.ReadFile("./expressions.json")

	if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
	}

	var c []Expression
	json.Unmarshal(file, &c)

	randomId := random(0, len(c))

	return c[randomId].toString()
}

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
