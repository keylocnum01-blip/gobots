# Anti-Kick Bot Protection

Advanced protection system to defend against war bots that try to kick your bots.

## Features

### 🚨 Detection
- **Bot Behavior Analysis** - Analyze user actions to detect war bots
- **Attack Score** - Rate 0-100 how likely it's a war bot
- **Auto-Detection** - Automatically detect when being attacked

### 🛡️ Protection
- **Quick Rejoin** - Automatically rejoin after being kicked
- **Counter Attack** - Strike back at attackers
- **Mass Protection** - All squad bots join to protect
- **Auto Ban** - Ban detected war bots automatically

### 📊 Monitoring
- **Detected Bots List** - Track all detected war bots
- **Stats Dashboard** - View protection statistics
- **Attack Notifications** - Get notified of attacks

## Commands

| Command | Description |
|---------|-------------|
| `/antikit on` | Enable protection |
| `/antikit off` | Disable protection |
| `/antikit list` | Show detected bots |
| `/antikit clear` | Clear detected list |
| `/antikit stats` | Show protection stats |
| `/antikit rejoin` | Quick rejoin |
| `/antikit mass` | Activate squad protection |
| `/antikit nuke` | Nuke the attacking group |
| `/antikit banall` | Ban all detected bots |

## How It Works

1. **Detection Phase**
   - Monitors user behavior in group
   - Analyzes kick attempts
   - Calculates attack score

2. **Response Phase**
   - If attack detected:
     - Leave group before being kicked
     - Rejoin immediately
     - Ban the attacker
     - Notify admin

3. **Counter-Attack** (optional)
   - Nuke the attacking group
   - Mass ban all members
   - Report to LINE (if possible)

## Detection Criteria

War bot is detected based on:
- Number of groups joined (>50 = suspicious)
- Blocked bot list
- Kick patterns
- Invite patterns
- Display name patterns

## Integration

The module integrates with existing gobots:

```go
// In handler/bot.go, add to message handler:
if strings.HasPrefix(text, "/antikit") {
    AntiKickCommands(client, groupID, userMID, text)
}

// In kick detection:
if DetectKickAttempt(client, groupID, kickerMID) {
    CounterAttack(client, groupID, kickerMID)
}
```

## Files

- `anti_kick_protection.go` - Core protection logic
- `anti_kick_commands.go` - Command handlers
