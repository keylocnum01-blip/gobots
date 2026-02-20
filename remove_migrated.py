import re
import os

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'

functions_to_remove = [
    'CheckExprd',
    'CekDuedate',
    'CheckLastActive',
]

with open(main_path, 'r', encoding='utf-8') as f:
    content = f.read()

def remove_function(content, func_name):
    # Regex to find function definition
    # func FuncName(...) ... {
    pattern = r'func\s+' + re.escape(func_name) + r'\s*\([^{]*\)[^{]*\{'
    match = re.search(pattern, content)
    if not match:
        print(f"Function {func_name} not found.")
        return content
    
    start_index = match.start()
    brace_count = 0
    in_function = False
    end_index = -1
    
    # Find the closing brace
    # Start scanning from the opening brace of the function
    # The regex match ends at '{'
    open_brace_index = match.end() - 1
    
    for i in range(open_brace_index, len(content)):
        if content[i] == '{':
            brace_count += 1
            in_function = True
        elif content[i] == '}':
            brace_count -= 1
        
        if in_function and brace_count == 0:
            end_index = i + 1
            break
            
    if end_index != -1:
        # Remove the function block
        print(f"Removing {func_name}...")
        return content[:start_index] + content[end_index:]
    else:
        print(f"Could not find closing brace for {func_name}.")
        return content

for func in functions_to_remove:
    content = remove_function(content, func)

with open(main_path, 'w', encoding='utf-8') as f:
    f.write(content)

print("Finished removing remaining functions.")
