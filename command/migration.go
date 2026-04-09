package command

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" // Driver wajib di-import kosong
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iqbalatma/gartisan/utils"
	"github.com/joho/godotenv"
)

func getDatabaseDSN() string {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Failed to read file .env: %s", err)
	}

	return os.Getenv("GARTISAN_DATABASE_URL")
}
func MigrateUp(arguments []string) {
	fmt.Println("Migration up")
	var dsn = getDatabaseDSN()

	migration, err := migrate.New(
		"file://database/migrations",
		dsn,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer migration.Close()
	if len(arguments) > 3 && arguments[2] == "force" {
		version, err := strconv.Atoi(arguments[3])
		if err != nil {
			log.Fatal("Migration version is not an integer!")
		}

		fmt.Printf("Forcing migration version to %d...\n", version)
		err = migration.Force(version)
		if err != nil {
			log.Fatal("Failed to run migrate force: ", err)
		}
		log.Println(utils.ANSI_GREEN + "Migrate force successfully !.")
		return
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println(utils.ANSI_GREEN + "Migration success")
}

func MigrateDown(arguments []string) {
	fmt.Println("Migration down")
	var dsn = getDatabaseDSN()

	migration, err := migrate.New(
		"file://database/migrations",
		dsn,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer migration.Close()

	if len(arguments) > 2 {
		step, err := strconv.Atoi(arguments[2])
		if err != nil {
			return
		}
		err = migration.Steps(-step)
	} else {
		err = migration.Down()
	}
	err = migration.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println(utils.ANSI_GREEN + "Migration success")
}

func MakeMigration(arguments []string) {
	fmt.Println("Make migration")
	fmt.Println("Creating migration file...")

	if len(arguments) < 3 {
		log.Fatal("Please enter migration name. Example: gartisan make:migration create_users_table")
	}

	migrationName := arguments[2]
	migrationName = strings.ReplaceAll(migrationName, " ", "_")

	timestamp := time.Now().Format("20060102150405")

	//next update use dynamic path for this
	migrationDir := "database/migrations"

	if err := os.MkdirAll(migrationDir, fs.ModePerm); err != nil {
		log.Fatal("Failed to create folder for migration: ", err)
	}

	upFileName := fmt.Sprintf("%s_%s.up.sql", timestamp, migrationName)
	downFileName := fmt.Sprintf("%s_%s.down.sql", timestamp, migrationName)

	upFilePath := filepath.Join(migrationDir, upFileName)
	err := os.WriteFile(upFilePath, []byte("-- Write your query here\n"), 0644)
	if err != nil {
		log.Fatal("Failed to create file up: ", err)
	}

	// Buat file DOWN
	downFilePath := filepath.Join(migrationDir, downFileName)
	err = os.WriteFile(downFilePath, []byte("-- Write your query here \n"), 0644)
	if err != nil {
		log.Fatal("Failed to create file down: ", err)
	}

	fmt.Println("----------------------------------------")
	log.Println("Migration files created successfully:")
	fmt.Println("->", upFilePath)
	fmt.Println("->", downFilePath)
	fmt.Println("----------------------------------------")
}
