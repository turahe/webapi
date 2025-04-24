package migrations

import (
	"context"

	"github.com/turahe/interpesona-data/internal/db/pgx"
)

func init() {
	Migrations = append(Migrations, createUserTable)
}

var createUserTable = &Migration{
	Name: "20230407151155_create_user_table",
	Up: func() error {
		_, err := pgx.GetPgxPool().Exec(context.Background(), `
			CREATE TABLE users (
				"id" UUID PRIMARY KEY,
				"username" VARCHAR(255) NOT NULL,
				"email" VARCHAR(255) NOT NULL UNIQUE,
			    "phone" VARCHAR(255) NULL UNIQUE,
			    "password" VARCHAR(255) NULL,
			    "email_verified_at" TIMESTAMP NULL,
			    "phone_verified_at" TIMESTAMP NULL,
			    "created_at" TIMESTAMP DEFAULT NOW(),
			    "updated_at" TIMESTAMP DEFAULT NOW(),
			    "deleted_at" TIMESTAMP NULL
			);

		`)

		if err != nil {
			return err
		}
		return nil

	},
	Down: func() error {
		_, err := pgx.GetPgxPool().Exec(context.Background(), `
			DROP TABLE IF EXISTS users;
		`)
		if err != nil {
			return err
		}

		return nil
	},
}
