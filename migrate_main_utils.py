import os
import re

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

funcs_to_remove = [
    "GetIP", "Checkserver", "MentionList", "SendMycreator", "SendMymaker",
    "SendMyseller", "SendMybuyer", "SendMyowner", "SendMymaster", "SendMyadmin",
    "SendMygowner", "SendMygadmin", "MaxRevision", "fmtDuration", "GetCodeprem",
    "SaveData", "MemUser", "MemBan", "IsBlacklist", "Ungban", "Addgban", "AddbanOp3",
    "Joinsave", "RandomToString", "sendBigImage"
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
    "GetIP(": "handler.GetIP(",
    "Checkserver(": "handler.Checkserver(",
    "MentionList(": "handler.MentionList(",
    "SendMycreator(": "handler.SendMycreator(",
    "SendMymaker(": "handler.SendMymaker(",
    "SendMyseller(": "handler.SendMyseller(",
    "SendMybuyer(": "handler.SendMybuyer(",
    "SendMyowner(": "handler.SendMyowner(",
    "SendMymaster(": "handler.SendMymaster(",
    "SendMyadmin(": "handler.SendMyadmin(",
    "SendMygowner(": "handler.SendMygowner(",
    "SendMygadmin(": "handler.SendMygadmin(",
    "MaxRevision(": "handler.MaxRevision(",
    "fmtDuration(": "handler.FmtDurations(",
    "GetCodeprem(": "handler.CheckPermission(",
    "SaveData(": "handler.SaveData(",
    "MemUser(": "handler.MemUser(",
    "MemBan(": "handler.MemBan(",
    "IsBlacklist(": "handler.IsBlacklist(",
    "Ungban(": "handler.Ungban(",
    "Addgban(": "handler.Addgban(",
    "AddbanOp3(": "handler.AddbanOp3(",
    "Joinsave(": "handler.Joinsave(",
    "RandomToString(": "handler.RandomToString(",
    "sendBigImage(": "handler.SendBigImage("
}

# Apply replacements
# We need to be careful not to replace definitions if any remain, but we removed them.
# Also avoid replacing handler.handler.Name
# And avoid replacing func definitions if I missed any (unlikely).

for old, new in replacements.items():
    content = content.replace(old, new)

# Fix double handler prefixes if any
content = content.replace("handler.handler.", "handler.")

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)

print("Migrated utils functions in main.go")
