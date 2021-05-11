package main

import (
	"context"
	"database/sql"
	"fmt"
	models "test_sqlboiler/my_models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestUsers(t *testing.T) {
	db, _ := sql.Open("mysql", "root:my-secret-pw@tcp(test:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	ctx := context.Background()

	// Begin a transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err.Error())
	}
	func() {
		defer func() {
			// Rollback or commit
			if r := recover(); r != nil {
				// fmt.Println(r.(error))
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()
		models.Users(qm.Where("username=?", "foo")).DeleteAll(ctx, tx)
		foo := models.User{
			Username: "foo",
			Passwd:   "bar",
		}
		err := foo.Insert(ctx, tx, boil.Infer())
		if err != nil {
			panic(err.Error())
		}
		one, err := models.Users(qm.Where("username=?", "foo")).One(ctx, tx)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(one)
	}()

	users, err := models.Users().All(ctx, db)
	if err != nil {
		panic(err.Error())
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
