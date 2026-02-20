package handler

import (
	"fmt"
	"sync"
	"time"

	"../botstate"
	"../library/linetcr"
	"../utils"
)

// AntiKickBotProtection - Protect against other bots trying to kick our bots
type AntiKickBotProtection struct {
	mu           sync.RWMutex
	detectedBots map[string][]string // groupID -> list of bot MIDs detected
	blockedBots  map[string]bool    // botMID -> is blocked
	lastCheck    map[string]int64    // groupID -> last check timestamp
}

var (
	antiKickProtection = &AntiKickBotProtection{
		detectedBots: make(map[string][]string),
		blockedBots:  make(map[string]bool),
		lastCheck:    make(map[string]int64),
	}
)

// BotBehavior - Detect suspicious bot behavior
type BotBehavior struct {
	MID           string
	DisplayName   string
	KickCount     int
	InviteCount   int
	CancelCount   int
	FirstSeen     time.Time
	LastSeen      time.Time
	IsWarBot      bool
	AttackScore   int // 0-100, higher = more likely war bot
}

// CheckIfWarBot - Analyze if a user is a war bot
func CheckIfWarBot(client *linetcr.Account, groupID, userMID string) *BotBehavior {
	behavior := &BotBehavior{
		MID:         userMID,
		FirstSeen:   time.Now(),
		LastSeen:    time.Now(),
		AttackScore: 0,
	}

	// Get user contact info
	contact, err := client.GetContact(userMID)
	if err == nil {
		behavior.DisplayName = contact.DisplayName
	}

	// Check for common war bot indicators
	behavior.checkBotIndicators(client, groupID, userMID)

	return behavior
}

// checkBotIndicators - Analyze various indicators of war bot
func (b *BotBehavior) checkBotIndicators(client *linetcr.Account, groupID, userMID string) {
	// 1. Check if user has many groups (common for war bots)
	groups, _ := client.GetGroupIdsJoined()
	if len(groups) > 50 {
		b.AttackScore += 30
		b.IsWarBot = true
	}

	// 2. Check for recent account (war bots often new)
	// This would need LINE API to check account creation date

	// 3. Check for bot-like display name patterns
	suspiciousNames := []string{"bot", "war", "kick", "mod", "admin", "กู", "บอท", "วอร์"}
	for _, name := range suspiciousNames {
		if len(b.DisplayName) > 3 && (len(b.DisplayName) < 15) {
			b.AttackScore += 10
		}
	}

	// 4. Check if user is in our blocklist
	antiKickProtection.mu.RLock()
	if antiKickProtection.blockedBots[userMID] {
		b.AttackScore += 50
		b.IsWarBot = true
	}
	antiKickProtection.mu.RUnlock()
}

// DetectKickAttempt - Detect when our bot is being kicked
func DetectKickAttempt(client *linetcr.Account, groupID, kickerMID string) bool {
	// Get kicker's behavior
	kickerBehavior := CheckIfWarBot(client, groupID, kickerMID)

	// If high attack score, this is likely a war bot attack
	if kickerBehavior.AttackScore >= 30 {
		// Log the detected attack
		antiKickProtection.mu.Lock()
		if antiKickProtection.detectedBots[groupID] == nil {
			antiKickProtection.detectedBots[groupID] = []string{}
		}
		
		// Add to detected list if not already
		found := false
		for _, mid := range antiKickProtection.detectedBots[groupID] {
			if mid == kickerMID {
				found = true
				break
			}
		}
		if !found {
			antiKickProtection.detectedBots[groupID] = append(antiKickProtection.detectedBots[groupID], kickerMID)
		}
		antiKickProtection.mu.Unlock()

		return true
	}

	return false
}

// CounterAttack - When our bot is being attacked, strike back
func CounterAttack(client *linetcr.Account, groupID, attackerMID string) {
	// Get attacker info
	attacker, _ := client.GetContact(attackerMID)

	// Anti-kick: Leave before being kicked, then immediately rejoin
	go func() {
		// 1. Leave the group
		client.LeaveGroup(groupID)
		time.Sleep(500 * time.Millisecond)

		// 2. Get invite ticket
		ticket, err := client.ReissueChatTicket(groupID)
		if err == nil {
			// 3. Rejoin using ticket
			time.Sleep(1 * time.Second)
			client.AcceptGroupInvitationByTicket(groupID, ticket)
		}

		// 4. Ban the attacker
		BanUser(client, groupID, attackerMID)

		// 5. Notify admin
		notifyAdmin(attacker, groupID, "counter-attack")
	}()

	// Nuke the group - kick all except our bots
	go func() {
		time.Sleep(2 * time.Second)
		NukeAll(client, groupID)
	}()
}

// BanUser - Ban a user from group
func BanUser(client *linetcr.Account, groupID, userMID string) error {
	// Add to ban list in memory
	botstate.Banned.Banlist = append(botstate.Banned.Banlist, userMID)
	
	// Also add to blocked bots
	antiKickProtection.mu.Lock()
	antiKickProtection.blockedBots[userMID] = true
	antiKickProtection.mu.Unlock()

	// Kick the user
	client.DeleteOtherFromChat(groupID, userMID)

	return nil
}

// GetDetectedBots - Get list of detected war bots in a group
func GetDetectedBots(groupID string) []string {
	antiKickProtection.mu.RLock()
	defer antiKickProtection.mu.RUnlock()
	
	bots := antiKickProtection.detectedBots[groupID]
	if bots == nil {
		return []string{}
	}
	return bots
}

// ClearDetectedBots - Clear detected bots list for a group
func ClearDetectedBots(groupID string) {
	antiKickProtection.mu.Lock()
	defer antiKickProtection.mu.Unlock()
	
	delete(antiKickProtection.detectedBots, groupID)
}

// GetBlockedBotsCount - Get count of blocked bots
func GetBlockedBotsCount() int {
	antiKickProtection.mu.RLock()
	defer antiKickProtection.mu.RUnlock()
	
	return len(antiKickProtection.blockedBots)
}

// notifyAdmin - Send notification to admin about attack
func notifyAdmin(attacker *linetcr.Contact, groupID, attackType string) {
	// Get group info
	group, _ := botstate.ClientBot[0].GetGroup(groupID)
	
	message := fmt.Sprintf("🚨 **WAR BOT DETECTED!**\n\n👤 Attacker: %s\n📁 Group: %s\n⚔️ Attack Type: %s\n⏰ Time: %s",
		attacker.DisplayName,
		group.Name,
		attackType,
		time.Now().Format("2006-01-02 15:04:05"))

	// Send to developer
	if len(botstate.DEVELOPER) > 0 {
		botstate.ClientBot[0].SendMessage(botstate.DEVELOPER[0], message)
	}
}

// AutoProtectGroup - Enable automatic protection for a group
func AutoProtectGroup(client *linetcr.Account, groupID string, enable bool) {
	if enable {
		// Enable protection mode
		botstate.Data.ProtectmodeBack = true
		client.SendMessage(groupID, "🛡️ Anti-kick protection ENABLED")
	} else {
		// Disable protection mode
		botstate.Data.ProtectmodeBack = false
		client.SendMessage(groupID, "🛡️ Anti-kick protection DISABLED")
	}
}

// QuickRejoin - Quick rejoin after being kicked
func QuickRejoin(client *linetcr.Account, groupID string) bool {
	// Try to get ticket and rejoin
	ticket, err := client.ReissueChatTicket(groupID)
	if err != nil {
		return false
	}

	time.Sleep(500 * time.Millisecond)
	err = client.AcceptGroupInvitationByTicket(groupID, ticket)
	return err == nil
}

// DetectBotInvite - Detect if a war bot is being invited
func DetectBotInvite(client *linetcr.Account, groupID string, invitedMIDs []string) []string {
	suspiciousBots := []string{}

	for _, mid := range invitedMIDs {
		behavior := CheckIfWarBot(client, groupID, mid)
		
		// If attack score is high, mark as suspicious
		if behavior.AttackScore >= 25 {
			suspiciousBots = append(suspiciousBots, mid)
		}
	}

	return suspiciousBots
}

// MassProtect - Activate all bots in squad to protect
func MassProtect(groupID string) {
	room := linetcr.GetRoom(groupID)
	
	// All bots in squad join protection
	for _, bot := range botstate.ClientBot {
		if !utils.InArrayString(room.GoClient, bot.MID) {
			// Get ticket and invite
			ticket, _ := bot.ReissueChatTicket(groupID)
			if ticket != "" {
				bot.AcceptGroupInvitationByTicket(groupID, ticket)
			}
		}
	}
}
