package main

import (
	"careerhub-dataprovider/test/tinit"
	"testing"
)

func TestReset(t *testing.T) {
	tinit.InitCompanyRepo(t)
	tinit.InitJobPostingRepo(t)

	tinit.InitGrpcClient(t)
	tinit.InitGrpcClient(t)
	tinit.InitGrpcClient(t)
}
