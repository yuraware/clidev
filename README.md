# cli-builder

Turn any OpenAPI spec into a fully-featured CLI — no code generation required.

`cli-builder` reads an OpenAPI spec and produces a declarative `cli-schema.yaml` file. A companion runtime (`runner`) interprets that file at startup and builds a complete [Cobra](https://cobra.dev)-powered CLI: commands, subcommands, flags, auth, and table output — all driven by the schema.

```
OpenAPI spec (.oas.json)
       │
  builder generate          →   cli-schema.yaml   (human-editable)
                                       │
                               runner --form ...   →   $ acs apps list --limit 5
```

The `cli-schema.yaml` is the product. It is human-readable, version-controllable, and can be edited by hand to add auth, rename commands, restrict flags, or add commands that aren't in the spec.

---

## Quick start

```bash
# 1. Generate a cli-schema from an OpenAPI spec
go run ./cmd/builder generate \
  --spec sample-api/acs/openapi.oas.json \
  --out sample-api/acs/cli-schema.yaml

# 2. Run any command defined in the schema
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml --help
```

Or build standalone binaries first:

```bash
go build -o bin/builder ./cmd/builder
go build -o bin/runner  ./cmd/runner

./bin/builder generate --spec sample-api/acs/openapi.oas.json --out sample-api/acs/cli-schema.yaml
./bin/runner  --form sample-api/acs/cli-schema.yaml --help
```

---

## App Store Connect API example

The repo ships with the [App Store Connect API](https://developer.apple.com/documentation/appstoreconnectapi) spec (v4.3, 923 paths, 1 208 operations) as a sample.

### 1. Configure auth

Edit `sample-api/acs/cli-schema.yaml` and set the auth block (already done in this repo):

```yaml
auth:
  type: apple_jwt
  config:
    key_id:     { env: APP_STORE_KEY_ID }
    issuer_id:  { env: APP_STORE_ISSUER_ID }
    private_key: { env: APP_STORE_PRIVATE_KEY }
    audience:   { value: appstoreconnect-v1 }
```

Then export your credentials (from **App Store Connect → Users and Access → Keys**):

```bash
export APP_STORE_KEY_ID="XXXXXXXXXX"
export APP_STORE_ISSUER_ID="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
export APP_STORE_PRIVATE_KEY="$(cat ~/AuthKey_XXXXXXXXXX.p8)"
```

### 2. Run commands

```bash
# List your apps (table output)
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps list --limit 5

# Filter by name
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps list --filter-name "My App"

# Get a single app by ID
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps get <app-id>

# List builds for an app
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml builds list --filter-app <app-id> --limit 10

# List beta groups
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml betaGroups list

# JSON output
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps list --limit 3 -o json

# Select specific fields
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps list \
  --fields-apps "name,bundleId,sku" --limit 10

# List app events for a specific app
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps appEvents list <app-id>

# Explore all available resources
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml --help
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps --help
```

---

## cli-schema format

A `cli-schema.yaml` (or `.json`) file describes the entire CLI surface. Example:

```yaml
name: acs
version: "1.0"
description: App Store Connect CLI
base_url: https://api.appstoreconnect.apple.com

auth:
  type: apple_jwt           # none | bearer | api_key | apple_jwt
  config:
    key_id:      { env: APP_STORE_KEY_ID }
    issuer_id:   { env: APP_STORE_ISSUER_ID }
    private_key: { env: APP_STORE_PRIVATE_KEY }

output:
  default_format: table     # table | json

commands:
  apps:
    description: Manage your apps
    commands:
      list:
        description: List all apps
        action:
          method: GET
          path: /v1/apps
          operation_id: apps_getCollection
        parameters:
          - flag: --filter-name
            query: "filter[name]"
            type: string_array
          - flag: --limit
            query: limit
            type: integer
            default: 10
            max: 200
          - flag: --fields
            query: "fields[apps]"
            type: enum_array
            values: [name, bundleId, sku, primaryLocale]

      get:
        description: Get app by ID
        action:
          method: GET
          path: /v1/apps/{id}
          operation_id: apps_getInstance
        args:
          - name: id
            required: true

      create:
        description: Create a new app
        action:
          method: POST
          path: /v1/apps
          operation_id: apps_createInstance
        body:
          format: json_api
          resource_type: apps
          attributes:
            - flag: --bundle-id
              field: bundleId
              type: string
              required: true
            - flag: --name
              field: name
              type: string
              required: true
```

### Parameter types

| Type | Description | Example flag |
|---|---|---|
| `string` | Single string value | `--filter-name "My App"` |
| `string_array` | Comma-separated strings | `--filter-id "a,b,c"` |
| `integer` | Integer value | `--limit 20` |
| `boolean` | Boolean flag | `--exists-game-center-enabled-versions true` |
| `enum` | One of a fixed set | `--sort name` |
| `enum_array` | Comma-separated enum values | `--fields-apps "name,bundleId"` |

### Auth types

| Type | Required config keys |
|---|---|
| `none` | — |
| `bearer` | `token` |
| `api_key` | `key`, `header` (optional, default `X-API-Key`) |
| `apple_jwt` | `key_id`, `issuer_id`, `private_key`, `audience` |

Each config value can be supplied as an env var, a literal, or a file path:

```yaml
private_key: { env: MY_ENV_VAR }
private_key: { value: "literal-value" }
private_key: { file: /path/to/key.p8 }
```

---

## Project structure

```
cmd/
  builder/    builder generate | builder schema
  runner/     runner --form <cli-schema.yaml> <command...>
internal/
  oas/        OpenAPI 3.x spec parser
  cliSchema/  CLIForm types, loader, generator, JSON Schema emitter
  runtime/    Cobra command tree builder, HTTP executor, auth providers
  output/     Table and JSON response formatters
schema/
  cli-schema.schema.json   JSON Schema for cli-schema files (IDE support)
sample-api/
  acs/        App Store Connect API spec + generated cli-schema
```

---

## Adding a new API

```bash
# Any OpenAPI 3.x spec works
builder generate --spec path/to/your-api.json --out your-api-cli-schema.yaml

# Optionally rename the CLI and configure auth by editing the YAML, then run:
runner --form your-api-cli-schema.yaml --help
```
