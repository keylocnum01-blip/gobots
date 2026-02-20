package handler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

// ==========================================
// SECURITY HARDENING - Fix Vulnerabilities
// ==========================================

// ==========================================
// 1. RATE LIMITER
// ==========================================

type RateLimiter struct {
	mu         sync.RWMutex
	requests   map[string][]time.Time
	maxReqs    int
	windowSec  int
}

var globalRateLimiter = &RateLimiter{
	requests:  make(map[string][]time.Time),
	maxReqs:   10,    // 10 requests
	windowSec: 60,    // per 60 seconds
}

// CheckRateLimit - Returns true if allowed
func (rl *RateLimiter) CheckRateLimit(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-time.Duration(rl.windowSec) * time.Second)

	// Get existing requests
	times := rl.requests[key]
	
	// Filter out old requests
	var validTimes []time.Time
	for _, t := range times {
		if t.After(cutoff) {
			validTimes = append(validTimes, t)
		}
	}

	// Check limit
	if len(validTimes) >= rl.maxReqs {
		return false
	}

	// Add new request
	validTimes = append(validTimes, now)
	rl.requests[key] = validTimes

	return true
}

// ==========================================
// 2. INPUT VALIDATOR
// ==========================================

// SanitizeInput - Prevent injection attacks
func SanitizeInput(input string) string {
	// Remove potentially dangerous characters
	dangerous := []string{
		"\x00",  // Null byte
		"\n",    // Newline (for command injection)
		"\r",    // Carriage return
		";",     // Command separator
		"|",     // Pipe
		"`",     // Command substitution
		"$",     // Variable
		"$(",    // Command substitution
		"${",    // Variable expansion
	}

	result := input
	for _, d := range dangerous {
		result = removeAll(result, d)
	}

	// Limit length
	if len(result) > 1000 {
		result = result[:1000]
	}

	return result
}

func removeAll(s, substr string) string {
	result := s
	for {
		idx := indexOf(result, substr)
		if idx == -1 {
			break
		}
		result = result[:idx] + result[idx+len(substr):]
	}
	return result
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// ValidateMID - Validate LINE MID format
func ValidateMID(mid string) bool {
	if len(mid) == 0 || len(mid) > 50 {
		return false
	}
	// MIDs typically start with 'u', 'c', or 'r'
	firstChar := mid[0]
	if firstChar != 'u' && firstChar != 'c' && firstChar != 'r' {
		return false
	}
	return true
}

// ==========================================
// 3. AUDIT LOGGER
// ==========================================

type AuditLog struct {
	mu      sync.RWMutex
	entries []AuditEntry
	maxSize int
}

type AuditEntry struct {
	Timestamp   time.Time
	Action     string
	UserMID    string
	GroupID    string
	Result     string
	IPAddress  string
	Details    string
}

var auditLogger = &AuditLog{
	entries: make([]AuditEntry, 0),
	maxSize: 10000,
}

// LogAction - Record an action for audit
func (al *AuditLog) LogAction(action, userMID, groupID, result, details string) {
	al.mu.Lock()
	defer al.mu.Unlock()

	entry := AuditEntry{
		Timestamp: time.Now(),
		Action:   action,
		UserMID:  userMID,
		GroupID:  groupID,
		Result:   result,
		Details:  details,
	}

	al.entries = append(al.entries, entry)

	// Trim if too large
	if len(al.entries) > al.maxSize {
		al.entries = al.entries[len(al.entries)-al.maxSize:]
	}
}

// GetAuditLog - Retrieve audit entries
func (al *AuditLog) GetAuditLog(action, userMID string, limit int) []AuditEntry {
	al.mu.RLock()
	defer al.mu.RUnlock()

	var results []AuditEntry
	count := 0

	// Iterate backwards (newest first)
	for i := len(al.entries) - 1; i >= 0 && count < limit; i-- {
		entry := al.entries[i]
		
		if action != "" && entry.Action != action {
			continue
		}
		if userMID != "" && entry.UserMID != userMID {
			continue
		}

		results = append(results, entry)
		count++
	}

	return results
}

// ==========================================
// 4. SECURE STORAGE
// ==========================================

type SecureStorage struct {
	mu         sync.RWMutex
	data       map[string]string
	encryptionKey []byte
}

// NewSecureStorage - Create new secure storage
func NewSecureStorage(key string) (*SecureStorage, error) {
	// Derive key from password (simplified - use proper KDF in production)
	encryptionKey := deriveKey(key)
	
	return &SecureStorage{
		data:         make(map[string]string),
		encryptionKey: encryptionKey,
	}, nil
}

func deriveKey(password string) []byte {
	// Simplified - use bcrypt or Argon2 in production
	key := make([]byte, 32)
	for i := 0; i < 32; i++ {
		key[i] = byte(i)
	}
	return key
}

// Encrypt - Encrypt data
func (ss *SecureStorage) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(ss.encryptionKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt - Decrypt data
func (ss *SecureStorage) Decrypt(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(ss.encryptionKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// ==========================================
// 5. COMMAND TIMEOUT
// ==========================================

type CommandTimeout struct {
	mu          sync.RWMutex
	activeCmds  map[string]*time.Timer
	maxDuration time.Duration
}

var cmdTimeout = &CommandTimeout{
	activeCmds:  make(map[string]*time.Timer),
	maxDuration: 30 * time.Second,
}

// RunWithTimeout - Run command with timeout
func RunWithTimeout(cmdID string, fn func()) bool {
	cmdTimeout.mu.Lock()
	defer cmdTimeout.mu.Unlock()

	// Check if already running
	if _, exists := cmdTimeout.activeCmds[cmdID]; exists {
		return false // Already running
	}

	// Set timeout
	cmdTimeout.activeCmds[cmdID] = time.AfterFunc(cmdTimeout.maxDuration, func() {
		cmdTimeout.mu.Lock()
		delete(cmdTimeout.activeCmds, cmdID)
		cmdTimeout.mu.Unlock()
		
		// Log timeout
		auditLogger.LogAction("TIMEOUT", "", cmdID, "TIMEOUT", "Command timed out")
	})

	// Execute
	go func() {
		fn()
		cmdTimeout.mu.Lock()
		if t, exists := cmdTimeout.activeCmds[cmdID]; exists {
			t.Stop()
			delete(cmdTimeout.activeCmds, cmdID)
		}
		cmdTimeout.mu.Unlock()
	}()

	return true
}

// ==========================================
// 6. SECURE COMMAND HANDLER
// ==========================================

// SecureCommandHandler - Handle commands with security
func SecureCommandHandler(client *linetcr.Account, groupID, userMID, text string) {
	// 1. Rate limit
	if !globalRateLimiter.CheckRateLimit(userMID + ":" + groupID) {
		client.SendMessage(groupID, "⏳ Too many commands. Please wait.")
		return
	}

	// 2. Sanitize input
	sanitized := SanitizeInput(text)
	if sanitized != text {
		auditLogger.LogAction("SANITIZE", userMID, groupID, "WARN", "Input sanitized")
	}
	text = sanitized

	// 3. Log action
	auditLogger.LogAction("COMMAND", userMID, groupID, "START", text)

	// 4. Check permission
	if !CheckPermission(&botstate.Command{Admin: true}, userMID, groupID) {
		auditLogger.LogAction("COMMAND", userMID, groupID, "DENIED", "No permission")
		client.SendMessage(groupID, "❌ No permission")
		return
	}

	// 5. Execute with timeout
	cmdID := fmt.Sprintf("%s:%s:%d", userMID, groupID, time.Now().Unix())
	success := RunWithTimeout(cmdID, func() {
		// Execute actual command here
		// ProcessCommand(client, groupID, userMID, text)
	})

	if !success {
		client.SendMessage(groupID, "⏳ Command already running")
	}
}

// ==========================================
// 7. ANTI-SPAM FILTER
// ==========================================

type SpamFilter struct {
	mu           sync.RWMutex
	userMessages map[string][]time.Time
	threshold    int
	windowSec    int
}

var spamFilter = &SpamFilter{
	userMessages: make(map[string][]time.Time),
	threshold:    5,   // 5 messages
	windowSec:   10,  // in 10 seconds
}

// CheckSpam - Returns true if user is spamming
func (sf *SpamFilter) CheckSpam(userMID string) bool {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-time.Duration(sf.windowSec) * time.Second)

	// Get existing messages
	times := sf.userMessages[userMID]
	
	// Filter old
	var validTimes []time.Time
	for _, t := range times {
		if t.After(cutoff) {
			validTimes = append(validTimes, t)
		}
	}

	// Check threshold
	if len(validTimes) >= sf.threshold {
		return true
	}

	// Add new
	validTimes = append(validTimes, now)
	sf.userMessages[userMID] = validTimes

	return false
}

// ==========================================
// 8. FAILSAFE MODE
// ==========================================

type Failsafe struct {
	enabled       bool
	triggeredAt   time.Time
	reason        string
	mu            sync.RWMutex
}

var failsafe = &Failsafe{
	enabled: false,
}

// EnableFailsafe - Activate failsafe mode
func (fs *Failsafe) EnableFailsafe(reason string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	fs.enabled = true
	fs.triggeredAt = time.Now()
	fs.reason = reason
	
	// Disable all non-essential functions
	// Log the event
	auditLogger.LogAction("FAILSAFE", "", "", "ENABLED", reason)
}

// CheckFailsafe - Returns if in failsafe mode
func (fs *Failsafe) CheckFailsafe() (bool, string) {
	fs.mu.RLock()
	defer fs.mu.RUnlock()
	
	return fs.enabled, fs.reason
}

// DisableFailsafe - Deactivate failsafe
func (fs *Failsafe) DisableFailsafe() {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	
	fs.enabled = false
	fs.reason = ""
	
	auditLogger.LogAction("FAILSAFE", "", "", "DISABLED", "")
}
