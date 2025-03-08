package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"recipes.krogowski.dev/internal/consts"
	"recipes.krogowski.dev/internal/core"
	"recipes.krogowski.dev/internal/middleware"
	"recipes.krogowski.dev/internal/repository"
	"recipes.krogowski.dev/internal/validator"
)

type recipeCreateHandler struct {
	recipes         repository.RecipeRepository
	ingredients     repository.IngredientRepository
	ingredientsList repository.IngredientsListRepository
	units           repository.UnitRepository
	requestHandler
}

func NewRecipeCreateHandler(
	env core.Env,
	recipes repository.RecipeRepository,
	ingredientsList repository.IngredientsListRepository,
	ingredients repository.IngredientRepository,
	units repository.UnitRepository,
) recipeCreateHandler {
	return recipeCreateHandler{
		recipes:         recipes,
		ingredients:     ingredients,
		ingredientsList: ingredientsList,
		units:           units,
		requestHandler:  requestHandler{Env: env},
	}
}

type recipieCreateForm struct {
	Title         string
	Description   string
	Ingredients   []string
	Units         []string
	Amount        []string
	ThumbnailFile multipart.File
	validator.Validator
}

// 3 MB
const MAX_UPLOAD_SIZE = 3 << 20

func (h *recipeCreateHandler) post(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	form := recipieCreateForm{}

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		form.AddFieldError("thumbnailFile", validator.FieldErr.ErrFileToBig())
	}

	form.Title = r.PostForm.Get("title")
	form.Description = r.PostForm.Get("description")
	form.Units = r.PostForm["unit_id"]
	form.Amount = r.PostForm["amount"]
	form.Ingredients = r.PostForm["ingredient_id"]

	file, fileHeader, err := r.FormFile("thumbnailFile")
	if err != nil {
		// cannot retrieve file
		h.serverError(w, r, err)
		return
	}
	defer file.Close()

	form.ThumbnailFile = file

	form.CheckField(validator.FileNotBlank(*fileHeader), "thumbnailFile", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.FileTypeNotAllowed(*fileHeader), "thumbnailFile", validator.FieldErr.ErrFileNotAllowed("png, jpeg"))
	form.CheckField(validator.NotBlank(form.Title), "title", validator.FieldErr.ErrNotBlank())
	form.CheckField(validator.NotBlank(form.Description), "description", validator.FieldErr.ErrNotBlank())

	for i, _ := range form.Ingredients {
		unit := form.Units[i]
		ingredient := form.Ingredients[i]
		amount := form.Amount[i]

		index := fmt.Sprintf("i-%d", i)
		form.CheckField(validator.NotBlank(amount), index, validator.FieldErr.ErrNotBlank())
		form.CheckField(validator.NotBlank(ingredient), index, validator.FieldErr.ErrNotBlank())
		form.CheckField(validator.NotBlank(unit), index, validator.FieldErr.ErrNotBlank())
	}

	if !form.IsValid() {
		data := h.Tmpl.NewData(r)

		ingredients, err := h.ingredients.List()
		if err != nil {
			h.serverError(w, r, err)
			return
		}

		units, err := h.units.List()
		if err != nil {
			h.serverError(w, r, err)
			return
		}

		data.Ingredients = ingredients
		data.Units = units
		data.Form = form

		h.render(w, r, http.StatusUnprocessableEntity, "recipeCreate.tmpl", data)
		return
	}

	filename := fmt.Sprintf("%s_%d_%s", form.Title, time.Now().UnixNano(), fileHeader.Filename)
	uploadDir := "./files"

	filepath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(filepath)
	if err != nil {
		// cannot save file
		h.serverError(w, r, err)
		return
	}
	defer dst.Close()

	filesize, err := io.Copy(dst, file)
	if err != nil {
		// cannot copy file
		h.serverError(w, r, err)
		return
	}

	h.Logger.Info("Saved path", "filename", filename, "filepath", filepath, "fileSize", filesize)

	userId := h.Session.GetUserId(r)
	id, err := h.recipes.Insert(form.Title, form.Description, userId, filename, filepath)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	err = h.ingredientsList.Insert(form.Ingredients, form.Units, form.Amount, id)

	if err != nil {
		h.serverError(w, r, err)
		return
	}

	h.Session.SetFlashMsg(r, consts.MsgRecipeCreated)

	pagePath := fmt.Sprintf("/recipes/%d", id)
	http.Redirect(w, r, pagePath, http.StatusSeeOther)
}

func (h *recipeCreateHandler) get(w http.ResponseWriter, r *http.Request) {
	units, err := h.units.List()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	ingredients, err := h.ingredients.List()
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	data := h.Tmpl.NewData(r)
	data.Form = recipieCreateForm{}

	data.Units = units
	data.Ingredients = ingredients

	h.render(w, r, http.StatusOK, "recipeCreate.tmpl", data)
}

func (h *recipeCreateHandler) ingredientAutocomplete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		Query string `json:"query"`
	}

	if err := decoder.Decode(&data); err != nil {
		h.serverError(w, r, err)
		return
	}

	ingredients, err := h.ingredients.Search(data.Query)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	json, err := json.Marshal(ingredients)
	if err != nil {
		h.serverError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (h *recipeCreateHandler) RegisterRoute(mux *http.ServeMux, midw *middleware.Midw) {
	mux.Handle("GET /recipes/create", midw.Protected.ThenFunc(h.get))
	mux.Handle("POST /recipes/create", midw.Protected.ThenFunc(h.post))

	// TODO add proper middlewares
	mux.HandleFunc("POST /ingredients-autocomplete", h.ingredientAutocomplete)
}
