package model

type MockBookModel struct {
	books []Book
}

func (m *MockBookModel) Get() []Book {
	return m.books
}

func (m *MockBookModel) Insert(book Book) error {
	m.books = append(m.books, book)
	return nil
}

func NewMockBookModel() *MockBookModel {
	return &MockBookModel{
		books: []Book{},
	}
}
