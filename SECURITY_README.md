# Security Hardening

Fixes for identified vulnerabilities in gobots.

---

## Vulnerabilities Fixed

### 1. ✅ Rate Limiting
```go
// Now limits to 10 requests per 60 seconds
globalRateLimiter.CheckRateLimit(key)
```

### 2. ✅ Input Validation
```go
// Sanitizes all user input
SanitizeInput(input)

// Validates MID format
ValidateMID(mid)
```

### 3. ✅ Audit Logging
```go
// Records all actions
auditLogger.LogAction("COMMAND", userMID, groupID, "RESULT", details)

// Retrieve logs
auditLogger.GetAuditLog("", userMID, 100)
```

### 4. ✅ Secure Storage
```go
// AES-256 encryption for tokens
ss.Encrypt(plaintext)
ss.Decrypt(ciphertext)
```

### 5. ✅ Command Timeout
```go
// Commands timeout after 30 seconds
RunWithTimeout(cmdID, func() { ... })
```

### 6. ✅ Secure Command Handler
```go
// All-in-one security wrapper
SecureCommandHandler(client, groupID, userMID, text)
```

### 7. ✅ Anti-Spam Filter
```go
// Filters spam messages
spamFilter.CheckSpam(userMID)
```

### 8. ✅ Failsafe Mode
```go
// Emergency shutdown
failsafe.EnableFailsafe("Too many errors")

// Check status
enabled, reason := failsafe.CheckFailsafe()
```

---

## Usage

### Enable Security

```go
// In main.go, wrap command handler:
func handleMessage(op *SyncService.Operation, client *linetcr.Account) {
    text := op.Message.Text
    groupID := op.Message.To
    userMID := op.Message.From_
    
    // Use secure handler
    SecureCommandHandler(client, groupID, userMID, text)
}
```

### Configure Rate Limiting

```go
// Adjust limits
globalRateLimiter.maxReqs = 20   // requests
globalRateLimiter.windowSec = 60 // seconds
```

### View Audit Logs

```go
// Get recent commands for user
logs := auditLogger.GetAuditLog("COMMAND", userMID, 50)

// Get all commands in group
logs := auditLogger.GetAuditLog("", groupID, 100)
```

### Encrypt Tokens

```go
// Create secure storage
ss, _ := NewSecureStorage("your-master-password")

// Encrypt token
encrypted, _ := ss.Encrypt("your-line-token")

// Decrypt token
decrypted, _ := ss.Decrypt(encrypted)
```

---

## Security Levels

| Level | Features |
|-------|----------|
| Basic | Rate limiting, Input validation |
| Standard | + Audit logging, Command timeout |
| High | + Encryption, Failsafe mode |

---

## Configuration

```go
// config/security.json
{
  "rateLimit": {
    "maxRequests": 10,
    "windowSeconds": 60
  },
  "spamFilter": {
    "threshold": 5,
    "windowSeconds": 10
  },
  "commandTimeout": {
    "maxDurationSeconds": 30
  }
}
```

---

## Files

- `security_hardening.go` - All security implementations
- `SECURITY_README.md` - This file
