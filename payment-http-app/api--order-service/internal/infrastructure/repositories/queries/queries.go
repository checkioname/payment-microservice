package queries

const (
	SelectOrderQuery = `SELECT * FROM order`

	RegisterOrder = `
	INSERT INTO order (ID, CLIENT_ID, CREATION_DATE, STATUS, TOTAL_AMOUNT) 
	VALUES ($1, $2, NOW(), $3, $4)
	ON CONFLICT (ID)
	DO UPDATE
	SET 
		STATUS = EXCLUDED.STATUS
	`
)
