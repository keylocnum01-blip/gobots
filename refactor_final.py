import os
import re

def read_file(path):
    with open(path, 'r', encoding='utf-8') as f:
        return f.read()

def write_file(path, content):
    with open(path, 'w', encoding='utf-8') as f:
        f.write(content)

def extract_function(content, func_name):
    pattern = re.compile(r'func ' + func_name + r'\s*\(.*?\)\s*\{(?:[^{}]*|\{(?:[^{}]*|\{[^{}]*\})*\})*\}', re.DOTALL)
    match = pattern.search(content)
    if match:
        return match.group(0)
    return None

def remove_function(content, func_name):
    pattern = re.compile(r'func ' + func_name + r'\s*\(.*?\)\s*\{(?:[^{}]*|\{(?:[^{}]*|\{[^{}]*\})*\})*\}', re.DOTALL)
    return pattern.sub('', content)

def replace_calls(content, old_name, new_name):
    # Matches function calls: name(
    pattern = re.compile(r'\b' + old_name + r'\(')
    return pattern.sub(new_name + '(', content)

def modify_code(code, replacements):
    for old, new in replacements.items():
        code = code.replace(old, new)
    return code

def main():
    root = r'c:\Users\Home\Desktop\LineBotProtect\gobots'
    main_path = os.path.join(root, 'main.go')
    mod_path = os.path.join(root, 'handler', 'moderation.go')
    helpers_path = os.path.join(root, 'utils', 'helpers.go')

    main_content = read_file(main_path)
    mod_content = read_file(mod_path)
    helpers_content = read_file(helpers_path)

    # --- 1. Extract and Move to handler/moderation.go ---
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
        'handler.MemBan': 'MemBan', # Same package
        'utils.InArrayString': 'utils.InArrayString',
        'InArrayInt64': 'utils.InArrayInt64',
        'qrGo22': 'QrGo22',
        'GetKorban': 'botstate.GetKorban',
        'stringToInt': 'botstate.StringToInt',
        'valid.Abs': 'valid.Abs', # Assuming valid is imported or will be
        'func kickPelaku': 'func KickPelaku',
        'func kickProtect': 'func KickProtect',
        'kickPelaku(': 'KickPelaku(',
        'kickProtect(': 'KickProtect(',
        'FastKick(': 'FastKick(', # Already in handler
        'FastCancel(': 'FastCancel(' # Already in handler
    }
    
    # Add imports if missing
    if '"math"' not in mod_content and ('valid.Abs' in main_content):
         # Wait, valid.Abs might be from a library. Assuming it works if we keep imports.
         pass

    new_mod_funcs = ""
    for func, new_name in funcs_to_mod.items():
        code = extract_function(main_content, func)
        if code:
            # Apply replacements
            code = modify_code(code, replacements_mod)
            # Rename definition if needed
            if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
            
            # Check if already exists to avoid duplicates
            if f'func {new_name}' not in mod_content:
                new_mod_funcs += "\n\n" + code
                print(f"Extracted {func} -> {new_name}")
            else:
                print(f"Skipping {new_name}, already in moderation.go")

    if new_mod_funcs:
        mod_content += new_mod_funcs
        write_file(mod_path, mod_content)

    # --- 2. Extract and Move to utils/helpers.go ---
    funcs_to_help = {
        'randomToString': 'RandomToString'
    }
    
    replacements_help = {
        'stringToInt': 'botstate.StringToInt' # Wait, stringToInt is in botstate? 
        # utils/helpers.go might not import botstate.
        # If randomToString depends on stringToInt, and stringToInt is in botstate, 
        # utils/helpers.go needs to import botstate.
        # But botstate imports utils. Circular dependency!
        # Solution: Move stringToInt to utils/helpers.go as well?
        # Or keep randomToString in handler/helpers.go (which can import botstate).
    }

    # Let's check imports. botstate imports utils.
    # So utils cannot import botstate.
    # So randomToString CANNOT go to utils/helpers.go if it uses botstate.StringToInt.
    # Where is stringToInt defined?
    # main.go:202: StringToInt = []rune("01") in botstate/state.go?
    # I saw StringToInt in botstate/state.go line 202.
    # So randomToString must go to handler/helpers.go (or similar) or I move StringToInt to utils.
    
    # Decision: Move randomToString to handler/helpers.go instead of utils/helpers.go
    # because handler imports botstate.
    handler_helpers_path = os.path.join(root, 'handler', 'helpers.go')
    handler_helpers_content = read_file(handler_helpers_path)
    
    replacements_handler_help = {
        'stringToInt': 'botstate.StringToInt'
    }

    new_handler_help_funcs = ""
    code = extract_function(main_content, 'randomToString')
    if code:
        code = modify_code(code, replacements_handler_help)
        code = code.replace('func randomToString', 'func RandomToString', 1)
        if 'func RandomToString' not in handler_helpers_content:
             new_handler_help_funcs += "\n\n" + code
             print(f"Extracted randomToString -> RandomToString (handler/helpers.go)")

    if new_handler_help_funcs:
        handler_helpers_content += new_handler_help_funcs
        write_file(handler_helpers_path, handler_helpers_content)


    # --- 3. Remove from main.go and Update Calls ---
    
    # Functions to remove
    to_remove = [
        'KickCancelProtect', 'kickPelaku', 'kickProtect', 'GhostEnd',
        'Purgemode', 'KickAllBan', 'IsBlacklist', 'IsBlacklist2',
        'KickCancelBan129', 'FastKick', 'FastCancel', 'NodeBans', 'PurgeFaster', 'KickCancel',
        'KickMemBan', 'CanMemBan',
        'Ungban', 'Addgban', 'Joinsave', 'AddbanOp3',
        'panicHandle', 'GetMentionData', 'randomString', 'randomToString', 'IndexOf', 'Checkmulti'
    ]
    
    for func in to_remove:
        main_content = remove_function(main_content, func)
        print(f"Removed {func} from main.go")

    # Replace calls
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
        'randomToString': 'handler.RandomToString', # Moved to handler
        'IndexOf': 'utils.IndexOf',
        'Checkmulti': 'utils.Checkmulti'
    }

    for old, new in call_replacements.items():
        main_content = replace_calls(main_content, old, new)
    
    write_file(main_path, main_content)
    print("Main.go updated.")

if __name__ == '__main__':
    main()
