package handler

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"../botstate"
	"../library/linetcr"
	"../utils"
	"runtime"
)

func Setinvitetomsg(client *linetcr.Account, to string, invits []string) []string {
	defer utils.PanicHandle("Setinvitetomsg")
	bans := []string{}
	news := []string{}
	room := linetcr.GetRoom(to)
	exe := room.HaveClient
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == false && linetcr.IsPending(client, to, cc) == false {
			if !MemBan(to, cc) {
				if linetcr.IsFriends(client, cc) == false {
					client.FindAndAddContactsByMidV2(cc)
					time.Sleep(250 * time.Millisecond)
				}
				news = append(news, cc)
			} else {
				bans = append(bans, cc)

			}
		}
	}
	if len(news) != 0 && len(exe) != 0 {
		celek := len(news)
		no := 0
		bat := 5
		ClAct := len(exe)
		if ClAct != 0 {
			if celek < bat {
				for _, cl := range exe {
					cl.GetRecommendationIds()
					for _, mid := range news {
						linetcr.AddContact3(cl, mid)
					}
					fl, _ := cl.GetAllContactIds()
					bb := []string{}
					for _, mid := range news {
						if utils.InArrayString(fl, mid) {
							bb = append(bb, mid)
							news = utils.RemoveString(news, mid)
						}
					}
					if len(bb) != 0 {
						cl.InviteIntoGroupNormal(to, bb)
					}
					if len(news) == 0 {
						break
					}
				}
			} else {
				hajar := []string{}
				z := celek / bat
				y := z + 1
				for i := 0; i < y; i++ {
					if no >= ClAct {
						no = 0
					}
					client := exe[no]
					if i == z {
						hajar = news[i*bat:]
					} else {
						hajar = news[i*bat : (i+1)*bat]
					}
					if len(hajar) != 0 {
						client.GetRecommendationIds()
						for _, mid := range hajar {
							linetcr.AddContact3(client, mid)
						}
						fl, _ := client.GetAllContactIds()
						bb := []string{}
						for _, mid := range hajar {
							if utils.InArrayString(fl, mid) {
								bb = append(bb, mid)
							}
						}
						if len(bb) != 0 {
							client.InviteIntoGroupNormal(to, bb)
						}
					}
					no += 1
				}
			}
		}
	}
	return bans
}

func Setkickto(client *linetcr.Account, to string, invits []string) {
	defer utils.PanicHandle("Setkickto")
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, []string{cc})
		}
	}

}

func Setcancelto(client *linetcr.Account, to string, invits []string) {
	defer utils.PanicHandle("Setcancelto")
	for _, x := range invits {
		if linetcr.IsPending(client, to, x) == true {
			client.CancelChatInvitations(to, []string{x})
		}
	}
}

func AutojoinQr22(client *linetcr.Account, to string) {
	numb := len(botstate.ClientBot)
	if numb > 0 && numb <= len(botstate.ClientBot) {
		GetSquad(client, to)
		room := linetcr.GetRoom(to)
		aa := len(room.Client)
		if aa > numb {
			c := aa - numb
			ca := 0
			list := append([]*linetcr.Account{}, room.Client...)
			sort.Slice(list, func(i, j int) bool {
				return list[i].KickPoint > list[j].KickPoint
			})
			for _, o := range list {
				o.LeaveGroup(to)
				ca = ca + 1
				if ca == c {
					break
				}
			}
			GetSquad(client, to)
		} else if aa < numb {
			ti, botstate.Err := client.ReissueChatTicket(to)
			if botstate.Err == nil {
				go client.UpdateChatQrV2(to, false)
				all := []*linetcr.Account{}
				room := linetcr.GetRoom(to)
				cuk := room.Client
				for _, x := range botstate.ClientBot {
					if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
						all = append(all, x)
					}
				}
				sort.Slice(all, func(i, j int) bool {
					return all[i].KickPoint < all[j].KickPoint
				})
				g := numb - aa
				var wg sync.WaitGroup
				wi := GetSquad(client, to)
				for i := 0; i < len(all); i++ {
					if i == g {
						break
					}
					l := all[i]
					if l != client && !linetcr.InArrayCl(wi, l) {
						wg.Add(1)
						go func() {
							l.AcceptTicket(to, ti)
							QrKick(client, to)
							wg.Done()
						}()
					}
				}
				wg.Wait()
				client.UpdateChatQrV2(to, true)
				GetSquad(client, to)
			}
		}
	}
}

func AutojoinQr(client *linetcr.Account, to string) {
	defer utils.PanicHandle("AutojoinQr")
	ti, botstate.Err := client.ReissueChatTicket(to)
	if botstate.Err == nil {
		go client.UpdateChatQrV2(to, false)
		all := []*linetcr.Account{}
		room := linetcr.GetRoom(to)
		cuk := room.Client
		for _, x := range botstate.ClientBot {
			if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
				all = append(all, x)
			}
		}
		sort.Slice(all, func(i, j int) bool {
			return all[i].KickPoint < all[j].KickPoint
		})
		var wg sync.WaitGroup
		wi := GetSquad(client, to)
		for i := 0; i < len(all); i++ {
			l := all[i]
			if l != client && !linetcr.InArrayCl(wi, l) {
				wg.Add(1)
				go func() {
					l.AcceptTicket(to, ti)
					wg.Done()
				}()
			}
		}
		wg.Wait()
		client.UpdateChatQrV2(to, true)
		GetSquad(client, to)
	}
}

func QrGo(cl *linetcr.Account, cans []*linetcr.Account, to string) {
	defer utils.PanicHandle("QR_go")
	Room := linetcr.GetRoom(to)
	mes := make(chan bool)
	go func() {
		botstate.Err := cl.UpdateChatQrV2(to, false)
		if botstate.Err != nil {
			mes <- false
		} else {
			mes <- true
		}
	}()
	Room.Qr = false
	var ticket string
	link, botstate.Err := cl.ReissueChatTicket(to)
	if botstate.Err == nil {
		ticket = link
	} else {
		ticket = "error"
	}
	var wg sync.WaitGroup
	if ticket != "error" && ticket != "" {
		ok := <-mes
		if !ok {
			return
		}
		for _, cc := range cans {
			wg.Add(1)
			go func(c *linetcr.Account) {
				botstate.Err := c.AcceptTicket(to, ticket)
				if botstate.Err != nil {
					fmt.Println(botstate.Err)
				}
				wg.Done()
			}(cc)
		}
		wg.Wait()
		Room.Qr = true
	}
	if Room.Qr {
		go func() {
			botstate.Err := cl.UpdateChatQrV2(to, true)
			if botstate.Err != nil {
				mes <- true
			} else {
				mes <- false
			}
		}()
	}
}

func AcceptTicketSimple(client *linetcr.Account, to string, ticketId string){
    runtime.GOMAXPROCS(botstate.Cpu)
    go func(){
        client.AcceptTicket(to, ticketId)
    }()
    go func(){
        var wg sync.WaitGroup
        wg.Add(len(botstate.Banned.Banlist))
        for i:=0; i<len(botstate.Banned.Banlist); i++ {
            go func(i int){
                defer wg.Done()
                if linetcr.IsMembers(client, to, botstate.Banned.Banlist[i]) {
                    client.DeleteOtherFromChats(to, []string{botstate.Banned.Banlist[i]})
                }
            }(i)
        }
        wg.Wait()
    }()
    time.Sleep(500 * time.Nanosecond)            
}

func WarQr(client *linetcr.Account, to string){
	runtime.GOMAXPROCS(botstate.Cpu)
	GetSquad(client, to)
    chat := client.GetChat([]string{to}, true, false)
    if chat != nil {
        cek := chat.Chats[0].Extra.GroupExtra.PreventedJoinByTicket
        if cek == true{go client.UpdateChatQrV2(to, false);time.Sleep(100 * time.Millisecond)}	
        ticket, botstate.Err := client.ReissueChatTickets(to)
        if botstate.Err == nil {
            link := ticket
            all := []*linetcr.Account{}
            room := linetcr.GetRoom(to)
            cuk := room.Client
            for _, x := range botstate.ClientBot {
                if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
                    all = append(all, x)}}
            sort.Slice(all, func(i, j int) bool {return all[i].KickPoint < all[j].KickPoint})
            var wg sync.WaitGroup
            wi := GetSquad(client, to)
            for i := 0; i < len(all); i++ {
                l := all[i]
                if l != client && !linetcr.InArrayCl(wi, l) {
                    wg.Add(1)
                    go func() {
                        SpamAcceptQR(l, to, link)
                        wg.Done()
                    }()
                }
            }
            wg.Wait()  
        }
        runtime.GOMAXPROCS(1)
    }
}

func SpamAcceptQR(client *linetcr.Account, to string, ticketId string) {
	runtime.GOMAXPROCS(100)
	go func(){AcceptTicketSimple(client, to, ticketId)}()
	go func(){WarQr(client, to)}()
	time.Sleep(200 * time.Nanosecond)
}

// Hstg moved to helpers.go

func QrGo22(client *linetcr.Account, cans []*linetcr.Account, to string) {
	numb := len(botstate.ClientBot)
	if numb > 0 && numb <= len(botstate.ClientBot) {
		GetSquad(client, to)
		room := linetcr.GetRoom(to)
		aa := len(room.Client)
		if aa > numb {
			c := aa - numb
			ca := 0
			list := append([]*linetcr.Account{}, room.Client...)
			sort.Slice(list, func(i, j int) bool {
				return list[i].KickPoint > list[j].KickPoint
			})
			for _, o := range list {
				o.LeaveGroup(to)
				ca = ca + 1
				if ca == c {
					break
				}
			}
			GetSquad(client, to)
		} else if aa < numb {
			ti, botstate.Err := client.ReissueChatTicket(to)
			if botstate.Err == nil {
				go client.UpdateChatQrV2(to, false)
				all := []*linetcr.Account{}
				room := linetcr.GetRoom(to)
				cuk := room.Client
				for _, x := range botstate.ClientBot {
					if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
						all = append(all, x)
					}
				}
				sort.Slice(all, func(i, j int) bool {
					return all[i].KickPoint < all[j].KickPoint
				})
				g := numb - aa
				var wg sync.WaitGroup
				wi := GetSquad(client, to)
				for i := 0; i < len(all); i++ {
					if i == g {
						break
					}
					l := all[i]
					if l != client && !linetcr.InArrayCl(wi, l) {
						wg.Add(1)
						go func() {
							l.AcceptTicket(to, ti)
							QrKick(client, to)
							wg.Done()
						}()
					}
				}
				wg.Wait()
				client.UpdateChatQrV2(to, true)
				GetSquad(client, to)
			}
		}
	}
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
		NukeAll(client, Group)
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
			cl := GetKorban(mid)
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
			NukeAll(client, Group)
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

func CancelJoinBot(Client *linetcr.Account, Group string) {
	defer utils.PanicHandle("CanceljoinBot")
	_, _, pind := Client.GetChatList(Group)
	for _, i := range pind {
		if linetcr.IsPending(Client, Group, i) == true {
			Client.CancelChatInvitations(Group, []string{i})
		}
	}
}

func NukJoin(Client *linetcr.Account, Optime int64, Group string) {
	defer utils.PanicHandle("Nukejoin")
	_, ok := botstate.Ceknuke.Get(Optime)
	if !ok {
		botstate.Ceknuke.Set(Optime, 1)
	} else {
		return
	}
	exe, list := Checkuser(Client, Group)
	if exe != nil {
		no := 0
		i := 0
		lm := len(list)
		acts := []*linetcr.Account{}
		var cl *linetcr.Account
		for ; i < lm; i++ {
			if no >= len(exe) {
				no = 0
			}
			acts = append(acts, exe[no])
			no += 1
		}
		for n, target := range list {
			go func(n int, target string) {
				cl = acts[n]
				cl.DeleteOtherFromChats(Group, []string{target})
			}(n, target)
		}
		_, _, pind := Client.GetChatList(Group)
		for _, p := range pind {
			if MemUser(Group, p) {
				if linetcr.IsPending(Client, Group, p) == true {
					Client.CancelChatInvitations(Group, []string{p})
				}
			}
		}
	}
}

func AutoPurgeEnd(client *linetcr.Account, Group string, mem []string) {
	defer utils.PanicHandle("AutopurgeEnd")
	for _, target := range mem {
		client.DeleteOtherFromChats(Group, []string{target})
	}
}

func CancelEnd(client *linetcr.Account, Group string, mem []string) {
	defer utils.PanicHandle("CancelEnd")
	for _, target := range mem {
		client.CancelChatInvitations(Group, []string{target})
	}
}

func SetPurgeAllN(client *linetcr.Account, to string, invits []string) {
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == true {
			client.DeleteOtherFromChats(to, []string{cc})
		} else if linetcr.IsPending(client, to, cc) == true {
			client.CancelChatInvitations(to, []string{cc})
		}
	}

}

func SetInviteTo(client *linetcr.Account, to string, invits []string) {
	news := []string{}
	for _, cc := range invits {
		if linetcr.IsMembers(client, to, cc) == false && linetcr.IsPending(client, to, cc) == false {
			news = append(news, cc)
		}
	}
	if len(news) != 0 {
		client.InviteIntoChatPollVer(to, news)
	}
}

func InvBackup(exe *linetcr.Account, to string, oke []string, korban string) {
	exe.InviteIntoGroupNormal(to, []string{korban})
}

func OpenQR(exe []*linetcr.Account, to string, mes chan bool) {
	defer utils.PanicHandle("QR_backupupdate")
	Room := linetcr.GetRoom(to)
	Room.Qr = false
	for _, cl := range exe {
		if botstate.Err == nil {
			mes <- true
			return
		}
	}
	mes <- false
}

func GetTicket(exe []*linetcr.Account, to string, lnk chan string) {
	defer utils.PanicHandle("gettiket")
	ClAct := len(exe)
	if ClAct > 1 {
		for i := ClAct - 1; i >= 0; i-- {
			cls := exe[i]
			link, botstate.Err := cls.ReissueChatTicket(to)
			if botstate.Err == nil {
				lnk <- link
				return
			}
		}
		lnk <- "error"
		return

	} else {
		link, botstate.Err := exe[0].ReissueChatTicket(to)
		if botstate.Err == nil {
			lnk <- link
		} else {
			lnk <- "error"
		}
		return
	}
}

func QrBackup(exe []*linetcr.Account, to string, oke []string) {
	defer utils.PanicHandle("qrBackup")
	lnk := make(chan string)
	Room := linetcr.GetRoom(to)
	mes := make(chan bool)
	go OpenQR(exe, to, mes)
	go GetTicket(exe, to, lnk)
	bot := linetcr.GetRoom(to).Bot
	cans := []*linetcr.Account{}
	for _, mid := range bot {
		if !utils.InArrayString(oke, mid) {
			cl := GetKorban(mid)
			if cl.Limited {
				cans = append(cans, cl)
			}
		}
	}
	var wg sync.WaitGroup
	linetcr.GetRoom(to).Purge = false
	ticket := <-lnk
	if len(ticket) > 5 {
		ok := <-mes
		if ok {
			for _, cc := range cans {
				wg.Add(1)
				go func(c *linetcr.Account) {
					c.AcceptTicket(to, ticket)
					wg.Done()
				}(cc)
			}
			wg.Wait()
			Room.Qr = true
		} else {
			Room.Qr = true
		}

	} else {
		Room.Qr = true
	}
}

func GroupBackupInv2(client *linetcr.Account, to string) {
	defer utils.PanicHandle("groupBackupInv2")
	all := []string{}
	GetSquad(client, to)
	room := linetcr.GetRoom(to)
	cuk := room.Client
	exe := []*linetcr.Account{}
	for _, x := range botstate.ClientBot {
		if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
			exe = append(exe, x)
		}
	}
	for _, x := range botstate.ClientBot {
		if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
			all = append(all, x.MID)
		} else {
			break
		}
	}
	ClAct := len(exe)
	if ClAct != 0 {
		sort.Slice(exe, func(i, j int) bool {
			return exe[i].KickPoint < exe[j].KickPoint
		})
		for _, a := range all {
			exe[0].InviteIntoGroupNormal(to, []string{a})
	            	time.Sleep(250 * time.Millisecond)
			GetSquad(client, to)
		}
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}

func GroupBackupInv(client *linetcr.Account, to string, optime int64, korban string) {
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
			InvBackup(exe[0], to, oke, korban)
		} else if botstate.ModeBackup == "qr" {
			QrBackup(exe, to, oke)
		}
		linetcr.SetAva(to, oke)
	}
	runtime.GOMAXPROCS(botstate.Cpu)
}