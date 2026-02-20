package handler

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"

	"../botstate"
	"../library/linetcr"
	"../utils"
	valid "github.com/asaskevich/govalidator"
)

func NukeAll(Client *linetcr.Account, Group string) {
	defer linetcr.PanicOnly()
	memlist := []string{}
	_, memlists := Client.GetGroupMember(Group)
	act := []*linetcr.Account{}
	for mid, _ := range memlists {
		if MemUser(Group, mid) {
			memlist = append(memlist, mid)
		} else if utils.InArrayString(botstate.Squadlist, mid) {
			cl := botstate.GetKorban(mid)
			if !cl.Limited {
				act = append(act, cl)
			}
		}
	}
	lact := len(act)
	if lact == 0 {
		return
	} else {
		sort.Slice(act, func(i, j int) bool {
			return act[i].KickPoint < act[j].KickPoint
		})
		celek := len(memlist)
		if celek < botstate.MaxKick || lact == 1 {
			cl := act[0]
			for _, mem := range memlist {
				go cl.DeleteOtherFromChat(Group, mem)
			}
		} else {
			hajar := []string{}
			z := celek / botstate.MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= lact {
					no = 0
				}
				go func(Group string, no int, i int, z int, memlist []string, act []*linetcr.Account) {
					Client = act[no]
					if i == z {
						hajar = memlist[i*botstate.MaxKick:]
					} else {
						hajar = memlist[i*botstate.MaxKick : (i+1)*botstate.MaxKick]
					}
					if len(hajar) != 0 {
						for _, target := range hajar {
							go Client.DeleteOtherFromChat(Group, target)
						}
					}
				}(Group, no, i, z, memlist, act)
				no += 1
			}
		}
		linetcr.GetRoom(Group).HaveClient = act
	}
}

func KickCancelBan129(client *linetcr.Account, to string) {
	go runtime.Gosched()
	x, _ := client.NewGetChat(to)
	grup := x.Chats[0].Extra.GroupExtra.MemberMids
	g, _ := client.NewGetChat(to)
	for _, v := range client.Backup {
		if _, cek := grup[v]; cek {
			if v == client.MID {
				var Batas = 0
				if g != nil {
					targets := g.Chats[0].Extra.GroupExtra.MemberMids
					targets2 := g.Chats[0].Extra.GroupExtra.InviteeMids
					listMid := []string{}
					listMid2 := []string{}
					for cok := range targets {
						listMid = append(listMid, cok)
					}
					for cok := range targets2 {
						listMid2 = append(listMid2, cok)
					}
					for v := range botstate.Banned.Banlist {
						go runtime.Gosched()
						if contains(listMid, botstate.Banned.Banlist[v]) {
							go runtime.Gosched()
							go func(v string) { client.NewkickGroup(to, v) }(botstate.Banned.Banlist[v])
							Batas = Batas + 1
							if int64(Batas) >= int64(botstate.MaxKick) {
								Batas = 0
								break
							}
						} else if contains(listMid2, botstate.Banned.Banlist[v]) {
							go runtime.Gosched()
							go func(v string) { client.NewcancelGroup(to, v) }(botstate.Banned.Banlist[v])
							Batas = Batas + 1
							if int64(Batas) >= int64(botstate.MaxCancel) {
								Batas = 0
								break
							}
						}
					}
				}
				break
			} else {
				continue
			}
		}
	}
}

func KickBan132V2(client *linetcr.Account, to string) {
	var Batas = 0
	chat, _ := client.NewGetChat(to)
	if chat != nil {
		memb := chat.Chats[0].Extra.GroupExtra.MemberMids
		for x := range client.Backup {
			if _, blog := memb[client.Backup[x]]; blog {
				if client.MID == client.Backup[x] {
					go func() {
						for mid := range memb {
							if IsBlacklist(client, mid) == true {
								go func(mid string) { go client.NewkickGroup(to, mid) }(mid)
								Batas = Batas + 1
								if int64(Batas) >= int64(botstate.MaxKick) {
									Batas = 0
									break
								}
							}
						}
					}()
					break
				} else {
					continue
				}
			}
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func CancelBan125V2(client *linetcr.Account, to string) {
	var Batas = 0
	g, _ := client.NewGetChat(to)
	if g != nil {
		memb := g.Chats[0].Extra.GroupExtra.MemberMids
		memb1 := g.Chats[0].Extra.GroupExtra.InviteeMids
		for x := range client.Backup {
			if _, blog := memb[client.Backup[x]]; blog {
				if client.MID == client.Backup[x] {
					var wg sync.WaitGroup
					wg.Add(len(memb1))
					go func() {
						for mid := range memb1 {
							if IsBlacklist(client, mid) == true {
								go func(mid string) {
									defer wg.Done()
									go client.NewcancelGroup(to, mid)
								}(mid)
								Batas = Batas + 1
								if int64(Batas) >= int64(botstate.MaxCancel) {
									Batas = 0
									break
								}
							}
						}
					}()
					wg.Wait()
					break
				} else {
					continue
				}
			}
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func KickBan132(client *linetcr.Account, to string, user string) {
	if len(user) > 10 {
		go func() { FastKick(client, to) }()
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func CancelBan125(client *linetcr.Account, to string, korban []string) {
	if len(korban) > 10 {
		go func() { FastCancel(client, to) }()
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func FastKick(client *linetcr.Account, to string) {
	var Batas = 0
	c, _ := client.GetChats([]string{to})
	mem := c.Chats[0].Extra.GroupExtra.MemberMids
	for k, _ := range mem {
		if IsBlacklist(client, k) == true {
			go func(k string) {
				go client.DeleteOtherFromChats(to, []string{k})
			}(k)
			Batas = Batas + 1
			if Batas >= botstate.MaxKick-1 {
				Batas = 0
				break
			}
		}
	}
}

func FastCancel(client *linetcr.Account, to string) {
	var Batas = 0
	c, _ := client.GetChats([]string{to})
	pend := c.Chats[0].Extra.GroupExtra.InviteeMids
	for k, _ := range pend {
		if IsBlacklist(client, k) == true {
			go func(k string) {
				go client.CancelChatInvitations(to, []string{k})
			}(k)
			Batas = Batas + 1
			if Batas >= botstate.MaxCancel-1 {
				Batas = 0
				break
			}
		}
	}
}

func NodeBans(client *linetcr.Account, to string, korban []string) {
	if len(korban) > 3 {
		go func() { FastKick(client, to); FastCancel(client, to) }()
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func PurgeFaster(client *linetcr.Account, to string) {
	for x := range botstate.Squadlist {
		if linetcr.IsMembers(client, to, botstate.Squadlist[x]) == true {
			if client.MID == botstate.Squadlist[x] {
				go func() { FastKick(client, to); FastCancel(client, to) }()
			}
			break
		} else {
			continue
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func KickMemBan(client *linetcr.Account, to string) {
	var Batas = 0
	c, _ := client.GetChatsV2(to)
	mem := c.Extra.GroupExtra.MemberMids
	for k, _ := range mem {
		if MemBan(to, k) == true {
			go func(k string) { go client.DeleteOtherFromChats(to, []string{k}) }(k)
			Batas = Batas + 1
			if Batas >= botstate.MaxKick-1 {
				Batas = 0
				break
			}
		}
	}
}

func CanMemBan(client *linetcr.Account, to string) {
	var Batas = 0
	c, _ := client.GetChatsV2(to)
	pend := c.Extra.GroupExtra.InviteeMids
	for k, _ := range pend {
		if MemBan(to, k) == true {
			go func(k string) { go client.DeleteOtherFromChats(to, []string{k}) }(k)
			Batas = Batas + 1
			if Batas >= botstate.MaxKick-1 {
				Batas = 0
				break
			}
		}
	}
}

func KickCancel(client *linetcr.Account, to string) {
	for _, v := range botstate.Banned.Banlist {
		if botstate.Banned.GetBan(v) {
			if linetcr.IsMembers(client, to, v) == true {
				go func(v string) { client.DeleteOtherFromChats(to, []string{v}) }(v)
			}
			if linetcr.IsPending(client, to, v) == true {
				if botstate.Banned.GetBan(v) {
					go func(v string) { client.CancelChatInvitations(to, []string{v}) }(v)
				}
			}
		}
	}
}





func Autokickban(client *linetcr.Account, to string, target string) {
	if botstate.AutokickBan {
		gr, _ := client.GetGroupIdsJoined()
		for _, aa := range gr {
			go client.DeleteOtherFromChats(aa, []string{target})
			go client.CancelChatInvitations(aa, []string{target})
		}
	}
}

func Purgemode(Client *linetcr.Account, Group string) {
	defer utils.PanicHandle("Purgemode")
	_, memlists := Client.GetGroupMember(Group)
	for target, _ := range memlists {
		if MemBan(Group, target) {
			go Client.DeleteOtherFromChat(Group, target)
		}
	}
}

func KIckbansPurges(client *linetcr.Account, group string) {
	defer linetcr.PanicOnly()
	gr, _ := client.GetGroupIdsJoined()
	for _, aa := range gr {
		c, _ := client.GetChats([]string{aa})
		zxc := c.Chats[0].Extra.GroupExtra.MemberMids
		for k, _ := range zxc {
			if IsBlacklist(client, k) == true {
				go func(k string) {
					go client.DeleteOtherFromChats(aa, []string{k})
				}(k)
			}
		}
		_, _, pind := client.GetChatList(aa)
		for _, i := range pind {
			if IsBlacklist(client, i) == true {
				go func(i string) {
					go client.CancelChatInvitations(aa, []string{i})
				}(i)
			}
		}
	}
	client.SendMessage(group, botstate.Fancy("Success nukebanlist"))
}

func KickAllBan(client *linetcr.Account, to string) {
	_, mem, pend := client.GetChatList(to)
	for _, k := range mem {
		if IsBlacklist(client, k) == true {
			go func(k string) { go client.DeleteOtherFromChats(to, []string{k}) }(k)
		}
	}
	for _, i := range pend {
		if IsBlacklist(client, i) == true {
			go func(i string) { go client.CancelChatInvitations(to, []string{i}) }(i)
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func KIckbansPurges1(client *linetcr.Account, group string) {
	defer linetcr.PanicOnly()
	gr, _ := client.GetGroupIdsJoined()
	nus := []string{}
	list := ""
	list += fmt.Sprintf("Purged %v groups: \n", len(gr))
	for num, aa := range gr {
		num++
		for _, v := range botstate.Banned.Banlist {
			if linetcr.IsMembers(client, aa, v) == true {
				if botstate.Banned.GetBan(v) {
					go func(v string) { client.DeleteOtherFromChats(aa, []string{v}) }(v)
					if linetcr.IsPending(client, aa, v) == true {
						client.CancelChatInvitations(aa, []string{v})
					}
					x, botstate.Err := client.GetContact(v)
					rengs := strconv.Itoa(num)
					if botstate.Err != nil {
						list += "\n " + rengs + ". Closed Account"
					} else {
						nus = append(nus, v)
						list += "\n " + rengs + ". " + x.DisplayName
					}
				}
			}
		}
	}
	list += fmt.Sprintf("\n\nTotal kicks: %v.", len(nus))
	client.SendMessage(group, botstate.Fancy(list))
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

func JKickFuck(client *linetcr.Account, to string) {
	c, _ := client.GetGroup3(to)
	zxc := c.Extra.GroupExtra.MemberMids
	for k, _ := range zxc {
		if IsBlacklist(client, k) == true {
			go func(to string, k string) {
				go client.DeleteOtherFromChats(to, []string{k})
			}(to,k)
		}
	}
}
func JCancelFuck(client *linetcr.Account, to string) {
	c, _ := client.GetGroup3(to)
	zxc := c.Extra.GroupExtra.InviteeMids
	for k, _ := range zxc {
		if IsBlacklist(client, k) == true {
			go func(to string, k string) {
				go client.CancelChatInvitations(to, []string{k})
			}(to,k)
		}
	}
}
func KickCansWar(client *linetcr.Account, to string) {
	go func() {
		JKickFuck(client, to)
	}()
	go func() {
		JCancelFuck(client, to)
	}()
}

func CheckUnbanBots(client *linetcr.Account, to string, targets []string, pl int, sinder string) {
	room := linetcr.GetRoom(to)
	target := []string{}
	for _, from := range targets {
		if botstate.Banned.GetFuck(from) {
			target = append(target, from)
			botstate.Banned.DelFuck(from)
		} else if botstate.Banned.GetBan(from) {
			target = append(target, from)
			botstate.Banned.DelBan(from)
		} else if utils.InArrayString(room.Gban, from) {
			target = append(target, from)
			Ungban(to, from)
		} else if botstate.Banned.GetMute(from) {
			target = append(target, from)
			botstate.Banned.DelMute(from)
		}
	}
	if len(target) != 0 {
		list := ""
		if pl == 1 {
			list += "Removed from banlist:\n"
		} else if pl == 2 {
			list += "Removed from fucklist:\n"
		} else if pl == 3 {
			list += "Removed from gbanlist:\n"
		} else if pl == 4 {
			list += "Removed from mutelist:\n"
		}
		for i := range target {
			list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
		}
		client.SendPollMention(to, list, target)
		if pl == 1 {
			LogAccess(client, to, sinder, "unban", target, 2)
		} else if pl == 2 {
			LogAccess(client, to, sinder, "unfuck", target, 2)
		} else if pl == 3 {
			LogAccess(client, to, sinder, "ungban", target, 2)
		} else if pl == 4 {
			LogAccess(client, to, sinder, "unmute", target, 2)
		}
	} else {
		list := ""
		if pl == 1 {
			list += "User(s) not in banlist.\n"
		} else if pl == 2 {
			list += "User(s) not in fucklist.\n"
		} else if pl == 3 {
			list += "User(s) not in gbanlist.\n"
		} else if pl == 4 {
			list += "User(s) not in mutelist.\n"
		}
		client.SendMessage(to, botstate.Fancy(list))
	}
}

func Purgesip(Group string, cl *linetcr.Account) {
	defer utils.PanicHandle("purgesip")
	mem := make(chan []string)
	go func(m chan []string) {
		memlistss := []string{}
		_, memlists := cl.GetGroupMember(Group)
		for target := range memlists {
			if MemBan(Group, target) {
				memlistss = append(memlistss, target)
			}
		}
		m <- memlistss
	}(mem)
	Cans := linetcr.Actor(Group)
	ClAct := len(Cans)
	hajar := []string{}
	var client *linetcr.Account
	memlist := <-mem
	celek := len(memlist)
	if celek > botstate.MaxKick {
		if ClAct != 0 {
			z := celek / botstate.MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= ClAct {
					no = 0
				}
				if i != 0 {
					client = Cans[no]
				} else {
					client = cl
				}
				if i == z {
					hajar = memlist[i*botstate.MaxKick:]
				} else {
					hajar = memlist[i*botstate.MaxKick : (i+1)*botstate.MaxKick]
				}
				if len(hajar) != 0 {
					for _, target := range hajar {
						go client.DeleteOtherFromChats(Group, []string{target})
					}
				}
				no += 1
			}
		} else if !cl.Limited {
			for _, target := range memlist {
				go cl.DeleteOtherFromChats(Group, []string{target})
			}
		}
	} else if !cl.Limited {
		for _, target := range memlist {
			go cl.DeleteOtherFromChats(Group, []string{target})
		}
	}
}

func KickBl(client *linetcr.Account, to string) {
	defer utils.PanicHandle("detectBl")
	memlist := []string{}
	_, memlists := client.GetGroupMember(to)
	act := []*linetcr.Account{}
	for mid, _ := range memlists {
		if MemBan(to, mid) {
			memlist = append(memlist, mid)
		} else if utils.InArrayString(botstate.Squadlist, mid) {
			cl := botstate.GetKorban(mid)
			if cl.Limited {
				act = append(act, cl)
			}
		}
	}
	lact := len(act)
	if lact == 0 {
		return
	} else {
		sort.Slice(act, func(i, j int) bool {
			return act[i].KickPoint < act[j].KickPoint
		})
		celek := len(memlist)
		if celek < botstate.MaxKick || lact == 1 {
			cl := act[0]
			for _, mem := range memlist {
				go cl.DeleteOtherFromChat(to, mem)
			}
		} else {
			hajar := []string{}
			z := celek / botstate.MaxKick
			y := z + 1
			no := 0
			for i := 0; i < y; i++ {
				if no >= lact {
					no = 0
				}
				go func(to string, no int, i int, z int, memlist []string, act []*linetcr.Account) {
					client = act[no]
					if i == z {
						hajar = memlist[i*botstate.MaxKick:]
					} else {
						hajar = memlist[i*botstate.MaxKick : (i+1)*botstate.MaxKick]
					}
					if len(hajar) != 0 {
						for _, target := range hajar {
							go client.DeleteOtherFromChat(to, target)
						}
					}
				}(to, no, i, z, memlist, act)
				no += 1
			}
		}
	}
}
func AcceptWar(client *linetcr.Account, to string, korban []string) {
	go func() {
		AccGroup(client, to);JoinLlinetcrBan(client, to);for _, i := range korban {go func(i string) {go client.InviteIntoGroupNormal(to, []string{i})}(i)}
	}()
}




func KickBan129(client *linetcr.Account, to string) {
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
}

func GroupBackupWar(client *linetcr.Account, to string) {
	for x := range botstate.Squadlist {
        if linetcr.IsMembers(client, to, botstate.Squadlist[x]) == true {
            if client.MID == botstate.Squadlist[x] {
                go func() {KickCansWar(client, to)}()
            }
            break
        } else {
            continue
        }
	}
	runtime.GOMAXPROCS(100)
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

func KickPelaku(client *linetcr.Account, to, pelaku string) {
	defer utils.PanicHandle("kickPelaku")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := botstate.GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)};oke = append(oke, mid)}
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

func KickProtect(client *linetcr.Account, to, pelaku string) {
	defer utils.PanicHandle("kickProtect")
	runtime.GOMAXPROCS(botstate.Cpu)
	Room := linetcr.GetRoom(to)
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	oke := []string{}
	for mid, _ := range memlist {if utils.InArrayString(botstate.Squadlist, mid) {cl := botstate.GetKorban(mid);if cl.Limited == false {exe = append(exe, cl)};oke = append(oke, mid)}
	}
	if len(exe) != 0 {
		sort.Slice(exe, func(i, j int) bool {return exe[i].KickPoint < exe[j].KickPoint})
		Room.HaveClient = exe
		if _, ok := memlist[pelaku]; ok {exe[0].DeleteOtherFromChats(to, []string{pelaku})}
	}
	linetcr.SetAva(to, oke)
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
					QrGo22(cls[0], bots, Group)
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
					for mid, tt := range memberMids {ct := float64(createdTime/1000 - tt/1000);if math.Abs(ct) <= 10 {if MemUser(Group, mid) {botstate.Banned.AddBan(mid);ban = append(ban, mid)}}
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

func IsBlacklist(client *linetcr.Account, from string) bool {
	if contains(botstate.Banned.Banlist, from) == true {
		return true
	}
	return false
}

func IsBlacklist2(client *linetcr.Account, from string) bool {
	if contains(botstate.Banned.Locklist, from) == true {
		return true
	}
	return false
}

func RemBanFriends(client *linetcr.Account, to string) {
	defer utils.PanicHandle("RemBanFriends")
	if botstate.AllowDoOnce == 0 {
		donedel := []string{}
		for _, cl := range botstate.ClientBot {
			friendz, _ := cl.GetAllContactIds()
			for _, con := range friendz {
				if !linetcr.InArrayCl(linetcr.KickBanChat, cl) && !cl.Frez {
					r, _ := cl.GetHomeProfile(con)
					if linetcr.GetBannedChat(r) == 1 {
						cl.UnFriend(con)
						if !utils.InArrayString(donedel, con) {
							donedel = append(donedel, con)
						}
					}
				} else {
					r, _ := client.GetHomeProfile(con)
					if linetcr.GetBannedChat(r) == 1 {
						cl.UnFriend(con)
						if !utils.InArrayString(donedel, con) {
							donedel = append(donedel, con)
						}
					}
				}
			}
		}
		if len(donedel) == 0 {
			client.SendMessage(to, botstate.Fancy("Nothing Deleted || No C_Ban Friends\n"))
		} else {
			DataMention(to, "✭ Unfriend Banz ✭\n", donedel)
		}
		botstate.AllowDoOnce++
	} else {
		client.SendMessage(to, botstate.Fancy("No botstate.Data For C_Ban Friends"))
	}
}

func ClearCon() {
	n := 0
	for _, cl := range botstate.ClientBot {
		fl, _ := cl.GetAllContactIds()
		for _, x := range fl {
			if !utils.InArrayString(botstate.Squadlist, x) && !utils.InArrayString(botstate.DEVELOPER, x){
				cl.UnFriend(x)
				time.Sleep(2 * time.Second)
			}	
		}
		n += 1
	}
}

func BanAll(memlist []string, Group string) {
	ilen := len(memlist)
	for i := 0; i < ilen; i++ {
             if MemUser(Group, memlist[i]) {
	            botstate.Banned.AddBan(memlist[i])
	      }
       }
}

func Checklistexpel(client *linetcr.Account, to string, targets []string, pl int, sinder string) {
	Room := linetcr.GetRoom(to)
	if len(targets) > 1 {
		target := []string{}
		conts := 0
		conts2 := 0
		for _, from := range targets {
			if utils.InArrayString(botstate.DEVELOPER, from) {
				if !utils.InArrayString(botstate.DEVELOPER, sinder) {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetCreator(from) {
				if utils.InArrayString(botstate.DEVELOPER, sinder) {
					target = append(target, from)
					botstate.UserBot.DelCreator(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetMaker(from) {
				if SendMycreator(sinder) {
					target = append(target, from)
					botstate.UserBot.DelMaker(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetSeller(from) {
				if SendMymaker(sinder) {
					target = append(target, from)
					botstate.UserBot.DelSeller(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetBuyer(from) {
				if SendMyseller(sinder) {
					target = append(target, from)
					botstate.UserBot.DelBuyer(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetOwner(from) {
				if SendMybuyer(sinder) {
					target = append(target, from)
					botstate.UserBot.DelOwner(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetMaster(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					botstate.UserBot.DelMaster(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetAdmin(from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					botstate.UserBot.DelAdmin(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if utils.InArrayString(Room.Gowner, from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					Room.Gowner = utils.RemoveString(Room.Gowner, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if utils.InArrayString(Room.Gadmin, from) {
				if SendMygowner(to, sinder) {
					target = append(target, from)
					Room.Gadmin = utils.RemoveString(Room.Gadmin, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetBot(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					botstate.UserBot.DelBot(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
		}
		if len(target) != 0 {
			list := ""
			if pl == 1 {
				list += "Expeled from Buyer\n"
			} else if pl == 2 {
				list += "Expeled from Owner\n"
			} else if pl == 3 {
				list += "Expeled from Master\n"
			} else if pl == 4 {
				list += "Expeled from Admin\n"
			} else if pl == 5 {
				list += "Expeled from Bot\n"
			} else if pl == 6 {
				list += "Expeled from Gowner\n"
			} else if pl == 7 {
				list += "Expeled from Gadmin\n"
			} else if pl == 8 {
				list += "Expeled from Access\n"
			} else if pl == 9 {
				list += "Expeled from Maker\n"
			} else if pl == 10 {
				list += "Expeled from Creator\n"
			} else if pl == 17 {
				list += "Expeled from Seller\n"
			}
			for i := range target {
				list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
			}
			client.SendPollMention(to, list, target)
			if pl == 2 {
				LogAccess(client, to, sinder, "unowner", target, 2)
			} else if pl == 3 {
				LogAccess(client, to, sinder, "unmaster", target, 2)
			} else if pl == 4 {
				LogAccess(client, to, sinder, "unadmin", target, 2)
			} else if pl == 5 {
				LogAccess(client, to, sinder, "unbot", target, 2)
			} else if pl == 6 {
				LogAccess(client, to, sinder, "ungowner", target, 2)
			} else if pl == 7 {
				LogAccess(client, to, sinder, "ungadmin", target, 2)
			} else if pl == 8 {
				LogAccess(client, to, sinder, "expel", target, 2)
			}
		} else if conts != 0 {
			list := "Sorry, your grade is too low.\n"
			client.SendMessage(to, botstate.Fancy(list))
		} else if conts2 != 0 {
			list := "Users not have access.\n"
			client.SendMessage(to, botstate.Fancy(list))
		}
	} else {
		target := []string{}
		conts := 0
		conts2 := 0
		for _, from := range targets {
			if utils.InArrayString(botstate.DEVELOPER, from) {
				if !utils.InArrayString(botstate.DEVELOPER, sinder) {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetCreator(from) {
				if utils.InArrayString(botstate.DEVELOPER, sinder) {
					target = append(target, from)
					botstate.UserBot.DelCreator(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetMaker(from) {
				if SendMycreator(sinder) {
					target = append(target, from)
					botstate.UserBot.DelMaker(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetSeller(from) {
				if SendMymaker(sinder) {
					target = append(target, from)
					botstate.UserBot.DelSeller(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetBuyer(from) {
				if SendMyseller(sinder) {
					target = append(target, from)
					botstate.UserBot.DelBuyer(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetOwner(from) {
				if SendMybuyer(sinder) {
					target = append(target, from)
					botstate.UserBot.DelOwner(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetMaster(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					botstate.UserBot.DelMaster(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetAdmin(from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					botstate.UserBot.DelAdmin(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if utils.InArrayString(Room.Gowner, from) {
				if SendMyadmin(sinder) {
					target = append(target, from)
					Room.Gowner = utils.RemoveString(Room.Gowner, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if utils.InArrayString(Room.Gadmin, from) {
				if SendMygowner(to, sinder) {
					target = append(target, from)
					Room.Gadmin = utils.RemoveString(Room.Gadmin, from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
			if botstate.UserBot.GetBot(from) {
				if SendMyowner(sinder) {
					target = append(target, from)
					botstate.UserBot.DelBot(from)
				} else {
					conts++
				}
			} else {
				conts2++
			}
		}
		if len(target) != 0 {
			list := ""
			if pl == 1 {
				list += "Removed from Buyer\n"
			} else if pl == 2 {
				list += "Removed from Owner\n"
			} else if pl == 3 {
				list += "Removed from Master\n"
			} else if pl == 4 {
				list += "Expeled from Admin\n"
			} else if pl == 5 {
				list += "Expeled from Bot\n"
			} else if pl == 6 {
				list += "Expeled from Gowner\n"
			} else if pl == 7 {
				list += "Expeled from Gadmin\n"
			} else if pl == 8 {
				list += "Expeled from Access\n"
			} else if pl == 9 {
				list += "Expeled from Maker\n"
			} else if pl == 10 {
				list += "Expeled from Creator\n"
			} else if pl == 17 {
				list += "Expeled from Seller\n"
			}
			for i := range target {
				list += "\n" + strconv.Itoa(i+1) + ". " + "@!"
			}
			client.SendPollMention(to, list, target)
			if pl == 2 {
				LogAccess(client, to, sinder, "unowner", target, 2)
			} else if pl == 3 {
				LogAccess(client, to, sinder, "unmaster", target, 2)
			} else if pl == 4 {
				LogAccess(client, to, sinder, "unadmin", target, 2)
			} else if pl == 5 {
				LogAccess(client, to, sinder, "unbot", target, 2)
			} else if pl == 6 {
				LogAccess(client, to, sinder, "ungowner", target, 2)
			} else if pl == 7 {
				LogAccess(client, to, sinder, "ungadmin", target, 2)
			} else if pl == 8 {
				LogAccess(client, to, sinder, "expel", target, 2)
			}
		} else if conts != 0 {
			list := "Sorry, your grade is too low.\n"
			client.SendMessage(to, botstate.Fancy(list))
		} else if conts2 != 0 {
			list := "Users not have access.\n"
			client.SendMessage(to, botstate.Fancy(list))
		}
	}
}

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


func CancelAll(client *linetcr.Account, mem []string, to string) {
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

func GetFuckV1(client *linetcr.Account, vo string, to string) {
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

func GetFuck(cl *linetcr.Account, vo string, Group string) {
	defer utils.PanicHandle("getfuck")
	runtime.GOMAXPROCS(botstate.Cpu)
	if MemBan(Group, vo) {
		cl.CancelChatInvitations(Group, []string{vo})
	}
}

func CancelAllCek(client *linetcr.Account, mem []string, to string) {
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
			go GetFuck(Cans[no], target, to)
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

func PurgeAct(Group string, cl *linetcr.Account) {
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

func AutoKickBan(client *linetcr.Account, to string, target string) {
	if botstate.AutokickBan {
		gr, _ := client.GetGroupIdsJoined()
		for _, aa := range gr {
			go client.DeleteOtherFromChats(aa, []string{target})
			go client.CancelChatInvitations(aa, []string{target})
		}
	}
}
