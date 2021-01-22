package integration

import (
	"testing"
)

func Test_createGame(t *testing.T) {
	// pgURL, err, cleanUp := startDatabase()
	// defer cleanUp()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// db, err := dburl.Open(pgURL.String())
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if err := db.Ping(); err != nil {
	// 	t.Fatal(err)
	// }
	// ready := make(chan bool)
	// server := restapi.Server{
	// 	DB:    db,
	// 	Port:  "8080",
	// 	Ready: ready,
	// }
	// go server.Start()
	// <-ready

	// url := "http://localhost:8080/v1/games"

	// client := http.Client{
	// 	Timeout: time.Second * 30, // Timeout after 2 seconds
	// }

	// req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(`{"player1": "Rick","player2": "Morty"}`))
	// if err != nil {
	// 	t.Fatalf("err: %v", err)
	// }

	// req.Header.Set("Content-Type", "application/json")

	// res, err := client.Do(req)
	// if err != nil {
	// 	t.Fatalf("err: %v", err)
	// }
	// defer res.Body.Close()
	// bodyBytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// bodyString := string(bodyBytes)
	// if res.StatusCode != http.StatusCreated {
	// 	t.Fatalf("res: %v", bodyString)
	// 	t.Fail()
	// }
}
