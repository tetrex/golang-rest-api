package book

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	e "golang-rest-api/api/resource/common/err"
)

type API struct {
	repository *Repository
	validator  *validator.Validate
}

func New(db *gorm.DB, v *validator.Validate) *API {
	return &API{
		repository: NewRepository(db),
		validator:  v,
	}
}

// List godoc
//
//	@summary		List books
//	@description	List books
//	@tags			books
//	@accept			json
//	@produce		json
//	@success		200	{array}		DTO
//	@failure		500	{object}	err.Error
//	@router			/books [get]
func (a *API) List(c echo.Context) error {
	books, err := a.repository.List()
	if err != nil {
		return c.String(http.StatusBadRequest, string(e.RespDBDataAccessFailure))
	}

	if len(books) == 0 {
		return c.String(http.StatusNoContent, string(e.RespNoDataFound))
	}

	u := books.ToDto()
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(u)

}

// Create godoc
//
//	@summary		Create book
//	@description	Create book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			body	body	Form	true	"Book form"
//	@success		201
//	@failure		400	{object}	err.Error
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/books [post]
func (a *API) Create(c echo.Context) error {
	form := new(Form)
	if err := c.Bind(form); err != nil {
		return c.String(http.StatusBadRequest, string(e.RespJSONDecodeFailure))
	}

	newBook := form.ToModel()
	newBook.ID = uuid.New()

	_, err := a.repository.Create(newBook)
	if err != nil {
		return c.String(http.StatusBadRequest, string(e.RespDBDataInsertFailure))
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return c.String(http.StatusCreated, newBook.ID.String())
}

// Read godoc
//
//	@summary		Read book
//	@description	Read book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id	path		string	true	"Book ID"
//	@success		200	{object}	DTO
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/books/{id} [get]
func (a *API) Read(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, string(e.RespInvalidURLParamID))
	}
	book, err := a.repository.Read(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, string(e.RespNoDataFound))
		}

		return c.String(http.StatusBadRequest, string(e.RespDBDataAccessFailure))
	}

	dto := book.ToDto()

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(dto)
}

// Update godoc
//
//	@summary		Update book
//	@description	Update book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id		path	string	true	"Book ID"
//	@param			body	body	Form	true	"Book form"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/books/{id} [put]
func (a *API) Update(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, string(e.RespInvalidURLParamID))
	}

	form := &Form{}
	if err := json.NewDecoder(c.Request().Body).Decode(form); err != nil {
		return c.String(http.StatusBadRequest, string(e.RespJSONDecodeFailure))
	}

	book := form.ToModel()
	book.ID = id

	rows, err := a.repository.Update(book)
	if err != nil {
		return c.String(http.StatusBadRequest, string(e.RespDBDataUpdateFailure))
	}
	if rows == 0 {
		return c.String(http.StatusBadRequest, string(e.RespNoDataFound))
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder((c.Response())).Encode(book)
}

// Delete godoc
//
//	@summary		Delete book
//	@description	Delete book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			id	path	string	true	"Book ID"
//	@success		200
//	@failure		400	{object}	err.Error
//	@failure		404
//	@failure		500	{object}	err.Error
//	@router			/books/{id} [delete]
func (a *API) Delete(c echo.Context) error {

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, string(e.RespInvalidURLParamID))
	}

	rows, err := a.repository.Delete(id)
	if err != nil {
		return c.String(http.StatusBadRequest, string(e.RespDBDataRemoveFailure))
	}
	if rows == 0 {
		return c.String(http.StatusBadRequest, string(e.RespNoDataFound))
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)

	return c.String(http.StatusOK, "ok")
}
