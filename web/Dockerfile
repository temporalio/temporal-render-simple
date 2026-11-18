FROM temporalio/ui:2.58.1

# Render supplies the private-service hostname separately from its gRPC port.
ENTRYPOINT ["/bin/sh", "-c", "export TEMPORAL_ADDRESS=$TEMPORAL_GRPC_HOST:7233; exec ./start-ui-server.sh"]
