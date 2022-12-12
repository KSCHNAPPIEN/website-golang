package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func main() {

	// Maak een nieuwe ServeMux-struct
	mux := http.NewServeMux()

	// Deze route stuurt alle requests naar de "/" URL naar de
	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/Locatie", LocatieHandler)

	mux.HandleFunc("/Login", LoginHandler)

	http.ListenAndServe(":8080", mux)

	// Lees het JSON-bestand met de configuratie-instellingen voor de database
	configBytes, err := ioutil.ReadFile("Db.json")
	if err != nil {
		log.Fatal(err)
	}

	// Maak een nieuwe struct voor de configuratie-instellingen
	var config DatabaseConfig
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatal(err)
	}

	// Maak de verbinding met de database met de opgegeven configuratie-instellingen
	dbc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", dbc)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Controleer of de verbinding met de database werkt door een testquery uit te voeren
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
        <h1>Fonteyn Vakantieparken</h1>
        <p>Vakantie parken voor een onvergeetelijke vakantie.</p>
        <ul>
            <li><a href="/Locatie">Locatie</a></li>
            <li><a href="/Login">Login</a></li>
        </ul>
    `)
}

func LocatieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
        <h1>Locatie</h1>
        <p>Wij zijn beschikbaar in meerder landen.</p>
        <ul>
            <li><a href="/">Home</a></li>
            <li><a href="/Login">Login</a></li>
        </ul>
    `)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// Serve the login form
		fmt.Fprintln(w, `
            <h1>Login</h1>
            <p>Login with your credentials.</p>
            <ul>
                <li><a href="/">Home</a></li>
                <li><a href="/Locatie">Location</a></li>
            </ul>
            <form action="/Login" method="POST">
                <label for="username">Username:</label>
                <input type="text" id="username" name="username"><br>
                <label for="password">Password:</label>
                <input type="password" id="password" name="password"><br>
                <input type="submit" value="Submit">
            </form>
        `)
	}
}
