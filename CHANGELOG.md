## Unreleased

## 0.10.2
### October 25, 2023

IMPROVEMENTS:
* Improve handing of fetching API keys for revocation of leases [GH-59](https://github.com/hashicorp/vault-plugin-secrets-mongodbatlas/pull/59)

## 0.10.1
### September 1, 2023

IMPROVEMENTS:
* Update dependencies:
  * github.com/hashicorp/vault/api v1.9.1 -> v1.9.2
  * github.com/hashicorp/vault/sdk v0.9.0 -> v0.9.2
  * github.com/mongodb-forks/digest v1.0.4 -> v1.0.5
  * go.mongodb.org/atlas v0.25.0 -> v0.33.0

## 0.10.0
### May 23, 2023

IMPROVEMENTS:
* Update Go version to 1.20.2
* Add display attributes for OpenAPI OperationID
* enable plugin multiplexing [GH-35](https://github.com/hashicorp/vault-plugin-secrets-mongodbatlas/pull/35)
* update dependencies
  * `github.com/hashicorp/vault/api` v1.9.1
  * `github.com/hashicorp/vault/sdk` v0.8.1
  * `go.mongodb.org/atlas` v0.25.0

## 0.9.1
### February 9, 2023

Bug Fixes:
* Fix a bug that did not allow WAL rollback to handle partial failures when
  creating API keys [GH-32](https://github.com/hashicorp/vault-plugin-secrets-mongodbatlas/pull/32)

Improvements:
* Update dependencies [GH-33](https://github.com/hashicorp/vault-plugin-secrets-mongodbatlas/pull/33)
  * github.com/hashicorp/vault/api v1.8.3
  * github.com/hashicorp/vault/sdk v0.7.0

## 0.9.0
### February 6, 2023

Improvements:
* Change how the Vault version is acquired for the user agent string. This
  change is transparent to users.
