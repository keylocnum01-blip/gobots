import os

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\main.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

# Fix handler.Autoset -> handler.AutoSet
content = content.replace("handler.Autoset()", "handler.AutoSet()")

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)

print("Fixed handler calls in main.go")
