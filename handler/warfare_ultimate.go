package handler

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"../botstate"
	"../library/linetcr"
	"../library/SyncService"
	"../utils"
)

// ==========================================
// ULTIMATE WARfare SYSTEM - Advanced Logic
// ==========================================

// WarConfig - Configuration for warfare
type WarConfig struct {
	AutoRejoin       bool // Auto rejoin when kicked
	CounterAttack    bool // Strike back when attacked
	MultiNuke        bool // Use multiple bots for nuke
	PsychologicalWar bool // Demoralize enemy
	DeepRecon        bool // Deep reconnaissance
	MaxNukePerWave   int  // Members to kick per wave
	NukeDelayMs      int  // Delay between kicks
}

var warConfig = &WarConfig{
	AutoRejoin:       true,
	CounterAttack:    true,
	MultiNuke:        true,
	PsychologicalWar: true,
	DeepRecon:        true,
	MaxNukePerWave:   10,
	NukeDelayMs:      100,
}

// WarIntelligence - Intelligence gathered about enemy
type WarIntelligence struct {
	GroupID      string
	GroupName    string
	EnemyBots    []string
	EnemyAdmins  []string
	SquadLeader  string
	MemberCount  int
	BotCount     int
	Weaknesses   []string
	Strengths    []string
	ThreatLevel  int // 1-10
	LastUpdate   time.Time
	AttackHistory []AttackRecord
}

type AttackRecord struct {
	Timestamp time.Time
	Type      string // "kick", "nuke", "spam", "invite_flood"
	Attacker  string
	Result    string // "success", "failed", "partial"
}

// WarRoom - Manages warfare for a group
type WarRoom struct {
	mu              sync.RWMutex
	GroupID         string
	IsWarMode       bool
	EnemyGroupID    string
	Intelligence    *WarIntelligence
	DefenseLevel    int // 1-5
	AttackLevel     int // 1-5
	LastAttack      time.Time
	SquadPositions  map[int]string // bot index -> position
}

var warRooms = make(map[string]*WarRoom)
var warRoomsMu sync.RWMutex

// ==========================================
// PHASE 1: EARLY WARNING & DETECTION
// ==========================================

// AdvancedWarBotDetector - Multi-layered detection
type AdvancedWarBotDetector struct {
	mu              sync.RWMutex
	SuspiciousUsers map[string]*WarBotProfile
}

type WarBotProfile struct {
	MID             string
	DisplayName     string
	GroupCount      int
	KickAttempts    int
	InviteFloods    int
	CancelAttempts  int
	AccountAgeDays  int
	RiskScore       int // 0-100
	Indicators      []string
	FirstSuspected  time.Time
	LastSuspected   time.Time
	ConfirmedWarBot bool
}

var advancedDetector = &AdvancedWarBotDetector{
	SuspiciousUsers: make(map[string]*WarBotProfile),
}

// AnalyzeUserThreat - Comprehensive threat analysis
func AnalyzeUserThreat(client *linetcr.Account, groupID, userMID string) *WarBotProfile {
	profile := &WarBotProfile{
		MID:            userMID,
		FirstSuspected: time.Now(),
		LastSuspected:  time.Now(),
	}

	// Get user contact
	contact, _ := client.GetContact(userMID)
	if contact != nil {
		profile.DisplayName = contact.DisplayName
	}

	// LAYER 1: Group Count Analysis
	groups, _ := client.GetGroupIdsJoined()
	profile.GroupCount = len(groups)
	if profile.GroupCount > 100 {
		profile.RiskScore += 40
		profile.Indicators = append(profile.Indicators, "excessive_groups")
	} else if profile.GroupCount > 50 {
		profile.RiskScore += 25
		profile.Indicators = append(profile.Indicators, "many_groups")
	}

	// LAYER 2: Behavior Pattern Analysis
	// Check recent operations from this user
	lastOps := getRecentOperations(userMID, 10*time.Minute)
	for _, op := range lastOps {
		switch op.Type {
		case 133: // Kick
			profile.KickAttempts++
			profile.RiskScore += 15
		case 124: // Invite
			profile.InviteFloods++
			profile.RiskScore += 5
		case 126: // Cancel
			profile.CancelAttempts++
			profile.RiskScore += 10
		}
	}

	// LAYER 3: Account Age (if detectable)
	// War bots often have new accounts
	if profile.AccountAgeDays < 7 {
		profile.RiskScore += 30
		profile.Indicators = append(profile.Indicators, "new_account")
	}

	// LAYER 4: Name Analysis
	suspiciousPatterns := []string{"war", "kick", "mod", "bot", "admin", "กู", "บอท", "วอร์", "nuke", "raid"}
	lowerName := strings.ToLower(profile.DisplayName)
	for _, pattern := range suspiciousPatterns {
		if strings.Contains(lowerName, pattern) {
			profile.RiskScore += 10
			profile.Indicators = append(profile.Indicators, "suspicious_name")
			break
		}
	}

	// LAYER 5: Check against known war bot database
	if isKnownWarBot(userMID) {
		profile.RiskScore += 50
		profile.ConfirmedWarBot = true
		profile.Indicators = append(profile.Indicators, "known_warbot_db")
	}

	// LAYER 6: Invite Pattern Detection
	// War bots often invite multiple accounts quickly
	if profile.InviteFloods > 5 {
		profile.RiskScore += 25
		profile.Indicators = append(profile.Indicators, "invite_flood")
	}

	// Determine if confirmed
	if profile.RiskScore >= 50 {
		profile.ConfirmedWarBot = true
	}

	// Store in detector
	advancedDetector.mu.Lock()
	advancedDetector.SuspiciousUsers[userMID] = profile
	advancedDetector.mu.Unlock()

	return profile
}

// getRecentOperations - Get operations from user in recent time
func getRecentOperations(userMID string, duration time.Duration) []*SyncService.Operation {
	// This would need implementation based on operation storage
	// Placeholder
	return []*SyncService.Operation{}
}

// isKnownWarBot - Check against known war bot database
func isKnownWarBot(mid string) bool {
	knownBots := []string{
		// Add known war bot MIDs here
	}
	for _, bot := range knownBots {
		if mid == bot {
			return true
		}
	}
	return false
}

// ==========================================
// PHASE 2: DEFENSE STRATEGIES
// ==========================================

// DefenseStrategy - Multi-layered defense
type DefenseStrategy int

const (
	DefensePassive DefenseStrategy = iota // Basic protection
	DefenseActive                        // Active counter-measures
	DefenseAggressive                    // Pre-emptive strikes
	DefenseTotal                         // Total war mode
)

// ActivateDefense - Activate appropriate defense level
func ActivateDefense(client *linetcr.Account, groupID string, strategy DefenseStrategy) {
	switch strategy {
	case DefensePassive:
		// Basic: just accept invites and stay
		client.SendMessage(groupID, "🛡️ Defense: PASSIVE mode")

	case DefenseActive:
		// Active: detect and counter
		activateActiveDefense(client, groupID)
		client.SendMessage(groupID, "🛡️ Defense: ACTIVE mode")

	case DefenseAggressive:
		// Aggressive: pre-emptive strikes
		activateAggressiveDefense(client, groupID)
		client.SendMessage(groupID, "🛡️ Defense: AGGRESSIVE mode")

	case DefenseTotal:
		// Total war: everything goes
		activateTotalDefense(client, groupID)
		client.SendMessage(groupID, "🛡️ Defense: TOTAL WAR mode")
	}
}

func activateActiveDefense(client *linetcr.Account, groupID string) {
	// 1. Monitor all invites
	// 2. Auto-kick detected war bots
	// 3. Quick rejoin capability
	go monitorInvites(client, groupID)
}

func activateAggressiveDefense(client *linetcr.Account, groupID string) {
	// 1. Everything in active
	// 2. Plus: preemptive kick of suspicious members
	// 3. Counter-invite to overwhelm enemy
	go monitorInvites(client, groupID)
	go preemptiveCounterMeasures(client, groupID)
}

func activateTotalDefense(client *linetcr.Account, groupID string) {
	// 1. Everything in aggressive
	// 2. All bots in squad report to group
	// 3. Prepare for nuke
	activateAggressiveDefense(client, groupID)

	// Get all squad bots
	for i, bot := range botstate.ClientBot {
		bot.SendMessage(groupID, fmt.Sprintf("⚔️ Squad Bot %d reporting for duty!", i+1))
	}
}

// monitorInvites - Monitor and intercept malicious invites
func monitorInvites(client *linetcr.Account, groupID string) {
	// Implementation would hook into invite events
}

// preemptiveCounterMeasures - Pre-emptive strikes
func preemptiveCounterMeasures(client *linetcr.Account, groupID string) {
	// Get all group members
	_, members := client.GetGroupMember(groupID)

	// Analyze each member
	for mid := range members {
		profile := AnalyzeUserThreat(client, groupID, mid)
		if profile.ConfirmedWarBot {
			// Pre-emptive kick
			client.DeleteOtherFromChat(groupID, mid)
		}
	}
}

// QuickRejoinAdvanced - Advanced rejoin with multiple strategies
func QuickRejoinAdvanced(client *linetcr.Account, groupID string) bool {
	strategies := []func(*linetcr.Account, string) bool{
		// Strategy 1: Ticket rejoin
		func(c *linetcr.Account, gid string) bool {
			ticket, _ := c.ReissueChatTicket(gid)
			if ticket != "" {
				time.Sleep(500 * time.Millisecond)
				return c.AcceptGroupInvitationByTicket(gid, ticket) == nil
			}
			return false
		},
		// Strategy 2: Re-invite from friend
		func(c *linetcr.Account, gid string) bool {
			// Get friend in group to reinvite
			_, members := c.GetGroupMember(gid)
			for mid := range members {
				if c.MID != mid && !isBot(mid) {
					// Ask friend to reinvite
					c.SendMessage(mid, fmt.Sprintf("Please reinvite me to: %s", gid))
					break
				}
			}
			return false // Will rejoin when reinvited
		},
		// Strategy 3: QR code scan
		func(c *linetcr.Account, gid string) bool {
			// Generate new QR
			c.UpdateChatQrV2(gid, true)
			return false // QR shown, need manual scan
		},
	}

	// Try each strategy
	for _, strategy := range strategies {
		if strategy(client, groupID) {
			return true
		}
		time.Sleep(1 * time.Second)
	}

	return false
}

// ==========================================
// PHASE 3: COUNTER-ATTACK STRATEGIES
// ==========================================

// CounterAttackPlan - Comprehensive counter-attack
func CounterAttackPlan(client *linetcr.Account, groupID, enemyGroupID string) {
	// Phase 1: Intelligence Gathering
	intel := gatherIntelligence(client, enemyGroupID)

	// Phase 2: Weakness Analysis
	analyzeWeaknesses(intel)

	// Phase 3: Execute Counter-Attack
	executeCounterAttack(client, enemyGroupID, intel)
}

// gatherIntelligence - Deep reconnaissance on enemy group
func gatherIntelligence(client *linetcr.Account, groupID string) *WarIntelligence {
	intel := &WarIntelligence{
		GroupID:      groupID,
		LastUpdate:   time.Now(),
		AttackHistory: []AttackRecord{},
	}

	// Basic info
	group, _ := client.GetGroup(groupID)
	if group != nil {
		intel.GroupName = group.Name
		intel.MemberCount = len(group.MemberMids)
	}

	// Get all members
	_, members := client.GetGroupMember(groupID)

	// Analyze each member
	botCount := 0
	for mid := range members {
		profile := AnalyzeUserThreat(client, groupID, mid)
		if profile.ConfirmedWarBot || profile.RiskScore > 30 {
			intel.EnemyBots = append(intel.EnemyBots, mid)
			botCount++
		}
	}
	intel.BotCount = botCount

	// Determine threat level
	if botCount > 5 {
		intel.ThreatLevel = 10
	} else if botCount > 3 {
		intel.ThreatLevel = 7
	} else if botCount > 1 {
		intel.ThreatLevel = 5
	} else {
		intel.ThreatLevel = 3
	}

	return intel
}

// analyzeWeaknesses - Find weaknesses in enemy group
func analyzeWeaknesses(intel *WarIntelligence) {
	// Weakness 1: Few bots
	if intel.BotCount < 3 {
		intel.Weaknesses = append(intel.Weaknesses, "few_defenders")
	}

	// Weakness 2: Large group (chaos opportunity)
	if intel.MemberCount > 100 {
		intel.Weaknesses = append(intel.Weaknesses, "large_group_chaos")
	}

	// Weakness 3: No apparent leadership
	if len(intel.EnemyAdmins) == 0 {
		intel.Weaknesses = append(intel.Weaknesses, "no_clear_leader")
	}

	// Strengths
	if intel.BotCount > 5 {
		intel.Strengths = append(intel.Strengths, "strong_defense")
	}
}

// executeCounterAttack - Execute the counter-attack plan
func executeCounterAttack(client *linetcr.Account, groupID string, intel *WarIntelligence) {
	// Determine attack strategy based on intel
	if intel.ThreatLevel >= 8 {
		// High threat: total nuke
		totalNuke(client, groupID, intel)
	} else if intel.ThreatLevel >= 5 {
		// Medium threat: targeted strikes
		targetedStrike(client, groupID, intel)
	} else {
		// Low threat: simple harassment
		harassment(client, groupID, intel)
	}
}

// ==========================================
// PHASE 4: DESTRUCTION - NUKE SYSTEM
// ==========================================

// NukeConfig - Configuration for group destruction
type NukeConfig struct {
	WaveSize        int     // Members per wave
	WaveDelayMs     int     // Delay between waves
	UseAllBots      bool    // Use all squad bots
	PrioritizeBots  bool    // Kick bots first
	Randomize       bool    // Random kick order
	LeaveLast       bool    // Leave last (don't trap yourself)
	NotifyBefore    bool    // Notify before nuke
	Message         string  // Custom nuke message
}

var defaultNukeConfig = &NukeConfig{
	WaveSize:       10,
	WaveDelayMs:    100,
	UseAllBots:     true,
	PrioritizeBots: true,
	Randomize:      true,
	LeaveLast:      true,
	NotifyBefore:   true,
	Message:        "☠️ NUKE INITIATED ☠️",
}

// ExecuteNuke - Multi-wave group destruction
func ExecuteNuke(client *linetcr.Account, groupID string, config *NukeConfig) {
	if config == nil {
		config = defaultNukeConfig
	}

	// Get all members
	_, members := client.GetGroupMember(groupID)

	// Filter out our bots
	var targets []string
	for mid := range members {
		if !isOurBot(mid) {
			targets = append(targets, mid)
		}
	}

	// Prioritize bots if enabled
	if config.PrioritizeBots {
		targets = prioritizeBotTargets(client, groupID, targets)
	}

	// Randomize if enabled
	if config.Randomize {
		rand.Shuffle(len(targets), func(i, j int) {
			targets[i], targets[j] = targets[j], targets[i]
		})
	}

	// Notify before nuke
	if config.NotifyBefore {
		client.SendMessage(groupID, config.Message)
		time.Sleep(2 * time.Second)
	}

	// Execute waves
	totalKicked := 0
	waveSize := config.WaveSize
	if config.UseAllBots && len(botstate.ClientBot) > 1 {
		// Use multiple bots for faster nuke
		executeMultiBotNuke(groupID, targets, waveSize, config.WaveDelayMs)
	} else {
		// Single bot nuke
		for i := 0; i < len(targets); i += waveSize {
			end := i + waveSize
			if end > len(targets) {
				end = len(targets)
			}

			wave := targets[i:end]
			for _, mid := range wave {
				client.DeleteOtherFromChat(groupID, mid)
				totalKicked++
			}

			// Delay between waves
			if end < len(targets) {
				time.Sleep(time.Duration(config.WaveDelayMs) * time.Millisecond)
			}
		}
	}

	// Final message
	client.SendMessage(groupID, fmt.Sprintf("💥 NUKE COMPLETE: %d members kicked", totalKicked))
}

// executeMultiBotNuke - Coordinated nuke from multiple bots
func executeMultiBotNuke(groupID string, targets []string, waveSize, delayMs int) {
	var wg sync.WaitGroup

	// Distribute targets among bots
	botCount := len(botstate.ClientBot)
	chunkSize := len(targets) / botCount
	if chunkSize == 0 {
		chunkSize = 1
	}

	for i, bot := range botstate.ClientBot {
		start := i * chunkSize
		end := start + chunkSize
		if i == botCount-1 {
			end = len(targets) // Last bot gets remainder
		}

		if start >= len(targets) {
			break
		}

		wg.Add(1)
		go func(b *linetcr.Account, t []string) {
			defer wg.Done()
			for _, mid := range t {
				b.DeleteOtherFromChat(groupID, mid)
				time.Sleep(time.Duration(delayMs) * time.Millisecond)
			}
		}(bot, targets[start:end])
	}

	wg.Wait()
}

// prioritizeBotTargets - Put bots first in kick order
func prioritizeBotTargets(client *linetcr.Account, groupID string, targets []string) []string {
	type scoredTarget struct {
		mid   string
		score int
	}

	var scored []scoredTarget

	for _, mid := range targets {
		profile := AnalyzeUserThreat(client, groupID, mid)
		score := profile.RiskScore

		// Bots get higher score (kick first)
		if profile.ConfirmedWarBot {
			score += 100
		}

		scored = append(scored, scoredTarget{mid: mid, score: score})
	}

	// Sort by score descending
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].score > scored[j].score
	})

	// Extract sorted targets
	result := make([]string, len(scored))
	for i, s := range scored {
		result[i] = s.mid
	}

	return result
}

// isOurBot - Check if MID belongs to our squad
func isOurBot(mid string) bool {
	for _, bot := range botstate.ClientBot {
		if bot.MID == mid {
			return true
		}
	}
	return false
}

// isBot - Check if MID is a bot (any bot)
func isBot(mid string) bool {
	// Check if it's our bot
	if isOurBot(mid) {
		return true
	}

	// Check against known bots
	// This would need more sophisticated detection
	return false
}

// ==========================================
// PHASE 5: PSYCHOLOGICAL WARFARE
// ==========================================

// PsychologicalWarfare - Demoralize enemy
func PsychologicalWarfare(client *linetcr.Account, groupID string) {
	messages := []string{
		"😈 Your group is about to be destroyed...",
		"💀 Any last words?",
		"🚀 Nuke incoming...",
		"⚠️ This is your final warning",
		"🏴‍☠️ Surrender or be destroyed",
	}

	for _, msg := range messages {
		client.SendMessage(groupID, msg)
		time.Sleep(3 * time.Second)
	}
}

// RaidWithMessage - Nuke with demoralizing messages
func RaidWithMessage(client *linetcr.Account, groupID string, message string) {
	client.SendMessage(groupID, message)
	time.Sleep(2 * time.Second)
	ExecuteNuke(client, groupID, nil)
}

// ==========================================
// PHASE 6: COMMAND HANDLERS
// ==========================================

// WarCommands - Handle war-related commands
func WarCommands(client *linetcr.Account, groupID, userMID, text string) {
	if !CheckPermission(&botstate.Command{Admin: true}, userMID, groupID) {
		client.SendMessage(groupID, "❌ No permission")
		return
	}

	parts := strings.Fields(text)
	if len(parts) < 2 {
		showWarHelp(client, groupID)
		return
	}

	command := strings.ToLower(parts[1])

	switch command {
	case "scan":
		// Scan group for threats
		intel := gatherIntelligence(client, groupID)
		client.SendMessage(groupID, fmt.Sprintf("📊 Scan Complete:\n👥 Members: %d\n🤖 Bots: %d\n⚠️ Threat: %d/10",
			intel.MemberCount, intel.BotCount, intel.ThreatLevel))

	case "defense":
		// Set defense level
		if len(parts) < 3 {
			client.SendMessage(groupID, "Usage: /war defense <passive|active|aggressive|total>")
			return
		}
		level := strings.ToLower(parts[2])
		var strategy DefenseStrategy
		switch level {
		case "passive":
			strategy = DefensePassive
		case "active":
			strategy = DefenseActive
		case "aggressive":
			strategy = DefenseAggressive
		case "total":
			strategy = DefenseTotal
		default:
			client.SendMessage(groupID, "Invalid level")
			return
		}
		ActivateDefense(client, groupID, strategy)

	case "nuke":
		// Execute nuke
		ExecuteNuke(client, groupID, nil)

	case "raid":
		// Nuke with message
		msg := "💥 RAID INITIATED!"
		if len(parts) > 2 {
			msg = strings.Join(parts[2:], " ")
		}
		RaidWithMessage(client, groupID, msg)

	case "counter":
		// Counter-attack
		if len(parts) < 3 {
			client.SendMessage(groupID, "Usage: /war counter <group_id>")
			return
		}
		enemyGroup := parts[2]
		CounterAttackPlan(client, groupID, enemyGroup)

	case "rejoin":
		// Quick rejoin
		if QuickRejoinAdvanced(client, groupID) {
			client.SendMessage(groupID, "✅ Successfully rejoined!")
		} else {
			client.SendMessage(groupID, "❌ Could not rejoin")
		}

	case "status":
		// War status
		warRoomsMu.RLock()
		room := warRooms[groupID]
		warRoomsMu.RUnlock()

		status := "🟢 PEACE"
		if room != nil && room.IsWarMode {
			status = "🔴 WAR MODE"
		}

		client.SendMessage(groupID, fmt.Sprintf("⚔️ Status: %s\n🤖 Squad: %d bots",
			status, len(botstate.ClientBot)))

	default:
		showWarHelp(client, groupID)
	}
}

func showWarHelp(client *linetcr.Account, groupID string) {
	message := `⚔️ **WARfare Commands**

📊 /war scan - Scan group threats
🛡️ /war defense <mode> - Set defense
💥 /war nuke - Destroy group
🚀 /war raid <msg> - Nuke with message
⚔️ /war counter <group> - Counter-attack
🔄 /war rejoin - Quick rejoin
📡 /war status - War status

Defense Modes:
- passive  = Basic protection
- active   = Auto counter
- aggressive = Pre-emptive
- total    = Full war mode`

	client.SendMessage(groupID, message)
}

// ==========================================
// PHASE 7: INVITE WARFARE
// ==========================================

// InviteWarfare - Flood enemy group with invites
func InviteWarfare(enemyGroupID, targetMID string, count int) {
	// Get all our bots
	bots := botstate.ClientBot

	for i := 0; i < count; i++ {
		botIndex := i % len(bots)
		bot := bots[botIndex]

		// Invite the target
		bot.InviteIntoGroupNormal(enemyGroupID, []string{targetMID})

		time.Sleep(500 * time.Millisecond)
	}
}

// CancelWave - Wave of cancels to disrupt enemy
func CancelWave(client *linetcr.Account, groupID string) {
	// Get pending invites
	group, _ := client.GetGroup(groupID)

	if group == nil || len(group.InviteeMids) == 0 {
		return
	}

	// Cancel all
	for _, mid := range group.InviteeMids {
		client.CancelChatInvitations(groupID, []string{mid})
	}
}

// ==========================================
// PHASE 8: SCOUT & EVADE
// ==========================================

// ScoutEnemy - Gather intel on enemy group
func ScoutEnemy(client *linetcr.Account, groupID string) *WarIntelligence {
	return gatherIntelligence(client, groupID)
}

// EvadeKick - Advanced kick evasion
func EvadeKick(client *linetcr.Account, groupID string) {
	// Strategy 1: Leave before kick completes
	go func() {
		client.LeaveGroup(groupID)
		time.Sleep(100 * time.Millisecond)
		QuickRejoinAdvanced(client, groupID)
	}()

	// Strategy 2: Spread across multiple groups (confuse enemy)
	// This would require coordination with squad
}
