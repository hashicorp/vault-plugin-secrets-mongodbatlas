---
layout: "api"
page_title: "MongoDB Atlas - Secrets Engines - HTTP API"
sidebar_title: "MongoDB Atlas"
sidebar_current: "docs-secrets-engines-mongodbatlas"
description: |-
  The MongoDB Atlas Secrets Engine for Vault generates MongoDB Atlas Programmatic API Keys dynamically.
---

# MongoDB Atlas Secrets Engine

The MongoDB Atlas Secrets Engine generates Programmatic API keys for MongoDB Atlas. This allows one to manage the lifecycle of these MongoDB Atlas secrets programmatically. The created MongoDB Atlas secrets are
time-based and are automatically revoked when the Vault lease expires, unless renewed. Vault will create a Programmatic API key for each lease scoped to the MongoDB Atlas project or organization denoted with the included role(s). An IP Whitelist may also be configured for the Programmatic API key with desired IPs and/or CIDR blocks.

The MongoDB Atlas Programmatic API Key Public and
Private Key is returned to the caller. To learn more about Programmatic API Keys visit the [Programmatic API Keys Doc](https://docs.atlas.mongodb.com/reference/api/apiKeys/).

## Configure Connection

In addition to the parameters defined by the Secrets Engines Backend, this plugin has a number of parameters to further configure a connection.

| Method   | Path                         |
| :--------------------------- | :--------------------- |
| `POST`   | `/mongodbatlas/config`     |


## Parameters

- `public_key` `(string: <required>)` – The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
- `private_key` `(string: <required>)` - The Private Programmatic API Key used to connect with MongoDB Atlas API.

### Sample Payload

```json
{
  "public_key": "aPublicKey",
  "private_key": "aPrivateKey",
}
```

### Sample Request
```bash
$ curl \
    --header "X-Vault-Token: ..." \
    --request POST \
    --data @payload.json \
    http://127.0.0.1:8200/mongodbatlas/config`
```

## Create/Update Programmatic API Key role
Programmatic API Key credential types create a Vault role to generate a Programmatic API Key at
either the MongoDB Atlas Organization or Project level with the designated role(s) for programmatic access. If a role with the name does not exist, it will be created. If the role exists, it will be updated with the new attributes.

| Method   | Path                         |
| :--------------------------- | :--------------------- |
| `POST`   | `/roles/:name`     |


## Parameters

`name` `(string <required>)` - Unique identifier name of the role name
`project_id` `(string <required>)` - Unique identifier for the organization to which the target API Key belongs. Use the /orgs endpoint to retrieve all organizations to which the authenticated user has access.
`roles` `(list [string] <required>)` - List of roles that the API Key needs to have. If the roles array is provided:

  -> **IMPORTANT:** Provide at least one role. Make sure all roles must be valid for the Organization or Project.

  -> **NOTE:** Include all roles that you want this API Key to have. Any roles not in this array are removed.

  - The Organization roles are:
    - `ORG_OWNER`
    - `ORG_MEMBER`
    - `ORG_GROUP_CREATOR`
    - `ORG_BILLING_ADMIN`
    - `ORG_READ_ONLY`

  - The Project roles are:
    - `GROUP_CHARTS_ADMIN`
    - `GROUP_CLUSTER_MANAGER`
    - `GROUP_DATA_ACCESS_ADMIN`
    - `GROUP_DATA_ACCESS_READ_ONLY`
    - `GROUP_DATA_ACCESS_READ_WRITE`
    - `GROUP_OWNER`
    - `GROUP_READ_ONLY`


`ip_addresses` `(list [string] <Optional>)` - IP address to be added to the whitelist for the API key. This field is mutually exclusive with the cidrBlock field.
`cidr_blocks` `(list [string] <Optional>)` - Whitelist entry in CIDR notation to be added for the API key. This field is mutually exclusive with the ipAddress field.

### Sample Payload

```json
{
  "project_id": "5cf5a45a9ccf6400e60981b6",
  "roles": ["GROUP_CLUSTER_MANAGER"],
  "cidr_blocks": ["192.168.1.3/32"],
  "ip_addresses": ["192.168.1.3", "192.168.1.4"]
}
```

```bash
$ curl \
    --header "X-Vault-Token: ..." \
    --request POST \
    --data @payload.json \
    http://127.0.0.1:8200/mongodbatlas/roles/test-programmatic-key
```

### Sample Response
```json
{
  "project_id": "5cf5a45a9ccf6400e60981b6",
  "roles": ["GROUP_CLUSTER_MANAGER"],
  "cidr_blocks": ["192.168.1.3/32"],
  "ip_addresses": ["192.168.1.3", "192.168.1.4"],
  "organization_id": "7cf5a45a9ccf6400e60981b7",
  "ttl": "0s",
  "max_ttl": "0s"
}

```

## Read Programmatic API Key role

| Method   | Path                         |
| :--------------------------- | :--------------------- |
| `Get`   | `/roles/:name`     |


## Parameters

`name` `(string <required>)` - Unique identifier name of the role name

### Sample Payload

```bash
$ curl \
    --header "X-Vault-Token: ..." \
    --request GET \
    --data @payload.json \
    http://127.0.0.1:8200/mongodbatlas/roles/test-programmatic-key
```

### Sample Response
```json
{
  "project_id": "5cf5a45a9ccf6400e60981b6",
  "roles": ["GROUP_CLUSTER_MANAGER"],
  "cidr_blocks": ["192.168.1.3/32"],
  "ip_addresses": ["192.168.1.3", "192.168.1.4"],
  "organization_id": "7cf5a45a9ccf6400e60981b7",
  "ttl": "0s",
  "max_ttl": "0s"
}
```

## List Programmatic API Key role

| Method   | Path                         |
| :--------------------------- | :--------------------- |
| `Get`   | `/roles`     |


### Sample Payload

```bash
$ curl \
    --header "X-Vault-Token: ..." \
    --request GET \
    --data @payload.json \
    http://127.0.0.1:8200/mongodbatlas/roles
```

### Sample Response
```json
[
  {
    "project_id": "5cf5a45a9ccf6400e60981b6",
    "roles": ["GROUP_CLUSTER_MANAGER"],
    "cidr_blocks": ["192.168.1.3/32"],
    "ip_addresses": ["192.168.1.3", "192.168.1.4"],
    "organization_id": "7cf5a45a9ccf6400e60981b7",
    "ttl": "0s",
    "max_ttl": "0s"
  },
  {
    "project_id": "5cf5a45a9ccf6400e60981b6",
    "roles": ["READ"],
    "cidr_blocks": ["192.168.1.3/35"],
    "ip_addresses": ["192.168.1.5", "192.168.1.6"],
    "organization_id": "7cf5a45a9ccf6400e60981b7",
    "ttl": "0s",
    "max_ttl": "0s"
  }
]

```

## Delete Programmatic API Key role

| Method   | Path                         |
| :--------------------------- | :--------------------- |
| `DELETE`   | `/roles/:name`     |


## Parameters

`name` `(string <required>)` - Unique identifier name of the role name

### Sample Payload

```bash
$ curl \
    --header "X-Vault-Token: ..." \
    --request DELETE \
    --data @payload.json \
    http://127.0.0.1:8200/mongodbatlas/roles/test-programmatic-key
```

### Sample Response
```json
{}
```

## Read Credential

### Sample Request

| Method   | Path                         |
| :--------------------------- | :--------------------- |
| `GET`   | `/creds/:name`     |

## Parameters
`name` `(string <required>)` - Unique identifier name of the credential

```bash
$ curl \
    --header "X-Vault-Token: ..." \
    http://127.0.0.1:8200/mongodbatlas/creds/0fLBv1c2YDzPlJB1PwsRRKHR
```

### Sample Response
```json
{
  "lease_duration": "20s",
  "lease_renewable": true,
  "description": "vault-test-1563980947-1318",
  "private_key": "905ae89e-6ee8-40rd-ab12-613t8e3fe836",
  "public_key": "klpruxce"
}
```