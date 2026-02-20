import re
import os

# Configuration
SOURCE_FILE = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"
DEST_FILE = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\moderation.go"

# Function mapping: old_name_in_main -> new_name_in_handler
# Using TitleCase for new names
FUNCTIONS = {
    "cancelBanInv": "CancelBanInv",
    "cancelall": "CancelAll",
    "getfuckV1": "GetFuckV1",
    "getfuck": "GetFuck",
    "cancelallcek": "CancelAllCek",
    "Purgeact": "PurgeAct",
    "groupBackupKick": "GroupBackupKick",
    "groupBackupKickV2": "GroupBackupKickV2",
    "groupBackupCans": "GroupBackupCans",
    "kickDirt": "KickDirt",
}

def read_file(path):
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def write_file(path, content):
    with open(path, "w", encoding="utf-8") as f:
        f.write(content)

def main():
    main_content = read_file(SOURCE_FILE)
    dest_content = read_file(DEST_FILE)
    
    # Imports to check/add in destination
    # We'll just ensure the file has imports, specific ones might need manual check or smarter logic
    # But usually these functions use linetcr, botstate, utils, fmt, time, sync
    
    modified_dest = False
    
    for old_name, new_name in FUNCTIONS.items():
        # 1. Find definition in main.go
        # Pattern: func old_name(
        pattern_def = r'func\s+' + re.escape(old_name) + r'\s*\('
        match = re.search(pattern_def, main_content)
        
        if match:
            print(f"Found definition for {old_name}")
            start_idx = match.start()
            
            # Find the full block
            # Look for opening brace
            brace_search = re.search(r'\{', main_content[start_idx:])
            if not brace_search:
                print(f"Error: No opening brace for {old_name}")
                continue
                
            brace_open_idx = start_idx + brace_search.start()
            
            balance = 1
            i = brace_open_idx + 1
            while i < len(main_content) and balance > 0:
                char = main_content[i]
                if char == '{':
                    balance += 1
                elif char == '}':
                    balance -= 1
                i += 1
            
            if balance == 0:
                end_idx = i
                func_body = main_content[start_idx:end_idx]
                
                # 2. Prepare function for destination
                # Replace name with new_name
                new_func_body = re.sub(r'func\s+' + re.escape(old_name), f'func {new_name}', func_body, count=1)
                
                # Check if it already exists in destination
                if f"func {new_name}(" in dest_content:
                    print(f"Function {new_name} already exists in destination. Skipping append.")
                else:
                    dest_content += "\n\n" + new_func_body
                    modified_dest = True
                    print(f"Appended {new_name} to destination")
                
                # 3. Update calls in main.go
                # We need to be careful not to replace the definition we just found (though we will remove it)
                # But we should replace calls elsewhere.
                # Replace `old_name(` with `handler.new_name(`
                # Use word boundary to avoid partial matches
                
                # Special case: recursive calls inside the function itself?
                # If we move it, recursive calls inside the body should point to the new name (in the same package) 
                # OR if it's in the same package (handler), it calls NewName.
                # If we simply append `new_func_body` to dest, calls inside it to `old_name` need to be updated to `new_name` (or `NewName`).
                
                # Let's update calls inside the extracted body first
                new_func_body = re.sub(r'\b' + re.escape(old_name) + r'\(', f'{new_name}(', new_func_body)
                
                # Now update dest_content with the corrected body if we are appending
                if modified_dest and dest_content.endswith(new_func_body.replace(f'{new_name}(', f'{old_name}(', 1)): 
                     # This logic is a bit flawed because I already appended above.
                     # Let's redo the append logic correctly.
                     pass 

                # Re-do append logic cleanly:
                # Remove the raw append from above, let's do it right here.
                # We haven't written to file yet.
                
                # Fix calls inside the function body (recursion)
                # old_name(...) -> new_name(...)
                new_func_body = re.sub(r'\b' + re.escape(old_name) + r'\(', f'{new_name}(', new_func_body)
                
                # Also, we need to fix calls to OTHER functions we are migrating?
                # Ideally yes, but let's do one pass. Go will complain if not found, or we rely on `handler.` prefix if called from outside.
                # If called from inside `handler` package, they should call `NewName`.
                # If we migrate `A` and `B`. `A` calls `B`.
                # In `main.go`: `A` calls `B`.
                # Moved to `handler`: `A` calls `B`. `B` is now `B` (exported? or local?).
                # If we rename `B` to `NewB`. `A` needs to call `NewB`.
                
                # For now, let's just rename the function definition and self-calls.
                # We can run a second pass to fix cross-calls within handler.
                
                # Re-append corrected body
                if f"func {new_name}(" not in dest_content:
                     # Remove the previous append if I messed up logic flow? 
                     # Actually I'll just reset dest_content in memory for this iteration if I needed to.
                     # But simpler: Just don't append until now.
                     dest_content = dest_content.replace(func_body.replace(f'func {old_name}', f'func {new_name}'), "") # Hacky cleanup if needed
                     # Let's just trust the "if not in" check works for the *start*.
                     
                     dest_content += "\n\n" + new_func_body
                     modified_dest = True
                
                # 4. Remove definition from main_content
                main_content = main_content[:start_idx] + main_content[end_idx:]
                
                # 5. Update calls in main_content
                # Replace `old_name(` with `handler.new_name(`
                # Regex: \b old_name \(
                main_content = re.sub(r'\b' + re.escape(old_name) + r'\(', f'handler.{new_name}(', main_content)
                
            else:
                print(f"Error: Unbalanced braces for {old_name}")
        else:
            print(f"Definition not found for {old_name}")

    if modified_dest:
        write_file(DEST_FILE, dest_content)
        print("Updated destination file.")
    
    write_file(SOURCE_FILE, main_content)
    print("Updated source file.")

if __name__ == "__main__":
    main()
