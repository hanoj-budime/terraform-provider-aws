#!/usr/bin/env bash

# Code generated by internal/generate/teamcity/provider_tests.go; DO NOT EDIT.

set -euo pipefail

go test \
    ./internal/acctest/... \
    ./internal/attrmap/... \
    ./internal/conns/... \
    ./internal/create/... \
    ./internal/enum/... \
    ./internal/envvar/... \
    ./internal/errs/... \
    ./internal/experimental/... \
    ./internal/flex/... \
    ./internal/framework/... \
    ./internal/function/... \
    ./internal/generate/... \
    ./internal/io/... \
    ./internal/json/... \
    ./internal/logging/... \
    ./internal/maps/... \
    ./internal/namevaluesfilters/... \
    ./internal/provider/... \
    ./internal/retry/... \
    ./internal/sdkv2/... \
    ./internal/semver/... \
    ./internal/slices/... \
    ./internal/sweep/... \
    ./internal/tags/... \
    ./internal/tfresource/... \
    ./internal/types/... \
    ./internal/vault/... \
    ./internal/verify/... \
    ./internal/yaml/... \
    -json
