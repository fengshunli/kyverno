package internal

import (
	"context"
	"net"

	"github.com/go-logr/logr"
	"github.com/kyverno/kyverno/pkg/tracing"
	"k8s.io/client-go/kubernetes"
)

func SetupTracing(logger logr.Logger, kubeClient kubernetes.Interface) context.CancelFunc {
	logger = logger.WithName("tracing").WithValues("enabled", tracingEnabled, "address", tracingAddress, "port", tracingPort, "creds", tracingCreds)
	if tracingEnabled {
		logger.Info("setup tracing...")
		shutdown, err := tracing.NewTraceConfig(
			logger,
			net.JoinHostPort(tracingAddress, tracingPort),
			tracingCreds,
			kubeClient,
		)
		checkError(logger, err, "failed to setup tracing")
		return shutdown
	}
	return func() {}
}
