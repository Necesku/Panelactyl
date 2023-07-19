package database

import "database/sql";
import "errors";
import _ "github.com/mattn/go-sqlite3";
import "teletvbis/panelactyl/backend/models";

var (
	ErrUserExists = errors.New("User already exists in the database!");
)

func FindUser(name string) (error, models.User) {
	var user models.User;
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

func CreateUser(username string, password string) (error) {
	err, _ := FindUser(username);
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			db, err := sql.Open("sqlite3", "data.db");
			if err != nil {
				return err;
			}
			defer db.Close();
			stmt, err := db.Prepare(`INSERT INTO users (name, password) VALUES (?, ?)`);
			if err != nil {
				return err;
			}
			_, err = stmt.Exec(username, password);
			return nil;
		}
		return err;
	}
	return ErrUserExists;
}