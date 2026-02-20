import os

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\helpers.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

new_func = """

func CheckPermission(num int, sinder string, group string) bool {
	if num == 0 {
		if utils.InArrayString(botstate.DEVELOPER, sinder) {
			return true
		}
	} else if num == 1 {
		if SendMycreator(sinder) {
			return true
		}
	} else if num == 2 {
		if SendMymaker(sinder) {
			return true
		}
	} else if num == 3 {
		if SendMyseller(sinder) {
			return true
		}
	} else if num == 4 {
		if SendMybuyer(sinder) {
			return true
		}
	} else if num == 5 {
		if SendMyowner(sinder) {
			return true
		}
	} else if num == 6 {
		if SendMymaster(sinder) {
			return true
		}
		return false
	} else if num == 7 {
		if SendMyadmin(sinder) {
			return true
		}
	} else if num == 8 {
		if SendMygowner(group, sinder) {
			return true
		}
	} else if num == 9 {
		if SendMygadmin(group, sinder) {
			return true
		}
	}
	return false
}
"""

if "func CheckPermission" not in content:
    with open(file_path, "a", encoding="utf-8") as f:
        f.write(new_func)
    print("Appended CheckPermission to handler/helpers.go")
else:
    print("CheckPermission already exists in handler/helpers.go")
