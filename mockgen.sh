#!/usr/bin/env bash

mockgen -destination pkg/testing/mocks/mock_doer.go -package=mocks -source pkg/testing/doer/doer.go 