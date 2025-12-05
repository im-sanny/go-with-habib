package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsStr, 10, 32)
	limit, _ := strconv.ParseInt(limitAsStr, 10, 32)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	prdCh := make(chan []*domain.Product)
	go func() {
		productList, err := h.svc.List(page, limit)
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		prdCh <- productList
	}()

	ch := make(chan int64)
	go func() {
		cnt, err := h.svc.Count()
		if err != nil {
			util.SendError(w, http.StatusInternalServerError, "Internal server error")
			return
		}
		ch <- cnt
	}()

	productList := <-prdCh
	totalCount := <-ch

	util.SendPage(w, productList, page, limit, totalCount)
}
