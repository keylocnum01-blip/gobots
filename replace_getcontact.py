import re
import os

file_path = r'c:\Users\Home\Desktop\LineBotProtect\gobots\main.go'

with open(file_path, 'r', encoding='utf-8') as f:
    lines = f.readlines()

new_lines = []
i = 0
replaced_count = 0

while i < len(lines):
    line = lines[i]
    # Check for start of pattern: new := client.Getcontactuser(xx)
    match_start = re.search(r'(\s*)new\s*:=\s*(\w+)\.Getcontactuser\(([^)]+)\)', line)
    
    if match_start:
        indent = match_start.group(1)
        client_var = match_start.group(2)
        id_var = match_start.group(3)
        
        # Check if following lines match the structure
        # We need to look ahead a few lines
        # Expecting:
        # if new != nil {
        #    ...
        # } else {
        #    x, _ := client.GetContact(...)
        #    list += ... x.DisplayName ...
        # }
        
        # Simplistic lookahead
        is_match = False
        append_line = ""
        block_end_index = -1
        
        try:
            # Check for 'if new != nil {'
            if 'if new != nil' in lines[i+1]:
                # Look for 'else {'
                for j in range(i+2, i+20): # Limit search range
                    if 'else {' in lines[j]:
                        # Look for GetContact in next line
                        if 'GetContact' in lines[j+1]:
                            # Look for append line in next line
                            if '+=' in lines[j+2] and 'DisplayName' in lines[j+2]:
                                append_line = lines[j+2]
                                # Check for closing brace
                                if '}' in lines[j+3]:
                                    block_end_index = j+3
                                    is_match = True
                                break
                
        except IndexError:
            pass
            
        if is_match:
            # Construct replacement
            # Replace x.DisplayName with name
            # And assume x is the variable name used for contact
            # usually x, _ := ...
            
            # Extract the variable name used for GetContact result from line j+1
            # x, _ := ...
            # block_end_index is absolute index of '}'
            # block_end_index - 2 is absolute index of 'x, _ := ...' (since block_end_index - 1 is append line)
            # wait, j+1 was 'GetContact', j+2 was append, j+3 was '}'
            # so block_end_index = j+3
            # block_end_index - 2 = j+1 ('GetContact' line)
            
            contact_var_match = re.search(r'(\w+),\s*_\s*:=', lines[block_end_index - 2]) 
            contact_var = "x"
            if contact_var_match:
                contact_var = contact_var_match.group(1)
            
            new_append_line = append_line.replace(f'{contact_var}.DisplayName', 'name')
            
            new_lines.append(f'{indent}name := handler.GetContactName({client_var}, {id_var})\n')
            new_lines.append(new_append_line)
            
            i = block_end_index + 1
            replaced_count += 1
            continue

    new_lines.append(line)
    i += 1

print(f"Replaced {replaced_count} occurrences.")

with open(file_path, 'w', encoding='utf-8') as f:
    f.writelines(new_lines)
