package presentation

import (
	"errors"
	"fmt"
	"net/http"
)


func Start(port uint16) error {
	http.HandleFunc("/", nil)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return errors.New("Error while starting Server")
	}
	return nil
}
