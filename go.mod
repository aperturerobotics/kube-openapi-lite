module k8s.io/kube-openapi

go 1.27.0

require (
	github.com/emicklei/go-restful/v3 v3.13.0
	github.com/go-openapi/jsonreference v1.0.0
	github.com/go-openapi/swag/jsonutils v0.27.1
	github.com/google/gnostic-models v0.7.0
	github.com/google/uuid v1.6.0
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
	github.com/stretchr/testify v1.12.1
	go.yaml.in/yaml/v2 v2.4.3
	go.yaml.in/yaml/v3 v3.0.5
	k8s.io/klog/v2 v2.140.0
	sigs.k8s.io/randfill v1.0.0
	sigs.k8s.io/structured-merge-diff/v7 v7.0.0
)

require (
	github.com/aperturerobotics/json-iterator-lite v1.1.0 // indirect
	github.com/aperturerobotics/protobuf-go-lite v0.19.1-0.20261028131920-f644656fc4a7 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v1.0.0 // indirect
	github.com/go-openapi/swag/conv v0.27.1 // indirect
	github.com/go-openapi/swag/pools v0.27.1 // indirect
	github.com/go-openapi/swag/typeutils v0.27.1 // indirect
)

replace github.com/google/gnostic-models => github.com/aperturerobotics/gnostic-models-lite v0.0.0-20261101174401-df56aa581c7f
