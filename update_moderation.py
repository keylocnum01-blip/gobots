
import os

file_path = r"c:\Users\Home\Desktop\LineBotProtect\gobots\handler\moderation.go"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

# Add import if not present
if 'valid "github.com/asaskevich/govalidator"' not in content:
    content = content.replace('	"../utils"', '	"../utils"\n	valid "github.com/asaskevich/govalidator"')

# Append functions
new_functions = r'''
func CancelEnemy(client *linetcr.Account, mem []string, to string) {
	defer utils.PanicHandle("CancelEnemy")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	cans := Room.HaveClient
	no := 0
	ah := 0
	if len(mem) > 50 {
		mem = mem[:50]
	}
	for _, target := range mem {
		go func(target string, no int) {
			cans[no].CancelChatInvitations(to, []string{target})
		}(target, no)
		if ah >= botstate.MaxCancel {
			no++
			if no >= len(cans) {
				no = 0
			}
			ah = 0
		}
		ah++
	}
}

func CancelBanInv(client *linetcr.Account, mem []string, to string) {
	defer utils.PanicHandle("cancelBanInv")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	Cans := Room.HaveClient
	clen := len(Cans)
	if clen != 0 {no := 0;ah := 0;if len(mem) > 50 {mem = mem[:50]};for _, target := range mem {var wg sync.WaitGroup;wg.Add(len(mem));go func(target string) {defer wg.Done();Cans[no].CancelChatInvitations(to, []string{target})}(target);if ah >= botstate.MaxCancel {no++;if no >= clen {no = 0};ah = 0};ah++}
	}
}

func Cancelall(client *linetcr.Account, mem []string, to string) {
	defer utils.PanicHandle("cancelall")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	cans := Room.HaveClient
	no := 0
	ah := 0
	if len(mem) > 50 {
		mem = mem[:50]
	}
	for _, target := range mem {
		go func(target string, no int) {
			cans[no].CancelChatInvitations(to, []string{target})
		}(target, no)
		if ah >= botstate.MaxCancel {
			no++
			if no >= len(cans) {
				no = 0
			}
			ah = 0
		}
		ah++
	}
}

func GetfuckV1(client *linetcr.Account, vo string, to string) {
	defer utils.PanicHandle("getfuck")
	runtime.GOMAXPROCS(botstate.Cpu)
	for _, v := range botstate.Banned.Banlist {
		if botstate.Banned.GetBan(v) {
			if linetcr.IsPending(client, to, v) == true {
				go func(v string) { 
					client.CancelChatInvitations(to, []string{v})
				}(v)
			}
			if linetcr.IsMembers(client, to, v) == true {
				if botstate.Banned.GetBan(v) {
					go func(v string) { 
						client.DeleteOtherFromChats(to, []string{v}) 
					}(v)
				}
			}
		}
	}
}

func Getfuck(cl *linetcr.Account, vo string, Group string) {
	defer utils.PanicHandle("getfuck")
	runtime.GOMAXPROCS(botstate.Cpu)
	if MemBan(Group, vo) {
		cl.CancelChatInvitations(Group, []string{vo})
	}
}


func Cancelallcek(client *linetcr.Account, mem []string, to string) {
	defer utils.PanicHandle("cancelallcek")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	Cans := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := botstate.GetKorban(mid)
			if cl.Limited == false {
				Cans = append(Cans, cl)
			}
			oke = append(oke, mid)
		}
	}
	if len(Cans) != 0 {
		sort.Slice(Cans, func(i, j int) bool {
			return Cans[i].KickPoint < Cans[j].KickPoint
		})
		Room.HaveClient = Cans
		no := 0
		ah := 0
		for _, target := range mem {
			go Getfuck(Cans[no], target, to)
			if ah >= botstate.MaxCancel {
				no++
				if no >= len(Cans) {
					no = 0
				}
				ah = 0
			}
			ah++
		}
	}
}

func Purgeact(Group string, cl *linetcr.Account) {
	defer utils.PanicHandle("purgeact")
	mem := make(chan []string)
	go func(m chan []string) {
		memlistss := []string{}
		_, memlists := cl.GetGroupMember(Group)
		for target, _ := range memlists {
			if MemBan(Group, target) {
				memlistss = append(memlistss, target)
			}
		}
		m <- memlistss
	}(mem)
	Cans := []*linetcr.Account{}
	for _, ym := range linetcr.Actor(Group) {
		if ym.Limited {
			Cans = append(Cans, ym)
		}
	}
	ClAct := len(Cans)
	if ClAct != 0 {
		no := 0
		memlist := <-mem
		for _, target := range memlist {
			if no >= ClAct {
				no = 0
			}
			cl = Cans[no]
			go cl.DeleteOtherFromChat(Group, target)
			no += 1
		}
	} else if cl.Limited {
		memlist := <-mem
		for _, target := range memlist {
			go cl.DeleteOtherFromChat(Group, target)
		}
	}
}
func GroupBackupKick(client *linetcr.Account, to, pelaku string, cek bool) {
	defer utils.PanicHandle("groupBackupKick")
	Room := linetcr.GetRoom(to)
	memlist, pending := client.GetChatListMap(to)
	ban := []string{}
	ban2 := []string{}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := botstate.GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)}} else if MemBan(to, mid) {ban = append(ban, mid)}
	}
	for mid2, _ := range pending {if MemBan(to, mid2) {ban2 = append(ban2, mid2)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if botstate.ProtectMode {
			chat := client.GetChat([]string{to}, true, true)
			if chat == nil { return }
			memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
			var createdTime int64
			for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
			}
			for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(to, mid) {botstate.Banned.AddBan(mid);ban = append(ban, mid)}}
			}
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban) != 0 {
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban2) != 0 {
			no := 0
			ah := 0
			for _, target2 := range ban2 {go func(target2 string, no int) {exe[no].CancelChatInvitations(to, []string{target2})}(target2, no);if ah >= botstate.MaxCancel {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
	//} else {
		//no := 0
		//ah := 0
		//if _, ok := memlist[pelaku]; ok {
			//exe[0].DeleteOtherFromChats(to, []string{pelaku});if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		//}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}
func GroupBackupKickV2(client *linetcr.Account, to, pelaku string, cek bool) {
	defer utils.PanicHandle("groupBackupKickV2")
	Room := linetcr.GetRoom(to)
	memlist, pending := client.GetChatListMap(to)
	ban := []string{}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := botstate.GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)}} else if MemBan(to, mid) {ban = append(ban, mid)}
	}
	for mid2, _ := range pending {if MemBan(to, mid2) {ban = append(ban, mid2)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if botstate.ProtectMode {
			chat := client.GetChat([]string{to}, true, true)
			if chat == nil { return }
			memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
			var createdTime int64
			for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
			}
			for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(to, mid) {botstate.Banned.AddBan(mid);ban = append(ban, mid)}}
			}
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban) != 0 {
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
			for _, target2 := range ban {go func(target2 string, no int) {exe[no].CancelChatInvitations(to, []string{target2})}(target2, no);if ah >= botstate.MaxCancel {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
	} else {
		no := 0
		ah := 0
		if _, ok := memlist[pelaku]; ok {
			exe[0].DeleteOtherFromChats(to, []string{pelaku});if ah >= botstate.MaxCancel {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}
func GroupBackupCans(client *linetcr.Account, to, pelaku string, cek bool) {
	defer utils.PanicHandle("groupBackupCans")
	Room := linetcr.GetRoom(to)
	memlist, _ := client.GetChatListMap(to)
	ban := []string{}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := botstate.GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)}} else if MemBan(to, mid) {ban = append(ban, mid)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if botstate.ProtectMode {
			chat := client.GetChat([]string{to}, true, true)
			if chat == nil { return }
			memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
			var createdTime int64
			for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
			}
			for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(to, mid) {botstate.Banned.AddBan(mid);ban = append(ban, mid)}}
			}
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
		if len(ban) != 0 {
			no := 0
			ah := 0
			for _, target := range ban {go func(target string, no int) {exe[no].DeleteOtherFromChats(to, []string{target})}(target, no);if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
			}
		}
	} else {
		no := 0
		ah := 0
		if _, ok := memlist[pelaku]; ok {
			exe[0].DeleteOtherFromChats(to, []string{pelaku});if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func CekKick(optime int64) bool {
	for _, tar := range botstate.Opkick {
		if tar == optime {
			return false
		}
	}
	return true
}

func Deljoin(user string) {
	for _, us := range botstate.Opjoin {
		if us == user {
		}
	}
}

func Cekjoin(optime string) bool {
	defer linetcr.PanicOnly()
	for _, tar := range botstate.Opjoin {
		if tar == optime {
			return false
		}
	}
	return true
}

func CekOp2(optime int64) bool {
	for _, tar := range botstate.Cekoptime {
		if tar == optime {
			return false
		}
	}
	return true
}

func KickDirt(client *linetcr.Account, to, pelaku string) {
	runtime.GOMAXPROCS(botstate.Cpu)
	cans := linetcr.Actor(to)
	for _, cl := range cans {
		if linetcr.GetRoom(to).Act(cl) {
			in := cl.DeleteOtherFromChat(to, pelaku)
			if in == 35 || in == 10 {
				continue
			} else {
				break
			}
		}
	}
}

func CekPurge(optime int64) bool {
	defer linetcr.PanicOnly()
	for _, tar := range botstate.PurgeOP {
		if tar == optime {
			return false
		}
	}
	return true
}
'''
content += "\n" + new_functions

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)
