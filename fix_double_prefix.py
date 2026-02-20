import re

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

# Replace handler.handler. with handler.
content = content.replace("handler.handler.", "handler.")

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)
