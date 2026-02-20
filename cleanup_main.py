
import os

file_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'

with open(file_path, 'r', encoding='utf-8') as f:
    lines = f.readlines()

new_lines = []
skip = False
for i, line in enumerate(lines):
    # Line numbers are 1-based in the editor, but 0-based in the list.
    # Line 57 is index 56.
    if i == 56: # Line 57
        if line.strip().startswith('var ('):
            skip = True
            print(f"Starting skip at line {i+1}: {line.strip()}")
    
    if skip:
        # Check if we reached func main()
        if line.strip().startswith('func main() {'):
            skip = False
            print(f"Ending skip at line {i+1}: {line.strip()}")
            new_lines.append(line)
        # If we are in the skip block, we don't append.
        # But we need to make sure we don't skip `func main()` itself.
    else:
        new_lines.append(line)

# Verify if we actually skipped something
if len(new_lines) < len(lines):
    print(f"Reduced lines from {len(lines)} to {len(new_lines)}")
    with open(file_path, 'w', encoding='utf-8') as f:
        f.writelines(new_lines)
    print("Successfully updated main.go")
else:
    print("No lines were skipped. Check the logic.")

