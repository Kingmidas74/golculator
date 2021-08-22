package storage

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
)

type Database struct {
	Client *sql.DB
}

type DBOperation struct {
	Id             uuid.UUID
	Name           string
	ArgumentsCount int
	Priority       int
	Code           string
	PreviousId     uuid.UUID
}

func (db *Database) Initialize(host, port, user, password, dbname string) error {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s  dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db.Client, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Up() error {
	const createExtension = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	const tableOperationCreationQuery = `CREATE TABLE IF NOT EXISTS operations
	(
		id uuid DEFAULT uuid_generate_v4(),
		name TEXT NOT NULL,
		created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		arguments_count INT,
		priority INT,
		code TEXT NOT NULL,
		previous_id uuid,
		CONSTRAINT operations_pkey PRIMARY KEY (id)
	)`

	if _, err := db.Client.Exec(createExtension); err != nil {
		return err
	}

	if _, err := db.Client.Exec(tableOperationCreationQuery); err != nil {
		return err
	}
	return nil
}

func (db *Database) GetOperations() ([]DBOperation, error) {

	query := `WITH temp AS (select name, MAX(created_date) as cd FROM public.operations GROUP BY name)        
        SELECT s.id, t.name, s.arguments_count, s.priority, s.code FROM temp t
        JOIN public.operations s ON t.cd = s.created_date AND t.name=s.name
        ORDER BY t.name DESC
       `

	rows, err := db.Client.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var operations []DBOperation

	for rows.Next() {
		var operation DBOperation
		if err := rows.Scan(&operation.Id, &operation.Name, &operation.ArgumentsCount, &operation.Priority, &operation.Code); err != nil {
			return nil, err
		}
		operations = append(operations, operation)
	}
	return operations, nil
}

func (db *Database) CreateOperation(o DBOperation) {
	db.Client.QueryRow(
		"INSERT INTO public.operations(name, arguments_count, priority, code) VALUES($1, $2, $3, $4) RETURNING id",
		o.Name, o.ArgumentsCount, o.Priority, o.Code)
}
