import re
import os

FILE_PATH = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\actions.go"
FUNCTION_NAME = "NotifBot"

def read_file(path):
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def write_file(path, content):
    with open(path, "w", encoding="utf-8") as f:
        f.write(content)

def main():
    content = read_file(FILE_PATH)
    
    pattern_def = r'func\s+' + re.escape(FUNCTION_NAME) + r'\s*\('
    match = re.search(pattern_def, content)
    
    if match:
        print(f"Found definition for {FUNCTION_NAME}")
        start_idx = match.start()
        
        brace_search = re.search(r'\{', content[start_idx:])
        if not brace_search:
            print(f"Error: No opening brace for {FUNCTION_NAME}")
            return
            
        brace_open_idx = start_idx + brace_search.start()
        
        balance = 1
        i = brace_open_idx + 1
        while i < len(content) and balance > 0:
            char = content[i]
            if char == '{':
                balance += 1
            elif char == '}':
                balance -= 1
            i += 1
        
        if balance == 0:
            end_idx = i
            # Remove function
            content = content[:start_idx] + content[end_idx:]
            print(f"Removed {FUNCTION_NAME} from {FILE_PATH}")
            write_file(FILE_PATH, content)
        else:
            print(f"Error: Unbalanced braces for {FUNCTION_NAME}")
    else:
        print(f"Definition not found for {FUNCTION_NAME}")

if __name__ == "__main__":
    main()
