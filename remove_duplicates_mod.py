import os

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\moderation.go"
old_str = """func AccKickBan(client *linetcr.Account, to string) {
	go func() {
		AccGroup(client, to);JoinLlinetcrBan(client, to)
	}()
}

func AccGroup(client *linetcr.Account, to string) {
	go func() {client.NewacceptGroup(to)
	}()
}

func JoinLlinetcrBan(client *linetcr.Account, to string) {
	_, mem, pend := client.GetChatList(to)
	for _, k := range mem {
		if IsBlacklist(client, k) == true {
			go func(k string) {go client.DeleteOtherFromChats(to, []string{k})}(k)
		}
	}
	for _, i := range pend {
		if IsBlacklist(client, i) == true {
			go func(i string) {go client.CancelChatInvitations(to, []string{i})}(i)
		}
	}
	runtime.GOMAXPROCS(100)
}"""

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

if old_str in content:
    new_content = content.replace(old_str, "")
    with open(file_path, "w", encoding="utf-8") as f:
        f.write(new_content)
    print("Successfully removed duplicates from handler/moderation.go")
else:
    print("Could not find the duplicate block in handler/moderation.go")
