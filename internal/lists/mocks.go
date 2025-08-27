package lists

type MockListRepo struct {
	count int
	lists []List
	err   error
}

func (m *MockListRepo) Count() (int, error) {
	return m.count, m.err
}

func (m *MockListRepo) GetAll() ([]List, error) {
	return m.lists, m.err
}

func (m *MockListRepo) Add(name string) (int, error) {
	return 0, nil
}

func (m *MockListRepo) Delete(id int) error {
	return nil
}

func (m *MockListRepo) GetActive() (*List, error) {
	return nil, nil
}

func (m *MockListRepo) SetActive(id int, name string) (*List, error) {
	return nil, nil
}

func (m *MockListRepo) GetId(name string) (int, error) {
	return 0, nil
}
