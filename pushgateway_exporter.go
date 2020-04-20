package prometheus

import (
	"github.com/prometheus/client_golang/prometheus/push"
	"go.opencensus.io/stats/view"
	"log"
	"sync/atomic"
	"time"
)

// PushGatewayExporter exporter for the Prometheus's push gateway
type PushGatewayExporter struct {
	service    string
	gatewayURL string
	duration   time.Duration
	prometheus *Exporter
	pusher     *push.Pusher
	stopCh     chan bool
	finishCh   chan bool
	locker     uint32
}

// NewPushGatewayExporter returns the new PushGatewayExporter instance
func NewPushGatewayExporter(service string, gatewayURL string, duration time.Duration) (*PushGatewayExporter, error) {
	prometheus, err := NewExporter(Options{
		Namespace: service,
	})
	if err != nil {
		return nil, err
	}

	pusher := push.New(gatewayURL, service).Gatherer(prometheus.g)
	return &PushGatewayExporter{
		service:    service,
		gatewayURL: gatewayURL,
		duration:   duration,
		prometheus: prometheus,
		pusher:     pusher,
		stopCh:     make(chan bool),
		finishCh:   make(chan bool),
	}, nil
}

// ExportView implements the views interface
// Deprecated: don't need to do anything. prometheus uses the metricexport.Reader interface.
// which is implemented in: collector > metricExporter > go.opencensus.io/metric/metricexport.ReadAndExport
func (p *PushGatewayExporter) ExportView(viewData *view.Data) {
}

// Run start the exporter
func (p *PushGatewayExporter) Run() {
	ticker := time.NewTicker(p.duration)

	go func() {
		for {
			select {
			case <-p.stopCh:
				ticker.Stop()
				p.finishCh <- true
				return
			case <-ticker.C:
				p.push()
			}
		}
	}()
}

// Close closes the exporter
func (p *PushGatewayExporter) Close() {
	p.stopCh <- true
	<-p.finishCh
	close(p.stopCh)
	close(p.finishCh)
}

func (p *PushGatewayExporter) push() {
	// avoid job is crashed when exception happens
	defer func() {
		if err := recover(); err != nil {
			log.Println(p.service, "Could not push to the Pushgateway", p.gatewayURL, err)
		}
	}()

	// another job is running. avoid sending to many requests to the server
	if !atomic.CompareAndSwapUint32(&p.locker, 0, 1) {
		return
	}
	defer atomic.StoreUint32(&p.locker, 0)

	err := p.pusher.Add()
	if err != nil {
		log.Println(p.service, "Could not push to the Pushgateway", p.gatewayURL, err)
	}
}
