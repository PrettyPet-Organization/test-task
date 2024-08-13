package handler

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
	"small_warehouse/internal/model"

	"github.com/go-chi/chi"
)

type GoodsService interface {
	GetGood(key string) (model.Good, error)
	GetGoods() []model.Good
	CreateOrUpdateGood(t model.Good) error
	GetGoodsCount() int
}

type Handler struct {
	serv  GoodsService
	gtmpl template.Template
}

func NewGoods(serv GoodsService) Handler {
	return Handler{
		serv:  serv,
		gtmpl: *template.Must(template.ParseFiles("internal/views/good.html", "internal/views/goods.html")),
	}
}

func (h *Handler) ViewGood(w http.ResponseWriter, r *http.Request) {
	good_id := chi.URLParam(r, "id")
	if good_id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	model, err := h.serv.GetGood(good_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.gtmpl.ExecuteTemplate(w, "good", model)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		slog.Warn("template execution", slog.Any("failed", err))
		return
	}
}

func (h *Handler) ViewGoods(w http.ResponseWriter, r *http.Request) {
	goods := h.serv.GetGoods()
	if len(goods) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := h.gtmpl.ExecuteTemplate(w, "goods", goods)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetGood(w http.ResponseWriter, r *http.Request) {
	good_id := chi.URLParam(r, "id")
	if good_id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	model, err := h.serv.GetGood(good_id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	serialized, err := json.Marshal(&model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

func (h *Handler) GetGoods(w http.ResponseWriter, r *http.Request) {
	goods := h.serv.GetGoods()

	byt, err := json.Marshal(goods)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(byt)
}

func (h *Handler) GetGoodsCount(w http.ResponseWriter, r *http.Request) {
	v := h.serv.GetGoodsCount()

	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func (h *Handler) CreateRandomGood(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	typ := q.Get("typ")
	category := q.Get("category")

	model := model.NewGood(typ, category)
	err := h.serv.CreateOrUpdateGood(*model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	serialized, err := json.Marshal(&model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(serialized)
}
func (h *Handler) CreateOrUpdateGood(w http.ResponseWriter, r *http.Request) {
	var model model.Good
	err := json.NewDecoder(r.Body).Decode(&model)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.serv.CreateOrUpdateGood(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
