FROM gcr.io/distroless/static-debian12:nonroot

USER 20000:20000
ADD --chmod=555 external-dns-bunny-webhook /opt/external-dns-bunny-webhook/app

ENTRYPOINT ["/opt/external-dns-bunny-webhook/app"]
