#!/usr/bin/env bash

set -ex

cd cmd/vault-plugin-secrets-mongodbatlas/ && GOOS=linux CGO_ENABLED=0 go build -o vault-plugin-secrets-mongodbatlas
cd ../../ && mv cmd/vault-plugin-secrets-mongodbatlas/vault-plugin-secrets-mongodbatlas .

docker kill vaultplg 2>/dev/null || true
tmpdir=$(mktemp -d vaultplgXXXXXX)
mkdir "$tmpdir/data"
chmod 777 "$tmpdir/data"
docker run --rm -d -p8200:8200 --name vaultplg -v "$(pwd)/$tmpdir/data":/data -v $(pwd):/vault/plugins --cap-add=IPC_LOCK -e 'VAULT_LOCAL_CONFIG=
{
  "backend": {"file": {"path": "/data"}},
  "listener": [{"tcp": {"address": "0.0.0.0:8200", "tls_disable": true}}],
  "plugin_directory": "/vault/plugins",
  "log_level": "debug",
  "disable_mlock": true,
  "api_addr": "http://localhost:8200"
}
' vault server
sleep 1

export VAULT_ADDR=http://localhost:8200

initoutput=$(vault operator init -key-shares=1 -key-threshold=1 -format=json)
vault operator unseal $(echo "$initoutput" | jq -r .unseal_keys_hex[0])

export VAULT_TOKEN=$(echo "$initoutput" | jq -r .root_token)

vault write sys/plugins/catalog/secret/vault-plugin-secrets-mongodbatlas \
    sha_256=$(shasum -a 256 vault-plugin-secrets-mongodbatlas | cut -d' ' -f1) \
    command="vault-plugin-secrets-mongodbatlas"

vault secrets enable \
    -path="mongodbatlas" \
    -plugin-name="vault-plugin-secrets-mongodbatlas" plugin