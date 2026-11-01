# Kube OpenAPI lite

This fork retains the upstream module path `k8s.io/kube-openapi`.
It keeps the OpenAPI v2 and v3 schema models, builders, HTTP handlers, caching,
schema conversion and mutation, REST adapters, and document validation visitors.
Its gnostic dependency uses protobuf-go-lite. Schema reference wrapping and
JSON/YAML value dispatch do not use reflection.

The schema generators, validation engine, format validators, aggregator,
conversion command, integration programs, scripts and upstream automation are
omitted. The proto/testing helpers remain solely to support dependency tests
and module tidying. Other test helper packages are omitted; the retained handler
fixture supports the HTTP/protobuf tests.

Generic Go model naming still uses reflection. Standard JSON, YAML, HTTP and
logging dependencies retain their own reflection and platform-specific unsafe
implementation details. This is a focused subset rather than a replacement for
every upstream package.
