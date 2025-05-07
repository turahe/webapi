package migrations

import (
	"context"

	"webapi/internal/db/pgx"
)

func init() {
	Migrations = append(Migrations, createSettingTable)
}

var createSettingTable = &Migration{
	Name: "20250507232044_create_setting_table",
	Up: func() error {
		_, err := pgx.GetPgxPool().Exec(context.Background(), `
			CREATE TABLE settings (
				 "id" UUID NOT NULL,
				"model_type" VARCHAR(255) NOT NULL,
				"model_id" VARCHAR(255) NOT NULL,
			    "key" VARCHAR(255) NULL UNIQUE,
			    "value" VARCHAR(255) NULL,
			    "created_by" UUID NULL,
			    "updated_by" UUID NULL,
			    "deleted_by" UUID NULL,
			    "deleted_at" TIMESTAMP DEFAULT NOW(),
			    "created_at" TIMESTAMP DEFAULT NOW(),
			    "updated_at" TIMESTAMP NULL,
			    CONSTRAINT "setting_pkey" PRIMARY KEY ("id"),
			    CONSTRAINT "setting_created_by_foreign" FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON DELETE SET NULL ON UPDATE NO ACTION,
			    CONSTRAINT "setting_deleted_by_foreign" FOREIGN KEY ("deleted_by") REFERENCES "users" ("id") ON DELETE SET NULL ON UPDATE NO ACTION,
			    CONSTRAINT "setting_updated_by_foreign" FOREIGN KEY ("updated_by") REFERENCES "users" ("id") ON DELETE SET NULL ON UPDATE NO ACTION
			    
			);
		`)

		if err != nil {
			return err
		}
		return nil

	},
	Down: func() error {
		_, err := pgx.GetPgxPool().Exec(context.Background(), `
			DROP TABLE IF EXISTS settings;
		`)
		if err != nil {
			return err
		}

		return nil
	},
}
