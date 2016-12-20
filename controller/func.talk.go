package controller

import (
	"fmt"
	"github.com/training_project/database"
	"github.com/training_project/model"
	"log"
)

func GetTalks(productId int64) []model.Talk {
	query := fmt.Sprintf(
		`
		Select
			*
		From
			test_table	
		where 
			idtest_table = %d
		`,
		productId)

	rows, err := database.MysqlDb.Query(query)

	if err != nil {
		log.Fatal("error ", err)
	}

	talks := []model.Talk{}
	//get each one of object from db
	for rows.Next() {
		t := model.Talk{}
		//get data from db and insert to Object
		if err := rows.Scan(
			&t.ProductId, &t.Message,
		); err != nil {
			log.Fatal("error = ", err)
		}

		talks = append(talks, t)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return talks
}
