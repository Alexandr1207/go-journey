package httpapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"myfirstgo/week07/day03-middleware/internal/user"
)

type userResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Warning string `json:"warning,omitempty"`
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Handler struct {
	svc *user.Service
}

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type errorResponse struct {
	Error errorBody `json:"error"`
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	err := errorBody{Code: code, Message: message}
	res := errorResponse{Error: err}
	writeJSON(w, status, res)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "invalid JSON body")
		return
	}

	u, err := h.svc.Register(req.Name, req.Email)
	resp := toResponse(u)

	switch {
	case err == nil:
		writeJSON(w, 201, resp)
	case errors.Is(err, user.ErrNotifyFailed):
		resp.Warning = "welcome notification was not sent"
		writeJSON(w, 201, resp)
	case errors.Is(err, user.ErrEmptyName):
		writeError(w, http.StatusBadRequest, "empty_name", "name must not be empty")
	case errors.Is(err, user.ErrMailTaken):
		writeError(w, http.StatusConflict, "email_taken", "email already taken")
	case errors.Is(err, user.ErrInvalidEmail):
		writeError(w, http.StatusBadRequest, "invalid_email", "invalid email address")
	default:
		log.Printf("create user: %v", err)
		writeError(w, http.StatusInternalServerError, "internal", "internal error")
	}
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}

func toResponse(u user.User) userResponse {
	return userResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	u, err := h.svc.Find(email)

	switch {
	case err == nil:
		writeJSON(w, 200, toResponse(u))
	case errors.Is(err, user.ErrNotFound):
		writeError(w, http.StatusNotFound, "not_found", "user not found")
	default:
		log.Printf("get user: %v", err)
		writeError(w, 500, "internal", "internal error")
	}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		lst, err := h.svc.Users()
		if err != nil {
			log.Printf("get users: %v", err)
			writeError(w, 500, "internal", "inernal error")
			return
		}
		res := make([]userResponse, 0, len(lst))

		for _, val := range lst {
			res = append(res, toResponse(val))
		}

		writeJSON(w, http.StatusOK, res)
	})
	mux.HandleFunc("POST /users", h.createUser)
	mux.HandleFunc("GET /users/{email}", h.getUser)

	return Recover(RequestID(Logging(mux)))
}
