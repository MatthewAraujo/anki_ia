package anki

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/MatthewAraujo/anki_ia/types"
	"github.com/MatthewAraujo/anki_ia/utils"
	"github.com/gorilla/mux"
)

var logger = utils.NewParentLogger("Rota api/v1/anki")

type Handler struct {
	Service types.AnkiService
}

func NewHandler(Service types.AnkiService) *Handler {
	return &Handler{
		Service: Service,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.CreateAnki).Methods(http.MethodPost)
}

func (h *Handler) CreateAnki(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Creating anki")

	_, err := utils.ParseMultipartForm(r)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	file, name, err := FileUploadHandler(r)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer file.Close()

	payload := types.CreateAnkiPayload{
		File: file,
		Name: name,
	}

	questions, status, err := h.Service.CreateAnki(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, map[string]string{"response": questions})
}

func FileUploadHandler(r *http.Request) (multipart.File, string, error) {
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, "", fmt.Errorf("erro ao obter arquivo do formulário: %w", err)
	}

	return file, header.Filename, nil
}
