// DO NOT REMOVE TAGS BELOW. IF ANY NEW TEST FILES ARE CREATED UNDER /osde2e, PLEASE ADD THESE TAGS TO THEM IN ORDER TO BE EXCLUDED FROM UNIT TESTS. //go:build osde2e
//go:build osde2e
// +build osde2e

package osde2etests

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	prometheusop "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/openshift/osde2e-common/pkg/clients/openshift"
	. "github.com/openshift/osde2e-common/pkg/gomega/assertions"
	. "github.com/openshift/osde2e-common/pkg/gomega/matchers"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var _ = Describe("Route Monitor Operator", Ordered, func() {
	var (
		k8s               *openshift.Client
		operatorNamespace = "openshift-route-monitor-operator"
		deploymentName    = "route-monitor-operator-controller-manager"
		operatorName      = "route-monitor-operator"
		consoleNamespace  = operatorNamespace
		consoleName       = "console"
	)
	const (
		defaultDesiredReplicas int32 = 1
	)
	BeforeAll(func(ctx context.Context) {
		log.SetLogger(GinkgoLogr)
		var err error
		k8s, err = openshift.New(GinkgoLogr)
		Expect(err).ShouldNot(HaveOccurred(), "unable to setup k8s client")

	})

	It("is installed", func(ctx context.Context) {

		By("checking the deployment exists and is available")
		EventuallyDeployment(ctx, k8s, deploymentName, operatorNamespace).Should(BeAvailable())
	})

	It("can be upgraded", func(ctx context.Context) {
		By("forcing operator upgrade")
		err := k8s.UpgradeOperator(ctx, operatorName, operatorNamespace)
		Expect(err).NotTo(HaveOccurred(), "operator upgrade failed")
	})

	Context("rmo Route Monitor Operator regression for console", func() {
		It("has all of the required resources", func(ctx context.Context) {

			promclient, err := prometheusop.NewForConfig(k8s.GetConfig())
			Expect(err).ShouldNot(HaveOccurred(), "failed to configure Prometheus-operator clientset")

			_, err = promclient.MonitoringV1().ServiceMonitors(consoleNamespace).Get(ctx, consoleName, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred(), "Could not get console serviceMonitor")
			_, err = promclient.MonitoringV1().PrometheusRules(consoleNamespace).Get(ctx, consoleName, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred(), "Could not get console prometheusRule")

			//results, err := prom.InstantQuery(context.Background(), `up{job="route-monitor-operator"}`)
			//Expect(err).NotTo(HaveOccurred(), "failed to query prometheus")
			//Expect(results).NotTo(BeNil(), "No results for prometheus exporter")
			//Expect(results.Len()).To(BeNumerically(">=", 0), "No metrics returned for the route-monitor-operator job")
			//
			//query := `count(up{namespace="` + consoleNamespace + `", service="` + consoleName + `"})`
			//srvresult, err := prom.InstantQuery(context.Background(), query)
			//Expect(err).NotTo(HaveOccurred(), "Could not get console serviceMonitor")
			//Expect(srvresult).NotTo(BeNil(), "No results for ServiceMonitor query")
			//Expect(srvresult.Len()).To(BeNumerically(">=", 0), "ServiceMonitor is not active")
			//
			//ruleQuery := `count(prometheus_rule_group_last_duration_seconds{namespace="` + consoleNamespace + `", rule_group="` + consoleName + `"})`
			//ruleResult, err := prom.InstantQuery(context.Background(), ruleQuery)
			//Expect(err).NotTo(HaveOccurred(), "Could not get console prometheusRule")
			//Expect(ruleResult).NotTo(BeNil(), "No results for PrometheusRule query")
			//Expect(ruleResult.Len()).To(BeNumerically(">=", 0), "PrometheusRule is not active")
		})
	})
	/*
			// TODO: implement testRouteMonitorCreationWorks
		 	Context("rmo Route Monitor Operator integration test", func(ctx context.Context) {
	*/
})
