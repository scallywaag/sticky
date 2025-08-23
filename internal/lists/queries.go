package lists

const (
	GetActiveListSQL = `
		SELECT l.id, l.name
	    FROM state s
	    JOIN lists l ON s.list_id = l.id
	    WHERE s.key = 'active';
	`
)
