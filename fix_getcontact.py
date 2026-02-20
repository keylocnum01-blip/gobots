import re

file_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\handler\info.go'

with open(file_path, 'r', encoding='utf-8') as f:
    content = f.read()

# Pattern explanation:
# Capture client var name: (\w+)
# Capture target var name: (\w+)
# Match the check for error (new != nil)
# Capture the "Closed Account" block content: (.*?)
# Match the else block start
# Match the redundant GetContact call: \w+, _ := \1\.GetContact\(\2\)
# Capture the success block content: (.*?)
# End of else block (assuming simple indentation or single level brace, but regex .*? is non-greedy)
# We need to be careful about matching braces.

# Simplified approach:
# We look for the specific structure:
# new := client.Getcontactuser(xx)
# if new != nil {
#     ...
# } else {
#     x, _ := client.GetContact(xx)
#     ...
# }

# We can replace:
# new := (\w+)\.Getcontactuser\((\w+)\)
# if new != nil \{
# with:
# x, err := \1.GetContact(\2)
# if err != nil {

# And remove:
# \w+, _ := \1\.GetContact\(\2\)
# inside the else block.

# Step 1: Replace the setup and check
# Note: 'new' variable name might be used elsewhere, but here it's locally scoped.
# We replace `new := client.Getcontactuser(xx)` with `x, err := client.GetContact(xx)`
# And `if new != nil` with `if err != nil`

# Regex 1:
# new := (\w+)\.Getcontactuser\((\w+)\)\s+if new != nil
# Replace with:
# x, err := \1.GetContact(\2)\n\tif err != nil

pattern1 = r'new\s*:=\s*(\w+)\.Getcontactuser\((\w+)\)\s*if\s*new\s*!=\s*nil'
replacement1 = r'x, err := \1.GetContact(\2)\n\t\t\tif err != nil' 
# Indentation might be tricky. The original code has indentation.
# Let's try to preserve whitespace.

pattern1_flexible = r'(\s*)new\s*:=\s*(\w+)\.Getcontactuser\((\w+)\s*\)\s*(\s*)if\s*new\s*!=\s*nil'
# Replacement: \1x, err := \2.GetContact(\3)\4if err != nil

# Step 2: Remove the redundant GetContact in else block
# x, _ := client.GetContact(xx)
# We can just remove this line if it exists inside the else block.
# Or better, we can replace it with nothing, but we need to make sure we don't break logic if 'x' was not defined before.
# In Step 1, we defined 'x' in the outer scope (before if).
# So `x, _ := ...` redeclaring x might be an issue or shadowing.
# Wait, `x, err := ...` declares x.
# Inside else, `x, _ := ...` declares a NEW x (shadowing).
# If we remove `x, _ := ...`, the inner code will use the outer `x`.
# This is exactly what we want!

pattern2 = r'(\s*)x,\s*_\s*:=\s*(\w+)\.GetContact\((\w+)\s*\)'
# We should only remove this if it follows the pattern we just changed.
# But simply removing it globally might be safe if we are sure 'x' is available.
# In the original code:
# else {
#    x, _ := client.GetContact(xx)
#    ...
# }
# 'x' is local to else block.
# With change:
# x, err := ...
# if err != nil { ... } else {
#    (removed x definition)
#    ... use x ...
# }
# 'x' is now in scope of else block (defined before if).

# So the plan:
# 1. Regex replace `new := client.Getcontactuser(xx)\n\s*if new != nil` -> `x, err := client.GetContact(xx)\n\s*if err != nil`
# 2. Regex remove `x, _ := client.GetContact(xx)` (and the newline)

def replacer1(match):
    indent = match.group(1)
    client = match.group(2)
    arg = match.group(3)
    indent2 = match.group(4) # whitespace between statement and if
    # Use 'x' and 'err' as variable names. 
    # Caution: if 'x' is already used in outer scope, this shadows it. 
    # But in the loops shown, 'x' seems to be the standard name for contact.
    return f'{indent}x, err := {client}.GetContact({arg}){indent2}if err != nil'

content_new = re.sub(pattern1_flexible, replacer1, content)

# Remove redundant GetContact calls
# Matches: whitespace + x, _ := client.GetContact(arg)
pattern2 = r'(\s*)x,\s*_\s*:=\s*\w+\.GetContact\(\w+\)'
content_new = re.sub(pattern2, '', content_new)

if content != content_new:
    print("Replaced occurrences.")
    with open(file_path, 'w', encoding='utf-8') as f:
        f.write(content_new)
else:
    print("No changes made.")

