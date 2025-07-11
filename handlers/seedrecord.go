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

	return c.SendString(fmt.Sprintf("// Successfully seeded %d new records", end-start+1))
}
