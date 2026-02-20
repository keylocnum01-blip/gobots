import os
import re

def read_file(path):
    with open(path, 'r', encoding='utf-8') as f:
        return f.read()

def write_file(path, content):
    with open(path, 'w', encoding='utf-8') as f:
        f.write(content)

def main():
    root = r'c:\Users\Home\Desktop\LineBotProtect\gobots'
    main_path = os.path.join(root, 'main.go')
    main_content = read_file(main_path)
    
    # 1. Add config import
    if '"./config"' not in main_content:
        # Try to find imports block
        if 'import (' in main_content:
            main_content = main_content.replace('import (', 'import (\n\t"./config"', 1)
            print("Added config import")
        else:
            print("Could not find import block to add config import")
    
    # 2. Replace kickop struct initialization
    main_content = main_content.replace('&kickop{', '&config.Kickop{')
    
    # 3. Replace method calls
    # Map lowercase methods to PascalCase
    methods = {
        'cek': 'Cek',
        'ceko': 'Ceko',
        'del': 'Del',
        'ceki': 'Ceki',
        'deli': 'Deli',
        'clear': 'Clear'
    }
    
    for old, new in methods.items():
        # Replace filterWar.method(
        pattern = r'filterWar\.' + re.escape(old) + r'\('
        replacement = 'filterWar.' + new + '('
        main_content = re.sub(pattern, replacement, main_content)
        print(f"Replaced filterWar.{old} with filterWar.{new}")
        
        # Also check if there are other variables of type kickop, but filterWar is the main one found
    
    write_file(main_path, main_content)
    print("Finished updating main.go")

if __name__ == '__main__':
    main()
