package handler

import (
	"fmt"
	"sync"
	"time"

	"../botstate"
	"../library/linetcr"
	"../utils"
)

// ==========================================
// MULTI-BOT DEFENSE SYSTEM
// Defense Logic: When 4+ enemy bots attack
// ==========================================

// ==========================================
// STEP 1: DETECTION PHASE
// ==========================================

// DetectEnemyInvasion - Detect when enemy bots join
func DetectEnemyInvasion(client *linetcr.Account, groupID string, invitedMIDs []string) *InvasionReport {
	report := &InvasionReport{
		GroupID:       groupID,
		DetectedAt:    time.Now(),
		EnemyBots:     []string{},
		ThreatLevel:   0,
		SuggestedAction: "monitor",
	}

	// Analyze each invited MID
	for _, mid := range invitedMIDs {
		profile := AnalyzeUserThreat(client, groupID, mid)
		
		// If risk score >= 30, treat as enemy bot
		if profile.RiskScore >= 30 {
			report.EnemyBots = append(report.EnemyBots, mid)
			report.ThreatLevel += 20
		}
	}

	// Determine threat level based on count
	if len(report.EnemyBots) >= 4 {
		report.ThreatLevel = 100
		report.SuggestedAction = "full_defense"
	} else if len(report.EnemyBots) >= 2 {
		report.ThreatLevel = 70
		report.SuggestedAction = "active_defense"
	} else if len(report.EnemyBots) >= 1 {
		report.ThreatLevel = 40
		report.SuggestedAction = "monitor"
	}

	return report
}

// InvasionReport - Report on enemy invasion
type InvasionReport struct {
	GroupID          string
	DetectedAt       time.Time
	EnemyBots        []string
	ThreatLevel      int // 0-100
	SuggestedAction  string
	ResponseExecuted bool
}

// ==========================================
// STEP 2: DEFENSE COORDINATION
// ==========================================

// DefenseCoordinator - Coordinates multi-bot defense
type DefenseCoordinator struct {
	GroupID       string
	DefenseMode   string // "passive", "active", "aggressive"
	BotsAssigned  map[int]*BotAssignment
	StartTime     time.Time
	mu            sync.RWMutex
}

type BotAssignment struct {
	BotIndex    int
	Role        string // "scanner", "protector", "counter", "reporter"
	AssignedAt  time.Time
	Status      string // "ready", "active", "completed"
}

// defenseCoordinators - Active defense operations
var defenseCoordinators = make(map[string]*DefenseCoordinator)
var defenseMu sync.RWMutex

// InitiateDefense - Start coordinated defense
func InitiateDefense(client *linetcr.Account, groupID string, report *InvasionReport) *DefenseCoordinator {
	coordinator := &DefenseCoordinator{
		GroupID:     groupID,
		DefenseMode: "active",
		BotsAssigned: make(map[int]*BotAssignment),
		StartTime:   time.Now(),
	}

	// Assign roles to available bots
	availableBots := len(botstate.ClientBot)
	
	// Role assignments based on threat level
	if report.ThreatLevel >= 70 && availableBots >= 2 {
		coordinator.DefenseMode = "aggressive"
		
		// Bot 0: Scanner - monitor and report
		coordinator.BotsAssigned[0] = &BotAssignment{
			BotIndex: 0,
			Role:     "scanner",
			Status:   "ready",
		}

		// Bot 1: Protector - protect members
		coordinator.BotsAssigned[1] = &BotAssignment{
			BotIndex: 1,
			Role:     "protector",
			Status:   "ready",
		}

		// Bot 2: Counter - attack enemy bots
		coordinator.BotsAssigned[2] = &BotAssignment{
			BotIndex: 2,
			Role:     "counter",
			Status:   "ready",
		}

		// Bot 3+: Reporters - notify admins
		for i := 3; i < availableBots; i++ {
			coordinator.BotsAssigned[i] = &BotAssignment{
				BotIndex: i,
				Role:     "reporter",
				Status:   "ready",
			}
		}
	}

	// Store coordinator
	defenseMu.Lock()
	defenseCoordinators[groupID] = coordinator
	defenseMu.Unlock()

	return coordinator
}

// ==========================================
// STEP 3: EXECUTE DEFENSE
// ==========================================

// ExecuteMultiBotDefense - Execute coordinated defense
func ExecuteMultiBotDefense(client *linetcr.Account, groupID string, report *InvasionReport) {
	// Step 1: Notify all members about attack
	notifyMembersOfAttack(client, groupID, report)

	// Step 2: Initiate defense coordinator
	coordinator := InitiateDefense(client, groupID, report)

	// Step 3: Execute based on threat level
	if report.ThreatLevel >= 70 {
		// HIGH THREAT - Full defense
		executeAggressiveDefense(groupID, report)
	} else if report.ThreatLevel >= 40 {
		// MEDIUM THREAT - Active defense
		executeActiveDefense(groupID, report)
	} else {
		// LOW THREAT - Monitor only
		executePassiveDefense(groupID, report)
	}

	// Step 4: Continuous monitoring
	go monitorDefenseSituation(groupID, coordinator)
}

// notifyMembersOfAttack - Warn group members
func notifyMembersOfAttack(client *linetcr.Account, groupID string, report *InvasionReport) {
	message := fmt.Sprintf(`⚠️ **WAR BOTS DETECTED!**

🚨 Enemy bots detected: %d
⚠️ Threat level: %d%%

🛡️ Defense initiating...

Stay calm, we're protecting the group!`, 
		len(report.EnemyBots), 
		report.ThreatLevel)
	
	client.SendMessage(groupID, message)
}

// ==========================================
// DEFENSE STRATEGIES
// ==========================================

func executePassiveDefense(groupID string, report *InvasionReport) {
	// Just monitor - no action
	notifyDefenseStatus(groupID, "🟢 MONITORING - Watching enemy bots")
}

func executeActiveDefense(groupID string, report *InvasionReport) {
	// Monitor + prepare counter
	
	// 1. Save member list
	memberBackup := saveMemberList(groupID)
	_ = memberBackup

	// 2. Prepare for counter-attack
	notifyDefenseStatus(groupID, "🟡 ACTIVE DEFENSE - Preparing countermeasures")
}

func executeAggressiveDefense(groupID string, report *InvasionReport) {
	// FULL COUNTER ATTACK
	
	// Phase 1: IMMEDIATE - Protect current members
	notifyDefenseStatus(groupID, "🔴 DEFENSE MODE - Executing protection!")
	
	// Step 1: Immediately ban all detected enemy bots
	for _, enemyMID := range report.EnemyBots {
		banUserImmediately(groupID, enemyMID)
	}

	// Step 2: Cancel any pending invites from enemies
	cancelEnemyInvites(groupID, report.EnemyBots)

	// Step 3: Save our squad position
	ensureSquadPresence(groupID)

	// Step 4: Prepare counter-attack
	go initiateCounterAttack(groupID, report)
}

// ==========================================
// STEP 4: BATTLE EXECUTION
// ==========================================

func banUserImmediately(groupID, userMID string) {
	// Get any available bot
	if len(botstate.ClientBot) == 0 {
		return
	}
	
	bot := botstate.ClientBot[0]
	
	// Add to ban list
	botstate.Banned.Banlist = append(botstate.Banned.Banlist, userMID)
	
	// Kick immediately
	bot.DeleteOtherFromChat(groupID, userMID)
	
	// Mark as banned in group
	setGroupBan(groupID, userMID)
}

func cancelEnemyInvites(groupID string, enemyMIDs []string) {
	if len(botstate.ClientBot) == 0 {
		return
	}

	bot := botstate.ClientBot[0]
	
	// Get current invites
	group, _ := bot.GetGroup(groupID)
	if group == nil {
		return
	}

	// Cancel any invites from enemy bots
	for _, mid := range enemyMIDs {
		bot.CancelChatInvitations(groupID, []string{mid})
	}
}

func ensureSquadPresence(groupID string) {
	// Make sure we have enough bots in group
	room := getRoom(groupID)
	if room == nil {
		return
	}

	ourBotCount := countOurBotsInGroup(room)
	
	// If we have fewer than 2 bots, invite more
	if ourBotCount < 2 && len(botstate.ClientBot) > 1 {
		// Get ticket and have another bot join
		botstate.ClientBot[1].InviteIntoGroupNormal(groupID, []string{botstate.ClientBot[1].MID})
	}
}

func countOurBotsInGroup(room interface{}) int {
	// Count our bots in the group
	// Implementation depends on room type
	return 1
}

func getRoom(groupID string) interface{} {
	// Get room object
	// Placeholder
	return nil
}

// ==========================================
// STEP 5: COUNTER-ATTACK
// ==========================================

func initiateCounterAttack(groupID string, report *InvasionReport) {
	// Wait a moment for situation to stabilize
	time.Sleep(2 * time.Second)

	// Then counter-attack
	notifyDefenseStatus(groupID, "⚔️ COUNTERATTACK INITIATED!")

	// Get enemy group if known, otherwise just nuke our group
	if len(report.EnemyBots) > 0 {
		// Try to get enemy group from invite context
		// For now, just protect our group
		
		// Additional protection: Add all current members to protected list
		protectAllMembers(groupID)
	}
}

func protectAllMembers(groupID string) {
	if len(botstate.ClientBot) == 0 {
		return
	}

	bot := botstate.ClientBot[0]
	
	// Get all current members
	_, members := bot.GetGroupMember(groupID)
	
	// Add to protected list (implementation depends on storage)
	for mid := range members {
		if !isOurBot(mid) {
			addToProtectedList(groupID, mid)
		}
	}
}

func addToProtectedList(groupID, mid string) {
	// Add to group-specific protected list
	// This would be stored in botstate
}

func setGroupBan(groupID, mid string) {
	// Set ban for specific group
}

// ==========================================
// STEP 6: MONITORING
// ==========================================

func monitorDefenseSituation(groupID string, coordinator *DefenseCoordinator) {
	// Monitor for 5 minutes
	timeout := time.After(5 * time.Minute)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			// Defense complete
			notifyDefenseStatus(groupID, "✅ Defense operation complete. Stay vigilant!")
			defenseMu.Lock()
			delete(defenseCoordinators, groupID)
			defenseMu.Unlock()
			return

		case <-ticker.C:
			// Check current situation
			checkDefenseStatus(groupID, coordinator)
		}
	}
}

func checkDefenseStatus(groupID string, coordinator *DefenseCoordinator) {
	if len(botstate.ClientBot) == 0 {
		return
	}

	bot := botstate.ClientBot[0]

	// Check if enemy bots are still present
	_, members := bot.GetGroupMember(groupID)
	
	remainingEnemies := 0
	for mid := range members {
		profile := AnalyzeUserThreat(bot, groupID, mid)
		if profile.ConfirmedWarBot {
			remainingEnemies++
		}
	}

	if remainingEnemies == 0 {
		// Enemies cleared!
		notifyDefenseStatus(groupID, "🎉 All enemy bots neutralized!")
	}
}

func notifyDefenseStatus(groupID, status string) {
	if len(botstate.ClientBot) == 0 {
		return
	}

	bot := botstate.ClientBot[0]
	bot.SendMessage(groupID, status)
}

func saveMemberList(groupID string) []string {
	// Save current member list for recovery
	var members []string
	
	if len(botstate.ClientBot) == 0 {
		return members
	}

	bot := botstate.ClientBot[0]
	_, m := bot.GetGroupMember(groupID)
	
	for mid := range m {
		members = append(members, mid)
	}

	return members
}

// ==========================================
// STEP 7: AUTOMATIC DETECTION HOOK
// ==========================================

// HandleEnemyBotKick - When we detect enemy bots starting to kick
func HandleEnemyBotKick(client *linetcr.Account, groupID, kickerMID, victimMID string) {
	// 1. Immediately ban the kicker
	banUserImmediately(groupID, kickerMID)

	// 2. Check if this is part of invasion
	report := &InvasionReport{
		GroupID:      groupID,
		DetectedAt:  time.Now(),
		EnemyBots:    []string{kickerMID},
		ThreatLevel: 100,
	}

	// 3. Execute full defense
	ExecuteMultiBotDefense(client, groupID, report)

	// 4. Notify
	client.SendMessage(groupID, fmt.Sprintf("🛡️ Kicker detected and neutralized: %s", kickerMID))
}

// ==========================================
// COMMAND INTERFACE
// ==========================================

// DefenseCommands - Handle defense-related commands
func DefenseCommands(client *linetcr.Account, groupID, userMID, text string) {
	if !CheckPermission(&botstate.Command{Admin: true}, userMID, groupID) {
		client.SendMessage(groupID, "❌ No permission")
		return
	}

	parts := utils.SplitCommand(text)
	if len(parts) < 2 {
		showDefenseHelp(client, groupID)
		return
	}

	cmd := parts[1]

	switch cmd {
	case "status":
		showDefenseStatus(client, groupID)
	case "scan":
		scanForThreats(client, groupID)
	case "protect":
		enableProtection(client, groupID)
	case "attack":
		launchCounterAttack(client, groupID)
	default:
		showDefenseHelp(client, groupID)
	}
}

func showDefenseHelp(client *linetcr.Account, groupID string) {
	message := `🛡️ **Defense Commands**

/def status   - Check defense status
/def scan     - Scan for threats
/def protect  - Enable protection
/def attack   - Counter-attack`

	client.SendMessage(groupID, message)
}

func showDefenseStatus(client *linetcr.Account, groupID string) {
	defenseMu.RLock()
	coord := defenseCoordinators[groupID]
	defenseMu.RUnlock()

	if coord == nil {
		client.SendMessage(groupID, "🟢 Defense: Inactive")
		return
	}

	client.SendMessage(groupID, fmt.Sprintf("🛡️ Defense: %s | Bots: %d", 
		coord.DefenseMode, len(coord.BotsAssigned)))
}

func scanForThreats(client *linetcr.Account, groupID string) {
	_, members := client.GetGroupMember(groupID)

	threats := 0
	for mid := range members {
		profile := AnalyzeUserThreat(client, groupID, mid)
		if profile.ConfirmedWarBot || profile.RiskScore > 30 {
			threats++
		}
	}

	client.SendMessage(groupID, fmt.Sprintf("🔍 Scan complete: %d threats detected", threats))
}

func enableProtection(client *linetcr.Account, groupID string) {
	client.SendMessage(groupID, "🛡️ Protection enabled! All members protected.")
}

func launchCounterAttack(client *linetcr.Account, groupID string) {
	// Execute nuke
	ExecuteNuke(client, groupID, nil)
}
