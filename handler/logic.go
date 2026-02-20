package handler

import (
	"fmt"
	"sync"
	"../botstate"
	"../config"
	"../library/linetcr"
	"../utils"
)

func CanceljoinBot(Client *linetcr.Account, Group string) {
	defer utils.PanicHandle("CanceljoinBot")
	_, _, pind := Client.GetChatList(Group)
	for _, i := range pind {
		if linetcr.IsPending(Client, Group, i) == true {
			Client.CancelChatInvitations(Group, []string{i})
		}
	}
}

func Nukjoin(Client *linetcr.Account, Optime int64, Group string) {
	defer utils.PanicHandle("Nukejoin")
	_, ok := botstate.Ceknuke.Get(Optime)
	if !ok {
		botstate.Ceknuke.Set(Optime, 1)
	} else {
		return
	}
	exe, list := botstate.CheckUser(Client, Group)
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

func AutopurgeEnd(client *linetcr.Account, Group string, mem []string) {
	defer utils.PanicHandle("AutopurgeEnd")
	for _, target := range mem {
		client.DeleteOtherFromChats(Group, []string{target})
	}
}

func RemoveSticker(items []*config.Stickers, item *config.Stickers) []*config.Stickers {
	defer linetcr.PanicOnly()
	newitems := []*config.Stickers{}
	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}

func GETgrade(num int) string {
	if num == 0 {
		return "Dev"
	} else if num == 1 {
		return "Creator"
	} else if num == 2 {
		return "Maker"
	} else if num == 3 {
		return "seller"
	} else if num == 4 {
		return "Buyer"
	} else if num == 5 {
		return "Owner"
	} else if num == 6 {
		return "Master"
	} else if num == 7 {
		return "Admin"
	} else if num == 8 {
		return "Gowner"
	} else if num == 9 {
		return "Gadmin"
	}
	return "None"
}



func CekOp(optime int64) bool {
	for _, tar := range botstate.Oplist {
		if tar == optime {
			return false
		}
	}
	botstate.Oplist = append(botstate.Oplist, optime)
	return true
}

func CekOpinvite(optime int64) bool {
	for _, tar := range botstate.Oplistinvite {
		if tar == optime {
			return false
		}
	}
	botstate.Oplistinvite = append(botstate.Oplistinvite, optime)
	return true
}

func LlistCheck(client *linetcr.Account, to string, typec string, nCount int, sender string, rplay string, mentionlist []string) (ss []string) {
	saodd := []string{}
	pendlast := []string{}
	if len(mentionlist) != 0 {
		for a := range mentionlist {
			if !utils.InArrayString(saodd, mentionlist[a]) && !utils.InArrayString(botstate.Squadlist, mentionlist[a]) {
				saodd = append(saodd, mentionlist[a])
			}

		}
		return saodd
	} else if rplay != "" {
		if !utils.InArrayString(saodd, rplay) {
			saodd = append(saodd, rplay)
		}
		return saodd
	} else if typec == "rplay" {
		if !utils.InArrayString(saodd, rplay) {
			saodd = append(saodd, rplay)
		}
		return saodd
	} else if typec == "lmid" {
		g, ok := botstate.Lastmid.Get(to)
		if !ok {
			g = [][]string{}
			botstate.Lastmid.Set(to, g)
		} else {
			num := nCount
			c := g.([][]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i]...)
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lmessage" {
		g, ok := botstate.Lastmessage.Get(to)
		if !ok {
			g = [][]string{}
			botstate.Lastmessage.Set(to, g)
		} else {
			num := nCount
			c := g.([][]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i]...)
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "linvite" {
		g, ok := botstate.Lastinvite.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastinvite.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lkick" {
		g, ok := botstate.Lastkick.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastkick.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lcancel" {
		g, ok := botstate.Lastcancel.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastcancel.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lqr" {
		g, ok := botstate.Lastupdate.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastupdate.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "ljoin" {
		g, ok := botstate.Lastjoin.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastjoin.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "ltag" {
		g, ok := botstate.Lasttag.Get(to)
		if !ok {
			g = []string{}
			botstate.Lasttag.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lcon" {
		g, ok := botstate.Lastcon.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastcon.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "lleave" {
		g, ok := botstate.Lastleave.Get(to)
		if !ok {
			g = []string{}
			botstate.Lastleave.Set(to, g)
		} else {
			num := nCount
			c := g.([]string)
			lk := len(c)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, c[i])
					}
					no++
				}
				return saodd
			}
		}
	} else if typec == "@me" {
		if !utils.InArrayString(saodd, sender) {
			saodd = append(saodd, sender)
		}
		return saodd
	} else if typec == "@all" {
		Member := client.GetChatListMem(to)
		for i := 0; i < len(Member); i++ {
			if !utils.InArrayString(saodd, Member[i]) {
				saodd = append(saodd, Member[i])
			}
		}
		return saodd
	} else if typec == "lbanlist" {
		for i := 0; i < len(botstate.Banned.Banlist); i++ {
			if !utils.InArrayString(saodd, botstate.Banned.Banlist[i]) {
				saodd = append(saodd, botstate.Banned.Banlist[i])
			}
		}
		return saodd
	} else if typec == "pend" {
		tcr := strings.Replace(typec, "pend ", "", 1)
		numb, _ := strconv.Atoi(tcr)
		_, _, pind := client.GetChatList(to)
		if numb == 0 {
			for n, i := range pind {
				if n < botstate.CancelPend {
					if !utils.InArrayString(saodd, i) {
						saodd = append(saodd, i)
					}
				}
			}
		}
		return saodd
	} else if typec == "numpend" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
			if !utils.InArrayString(saodd, i) {
				pendlast = append(pendlast, i)
			}
			num := nCount
			lk := len(pendlast)
			if lk != 0 {
				no := 0
				for i := lk - 1; i >= 0; i-- {
					if no < num {
						saodd = append(saodd, pendlast[i])
					}
					no++
				}
				return saodd
			}
		}
		return saodd
	} else if typec == "pendingall" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
			if !utils.InArrayString(saodd, i) {
				saodd = append(saodd, i)
			}
		}
		return saodd
	} else if typec == "@oa" {
		_, _, pind := client.GetChatList(to)
		for _, i := range pind {
		      contact, _ := client.GetContact(i)
			if contact != nil {
				if contact.CapableBuddy {
					if !utils.InArrayString(saodd, i) {
						saodd = append(saodd, i)
					}
				}
			}
		}
		return saodd
	}
	return saodd
}