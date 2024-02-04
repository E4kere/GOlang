package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Player struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Height   int    `json:"appearances"`
}

var players []Player

func init() {
	players = []Player{
		{ID: 1, Name: "David de Gea", Position: "Goalkeeper", Height: 192},
		{ID: 2, Name: "Victor Lindelöf", Position: "Defender", Height: 190},
		{ID: 3, Name: "Raphael Varane", Position: "Defender", Height: 192},
	}
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, position, appearances FROM players")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.ID, &player.Name, &player.Position, &player.Height)
		if err != nil {
			fmt.Println(err)
			return
		}
		players = append(players, player)
	}

	json.NewEncoder(w).Encode(players)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var player Player
	err = db.QueryRow("SELECT id, name, position, appearances FROM players WHERE id = ?", id).Scan(&player.ID, &player.Name, &player.Position, &player.Height)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(player)
}
