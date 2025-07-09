# GOAT Framework Security Implementation TODO

This document outlines security features that should be implemented in GOAT applications. Each section includes recommended packages and implementation code.

## TODO: Security Headers

### Add Secure Package

```bash
go get github.com/unrolled/secure
```

```go
// pkg/middleware/security.go - CREATE THIS FILE
package middleware

import (
    "github.com/unrolled/secure"
    "net/http"
)

func SecurityHeaders(isDevelopment bool) func(http.HandlerFunc) http.HandlerFunc {
    secureMiddleware := secure.New(secure.Options{
        AllowedHosts:          []string{}, // Add your domains in production
        SSLRedirect:           !isDevelopment,
        SSLHost:               "", // Your SSL host
        STSSeconds:            31536000,
        STSIncludeSubdomains:  true,
        FrameDeny:             true,
        ContentTypeNosniff:    true,
        BrowserXssFilter:     true,
        ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'",
        ReferrerPolicy:        "strict-origin-when-cross-origin",
    })

    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            err := secureMiddleware.Process(w, r)
            if err != nil {
                return
            }
            next(w, r)
        }
    }
}
```

## TODO: Rate Limiting

### Add Rate Limiting Package

```bash
go get golang.org/x/time/rate
```

```go
// pkg/middleware/ratelimit.go - CREATE THIS FILE
package middleware

import (
    "golang.org/x/time/rate"
    "net/http"
    "sync"
    "time"
)

type IPRateLimiter struct {
    ips map[string]*rate.Limiter
    mu  *sync.RWMutex
    r   rate.Limit
    b   int
}

func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
    i := &IPRateLimiter{
        ips: make(map[string]*rate.Limiter),
        mu:  &sync.RWMutex{},
        r:   r,
        b:   b,
    }
    return i
}

func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
    i.mu.Lock()
    defer i.mu.Unlock()

    limiter := rate.NewLimiter(i.r, i.b)
    i.ips[ip] = limiter
    return limiter
}

func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
    i.mu.Lock()
    limiter, exists := i.ips[ip]
    if !exists {
        i.mu.Unlock()
        return i.AddIP(ip)
    }
    i.mu.Unlock()
    return limiter
}

func RateLimit(limiter *IPRateLimiter) func(http.HandlerFunc) http.HandlerFunc {
    return func(next http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            limiter := limiter.GetLimiter(r.RemoteAddr)
            if !limiter.Allow() {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
            next(w, r)
        }
    }
}

// Cleanup old IPs periodically
func (i *IPRateLimiter) StartCleanup() {
    go func() {
        for {
            time.Sleep(time.Minute)
            i.mu.Lock()
            for ip, limiter := range i.ips {
                if limiter.Tokens() == float64(i.b) {
                    delete(i.ips, ip)
                }
            }
            i.mu.Unlock()
        }
    }()
}
```

## TODO: HTML Sanitization

### Add Bluemonday Package

```bash
go get github.com/microcosm-cc/bluemonday
```

```go
// pkg/sanitize/sanitize.go - CREATE THIS FILE
package sanitize

import (
    "github.com/microcosm-cc/bluemonday"
    "html"
)

var (
    strictPolicy = bluemonday.StrictPolicy()
    ugcPolicy    = bluemonday.UGCPolicy()
)

// For user-generated content (allows basic formatting)
func UGCContent(input string) string {
    return ugcPolicy.Sanitize(input)
}

// For strict sanitization (strips all HTML)
func StrictContent(input string) string {
    return strictPolicy.Sanitize(input)
}

// For basic HTML escaping
func EscapeHTML(input string) string {
    return html.EscapeString(input)
}

// Use in templ templates for trusted content
templ SafeUserContent(content string) {
    @templ.Raw(UGCContent(content))
}
```

## TODO: File Upload Security

### Add File Type Detection

```bash
go get github.com/h2non/filetype
```

```go
// pkg/upload/upload.go - CREATE THIS FILE
package upload

import (
    "errors"
    "io"
    "mime/multipart"
    "os"
    "path/filepath"
    "strings"

    "github.com/google/uuid"
    "github.com/h2non/filetype"
)

type FileUploader struct {
    MaxSize       int64
    AllowedTypes  []string
    UploadDir     string
}

func NewFileUploader(uploadDir string) *FileUploader {
    return &FileUploader{
        MaxSize: 5 * 1024 * 1024, // 5MB
        AllowedTypes: []string{
            "image/jpeg",
            "image/png",
            "image/gif",
            "image/webp",
        },
        UploadDir: uploadDir,
    }
}

func (fu *FileUploader) ValidateAndSave(file multipart.File, header *multipart.FileHeader) (string, error) {
    // Check file size
    if header.Size > fu.MaxSize {
        return "", errors.New("file too large")
    }

    // Read file content for type detection
    buffer := make([]byte, 512)
    _, err := file.Read(buffer)
    if err != nil {
        return "", err
    }

    // Detect actual file type
    kind, err := filetype.Match(buffer)
    if err != nil {
        return "", errors.New("unable to detect file type")
    }

    // Validate file type
    validType := false
    for _, allowedType := range fu.AllowedTypes {
        if kind.MIME.Value == allowedType {
            validType = true
            break
        }
    }
    if !validType {
        return "", errors.New("file type not allowed")
    }

    // Reset file pointer
    file.Seek(0, 0)

    // Generate secure filename
    secureFilename := uuid.New().String() + "." + kind.Extension

    // Create upload directory
    if err := os.MkdirAll(fu.UploadDir, 0755); err != nil {
        return "", err
    }

    // Create file path
    filePath := filepath.Join(fu.UploadDir, secureFilename)
    filePath = filepath.Clean(filePath)

    // Prevent path traversal
    if strings.Contains(filePath, "..") {
        return "", errors.New("invalid file path")
    }

    // Save file
    dst, err := os.Create(filePath)
    if err != nil {
        return "", err
    }
    defer dst.Close()

    _, err = io.Copy(dst, file)
    if err != nil {
        os.Remove(filePath)
        return "", err
    }

    return secureFilename, nil
}
```

## TODO: Password Strength Validation

### Add Password Validation Package

```bash
go get github.com/wagslane/go-password-validator
```

```go
// pkg/auth/password.go - CREATE THIS FILE
package auth

import (
    "errors"
    "strings"

    "github.com/wagslane/go-password-validator"
    "golang.org/x/crypto/bcrypt"
)

const (
    MinPasswordEntropy = 60 // Adjust based on your security requirements
)

var commonPasswords = []string{
    "password", "123456", "password123", "admin", "user", "login",
    "welcome", "monkey", "dragon", "master", "shadow", "qwerty",
}

func ValidatePassword(password string) error {
    // Check minimum entropy
    err := passwordvalidator.Validate(password, MinPasswordEntropy)
    if err != nil {
        return errors.New("password is too weak")
    }

    // Check against common passwords
    lowerPassword := strings.ToLower(password)
    for _, common := range commonPasswords {
        if strings.Contains(lowerPassword, common) {
            return errors.New("password contains common words")
        }
    }

    return nil
}

func HashPassword(password string) (string, error) {
    if err := ValidatePassword(password); err != nil {
        return "", err
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

func VerifyPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

## TODO: Structured Logging

### Add Structured Logging Package

```bash
go get github.com/rs/zerolog
```

```go
// pkg/security/logging.go - CREATE THIS FILE
package security

import (
    "os"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func init() {
    // Configure zerolog
    zerolog.TimeFieldFormat = time.RFC3339
    if os.Getenv("ENVIRONMENT") == "development" {
        log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    }
}

type SecurityLogger struct {
    logger zerolog.Logger
}

func NewSecurityLogger() *SecurityLogger {
    return &SecurityLogger{
        logger: log.With().Str("component", "security").Logger(),
    }
}

func (sl *SecurityLogger) LogFailedLogin(email, ip, userAgent string) {
    sl.logger.Warn().
        Str("event", "failed_login").
        Str("email", email).
        Str("ip", ip).
        Str("user_agent", userAgent).
        Msg("Failed login attempt")
}

func (sl *SecurityLogger) LogSuccessfulLogin(userID, ip, userAgent string) {
    sl.logger.Info().
        Str("event", "successful_login").
        Str("user_id", userID).
        Str("ip", ip).
        Str("user_agent", userAgent).
        Msg("Successful login")
}

func (sl *SecurityLogger) LogPasswordChange(userID, ip string) {
    sl.logger.Info().
        Str("event", "password_change").
        Str("user_id", userID).
        Str("ip", ip).
        Msg("Password changed")
}

func (sl *SecurityLogger) LogSuspiciousActivity(userID, ip, activity string) {
    sl.logger.Warn().
        Str("event", "suspicious_activity").
        Str("user_id", userID).
        Str("ip", ip).
        Str("activity", activity).
        Msg("Suspicious activity detected")
}

func (sl *SecurityLogger) LogCSRFViolation(ip, userAgent string) {
    sl.logger.Error().
        Str("event", "csrf_violation").
        Str("ip", ip).
        Str("user_agent", userAgent).
        Msg("CSRF token validation failed")
}

func (sl *SecurityLogger) LogRateLimitExceeded(ip string) {
    sl.logger.Warn().
        Str("event", "rate_limit_exceeded").
        Str("ip", ip).
        Msg("Rate limit exceeded")
}
```

## TODO: Environment Configuration

### Add Configuration Package

```bash
go get github.com/kelseyhightower/envconfig
```

```go
// pkg/config/security.go - CREATE THIS FILE
package config

import (
    "log"
    "time"

    "github.com/kelseyhightower/envconfig"
)

type SecurityConfig struct {
    Environment        string        `envconfig:"ENVIRONMENT" default:"development"`
    UseHTTPS          bool          `envconfig:"USE_HTTPS" default:"false"`
    SessionTimeout    time.Duration `envconfig:"SESSION_TIMEOUT" default:"24h"`
    RateLimitRPS      int           `envconfig:"RATE_LIMIT_RPS" default:"10"`
    RateLimitBurst    int           `envconfig:"RATE_LIMIT_BURST" default:"20"`
    MaxFileSize       int64         `envconfig:"MAX_FILE_SIZE_MB" default:"5"`
    UploadDir         string        `envconfig:"UPLOAD_DIR" default:"uploads"`
    PasswordEntropy   float64       `envconfig:"PASSWORD_ENTROPY" default:"60"`
    CSRFTokenLength   int           `envconfig:"CSRF_TOKEN_LENGTH" default:"32"`
}

func LoadSecurityConfig() *SecurityConfig {
    var config SecurityConfig
    err := envconfig.Process("", &config)
    if err != nil {
        log.Fatal("Failed to load security config:", err)
    }

    // Convert MB to bytes
    config.MaxFileSize = config.MaxFileSize * 1024 * 1024

    return &config
}

func (sc *SecurityConfig) IsProduction() bool {
    return sc.Environment == "production"
}

func (sc *SecurityConfig) GetCookieSettings() (secure bool, sameSite http.SameSite) {
    if sc.UseHTTPS {
        return true, http.SameSiteStrictMode
    }
    return false, http.SameSiteLaxMode
}
```

## TODO: Request ID Middleware

### Add Request ID Package

```bash
go get github.com/google/uuid
```

```go
// pkg/middleware/requestid.go - CREATE THIS FILE
package middleware

import (
    "context"
    "net/http"

    "github.com/google/uuid"
)

const RequestIDKey = "request_id"

func RequestID(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        requestID := r.Header.Get("X-Request-ID")
        if requestID == "" {
            requestID = uuid.New().String()
        }

        // Add to response header
        w.Header().Set("X-Request-ID", requestID)

        // Add to context
        ctx := context.WithValue(r.Context(), RequestIDKey, requestID)
        next(w, r.WithContext(ctx))
    }
}

func GetRequestID(ctx context.Context) string {
    if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
        return requestID
    }
    return ""
}
```

## TODO: Integration Example

### Main Application Setup

```go
// In your main.go - MODIFY EXISTING
package main

import (
    "golang.org/x/time/rate"

    "your-app/pkg/config"
    "your-app/pkg/middleware"
    "your-app/pkg/security"
    "your-app/pkg/upload"
)

func main() {
    // Load security configuration
    securityConfig := config.LoadSecurityConfig()

    // Initialize security logger
    securityLogger := security.NewSecurityLogger()

    // Initialize file uploader
    fileUploader := upload.NewFileUploader(securityConfig.UploadDir)

    // Initialize rate limiter
    rateLimiter := middleware.NewIPRateLimiter(
        rate.Limit(securityConfig.RateLimitRPS),
        securityConfig.RateLimitBurst,
    )
    rateLimiter.StartCleanup()

    // Setup router with security middlewares
    router := server.NewRouter()

    // Add security middlewares
    router.Use(middleware.RequestID)
    router.Use(middleware.SecurityHeaders(securityConfig.Environment == "development"))
    router.Use(middleware.ValidationMiddleware)
    router.Use(middleware.RateLimit(rateLimiter))

    // Your existing routes...

    // Start server
    server := server.NewServer(":8080", &router.Mux)
    server.Start()
}
```

## TODO: Environment Variables

### Create .env.example

```bash
# .env.example - CREATE THIS FILE
ENVIRONMENT=development
USE_HTTPS=false
SESSION_TIMEOUT=24h
RATE_LIMIT_RPS=10
RATE_LIMIT_BURST=20
MAX_FILE_SIZE_MB=5
UPLOAD_DIR=uploads
PASSWORD_ENTROPY=60
CSRF_TOKEN_LENGTH=32
```

## Implementation Priority

### HIGH Priority

1. **Security headers** (`github.com/unrolled/secure`)
2. **Input validation** (`github.com/go-playground/validator/v10`)
3. **Rate limiting** (`golang.org/x/time/rate`)
4. **Configuration** (`github.com/kelseyhightower/envconfig`)

### MEDIUM Priority

5. **Password validation** (`github.com/wagslane/go-password-validator`)
6. **File upload security** (`github.com/h2non/filetype`)
7. **Structured logging** (`github.com/rs/zerolog`)

### LOW Priority

8. **HTML sanitization** (`github.com/microcosm-cc/bluemonday`)
9. **Request ID tracking** (already have `github.com/google/uuid`)

## Package Summary

**Total new dependencies: 7 minimal packages**

- `github.com/unrolled/secure` - Security headers
- `github.com/go-playground/validator/v10` - Input validation
- `golang.org/x/time/rate` - Rate limiting
- `github.com/kelseyhightower/envconfig` - Configuration
- `github.com/wagslane/go-password-validator` - Password strength
- `github.com/h2non/filetype` - File type detection
- `github.com/rs/zerolog` - Structured logging
- `github.com/microcosm-cc/bluemonday` - HTML sanitization (optional)

All packages are well-maintained, lightweight, and commonly used in Go applications.
