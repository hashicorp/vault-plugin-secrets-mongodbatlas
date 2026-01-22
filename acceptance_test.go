// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package mongodbatlas

import (
	"context"
	"os"
	"testing"
	"time"

	log "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/sdk/helper/logging"
	"github.com/hashicorp/vault/sdk/logical"
)

const (
	envVarRunAccTests    = "VAULT_ACC"
	envVarPrivateKey     = "ATLAS_PRIVATE_KEY"
	envVarPublicKey      = "ATLAS_PUBLIC_KEY"
	envVarProjectID      = "ATLAS_PROJECT_ID"
	envVarOrganizationID = "ATLAS_ORGANIZATION_ID"
)

var runAcceptanceTests = os.Getenv(envVarRunAccTests) == "1"

func TestAcceptanceProgrammaticAPIKey(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRole)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

func TestAcceptanceProgrammaticAPIKey_WithProjectID(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithProjectID)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

func TestAcceptanceProgrammaticAPIKey_WithProjectIDRenew(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithProjectIDWithTTL)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("check lease for programmatic API key cred", acceptanceTestEnv.CheckLease)
	t.Run("extend programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeysWithExtendedLease)
	t.Run("check lease for programmatic API key cred", acceptanceTestEnv.CheckExtendedLease)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

// requires IP to be set in the API key access list
// see https://www.mongodb.com/docs/atlas/configure-api-access/
func TestAcceptanceProgrammaticAPIKey_ProjectWithIPAccesslist(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleProjectWithIP)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

// requires IP to be set in the API key access list
// see https://www.mongodb.com/docs/atlas/configure-api-access/
func TestAcceptanceProgrammaticAPIKey_WithIPAccesslist(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithIP)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

// requires IP to be set in the API key access list
// see https://www.mongodb.com/docs/atlas/configure-api-access/
func TestAcceptanceProgrammaticAPIKey_WithCIDRAccesslist(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithCIDR)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

// requires IP to be set in the API key access list
// see https://www.mongodb.com/docs/atlas/configure-api-access/
func TestAcceptanceProgrammaticAPIKey_WithCIDRAndIPAccesslist(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithCIDRAndIP)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

// requires IP to be set in the API key access list
// see https://www.mongodb.com/docs/atlas/configure-api-access/
func TestAcceptanceProgrammaticAPIKey_AssignToProject(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithProjectIDAndOrgID)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

func TestAcceptanceProgrammaticAPIKey_WithTTL(t *testing.T) {
	if !runAcceptanceTests {
		t.SkipNow()
	}

	acceptanceTestEnv, err := newAcceptanceTestEnv()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("add config", acceptanceTestEnv.AddConfig)
	t.Run("add programmatic API Key role with TTL", acceptanceTestEnv.AddProgrammaticAPIKeyRoleWithTTL)
	t.Run("read programmatic API key cred", acceptanceTestEnv.ReadProgrammaticAPIKeyRule)
	t.Run("check lease for programmatic API key cred", acceptanceTestEnv.CheckLease)
	t.Run("renew programmatic API key creds", acceptanceTestEnv.RenewProgrammaticAPIKeys)
	t.Run("revoke programmatic API key creds", acceptanceTestEnv.RevokeProgrammaticAPIKeys)
}

func newAcceptanceTestEnv() (*testEnv, error) {
	ctx := context.Background()

	maxLease, _ := time.ParseDuration("60s")
	defaultLease, _ := time.ParseDuration("30s")
	conf := &logical.BackendConfig{
		System: &logical.StaticSystemView{
			DefaultLeaseTTLVal: defaultLease,
			MaxLeaseTTLVal:     maxLease,
		},
		Logger: logging.NewVaultLogger(log.Debug),
	}
	b, err := Factory(ctx, conf)
	if err != nil {
		return nil, err
	}

	te := &testEnv{
		PublicKey:      os.Getenv(envVarPublicKey),
		PrivateKey:     os.Getenv(envVarPrivateKey),
		ProjectID:      os.Getenv(envVarProjectID),
		OrganizationID: os.Getenv(envVarOrganizationID),
		Backend:        b,
		Context:        ctx,
		Storage:        &logical.InmemStorage{},
	}
	if err := te.validate(); err != nil {
		return nil, err
	}

	return te, nil
}
