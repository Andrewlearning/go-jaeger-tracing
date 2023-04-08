# go-jaeger-tracing
A simple demo for using tracing across a few services

# Introduce
### The calling chain

`a service -> b service -> c service`

###  Three IDs
TraceID: A trace has a unique trace ID that identifies the entire trace
ParentID: Span ID of previous service \
SpanID:  Span ID is generate by current service that identifies this service itself

# Expected outcome
```
2023/04/08 15:50:36 [Service A] traceID: 18c4e30a5910e5b3, parentID: 0000000000000000, spanID:18c4e30a5910e5b3
2023/04/08 15:50:36 [Service B] traceID: 18c4e30a5910e5b3, parentID: 18c4e30a5910e5b3, spanID:2a3d70d0e33ac13a
2023/04/08 15:50:36 [Service C] traceID: 18c4e30a5910e5b3, parentID: 2a3d70d0e33ac13a, spanID:6e8e71cbadf1eff9
```