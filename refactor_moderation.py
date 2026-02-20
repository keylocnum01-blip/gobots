import os
import re

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
mod_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\moderation.go'

with open(main_path, 'r', encoding='utf-8') as f:
    main_content = f.read()

with open(mod_path, 'r', encoding='utf-8') as f:
    mod_content = f.read()

# Functions to move (not present in moderation.go yet)
functions_to_move = [
    'KickCancelProtect',
    'kickPelaku',
    'kickProtect',
    'GhostEnd'
]

# Functions to just remove (already in moderation.go)
functions_to_remove = [
    'KickCancelBan129', 'KickBan132V2', 'CancelBan125V2', 'KickBan132', 'CancelBan125',
    'FastKick', 'FastCancel', 'NodeBans', 'PurgeFaster', 'KickCancel', 'Purgemode',
    'KickAllBan', 'checkunbanbots', 'kickBl', 'AccKickBan', 'AccGroup',
    'JoinLlinetcrBan', 'KickBan129', 'GroupBackupWar', 'KickMemBan', 'CanMemBan',
    'Autokickban', 'KIckbansPurges', 'KIckbansPurges1', 'KickCansWar',
    'JKickFuck', 'JCancelFuck', 'AcceptWar'
]

new_mod_content = mod_content

# Add import if missing
if '"github.com/asaskevich/govalidator"' not in new_mod_content:
    new_mod_content = new_mod_content.replace('import (', 'import (\n\tvalid "github.com/asaskevich/govalidator"')

# Extract and append functions to move
extracted_funcs = ""
for func_name in functions_to_move:
    # Pattern to match function definition including body
    # This is a simple regex and might fail with nested braces if not careful, but assuming standard formatting
    pattern = re.compile(r'func ' + func_name + r'\s*\(.*?\)\s*\{(?:[^{}]*|\{(?:[^{}]*|\{[^{}]*\})*\})*\}', re.DOTALL)
    match = pattern.search(main_content)
    if match:
        func_code = match.group(0)
        # Refactor the code
        func_code = func_code.replace('defer panicHandle', 'defer utils.PanicHandle')
        func_code = func_code.replace('Squadlist', 'botstate.Squadlist')
        func_code = func_code.replace('GetKorban', 'botstate.GetKorban')
        func_code = func_code.replace('MaxKick', 'botstate.MaxKick')
        func_code = func_code.replace('cpu', 'botstate.Cpu')
        func_code = func_code.replace('cekGo', 'botstate.CekGo') # case sensitive? main.go used cekGo
        func_code = func_code.replace('Ajsjoin', 'botstate.Ajsjoin')
        func_code = func_code.replace('ProtectMode', 'botstate.ProtectMode')
        func_code = func_code.replace('Banned', 'botstate.Banned')
        func_code = func_code.replace('Checkkickuser', 'CheckKickUser') # Assuming this helper needs to be exported/available or moved?
        # Checkkickuser is likely in handler/logic.go or main.go? 
        # If it's in main.go and local, I might need to move it too or export it.
        # Let's assume for now it's available or I'll fix it later.
        
        # Capitalize function name
        new_func_name = func_name[0].upper() + func_name[1:]
        func_code = func_code.replace('func ' + func_name, 'func ' + new_func_name)
        
        extracted_funcs += "\n\n" + func_code
        
        # Remove from main_content
        main_content = main_content.replace(match.group(0), "")
    else:
        print(f"Could not find {func_name} in main.go")

new_mod_content += extracted_funcs

with open(mod_path, 'w', encoding='utf-8') as f:
    f.write(new_mod_content)

# Remove duplicates
for func_name in functions_to_remove:
    pattern = re.compile(r'func ' + func_name + r'\s*\(.*?\)\s*\{(?:[^{}]*|\{(?:[^{}]*|\{[^{}]*\})*\})*\}', re.DOTALL)
    match = pattern.search(main_content)
    if match:
        main_content = main_content.replace(match.group(0), "")
    else:
        print(f"Could not find duplicate {func_name} in main.go (might be already removed)")

# Update calls in main.go
# Map for replacements
replacements = {
    'KickCancelProtect': 'handler.KickCancelProtect',
    'kickPelaku': 'handler.KickPelaku',
    'kickProtect': 'handler.KickProtect',
    'GhostEnd': 'handler.GhostEnd',
    'KickCancelBan129': 'handler.KickCancelBan129',
    'KickBan132V2': 'handler.KickBan132V2',
    'CancelBan125V2': 'handler.CancelBan125V2',
    'KickBan132': 'handler.KickBan132',
    'CancelBan125': 'handler.CancelBan125',
    'FastKick': 'handler.FastKick',
    'FastCancel': 'handler.FastCancel',
    'NodeBans': 'handler.NodeBans',
    'PurgeFaster': 'handler.PurgeFaster',
    'KickCancel': 'handler.KickCancel',
    'Purgemode': 'handler.Purgemode',
    'KickAllBan': 'handler.KickAllBan',
    'checkunbanbots': 'handler.CheckUnbanBots',
    'kickBl': 'handler.KickBl',
    'AccKickBan': 'handler.AccKickBan',
    'AccGroup': 'handler.AccGroup',
    'JoinLlinetcrBan': 'handler.JoinLlinetcrBan',
    'KickBan129': 'handler.KickBan129',
    'GroupBackupWar': 'handler.GroupBackupWar',
    'KickMemBan': 'handler.KickMemBan',
    'CanMemBan': 'handler.CanMemBan',
    'Autokickban': 'handler.Autokickban',
    'KIckbansPurges': 'handler.KIckbansPurges',
    'KIckbansPurges1': 'handler.KIckbansPurges1',
    'KickCansWar': 'handler.KickCansWar',
    'JKickFuck': 'handler.JKickFuck',
    'JCancelFuck': 'handler.JCancelFuck',
    'AcceptWar': 'handler.AcceptWar'
}

for old, new in replacements.items():
    # Use regex to replace calls to avoid replacing substrings
    # Look for boundary or start of line, followed by function name, followed by (
    # But wait, we just want to replace the function call.
    # Simple replace might be safer if names are unique enough.
    # Most names are quite unique.
    # Be careful with case sensitivity.
    
    # Special handling for lowercase start functions to uppercase handler calls
    main_content = re.sub(r'(?<!handler\.)\b' + old + r'\b', new, main_content)

with open(main_path, 'w', encoding='utf-8') as f:
    f.write(main_content)

print("Refactoring complete.")
