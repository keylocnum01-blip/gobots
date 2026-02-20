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
    
    # Find the first opening brace after start_index
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
    # Find "func func_name"
    # Allow for variations like "func  func_name ("
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
    helpers_path = os.path.join(root, 'handler', 'helpers.go') # Using handler/helpers.go

    main_content = read_file(main_path)
    mod_content = read_file(mod_path)
    helpers_content = read_file(helpers_path)

    # 1. Move to handler/moderation.go
    funcs_to_mod = {
        'KickCancelProtect': 'KickCancelProtect',
        'kickPelaku': 'KickPelaku',
        'kickProtect': 'KickProtect',
        'GhostEnd': 'GhostEnd',
        'Purgemode': 'Purgemode',
        'KickAllBan': 'KickAllBan',
        'IsBlacklist': 'IsBlacklist',
        'IsBlacklist2': 'IsBlacklist2'
    }

    replacements_mod = {
        'defer panicHandle': 'defer utils.PanicHandle',
        'Squadlist': 'botstate.Squadlist',
        'cekGo': 'botstate.CekGo',
        'Ajsjoin': 'botstate.Ajsjoin',
        'ProtectMode': 'botstate.ProtectMode',
        'Banned': 'botstate.Banned',
        'handler.MemBan': 'MemBan',
        'utils.InArrayString': 'utils.InArrayString',
        'InArrayInt64': 'utils.InArrayInt64',
        'qrGo22': 'QrGo22',
        'GetKorban': 'botstate.GetKorban',
        'stringToInt': 'botstate.StringToInt',
        'valid.Abs': 'valid.Abs',
        'func kickPelaku': 'func KickPelaku',
        'func kickProtect': 'func KickProtect',
        'kickPelaku(': 'KickPelaku(',
        'kickProtect(': 'KickProtect(',
        'FastKick(': 'FastKick(', 
        'FastCancel(': 'FastCancel('
    }

    extracted_mod = []
    
    for func, new_name in funcs_to_mod.items():
        code, main_content = extract_and_remove(main_content, func)
        if code:
            for old, new in replacements_mod.items():
                code = code.replace(old, new)
            if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
            
            if f'func {new_name}' not in mod_content:
                extracted_mod.append(code)
                print(f"Extracted {func}")
            else:
                print(f"Skipping {new_name}, already exists")
    
    if extracted_mod:
        mod_content += "\n\n" + "\n\n".join(extracted_mod)
        write_file(mod_path, mod_content)

    # 2. Move to handler/helpers.go
    funcs_to_help = {
        'randomToString': 'RandomToString'
    }
    replacements_help = {
        'stringToInt': 'botstate.StringToInt'
    }
    
    extracted_help = []
    for func, new_name in funcs_to_help.items():
        code, main_content = extract_and_remove(main_content, func)
        if code:
             for old, new in replacements_help.items():
                code = code.replace(old, new)
             if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
             
             if f'func {new_name}' not in helpers_content:
                extracted_help.append(code)
                print(f"Extracted {func}")
    
    if extracted_help:
        helpers_content += "\n\n" + "\n\n".join(extracted_help)
        write_file(helpers_path, helpers_content)

    # 3. Remove other duplicates
    to_remove = [
        'KickCancelBan129', 'FastKick', 'FastCancel', 'NodeBans', 'PurgeFaster', 'KickCancel',
        'KickMemBan', 'CanMemBan',
        'Ungban', 'Addgban', 'Joinsave', 'AddbanOp3',
        'panicHandle', 'GetMentionData', 'randomString', 'IndexOf', 'Checkmulti'
    ]
    
    for func in to_remove:
        code, main_content = extract_and_remove(main_content, func)
        if code:
            print(f"Removed {func}")

    # 4. Replace calls
    call_replacements = {
        'KickCancelProtect': 'handler.KickCancelProtect',
        'kickPelaku': 'handler.KickPelaku',
        'kickProtect': 'handler.KickProtect',
        'GhostEnd': 'handler.GhostEnd',
        'Purgemode': 'handler.Purgemode',
        'KickAllBan': 'handler.KickAllBan',
        'IsBlacklist': 'handler.IsBlacklist',
        'IsBlacklist2': 'handler.IsBlacklist2',
        'KickCancelBan129': 'handler.KickCancelBan129',
        'FastKick': 'handler.FastKick',
        'FastCancel': 'handler.FastCancel',
        'NodeBans': 'handler.NodeBans',
        'PurgeFaster': 'handler.PurgeFaster',
        'KickCancel': 'handler.KickCancel',
        'KickMemBan': 'handler.KickMemBan',
        'CanMemBan': 'handler.CanMemBan',
        'Ungban': 'handler.Ungban',
        'Addgban': 'handler.Addgban',
        'Joinsave': 'handler.Joinsave',
        'AddbanOp3': 'handler.AddbanOp3',
        'panicHandle': 'utils.PanicHandle',
        'GetMentionData': 'utils.GetMentionData',
        'randomString': 'utils.RandomString',
        'randomToString': 'handler.RandomToString',
        'IndexOf': 'utils.IndexOf',
        'Checkmulti': 'utils.Checkmulti'
    }

    for old, new in call_replacements.items():
        main_content = replace_calls(main_content, old, new)
        
    write_file(main_path, main_content)
    print("Done.")

if __name__ == '__main__':
    main()
