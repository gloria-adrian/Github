package main

import(
"fmt"
"log"
"net/http"
"encoding/json"
)
type Article struct{
  Title string`json:"Title"`
  Desc string`json:"desc"`
  Content string`json:"content"`
}
 var Articles []Article

func homePage(w http.ResponseWriter,r *http.Request){
fmt.Fprintf(w,"welcome to the HomePage!")
fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter,r *http.Request){
  fmt.Println("Endpoint Hit:returnAllArticles")
  json.NewEncoder(w).Encode(Articles)
}

func handleRequests(){
http.HandleFunc("/", homePage)
http.HandleFunc("/articles", returnAllArticles)
log.Fatal(http.ListenAndServe(":10000",nil))
}

func main() {
  Articles=[]Article{
  Article{Title:"Hello",Desc: "article Description",Content: "Article Content"},
  Article{Title:"Hello 2",Desc:"Article Description",Content:"Article Content"},
}
handleRequests()
handleRequests()
}
