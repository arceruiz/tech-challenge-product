package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"tech-challenge-product/internal/canonical"
	"tech-challenge-product/internal/channels/rest"
	"tech-challenge-product/internal/mocks"
	"tech-challenge-product/internal/service"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterGroup(t *testing.T) {
	endpoint := "/product"

	type Given struct {
		group          *echo.Group
		paymenyService service.ProductService
	}
	type Expected struct {
		err        assert.ErrorAssertionFunc
		statusCode int
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given valid group, should register endpoints successfully": {
			given: Given{
				group:          echo.New().Group("/product"),
				paymenyService: &mocks.ProductServiceMock{},
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusNotFound,
			},
		},
	}

	for _, tc := range tests {
		p := rest.NewProductChannel(tc.given.paymenyService)
		p.RegisterGroup(tc.given.group)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, endpoint+"/123", nil)
		e := echo.New()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues("123")

		e.ServeHTTP(rec, req)
		statusCode := rec.Result().StatusCode

		assert.Equal(t, tc.expected.statusCode, statusCode)
	}
}

func TestCreate(t *testing.T) {
	endpoint := "/product"

	type Given struct {
		request        *http.Request
		paymenyService service.ProductService
	}
	type Expected struct {
		err        assert.ErrorAssertionFunc
		statusCode int
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given normal json income must process normally": {
			given: Given{
				request:        createJsonRequest(http.MethodPost, endpoint, rest.ProductRequest{}),
				paymenyService: mockProductServiceForCreate(canonical.Product{}, canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusCreated,
			},
		},
		"given wrong format must return error": {
			given: Given{
				request:        createRequest(http.MethodPost, endpoint),
				paymenyService: mockProductServiceForCreate(canonical.Product{}, canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusBadRequest,
			},
		},
		"given invalid data, must return bad request": {
			given: Given{
				request:        createRequest(http.MethodPost, endpoint),
				paymenyService: mockProductServiceForCreate(canonical.Product{}, canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()
		err := rest.NewProductChannel(tc.given.paymenyService).Add(echo.New().NewContext(tc.given.request, rec))
		statusCode := rec.Result().StatusCode

		assert.Equal(t, tc.expected.statusCode, statusCode)

		tc.expected.err(t, err)
	}
}

func TestUpdate(t *testing.T) {
	endpoint := "/product"

	type Given struct {
		request        *http.Request
		pathParamID    string
		paymenyService service.ProductService
	}
	type Expected struct {
		err        assert.ErrorAssertionFunc
		statusCode int
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given normal json income must process normally": {
			given: Given{
				pathParamID:    "valid_ID",
				request:        createJsonRequest(http.MethodPost, endpoint, rest.ProductRequest{}),
				paymenyService: mockProductServiceForUpdate("valid_ID", canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusOK,
			},
		},
		"given wrong format must return error": {
			given: Given{
				pathParamID:    "valid_ID",
				request:        createRequest(http.MethodPost, endpoint),
				paymenyService: mockProductServiceForUpdate("valid_ID", canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusBadRequest,
			},
		},
		"given invalid data, must return bad request": {
			given: Given{
				pathParamID:    "invalid_ID",
				request:        createRequest(http.MethodPost, endpoint),
				paymenyService: mockProductServiceForUpdate("valid_ID", canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()
		e := echo.New().NewContext(tc.given.request, rec)
		e.SetPath("/:id")
		e.SetParamNames("id")
		e.SetParamValues(tc.given.pathParamID)
		err := rest.NewProductChannel(tc.given.paymenyService).Update(e)
		statusCode := rec.Result().StatusCode

		assert.Equal(t, tc.expected.statusCode, statusCode)

		tc.expected.err(t, err)
	}
}

func TestRemove(t *testing.T) {
	endpoint := "/product"

	type Given struct {
		request        *http.Request
		pathParamID    string
		paymenyService service.ProductService
	}
	type Expected struct {
		err        assert.ErrorAssertionFunc
		statusCode int
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given normal json income must process normally": {
			given: Given{
				pathParamID:    "valid_ID",
				request:        createRequest(http.MethodPost, endpoint),
				paymenyService: mockProductServiceForRemove("valid_ID", canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusOK,
			},
		},
		"given wrong format must return error": {
			given: Given{
				pathParamID:    "invalid_ID",
				request:        createRequest(http.MethodPost, endpoint),
				paymenyService: mockProductServiceForRemove("valid_ID", canonical.Product{}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusNotFound,
			},
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()
		e := echo.New().NewContext(tc.given.request, rec)
		e.SetPath("/:id")
		e.SetParamNames("id")
		e.SetParamValues(tc.given.pathParamID)
		err := rest.NewProductChannel(tc.given.paymenyService).Remove(e)
		statusCode := rec.Result().StatusCode

		assert.Equal(t, tc.expected.statusCode, statusCode)

		tc.expected.err(t, err)
	}
}

func TestGet(t *testing.T) {
	endpoint := "/product/"

	type Given struct {
		request        *http.Request
		pathParamKey   string
		pathParamValue string
		paymenyService service.ProductService
	}
	type Expected struct {
		err        assert.ErrorAssertionFunc
		statusCode int
	}
	tests := map[string]struct {
		given    Given
		expected Expected
	}{
		"given clean request returns valid product and status 200": {
			given: Given{
				request: createRequest(http.MethodGet, endpoint),
				paymenyService: mockProductServiceForGetAll("1234", []canonical.Product{{
					ID: "1234",
				}}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusOK,
			},
		},
		"given valid id returns valid product and status 200": {
			given: Given{
				request:        createRequest(http.MethodGet, endpoint),
				pathParamKey:   "id",
				pathParamValue: "1234",
				paymenyService: mockProductServiceForGetByID("1234", &canonical.Product{
					ID: "1234",
				}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusOK,
			},
		},
		"given valid category returns valid product and status 200": {
			given: Given{
				request:        createRequest(http.MethodGet, endpoint),
				pathParamKey:   "category",
				pathParamValue: "valid_category",
				paymenyService: mockProductServiceForGetByCategory("valid_category", []canonical.Product{{
					ID: "1234",
				}}),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusOK,
			},
		},
		"given empty id returns no product and status 400": {
			given: Given{
				request:        createRequest(http.MethodGet, endpoint),
				paymenyService: mockProductServiceForGetAll_error("1234", nil),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusInternalServerError,
			},
		},
		"given invalic id returns no product and status 404": {
			given: Given{
				request:        createRequest(http.MethodGet, endpoint),
				paymenyService: mockProductServiceForGetAll("1234", nil),
			},
			expected: Expected{
				err:        assert.NoError,
				statusCode: http.StatusNotFound,
			},
		},
	}

	for _, tc := range tests {
		rec := httptest.NewRecorder()
		e := echo.New().NewContext(tc.given.request, rec)

		if tc.given.pathParamKey != "" {
			e.QueryParams().Add(tc.given.pathParamKey, tc.given.pathParamValue)
		}
		err := rest.NewProductChannel(tc.given.paymenyService).Get(e)
		statusCode := rec.Result().StatusCode

		assert.Equal(t, tc.expected.statusCode, statusCode)

		tc.expected.err(t, err)
	}
}

func mockProductServiceForRemove(id string, productReturned canonical.Product) *mocks.ProductServiceMock {
	mockProductSvc := new(mocks.ProductServiceMock)

	mockProductSvc.
		On("Remove", mock.Anything, id).
		Return(nil)

	mockProductSvc.
		On("Remove", mock.Anything, "invalid_ID").
		Return(errors.New(""))

	return mockProductSvc
}

func mockProductServiceForUpdate(id string, productReturned canonical.Product) *mocks.ProductServiceMock {
	mockProductSvc := new(mocks.ProductServiceMock)

	mockProductSvc.
		On("Update", mock.Anything, id, productReturned).
		Return(nil)

	mockProductSvc.
		On("Update", mock.Anything, "invalid_ID", productReturned).
		Return(errors.New(""))

	return mockProductSvc
}

func mockProductServiceForGetByCategory(category string, productReturned []canonical.Product) *mocks.ProductServiceMock {
	mockProductSvc := new(mocks.ProductServiceMock)

	mockProductSvc.
		On("GetByCategory", mock.Anything, category).
		Return(productReturned, nil)

	return mockProductSvc
}

func mockProductServiceForGetByID(productID string, productReturned *canonical.Product) *mocks.ProductServiceMock {
	mockProductSvc := new(mocks.ProductServiceMock)

	mockProductSvc.
		On("GetByID", mock.Anything, productID).
		Return(productReturned, nil)

	return mockProductSvc
}

func mockProductServiceForGetAll(productID string, productReturned []canonical.Product) *mocks.ProductServiceMock {
	mockProductSvc := new(mocks.ProductServiceMock)

	mockProductSvc.
		On("GetAll", mock.Anything).
		Return(productReturned, nil)

	return mockProductSvc
}

func mockProductServiceForGetAll_error(productID string, productReturned []canonical.Product) *mocks.ProductServiceMock {
	mockProductSvc := new(mocks.ProductServiceMock)

	mockProductSvc.
		On("GetAll", mock.Anything).
		Return(productReturned, errors.New(""))

	return mockProductSvc
}

func mockProductServiceForCreate(productReceived, productReturned canonical.Product) *mocks.ProductServiceMock {
	mockProductSvc := new(mocks.ProductServiceMock)
	mockProductSvc.On("Create", mock.Anything, &productReceived).Return(&productReturned, nil)
	mockProductSvc.On("Create", mock.Anything, &canonical.Product{
		ID:          "",
		Name:        "",
		Description: "",
		Price:       0,
		Category:    "",
		Status:      0,
		ImagePath:   "",
	}).Return(&productReturned, errors.New(""))
	return mockProductSvc
}

func createRequest(method, endpoint string) *http.Request {
	req := createJsonRequest(method, endpoint, nil)
	req.Header.Del("Content-Type")
	return req
}

func createJsonRequest(method, endpoint string, request interface{}) *http.Request {
	json, _ := json.Marshal(request)
	req := httptest.NewRequest(method, endpoint, bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	return req
}
