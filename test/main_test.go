package main

// import (
// 	"careerhub-dataprovider/careerhub/provider/vars"
// 	"careerhub-dataprovider/test/tinit"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestReset(t *testing.T) {
// 	tinit.InitCompanyRepo(t)
// 	tinit.InitJobPostingRepo(t)
// 	envVars, err := vars.Variables()
// 	require.NoError(t, err)

// 	tinit.InitSQS(t, envVars.JobPostingQueue)
// 	tinit.InitSQS(t, envVars.CompanyQueue)
// 	tinit.InitSQS(t, envVars.ClosedQueue)
// }
