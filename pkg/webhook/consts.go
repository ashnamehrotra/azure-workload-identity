package webhook

// Annotations and labels defined in service account
const (
	// UsePodIdentityLabel represents the service account is to be used for pod identity
	UsePodIdentityLabel = "azure.pod.identity/use"
	// ClientIDAnnotation represents the clientID to be used with pod
	ClientIDAnnotation = "azure.pod.identity/client-id"
	// TenantIDAnnotation represent the tenantID to be used with pod
	TenantIDAnnotation = "azure.pod.identity/tenant-id"
	// ServiceAccountTokenExpiryAnnotation represents the expirationSeconds for projected service account token
	// [OPTIONAL] field. User might want to configure this to prevent any downtime caused by errors during service account token refresh.
	// Kubernetes service account token expiry will not be correlated with AAD tokens. AAD tokens expiry will be 24h.
	ServiceAccountTokenExpiryAnnotation = "azure.pod.identity/service-account-token-expiration"
	// SkipContainersAnnotation represents list of containers to skip added projected service account token volume
	// By default, the projected service account token volume will be added to all containers if the service account is annotated with clientID
	SkipContainersAnnotation = "azure.pod.identity/skip-containers"

	// DefaultServiceAccountTokenExpiration is the default service account token expiration in seconds
	DefaultServiceAccountTokenExpiration = int64(86400)
	// MinServiceAccountTokenExpiration is the minimum service account token expiration in seconds
	MinServiceAccountTokenExpiration = int64(3600)
)

// Environment variables injected in the pod
const (
	AzureClientIDEnvVar = "AZURE_CLIENT_ID"
	AzureTenantIDEnvVar = "AZURE_TENANT_ID"
	TokenFilePathEnvVar = "TOKEN_FILE_PATH"
)
