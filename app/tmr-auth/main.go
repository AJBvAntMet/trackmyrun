package main

import (
	"fmt"
	"net/http"
	"os"

	chiprometheus "github.com/jamscloud/chi-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vantmet/trackmyrun/pkg/auth"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

var Version string

var appVersion prometheus.Gauge

func main() {
	appVersion := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "app_version_info",
		Help:        "App info at buildtime",
		ConstLabels: prometheus.Labels{"version": Version},
	})
	prometheus.Register(appVersion)
	appVersion.Set(1)

	cognitoClient := auth.Init()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger, middleware.WithValue("CognitoClient", cognitoClient))
	r.Use(middleware.Heartbeat("/ping"))
	m := chiprometheus.NewMiddleware("serviceName")
	r.Use(m)
	r.Use(middleware.Recoverer)

	r.Handle("/metrics", promhttp.Handler())
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// r.Post("/signup", signUp)

	// r.Post("/signin", signIn)

	// r.Get("/verify", verifyToken)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("starting server on port %s!", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
