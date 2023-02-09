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
