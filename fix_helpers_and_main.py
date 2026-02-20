import re
import os

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
helpers_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\helpers.go'

functions_moved = [
    'MemBan', 'MemBan2', 'MemUser', 'MemUserN', 
    'IsBlacklist', 'IsBlacklist2', 'IsMember', 'GetCodeprem', 'GetSquad',
    'SendMycreator', 'SendMymaker', 'SendMyseller', 'SendMybuyer', 
    'SendMyowner', 'SendMymaster', 'SendMyadmin', 'SendMygowner', 'SendMygadmin',
    'qrGo22', 'QrKick'
]

# --- Step 1: Fix handler/helpers.go ---

with open(helpers_path, 'r', encoding='utf-8') as f:
    helpers_content = f.read()

# Add sync import if missing
if '"sync"' not in helpers_content:
    helpers_content = helpers_content.replace('import (', 'import (\n\t"sync"')

# Remove duplicate GetCodeprem stub
# Look for the stub pattern: func GetCodeprem... return false }
stub_pattern = re.compile(r'func GetCodeprem.*?return false\s*}', re.DOTALL)
# We want to keep the real implementation which is longer. The stub is short.
# Let's just find all matches and remove the one that is short (lines 127-133 approx)
matches = list(stub_pattern.finditer(helpers_content))
if len(matches) > 1:
    # Assuming the first one is the stub (it appeared first in the file)
    # Double check if it looks like a stub (short length)
    first_match = matches[0]
    if len(first_match.group(0)) < 200: # Stub is small
        helpers_content = helpers_content[:first_match.start()] + helpers_content[first_match.end():]

# Append QrKick if not present
if 'func QrKick' not in helpers_content:
    qrkick_code = """
func QrKick(client *linetcr.Account, to string) {
	if len(botstate.Banned.Banlist) != 0 {
		for _, v := range botstate.Banned.Banlist {
			if linetcr.IsMembers(client, to, v) == true {go func() {client.DeleteOtherFromChats(to, []string{v})}()}
			if linetcr.IsPending(client, to, v) == true {go func() {client.CancelChatInvitations(to, []string{v})}()}
		}
	}
}
"""
    helpers_content += qrkick_code

with open(helpers_path, 'w', encoding='utf-8') as f:
    f.write(helpers_content)
print("Updated handler/helpers.go")

# --- Step 2: Clean up main.go ---

with open(main_path, 'r', encoding='utf-8') as f:
    main_content = f.read()

# Helper to remove function definition
def remove_function_def(content, func_name):
    # Find start of function
    pattern = re.compile(r'func ' + func_name + r'\s*\(', re.MULTILINE)
    match = pattern.search(content)
    if not match:
        return content
    
    start_index = match.start()
    # Find matching brace for the function body
    brace_count = 0
    end_index = -1
    found_start_brace = False
    
    # Start searching from the function definition line
    for i in range(start_index, len(content)):
        if content[i] == '{':
            brace_count += 1
            found_start_brace = True
        elif content[i] == '}':
            brace_count -= 1
            if found_start_brace and brace_count == 0:
                end_index = i + 1
                break
    
    if end_index != -1:
        # Remove the code block
        return content[:start_index] + content[end_index:]
    return content

for func in functions_moved:
    main_content = remove_function_def(main_content, func)

# Replace calls
# We use regex \bFuncName\( to replace with handler.FuncName(
# But we must ensure we don't replace definitions (which are already removed)
# or other things.
for func in functions_moved:
    pattern = re.compile(r'\b' + func + r'\(')
    main_content = pattern.sub('handler.' + func + '(', main_content)

# Add handler import
if 'gobots/handler"' not in main_content:
    # Try to find where to insert
    # Look for "gobots/botstate" and add after
    if '"gobots/botstate"' in main_content:
        main_content = main_content.replace('"gobots/botstate"', '"gobots/botstate"\n\t"gobots/handler"')
    elif '"./botstate"' in main_content: # It might use relative path in main.go?
        main_content = main_content.replace('"./botstate"', '"./botstate"\n\t"./handler"')
    else:
        # Just find import (
        main_content = main_content.replace('import (', 'import (\n\t"./handler"')

with open(main_path, 'w', encoding='utf-8') as f:
    f.write(main_content)
print("Updated main.go")
