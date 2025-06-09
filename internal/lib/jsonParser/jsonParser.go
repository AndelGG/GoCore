package jsonParser

import (
	"awesomeProject/internal/domain"
	"encoding/json"
	"net/http"
)

func JsonParser(r *http.Request) (domain.ServiceMessage, error) {
	var data domain.JsonParser

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return domain.ServiceMessage{}, nil
	}

	resp, err := data.Parse()
	if err != nil {
		return domain.ServiceMessage{}, nil
	}

	return resp, nil
}
