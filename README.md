# LinkinAutomator

A stealthy, Go-based LinkedIn automation tool using `go-rod`.

## Features (Assignment Status)

- [x] **Authentication** (Login, Session Persistence, Security Checks)
- [x] **Search & Targeting** (Keywords, Pagination, URL Extraction)
- [x] **Connection Requests** (Personalized Notes, Daily Limits)
- [x] **Messaging System** (Structure Implemented)
- [x] **Stealth & Anti-Bot**
    - [x] Human-like Mouse (Bezier Curves)
    - [x] Randomized Timings
    - [x] Fingerprint Masking
    - [x] Random Scrolling
    - [x] Realistic Typing
    - [x] Rate Limiting
- [x] **Configurable** (YAML/.env support)

## Setup

1. **Prerequisites**: 
   - Go 1.21+
   - Google Chrome installed.

2. **Installation**:
   ```bash
   git clone https://github.com/yourusername/linkin-automator.git
   cd linkin-automator
   go mod download
   ```

3. **Configuration**:
   Copy `.env.example` to `.env` and fill in your credentials:
   ```bash
   cp .env.example .env
   ```
   
   Edit `.env`:
   ```properties
   LINKEDIN_USERNAME=your_email
   LINKEDIN_PASSWORD=your_password
   HEADLESS=false  # Set to true for background operation
   ```

## Usage

Run the tool:
```bash
go run cmd/main.go
```

The tool will:
1. Log in (or use saved session).
2. Search for "Software Engineer" (default in `main.go`).
3. Collect top 10 profiles.
4. Visit each profile and send a connection request with a note.

## Structure

- `cmd/main.go`: Entry point.
- `internal/browser`: Rod wrapper with stealth.
- `internal/auth`: Login logic.
- `internal/search`: Search & scraping.
- `internal/connection`: Connection logic.
- `pkg/stealth`: Core stealth primitives (Mouse, Input, Scroll).

## Disclaimer

This tool is for educational purposes and Proof of Concept (POC) only. Automated interaction with LinkedIn violates their Terms of Service and can lead to account restriction or banning. Use at your own risk.
