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
        # Try method syntax if simple func fails (though for this script we mostly target simple funcs)
        regex_method = r'func\s+\([^)]+\)\s+' + re.escape(func_name) + r'\s*\('
        match = re.search(regex_method, content)
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
    back_path = os.path.join(root, 'handler', 'background.go')
    logic_path = os.path.join(root, 'handler', 'logic.go')

    main_content = read_file(main_path)
    mod_content = read_file(mod_path)
    help_content = read_file(help_path)
    back_content = read_file(back_path)
    logic_content = read_file(logic_path)

    # 1. Moderation
    funcs_mod = {
        'BanAll': 'BanAll',
        'Checklistexpel': 'Checklistexpel'
    }
    repl_mod = {
        'Banned': 'botstate.Banned',
        'MemUser': 'MemUser', # In helpers, but moderation imports helpers? No, same package handler
        'DEVELOPER': 'botstate.DEVELOPER',
        'UserBot': 'botstate.UserBot',
        'SendMycreator': 'SendMycreator',
        'SendMymaker': 'SendMymaker',
        'SendMyseller': 'SendMyseller',
        'SendMybuyer': 'SendMybuyer',
        'SendMyowner': 'SendMyowner',
        'SendMymaster': 'SendMymaster',
        'SendMyadmin': 'SendMyadmin',
        'SendMygowner': 'SendMygowner',
        'handler.LogAccess': 'LogAccess', # It is in logging.go (handler package)
        'fancy(': 'botstate.Fancy(',
        'utils.InArrayString': 'utils.InArrayString',
        'utils.RemoveString': 'utils.RemoveString',
        'defer linetcr.PanicOnly()': 'defer utils.PanicHandle("Checklistexpel")',
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
        'botDuration': 'BotDuration',
        'AppendLastD': 'AppendLastD',
        'AppendLast': 'AppendLast'
    }
    repl_help = {
        'defer linetcr.PanicOnly()': 'defer utils.PanicHandle("Helper")',
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


    # 3. Background
    funcs_back = {
        'LogGet': 'LogGet',
        'savejoin': 'SaveJoin'
    }
    repl_back = {
        'Lastinvite': 'botstate.Lastinvite',
        'Lastkick': 'botstate.Lastkick',
        'Lastjoin': 'botstate.Lastjoin',
        'Lastcancel': 'botstate.Lastcancel',
        'Lastupdate': 'botstate.Lastupdate',
        'Lastleave': 'botstate.Lastleave',
        'Lasttag': 'botstate.Lasttag',
        'Lastcon': 'botstate.Lastcon',
        'Laststicker': 'botstate.Laststicker',
        'Lastmid': 'botstate.Lastmid',
        'Lastmessage': 'botstate.Lastmessage',
        'Squadlist': 'botstate.Squadlist',
        'Detectjoin': 'botstate.Detectjoin',
        'Banned': 'botstate.Banned',
        'defer linetcr.PanicOnly()': 'defer utils.PanicHandle("LogGet")',
        'AppendLast': 'AppendLast', # in handler package
        'AppendLastD': 'AppendLastD', # in handler package
        'AppendLastSticker': 'AppendLastSticker', # need to check if exists
        'MentionList': 'MentionList', # need to check
        'MemUser': 'MemUser',
    }
    
    extracted_back = []
    for func, new_name in funcs_back.items():
        code, main_content = extract_and_remove(main_content, func)
        if code:
             for old, new in repl_back.items():
                 code = code.replace(old, new)
             if func != new_name:
                code = code.replace(f'func {func}', f'func {new_name}', 1)
             
             if f'func {new_name}' not in back_content:
                 extracted_back.append(code)
                 print(f"Extracted {func} -> {new_name} (background)")

    if extracted_back:
        back_content += "\n\n" + "\n\n".join(extracted_back)
        write_file(back_path, back_content)


    # 4. Logic
    funcs_logic = {
        'cekOpinvite': 'CekOpinvite',
        'LlistCheck': 'LlistCheck'
    }
    repl_logic = {
        'oplistinvite': 'botstate.Oplistinvite',
        'Squadlist': 'botstate.Squadlist',
        'Lastmid': 'botstate.Lastmid',
        'Lastmessage': 'botstate.Lastmessage',
        'Lastinvite': 'botstate.Lastinvite',
        'Lastkick': 'botstate.Lastkick',
        'Lastcancel': 'botstate.Lastcancel',
        'Lastupdate': 'botstate.Lastupdate',
        'Lastjoin': 'botstate.Lastjoin',
        'Lasttag': 'botstate.Lasttag',
        'Lastcon': 'botstate.Lastcon',
        'Lastleave': 'botstate.Lastleave',
        'Banned': 'botstate.Banned',
        'CancelPend': 'botstate.CancelPend',
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


    # 5. Remove kickop methods
    kickop_methods = ['ceko', 'cek', 'del', 'ceki', 'deli', 'clear']
    for method in kickop_methods:
        regex = r'func\s+\(self\s+\*kickop\)\s+' + re.escape(method) + r'\s*\('
        match = re.search(regex, main_content)
        if match:
             code, main_content = extract_and_remove(main_content, method) # This might fail if extract_and_remove doesn't handle methods well with name
             # But extract_and_remove uses regex for simple func. I need to use the method regex.
             # Actually, let's just use regex sub for these since they are specific.
             pass
    
    # Custom removal for kickop methods
    for method in kickop_methods:
         regex = r'func\s+\(self\s+\*kickop\)\s+' + re.escape(method) + r'\s*\([^\)]*\)\s*[^{]*'
         match = re.search(regex, main_content)
         if match:
             start = match.start()
             end = find_func_end(main_content, start)
             if end != -1:
                 main_content = main_content[:start] + main_content[end:]
                 print(f"Removed kickop method: {method}")

    # 6. Update calls in main.go
    calls = {
        'BanAll': 'handler.BanAll',
        'Checklistexpel': 'handler.Checklistexpel',
        'botDuration': 'handler.BotDuration',
        'AppendLastD': 'handler.AppendLastD',
        'AppendLast': 'handler.AppendLast',
        'LogGet': 'handler.LogGet',
        'savejoin': 'handler.SaveJoin',
        'cekOpinvite': 'handler.CekOpinvite',
        'LlistCheck': 'handler.LlistCheck'
    }
    
    for old, new in calls.items():
        main_content = replace_calls(main_content, old, new)

    write_file(main_path, main_content)
    print("Done.")

if __name__ == '__main__':
    main()
