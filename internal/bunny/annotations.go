package bunny

import (
	"fmt"
	"strconv"

	"sigs.k8s.io/external-dns/endpoint"
)

const (
	providerSpecificMonitorType = "webhook/bunny-monitor-type"
	providerSpecificWeight      = "webhook/bunny-weight"
)

type providerSpecificOptions struct {
	MonitorType MonitorType
	Weight      int
}

func providerSpecificOptionsFromEndpoint(e *endpoint.Endpoint) (providerSpecificOptions, error) {
	opts := providerSpecificOptions{}

	if monitorType, ok := e.GetProviderSpecificProperty(providerSpecificMonitorType); ok {
		opts.MonitorType = MonitorTypeFromString(monitorType)
	}

	if weight, ok := e.GetProviderSpecificProperty(providerSpecificWeight); ok {
		var err error
		opts.Weight, err = strconv.Atoi(weight)
		if err != nil {
			return opts, fmt.Errorf("weight %s is not parseable as an integer: %w", weight, err)
		}
	}

	if opts.Weight == 0 {
		opts.Weight = 100
	}

	return opts, nil
}

func providerSpecificOptionsFromRecord(r *Record) *providerSpecificOptions {
	opts := &providerSpecificOptions{
		MonitorType: r.MonitorType,
		Weight:      r.Weight,
	}

	return opts
}

func (p *providerSpecificOptions) ApplyToEndpoint(e *endpoint.Endpoint) {
	e.WithProviderSpecific(providerSpecificMonitorType, p.MonitorType.String())
	e.WithProviderSpecific(providerSpecificWeight, strconv.Itoa(p.Weight))
}
