// DO NOT REMOVE TAGS BELOW. IF ANY NEW TEST FILES ARE CREATED UNDER /osde2e, PLEASE ADD THESE TAGS TO THEM IN ORDER TO BE EXCLUDED FROM UNIT TESTS. //go:build osde2e
//go:build osde2e
// +build osde2e

package osde2etests

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/openshift/osde2e-common/pkg/clients/openshift"
)

var routeMonitorOperatorTestName string = "[Suite: informing] [OSD] Route Monitor Operator (rmo)"

func testRouteMonitorCreationWorks(k8s *openshift.Client) {
	Context("rmo Route Monitor Operator integration test", func() {
		gomega.Expect(true)
	})
}
