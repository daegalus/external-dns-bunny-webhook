# ExternalDNS - Bunny Webhook Provider

ExternalDNS is a Kubernetes add-on for automatically managing Domain Name System (DNS) records for
Kubernetes services by using different DNS providers. By default, Kubernetes manages DNS records
internally, but ExternalDNS takes this functionality a step further by delegating the management of
DNS records to an external DNS provider such as this one.

This repository contains a provider that implements an ExternalDNS webhook provider for [Bunny.net](https://bunny.net).

## Important

This provider is not officially supported by [Bunny.net](https://bunny.net), but is maintained by the team at Contaim Labs
for the community. If you encounter any issues, please open an issue on this repository. If you have any questions
about Bunny.net, please reach out to their support team.

## Deployment

You can deploy the provider using any Kubernetes deployment method, such as Helm or kubectl. Examples for the official
external-dns Helm chart are provided below.

### External DNS Helm Chart

The default configuration is designed to work seamlessly with the official ExternalDNS Helm chart. Be sure to create the
`external-dns-bunny-secret` secret with the `api-key` key containing your Bunny.net API key or modify the configuration
to use a different secret or method of providing the API key.

The values file should look similar to the following:

```yaml
namespace: external-dns
provider:
  name: webhook
  webhook:
    image:
      repository: ghcr.io/contain-labs/external-dns-bunny-webhook
      tag: v0.1.0
    env:
      - name: BUNNY_API_KEY
        valueFrom:
          secretKeyRef:
            name: external-dns-bunny-secret
            key: api-key
```

To deploy the provider using the Helm chart, add the repository. You can skip this step if you already have the
repository added.

```shell
helm repo add external-dns https://kubernetes-sigs.github.io/external-dns/
```

Once the repository is added, install the chart with the values file.

```shell
helm upgrade --install external-dns external-dns/external-dns --version 1.15.0
```

Additional configuration options are available below and may be set using environment variables.

## Configuration

The provider can be configured using the following environment variables:

| Environment Variable | Required | Description | Default |
|----------------------|----------|-------------|---------|
| `BUNNY_API_KEY` | Yes | The API key used to authenticate with the Bunny.net API. | |
| `BUNNY_DRY_RUN` | No | If set to `true`, the provider will not make any changes to the DNS records. | `false` |
| `WEBHOOK_HOST` | No | The host to use for the webhook endpoint. | `localhost` |
| `WEBHOOK_PORT` | No | The port to use for the webhook endpoint. | `8888` |
| `WEBHOOK_READ_TIMEOUT` | No | The read timeout for the webhook endpoint. | `60s` |
| `WEBHOOK_WRITE_TIMEOUT` | No | The write timeout for the webhook endpoint. | `60s` |
| `HEALTH_HOST` | No | The host to use for the health endpoint. | `0.0.0.0` |
| `HEALTH_PORT` | No | The port to use for the health endpoint. | `8080` |
| `HEALTH_READ_TIMEOUT` | No | The read timeout for the health endpoint. | `60s` |
| `HEALTH_WRITE_TIMEOUT` | No | The write timeout for the health endpoint. | `60s` |
