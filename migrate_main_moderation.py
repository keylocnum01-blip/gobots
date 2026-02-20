import os

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

funcs_to_remove = [
    "nukeAll", "KickBan132V2", "CancelBan125V2", "KickBan132", "CancelBan125",
    "FastKick", "FastCancel", "NodeBans", "PurgeFaster", "KickMemBan", "CanMemBan",
    "KickCancel", "KickCancelBan129", "KickBan129", "AccKickBan", "AccGroup",
    "JoinLlinetcrBan", "QrKick"
]

# 1. Remove function definitions
lines = content.split('\n')
new_lines = []
skip = False
brace_count = 0

for line in lines:
    stripped = line.strip()
    
    # Check if this line starts a function we want to remove
    start_func = False
    for func in funcs_to_remove:
        if stripped.startswith(f"func {func}("):
            start_func = True
            break
    
    if start_func:
        skip = True
        brace_count = 0
    
    if skip:
        brace_count += line.count('{')
        brace_count -= line.count('}')
        if brace_count == 0:
            skip = False
        continue
    
    new_lines.append(line)

content = '\n'.join(new_lines)

# 2. Update calls
replacements = {
    "nukeAll(": "handler.NukeAll(",
    "KickBan132V2(": "handler.KickBan132V2(",
    "CancelBan125V2(": "handler.CancelBan125V2(",
    "KickBan132(": "handler.KickBan132(",
    "CancelBan125(": "handler.CancelBan125(",
    "FastKick(": "handler.FastKick(",
    "FastCancel(": "handler.FastCancel(",
    "NodeBans(": "handler.NodeBans(",
    "PurgeFaster(": "handler.PurgeFaster(",
    "KickMemBan(": "handler.KickMemBan(",
    "CanMemBan(": "handler.CanMemBan(",
    "KickCancel(": "handler.KickCancel(",
    "KickCancelBan129(": "handler.KickCancelBan129(",
    "KickBan129(": "handler.KickBan129(",
    "AccKickBan(": "handler.AccKickBan(",
    "AccGroup(": "handler.AccGroup(",
    "JoinLlinetcrBan(": "handler.JoinLlinetcrBan(",
    "QrKick(": "handler.QrKick("
}

for old, new in replacements.items():
    content = content.replace(old, new)

# Fix double handler prefixes if any
content = content.replace("handler.handler.", "handler.")

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)

print("Migrated moderation functions in main.go")
