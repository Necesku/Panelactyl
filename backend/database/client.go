package database

import "database/sql";
import _ "github.com/mattn/go-sqlite3";
import "teletvbis/panelactyl/backend/modals";

func FindUser(name string) (error, modals.User) {
	var user modals.User;
	db, err := sql.Open("sqlite3", "data.db");
	if err != nil {
		return err, user;
	}
	defer db.Close();
	stmt, err := db.Prepare(`SELECT * FROM users WHERE name = ?`);
	if err != nil {
		return err, user;
	}
	defer stmt.Close();
	err = stmt.QueryRow(name).Scan(&user.Name, &user.Password);
	if err != nil {
		return err, user;
	};
	return nil, user;
}