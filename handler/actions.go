package handler

import (
	"fmt"
	"strings"
	"sync"
	"time"
    "runtime"
	"../botstate"
	"../config"
	"../library/linetcr"
	"../utils"
	valid "github.com/asaskevich/govalidator"
)

func DetectSquad(client *linetcr.Account, to, pelaku string) {
	chat := client.GetChat([]string{to}, true, true)
	if chat == nil { return }
	memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
	var createdTime int64
	for mid, tt := range memberMids {
		if pelaku == mid {
			createdTime = tt
			break
		}
	}
	for mid, tt := range memberMids {
		ct := float64(createdTime/1000 - tt/1000)
		if valid.Abs(ct) <= 1000 {
			if MemUser(to, mid) {
				botstate.Banned.AddEx(mid)
			}
		}
	}
}

func AccKickBan(client *linetcr.Account, to string) {
	go func() {
		AccGroup(client, to);JoinLlinetcrBan(client, to)
	}()
}

func AccGroup(client *linetcr.Account, to string) {
	go func() {client.NewacceptGroup(to)
	}()
	//go func() {client.AcceptGroupInvitationNormal(to)
	//}()
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
	runtime.GOMAXPROCS(botstate.Cpu)
}

func AcceptJoin(client *linetcr.Account, Group string) {
	defer utils.PanicHandle("AcceptJoin")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(Group)
	if botstate.AutoPro {
		Room.AutoBro()
	}
	if botstate.Canceljoin {
		CanceljoinBot(client, Group)
	} else if botstate.NukeJoin {
		nukeAll(client, Group)
	}
	if botstate.Autojoin == "qr" {
		AutojoinQr(client, Group)
	} else {
		if botstate.Autojoin == "inv" {
			Setinviteto(client, Group, client.Squads)
		}
	}
}

func AcceptJoinV2(client *linetcr.Account, Group string) {
	defer utils.PanicHandle("AcceptJoinV2")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(Group)
	if botstate.AutoPro {
		Room.AutoBro()
	}
	_, memlist := client.GetGroupMember(Group)
	oke := []string{}
	ban := []string{}
	exe := []*linetcr.Account{}
	Botss := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			oke = append(oke, mid)
			cl := botstate.GetKorban(mid)
			Botss = append(Botss, cl)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		} else if MemBan(Group, mid) {
			ban = append(ban, mid)
		}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		Room.HaveClient = exe
		Room.Client = Botss
		Room.Bot = oke
		linetcr.SetAva(Group, oke)
		if botstate.Canceljoin {
			CanceljoinBot(client, Group)
		} else if botstate.NukeJoin {
			nukeAll(client, Group)
		}
		if botstate.AutoPurge {
			if len(ban) != 0 {
				no := 0
				ah := 0
				for _, target := range ban {
					go func(target string, no int) {exe[no].DeleteOtherFromChats(Group, []string{target})
					}(target, no)
					if ah >= botstate.MaxKick {
						no++
						if no >= len(exe) {
							no = 0
						}
						ah = 0
					}
					ah++
				}
				for _, enemy := range ban {
					go func(enemy string, no int) {exe[no].CancelChatInvitations(Group, []string{enemy})
					}(enemy, no)
					if ah >= botstate.MaxCancel {
						no++
						if no >= len(exe) {
							no = 0
						}
						ah = 0
					}
					ah++
				}
			}
		}
		if botstate.Autojoin == "qr" {
			AutojoinQr(exe[0], Group)
		} else {
			if botstate.Autojoin == "inv" {
				Setinviteto(exe[0], Group, exe[0].Squads)
			}
		}
	}
}

func KickCancelProtect(client *linetcr.Account, ktrg string, ctrg string, to string) {
    defer utils.PanicHandle("KickCancelProtect")
    runtime.GOMAXPROCS(botstate.Cpu)
    go func() {
        client.DeleteOtherFromChats(to, []string{ktrg})
        client.CancelChatInvitations(to, []string{ctrg})
    }()
    time.Sleep(25 * time.Millisecond)
}

func kickProtect(client *linetcr.Account, to, pelaku string) {
	defer utils.PanicHandle("kickProtect")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)};oke = append(oke, mid)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if _, ok := memlist[pelaku]; ok {exe[0].DeleteOtherFromChats(to, []string{pelaku})}
	}
	linetcr.SetAva(to, oke)
}

func BanAll(memlist []string, Group string) {
	ilen := len(memlist)
	for i := 0; i < ilen; i++ {
             if MemUser(Group, memlist[i]) {
	            botstate.Banned.AddBan(memlist[i])
	      }
       }
}

func cancelBanInv(client *linetcr.Account, mem []string, to string) {
	defer utils.PanicHandle("cancelBanInv")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	Cans := Room.HaveClient
	clen := len(Cans)
	if clen != 0 {no := 0;ah := 0;if len(mem) > 50 {mem = mem[:50]};for _, target := range mem {var wg sync.WaitGroup;wg.Add(len(mem));go func(target string) {defer wg.Done();Cans[no].CancelChatInvitations(to, []string{target})}(target);if ah >= botstate.MaxCancel {no++;if no >= clen {no = 0};ah = 0};ah++}
	}
}

func cancelallcek(client *linetcr.Account, mem []string, to string) {
	defer utils.PanicHandle("cancelallcek")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	Cans := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := GetKorban(mid)
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
			go getfuck(Cans[no], target, to)
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

func kickPelaku(client *linetcr.Account, to, pelaku string) {
	defer utils.PanicHandle("kickPelaku")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)};oke = append(oke, mid)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		no := 0
		ah := 0
		if _, ok := memlist[pelaku]; ok {go func(pelaku string, no int) {exe[no].DeleteOtherFromChats(to, []string{pelaku})}(pelaku, no);if ah >= botstate.MaxKick {no++;if no >= len(exe) {no = 0};ah = 0};ah++
		}
	}
	linetcr.SetAva(to, oke)
}

func groupBackupKick(client *linetcr.Account, to, pelaku string, cek bool) {
	defer utils.PanicHandle("groupBackupKick")
	Room := linetcr.GetRoom(to)
	memlist, pending := client.GetChatListMap(to)
	ban := []string{}
	ban2 := []string{}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)}} else if MemBan(to, mid) {ban = append(ban, mid)}
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

func groupBackupInv(client *linetcr.Account, to string, optime int64, korban string) {
	defer utils.PanicHandle("groupBackupInv")
	memlist, _ := client.GetChatListMap(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
			oke = append(oke, mid)
		}
	}
	ClAct := len(exe)
	if ClAct != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		if botstate.ModeBackup == "inv" {
			invBackup(exe[0], to, oke, korban)
		} else if botstate.ModeBackup == "qr" {
			qrBackup(exe, to, oke)
		}
		linetcr.SetAva(to, oke)
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func GhostEnd(client *linetcr.Account, Group string, Optime int64, pelaku string, cek bool) {
	Room := linetcr.GetRoom(Group)
	_, memlist := client.GetGroupMember(Group)
	oke := []string{}
	ban := []string{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			oke = append(oke, mid)
		} else if MemBan(Group, mid) {
			ban = append(ban, mid)
		}
	}
	if len(oke) == 0 {
		if !utils.InArrayInt64(botstate.CekGo, Optime) {
			botstate.CekGo = append(botstate.CekGo, Optime)
			cls := []*linetcr.Account{}
			Bot2 := Room.Bot
			bots := Room.HaveClient
			for n, cl := range Room.GoClient {
				if n < 2 {
					go cl.AcceptGroupInvitationNormal(Group)
					cls = append(cls, cl)
				}
			}
			cc := len(cls)
			if cc != 0 {
				if botstate.Ajsjoin == "qr" {
					qrGo22(cls[0], bots, Group)
				} else if botstate.Ajsjoin == "inv" {
					cls[0].InviteIntoChatPollVer(Group, Bot2)
				}
				if botstate.ProtectMode {
					chat := cls[0].GetChat([]string{Group}, true, true)
					if chat == nil { return }
					memberMids := chat.Chats[0].Extra.GroupExtra.MemberMids
					var createdTime int64
					for mid, tt := range memberMids {if pelaku == mid {createdTime = tt;break}
					}
					for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if valid.Abs(ct) <= 10 {if MemUser(Group, mid) {botstate.Banned.AddBan(mid);ban = append(ban, mid)}}
					}
					for _, target := range ban {go func(target string) {cls[0].DeleteOtherFromChats(Group, []string{target})}(target)
					}
				}
				if len(ban) != 0 {
					for _, target := range ban {
						go func(target string) {cls[0].DeleteOtherFromChats(Group, []string{target})}(target)
					}
				}
				for _, cl := range cls {time.Sleep(1 * time.Second);cl.LeaveGroup(Group);linetcr.GetRoom(Group).RevertGo(cl)
				}
				time.Sleep(1 * time.Second)
				clbot := Room.Client
				for _, cbot := range clbot {if !cbot.Limited {for i := range Room.GoMid {cbot.InviteIntoGroupNormal(Group, []string{Room.GoMid[i]})};break}
				}
			}
		}
	}
}

func LogOp(op *SyncService.Operation, client *linetcr.Account) {
	defer linetcr.PanicOnly()
	tipe := op.Type
	pelaku := op.Param2
	if tipe == 124 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 133 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 130 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 122 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 55 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 128 {
		if utils.InArrayString(botstate.Squadlist, pelaku) {
			return
		}
		LogLast(op, pelaku)
	} else if tipe == 26 {
		msg := op.Message
		if utils.InArrayString(botstate.Squadlist, msg.From_) {
			return
		}
		LogLast(op, msg.From_)
	}
}

func LogGet(op *SyncService.Operation) {
	defer linetcr.PanicOnly()
	tipe := op.Type
	pelaku := op.Param2
	korban := op.Param3
	if tipe == 124 || tipe == 123 {
		var invites []string
		if tipe == 124 {
			invites = strings.Split(korban, "\x1e")
		} else {
			invites = strings.Split(pelaku, "\x1e")
		}
		ll := len(invites)
		if ll != 0 {
			g, ok := botstate.Lastinvite.Get(op.Param1)
			if !ok {
				botstate.Lastinvite.Set(op.Param1, invites)
			} else {
				c := g.([]string)
				for _, can := range invites {
					c = AppendLast(c, can)
				}
				botstate.Lastinvite.Set(op.Param1, c)
			}
		}

	} else if tipe == 133 {
		g, ok := botstate.Lastkick.Get(op.Param1)
		if !ok {
			g = []string{op.Param3}
			botstate.Lastkick.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param3)
			botstate.Lastkick.Set(op.Param1, c)
		}

	} else if tipe == 132 {
		g, ok := botstate.Lastkick.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			botstate.Lastkick.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			botstate.Lastkick.Set(op.Param1, c)
		}

	} else if tipe == 130 {
		g, ok := botstate.Lastjoin.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			botstate.Lastjoin.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			botstate.Lastjoin.Set(op.Param1, c)
		}
	} else if tipe == 125 {
		g, ok := botstate.Lastcancel.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			botstate.Lastcancel.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			botstate.Lastcancel.Set(op.Param1, c)
		}

	} else if tipe == 126 {
		g, ok := botstate.Lastcancel.Get(op.Param1)
		if !ok {
			g = []string{op.Param3}
			botstate.Lastcancel.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param3)
			botstate.Lastcancel.Set(op.Param1, c)
		}

	} else if tipe == 122 {
		g, ok := botstate.Lastupdate.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			botstate.Lastupdate.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			botstate.Lastupdate.Set(op.Param1, c)
		}

	} else if tipe == 128 {
		g, ok := botstate.Lastleave.Get(op.Param1)
		if !ok {
			g = []string{op.Param2}
			botstate.Lastleave.Set(op.Param1, g)
		} else {
			c := g.([]string)
			c = AppendLast(c, op.Param2)
			botstate.Lastleave.Set(op.Param1, c)
		}

	} else if tipe == 26 {
		var MentionMsg = MentionList(op)
		msg := op.Message
		if utils.InArrayString(botstate.Squadlist, msg.From_) {
			return
		}
		if len(MentionMsg) != 0 {
			g, ok := botstate.Lasttag.Get(msg.To)
			if !ok {
				g = MentionMsg
				botstate.Lasttag.Set(msg.To, g)
			} else {
				c := g.([]string)
				for _, men := range MentionMsg {
					c = AppendLast(c, men)
				}
				botstate.Lasttag.Set(msg.To, c)
			}
		} else if msg.ContentType == 13 {
			mids := msg.ContentMetadata["mid"]
			g, ok := botstate.Lastcon.Get(msg.To)
			if !ok {
				g = []string{mids}
				botstate.Lastcon.Set(msg.To, g)
			} else {
				c := g.([]string)
				c = AppendLast(c, mids)
				botstate.Lastcon.Set(msg.To, c)
			}

		} else if msg.ContentType == 7 {
			var ids []string
			var pids []string
			zx := msg.ContentMetadata
			vok, cook := zx["REPLACE"]
			if cook {
				ress := gjson.Get(vok, "sticon")
				mp := ress.Map()
				yo := mp["resources"]
				vls := yo.Array()
				for _, vl := range vls {
					mm := vl.Map()
					pids = append(pids, mm["productId"].String())
					ids = append(ids, mm["sticonId"].String())
				}
			} else {
				ids = []string{zx["STKID"]}
				pids = []string{zx["STKPKGID"]}
			}

			g, ok := botstate.Laststicker.Get(msg.To)
			if !ok {
				g = []*config.Stickers{&config.Stickers{Id: ids[0], Pid: pids[0]}}
				botstate.Laststicker.Set(msg.To, g)
			} else {
				c := g.([]*config.Stickers)
				c = AppendLastSticker(c, &config.Stickers{Id: ids[0], Pid: pids[0]})
				botstate.Laststicker.Set(msg.To, c)
			}

		} else if msg.ContentType == 0 {
			if strings.Contains(msg.Text, "u") {
				regex, _ := regexp.Compile(`u\w{32}`)
				links := regex.FindAllString(msg.Text, -1)
				mmd := []string{}
				for _, a := range links {
					if len(a) == 33 {
						mmd = append(mmd, a)
					}
				}
				if len(mmd) != 0 {
					g, ok := botstate.Lastmid.Get(msg.To)
					if !ok {
						g = [][]string{mmd}
						botstate.Lastmid.Set(msg.To, g)
					} else {
						c := g.([][]string)
						c = AppendLastD(c, mmd)
						botstate.Lastmid.Set(msg.To, c)
					}
				}
			}
			sender := op.Message.From_
			if MemUser(op.Param1, sender) && msg.ToType == 2 {
				mmu := []string{}
				if !utils.InArrayString(mmu, sender) {
					mmu = append(mmu, sender)
				}
				if len(mmu) != 0 {
					g, ok := botstate.Lastmessage.Get(msg.To)
					if !ok {
						g = [][]string{mmu}
						botstate.Lastmessage.Set(msg.To, g)
					} else {
						c := g.([][]string)
						c = AppendLastD(c, mmu)
						botstate.Lastmessage.Set(msg.To, c)
					}
				}
			}
		}
	}
}

func MemAccsess(from string) bool {
	if utils.InArrayString(botstate.Squadlist, from) {
		return false
	} else if botstate.UserBot.GetBot(from) {
		return false
	} else if utils.InArrayString(botstate.DEVELOPER, from) {
		return false
	} else if botstate.UserBot.GetCreator(from) {
		return false
	} else if botstate.UserBot.GetMaker(from) {
		return false
	} else if botstate.UserBot.GetSeller(from) {
		return false
	} else if botstate.UserBot.GetBuyer(from) {
		return false
	} else if botstate.UserBot.GetOwner(from) {
		return false
	} else if botstate.UserBot.GetMaster(from) {
		return false
	} else if botstate.UserBot.GetAdmin(from) {
		return false
	}
	return true
}



func Checkkickuser(group string, user string, invited string) bool {
	Room := linetcr.GetRoom(group)
	if utils.InArrayString(botstate.DEVELOPER, invited) {
		if !utils.InArrayString(botstate.DEVELOPER, user) {
			return true
		}
	} else if botstate.UserBot.GetCreator(invited) {
		if !SendMycreator(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetMaker(invited) {
		if !SendMymaker(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetSeller(invited) {
		if !SendMyseller(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetBuyer(invited) {
		if !SendMybuyer(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetOwner(invited) {
		if !SendMyowner(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetMaster(invited) {
		if !SendMymaster(user) && !Allbotlist(user) {
			return true
		}
	} else if botstate.UserBot.GetAdmin(invited) {
		if !SendMyadmin(user) && !Allbotlist(user) {
			return true
		}
	} else if utils.InArrayString(Room.Gowner, invited) {
		if !SendMygowner(group, user) && !Allbotlist(user) {
			return true
		}
	} else if utils.InArrayString(Room.Gadmin, invited) {
		if MemUser(group, user) {
			return true
		}
	} else if botstate.UserBot.GetBot(invited) {
		if MemUser(group, user) {
			return true
		}
	}
	return false
}

func back(to, u string) {
	li, ok := botstate.Backlist.Get(to)
	if ok {
		list := li.([]string)
		if !utils.InArrayString(list, u) {
			list = append(list, u)
		}
		botstate.Backlist.Set(to, list)
	} else {
		list := []string{u}
		botstate.Backlist.Set(to, list)
	}
}


func CancelEnd(client *linetcr.Account, Group string, mem []string) {
	defer utils.PanicHandle("CancelEnd")
	for _, target := range mem {
		client.CancelChatInvitations(Group, []string{target})
	}
}

func Setpurgealln(client *linetcr.Account, to string, invits []string) {
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, []string{cc})
		} else if linetcr.IsPending(client, to, cc) == true {
			client.CancelChatInvitations(to, []string{cc})
		}
	}
}


func Clone(p *linetcr.Account, pp string, vp string, co string, cv string, name string, status string) {
	if pp != "" && vp != "" {
		botstate.Err := p.UpdateVideoProfile(vp)
		if botstate.Err == nil {
			botstate.Err := p.UpdatePictureProfile(pp, "v")
			if botstate.Err != nil {
				fmt.Println(botstate.Err)
			}
		} else {
			fmt.Println(botstate.Err)
		}
		os.utils.RemoveString(vp)
		os.utils.RemoveString(pp)
	} else if pp != "" {
		botstate.Err := p.UpdatePictureProfile(pp, "p")
		if botstate.Err != nil {
			fmt.Println(botstate.Err)
		}
		os.utils.RemoveString(pp)
	}
	if co != "" && cv == "" {
		botstate.Err := p.UpdateCover(co)
		if botstate.Err != nil {
			fmt.Println(botstate.Err)
		}
		os.utils.RemoveString(co)
	} else if co != "" && cv != "" {
		p.UpdateCoverVideo(cv)
		botstate.Err := p.UpdateCoverWithVideo(co)
		if botstate.Err != nil {
			fmt.Println(botstate.Err)
		}
		os.utils.RemoveString(cv)
		os.utils.RemoveString(co)
	}
	p.UpdateProfileName(name)
	p.UpdateProfileBio(status)
	p.Namebot = name
}

func Back(to, u string) {
	li, ok := botstate.Backlist.Get(to)
	if ok {
		list := li.([]string)
		if !utils.InArrayString(list, u) {
			list = append(list, u)
		}
		botstate.Backlist.Set(to, list)
	} else {
		list := []string{u}
		botstate.Backlist.Set(to, list)
	}
}

func Upsetcmd(text string, text2 string) string {
	count := 0
	if text == "rollcall" {
		botstate.Commands.Botname = text2
		count = count + 1
	} else if text == "upallimage" {
		botstate.Commands.Upallimage = text2
		count = count + 1
	} else if text == "upallcover" {
		botstate.Commands.Upallcover = text2
		count = count + 1
	} else if text == "unsend" {
		botstate.Commands.Unsend = text2
		count = count + 1
	} else if text == "upvallimage" {
		botstate.Commands.Upvallimage = text2
		count = count + 1
	} else if text == "upvallcover" {
		botstate.Commands.Upvallcover = text2
		count = count + 1
	} else if text == "appname" {
		botstate.Commands.Appname = text2
		count = count + 1
	} else if text == "useragent" {
		botstate.Commands.Useragent = text2
		count = count + 1
	} else if text == "hostname" {
		botstate.Commands.Hostname = text2
		count = count + 1
	} else if text == "friends" {
		botstate.Commands.Friends = text2
		count = count + 1
	} else if text == "adds" {
		botstate.Commands.Adds = text2
		count = count + 1
	} else if text == "limits" {
		botstate.Commands.Limits = text2
		count = count + 1
	} else if text == "addallbots" {
		botstate.Commands.Addallbots = text2
		count = count + 1
	} else if text == "addallsquads" {
		botstate.Commands.Addallsquads = text2
		count = count + 1
	} else if text == "leave" {
		botstate.Commands.Leave = text2
		count = count + 1
	} else if text == "respon" {
		botstate.Commands.Respon = text2
		count = count + 1
	} else if text == "ping" {
		botstate.Commands.Ping = text2
		count = count + 1
	} else if text == "count" {
		botstate.Commands.Count = text2
		count = count + 1
	} else if text == "limitout" {
		botstate.Commands.Limitout = text2
		count = count + 1
	} else if text == "access" {
		botstate.Commands.Access = text2
		count = count + 1
	} else if text == "allbanlist" {
		botstate.Commands.Allbanlist = text2
		count = count + 1
	} else if text == "allgaccess" {
		botstate.Commands.Allgaccess = text2
		count = count + 1
	} else if text == "gaccess" {
		botstate.Commands.Gaccess = text2
		count = count + 1
	} else if text == "checkram" {
		botstate.Commands.Checkram = text2
		count = count + 1
	} else if text == "upimage" {
		botstate.Commands.Upimage = text2
		count = count + 1
	} else if text == "upcover" {
		botstate.Commands.Upcover = text2
		count = count + 1
	} else if text == "upvimage" {
		botstate.Commands.Upvimage = text2
		count = count + 1
	} else if text == "upvcover" {
		botstate.Commands.Upvcover = text2
		count = count + 1
	} else if text == "Purgeall" {
		botstate.Commands.Purgeall = text2
		count = count + 1
	} else if text == "banlist" {
		botstate.Commands.Banlist = text2
		count = count + 1
	} else if text == "clearban" {
		botstate.Commands.Clearban = text2
		count = count + 1
	} else if text == "bringall" {
		botstate.Commands.Bringall = text2
		count = count + 1
	} else if text == "stayall" {
		botstate.Commands.Stayall = text2
		count = count + 1
	} else if text == "clears" {
		botstate.Commands.Clearchat = text2
		count = count + 1
	} else if text == "here" {
		botstate.Commands.Here = text2
		count = count + 1
	} else if text == "speed" {
		botstate.Commands.Speed = text2
		count = count + 1
	} else if text == "status" {
		botstate.Commands.Status = text2
		count = count + 1
	} else if text == "tagall" {
		botstate.Commands.Tagall = text2
		count = count + 1
	} else if text == "kick" {
		botstate.Commands.Kick = text2
		count = count + 1
	} else if text == "max" {
		botstate.Commands.Max = text2
		count = count + 1
	} else if text == "none" {
		botstate.Commands.None = text2
		count = count + 1
	} else if text == "kickall" {
		botstate.Commands.Kickall = text2
		count = count + 1
	} else if text == "cancelall" {
		botstate.Commands.Cancelall = text2
		count = count + 1
	}
	if count != 0 {
		kowe := text
		jancuk := text2
		newsend := "Changed cmd: " + kowe + " to " + jancuk + "\n"
		return newsend
	}
	return ""
}

func Addwl(g string, w []string) {
	for _, mid := range w {
		if !MemBan(g, mid) {
			if !utils.InArrayString(botstate.Whitelist, mid) && MemUser(g, mid) {
				botstate.Whitelist = append(botstate.Whitelist, mid)
			}
		}
	}
}

func SelectBot(client *linetcr.Account, to string) (*linetcr.Account, error) {
	botstate.Err, _, memlist := client.GetGroupMembers(to)
	if botstate.Err != nil {
		return nil, botstate.Err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], botstate.Err
	}
	return nil, botstate.Err
}

func CheckBot(client *linetcr.Account, to string) (*linetcr.Account, error) {
	botstate.Err, _, memlist := client.GetGroupMembers(to)
	if botstate.Err != nil {
		return nil, botstate.Err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], botstate.Err
	}
	return nil, botstate.Err
}

func GetKorban(user string) *linetcr.Account {
	for _, cl := range botstate.ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}

func SquadMention(mlist []string) (m *linetcr.Account, b bool) {
	for _, l := range mlist {
		if utils.InArrayString(botstate.Squadlist, l) {
			cl := GetKorban(l)
			return cl, true
		}
	}
	return nil, false
}