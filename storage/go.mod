module cloud.google.com/go/storage

go 1.22.0

toolchain go1.23.4

retract [v1.25.0, v1.27.0] // due to https://github.com/googleapis/google-cloud-go/issues/6857

require (
	cloud.google.com/go v0.118.0
	cloud.google.com/go/compute/metadata v0.6.0
	cloud.google.com/go/iam v1.3.1
	cloud.google.com/go/longrunning v0.6.4
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.49.0
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/googleapis/gax-go/v2 v2.14.1
	go.opentelemetry.io/contrib/detectors/gcp v1.33.0
	go.opentelemetry.io/otel v1.33.0
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.29.0
	go.opentelemetry.io/otel/sdk v1.33.0
	go.opentelemetry.io/otel/sdk/metric v1.31.0
  go.opentelemetry.io/otel/trace v1.33.0
	golang.org/x/oauth2 v0.25.0
	golang.org/x/sync v0.10.0
	google.golang.org/api v0.217.0
	google.golang.org/genproto v0.0.0-20241118233622-e639e219e697
	google.golang.org/genproto/googleapis/api v0.0.0-20250115164207-1a7da9e5054f
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f
	google.golang.org/grpc v1.69.4
	google.golang.org/protobuf v1.36.3
)

require (
	cel.dev/expr v0.16.2 // indirect
	cloud.google.com/go/auth v0.14.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.7 // indirect
	cloud.google.com/go/monitoring v1.21.2 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.25.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.49.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78 // indirect
	github.com/envoyproxy/go-control-plane/envoy v1.32.3 // indirect
	github.com/envoyproxy/go-control-plane/ratelimit v0.1.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/martian/v3 v3.3.3 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel/metric v1.33.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.9.0 // indirect
)
