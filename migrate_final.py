import os
import re

def read_file(path):
    with open(path, 'r', encoding='utf-8') as f:
        return f.read()

def write_file(path, content):
    with open(path, 'w', encoding='utf-8') as f:
        f.write(content)

def find_func_end(content, start_index):
    brace_count = 0
    in_func = False
    for i in range(start_index, len(content)):
        if content[i] == '{':
            brace_count += 1
            in_func = True
        elif content[i] == '}':
            brace_count -= 1
        
        if in_func and brace_count == 0:
            return i + 1
    return -1

def extract_and_remove(content, func_name):
    regex = r'func\s+' + re.escape(func_name) + r'\s*\('
    match = re.search(regex, content)
    if not match:
        return None, content
    
    start_pos = match.start()
    end_pos = find_func_end(content, start_pos)
    
    if end_pos == -1:
        print(f"Error: Could not find end of function {func_name}")
        return None, content
    
    func_code = content[start_pos:end_pos]
    new_content = content[:start_pos] + content[end_pos:]
    return func_code, new_content

def replace_calls(content, old_name, new_name):
    # Matches function calls: name(
    pattern = re.compile(r'\b' + old_name + r'\(')
    return pattern.sub(new_name + '(', content)

def main():
    root = r'c:\Users\Home\Desktop\LineBotProtect\gobots'
    main_path = os.path.join(root, 'main.go')
    mod_path = os.path.join(root, 'handler', 'moderation.go')
    help_path = os.path.join(root, 'handler', 'helpers.go')
    info_path = os.path.join(root, 'handler', 'info.go')
    logic_path = os.path.join(root, 'handler', 'logic.go')

    main_content = read_file(main_path)
    mod_content = read_file(mod_path)
    help_content = read_file(help_path)
    info_content = read_file(info_path)
    logic_content = read_file(logic_path)

    # 1. Moderation
    funcs_mod = {
        'RemBanFriends': 'RemBanFriends',
        'clearCon': 'ClearCon'
    }
    repl_mod = {
        'AllowDoOnce': 'botstate.AllowDoOnce',
        'ClientBot': 'botstate.ClientBot',
        'Squadlist': 'botstate.Squadlist',
        'DEVELOPER': 'botstate.DEVELOPER',
        'defer utils.PanicHandle': 'defer utils.PanicHandle', # ensure utils
        'fancy(': 'botstate.Fancy(',
        'DataMention(': 'utils.DataMention(', # Will move DataMention to utils/helpers or similar
        'linetcr.KickBanChat': 'linetcr.KickBanChat',
        'linetcr.GetBannedChat': 'linetcr.GetBannedChat',
        'linetcr.BanChatAdd': 'linetcr.BanChatAdd'
    }
    
    extracted_mod = []
    for func, new_name in funcs_mod.items():
        code, main_content = extract_and_remove(main_content, func)
        if code:
             for old, new in repl_mod.items():
                 code = code.replace(old, new)
             if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
             
             if f'func {new_name}' not in mod_content:
                 extracted_mod.append(code)
                 print(f"Extracted {func} -> {new_name} (moderation)")
    
    if extracted_mod:
        mod_content += "\n\n" + "\n\n".join(extracted_mod)
        write_file(mod_path, mod_content)


    # 2. Helpers
    funcs_help = {
        'DataMention': 'DataMention',
        'SelectallBot': 'SelectallBot'
    }
    repl_help = {
        'Squadlist': 'botstate.Squadlist',
        'GetKorban': 'botstate.GetKorban',
        'defer utils.PanicHandle': 'defer utils.PanicHandle',
    }
    
    extracted_help = []
    for func, new_name in funcs_help.items():
        code, main_content = extract_and_remove(main_content, func)
        if code:
             for old, new in repl_help.items():
                 code = code.replace(old, new)
             if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
             
             if f'func {new_name}' not in help_content:
                 extracted_help.append(code)
                 print(f"Extracted {func} -> {new_name} (helpers)")
    
    if extracted_help:
        help_content += "\n\n" + "\n\n".join(extracted_help)
        write_file(help_path, help_content)


    # 3. Info
    funcs_info = {
        'GenerateTimeLog': 'GenerateTimeLog'
    }
    repl_info = {
        'fancy(': 'botstate.Fancy('
    }
    
    extracted_info = []
    for func, new_name in funcs_info.items():
        code, main_content = extract_and_remove(main_content, func)
        if code:
             for old, new in repl_info.items():
                 code = code.replace(old, new)
             if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
             
             if f'func {new_name}' not in info_content:
                 extracted_info.append(code)
                 print(f"Extracted {func} -> {new_name} (info)")

    if extracted_info:
        info_content += "\n\n" + "\n\n".join(extracted_info)
        write_file(info_path, info_content)


    # 4. Logic
    funcs_logic = {
        'cekOp': 'CekOp'
    }
    repl_logic = {
        'oplist': 'botstate.Oplist'
    }
    
    extracted_logic = []
    for func, new_name in funcs_logic.items():
        code, main_content = extract_and_remove(main_content, func)
        if code:
             for old, new in repl_logic.items():
                 code = code.replace(old, new)
             if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
             
             if f'func {new_name}' not in logic_content:
                 extracted_logic.append(code)
                 print(f"Extracted {func} -> {new_name} (logic)")

    if extracted_logic:
        logic_content += "\n\n" + "\n\n".join(extracted_logic)
        write_file(logic_path, logic_content)


    # 5. Remove ReloginProgram (already in system.go)
    code, main_content = extract_and_remove(main_content, 'ReloginProgram')
    if code:
        print("Removed ReloginProgram from main.go")


    # 6. Update calls in main.go
    calls = {
        'RemBanFriends': 'handler.RemBanFriends',
        'clearCon': 'handler.ClearCon',
        'DataMention': 'handler.DataMention',
        'SelectallBot': 'handler.SelectallBot',
        'GenerateTimeLog': 'handler.GenerateTimeLog',
        'cekOp': 'handler.CekOp',
        'ReloginProgram': 'handler.ReloginProgram'
    }
    
    for old, new in calls.items():
        main_content = replace_calls(main_content, old, new)

    write_file(main_path, main_content)
    print("Done.")

if __name__ == '__main__':
    main()
