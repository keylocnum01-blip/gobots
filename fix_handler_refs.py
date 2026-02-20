
import os
import re

handler_dir = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler'

for filename in os.listdir(handler_dir):
    if filename.endswith(".go"):
        filepath = os.path.join(handler_dir, filename)
        with open(filepath, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Replace handler.Func( with Func(
        # We assume function names start with Uppercase.
        # Pattern: handler\.([A-Z][a-zA-Z0-9_]*)
        
        def replace_func(match):
            func_name = match.group(1)
            return func_name
        
        new_content = re.sub(r'handler\.([A-Z][a-zA-Z0-9_]*)', replace_func, content)
        
        if new_content != content:
            print(f"Updating {filename}...")
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(new_content)
        else:
            print(f"No changes in {filename}")

print("Handler refs fixed.")
