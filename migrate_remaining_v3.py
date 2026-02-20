import re
import os

MAIN_GO = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"
HANDLER_HELPERS = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\helpers.go"
HANDLER_SYSTEM = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\system.go"
HANDLER_MODERATION = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\moderation.go"

MIGRATIONS = [
    {
        "name": "CheckBan",
        "target": HANDLER_HELPERS
    },
    {
        "name": "LeaveallGroups",
        "target": HANDLER_SYSTEM
    },
    {
        "name": "checkunbanbots",
        "target": HANDLER_MODERATION
    },
    {
        "name": "fmtDurations",
        "target": HANDLER_HELPERS
    },
    {
        "name": "kickBl",
        "target": HANDLER_MODERATION
    }
]

REMOVALS_ONLY = [
    "InfoCreator",
    "AllBanList",
    "Cekbanwhois",
    "CheckExprd",
    "CekDuedate",
    "Setkickto",
    "AutojoinQr22",
    "AutojoinQr",
    "qrGo",
    "hstg",
    "KIckbansPurges",
    "KIckbansPurges1",
    "JKickFuck",
    "JCancelFuck",
    "KickCansWar",
    "AcceptWar"
]

def read_file(path):
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def write_file(path, content):
    with open(path, "w", encoding="utf-8") as f:
        f.write(content)

def extract_function(content, func_name):
    pattern = r'func\s+' + re.escape(func_name) + r'\s*\('
    match = re.search(pattern, content)
    if not match:
        return None, content
    
    start_idx = match.start()
    brace_search = re.search(r'\{', content[start_idx:])
    if not brace_search:
        return None, content
        
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
        func_body = content[start_idx:end_idx]
        new_content = content[:start_idx] + content[end_idx:]
        return func_body, new_content
    return None, content

def clean_body(body):
    # Remove 'handler.' prefix as functions are now in the same package
    body = body.replace("handler.", "")
    # Remove 'botstate.' prefix if we are in botstate package? No, we are in handler.
    # We should keep 'botstate.' prefix.
    # Check imports. 'handler' files import 'botstate' and 'linetcr'.
    # Ensure TitleCase for exported functions if needed.
    # But checkunbanbots is lowercase. If used across files in handler, it's fine.
    # If used from main.go, main.go will need to call handler.checkunbanbots (exported? No, must be Checkunbanbots).
    # If checkunbanbots is lowercase, it is unexported. If main.go uses it, it will break.
    # But main.go is where it came from.
    # If main.go calls it, it must be exported.
    # But main.go is the root package 'main'. 'handler' is package 'handler'.
    # So if I move it to handler, I MUST export it (Capitalize) if I want main.go to call it?
    # Or if main.go no longer needs it?
    # 'checkunbanbots' seems to be a command handler helper.
    # If I rename it to CheckUnbanBots, I must update calls.
    # For now, let's keep name, but if it's unexported, main.go can't call it.
    # Wait, if main.go calls it, I need to update main.go to use handler.Checkunbanbots.
    # But I am removing it from main.go.
    # Does main.go call it?
    # I should rename it to CheckUnbanBots to be safe and exported.
    
    if body.startswith("func checkunbanbots"):
        body = body.replace("func checkunbanbots", "func CheckUnbanBots", 1)
    if body.startswith("func kickBl"):
        body = body.replace("func kickBl", "func KickBl", 1)
    if body.startswith("func fmtDurations"):
        body = body.replace("func fmtDurations", "func FmtDurations", 1)
        
    return body

def main():
    main_content = read_file(MAIN_GO)
    
    # Handle Migrations
    for item in MIGRATIONS:
        func_name = item["name"]
        target_path = item["target"]
        
        print(f"Migrating {func_name} to {target_path}...")
        func_body, main_content = extract_function(main_content, func_name)
        
        if func_body:
            target_content = read_file(target_path)
            
            # Check if already exists in target
            check_name = func_name
            if func_name == "checkunbanbots": check_name = "CheckUnbanBots"
            if func_name == "kickBl": check_name = "KickBl"
            if func_name == "fmtDurations": check_name = "FmtDurations"
            
            if f"func {check_name}" in target_content:
                print(f"  Function {check_name} already exists in target. Skipping append.")
            else:
                cleaned_body = clean_body(func_body)
                target_content += "\n\n" + cleaned_body
                write_file(target_path, target_content)
                print(f"  Appended {check_name} to {target_path}")
        else:
            print(f"  Function {func_name} not found in main.go")

    # Handle Removals
    for func_name in REMOVALS_ONLY:
        print(f"Removing {func_name} from main.go...")
        func_body, main_content = extract_function(main_content, func_name)
        if func_body:
            print(f"  Removed {func_name}")
        else:
            print(f"  {func_name} not found in main.go")

    # Write back main.go
    write_file(MAIN_GO, main_content)
    print("Done.")

if __name__ == "__main__":
    main()
