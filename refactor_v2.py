import os
import re

# File paths
main_go_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
handler_actions_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\actions.go'

# 1. Append functions to handler/actions.go
append_code = """
func CancelEnd(client *linetcr.Account, Group string, mem []string) {
	defer utils.PanicHandle("CancelEnd")
	for _, target := range mem {
		client.CancelChatInvitations(Group, []string{target})
	}
}

func Setpurgealln(client *linetcr.Account, to string, invits []string) {
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, []string{cc})
		} else if linetcr.IsPending(client, to, cc) == true {
			client.CancelChatInvitations(to, []string{cc})
		}
	}
}
"""

with open(handler_actions_path, 'r', encoding='utf-8') as f:
    content = f.read()

if "func CancelEnd" not in content:
    with open(handler_actions_path, 'a', encoding='utf-8') as f:
        f.write(append_code)
    print("Appended functions to handler/actions.go")
else:
    print("CancelEnd already in handler/actions.go")

# 2. Refactor main.go
with open(main_go_path, 'r', encoding='utf-8') as f:
    main_content = f.read()

# Functions to remove (regex patterns)
funcs_to_remove = [
    r'func fmtDuration\(d time\.Duration\) string \{[\s\S]*?\n\}',
    r'func GetKorban\(user string\) \*linetcr\.Account \{[\s\S]*?\n\}',
    r'func SelectBot\(client \*linetcr\.Account, to string\) \(\*linetcr\.Account, error\) \{[\s\S]*?\n\}',
    r'func CheckBot\(client \*linetcr\.Account, to string\) \(\*linetcr\.Account, error\) \{[\s\S]*?\n\}',
    r'func squadMention\(mlist \[\]string\) \(m \*linetcr\.Account, b bool\) \{[\s\S]*?\n\}',
    r'func RemoveString\(s \[\]string, r string\) \[\]string \{[\s\S]*?\n\}',
    r'func CancelEnd\(client \*linetcr\.Account, Group string, mem \[\]string\) \{[\s\S]*?\n\}',
    r'func Setpurgealln\(client \*linetcr\.Account, to string, invits \[\]string\) \{[\s\S]*?\n\}',
]

for pattern in funcs_to_remove:
    main_content = re.sub(pattern, '', main_content)

# Replacements
replacements = {
    'fmtDuration(': 'utils.FmtDuration(',
    'GetKorban(': 'botstate.GetKorban(',
    'SelectBot(': 'botstate.SelectBot(',
    'CheckBot(': 'botstate.CheckBot(',
    'squadMention(': 'botstate.SquadMention(',
    'RemoveString(': 'utils.RemoveString(',
    'CancelEnd(': 'handler.CancelEnd(',
    'Setpurgealln(': 'handler.Setpurgealln(',
}

for old, new in replacements.items():
    main_content = main_content.replace(old, new)

# Fix double prefixes if any (e.g. handler.handler.CancelEnd)
main_content = main_content.replace('handler.handler.', 'handler.')
main_content = main_content.replace('botstate.botstate.', 'botstate.')
main_content = main_content.replace('utils.utils.', 'utils.')

with open(main_go_path, 'w', encoding='utf-8') as f:
    f.write(main_content)

print("Refactored main.go")
