package httpserver

import (
	"kopuro/service"
	"kopuro/view"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	httpPort        string
	jsonFileService service.JsonFileService
}

func NewServer(httpPort string, jsonFileService service.JsonFileService) Server {
	return Server{httpPort, jsonFileService}
}

func (s Server) Start() {
	router := httprouter.New()
	router.POST("/json/write", s.writeJSONFile)
	router.GET("/json/read", s.readJSONFile)

	http.ListenAndServe(":" + s.httpPort, router)
}

func (s Server) writeJSONFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	var request view.WriteJsonFileRequest
	err := decoder.Decode(&request)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{
			"error": "Cannot parse request into JSON - %s"
		}`, err.Error())))
		return
	}

	if len(request.Filename) <= 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
			"error": "Filename cannot be empty"
		}`))
		return
	}

	err = s.jsonFileService.WriteJSONFile(request)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{
			"error": "%s"
		}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s Server) readJSONFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filename := r.URL.Query().Get("filename")

	if len(filename) <= 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{
			"error": "Filename cannot be empty"
		}`))
		return
	}

	res, err := s.jsonFileService.ReadJSONFile(filename)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{
			"error": "%s"
		}`, err.Error())))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}