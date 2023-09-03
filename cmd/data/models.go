package data

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeout = 3 * time.Second

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		Product: Product{},
	}
}

type Models struct {
	Product Product
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Status      int       `json:"status"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
}

func (p *Product) GetAll() ([]*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, price, status, created_on, updated_on from products`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Status,
			&product.CreatedOn,
			&product.UpdatedOn,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}
	return products, nil
}

func (p *Product) GetByID(id int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, name, description, price, status, created_on, updated_on from products where id = $1`

	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.Status,
		&p.CreatedOn,
		&p.UpdatedOn,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `insert into products (name, description, price, status, created_on, updated_on) values ($1, $2, $3, $4, $5, $6) returning id`

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		p.Name,
		p.Description,
		p.Price,
		p.Status,
		time.Now(),
		time.Now(),
	).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `update products set name = $1, description = $2, price = $3, status = $4, updated_on = $5 where id = $6`

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		p.Name,
		p.Description,
		p.Price,
		p.Status,
		time.Now(),
		p.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// update status to 3 (deleted)
	query := `update products set status = 3, where id = $1`

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, p.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) Undelete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// update status to 1 (active)
	query := `update products set status = 1, where id = $1`

	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, p.ID)

	if err != nil {
		return err
	}

	return nil
}
