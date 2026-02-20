import os
import re

# Mappings: function_name -> (target_file, new_name)
# If new_name is None, use TitleCase of function_name
mappings = {
    # actions.go
    "clone": ("handler/actions.go", "Clone"),
    "back": ("handler/actions.go", "Back"),
    "Upsetcmd": ("handler/actions.go", "Upsetcmd"),
    "addwl": ("handler/actions.go", "Addwl"),
    "SelectBot": ("handler/actions.go", "SelectBot"),
    "CheckBot": ("handler/actions.go", "CheckBot"),
    "GetKorban": ("handler/actions.go", "GetKorban"),
    "squadMention": ("handler/actions.go", "SquadMention"),
    "detectSquad": ("handler/actions.go", "DetectSquad"), # Already exists, just update call/remove
    
    # info.go
    "GetComs": ("handler/info.go", "GetComs"),
    "CheckAccount": ("handler/info.go", "CheckAccount"),
    "Checkuser": ("handler/info.go", "CheckUser"),
    "Allbotlist": ("handler/info.go", "AllBotList"),
    
    # contact.go
    "Checkqr": ("handler/contact.go", "CheckQr"),
    "Cmdlistcheck": ("handler/contact.go", "CmdListCheck"),
    "addCon": ("handler/contact.go", "AddCon"),
    "addConSq": ("handler/contact.go", "AddConSq"),
    "addConSqV2": ("handler/contact.go", "AddConSqV2"),
    "addConSingle": ("handler/contact.go", "AddConSingle"),
    
    # invite.go
    "AcceptJoin": ("handler/invite.go", "AcceptJoin"),
    "AcceptJoinV2": ("handler/invite.go", "AcceptJoinV2"),
    "CanceljoinBot": ("handler/invite.go", "CancelJoinBot"),
    "Nukjoin": ("handler/invite.go", "NukJoin"),
    "AutopurgeEnd": ("handler/invite.go", "AutoPurgeEnd"),
    "CancelEnd": ("handler/invite.go", "CancelEnd"),
    "Setpurgealln": ("handler/invite.go", "SetPurgeAllN"),
    
    # helpers.go
    "InArrayChat": ("handler/helpers.go", "InArrayChat"),
    "InArrayInt64": ("handler/helpers.go", "InArrayInt64"),
    "contains": ("handler/helpers.go", "Contains"),
    "bToMb": ("handler/helpers.go", "BToMb"),
    "TimeDown": ("handler/helpers.go", "TimeDown"),
    "StripOut": ("handler/helpers.go", "StripOut"),
    "gettxt": ("handler/helpers.go", "GetTxt"),
    "CheckMessage": ("handler/helpers.go", "CheckMessage"),
    "GETgrade": ("handler/helpers.go", "GetGrade"),
    "RemoveSticker": ("handler/helpers.go", "RemoveSticker"),
    "AppendLastSticker": ("handler/helpers.go", "AppendLastSticker"),
    "getArg": ("handler/helpers.go", "GetArg"),
    "getKey": ("handler/helpers.go", "GetKey"),
    "MemAccsess": ("handler/helpers.go", "MemAccsess"),
    "MemUserN": ("handler/helpers.go", "MemUserN"),
    "MemEx": ("handler/helpers.go", "MemEx"),
    "Checkkickuser": ("handler/helpers.go", "CheckKickUser"),
    
    # system.go
    "SaveBackup": ("handler/system.go", "SaveBackup"),
    "gracefulShutdown": ("handler/system.go", "GracefulShutdown"), # Already exists
    "Resprem": ("handler/system.go", "Resprem"), # Already exists
    
    # logging.go
    "LogOp": ("handler/logging.go", "LogOp"),
    "LogLast": ("handler/logging.go", "LogLast"), # Already exists
}

main_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"
root_dir = r"c:\Users\Home\Desktop\LineBotProtect\gobots"

# 1. Read main.go
with open(main_path, "r", encoding="utf-8") as f:
    main_content = f.read()

# 2. Extract functions
lines = main_content.split('\n')
extracted_funcs = {} # func_name -> lines

current_func = None
brace_count = 0
func_lines = []

for line in lines:
    stripped = line.strip()
    
    # Detect start of function
    if stripped.startswith("func ") and "{" in line:
        # Extract name
        parts = stripped.split(" ")
        if len(parts) >= 2:
            name_part = parts[1]
            func_name = name_part.split("(")[0]
            
            if func_name in mappings:
                current_func = func_name
                brace_count = 0
                func_lines = []
    
    if current_func:
        func_lines.append(line)
        brace_count += line.count('{')
        brace_count -= line.count('}')
        
        if brace_count == 0:
            extracted_funcs[current_func] = "\n".join(func_lines)
            current_func = None
            func_lines = []

# 3. Append to target files
for func_name, (target_rel, new_name) in mappings.items():
    if func_name not in extracted_funcs:
        # It might be that the function is already removed or I missed it.
        # Or it is one of the "Already exists" ones.
        continue
    
    # Check if target function already exists in target file
    target_path = os.path.join(root_dir, target_rel)
    if not os.path.exists(target_path):
        print(f"Warning: Target file {target_path} does not exist.")
        continue
        
    with open(target_path, "r", encoding="utf-8") as f:
        target_content = f.read()
    
    if f"func {new_name}(" in target_content:
        print(f"Function {new_name} already exists in {target_rel}. Skipping append.")
    else:
        # Modify function definition to use new name
        func_code = extracted_funcs[func_name]
        func_code = func_code.replace(f"func {func_name}(", f"func {new_name}(", 1)
        
        # Append
        with open(target_path, "a", encoding="utf-8") as f:
            f.write("\n\n" + func_code)
        print(f"Appended {new_name} to {target_rel}")

# 4. Remove from main.go
new_main_lines = []
skip = False
brace_count = 0

for line in lines:
    stripped = line.strip()
    
    start_func = False
    for func_name in mappings:
        if stripped.startswith(f"func {func_name}("):
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
        
    new_main_lines.append(line)

main_content = "\n".join(new_main_lines)

# 5. Update calls in main.go
for func_name, (target_rel, new_name) in mappings.items():
    # Basic replacement: func_name( -> handler.new_name(
    # But watch out for substrings.
    # Regex is safer: \bfunc_name\(
    
    pattern = r'\b' + re.escape(func_name) + r'\('
    replacement = f"handler.{new_name}("
    main_content = re.sub(pattern, replacement, main_content)

# Fix double handler prefixes
main_content = main_content.replace("handler.handler.", "handler.")

with open(main_path, "w", encoding="utf-8") as f:
    f.write(main_content)

print("Migration complete.")
