package main

import (
	"net/http"

	"github.com/ahmetilboga2004/go-blog/pkg/utils"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":4000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		utils.Log(utils.ERROR, "Sunucu başlatılırken bir hata oluştu %v", err)
	} else {
		utils.Log(utils.INFO, "Sunucu başlatıldı")
	}
}
