package handler

import (
	"fmt"
	"strings"
	"time"

	"../botstate"
	"../library/linetcr"
	"../utils"
)

// AntiKickCommands - Handle anti-kick protection commands
func AntiKickCommands(client *linetcr.Account, groupID, userMID, text string) {
	// Check if user has permission
	if !CheckPermission(&botstate.Command{Admin: true}, userMID, groupID) {
		client.SendMessage(groupID, "❌ No permission")
		return
	}

	// Parse command
	parts := strings.Fields(text)
	if len(parts) < 2 {
		client.SendMessage(groupID, "Usage: /antikit <command>")
		return
	}

	command := strings.ToLower(parts[1])

	switch command {
	case "on":
		// Enable anti-kick protection
		AutoProtectGroup(client, groupID, true)

	case "off":
		// Disable anti-kick protection
		AutoProtectGroup(client, groupID, false)

	case "list":
		// Show detected bots
		bots := GetDetectedBots(groupID)
		if len(bots) == 0 {
			client.SendMessage(groupID, "✅ No war bots detected")
			return
		}

		message := "🚨 **Detected War Bots:**\n\n"
		for i, mid := range bots {
			contact, _ := client.GetContact(mid)
			name := mid
			if contact != nil {
				name = contact.DisplayName
			}
			message += fmt.Sprintf("%d. %s\n", i+1, name)
		}
		client.SendMessage(groupID, message)

	case "clear":
		// Clear detected bots
		ClearDetectedBots(groupID)
		client.SendMessage(groupID, "✅ Detected bots cleared")

	case "stats":
		// Show protection stats
		blockedCount := GetBlockedBotsCount()
		detectedCount := len(GetDetectedBots(groupID))

		message := fmt.Sprintf("📊 **Protection Stats:**\n\n🛡️ Blocked Bots: %d\n🚨 Detected: %d\n⏰ Status: %s",
			blockedCount,
			detectedCount,
			getProtectionStatus())
		client.SendMessage(groupID, message)

	case "rejoin":
		// Quick rejoin if kicked
		success := QuickRejoin(client, groupID)
		if success {
			client.SendMessage(groupID, "✅ Rejoined successfully!")
		} else {
			client.SendMessage(groupID, "❌ Could not rejoin")
		}

	case "mass":
		// Activate all squad bots
		MassProtect(groupID)
		client.SendMessage(groupID, "🚀 Squad protection activated!")

	case "nuke":
		// Nuke the group
		NukeAll(client, groupID)
		client.SendMessage(groupID, "💥 Nuke activated!")

	case "banall":
		// Ban all detected bots
		bots := GetDetectedBots(groupID)
		for _, mid := range bots {
			BanUser(client, groupID, mid)
		}
		client.SendMessage(groupID, fmt.Sprintf("✅ Banned %d war bots", len(bots)))

	default:
		client.SendMessage(groupID, "Unknown command. Available: on, off, list, clear, stats, rejoin, mass, nuke, banall")
	}
}

// getProtectionStatus - Get current protection status
func getProtectionStatus() string {
	if botstate.Data.ProtectmodeBack {
		return "🟢 ACTIVE"
	}
	return "🔴 INACTIVE"
}

// WarBotHelp - Show help for anti-kick commands
func WarBotHelp(client *linetcr.Account, groupID string) {
	message := `🛡️ **Anti-Kick Protection Commands**

📖 Available Commands:

• /antikit on - Enable protection
• /antikit off - Disable protection  
• /antikit list - Show detected bots
• /antikit clear - Clear detected list
• /antikit stats - Show protection stats
• /antikit rejoin - Quick rejoin
• /antikit mass - Activate squad
• /antikit nuke - Nuke group
• /antikit banall - Ban all detected

🔒 Features:
- Auto-detect war bots
- Auto-ban attackers
- Quick rejoin after kick
- Squad protection mode
- Attack notifications`

	client.SendMessage(groupID, message)
}
