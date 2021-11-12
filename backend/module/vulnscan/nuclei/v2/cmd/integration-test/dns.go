package main

import (
	testutils2 "backend/module/vulnscan/nuclei/v2/internal/testutils"
)

var dnsTestCases = map[string]testutils2.TestCase{
	"dns/basic.yaml": &dnsBasic{},
}

type dnsBasic struct{}

// Execute executes a test case and returns an error if occurred
func (h *dnsBasic) Execute(filePath string) error {
	var routerErr error

	results, err := testutils2.RunNucleiAndGetResults(filePath, "one.one.one.one", debug)
	if err != nil {
		return err
	}
	if routerErr != nil {
		return routerErr
	}
	if len(results) != 1 {
		return errIncorrectResultsCount(results)
	}
	return nil
}
