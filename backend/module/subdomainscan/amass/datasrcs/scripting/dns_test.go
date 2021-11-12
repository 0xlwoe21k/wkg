// Copyright 2021 Jeff Foley. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package scripting

import (
	"strings"
	"testing"
	"time"

	"backend/module/subdomainscan/amass/requests"
)

func TestResolve(t *testing.T) {
	expected := []bool{false, false, true, false, false, true, true, false, true, true, true, true}
	ctx, sys := setupMockScriptEnv(`
		name="resolve"
		type="testing"

		function vertical(ctx, domain)
			local tests = {
				{ctx, "", "A"},
				{ctx, "www.owasp.org", ""},
				{ctx, "www.owasp.org", "AAAA"},
				{ctx, "www.owasp.org", "PTR"},
				{ctx, "www.sciencedirect.com.ezproxy.utica.edu", "CNAME"},
				{ctx, "ezproxy.utica.edu", "A"},
				{ctx, "_sips._tcp.utica.edu", "SRV"},
				{ctx, "bestsecurity.owasp.org", "A"},
				{ctx, "owasp.org", "NS"},
				{ctx, "owasp.org", "MX"},
				{ctx, "owasp.org", "TXT"},
				{ctx, "owasp.org", "SOA"},
			}

			for _, t in ipairs(tests) do
				local resp, err = resolve(t[1], t[2], t[3])
				if (err ~= nil and err ~= "") then
					log(ctx, "Error: " .. err)
				elseif #resp > 0 then
					log(ctx, resp[1].rrdata)
				end
    		end
		end
	`)
	if ctx == nil || sys == nil {
		t.Fatal("Failed to initialize the scripting environment")
	}
	defer func() { _ = sys.Shutdown() }()

	cfg, bus, err := requests.ContextConfigBus(ctx)
	if err != nil {
		t.Fatal("Failed to obtain the config and event bus")
	}

	numtests := len(expected)
	ch := make(chan string, numtests)
	fn := func(msg string) {
		ch <- msg
	}

	bus.Subscribe(requests.LogTopic, fn)
	defer bus.Unsubscribe(requests.LogTopic, fn)

	cfg.AddDomain("owasp.org")
	sys.DataSources()[0].Request(ctx, &requests.DNSRequest{Domain: "owasp.org"})
	// Read the initial message generated by the DNS Request
	<-ch

	timer := time.NewTimer(time.Duration(numtests*15) * time.Second)
	defer timer.Stop()
loop:
	for i := 0; i < numtests; i++ {
		select {
		case <-timer.C:
			t.Error("The test timed out")
			break loop
		case msg := <-ch:
			if ans := strings.Split(msg, "Error: "); len(ans) > 1 == expected[i] {
				t.Error(ans[len(ans)-1])
			}
		}
	}
}
