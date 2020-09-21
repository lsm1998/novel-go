package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestQuery(t *testing.T) {
	db, err := sql.Open("mysql", "lsm:123456@tcp(47.103.211.234:9696)/novel")
	if err != nil {
		panic(err)
	}

	row, _ := db.Query("select count(*) from test_shard_hash")
	var count int
	if row.Next() {
		row.Scan(&count)
	}
	fmt.Println(count)
}

func TestInsert(t *testing.T) {
	db, err := sql.Open("mysql", "lsm:123456@tcp(47.103.211.234:9696)/novel")
	if err != nil {
		panic(err)
	}

	tx, _ := db.Begin()
	stmt, err := tx.Prepare("insert into test_shard_hash(id,str,f,e,u,i) values(?,?,?,?,?,?);")
	if err != nil {
		panic(err)
	}
	fmt.Println(stmt.Exec(101, "haha", 2.1, "技术部", 4, 18))
	_ = tx.Commit()
}
