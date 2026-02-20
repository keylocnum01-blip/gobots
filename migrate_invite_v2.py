import re
import os

# Configuration
SOURCE_FILE = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"
DEST_FILE = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\invite.go"

# Function mapping: old_name_in_main -> new_name_in_handler
# Using TitleCase for new names
FUNCTIONS = {
    "Setinviteto": "SetInviteTo",
    "invBackup": "InvBackup",
    "openqr": "OpenQR",
    "getTicket": "GetTicket",
    "qrBackup": "QrBackup",
    "groupBackupInv2": "GroupBackupInv2",
    "groupBackupInv": "GroupBackupInv",
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
    
    modified_dest = False
    
    for old_name, new_name in FUNCTIONS.items():
        # 1. Find definition in main.go
        pattern_def = r'func\s+' + re.escape(old_name) + r'\s*\('
        match = re.search(pattern_def, main_content)
        
        if match:
            print(f"Found definition for {old_name}")
            start_idx = match.start()
            
            # Find the full block
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
                new_func_body = re.sub(r'func\s+' + re.escape(old_name), f'func {new_name}', func_body, count=1)
                
                # Fix self-calls inside the function body
                new_func_body = re.sub(r'\b' + re.escape(old_name) + r'\(', f'{new_name}(', new_func_body)
                
                # Check if it already exists in destination
                if f"func {new_name}(" in dest_content:
                    print(f"Function {new_name} already exists in destination. Skipping append.")
                else:
                    dest_content += "\n\n" + new_func_body
                    modified_dest = True
                    print(f"Appended {new_name} to destination")
                
                # 3. Remove definition from main_content
                main_content = main_content[:start_idx] + main_content[end_idx:]
                
                # 4. Update calls in main_content
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
