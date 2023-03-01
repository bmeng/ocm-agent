package consts

const (
	// Listening port for the OCM Agent web service
	OCMAgentServicePort = 8081
	// Listening port for the OCM Agent metrics
	OCMAgentMetricsPort = 8383

	// Metrics path for OCM Agent service
	MetricsPath = "/metrics"
	// Ready probe path for OCM Agent web service
	ReadyzPath = "/readyz"
	// Live probe path for OCM Agent web service
	LivezPath = "/livez"
	// Alertmanager webhook receiver path
	WebhookReceiverPath = "/alertmanager-receiver"

	// Alertmanager webhook receiver path for fleet mode
	WebhookRHOBSReceiverPath = "/rhobs-receiver"

	// Service name for the sending service logs
	ServiceLogServiceName = "SREManualAction"
)
