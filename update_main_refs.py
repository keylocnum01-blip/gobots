import re
import os

MAIN_GO = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"

REPLACEMENTS = [
    (r'CheckBan\(', 'handler.CheckBan('),
    (r'LeaveallGroups\(', 'handler.LeaveallGroups('),
    (r'checkunbanbots\(', 'handler.CheckUnbanBots('),
    (r'fmtDurations\(', 'handler.FmtDurations('),
    (r'kickBl\(', 'handler.KickBl('),
    (r'InfoCreator\(', 'handler.InfoCreator('),
    (r'AllBanList\(', 'handler.AllBanList('),
    (r'Cekbanwhois\(', 'handler.Cekbanwhois('),
    # Also fix some potentially missed renames if they were lowercase in main.go but Uppercase in handler
    (r'handler\.checkunbanbots\(', 'handler.CheckUnbanBots('),
    (r'handler\.kickBl\(', 'handler.KickBl('),
    (r'handler\.fmtDurations\(', 'handler.FmtDurations('),
]

def read_file(path):
    with open(path, "r", encoding="utf-8") as f:
        return f.read()

def write_file(path, content):
    with open(path, "w", encoding="utf-8") as f:
        f.write(content)

def main():
    content = read_file(MAIN_GO)
    
    for old, new in REPLACEMENTS:
        # We need to be careful not to double prefix if already prefixed
        # But my regex is simple.
        # If I replace 'CheckBan(' with 'handler.CheckBan(', what if it was 'handler.CheckBan('?
        # It becomes 'handler.handler.CheckBan('.
        # So I should use regex that ensures no 'handler.' before.
        
        # Regex: (?<!handler\.)CheckBan\(
        pattern = r'(?<!handler\.)' + old
        
        # Also need to handle if it was defined as function (func CheckBan) which I already removed.
        # But if I missed removal, I don't want to change definition signature.
        # But I removed definitions.
        
        content = re.sub(pattern, new, content)
        
    write_file(MAIN_GO, content)
    print("Updated references in main.go")

if __name__ == "__main__":
    main()
