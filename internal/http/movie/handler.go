package movie

import (
	"encoding/json"
	"github.com/vageeshabr/cinema-service/internal/services"
	"net/http"
	"strconv"
)

type Handler struct {
	svc services.Movie
}

func New(svc services.Movie) *Handler {
	return &Handler{svc: svc}
}

type listResponse struct {
	Status string                 `json:"status"`
	Code   int                    `json:"code"`
	Data   map[string]interface{} `json:"data"`
	Meta   metaData               `json:"metaData"`
}

type metaData struct {
	Count int `json:"count"`
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Query().Get("page")
	page, err := strconv.Atoi(p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	movies, err := h.svc.Get(r.Context(), page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	count, err := h.svc.Count(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	res, _ := json.Marshal(listResponse{
		Status: "SUCCESS",
		Code:   200,
		Data:   map[string]interface{}{"movies": movies},
		Meta: metaData{
			Count: count,
		},
	})

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
