# Ultimate Warfare System

Advanced war bot system for group destruction and defense.

## Overview

This system provides comprehensive warfare capabilities:
- **8 Phases** of warfare
- **Multi-layered defense**
- **Coordinated attacks**
- **Psychological warfare**

---

## Phase 1: Early Warning & Detection

### Multi-Layer Detection System

```
Risk Score Calculation (0-100):

LAYER 1: Group Count Analysis
├── >100 groups = +40 risk
├── >50 groups = +25 risk
└── Normal = 0

LAYER 2: Behavior Pattern
├── Kick attempts ×15
├── Invite floods ×5
└── Cancel attempts ×10

LAYER 3: Account Age
└── <7 days = +30 risk

LAYER 4: Name Analysis
└── Suspicious keywords = +10 risk

LAYER 5: Known War Bot DB
└── Match found = +50 risk

LAYER 6: Invite Pattern
└── >5 invites/10min = +25 risk
```

**CONFIRMED WAR BOT: Risk Score ≥ 50**

---

## Phase 2: Defense Strategies

### Defense Levels

| Level | Name | Description |
|-------|------|-------------|
| 1 | Passive | Basic protection only |
| 2 | Active | Auto-detect and counter |
| 3 | Aggressive | Pre-emptive strikes |
| 4 | Total | Full war mode |

### Rejoin Strategies (in order)

1. **Ticket Rejoin** - Use reissued ticket
2. **Friend Reinvite** - Ask friend to reinvite
3. **QR Scan** - Show QR for manual scan

---

## Phase 3: Counter-Attack

### Intelligence Gathering

```go
// Gather data on enemy group:
- Member count
- Bot count  
- Bot identities
- Admin list
- Group name
- Weaknesses
- Strengths
- Threat level (1-10)
```

### Attack Decision Matrix

| Threat Level | Response |
|--------------|----------|
| 8-10 | Total Nuke |
| 5-7 | Targeted Strikes |
| 1-4 | Harassment |

---

## Phase 4: Nuke System

### Multi-Wave Destruction

```
Wave-based kicking:
├── Each wave: 10 members (configurable)
├── Delay between waves: 100ms
├── Multi-bot coordination
└── Prioritize bots first

Execution:
├── Pre-nuke notification
├── Wave 1: Bots (priority)
├── Wave 2: Active members
├── Wave 3: Remaining
└── Final: Completion report
```

### Multi-Bot Coordination

```
When UseAllBots = true:
├── Distribute targets among bots
├── Parallel execution
├── Synchronized timing
└── Maximum throughput
```

---

## Phase 5: Psychological Warfare

### Demoralization Messages

Pre-nuke:
```
"😈 Your group is about to be destroyed..."
"💀 Any last words?"
"🚀 Nuke incoming..."
"⚠️ This is your final warning"
"🏴‍☠️ Surrender or be destroyed"
```

### Raid Commands

```
/war raid <message> - Nuke with custom message
```

---

## Phase 6: Commands

### Available Commands

| Command | Description |
|---------|-------------|
| `/war scan` | Scan group for threats |
| `/war defense <mode>` | Set defense level |
| `/war nuke` | Destroy group |
| `/war raid <msg>` | Nuke with message |
| `/war counter <id>` | Counter-attack |
| `/war rejoin` | Quick rejoin |
| `/war status` | War status |

### Defense Modes

```
/war defense passive   - Basic
/war defense active   - Auto-counter
/war defense aggressive - Pre-emptive
/war defense total    - Full war
```

---

## Phase 7: Invite Warfare

### Attack Methods

1. **Invite Flood** - Invite same person multiple times
2. **Cancel Wave** - Cancel all pending invites
3. **Bot Spread** - Occupy group with bot accounts

---

## Phase 8: Scout & Evade

### Evasion Techniques

1. **Pre-emptive Leave** - Leave before kick
2. **Quick Rejoin** - Return immediately
3. **Position Shuffle** - Move between groups

---

## Usage Example

```go
// In message handler:
// 1. Detect threat
profile := AnalyzeUserThreat(client, groupID, userMID)
if profile.ConfirmedWarBot {
    // 2. Activate defense
    ActivateDefense(client, groupID, DefenseActive)
    
    // 3. Counter-attack if attacked
    CounterAttackPlan(client, groupID, enemyGroupID)
}

// 4. Execute nuke
/war nuke

// 5. Or targeted raid
/war raid Your group is doomed!
```

---

## Configuration

```go
WarConfig{
    AutoRejoin:       true,
    CounterAttack:    true,
    MultiNuke:        true,
    PsychologicalWar: true,
    DeepRecon:        true,
    MaxNukePerWave:   10,
    NukeDelayMs:      100,
}
```

---

## Risk Assessment

### Automatic Scoring

- **0-30**: Safe
- **31-49**: Suspicious  
- **50-69**: Likely War Bot
- **70-100**: Confirmed War Bot

### Indicators Tracked

- Excessive group membership
- Rapid kick/invite operations
- New account age
- Suspicious display names
- Known war bot database match
- Invite flood patterns

---

## Files

- `warfare_ultimate.go` - Complete warfare system
- `anti_kick_protection.go` - Anti-kick module
- `anti_kick_commands.go` - Command handlers
