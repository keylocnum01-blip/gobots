import re
import os

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
logic_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\logic.go'

functions_to_move = [
    'CanceljoinBot', 'Nukjoin', 'AutopurgeEnd', 'RemoveSticker', 'GETgrade'
]

functions_to_remove_only = [
    'panicHandle', 'GetMentionData'
]

with open(main_path, 'r', encoding='utf-8') as f:
    main_content = f.read()

extracted_code = ""

# Helper to find and extract function
def extract_function(content, func_name):
    pattern = re.compile(r'func ' + func_name + r'\s*\(', re.MULTILINE)
    match = pattern.search(content)
    if not match:
        return None, content
    
    start_index = match.start()
    brace_count = 0
    end_index = -1
    found_start_brace = False
    
    for i in range(start_index, len(content)):
        if content[i] == '{':
            brace_count += 1
            found_start_brace = True
        elif content[i] == '}':
            brace_count -= 1
            if found_start_brace and brace_count == 0:
                end_index = i + 1
                break
    
    if end_index != -1:
        func_code = content[start_index:end_index]
        return func_code, content # We don't remove yet to keep indices valid if we were doing in-place, but here we just extract
    return None, content

for func in functions_to_move:
    code, _ = extract_function(main_content, func)
    if code:
        extracted_code += code + "\n\n"

# Fix references in extracted code
extracted_code = extracted_code.replace('config.config.config.', 'config.')
extracted_code = extracted_code.replace('config.config.', 'config.')
extracted_code = extracted_code.replace('panicHandle', 'utils.PanicHandle')
# In RemoveSticker, replace 'config.Stickers' if needed (it might be correct already after above replacement)

# Create logic.go
logic_content = """package handler

import (
	"fmt"
	"sync"
	"../botstate"
	"../config"
	"../library/linetcr"
	"../utils"
)

""" + extracted_code

with open(logic_path, 'w', encoding='utf-8') as f:
    f.write(logic_content)
print(f"Created {logic_path}")

# Remove functions from main.go
def remove_function_def(content, func_name):
    pattern = re.compile(r'func ' + func_name + r'\s*\(', re.MULTILINE)
    match = pattern.search(content)
    if not match:
        return content
    start_index = match.start()
    brace_count = 0
    end_index = -1
    found_start_brace = False
    for i in range(start_index, len(content)):
        if content[i] == '{':
            brace_count += 1
            found_start_brace = True
        elif content[i] == '}':
            brace_count -= 1
            if found_start_brace and brace_count == 0:
                end_index = i + 1
                break
    if end_index != -1:
        return content[:start_index] + content[end_index:]
    return content

for func in functions_to_move + functions_to_remove_only:
    main_content = remove_function_def(main_content, func)

# Replace calls in main.go
# 1. logic functions -> handler.Func
for func in functions_to_move:
    pattern = re.compile(r'\b' + func + r'\(')
    main_content = pattern.sub('handler.' + func + '(', main_content)

# 2. utils functions -> utils.Func
main_content = main_content.replace('panicHandle(', 'utils.PanicHandle(')
main_content = main_content.replace('GetMentionData(', 'utils.GetMentionData(')

# Fix config.config...
main_content = main_content.replace('config.config.config.', 'config.')
main_content = main_content.replace('config.config.', 'config.')

with open(main_path, 'w', encoding='utf-8') as f:
    f.write(main_content)
print(f"Updated {main_path}")
