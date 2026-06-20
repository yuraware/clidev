# Streetlights Kafka API

**CLI name:** `streetlights-kafka-api` · **Version:** 1.0.0
**Spec format:** AsyncAPI 2.x / 3.x · AsyncAPI 3.1.0 · 4 channels · 4 operations
**Base URL:** kafka-secure://test.mykafkacluster.org:18092

## Authentication

**Type:** `basic`

SASL/SCRAM — set KAFKA_USERNAME and KAFKA_PASSWORD

| Config key | Source |
|---|---|
| `password` | env `KAFKA_PASSWORD` |
| `username` | env `KAFKA_USERNAME` |

## Generate

```bash
go run ./cmd/builder generate --spec sample-api/streetlights/asyncapi.yml --out sample-api/streetlights/cli-schema.yaml
```

## Usage

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml --help
```

## Example commands

Inform about environmental lighting conditions of a particular streetlight.

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightingmeasured subscribe <streetlightId>
```

send on smartylighting.streetlights.1.0.action.{streetlightId}.dim

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightsdim publish <streetlightId>
```

send on smartylighting.streetlights.1.0.action.{streetlightId}.turn.off

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightturnoff publish <streetlightId>
```

send on smartylighting.streetlights.1.0.action.{streetlightId}.turn.on

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightturnon publish <streetlightId>
```

## Commands

### lightingmeasured

The topic on which measured values may be produced and consumed.

#### subscribe

Inform about environmental lighting conditions of a particular streetlight.

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightingmeasured subscribe <streetlightId>
```

### lightsdim

#### publish

send on smartylighting.streetlights.1.0.action.{streetlightId}.dim

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightsdim publish <streetlightId>
```

### lightturnoff

#### publish

send on smartylighting.streetlights.1.0.action.{streetlightId}.turn.off

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightturnoff publish <streetlightId>
```

### lightturnon

#### publish

send on smartylighting.streetlights.1.0.action.{streetlightId}.turn.on

```bash
go run ./cmd/runner --form sample-api/streetlights/cli-schema.yaml lightturnon publish <streetlightId>
```

