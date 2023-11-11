package mapping

import (
	"database/sql"
	"gin-twitter/types"
	"log"
	"time"
)

func MapPerson(row *sql.Row) (*types.Person, error) {
	var id int
	var firstname string
	var lastname string
	var userName string
	var verified bool
	var joinedUnix string

	err := row.Scan(&id, &firstname, &lastname, &userName, &verified, &joinedUnix)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	layout := "2006-01-02 15:04:05.999999999-07:00"

	joinedTime, parseErr := time.Parse(layout, joinedUnix)

	if parseErr != nil {
		log.Println(parseErr)
		return nil, parseErr
	}

	return &types.Person{
		ID:        id,
		FirstName: firstname,
		LastName:  lastname,
		UserName:  userName,
		Verified:  verified,
		Joined:    &joinedTime,
	}, nil
}
