package bunny

import (
	"sigs.k8s.io/external-dns/endpoint"
)

func recordToEndpoint(domain string, record *Record) *endpoint.Endpoint {
	return endpoint.NewEndpointWithTTL(
		record.Name+"."+domain,
		record.Type.String(),
		endpoint.TTL(record.TTLSeconds),
		record.Value,
	)
}
