package lists

import "fmt"

type MockRepo struct {
	err   error
	count int
	lists []List

	addName  string
	deleteId int

	countCalled     bool
	getAllCalled    bool
	addCalled       bool
	deleteCalled    bool
	getActiveCalled bool
	setActiveCalled bool
	getIdCalled     bool
}

func (m *MockRepo) Count() (int, error) {
	m.countCalled = true
	return m.count, m.err
}

func (m *MockRepo) GetAll() ([]List, error) {
	m.getAllCalled = true
	return m.lists, m.err
}

func (m *MockRepo) Add(name string) (int, error) {
	m.addCalled = true
	m.addName = name
	return len(m.lists) + 1, m.err
}

func (m *MockRepo) Delete(id int) error {
	m.deleteCalled = true
	m.deleteId = id
	return nil
}

func (m *MockRepo) GetActive() (*List, error) {
	m.getActiveCalled = true
	return nil, nil
}

func (m *MockRepo) SetActive(id int, name string) (*List, error) {
	m.setActiveCalled = true
	return nil, nil
}

func (m *MockRepo) GetId(name string) (int, error) {
	m.getIdCalled = true
	for _, l := range m.lists {
		if l.Name == name {
			return l.Id, nil
		}
	}
	return 0, fmt.Errorf("list not found: %s", name)
}

func (m *MockRepo) GetFirst() (*List, error) {
	return &List{Id: m.lists[0].Id, Name: m.lists[0].Name}, nil
}
