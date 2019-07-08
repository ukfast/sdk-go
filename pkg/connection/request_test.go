package connection

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewValidationError_PopulatesMessage(t *testing.T) {
	err := NewValidationError("test error 1")

	assert.Equal(t, "test error 1", err.Message)
}

func TestAPIRequestBodyDefaultValidator(t *testing.T) {
	type teststruct struct {
		TestProperty string `validate:"required"`
	}

	t.Run("Valid_NoError", func(t *testing.T) {
		ts := teststruct{
			TestProperty: "testvalue1",
		}
		v := APIRequestBodyDefaultValidator{}

		err := v.Validate(ts)

		assert.Nil(t, err)
	})

	t.Run("Invalid_ReturnsValidationError", func(t *testing.T) {
		ts := teststruct{}
		v := APIRequestBodyDefaultValidator{}

		err := v.Validate(ts)

		assert.NotNil(t, err)
		assert.IsType(t, &ValidationError{}, err)
	})
}

func TestAPIRequestParameters_WithPagination_SetsPagination(t *testing.T) {
	p := APIRequestPagination{
		Page: 5,
	}
	params := APIRequestParameters{}

	params.WithPagination(p)

	assert.Equal(t, p, params.Pagination)
}

func TestAPIRequestParameters_WithSorting_SetsSorting(t *testing.T) {
	s := APIRequestSorting{
		Property: "testproperty1",
	}
	params := APIRequestParameters{}

	params.WithSorting(s)

	assert.Equal(t, s, params.Sorting)
}

func TestAPIRequestParameters_WithFilter_AddsFilter(t *testing.T) {
	f := APIRequestFiltering{
		Property: "testproperty1",
		Operator: EQOperator,
		Value:    []string{"testvalue1"},
	}
	params := APIRequestParameters{}

	params.WithFilter(f)

	assert.Len(t, params.Filtering, 1)
	assert.Equal(t, f, params.Filtering[0])
}

func TestAPIRequestParameters_WithFilters_AddsFilters(t *testing.T) {
	f1 := APIRequestFiltering{
		Property: "testproperty1",
		Operator: EQOperator,
		Value:    []string{"testvalue1"},
	}
	f2 := APIRequestFiltering{
		Property: "testproperty2",
		Operator: EQOperator,
		Value:    []string{"testvalue2"},
	}

	f := []APIRequestFiltering{f1, f2}
	params := APIRequestParameters{}

	params.WithFilters(f)

	assert.Len(t, params.Filtering, 2)
	assert.Equal(t, f1, params.Filtering[0])
	assert.Equal(t, f2, params.Filtering[1])
}

func TestAPIRequestParameters_Copy_ReturnsCopy(t *testing.T) {
	params := APIRequestParameters{
		Filtering: []APIRequestFiltering{
			APIRequestFiltering{
				Property: "testproperty1",
			},
			APIRequestFiltering{
				Operator: EQOperator,
				Value:    []string{"testproperty2valuea", "testproperty2valueb"},
			},
		},
		Sorting: APIRequestSorting{
			Property: "testproperty1",
		},
		Pagination: APIRequestPagination{
			Page: 4,
		},
	}

	newParams := params.Copy()

	assert.Len(t, newParams.Filtering, 2)
	assert.Equal(t, "testproperty1", newParams.Filtering[0].Property)
	assert.Equal(t, EQOperator, newParams.Filtering[1].Operator)
	assert.Equal(t, "testproperty2valueb", newParams.Filtering[1].Value[1])
	assert.Equal(t, "testproperty1", newParams.Sorting.Property)
	assert.Equal(t, 4, newParams.Pagination.Page)
}

func TestParseOperator(t *testing.T) {
	type testoperator struct {
		Operator         string
		ExpectedOperator APIRequestFilteringOperator
	}
	operators := []testoperator{
		testoperator{
			Operator:         "EQ",
			ExpectedOperator: EQOperator,
		},
		testoperator{
			Operator:         "LK",
			ExpectedOperator: LKOperator,
		},
		testoperator{
			Operator:         "GT",
			ExpectedOperator: GTOperator,
		},
		testoperator{
			Operator:         "LT",
			ExpectedOperator: LTOperator,
		},
		testoperator{
			Operator:         "IN",
			ExpectedOperator: INOperator,
		},
		testoperator{
			Operator:         "NEQ",
			ExpectedOperator: NEQOperator,
		},
		testoperator{
			Operator:         "NIN",
			ExpectedOperator: NINOperator,
		},
		testoperator{
			Operator:         "NLK",
			ExpectedOperator: NLKOperator,
		},
	}

	for _, operator := range operators {
		t.Run(fmt.Sprintf("%s_Exact_Parses", operator.Operator), func(t *testing.T) {
			parsedOperator, err := ParseOperator(operator.Operator)

			assert.Nil(t, err)
			assert.Equal(t, operator.ExpectedOperator, parsedOperator)
		})
	}

	t.Run("MixedCase_Parses", func(t *testing.T) {
		operator, err := ParseOperator("gT")

		assert.Nil(t, err)
		assert.Equal(t, GTOperator, operator)
	})

	t.Run("Invalid_Error", func(t *testing.T) {
		_, err := ParseOperator("invalidoperator")

		assert.NotNil(t, err)
	})
}

func TestAPIRequestFilteringOperator_String(t *testing.T) {
	t.Run("Valid_String", func(t *testing.T) {
		o := GTOperator

		s := o.String()

		assert.Equal(t, "gt", s)
	})

	t.Run("Invalid_Unknown", func(t *testing.T) {
		var o APIRequestFilteringOperator = 999

		s := o.String()

		assert.Equal(t, "Unknown", s)
	})
}

func TestNewPaginatedBase_ReturnsPaginatedBase(t *testing.T) {
	b := NewPaginatedBase(APIRequestParameters{}, APIResponseMetadataPagination{}, func(parameters APIRequestParameters) (Paginated, error) { return nil, nil })

	assert.NotNil(t, b)
}

func TestPaginatedBase_TotalPages_ReturnsTotalPages(t *testing.T) {
	b := PaginatedBase{
		pagination: APIResponseMetadataPagination{
			TotalPages: 10,
		},
	}

	totalPages := b.TotalPages()

	assert.Equal(t, 10, totalPages)
}

func TestPaginatedBase_CurrentPage(t *testing.T) {
	t.Run("GreaterThan0_ReturnsCurrentPage", func(t *testing.T) {
		b := PaginatedBase{
			parameters: APIRequestParameters{
				Pagination: APIRequestPagination{
					Page: 5,
				},
			},
		}

		currentPage := b.CurrentPage()

		assert.Equal(t, 5, currentPage)
	})

	t.Run("LessThan1_Returns1", func(t *testing.T) {
		b := PaginatedBase{
			parameters: APIRequestParameters{
				Pagination: APIRequestPagination{
					Page: 0,
				},
			},
		}

		currentPage := b.CurrentPage()

		assert.Equal(t, 1, currentPage)
	})
}

func TestPaginatedBase_Total_ReturnsTotal(t *testing.T) {
	b := PaginatedBase{
		pagination: APIResponseMetadataPagination{
			Total: 10,
		},
	}

	total := b.Total()

	assert.Equal(t, 10, total)
}

func TestPaginatedBase_First(t *testing.T) {
	b := PaginatedBase{
		getFunc: func(parameters APIRequestParameters) (Paginated, error) {
			assert.Equal(t, 1, parameters.Pagination.Page)

			return &PaginatedBase{
				parameters: parameters,
			}, nil
		},
	}

	firstPage, err := b.First()

	assert.Nil(t, err)
	assert.Equal(t, 1, firstPage.CurrentPage())
}

func TestPaginatedBase_Previous(t *testing.T) {
	t.Run("NotFirstPage_GetsPreviousPage", func(t *testing.T) {
		b := PaginatedBase{
			parameters: APIRequestParameters{
				Pagination: APIRequestPagination{
					Page: 5,
				},
			},
			getFunc: func(parameters APIRequestParameters) (Paginated, error) {
				assert.Equal(t, 4, parameters.Pagination.Page)

				return &PaginatedBase{
					parameters: parameters,
				}, nil
			},
		}

		previousPage, err := b.Previous()

		assert.Nil(t, err)
		assert.Equal(t, 4, previousPage.CurrentPage())
	})

	t.Run("FirstPage_ReturnsNil", func(t *testing.T) {
		b := PaginatedBase{
			parameters: APIRequestParameters{
				Pagination: APIRequestPagination{
					Page: 1,
				},
			},
		}

		previousPage, err := b.Previous()

		assert.Nil(t, err)
		assert.Nil(t, previousPage)
	})
}

func TestPaginatedBase_Next(t *testing.T) {
	t.Run("NotLastPage_GetsNextPage", func(t *testing.T) {
		b := PaginatedBase{
			parameters: APIRequestParameters{
				Pagination: APIRequestPagination{
					Page: 5,
				},
			},
			pagination: APIResponseMetadataPagination{
				TotalPages: 8,
			},
			getFunc: func(parameters APIRequestParameters) (Paginated, error) {
				assert.Equal(t, 6, parameters.Pagination.Page)

				return &PaginatedBase{
					parameters: parameters,
				}, nil
			},
		}

		nextPage, err := b.Next()

		assert.Nil(t, err)
		assert.Equal(t, 6, nextPage.CurrentPage())
	})

	t.Run("LastPage_ReturnsNil", func(t *testing.T) {
		b := PaginatedBase{
			parameters: APIRequestParameters{
				Pagination: APIRequestPagination{
					Page: 5,
				},
			},
			pagination: APIResponseMetadataPagination{
				TotalPages: 5,
			},
		}

		nextPage, err := b.Next()

		assert.Nil(t, err)
		assert.Nil(t, nextPage)
	})
}

func TestPaginatedBase_Last(t *testing.T) {
	b := PaginatedBase{
		pagination: APIResponseMetadataPagination{
			TotalPages: 10,
		},
		getFunc: func(parameters APIRequestParameters) (Paginated, error) {
			assert.Equal(t, 10, parameters.Pagination.Page)

			return &PaginatedBase{
				parameters: parameters,
			}, nil
		},
	}

	lastPage, err := b.Last()

	assert.Nil(t, err)
	assert.Equal(t, 10, lastPage.CurrentPage())
}

func TestNewRequestAll_ReturnsRequestAll(t *testing.T) {
	b := NewRequestAll(func(parameters APIRequestParameters) (Paginated, error) { return nil, nil }, func(Paginated) {})

	assert.NotNil(t, b)
}

func TestRequestAll_Invoke(t *testing.T) {
	t.Run("MultiplePages_ExpectedCalls", func(t *testing.T) {
		totalCalls := 0

		r := RequestAll{
			getFunc: func(parameters APIRequestParameters) (Paginated, error) {
				totalCalls++

				return &PaginatedBase{
					parameters: parameters,
					pagination: APIResponseMetadataPagination{
						TotalPages: 10,
					},
				}, nil
			},
			responseFunc: func(Paginated) {},
		}

		err := r.Invoke(APIRequestParameters{})

		assert.Nil(t, err)
		assert.Equal(t, 10, totalCalls)
	})

	t.Run("Error", func(t *testing.T) {
		totalCalls := 0
		testErr := errors.New("test error")
		r := RequestAll{
			getFunc: func(parameters APIRequestParameters) (Paginated, error) {
				totalCalls++

				return &PaginatedBase{}, testErr
			},
		}

		err := r.Invoke(APIRequestParameters{})

		assert.NotNil(t, err)
		assert.Equal(t, testErr, err)
		assert.Equal(t, 1, totalCalls)
	})
}

func TestInvokeRequestAll_CallsInvoke(t *testing.T) {
	totalCalls := 0
	result := InvokeRequestAll(func(parameters APIRequestParameters) (Paginated, error) {
		totalCalls++

		return &PaginatedBase{
			parameters: parameters,
			pagination: APIResponseMetadataPagination{
				TotalPages: 5,
			},
		}, nil
	}, func(Paginated) {}, APIRequestParameters{})

	assert.Nil(t, result)
	assert.Equal(t, 5, totalCalls)
}
