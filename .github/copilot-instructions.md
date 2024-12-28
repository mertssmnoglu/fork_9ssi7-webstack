# Development Guidelines

This document provides detailed instructions for developing applications using the WebStack framework.

## Architecture Overview

WebStack is designed to be a versatile Go-based framework that supports:
- Server-side rendering (SSR)
- Client-side reactivity
- REST API endpoints
- gRPC services

### Key Components

1. **Web Application (`/api/web`)**
   - Uses `templ` for server-side templating
   - HTMX for dynamic server interactions
   - Petite Vue (6kb) for client-side reactivity
   - Go's `embed` feature for static assets

2. **REST Server (`/api/rest`)**
   - Traditional REST API endpoints
   - JSON response handling
   - Middleware support

3. **gRPC Server (`/api/rpc`)**
   - Protocol buffer definitions
   - gRPC service implementations
   - Binary communication

## Creating a New Web Application

To create a new web application (e.g., admin panel):

1. Create a new directory under `/api`:
   ```bash
   mkdir -p api/admin
   ```

2. Create the necessary folder structure:
   ```
   api/admin/
   ├── handlers/
   ├── templates/
   ├── static/
   └── embed.go
   ```

3. Implement the `embed.go` file:
   ```go
   package admin

   import "embed"

   //go:embed all:static templates
   var content embed.FS
   ```

## Template Development

### Server-Side Components

Use `templ` for server-side components:

```go
// example.templ
package templates

templ UserProfile(user User) {
    <div class="profile">
        <h1>{user.Name}</h1>
        <div hx-get="/api/user/details" hx-trigger="load">
            Loading...
        </div>
    </div>
}
```

### Client-Side Components

Use Petite Vue for reactive components:

```html
<div v-scope="{ count: 0 }">
  <button @click="count++">Increment</button>
  <span>{{ count }}</span>
</div>
```

## Best Practices

1. **Component Organization**
   - Keep templates in `/templates` directory
   - Organize by feature/module
   - Use partial templates for reusability

2. **State Management**
   - Use Petite Vue for local state
   - Server state via HTMX requests
   - Keep client-side state minimal

3. **Performance**
   - Leverage SSR for initial load
   - Use HTMX for dynamic updates
   - Minimize client-side JavaScript

4. **Security**
   - Implement CSRF protection
   - Validate all inputs
   - Use secure headers

## Development Workflow

1. Start with server-side templates
2. Add HTMX interactions
3. Include Petite Vue where needed
4. Test thoroughly
5. Deploy using Go's build system

## Common Patterns

### Form Handling
```html
<form hx-post="/api/submit" hx-swap="outerHTML">
    <input type="text" name="username" v-model="form.username">
    <button type="submit">Submit</button>
</form>
```

### Dynamic Loading
```html
<div hx-get="/api/data" hx-trigger="revealed">
    Loading...
</div>
```

### Reactive Components
```html
<div v-scope="{ items: [] }" @mounted="fetchItems()">
    <ul>
        <li v-for="item in items">{{ item.name }}</li>
    </ul>
</div>
