package main

import (
	"fmt"
	"net/http"
)

func main() {
	//define handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title>Welcome</title>
			</head>	
			<body>
				<h1>Learn about me</h1>
				<p>Check out the other pages by adding /about or /hobbies to the end of the url</p>
			</body>
		</html>
		
		`)
	})

	//Sports Page

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
		<head>
			<title>About Page</title>
		</head>	
		<body>
			<h1>What to know about me</h1>
			<p>I was born in Arizona then lived in Missouri for 10 years then Texas until I came to Point Park!</p>
			<img src = "https://www.nps.gov/articles/000/images/1164.jpg?maxwidth=650&autorotate=false&quality=78&format=webp" height="250", width="250"
			<p>
		</body>
		</html>
		
		
		`)
	})
	//Hobby page

	http.HandleFunc("/hobbies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<title>My Hobbies</title>
			</head>	
			<body>
				<h1>What do I do on my free time?</h1>
				<p>Play Soccer (lots of soccer)</p>
				<p>Play video games</p>
				<p>Play golf</p>
				<p>Bartending</p>

				<img src="https://fivethirtyeight.com/wp-content/uploads/2022/11/GettyImages-1442587075-e1668806020544.jpg?w=1834" height="250", width="250"
				
			</body>
		</html>
		
		`)
	})

	//start the service

	fmt.Println("Server listening to port 8080...")
	http.ListenAndServe(":8080", nil)
}
