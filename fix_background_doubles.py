import os

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\background.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

# Fix double botstate prefixes
content = content.replace("botstate.botstate.", "botstate.")

# Fix potentially triple... just in case
content = content.replace("botstate.botstate.", "botstate.")

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)

print("Fixed double prefixes in handler/background.go")
