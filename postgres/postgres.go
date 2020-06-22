package postgres

import (
	"database/sql"
	"fmt"
	// postgres driver
	_ "github.com/lib/pq"
)

const(
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "blue42"
	DB_NAME     = "widgetdb"
)

type Db struct {
	*sql.DB 
}

// sourced from https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356 
func New() (*Db, error) {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}


type Feedback struct {
	ID        int
	Rating    int
	Details   string
	Email     string 
}

// sourced from https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356 
func (d *Db) GetFeedbacks( ) []Feedback{
	// Prepare query, takes a name argument, protects from sql injection
	stmt, err := d.Prepare("SELECT id, rating, details, email FROM feedback")
	fmt.Println("error ", err)
	if err != nil {
		fmt.Println("GetUserByName Preperation Err: ", err)
	}

	// Make query with our stmt, passing in name argument
	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("GetFeedbacks Query Err: ", err)
	}

	// Create slice of Feedbacks for our response
	var feedbacks []Feedback
	// Copy the columns from row into the values pointed at by r (Feedback)
	for rows.Next() {
		fdbk := Feedback{}
		err = rows.Scan(&fdbk.ID, &fdbk.Rating, &fdbk.Details, &fdbk.Email)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		feedbacks = append(feedbacks, fdbk)
	}

	return feedbacks
}

func (d *Db) CreateFeedback(rating int, email string, details string) Feedback{

	var lastInsertId int

	d.QueryRow("INSERT INTO authors(rating, email, details) VALUES($1, $2, $3) returning id;", rating, email, details).Scan(&lastInsertId)

	fbdkRet := Feedback{
		ID:        	lastInsertId,
		Details:    details,
		Email:     	email,
		Rating: 	rating,
	}

	return fbdkRet
}