package main

import ("fmt"
		"net/http")

func indexHandler(w http.ResponseWriter, r* http.Request){
	fmt.Fprintf(w, "<h1>I am here</h1>")
	fmt.Fprintf(w, "<p>Is is simple?</p>")
	fmt.Fprintf(w, "<p>Yes</p>")
	fmt.Fprintf(w, "<p>You can even add %s</p>","<strong>variable</strong>")

	fmt.Fprintf(w, `<p>Is is realy simple?</p>
					<p>Yes, it is realy simple</p>`)
}

func main(){
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}