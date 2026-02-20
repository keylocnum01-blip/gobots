import os
import re

def read_file(path):
    with open(path, 'r', encoding='utf-8') as f:
        return f.read()

def write_file(path, content):
    with open(path, 'w', encoding='utf-8') as f:
        f.write(content)

def remove_function(content, func_name):
    # Matches func FuncName(...) { ... } handling nested braces
    pattern = re.compile(r'func ' + func_name + r'\s*\(.*?\)\s*\{(?:[^{}]*|\{(?:[^{}]*|\{[^{}]*\})*\})*\}', re.DOTALL)
    return pattern.sub('', content)

def replace_calls(content, old_name, new_name):
    # Simple replace for function calls
    # Be careful not to replace substrings of other words
    pattern = re.compile(r'\b' + old_name + r'\(')
    return pattern.sub(new_name + '(', content)

def main():
    main_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'
    actions_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\actions.go'

    # --- 1. Process handler/actions.go ---
    print(f"Processing {actions_path}...")
    actions_content = read_file(actions_path)

    # Remove duplicates that are in moderation.go
    funcs_to_remove_actions = ['AccKickBan', 'AccGroup', 'JoinLlinetcrBan']
    for func in funcs_to_remove_actions:
        actions_content = remove_function(actions_content, func)
        print(f"Removed {func} from actions.go")

    # Rename local functions to exported
    # kickProtect -> KickProtect
    if 'func kickProtect' in actions_content:
        actions_content = actions_content.replace('func kickProtect', 'func KickProtect')
        actions_content = replace_calls(actions_content, 'kickProtect', 'KickProtect')
        print("Renamed kickProtect to KickProtect in actions.go")

    # kickPelaku -> KickPelaku
    if 'func kickPelaku' in actions_content:
        actions_content = actions_content.replace('func kickPelaku', 'func KickPelaku')
        actions_content = replace_calls(actions_content, 'kickPelaku', 'KickPelaku')
        print("Renamed kickPelaku to KickPelaku in actions.go")

    # Fix qrGo22 -> QrGo22 call (from invite.go)
    actions_content = replace_calls(actions_content, 'qrGo22', 'QrGo22')
    print("Updated qrGo22 calls to QrGo22 in actions.go")

    write_file(actions_path, actions_content)

    # --- 2. Process main.go ---
    print(f"Processing {main_path}...")
    main_content = read_file(main_path)

    # Remove functions that are now in actions.go
    funcs_to_remove_main = ['KickCancelProtect', 'kickPelaku', 'kickProtect', 'GhostEnd']
    for func in funcs_to_remove_main:
        main_content = remove_function(main_content, func)
        print(f"Removed {func} from main.go")

    # Remove duplicates explicitly requested
    duplicates = [
        'Setkickto', 'AutojoinQr22', 'AutojoinQr', 'qrGo', 'hstg',
        'KIckbansPurges', 'KIckbansPurges1', 'JKickFuck', 'JCancelFuck',
        'KickCansWar', 'AcceptWar'
    ]
    for func in duplicates:
        main_content = remove_function(main_content, func)
        print(f"Removed duplicate {func} from main.go")

    # Update calls to handler functions
    # kickProtect -> handler.KickProtect
    main_content = replace_calls(main_content, 'kickProtect', 'handler.KickProtect')
    # kickPelaku -> handler.KickPelaku
    main_content = replace_calls(main_content, 'kickPelaku', 'handler.KickPelaku')
    # GhostEnd -> handler.GhostEnd
    main_content = replace_calls(main_content, 'GhostEnd', 'handler.GhostEnd')
    # KickCancelProtect -> handler.KickCancelProtect
    main_content = replace_calls(main_content, 'KickCancelProtect', 'handler.KickCancelProtect')

    write_file(main_path, main_content)
    print("Done.")

if __name__ == '__main__':
    main()
