package parser

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// ExtractIDFromRequest извлекает ID из запроса.
func ExtractIDFromRequest(r *http.Request, fieldID string) (int32, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid or missing %s: %s", fieldID, idStr)
	}

	return int32(id), nil
}
