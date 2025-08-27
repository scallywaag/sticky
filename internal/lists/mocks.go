package lists

type MockRepo struct {
	count int
	lists []List
	err   error
}

func (m *MockRepo) Count() (int, error) {
	return m.count, m.err
}

func (m *MockRepo) GetAll() ([]List, error) {
	return m.lists, m.err
}

func (m *MockRepo) Add(name string) (int, error) {
	return 0, nil
}

func (m *MockRepo) Delete(id int) error {
	return nil
}

func (m *MockRepo) GetActive() (*List, error) {
	return nil, nil
}

func (m *MockRepo) SetActive(id int, name string) (*List, error) {
	return nil, nil
}

func (m *MockRepo) GetId(name string) (int, error) {
	return 0, nil
}
