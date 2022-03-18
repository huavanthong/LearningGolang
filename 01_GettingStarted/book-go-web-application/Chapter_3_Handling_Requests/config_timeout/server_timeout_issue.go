package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var layout *template.Template

func main() {
	router := http.NewServeMux()
	server := &http.Server{
		Addr:        ":8080",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
		// Step 1: we set time to write a response only 1s
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	router.HandleFunc("/", home)

	var err error
	layout, err = template.ParseFiles("./layout.html")
	if err != nil {
		fmt.Printf("Error1: %+v\n", err)
	}

	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	// Step 2: Server access to handler function success
	fmt.Println("responding")
	// Step 3: We select a template and execute it
	// Template will success if this handler complate process.
	err := layout.Execute(w, template.HTML(`World`))
	if err != nil {
		fmt.Printf("Error2: %+v\n", err)
	}

	// Step 4: However, we sleep this thread until 5s > time to write a response
	time.Sleep(5 * time.Second)
	// Conclusion:
	// 	-	This is a reason why if we will see many responding from output console.
	//	-	And you also know that server will try entry handler function many time

}
