package storage

type Storage interface {
	// CreateTable creates a table with the given name and schema.
	CreateTable(tableName, schema string) error

	// Get retrieves data from the specified table based on the provided criteria.
	Get(tableName, whereCondition string, args ...interface{}) (interface{}, error)

	// Create inserts a new record into the specified table.
	Create(tableName string, data interface{}) error

	// Update updates a record in the specified table based on the provided criteria.
	Update(tableName, setClause, whereCondition string, args ...interface{}) error

	// Delete removes records from the specified table based on the provided criteria.
	Delete(tableName, whereCondition string, args ...interface{}) error
}
