import re
import os

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
actions_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\actions.go'

functions_to_move = [
    'AccKickBan', 'AccGroup', 'JoinLlinetcrBan', 'AcceptJoin', 'AcceptJoinV2',
    'KickCancelProtect', 'kickProtect', 'BanAll', 'cancelBanInv', 'cancelallcek',
    'kickPelaku', 'groupBackupKick', 'groupBackupInv', 'GhostEnd', 'LogOp', 'LogGet',
    'MemAccsess', 'NotifBot', 'Checkkickuser', 'back'
]

with open(main_path, 'r', encoding='utf-8') as f:
    main_content = f.read()

extracted_code = ""

# Helper to extract function
def extract_function(content, func_name):
    pattern = re.compile(r'func ' + func_name + r'\s*\(', re.MULTILINE)
    match = pattern.search(content)
    if not match:
        return None, content
    
    start_index = match.start()
    brace_count = 0
    end_index = -1
    found_start_brace = False
    
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
        func_code = content[start_index:end_index]
        return func_code, content
    return None, content

for func in functions_to_move:
    code, _ = extract_function(main_content, func)
    if code:
        extracted_code += code + "\n\n"

# Fix references in extracted code
# logic functions moved previously are now handler.Func
# But since we are in handler package, we can call them directly?
# No, they are in logic.go (same package handler), so we call them directly.
# BUT, main.go currently has them as `handler.Func`. 
# If I copy code from main.go, it might have `handler.Func` calls if I replaced them earlier.
# If I call a function in `logic.go` from `actions.go` (same package), I should NOT use `handler.` prefix.
# So I need to strip `handler.` prefix for calls to functions that are in logic.go.
# List of logic functions: CanceljoinBot, Nukjoin, AutopurgeEnd, RemoveSticker, GETgrade
logic_funcs = ['CanceljoinBot', 'Nukjoin', 'AutopurgeEnd', 'RemoveSticker', 'GETgrade']
for func in logic_funcs:
    extracted_code = extracted_code.replace('handler.' + func, func)

# Also strip handler. prefix for helper functions if they are used (MemUser etc)
helper_funcs = ['MemBan', 'MemBan2', 'MemUser', 'MemUserN', 'IsBlacklist', 'IsBlacklist2', 'IsMember', 'GetCodeprem', 'GetSquad', 'SendMycreator', 'SendMymaker', 'SendMyseller', 'SendMybuyer', 'SendMyowner', 'SendMymaster', 'SendMyadmin', 'SendMygowner', 'SendMygadmin', 'qrGo22', 'QrKick']
for func in helper_funcs:
    extracted_code = extracted_code.replace('handler.' + func, func)

# Fix utils.utils typo
extracted_code = extracted_code.replace('utils.utils.', 'utils.')

# Create actions.go
actions_content = """package handler

import (
	"fmt"
	"strings"
	"sync"
	"time"
    "runtime"
	"../botstate"
	"../config"
	"../library/linetcr"
	"../utils"
)

""" + extracted_code

with open(actions_path, 'w', encoding='utf-8') as f:
    f.write(actions_content)
print(f"Created {actions_path}")

# Remove functions from main.go
def remove_function_def(content, func_name):
    pattern = re.compile(r'func ' + func_name + r'\s*\(', re.MULTILINE)
    match = pattern.search(content)
    if not match:
        return content
    start_index = match.start()
    brace_count = 0
    end_index = -1
    found_start_brace = False
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
        return content[:start_index] + content[end_index:]
    return content

for func in functions_to_move:
    main_content = remove_function_def(main_content, func)

# Replace calls in main.go
for func in functions_to_move:
    pattern = re.compile(r'\b' + func + r'\(')
    main_content = pattern.sub('handler.' + func + '(', main_content)

# Fix utils.utils typo in main.go
main_content = main_content.replace('utils.utils.', 'utils.')

with open(main_path, 'w', encoding='utf-8') as f:
    f.write(main_content)
print(f"Updated {main_path}")
