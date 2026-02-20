
import re

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'

funcs_to_remove = [
    'KickBan132V2', 'CancelBan125V2', 'KickBan132', 'CancelBan125',
    'QrKick', 'Ungban', 'Addgban', 'Joinsave', 'RandomToString',
    'MaxRevision', 'SendBigImage', 'SendMycreator', 'SendMymaker',
    'SendMyseller', 'SendMybuyer', 'CheckKickUser', 'MemEx',
    'MemUserN', 'MemAccsess', 'CancelAllCek', 'GetFuck', 'GetFuckV1',
    'CancelAll', 'PurgeAct'
]

with open(main_path, 'r', encoding='utf-8') as f:
    lines = f.readlines()

new_lines = []
skip = False
brace_count = 0
removed_funcs = []

i = 0
while i < len(lines):
    line = lines[i]
    stripped = line.strip()
    
    # Check for function start
    # We match `func Name(` at the start of the line (allowing for indentation? No, top level funcs shouldn't be indented usually)
    # But let's allow whitespace just in case.
    match = re.match(r'^\s*func\s+([A-Za-z0-9_]+)\(', line)
    
    if match and not skip:
        func_name = match.group(1)
        if func_name in funcs_to_remove:
            skip = True
            brace_count = 0
            removed_funcs.append(func_name)
            # print(f"Removing {func_name} at line {i+1}")
    
    if skip:
        brace_count += line.count('{')
        brace_count -= line.count('}')
        if brace_count <= 0:
            skip = False
            # We don't append the last line of the function either
    else:
        new_lines.append(line)
    
    i += 1

# Update content with replaced calls
content = "".join(new_lines)

for func in funcs_to_remove:
    # Replace calls: Name( -> handler.Name(
    # Avoid replacing func Name( (but we removed definitions so it should be fine)
    # Be careful with partial matches.
    
    # Regex for call: \bName\(
    # Use negative lookbehind to ensure we don't double prefix if it's already handler.Name (though unlikely in main.go unless partially migrated)
    # Also we want to match `Name(` but not `SomeName(`.
    pattern = r'(?<!handler\.)\b' + func + r'\('
    replacement = f'handler.{func}('
    content = re.sub(pattern, replacement, content)

with open(main_path, 'w', encoding='utf-8') as f:
    f.write(content)

print(f"Removed {len(removed_funcs)} functions: {', '.join(removed_funcs)}")
