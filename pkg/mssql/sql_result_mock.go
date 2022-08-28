package mssql

type MockResult struct{}

func (m *MockResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (m *MockResult) RowsAffected() (int64, error) {
	return 1, nil
}
