import re

def extract_and_append():
    with open('main.go', 'r', encoding='utf-8') as f:
        main_content = f.read()

    vars_to_extract = [
        "helpowner", "helpmaster", "helpadmin", "helpgowner", "helpgadmin", "details"
    ]

    extracted_code = ""

    for var in vars_to_extract:
        # Regex to capture "var = Type{...}" or "var = map[...]..."
        # Handling multiline.
        # Pattern: varname = ... (matching balanced braces is hard with regex, but indentation helps)
        # In main.go they are inside var ( ... ), so they look like:
        # \thelpowner = []string{
        # \t\t...
        # \t}
        
        pattern = r'(\t' + var + r' = .*?^\t})'
        match = re.search(pattern, main_content, re.MULTILINE | re.DOTALL)
        if match:
            code = match.group(1)
            # Capitalize the variable name
            # code starts with "\thelpowner"
            # We want "\tHelpowner"
            
            # Find the first occurrence of var and capitalize it
            new_code = code.replace(f"\t{var}", f"\t{var.capitalize()}", 1)
            extracted_code += new_code + "\n"
        else:
            print(f"Could not find {var}")

    if not extracted_code:
        print("Nothing extracted")
        return

    with open('botstate/state.go', 'r', encoding='utf-8') as f:
        state_content = f.read()

    # Find the last ')'
    last_paren_index = state_content.rfind(')')
    if last_paren_index == -1:
        print("Could not find closing parenthesis in state.go")
        return

    new_state_content = state_content[:last_paren_index] + extracted_code + state_content[last_paren_index:]

    with open('botstate/state.go', 'w', encoding='utf-8') as f:
        f.write(new_state_content)
    
    print("Appended extracted variables to botstate/state.go")

if __name__ == "__main__":
    extract_and_append()
