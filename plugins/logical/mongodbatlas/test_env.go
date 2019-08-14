package mongodbatlas

import (
	"context"
	"testing"

	"github.com/hashicorp/vault/sdk/logical"
)

type testEnv struct {
	PublicKey      string
	PrivateKey     string
	ProjectID      string
	OrganizationID string

	Backend logical.Backend
	Context context.Context
	Storage logical.Storage

	MostRecentSecret *logical.Secret
}

func (e *testEnv) AddConfig(t *testing.T) {
	req := &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      "config/root",
		Storage:   e.Storage,
		Data: map[string]interface{}{
			"public_key":  e.PublicKey,
			"private_key": e.PrivateKey,
		},
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
	if resp != nil {
		t.Fatal("expected nil response to represent a 204")
	}
}

func (e *testEnv) AddProgrammaticAPIKeyRole(t *testing.T) {
	roles := []string{"ORG_MEMBER"}
	req := &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      "roles/test-programmatic-key",
		Storage:   e.Storage,
		Data: map[string]interface{}{
			"credential_type": "org_programmatic_api_key",
			"organization_id": e.OrganizationID,
			"roles":           roles,
		},
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
}

func (e *testEnv) AddProgrammaticAPIKeyRoleWithIP(t *testing.T) {
	roles := []string{"ORG_MEMBER"}
	ips := []string{"192.168.1.1", "192.168.1.2"}
	req := &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      "roles/test-programmatic-key",
		Storage:   e.Storage,
		Data: map[string]interface{}{
			"credential_type": "org_programmatic_api_key",
			"organization_id": e.OrganizationID,
			"roles":           roles,
			"ip_addresses":    ips,
		},
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
}

func (e *testEnv) AddProgrammaticAPIKeyRoleWithCIDR(t *testing.T) {
	roles := []string{"ORG_MEMBER"}
	cidrBlocks := []string{"179.154.224.2/32"}
	req := &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      "roles/test-programmatic-key",
		Storage:   e.Storage,
		Data: map[string]interface{}{
			"credential_type": "org_programmatic_api_key",
			"organization_id": e.OrganizationID,
			"roles":           roles,
			"cidr_blocks":     cidrBlocks,
		},
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
}

func (e *testEnv) AddProgrammaticAPIKeyRoleWithProjectID(t *testing.T) {
	roles := []string{"ORG_MEMBER"}
	req := &logical.Request{
		Operation: logical.UpdateOperation,
		Path:      "roles/test-programmatic-key",
		Storage:   e.Storage,
		Data: map[string]interface{}{
			"credential_type": "project_programmatic_api_key",
			"roles":           roles,
			"project_id":      e.ProjectID,
		},
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
}

func (e *testEnv) ReadProgrammaticAPIKeyRule(t *testing.T) {
	req := &logical.Request{
		Operation: logical.ReadOperation,
		Path:      "creds/test-programmatic-key",
		Storage:   e.Storage,
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
	if resp == nil {
		t.Fatal("expected a response")
	}

	if resp.Data["public_key"] == "" {
		t.Fatal("failed to receive access_key")
	}
	if resp.Data["private_key"] == "" {
		t.Fatal("failed to receive secret_key")
	}
	e.MostRecentSecret = resp.Secret
}

func (e *testEnv) RenewProgrammaticAPIKeys(t *testing.T) {
	req := &logical.Request{
		Operation: logical.RenewOperation,
		Storage:   e.Storage,
		Secret:    e.MostRecentSecret,
		Data: map[string]interface{}{
			"lease_id": "foo",
		},
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
	if resp == nil {
		t.Fatal("expected a response")
	}
	if resp.Secret != e.MostRecentSecret {
		t.Fatalf("expected %+v but got %+v", e.MostRecentSecret, resp.Secret)
	}
}

func (e *testEnv) RevokeProgrammaticAPIKeys(t *testing.T) {
	req := &logical.Request{
		Operation: logical.RevokeOperation,
		Storage:   e.Storage,
		Secret:    e.MostRecentSecret,
		Data: map[string]interface{}{
			"lease_id": "foo",
		},
	}
	resp, err := e.Backend.HandleRequest(e.Context, req)
	if err != nil || (resp != nil && resp.IsError()) {
		t.Fatalf("bad: resp: %#v\nerr:%v", resp, err)
	}
	if resp != nil {
		t.Fatal("expected nil response to represent a 204")
	}
}
