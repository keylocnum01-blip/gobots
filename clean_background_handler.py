import os
import re

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\background.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

# 1. Remove duplicate Autoset (lowercase) function
# The pattern looks for func Autoset() { ... }
# We need to be careful to match the specific duplicate block we saw in the Read output.
# It started around line 355 and ended around line 658.
# It starts with "func Autoset() {" and ends before "func LogGet(op *SyncService.Operation) {"

# Let's try to identify the duplicate block by its signature and remove it.
# Note: The valid function is "func AutoSet() {" (PascalCase) at the top.
# The duplicate is "func Autoset() {" (lowercase) later in the file.

# Regex to find func Autoset() { ... } until the next func or end of file
# This is risky with regex if braces aren't balanced, but we know the structure from the read.
# The duplicate Autoset is followed by LogGet in the file content we saw.

# Alternative: Read the file, find the line "func Autoset() {", and find the matching closing brace?
# Or simpler: Split by "func " and remove the block that starts with "Autoset()".

lines = content.split('\n')
new_lines = []
skip = False
brace_count = 0
in_duplicate = False

for line in lines:
    if line.strip().startswith("func Autoset() {"):
        in_duplicate = True
        brace_count = 1
        continue
    
    if in_duplicate:
        brace_count += line.count('{')
        brace_count -= line.count('}')
        if brace_count == 0:
            in_duplicate = False
        continue
    
    new_lines.append(line)

content = '\n'.join(new_lines)

# 2. Fix variable references in the remaining code (AutoSet and LogGet)
# We need to replace local variables with botstate globals.
# Mappings based on botstate/state.go and usage in background.go

replacements = [
    (r'\bfilterWar\.', 'botstate.FilterWar.'),
    (r'\bbacklist\.', 'botstate.Backlist.'),
    (r'\bAutoPro\b', 'botstate.AutoPro'),
    (r'\bNkick\b', 'botstate.Nkick'),
    (r'\bfilterop\b', 'botstate.Filterop'),
    (r'\boplist\b', 'botstate.Oplist'),
    (r'\bCeknuke\b', 'botstate.Ceknuke'),
    (r'\bcekoptime\b', 'botstate.Cekoptime'),
    (r'\bPurgeOP\b', 'botstate.PurgeOP'),
    (r'\bfiltermsg\b', 'botstate.Filtermsg'),
    (r'\bopjoin\b', 'botstate.Opjoin'),
    (r'\bCekpurge\b', 'botstate.Cekpurge'),
    (r'\bAutoproN\b', 'botstate.AutoproN'),
    (r'\bcekGo\b', 'botstate.CekGo'),
    (r'\bbotleave\b', 'botstate.Botleave'),
    (r'\baclear\b', 'botstate.Aclear'),
    (r'\bTimeSave\b', 'botstate.TimeSave'),
    (r'\bTimeBackup\b', 'botstate.TimeBackup'),
    (r'\bTimeClear\b', 'botstate.TimeClear'),
    (r'\bAutoBc\b', 'botstate.AutoBc'),
    (r'\bTimeBc\b', 'botstate.TimeBc'),
    (r'\bTimeBroadcast\b', 'botstate.TimeBroadcast'),
    (r'\bClientBot\b', 'botstate.ClientBot'),
    (r'\bTypebc\b', 'botstate.Typebc'),
    (r'\bMsgBroadcast\b', 'botstate.MsgBroadcast'),
    (r'\bData\.', 'botstate.Data.'), # In SaveProHistory
    (r'\bAllowDoOnce\b', 'botstate.AllowDoOnce'), # In SaveProHistory and CheckChatBan
    # LogGet specific
    (r'\bLastinvite\.', 'botstate.Lastinvite.'), # Might already be correct in LogGet
]

# Apply replacements
for old, new in replacements:
    content = re.sub(old, new, content)

# 3. Fix method casing for filterWar.clear() -> .Clear()
content = content.replace("botstate.FilterWar.clear()", "botstate.FilterWar.Clear()")

# 4. Check for MemBan and MemUser calls. 
# Since they are in the same package (handler), MemBan and MemUser calls are valid without prefix if they are defined in helpers.go (package handler).
# However, in line 498 of original file: !handler.MemBan(room.Id, l)
# If we are in package handler, "handler.MemBan" is invalid (circular/self ref issue if not imported as such, but usually just MemBan).
# We should remove "handler." prefix if it exists.
content = content.replace("handler.MemBan", "MemBan")

# 5. Check imports
# Ensure "fmt", "os/exec", "sort", "time", "strings" are imported.
# Ensure "../botstate", "../library/linetcr", "../library/hashmap", "../utils", "../library/SyncService" are imported.
# LogGet uses SyncService.Operation, so we need that import.

imports_to_check = [
    '"fmt"',
    '"os/exec"',
    '"sort"',
    '"time"',
    '"strings"',
    '"../botstate"',
    '"../library/linetcr"',
    '"../library/hashmap"',
    '"../utils"',
    '"../library/SyncService"'
]

# Simple check to add missing imports
import_block_match = re.search(r'import \((.*?)\)', content, re.DOTALL)
if import_block_match:
    current_imports = import_block_match.group(1)
    new_imports = current_imports
    for imp in imports_to_check:
        if imp not in current_imports:
            new_imports += "\n\t" + imp
    content = content.replace(current_imports, new_imports)

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)

print("Refactored handler/background.go")
