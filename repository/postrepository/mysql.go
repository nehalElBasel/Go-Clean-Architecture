package postrepository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nehal1992/Go-Clean-Architecture/models"
	"github.com/nehal1992/Go-Clean-Architecture/repository"
	"github.com/nehal1992/Go-Clean-Architecture/serviceErrors"
)

type mysqlPostRepo struct {
	Conn *sql.DB
}

var tablename = "posts"

func NewMysqlPost(conn *sql.DB) repository.PostRepo {
	return &mysqlPostRepo{
		Conn: conn,
	}
}

func (mysql *mysqlPostRepo) List(ctx context.Context) (posts []models.Post, err error) {

	query := "SELECT * FROM " + tablename

	return mysql.fetch(ctx, query)
}

func (mysql *mysqlPostRepo) Create(ctx context.Context, p models.Post) (interface{}, error) {
	query := fmt.Sprintf("Insert %s SET title=?, author=?", tablename)

	stmt, err := mysql.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Title, p.Author)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
func (mysql *mysqlPostRepo) Get(ctx context.Context, id int) (post models.Post, err error) {
	// Execute the query
	query := fmt.Sprintf("Select * FROM %s where id=?", tablename)

	rows, err := mysql.fetch(ctx, query, id)
	if err != nil {
		return post, err
	}

	payload := models.Post{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return post, serviceErrors.ErrNotFound
	}

	return payload, nil
}

func (mysql *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]models.Post, error) {
	rows, err := mysql.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := []models.Post{}
	for rows.Next() {
		data := models.Post{}

		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Author,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (mysql *mysqlPostRepo) Update(ctx context.Context, id int, p models.Post) (err error) {
	fmt.Println("post", id)
	query := fmt.Sprintf("Update %s set title=?,author=? where id=?", tablename)
	// query := "Update " + tablename + " set title=?,author=? where id=?"

	stmt, err := mysql.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Author,
		id,
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}

func (mysql *mysqlPostRepo) Delete(ctx context.Context, id int) (bool, error) {
	query := fmt.Sprintf("Delete From %s Where id=?", tablename)

	stmt, err := mysql.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
