package handlers

import (
	"fmt"

	"go-fiber-pg/db"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
)

// 🔁 Seed smartly: keeps adding 100,000 new keys each time
func SeedRecords(c *fiber.Ctx) error {
	var existing int
	_ = db.DB.QueryRow(db.Ctx, "SELECT COUNT(*) FROM records").Scan(&existing)

	start := existing + 1
	end := existing + 100000

	for i := start; i <= end; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := gofakeit.HipsterSentence(6)

		_, err := db.DB.Exec(db.Ctx,
			`INSERT INTO records (key, value) 
			 VALUES ($1, $2) 
			 ON CONFLICT (key) 
			 DO UPDATE SET value = EXCLUDED.value`,
			key, value)

		if err != nil {
			return c.Status(500).SendString(fmt.Sprintf("?? Insert failed at %d: %v", i, err))
		}
	}

	return c.SendString(fmt.Sprintf("//Inserted/updated %d records", end-start+1))
}

func GetAllRecords(c *fiber.Ctx) error {
	rows, err := db.DB.Query(db.Ctx, "SELECT key, value FROM records")
	if err != nil {
		return c.Status(500).SendString("?? Error fetching records: " + err.Error())
	}
	defer rows.Close()

	type Record struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	var records []Record

	for rows.Next() {
		var r Record
		if err := rows.Scan(&r.Key, &r.Value); err != nil {
			return c.Status(500).SendString("??Scan error: " + err.Error())
		}
		records = append(records, r)
	}

	return c.JSON(records)
}

func CreateRecord(c *fiber.Ctx) error {
	type Record struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	var rec Record

	if err := c.BodyParser(&rec); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	_, err := db.DB.Exec(db.Ctx, "INSERT INTO records (key, value) VALUES ($1, $2)", rec.Key, rec.Value)
	if err != nil {
		return c.Status(500).SendString("Insert failed: " + err.Error())
	}
	return c.SendString("// Record created")
}

func GetRecord(c *fiber.Ctx) error {
	key := c.Params("key")
	var value string

	err := db.DB.QueryRow(db.Ctx, "SELECT value FROM records WHERE key=$1", key).Scan(&value)
	if err != nil {
		return c.Status(404).SendString("Record not found")
	}

	return c.JSON(fiber.Map{
		"key":   key,
		"value": value,
	})
}

func UpdateRecord(c *fiber.Ctx) error {
	key := c.Params("key")
	type Payload struct {
		Value string `json:"value"`
	}
	var data Payload

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).SendString("Invalid body")
	}

	cmdTag, err := db.DB.Exec(db.Ctx, "UPDATE records SET value=$1 WHERE key=$2", data.Value, key)
	if err != nil {
		return c.Status(500).SendString("Update failed: " + err.Error())
	}
	if cmdTag.RowsAffected() == 0 {
		return c.Status(404).SendString("Record not found")
	}

	return c.SendString("// Record updated")
}

func DeleteRecord(c *fiber.Ctx) error {
	key := c.Params("key")

	cmdTag, err := db.DB.Exec(db.Ctx, "DELETE FROM records WHERE key=$1", key)
	if err != nil {
		return c.Status(500).SendString("Delete failed: " + err.Error())
	}
	if cmdTag.RowsAffected() == 0 {
		return c.Status(404).SendString("Record not found")
	}

	return c.SendString("🗑️ Record deleted")
}
