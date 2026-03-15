# ArcNote

ArcNote is managed as a single repository containing frontend, backend, and infrastructure code.

## Repository Layout

```text
/
  frontend/              # Next.js application
  backend/               # Go API and application modules
  infra/                 # IaC and deployment manifests
    terraform/
      modules/
      envs/
    kubernetes/
      base/
      overlays/
  docs/                  # Architecture and design documents
  scripts/               # Local automation scripts
```

## Intent

- `frontend/` owns the web UI and BFF-facing client code.
- `backend/` owns the Go API, domain, application, and adapters.
- `infra/` owns reproducible infrastructure and deployment definitions.
- `docs/` keeps architectural decisions close to the codebase.

## Backend Layout

```text
backend/
  cmd/api/
  internal/
    domain/
    application/
    port/
    adapter/
    infrastructure/
  migrations/
  test/
```

## Infrastructure Layout

```text
infra/
  terraform/
    modules/
    envs/
      local/
      dev/
      prod/
  kubernetes/
    base/
    overlays/
      dev/
      prod/
```
