# WebStack

A comprehensive Go-based web application framework that combines server-side rendering, client-side reactivity, and API services (REST/gRPC) in a single cohesive stack.

## Features

- **Multi-Protocol Support**
  - REST API Server (`/api/rest`)
  - gRPC Server (`/api/rpc`)
  - Web Application Server (`/api/web`)

- **Web Stack**
  - Server-Side Rendering with [templ](https://templ.guide/)
  - Enhanced interactivity with [HTMX](https://htmx.org/)
  - Lightweight client-side reactivity using [Petite Vue](https://github.com/vuejs/petite-vue) (6kb)

- **Embedable Applications**
  - Utilize Go's `embed` feature for serving static assets
  - Support for multiple web applications (e.g., main app, admin panel)

## Architecture

```mermaid
graph TD
    A[Main Application] --> B[API Layer]
    B --> C[REST API]
    B --> D[gRPC API]
    B --> E[Web Application]
    
    E --> F[Templates]
    E --> G[Static Assets]
    E --> H[Handlers]
    
    F --> I[templ Engine]
    G --> J[JavaScript/CSS]
    H --> K[HTTP Routes]
    
    subgraph "Embedded Applications"
        L[Admin Panel]
        M[Other Apps...]
    end
    
    A --> L
    A --> M
```

## Rendering Strategy

```mermaid
graph LR
    A[Web Request] --> B{Route Type}
    B -->|Full Page| C[Server-Side Rendering]
    B -->|Dynamic Update| D[HTMX Request]
    B -->|Client Interaction| E[Petite Vue]
    
    C --> F[templ Templates]
    D --> G[Partial HTML]
    E --> H[Reactive Components]
    
    subgraph "Server-Side"
        F
        G
    end
    
    subgraph "Client-Side"
        H -->|Local State| I[6kb Vue Runtime]
        H -->|User Events| J[DOM Updates]
    end
```

## Project Structure

```
webstack/
├── api/
│   ├── rest/    # REST API server
│   ├── rpc/     # gRPC server
│   ├── web/     # Main web application
│   └── admin/   # Admin panel (example additional web app)
```

## Quick Start

1. Clone the repository

2. Install dependencies:
   ```bash
   make setup
   ```
   This will install both Go and Node.js dependencies.

3. Build and run:
   ```bash
   # Build the application
   make build

   # Run the application
   make run
   ```

4. Development mode:
   ```bash
   # Watch JavaScript changes
   make js-watch

   # Run the application
   make run
   ```

## Documentation

For detailed development instructions and guidelines, see [COPILOT-INSTRUCTIONS.md](./.github/copilot-instructions.md).

## License

[Apache 2.0](./LICENSE)
