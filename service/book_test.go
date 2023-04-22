package service

import (
	"errors"
	"ninth-learn/model"
	"ninth-learn/repository/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			Title:       "The Great Gatsby",
			Author:      "F. Scott Fitzgerald",
			Description: "A novel about the decadence of the Jazz Age",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "The Great Gatsby",
				Author:      "F. Scott Fitzgerald",
				Description: "A novel about the decadence of the Jazz Age",
				CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "The Great Gatsby",
			Author:      "F. Scott Fitzgerald",
			Description: "A novel about the decadence of the Jazz Age",
			CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
		},
	})

	testTable = append(testTable, testCase{
		name:      "unexpected error",
		wantError: true,
		input: model.Book{
			Title:       "To Kill a Mockingbird",
			Author:      "Harper Lee",
			Description: "A novel about racial injustice in the American South",
		},
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid title length",
		wantError: true,
		input: model.Book{
			Title:       "A",
			Author:      "J.D. Salinger",
			Description: "A novel about a disillusioned teenager",
		},
		expectedError: errors.New("invalid title length"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("invalid title length")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid author length",
		wantError: true,
		input: model.Book{
			Title:       "The Catcher in the Rye",
			Author:      "J",
			Description: "A novel about a disillusioned teenager",
		},
		expectedError: errors.New("invalid author length"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("invalid author length")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:      "invalid description length",
		wantError: true,
		input: model.Book{
			Title:       "1984",
			Author:      "George Orwell",
			Description: "A",
		},
		expectedError: errors.New("invalid description length"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(model.Book{}, errors.New("invalid description length")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_GetBookById(t *testing.T) {
	type testCase struct {
		name           string
		id             int64
		wantError      bool
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		id:        1,
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "The Hitchhiker's Guide to the Galaxy",
				Author:      "Douglas Adams",
				Description: "A comedic science fiction series",
				CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "The Hitchhiker's Guide to the Galaxy",
			Author:      "Douglas Adams",
			Description: "A comedic science fiction series",
			CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		id:            2,
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		id:            3,
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBookById(testCase.id)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_GetBooks(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		expectedResult []model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBooks().Return([]model.Book{
				{
					ID:          1,
					Title:       "The Hitchhiker's Guide to the Galaxy",
					Author:      "Douglas Adams",
					Description: "A comedic science fiction series",
					CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				},
				{
					ID:          2,
					Title:       "1984",
					Author:      "George Orwell",
					Description: "A dystopian novel",
					CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				},
			}, nil).Times(1)
		},
		expectedResult: []model.Book{
			{
				ID:          1,
				Title:       "The Hitchhiker's Guide to the Galaxy",
				Author:      "Douglas Adams",
				Description: "A comedic science fiction series",
				CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			},
			{
				ID:          2,
				Title:       "1984",
				Author:      "George Orwell",
				Description: "A dystopian novel",
				CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			},
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBooks().Return([]model.Book{}, errors.New("unexpected error")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBooks()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_UpdateBook(t *testing.T) {
	type testCase struct {
		name           string
		wantError      bool
		input          model.Book
		expectedResult model.Book
		expectedError  error
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	// Test case for updating an existing book
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input: model.Book{
			ID:          1,
			Title:       "The Great Gatsby",
			Author:      "F. Scott Fitzgerald",
			Description: "A novel about the decadence of the Jazz Age",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "The Great Gatsby",
				Author:      "F. Scott Fitzgerald",
				Description: "A novel about the decadence of the Jazz Age",
				CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2023, time.April, 23, 15, 0, 0, 0, time.UTC),
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "The Great Gatsby",
			Author:      "F. Scott Fitzgerald",
			Description: "A novel about the decadence of the Jazz Age",
			CreatedAt:   time.Date(2023, time.April, 22, 14, 0, 0, 0, time.UTC),
			UpdatedAt:   time.Date(2023, time.April, 23, 15, 0, 0, 0, time.UTC),
		},
	})

	// Test case for updating a non-existing book
	testTable = append(testTable, testCase{
		name:      "not found",
		wantError: true,
		input: model.Book{
			ID:          2,
			Title:       "To Kill a Mockingbird",
			Author:      "Harper Lee",
			Description: "A novel about racial injustice in the American South",
		},
		expectedError: errors.New("book not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(model.Book{}, errors.New("book not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.UpdateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_DeleteBook(t *testing.T) {
	type testCase struct {
		name          string
		wantError     bool
		input         int64
		expectedError error
		onBookRepo    func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		input:     1,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(nil).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "unexpected error",
		wantError:     true,
		input:         1,
		expectedError: errors.New("unexpected error"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("unexpected error")).Times(1)
		},
	})

	testTable = append(testTable, testCase{
		name:          "book not found",
		wantError:     true,
		input:         1,
		expectedError: errors.New("book not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(errors.New("book not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			err := service.DeleteBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
