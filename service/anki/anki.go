package anki

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/MatthewAraujo/anki_ia/repository"
	"github.com/MatthewAraujo/anki_ia/service/auth"
	"github.com/MatthewAraujo/anki_ia/types"
	"github.com/MatthewAraujo/anki_ia/utils"
	"github.com/gorilla/mux"
)

var logger = utils.NewParentLogger("Rota api/v1/anki")

type Handler struct {
	Service types.AnkiService
	store   repository.Queries
}

func NewHandler(Service types.AnkiService, store repository.Queries) *Handler {
	return &Handler{
		Service: Service,
		store:   store,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", auth.WithJWTAuth(h.CreateAnki, h.store)).Methods(http.MethodPost)
	router.HandleFunc("/{id}", auth.WithJWTAuth(h.GetAnkiById, h.store)).Methods(http.MethodGet)
	router.HandleFunc("/{id}/user", auth.WithJWTAuth(h.GetAnkisByUser, h.store)).Methods(http.MethodGet)
}

func (h *Handler) GetAnkisByUser(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Get Anki By User")
	userID := auth.GetUserIDFromContext(r.Context())
	if userID == 0 {
		logger.LogError(r.URL.Path, fmt.Errorf("user not authenticated"))
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("user not authenticated"))
		return
	}

	userId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error parsing id"))
		return
	}

	var payload types.GetAnkisByUserIDPayload

	payload = types.GetAnkisByUserIDPayload{
		UserID: int32(id),
	}

	questions, status, err := h.Service.GetAnkisByUserID(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, questions)
}

func (h *Handler) GetAnkiById(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Get Anki By ID")
	ankiId := mux.Vars(r)["id"]

	id, err := strconv.Atoi(ankiId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error parsing id"))
		return
	}

	var payload types.GetAnkiByIdPayload

	payload = types.GetAnkiByIdPayload{
		Id: int32(id),
	}

	questions, status, err := h.Service.GetAnkiById(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, questions)

}

func (h *Handler) CreateAnki(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.URL.Path, "Creating anki")

	// Obtém o userID do contexto
	userID := auth.GetUserIDFromContext(r.Context())
	if userID == 0 {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("user not authenticated"))
		return
	}

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
		File:   file,
		Name:   name,
		UserID: userID,
	}

	questions, status, err := h.Service.CreateAnki(&payload)
	if err != nil {
		logger.LogError(r.URL.Path, err)
		utils.WriteError(w, status, err)
		return
	}

	utils.WriteJSON(w, status, questions)
}

func FileUploadHandler(r *http.Request) (multipart.File, string, error) {
	file, header, err := r.FormFile("file")
	if err != nil {
		return nil, "", fmt.Errorf("erro ao obter arquivo do formulário: %w", err)
	}

	return file, header.Filename, nil
}
