import re
import os

main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
runbot_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\runbot.go'

with open(main_path, 'r', encoding='utf-8') as f:
    main_content = f.read()

# Helper to extract function
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
        return func_code, content
    return None, content

runbot_code, _ = extract_function(main_content, 'RunBot')

if runbot_code:
    # Remove handler. prefix
    runbot_code = runbot_code.replace('handler.', '')
    
    # Determine imports
    imports = [
        '"fmt"',
        '"sync"',
        '"strings"',
        '"time"',
        '"runtime"',
        '"../botstate"',
        '"../config"',
        '"../library/linetcr"',
        '"../utils"',
        '"../library/SyncService"',
    ]
    
    # Check for optional imports
    if 'ants.' in runbot_code: imports.append('"github.com/panjf2000/ants"')
    if 'mod.' in runbot_code: imports.append('mod "../library/modcompact"')
    if 'call.' in runbot_code: imports.append('call "../library/libcall/call"')
    if 'namegenerator.' in runbot_code: imports.append('namegenerator "../library/RandomName"')
    if 'talkservice.' in runbot_code: imports.append('talkservice "../library/linethrift"')
    if 'valid.' in runbot_code: imports.append('valid "github.com/asaskevich/govalidator"')
    if 'gjson.' in runbot_code: imports.append('"github.com/tidwall/gjson"')
    if 'mem.' in runbot_code: imports.append('"github.com/shirou/gopsutil/mem"')
    if 'regexp.' in runbot_code: imports.append('"regexp"')
    if 'log.' in runbot_code: imports.append('"log"')
    if 'osext.' in runbot_code: imports.append('"github.com/kardianos/osext"')
    if 'net.' in runbot_code: imports.append('"net"')
    if 'http.' in runbot_code: imports.append('"net/http"')
    if 'debug.' in runbot_code: imports.append('"runtime/debug"')
    if 'sort.' in runbot_code: imports.append('"sort"')
    if 'json.' in runbot_code: imports.append('"encoding/json"')
    if 'ioutil.' in runbot_code: imports.append('"io/ioutil"')
    if 'rand.' in runbot_code: imports.append('"math/rand"')
    if 'exec.' in runbot_code: imports.append('"os/exec"')
    if 'syscall.' in runbot_code: imports.append('"syscall"')
    if 'strconv.' in runbot_code: imports.append('"strconv"')
    if 'os.' in runbot_code: imports.append('"os"')
    if 'hashmap.' in runbot_code: imports.append('"../library/hashmap"')
    if 'unistyle.' in runbot_code: imports.append('"../library/unistyle"')

    # Create runbot.go
    runbot_content = f"""package handler

import (
    {chr(10).join(imports)}
)

{runbot_code}
"""
    with open(runbot_path, 'w', encoding='utf-8') as f:
        f.write(runbot_content)
    print(f"Created {runbot_path}")

    # Remove RunBot from main.go
    def remove_function_def(content, func_name):
        pattern = re.compile(r'func ' + func_name + r'\s*\(', re.MULTILINE)
        match = pattern.search(content)
        if not match: return content
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

    main_content = remove_function_def(main_content, 'RunBot')

    # Replace call
    main_content = main_content.replace('go RunBot(', 'go handler.RunBot(')

    with open(main_path, 'w', encoding='utf-8') as f:
        f.write(main_content)
    print(f"Updated {main_path}")
else:
    print("RunBot not found")
