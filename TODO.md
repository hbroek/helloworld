# Project TODO List

## 🚨 Critical Issues (High Priority)

- [ ] **Add CSS styling** to `frontend/www/`
  - Current: Purely inline styles throughout HTML
  - Suggested: Extract to external CSS file
  - Include responsive design
  - Add hover states and animations

- [ ] **Add CORS configuration** for cross-origin requests
  - Currently: No CORS middleware configured
  - Risk: Browser security policy may block cross-origin requests

- [ ] **Add input validation** to prevent XSS vulnerabilities
  - Validate/sanitize user input on both frontend and backend

## 🟠 High Priority

- [ ] **Add error logging** with structured logging
  - Currently: Minimal error output
  - Suggested: Add request logging, structured error messages

- [ ] **Implement rate limiting** for API endpoints
  - `/api/v1/name-generator` endpoint needs rate limiting
  - Prevent API abuse

- [ ] **Graceful shutdown handling**
  - Add signal handlers for SIGTERM/SIGINT
  - Allow proper cleanup before shutdown

## 🟡 Medium Priority

- [ ] **Check `dist/` directory contents**
  - Directory exists but appears empty
  - Determine purpose: builds? deployments?

- [ ] **Create README.md**
  - Setup instructions
  - API documentation
  - Development guide
  - Project architecture

- [ ] **Add .gitignore**
  - Go workspace cache (`/tmp`)
  - Build artifacts
  - Environment variables
  - `dist/` contents if auto-generated

- [ ] **Dockerize application**
  - Create Dockerfile
  - Define docker-compose.yml for local development

- [ ] **CI/CD pipeline**
  - GitHub Actions or similar
  - Automated testing on push
  - Linting and code quality checks
  - Build pipeline

- [ ] **Add health check endpoint visibility**
  - Consider adding metrics or status page

## 🟢 Low Priority (Nice to Have)

- [ ] **Add documentation comments**
  - Doc comments for exported functions
  - Architecture documentation

- [ ] **Add environment variables example**
  - `.env.example` file
  - Document required/optional env vars

- [ ] **Add security headers**
  - Content-Security-Policy
  - X-Content-Type-Options
  - X-Frame-Options
  - X-XSS-Protection

- [ ] **Add Prometheus/OpenTelemetry metrics**
  - Request counts
  - Response times
  - Error rates

- [ ] **Add Swagger/OpenAPI documentation**
  - Auto-generated API docs
  - Request/response schemas

- [ ] **Add user preferences/local storage**
  - Remember user's preferred name option
  - Local storage for settings

- [ ] **Add analytics tracking** (optional)
  - Privacy-compliant tracking
  - Usage statistics

---

## Testing Status

- [ ] Run `go test ./...` to verify tests pass
- [ ] Run all tests and fix any failures
- [ ] Add integration tests
- [ ] Add performance tests

---

## Notes

- Current build binary: `frontend_server` (8.9MB) is pre-built
- Go version: 1.26
- Project module: `continue-test`
- Default port: 8080 (configurable via environment variable)
- Static files served from: `frontend/www`
