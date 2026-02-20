import os
import re

# Mappings: Old variable name -> New variable name (with package prefix)
mappings = {
    "cekGo": "botstate.CekGo",
    "filtermsg": "botstate.FilterMsg",
    "Nkick": "botstate.Nkick",
    "UserBot": "botstate.UserBot",
    "Cekstaybot": "botstate.Cekstaybot",
    # Case sensitive checks for other variations if they exist, but Go is case sensitive.
}

# Files to process
files_to_process = [
    "main.go",
    "handler/system.go",
    "handler/helpers.go",
    "handler/info.go",
    "handler/actions.go",
    "handler/background.go",
    "handler/moderation.go",
    "handler/runbot.go",
    "handler/invite.go",
    "handler/contact.go",
    "handler/logging.go",
]

def process_file(file_path):
    if not os.path.exists(file_path):
        print(f"File not found: {file_path}")
        return

    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()

    original_content = content
    
    # 1. Update references
    for old, new in mappings.items():
        # Regex to match the variable name but not if it's already prefixed with botstate.
        # And ensure we match whole words.
        # Pattern: (not botstate\.)\bOldName\b
        
        # However, in main.go, we might have declarations.
        # If it's main.go, we handle declarations separately or just comment them out later.
        # For now, let's replace usages.
        
        # We need to be careful not to replace struct fields if they have the same name (unlikely for these specific names).
        
        # Replace usages:
        # Look for word boundary, OldName, word boundary.
        # Avoid matching if preceded by . (field access) unless it's the struct itself? 
        # Actually UserBot is a variable.
        
        # Negative lookbehind (?<!\.) to ensure it's not a field access like something.Nkick
        # Negative lookbehind (?<!botstate\.) to ensure it's not already migrated.
        
        pattern = r'(?<!\.)(?<!botstate\.)\b' + re.escape(old) + r'\b'
        
        # If we are in main.go, we don't want to replace the declaration line if we are going to remove it later.
        # But if we replace it, the declaration becomes "var botstate.Nkick = ...", which is invalid syntax.
        # So we should probably remove declarations first or identify them.
        
        content = re.sub(pattern, new, content)

    if content != original_content:
        # Add import botstate if missing and we added botstate references
        if "botstate." in content and '"./botstate"' not in content and '"../botstate"' not in content:
             # Add import
             # This is a simple heuristic, might need manual adjustment for import block
             if "package main" in content:
                 content = content.replace('import (', 'import (\n\t"./botstate"')
             else:
                 content = content.replace('import (', 'import (\n\t"../botstate"')
        
        with open(file_path, 'w', encoding='utf-8') as f:
            f.write(content)
        print(f"Updated references in {file_path}")

def remove_declarations_main():
    file_path = "main.go"
    if not os.path.exists(file_path):
        return

    with open(file_path, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    new_lines = []
    # Variables to remove declarations for
    vars_to_remove = mappings.keys()
    
    # We look for lines starting with "var Name =" or inside var ( ... Name = ... )
    
    for line in lines:
        stripped = line.strip()
        remove = False
        for var in vars_to_remove:
            # Check for "var VarName =" or "VarName =" (inside var block)
            # We already replaced references in main.go so they might look like "botstate.VarName ="
            # Wait, process_file replaced them in main.go too. 
            # So the declarations in main.go likely look like:
            # var botstate.Nkick = ... (invalid)
            # or botstate.Nkick = ... (inside var block)
            
            check_var = f"botstate.{var}"
            
            if stripped.startswith(f"var {check_var} ") or stripped.startswith(f"{check_var} ") or stripped.startswith(f"{check_var}=") or stripped.startswith(f"var {check_var}="):
                remove = True
                print(f"Removing declaration: {stripped}")
                break
        
        if not remove:
            new_lines.append(line)

    with open(file_path, 'w', encoding='utf-8') as f:
        f.writelines(new_lines)
    print("Removed declarations from main.go")

if __name__ == "__main__":
    for fp in files_to_process:
        process_file(fp)
    
    remove_declarations_main()
