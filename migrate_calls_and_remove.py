import re
import os

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"

functions_to_migrate = [
    "CancelEnemy", "GroupBackupKick", "CekOp2", "CekOp", "CekBanwhois", 
    "CekUnbanwhois", "CheckUnbanBots", "Purgesip", "KickBl", "KickCancelBan129", 
    "NukeAll", "KickMemBan", "CanMemBan", "KickCancel", "Autokickban", 
    "Purgemode", "KIckbansPurges", "AcceptWar", "AccKickBan", "AccGroup", 
    "JoinLlinetcrBan", "KickBan129", "GroupBackupWar", "KickCancelProtect", 
    "KickPelaku"
]

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

for func_name in functions_to_migrate:
    # 1. Replace calls, ensuring we don't replace the definition
    # Regex lookbehind or careful matching.
    # Match "func_name(" but NOT "func func_name("
    
    # Simple approach: Replace all, then revert the definition signature? 
    # Or better: Use a callback with re.sub
    
    pattern_call = r'(?<!func\s)\b' + re.escape(func_name) + r'\('
    content = re.sub(pattern_call, f'handler.{func_name}(', content)

    # 2. Remove definition
    # We need to match balanced braces.
    pattern_def_start = r'func\s+' + re.escape(func_name) + r'\s*\('
    
    match = re.search(pattern_def_start, content)
    if match:
        start_idx = match.start()
        # Find the opening brace of the function body
        # It might be on the same line or next
        brace_search = re.search(r'\{', content[start_idx:])
        if brace_search:
            brace_open_idx = start_idx + brace_search.start()
            
            # Find matching closing brace
            balance = 1
            i = brace_open_idx + 1
            while i < len(content) and balance > 0:
                if content[i] == '{':
                    balance += 1
                elif content[i] == '}':
                    balance -= 1
                i += 1
            
            if balance == 0:
                # Remove the function block including the signature
                # content = content[:start_idx] + content[i:]
                # We should probably remove empty lines around it too, but let's be safe first
                print(f"Removing function definition: {func_name}")
                content = content[:start_idx] + "\n" + content[i:]
            else:
                print(f"Could not find closing brace for {func_name}")
        else:
            print(f"Could not find opening brace for {func_name}")
    else:
        print(f"Could not find definition for {func_name}")

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)
