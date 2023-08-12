package models

import (
	"mvc/pkg/types"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCheckinDeniedCheckouts(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()

	// Define the expected rows for the query
	rows := sqlmock.NewRows([]string{"bookid", "bookname", "author", "copiesAvailable"}).
		AddRow(1, "a", "A", 4).
		AddRow(2, "b", "B", 0).
		AddRow(3, "c", "C", 12).
		AddRow(4, "d", "D", 0)

	// Set the expectation for the query
	mock.ExpectQuery("SELECT \\* FROM books WHERE copiesAvailable > 0").
		WillReturnRows(rows)

	// Call the function being tested
	var books types.Books
	books, err = FetchBooks(db)
	if err != nil {
		t.Fatalf("Error getting checkin denied checkouts: %v", err)
	}

	// Assert the results
	assert.Equal(t, 2, len(books.Books))

	assert.Equal(t, 1, books.Books[0].Bookid)
	assert.Equal(t, "a", books.Books[0].Bookname)
	assert.Equal(t, "A", books.Books[0].Author)
	assert.Equal(t, 4, books.Books[0].Quantity)

	assert.Equal(t, 3, books.Books[1].Bookid)
	assert.Equal(t, "c", books.Books[1].Bookname)
	assert.Equal(t, "C", books.Books[1].Author)
	assert.Equal(t, 12, books.Books[1].Quantity)

	// Check if the mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
