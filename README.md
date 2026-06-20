# cli-builder

Turn any API spec into a fully-featured CLI — no code generation required.

`cli-builder` reads an API spec in any supported format and produces a declarative `cli-schema.yaml` file. A companion runtime (`runner`) interprets that file at startup and builds a complete [Cobra](https://cobra.dev)-powered CLI: commands, subcommands, flags, auth, and table output — all driven by the schema.

```
API spec  (.proto / .graphql / .yaml / .json)
       │
  builder generate          →   cli-schema.yaml   (human-editable)
                               cli-schema.md     (description + example commands)
                                       │
                               runner --form ...   →   $ acs apps list --limit 5
```

The `cli-schema.yaml` is the product. It is human-readable, version-controllable, and can be edited by hand to add auth, rename commands, restrict flags, or add commands that aren't in the spec.

### Supported input formats

| Format | Extensions | Auth styles commonly used |
|---|---|---|
| **OpenAPI 3.x / Swagger** | `.json` `.yaml` `.yml` | Bearer, API key, OAuth 2.0, Apple JWT |
| **gRPC / Protocol Buffers** | `.proto` | OAuth 2.0, Bearer |
| **GraphQL SDL** | `.graphql` `.gql` | Bearer, API key |
| **AsyncAPI 2.x / 3.x** | `.yaml` `.json` | Basic, API key, Bearer, mTLS |

---

## Quick start

```bash
# 1. Generate a cli-schema from any API spec (format auto-detected)
#    Also writes cli-schema.schema.yaml and cli-schema.md alongside the output.
go run ./cmd/builder generate --spec <path-to-spec> --out cli-schema.yaml

# 2. Run any command defined in the schema
go run ./cmd/runner --form cli-schema.yaml --help
```

Or build standalone binaries first:

```bash
go build -o bin/builder ./cmd/builder
go build -o bin/runner  ./cmd/runner
```

---

## Sample APIs

The repo ships four sample APIs covering all supported formats.

---

### App Store Connect API — OpenAPI

**What it is:** Apple's official REST API for managing apps, builds, TestFlight, in-app purchases, and App Store submissions. Used by developers and CI/CD pipelines to automate App Store workflows without the browser UI.

**Spec:** OpenAPI 3.0.1 · 923 paths · 1 208 operations  
**Auth:** Apple-signed ES256 JWT (`apple_jwt`)  
**Source:** [developer.apple.com/documentation/appstoreconnectapi](https://developer.apple.com/documentation/appstoreconnectapi)

```bash
go run ./cmd/builder generate \
  --spec sample-api/acs/openapi.oas.json \
  --out  sample-api/acs/cli-schema.yaml
```

Set credentials (from **App Store Connect → Users and Access → Keys**):

```bash
export APP_STORE_KEY_ID="XXXXXXXXXX"
export APP_STORE_ISSUER_ID="xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
export APP_STORE_PRIVATE_KEY="$(cat ~/AuthKey_XXXXXXXXXX.p8)"
```

```bash
# List your apps
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps list --limit 5

# Filter by name, select specific fields
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps list \
  --filter-name "My App" --fields-apps "name,bundleId,sku"

# Get a single app by ID
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps get <app-id>

# List builds for an app
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml builds list \
  --filter-app <app-id> --limit 10

# List beta groups
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml betaGroups list

# List app events for a specific app
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps appEvents list <app-id>

# JSON output
go run ./cmd/runner --form sample-api/acs/cli-schema.yaml apps list -o json
```

---

### Google Cloud Pub/Sub — gRPC / Protocol Buffers

**What it is:** Google Cloud's managed publish-subscribe messaging service. Publishers send messages to topics; subscribers receive them via subscriptions. Widely used for event-driven architectures, stream processing, and decoupling microservices at scale.

**Spec:** Proto3 · 2 services (`Publisher`, `Subscriber`) · 20+ RPCs with HTTP transcoding annotations  
**Auth:** OAuth 2.0 Bearer token (`bearer`) — obtain via `gcloud auth print-access-token`  
**Source:** [googleapis/googleapis — google/pubsub/v1/pubsub.proto](https://github.com/googleapis/googleapis/blob/master/google/pubsub/v1/pubsub.proto)

```bash
go run ./cmd/builder generate \
  --spec sample-api/pubsub/pubsub.proto \
  --out  sample-api/pubsub/cli-schema.yaml
```

```bash
export GOOGLE_ACCESS_TOKEN="$(gcloud auth print-access-token)"
export PROJECT="projects/my-gcp-project"

# Create a topic
go run ./cmd/runner --form sample-api/pubsub/cli-schema.yaml \
  publisher create "$PROJECT/topics/my-topic"

# List topics in a project
go run ./cmd/runner --form sample-api/pubsub/cli-schema.yaml \
  publisher list "$PROJECT"

# Publish a message
go run ./cmd/runner --form sample-api/pubsub/cli-schema.yaml \
  publisher publish "$PROJECT/topics/my-topic"

# Create a subscription
go run ./cmd/runner --form sample-api/pubsub/cli-schema.yaml \
  subscriber create "$PROJECT/subscriptions/my-sub"

# Pull messages from a subscription
go run ./cmd/runner --form sample-api/pubsub/cli-schema.yaml \
  subscriber pull "$PROJECT/subscriptions/my-sub"
```

---

### GitHub API — GraphQL SDL

**What it is:** GitHub's GraphQL API (v4) for querying and mutating repository data: users, repositories, issues, pull requests, and more. GraphQL lets clients specify exactly which fields to return, reducing over-fetching compared to REST.

**Spec:** GraphQL SDL · 7 queries · 4 mutations  
**Auth:** Personal Access Token as `Bearer` (`bearer`) — create at github.com/settings/tokens  
**Source:** [docs.github.com/en/graphql](https://docs.github.com/en/graphql)

```bash
go run ./cmd/builder generate \
  --spec sample-api/github-gql/schema.graphql \
  --out  sample-api/github-gql/cli-schema.yaml
```

Update `base_url` in the generated schema:

```yaml
base_url: https://api.github.com
```

```bash
export GRAPHQL_TOKEN="ghp_xxxxxxxxxxxxxxxxxxxx"

# Look up a user
go run ./cmd/runner --form sample-api/github-gql/cli-schema.yaml \
  query user octocat

# Get a repository
go run ./cmd/runner --form sample-api/github-gql/cli-schema.yaml \
  query repository octocat hello-world

# Look up an organization
go run ./cmd/runner --form sample-api/github-gql/cli-schema.yaml \
  query organization github

# Search repositories
go run ./cmd/runner --form sample-api/github-gql/cli-schema.yaml \
  query search --query "language:go stars:>1000" --type REPOSITORY

# Create an issue
go run ./cmd/runner --form sample-api/github-gql/cli-schema.yaml \
  mutate create-issue --repository-id <id> --title "Bug: ..."

# Get authenticated user info
go run ./cmd/runner --form sample-api/github-gql/cli-schema.yaml \
  query viewer
```

---

### Smartylighting Streetlights — AsyncAPI

**What it is:** The canonical AsyncAPI example API — a smart city streetlight management system that uses Apache Kafka for event-driven communication. Streetlights publish lighting measurements (lux levels); a control system sends commands to turn lights on/off or dim them. Demonstrates the publish/subscribe pattern over a message broker rather than REST.

**Spec:** AsyncAPI 3.1.0 · 4 channels · 4 operations · SASL/SCRAM + mTLS security  
**Auth:** SASL/SCRAM (`basic`) for the Kafka broker  
**Source:** [asyncapi/spec — examples/streetlights-kafka-asyncapi.yml](https://github.com/asyncapi/spec/blob/master/examples/streetlights-kafka-asyncapi.yml)

```bash
go run ./cmd/builder generate \
  --spec sample-api/streetlights/asyncapi.yml \
  --out  sample-api/streetlights/cli-schema.yaml
```

```bash
export KAFKA_USERNAME="my-user"
export KAFKA_PASSWORD="my-password"

# Subscribe to lighting measurements from a streetlight
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml \
  lightingmeasured subscribe

# Publish a command to turn on a streetlight
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml \
  lightturnon publish

# Publish a command to dim a streetlight
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml \
  lightsdim publish --lumens 50

# Subscribe to turn-off events
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml \
  lightturnoff subscribe
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
  type: apple_jwt           # none | bearer | api_key | basic | oauth2 | apple_jwt
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
| `boolean` | Boolean flag | `--is-draft true` |
| `enum` | One of a fixed set | `--sort name` |
| `enum_array` | Comma-separated enum values | `--fields-apps "name,bundleId"` |

### Auth types

| Type | Use case | Required config keys |
|---|---|---|
| `none` | Public APIs | — |
| `bearer` | GitHub PAT, generic tokens | `token` |
| `api_key` | Header or query-param key | `key`, `header` (default `X-API-Key`), `in` (`header`\|`query`) |
| `basic` | Username + password, SASL | `username`, `password` |
| `oauth2` | Machine-to-machine, Google APIs | `client_id`, `client_secret`, `token_url`, `scopes` |
| `apple_jwt` | App Store Connect API | `key_id`, `issuer_id`, `private_key`, `audience` |

Each config value resolves from an env var, a literal, or a file:

```yaml
private_key: { env: MY_ENV_VAR }
private_key: { value: "literal-value" }
private_key: { file: /path/to/key.p8 }
```

---

## Project structure

```
cmd/
  builder/      builder generate | builder schema
  runner/       runner --form <cli-schema.yaml> <command...>
internal/
  formats/      Format auto-detection
  oas/          OpenAPI 3.x / Swagger parser
  proto/        Protocol Buffers (.proto) parser
  graphql/      GraphQL SDL parser
  asyncapi/     AsyncAPI 2.x / 3.x parser
  cliSchema/    CLIForm types, loader, generator, JSON Schema emitter
  runtime/      Cobra command tree builder, HTTP executor, auth providers
  output/       Table and JSON response formatters
schema/
  cli-schema.schema.json    JSON Schema for cli-schema files (IDE support)
sample-api/
  acs/            Apple App Store Connect  (OpenAPI)
  pubsub/         Google Cloud Pub/Sub     (gRPC / Proto3)
  github-gql/     GitHub API               (GraphQL SDL)
  streetlights/   Smartylighting           (AsyncAPI 3.x)
```

---

## Adding a new API

```bash
# Format is auto-detected from file extension and content
go run ./cmd/builder generate --spec path/to/your-api.proto --out my-api/cli-schema.yaml

# Force a specific format if needed
go run ./cmd/builder generate --spec api.yaml --format openapi --out cli-schema.yaml

# Explore the generated CLI
go run ./cmd/runner --form my-api/cli-schema.yaml --help
```
