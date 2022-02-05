package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

func (rs *RunnerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.Body)

	switch r.Method {
	case http.MethodPost:
		rs.processRun(w, r)
	case http.MethodGet:
		rs.showRuns(w, r)
	}

}

func (rs *RunnerServer) processRun(w http.ResponseWriter, r *http.Request) {
	var run Run
	err := json.NewDecoder(r.Body).Decode(&run)
	log.Printf("%v", run)
	if err != nil {
		log.Printf("%q", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rs.store.RecordRun(run)
	w.WriteHeader(http.StatusAccepted)
}

func (rs *RunnerServer) showRuns(w http.ResponseWriter, r *http.Request) {
	runs := rs.store.GetRunnerRuns()
	data := struct {
		PageTitle string
		Runs      []Run
	}{
		PageTitle: "My Latest Runs",
		Runs:      runs,
	}
	f := filepath.Join("html", "GetLatest.html")
	t, err := template.ParseFiles(f)

	if err == nil {
		t.Execute(w, data)
	} else {
		fmt.Printf("Template error: %q", err)
	}
}

type RunnerStore interface {
	GetRunnerRuns() []Run
	RecordRun(r Run)
}

type RunnerServer struct {
	store RunnerStore
}
