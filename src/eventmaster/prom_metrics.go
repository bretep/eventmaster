package main

import (
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpReqLatencies = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "eventmaster",
		Subsystem: "http_server",
		Name:      "request_latency",
		Help:      "Latency of http requests grouped by req path",
	}, []string{"path"})

	httpReqCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "eventmaster",
		Subsystem: "http_server",
		Name:      "request_total",
		Help:      "The count of http requests received grouped by req path",
	}, []string{"path"})

	httpRespCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "eventmaster",
		Subsystem: "http_server",
		Name:      "response_total",
		Help:      "The count of http responses issued classified by code and api endpoint",
	}, []string{"path", "code"})

	grpcReqLatencies = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "eventmaster",
		Subsystem: "grpc_server",
		Name:      "request_latency",
		Help:      "Latency of grpc requests grouped by method name",
	}, []string{"method"})

	grpcReqCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "eventmaster",
		Subsystem: "grpc_server",
		Name:      "request_total",
		Help:      "The count of grpc requests received grouped by method name",
	}, []string{"method"})

	grpcRespCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "eventmaster",
		Subsystem: "grpc_server",
		Name:      "response_total",
		Help:      "The count of grpc responses issued classified by method name and success",
	}, []string{"method", "success"})

	rsyslogReqLatencies = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "eventmaster",
		Subsystem: "rsyslog_server",
		Name:      "request_latency",
		Help:      "Latency of rsyslog requests",
	}, []string{})

	rsyslogReqCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "eventmaster",
		Subsystem: "rsyslog_server",
		Name:      "request_total",
		Help:      "The count of rsyslog requests received",
	}, []string{})

	eventStoreTimer = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "eventmaster",
		Subsystem: "event_store",
		Name:      "method_time",
		Help:      "Time of event store methods by method name",
	}, []string{"method"})

	eventStoreDbErrCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "eventmaster",
		Subsystem: "event_store",
		Name:      "db_error",
		Help:      "The count of db errors by db name and type of operation",
	}, []string{"db_name", "operation"})
)

func GetPromHandler() *httprouter.Router {
	r := httprouter.New()
	r.Handler("GET", "/metrics", promhttp.Handler())
	return r
}

func registerPromMetrics() error {
	regErr := prometheus.Register(httpReqLatencies)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			httpReqLatencies = c.ExistingCollector.(*prometheus.HistogramVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(httpReqCounter)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			httpReqCounter = c.ExistingCollector.(*prometheus.CounterVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(httpRespCounter)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			httpRespCounter = c.ExistingCollector.(*prometheus.CounterVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(grpcReqLatencies)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			grpcReqLatencies = c.ExistingCollector.(*prometheus.HistogramVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(grpcReqCounter)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			grpcReqCounter = c.ExistingCollector.(*prometheus.CounterVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(grpcRespCounter)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			grpcRespCounter = c.ExistingCollector.(*prometheus.CounterVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(rsyslogReqLatencies)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			rsyslogReqLatencies = c.ExistingCollector.(*prometheus.HistogramVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(rsyslogReqCounter)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			rsyslogReqCounter = c.ExistingCollector.(*prometheus.CounterVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(eventStoreTimer)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			eventStoreTimer = c.ExistingCollector.(*prometheus.HistogramVec)
		} else {
			return regErr
		}
	}

	regErr = prometheus.Register(eventStoreDbErrCounter)
	if regErr != nil {
		if c, ok := regErr.(prometheus.AlreadyRegisteredError); ok {
			eventStoreDbErrCounter = c.ExistingCollector.(*prometheus.CounterVec)
		} else {
			return regErr
		}
	}
	return nil
}

func trackTime(start time.Time) float64 {
	return float64(time.Since(start).Nanoseconds()) / float64(1000000)
}
