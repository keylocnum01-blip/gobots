package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
)

// API Response structure
type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Bot status
var (
	botMutex   sync.Mutex
	botRunning bool
)

func main() {
	// API Routes
	http.HandleFunc("/api/status", handleStatus)
	http.HandleFunc("/api/kick", handleKick)
	http.HandleFunc("/api/ban", handleBan)
	http.HandleFunc("/api/unban", handleUnban)
	http.HandleFunc("/api/cancel", handleCancel)
	http.HandleFunc("/api/ginfo", handleGroupInfo)

	// Serve static files
	http.HandleFunc("/", handleIndex)

	log.Println("LINE Bot API Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"service": "LINE Bot API",
		"status":  "running",
		"version": "1.0.0",
	})
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	botMutex.Lock()
	defer botMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "ok",
		Message: "Bot status",
		Data: map[string]bool{
			"running": botRunning,
		},
	})
}

func handleKick(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		GroupID string `json:"group_id"`
		UserMID string `json:"user_mid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body")
		return
	}

	if req.GroupID == "" || req.UserMID == "" {
		respondError(w, "group_id and user_mid are required")
		return
	}

	// Execute kick via gobots (this would need to be implemented)
	result := executeCommand("kick", req.GroupID, req.UserMID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "ok",
		Message: "Kick command sent",
		Data:    result,
	})
}

func handleBan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		GroupID string `json:"group_id"`
		UserMID string `json:"user_mid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body")
		return
	}

	result := executeCommand("ban", req.GroupID, req.UserMID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "ok",
		Message: "Ban command sent",
		Data:    result,
	})
}

func handleUnban(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		GroupID string `json:"group_id"`
		UserMID string `json:"user_mid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body")
		return
	}

	result := executeCommand("unban", req.GroupID, req.UserMID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "ok",
		Message: "Unban command sent",
		Data:    result,
	})
}

func handleCancel(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		GroupID string `json:"group_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, "Invalid request body")
		return
	}

	result := executeCommand("cancel", req.GroupID, "")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "ok",
		Message: "Cancel command sent",
		Data:    result,
	})
}

func handleGroupInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	groupID := r.URL.Query().Get("group_id")
	if groupID == "" {
		respondError(w, "group_id is required")
		return
	}

	result := executeCommand("ginfo", groupID, "")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "ok",
		Message: "Group info retrieved",
		Data:    result,
	})
}

func executeCommand(cmd, groupID, userMID string) map[string]string {
	result := map[string]string{
		"command":  cmd,
		"group_id": groupID,
		"user_mid": userMID,
		"status":   "pending",
	}

	// This is a placeholder - actual implementation would communicate
	// with the gobots process via some IPC mechanism
	// For now, just log the command
	log.Printf("Command: %s | Group: %s | User: %s", cmd, groupID, userMID)

	return result
}

func respondError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(APIResponse{
		Status:  "error",
		Message: message,
	})
}
