package util

// InfisicalAuthEnvKeys are the environment variables the CLI's own auth resolution reads. Commands
// that launch a child process (run, pam agentic access) drop these from the inherited environment, so
// a credential the CLI was handed is not passed on to everything the child spawns in turn.
var InfisicalAuthEnvKeys = []string{
	INFISICAL_TOKEN_NAME,
	INFISICAL_UNIVERSAL_AUTH_ACCESS_TOKEN_NAME,
	INFISICAL_GATEWAY_TOKEN_NAME_LEGACY, // bare "TOKEN", still read by GetInfisicalToken
	INFISICAL_UNIVERSAL_AUTH_CLIENT_ID_NAME,
	INFISICAL_UNIVERSAL_AUTH_CLIENT_SECRET_NAME,
	INFISICAL_JWT_NAME,
	INFISICAL_OIDC_AUTH_JWT_NAME,
	INFISICAL_LDAP_USERNAME,
	INFISICAL_LDAP_PASSWORD,
	INFISICAL_KUBERNETES_SERVICE_ACCOUNT_TOKEN_NAME,
	INFISICAL_GCP_IAM_SERVICE_ACCOUNT_KEY_FILE_PATH_NAME,
	INFISICAL_AUTH_METHOD_NAME,
	INFISICAL_MACHINE_IDENTITY_ID_NAME,
	INFISICAL_ORGANIZATION_ID,
	INFISICAL_VAULT_FILE_PASSPHRASE_ENV_NAME, // unlocks the file-backed keyring
	INFISICAL_BOOTSTRAP_EMAIL_NAME,
	INFISICAL_BOOTSTRAP_PASSWORD_NAME,
	INFISICAL_BOOTSTRAP_ORGANIZATION_NAME,
}

// IsInfisicalAuthEnvKey reports whether name is one of InfisicalAuthEnvKeys.
func IsInfisicalAuthEnvKey(name string) bool {
	for _, key := range InfisicalAuthEnvKeys {
		if key == name {
			return true
		}
	}

	return false
}
