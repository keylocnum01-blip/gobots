
import os

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
moderation_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\moderation.go'
helpers_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\helpers.go'

# Functions to migrate
funcs_to_migrate = [
    {
        'name': 'autokickban',
        'new_name': 'AutoKickBan',
        'target': moderation_path,
        'old_block': """func autokickban(client *linetcr.Account, to string, target string) {
	if botstate.AutokickBan {
		gr, _ := client.GetGroupIdsJoined()
		for _, aa := range gr {
			go client.DeleteOtherFromChats(aa, []string{target})
			go client.CancelChatInvitations(aa, []string{target})
		}
	}
}""",
        'new_content': """func AutoKickBan(client *linetcr.Account, to string, target string) {
	if botstate.AutokickBan {
		gr, _ := client.GetGroupIdsJoined()
		for _, aa := range gr {
			go client.DeleteOtherFromChats(aa, []string{target})
			go client.CancelChatInvitations(aa, []string{target})
		}
	}
}
"""
    },
    {
        'name': 'IsMember',
        'new_name': 'IsMember',
        'target': helpers_path,
        'old_block': """func IsMember(client *linetcr.Account, from string, groups string) bool {
	res := client.GetGroup(groups)
	memlist := res.Members
	for _, a := range memlist {
		if a.Mid == from {
			return true
			break
		}
	}
	return false
}""",
        'new_content': """func IsMember(client *linetcr.Account, from string, groups string) bool {
	res := client.GetGroup(groups)
	memlist := res.Members
	for _, a := range memlist {
		if a.Mid == from {
			return true
			break
		}
	}
	return false
}
"""
    },
    {
        'name': 'groupBackupWar',
        'new_name': 'GroupBackupWar',
        'target': moderation_path,
        'old_block': """func groupBackupWar(client *linetcr.Account, to string) {
	for x := range botstate.Squadlist {if linetcr.IsMembers(client, to, botstate.Squadlist[x]) == true {if client.MID == botstate.Squadlist[x] {go func() {handler.KickCansWar(client, to)}()};break} else {continue}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}""",
        'new_content': """func GroupBackupWar(client *linetcr.Account, to string) {
	for x := range botstate.Squadlist {
		if linetcr.IsMembers(client, to, botstate.Squadlist[x]) == true {
			if client.MID == botstate.Squadlist[x] {
				go func() {
					KickCansWar(client, to)
				}()
			}
			break
		} else {
			continue
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}
"""
    }
]

# Read main.go
with open(main_path, 'r', encoding='utf-8') as f:
    main_content = f.read()

# Fix apikey
if '"+apikey' in main_content:
    print("Fixing apikey...")
    main_content = main_content.replace('"+apikey', '"+botstate.Apikey')

# Process functions
for func in funcs_to_migrate:
    print(f"Migrating {func['name']}...")
    
    # Remove from main.go
    if func['old_block'] in main_content:
        main_content = main_content.replace(func['old_block'], "")
        print(f"Removed {func['name']} from main.go")
    else:
        print(f"Warning: Could not find block for {func['name']}")
        # Fallback: try to find with normalized whitespace or just proceed if already removed
    
    # Append to target
    # Check if already exists to avoid duplicates
    with open(func['target'], 'r', encoding='utf-8') as f:
        target_content = f.read()
    
    if f"func {func['new_name']}" not in target_content:
        with open(func['target'], 'a', encoding='utf-8') as f:
            f.write("\n\n" + func['new_content'])
        print(f"Appended {func['new_name']} to {func['target']}")
    else:
        print(f"Function {func['new_name']} already exists in target.")

    # Update references in main.go
    # Replace `name(` with `handler.NewName(`
    # Be careful not to replace `func name(` if it wasn't removed (but we tried to remove it)
    # Also handle cases where it might be `go name(` or `name(`
    
    # Simple replacement of function calls
    # regex: \bname\(
    import re
    pattern = r'\b' + re.escape(func['name']) + r'\('
    replacement = f"handler.{func['new_name']}("
    main_content = re.sub(pattern, replacement, main_content)

# Write main.go
with open(main_path, 'w', encoding='utf-8') as f:
    f.write(main_content)

print("Migration complete.")
