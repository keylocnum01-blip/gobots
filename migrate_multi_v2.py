import re
import os

# Configuration
SOURCE_FILE = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"
BASE_DIR = r"c:\Users\Home\Desktop\LineBotProtect\gobots"

# List of (old_name, dest_relative_path, new_name)
# new_name is optional, if None, use TitleCase(old_name)
MIGRATIONS = [
    ("AddContact2", "handler/contact.go", "AddContact2"),
    ("NotifBot", "handler/logging.go", "NotifBot"),
    ("LogFight", "handler/logging.go", "LogFight"),
    ("deljoin", "handler/info.go", "DelJoin"),
    ("cekKick", "handler/info.go", "CekKick"),
    ("CekPurge", "handler/info.go", "CekPurge"),
    ("cekjoin", "handler/info.go", "CekJoin"),
    ("getBot", "handler/info.go", "GetBot"),
    ("cekOp2", "handler/moderation.go", "CekOp2"),
]

def read_file(path):
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def write_file(path, content):
    with open(path, "w", encoding="utf-8") as f:
        f.write(content)

def main():
    main_content = read_file(SOURCE_FILE)
    
    # Cache dest contents
    dest_contents = {}
    modified_dests = set()
    
    for old_name, dest_rel, new_name in MIGRATIONS:
        dest_path = os.path.join(BASE_DIR, dest_rel)
        if dest_path not in dest_contents:
            if os.path.exists(dest_path):
                dest_contents[dest_path] = read_file(dest_path)
            else:
                print(f"Destination {dest_path} does not exist!")
                continue
        
        # 1. Find definition in main.go
        pattern_def = r'func\s+' + re.escape(old_name) + r'\s*\('
        match = re.search(pattern_def, main_content)
        
        if match:
            print(f"Found definition for {old_name}")
            start_idx = match.start()
            
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
                new_func_body = re.sub(r'\b' + re.escape(old_name) + r'\(', f'{new_name}(', new_func_body)
                
                # Check if it already exists in destination
                if f"func {new_name}(" in dest_contents[dest_path]:
                    print(f"Function {new_name} already exists in {dest_rel}. Skipping append.")
                else:
                    dest_contents[dest_path] += "\n\n" + new_func_body
                    modified_dests.add(dest_path)
                    print(f"Appended {new_name} to {dest_rel}")
                
                # 3. Remove definition from main_content
                main_content = main_content[:start_idx] + main_content[end_idx:]
                
                # 4. Update calls in main_content
                main_content = re.sub(r'\b' + re.escape(old_name) + r'\(', f'handler.{new_name}(', main_content)
                
            else:
                print(f"Error: Unbalanced braces for {old_name}")
        else:
            print(f"Definition not found for {old_name}")

    # Write changes
    for dest_path in modified_dests:
        write_file(dest_path, dest_contents[dest_path])
        print(f"Updated {dest_path}")
    
    write_file(SOURCE_FILE, main_content)
    print("Updated source file.")

if __name__ == "__main__":
    main()
