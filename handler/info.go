package handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"../botstate"
	"../library/SyncService"
	"../library/linetcr"
	"../utils"
)

func InfoCreator(client *linetcr.Account) string {
	list := ""
	if len(botstate.CREATOR) != 0 {
		cuh, _ := client.GetContacts(botstate.CREATOR)
		for _, prs := range cuh {
			name := prs.DisplayName
			list += fmt.Sprintf("\n ⚙️ Developer : %v", name)
		}
	}
	return list
}

func CheckExprd(s *linetcr.Account, to string, sender string) bool {
	base := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)
	d := fmt.Sprintf("%v", botstate.Data.Dalltime)
	has := strings.Split(d, "-")
	has2 := strings.Split(has[2], "T")
	yy, _ := strconv.Atoi(has[0])
	mm, _ := strconv.Atoi(has[1])
	timeup, _ := strconv.Atoi(has2[0])
	batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
	if batas.Before(base) {
		if !SendMycreator(sender) {
			s.SendMessage(to, botstate.Fancy("Sorry your bots is expired, Please Contact with our Creator to renew your squad. ;-)"))
			return false
		}
		return true
	}
	return true
}

func CekDuedate() time.Time {
	bod := string(botstate.Data.Dalltime)
	date, _ := time.Parse(time.RFC3339, bod)
	return date
}

func CheckLastActive(client *linetcr.Account, targets string) string {
	list := ""
	mek, tu := botstate.LastActive.Get(targets)
	if tu {
		asu := mek.(*SyncService.Operation)
		if asu.Type == 55 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := FmtDurations(ss)
			list += "- botstate.LastActive: " + sp + "\n- Type: Read Message\n- Group: " + names1 + "\n\n"
		} else if asu.Type == 124 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := FmtDurations(ss)
			invites := strings.Split(asu.Param3, "\x1e")
			nos := 0
			her := ""
			for _, ampemng := range invites {
				nos += 1
				pr, _ := client.GetContact(ampemng)
				her += fmt.Sprintf("\n  %v. %v", nos, pr.DisplayName)
			}
			list += "- botstate.LastActive: " + sp + "\n- Type: Invited member\n- Group: " + names1 + "\n- Target: " + her + "\n\n"
		} else if asu.Type == 133 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := FmtDurations(ss)
			pr, _ := client.GetContact(asu.Param3)
			list += "- botstate.LastActive: " + sp + "\n- Type : Kick member\n- Group: " + names1 + "\n- Target: " + pr.DisplayName + "\n\n"
		} else if asu.Type == 126 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := FmtDurations(ss)
			pr, _ := client.GetContact(asu.Param3)
			list += "- botstate.LastActive: " + sp + "\n- Type: Cancel member\n- Group: " + names1 + "\n- Target: " + pr.DisplayName + "\n\n"
		} else if asu.Type == 26 {
			msg := asu.Message
			if msg.ToType == 2 {
				names1, _ := client.GetGroupMember(msg.To)
				cok := asu.CreatedTime / 1000
				i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
				tm := time.Unix(i, 0)
				ss := time.Since(tm)
				sp := FmtDurations(ss)
				tx := ""
				if msg.ContentType == 0 {
					tx = msg.Text
				} else {
					tx = "Non Text Message"
				}
				list += "- botstate.LastActive: " + sp + "\n- Type: Send Message\n- Group: " + names1 + "\n- Message: " + tx + "\n\n"
			}
		} else if asu.Type == 130 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := FmtDurations(ss)
			list += "- botstate.LastActive: " + sp + "\n- Type: Join Group\n- Group: " + names1 + "\n\n"
		} else if asu.Type == 122 {
			names1, _ := client.GetGroupMember(asu.Param1)
			cok := asu.CreatedTime / 1000
			i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			tm := time.Unix(i, 0)
			ss := time.Since(tm)
			sp := FmtDurations(ss)
			var ti string
			if asu.Param3 == "4" {
				g, _ := client.GetGroup3(asu.Param1)
				if g.Extra.GroupExtra.PreventedJoinByTicket == false {
					ti = "Open qr"
				} else {
					ti = "Close qr"
				}
			} else if asu.Param3 == "1" {
				ti = "Change Group Name"
			}
			list += "- botstate.LastActive: " + sp + "\n- Type: Update Group\n- Group: " + names1 + "\n- Type: " + ti + "\n\n"
		}
	}
	return list
}

func Gtotal(client *linetcr.Account, gid string) string {
	list := ""
	GetSquad(client, gid)
	Room := linetcr.GetRoom(gid)
	_, mem, pending := client.GetChatList(gid)
	Glist := []string{}
	mGlist := []string{}
	for _, from := range mem {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !utils.InArrayString(Glist, from) {
				Glist = append(Glist, from)
			}
		}
	}
	for _, from := range pending {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			if !utils.InArrayString(mGlist, from) {
				mGlist = append(mGlist, from)
			}
		}
	}
	list += fmt.Sprintf("Gtotal: %s", Room.Name)
	pp := len(mGlist)
	list += "\n Member: 38"
	list += fmt.Sprintf("\n Pending: %v", pp)
	list += "\n Total: 38"
	list += "\n Freeinvite: 462"
	return list
}

func InfoGroup(client *linetcr.Account, gid string) string {
	list := ""
	GetSquad(client, gid)
	Room := linetcr.GetRoom(gid)
	_, mem, pending := client.GetChatList(gid)
	Developer := []string{}
	creator := []string{}
	buyer := []string{}
	owner := []string{}
	master := []string{}
	admin := []string{}
	gowner := []string{}
	gadmin := []string{}
	squad := []string{}
	bot := []string{}
	ban := []string{}
	fuck := []string{}
	mute := []string{}
	Gban := []string{}
	Glist := []string{}
	Maker := []string{}
	Seller := []string{}
	mGlist := []string{}

	processMember := func(from string, isPending bool) {
		if MemUser(gid, from) && !MemBan2(gid, from) {
			targetList := &Glist
			if isPending {
				targetList = &mGlist
			}
			if !utils.InArrayString(*targetList, from) {
				*targetList = append(*targetList, from)
			}
		} else if botstate.UserBot.GetCreator(from) {
			creator = append(creator, from)
		} else if botstate.UserBot.GetMaker(from) {
			Maker = append(Maker, from)
		} else if botstate.UserBot.GetSeller(from) {
			Seller = append(Seller, from)
		} else if utils.InArrayString(botstate.DEVELOPER, from) {
			Developer = append(Developer, from)
		} else if botstate.UserBot.GetBuyer(from) {
			buyer = append(buyer, from)
		} else if botstate.UserBot.GetOwner(from) {
			owner = append(owner, from)
		} else if botstate.UserBot.GetMaster(from) {
			master = append(master, from)
		} else if botstate.UserBot.GetAdmin(from) {
			admin = append(admin, from)
		} else if utils.InArrayString(Room.Gowner, from) {
			gowner = append(gowner, from)
		} else if utils.InArrayString(Room.Gadmin, from) {
			gadmin = append(gadmin, from)
		} else if botstate.UserBot.GetBot(from) {
			bot = append(bot, from)
		} else if botstate.Banned.GetFuck(from) {
			fuck = append(fuck, from)
		} else if botstate.Banned.GetBan(from) {
			ban = append(ban, from)
		} else if botstate.Banned.GetMute(from) {
			mute = append(mute, from)
		} else if utils.InArrayString(Room.Gban, from) {
			Gban = append(Gban, from)
		} else if utils.InArrayString(botstate.Squadlist, from) {
			squad = append(squad, from)
		}
	}

	for _, from := range mem {
		processMember(from, false)
	}
	for _, from := range pending {
		processMember(from, true)
	}
	list += fmt.Sprintf("Group Info: %s", Room.Name)
	if len(Glist) != 0 {
		list += "\n\nMember: \n"
		cuh, _ := client.GetContacts(Glist)
		for _, prs := range cuh {
			name := prs.DisplayName
			list += fmt.Sprintf("\n   %s", name)
		}
	}
	if len(mGlist) != 0 {
		chp, _ := client.GetContacts(mGlist)
		list += "\n\n Pending: \n"
		for _, prs := range chp {
			name := prs.DisplayName
			list += fmt.Sprintf("\n   %s", name)
		}
	}

	appendMembers := func(title string, mids []string) {
		if len(mids) != 0 {
			list += "\n" + title + "\n"
			for n, xx := range mids {
				rengs := strconv.Itoa(n + 1)
				name := GetContactName(client, xx)
				list += fmt.Sprintf("%s. %s\n", rengs, name)
			}
		}
	}

	if len(Glist)+len(mGlist) != len(pending)+len(mem) {
		list += "\n\nUsers have access:\n"
		appendMembers("\nExist in Developer:", Developer)
		appendMembers("\nExist in Creator:", creator)
		appendMembers("\nExist in Maker:", Maker)
		appendMembers("\nExist in Seller:", Seller)
		appendMembers("\nExist in Buyers:", buyer)
		appendMembers("\nExist in Owners:", owner)
		appendMembers("\nExist in Masters:", master)
		appendMembers("\nExist in Admins:", admin)
		appendMembers("\nExist in Gowners:", gowner)
		appendMembers("\nExist in Gadmins:", gadmin)
		appendMembers("\nExist in Botlist", bot)
		appendMembers("\nExist in Squad:", squad)
		appendMembers("Exist in Banlist:", ban)
		appendMembers("\nExist in Fucklist:", fuck)
		appendMembers("\nExist in Gbanlist:", Gban)
		appendMembers("\nExist in Mutelist:", mute)
	}
	return list
}

func CheckListAccess(client *linetcr.Account, group string, targets []string, pl int, sinder string) {
	Room := linetcr.GetRoom(group)
	if pl == 12 {
		countr := 0
		countr1 := 0
		list := "Account Info: \n\n"
		for n, xx := range targets {
			x, botstate.Err := client.GetContact(xx)
			if botstate.Err != nil {
				list += "Name: Closed Account \n"
			} else {
				list += fmt.Sprintf("Name: %v \n", x.DisplayName)
				status := "status: None\n\n"
				if utils.InArrayString(botstate.DEVELOPER, targets[n]) {
					status = "status: Developer\n\n"
				} else if botstate.UserBot.GetCreator(targets[n]) {
					status = "status: Creators\n\n"
				} else if botstate.UserBot.GetMaker(targets[n]) {
					status = "status: Makers\n\n"
				} else if botstate.UserBot.GetBuyer(targets[n]) {
					status = "status: Buyer\n\n"
				} else if botstate.UserBot.GetOwner(targets[n]) {
					status = "status: Owner\n\n"
				} else if botstate.UserBot.GetMaster(targets[n]) {
					status = "status: Master\n\n"
				} else if botstate.UserBot.GetAdmin(targets[n]) {
					status = "status: Admin\n\n"
				} else if utils.InArrayString(Room.Gowner, targets[n]) {
					status = "status: GroupOwnar\n\n"
				} else if utils.InArrayString(Room.Gadmin, targets[n]) {
					status = "status: GroupAdmin\n\n"
				} else if botstate.UserBot.GetBot(targets[n]) {
					status = "status: Bot\n\n"
				} else if botstate.Banned.GetFuck(targets[n]) {
					status = "status: Fuck\n\n"
				} else if botstate.Banned.GetBan(targets[n]) {
					status = "status: Ban\n\n"
				} else if botstate.Banned.GetMute(targets[n]) {
					status = "status: Mute\n\n"
				} else if utils.InArrayString(Room.Gban, targets[n]) {
					status = "status: Groupban\n\n"
				} else if utils.InArrayString(botstate.Squadlist, targets[n]) {
					status = "status: My team\n\n"
				} else if botstate.UserBot.GetSeller(targets[n]) {
					status = "status: My Seller\n\n"
				}
				list += status
				if !utils.InArrayString(botstate.CheckHaid, targets[n]) {
					new := CheckLastActive(client, targets[n])
					list += new
				}
				listGroup := "\nMember of:\n"
				listPinde := "\nPending of:\n"
				grs, _ := client.GetGroupIdsJoined()
				groups, _ := client.GetGroups(grs)
				for _, x := range groups {
					if linetcr.IsMembers(client, x.ChatMid, targets[n]) == true {
						countr = countr + 1
						nm, _, _ := client.GetChatList(x.ChatMid)
						listGroup += nm + "\n"
					}
					if linetcr.IsPending(client, x.ChatMid, targets[n]) == true {
						countr1 = countr1 + 1
						nm, _, _ := client.GetChatList(x.ChatMid)
						listPinde += nm + "\n"
					}
				}
				if countr != 0 {
					list += fmt.Sprintf("Groups: %v\n", countr)

				} else {
					list += "Groups: 0\n"
				}
				if countr1 != 0 {
					list += fmt.Sprintf("Pendings: %v\n", countr1)
				} else {
					list += "Pendings: 0\n"
				}
				if countr != 0 {
					if !utils.InArrayString(botstate.CheckHaid, targets[n]) {
						list += listGroup
					}
				}
				if countr1 != 0 {
					if !utils.InArrayString(botstate.CheckHaid, targets[n]) {
						list += listPinde
					}
				}

			}
		}
		client.SendMessage(group, botstate.Fancy(list))
	} else if pl == 16 {
		list := ""
		for n, xx := range targets {
			rengs := strconv.Itoa(n + 1)
			x, botstate.Err := client.GetContact(xx)
			if botstate.Err != nil {
				list += rengs + ". Closed Account \n"
			} else {
				list += fmt.Sprintf("%v. %v \n", rengs, x.DisplayName)
			}
		}
		client.SendMessage(group, botstate.Fancy(list))
	} else if pl == 14 {
		list := ""
		for n, xx := range targets {
			rengs := strconv.Itoa(n + 1)
			x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
				list += rengs + ". Closed Account \n"
			} else {
				list += fmt.Sprintf("%v. %v\n_%v\n", n+1, x.DisplayName, targets[n])

			}
		}
		client.SendMessage(group, botstate.Fancy(list))
	} else {
		if len(targets) > 1 {
			Developer := []string{}
			creator := []string{}
			buyer := []string{}
			owner := []string{}
			master := []string{}
			admin := []string{}
			gowner := []string{}
			gadmin := []string{}
			squad := []string{}
			bot := []string{}
			ban := []string{}
			fuck := []string{}
			mute := []string{}
			Gban := []string{}
			Glist := []string{}
			Maker := []string{}
			Seller := []string{}
			for _, from := range targets {
				if MemUser(group, from) && !MemBan2(group, from) {
					if !utils.InArrayString(Glist, from) {
						Glist = append(Glist, from)
					}
				} else if botstate.UserBot.GetCreator(from) {
					creator = append(creator, from)
				} else if botstate.UserBot.GetMaker(from) {
					Maker = append(Maker, from)
				} else if botstate.UserBot.GetSeller(from) {
					Seller = append(Seller, from)
				} else if utils.InArrayString(botstate.DEVELOPER, from) {
					Developer = append(Developer, from)
				} else if botstate.UserBot.GetBuyer(from) {
					buyer = append(buyer, from)
				} else if botstate.UserBot.GetOwner(from) {
					owner = append(owner, from)
				} else if botstate.UserBot.GetMaster(from) {
					master = append(master, from)
				} else if botstate.UserBot.GetAdmin(from) {
					admin = append(admin, from)
				} else if utils.InArrayString(Room.Gowner, from) {
					gowner = append(gowner, from)
				} else if utils.InArrayString(Room.Gadmin, from) {
					gadmin = append(gadmin, from)
				} else if botstate.UserBot.GetBot(from) {
					bot = append(bot, from)
				} else if botstate.Banned.GetFuck(from) {
					fuck = append(fuck, from)
				} else if botstate.Banned.GetBan(from) {
					ban = append(ban, from)
				} else if botstate.Banned.GetMute(from) {
					mute = append(mute, from)
				} else if utils.InArrayString(Room.Gban, from) {
					Gban = append(Gban, from)
				} else if utils.InArrayString(botstate.Squadlist, from) {
					squad = append(squad, from)
				}
			}
			list2 := ""
			if len(Glist) != 0 {
				if pl == 1 {
					list2 += "Promoted as Buyer:\n\n"
				} else if pl == 2 {
					list2 += "Promoted as Owner:\n\n"
				} else if pl == 3 {
					list2 += "Promoted as Master:\n\n"
				} else if pl == 4 {
					list2 += "Promoted as Admin:\n\n"
				} else if pl == 5 {
					list2 += "Promoted as Bot:\n\n"
				} else if pl == 6 {
					list2 += "Promoted as Gowner:\n\n"
				} else if pl == 7 {
					list2 += "Promoted as Gadmin\n\n"
				} else if pl == 8 {
					list2 += "Added to banlist:\n\n"
				} else if pl == 9 {
					list2 += "Added to fucklist:\n\n"
				} else if pl == 10 {
					list2 += "Added to gbanlist:\n\n"
				} else if pl == 11 {
					list2 += "Added to mutelist:\n\n"
				} else if pl == 13 {
					list2 += "Added to Makerlist:\n\n"
				} else if pl == 15 {
					list2 += "Added to Creatorlist:\n\n"
				} else if pl == 17 {
					list2 += "Added to Sellerlist:\n\n"
				} else if pl == 18 {
					list2 += "Added to Friendlist:\n\n"
				}
				for n, xx := range Glist {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list2 += rengs + ". Closed Account \n"
					} else {
						list2 += rengs + ". " + x.DisplayName + "\n"
						if pl == 1 {
							botstate.UserBot.AddBuyer(xx)
						} else if pl == 2 {
							botstate.UserBot.AddOwner(xx)
						} else if pl == 13 {
							botstate.UserBot.AddMaker(xx)
						} else if pl == 15 {
							botstate.UserBot.AddCreator(xx)
						} else if pl == 3 {
							botstate.UserBot.AddMaster(xx)
						} else if pl == 4 {
							botstate.UserBot.AddAdmin(xx)
						} else if pl == 5 {
							botstate.UserBot.AddBot(xx)
						} else if pl == 6 {
							Room.Gowner = append(Room.Gowner, xx)
						} else if pl == 7 {
							Room.Gadmin = append(Room.Gadmin, xx)
						} else if pl == 8 {
							botstate.Banned.AddBan(xx)
						} else if pl == 9 {
							botstate.Banned.AddBan(xx)
						} else if pl == 10 {
							Addgban(xx, group)
						} else if pl == 11 {
							botstate.Banned.AddBan(xx)
						} else if pl == 17 {
							botstate.UserBot.AddSeller(xx)
						}
					}
				}
				if pl == 2 {
					LogAccess(client, group, sinder, "owner", Glist, 2)
				} else if pl == 3 {
					LogAccess(client, group, sinder, "master", Glist, 2)
				} else if pl == 4 {
					LogAccess(client, group, sinder, "admin", Glist, 2)
				} else if pl == 5 {
					LogAccess(client, group, sinder, "bot", Glist, 2)
				} else if pl == 6 {
					LogAccess(client, group, sinder, "gowner", Glist, 2)
				} else if pl == 7 {
					LogAccess(client, group, sinder, "gadmin", Glist, 2)
				} else if pl == 8 {
					LogAccess(client, group, sinder, "ban", Glist, 2)
				} else if pl == 9 {
					LogAccess(client, group, sinder, "fuck", Glist, 2)
				} else if pl == 10 {
					LogAccess(client, group, sinder, "gban", Glist, 2)
				} else if pl == 11 {
					LogAccess(client, group, sinder, "mute", Glist, 2)
				}
			}
			list := "Users have access:\n"
			if len(Developer) != 0 {
				list += "\nExist in Developers:\n"
				for n, xx := range Developer {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(creator) != 0 {
				list += "\nExist in Creators:\n"
				for n, xx := range creator {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Maker) != 0 {
				list += "\nExist in Makers:\n"
				for n, xx := range Maker {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Seller) != 0 {
				list += "\nExist in Sellers:\n"
				for n, xx := range Seller {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(buyer) != 0 {
				list += "\nExist in Buyers:\n"
				for n, xx := range buyer {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(owner) != 0 {
				list += "\nExist in Owners:\n"
				for n, xx := range owner {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(master) != 0 {
				list += "\nExist in Masters:\n"
				for n, xx := range master {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(admin) != 0 {
				list += "\nExist in Admins:\n"
				for n, xx := range admin {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(gowner) != 0 {
				list += "\nExist in Gowners:\n"
				for n, xx := range gowner {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(gadmin) != 0 {
				list += "\nExist in Gadmins:\n"
				for n, xx := range gadmin {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(bot) != 0 {
				list += "\nExist in Botlist\n"
				for n, xx := range bot {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(squad) != 0 {
				list += "\nExist in Squads:\n"
				for n, xx := range squad {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(ban) != 0 {
				list += "Exist in Banlist:\n"
				for n, xx := range ban {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(fuck) != 0 {
				list += "\nExist in Fucklist:\n"
				for n, xx := range fuck {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(Gban) != 0 {
				list += "\nExist in Gbanlist:\n\n"
				for n, xx := range Gban {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if len(mute) != 0 {
				list += "\nExist in Mutelist:\n\n"
				for n, xx := range mute {
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
						list += rengs + ". Closed Account \n"
					} else {
						list += rengs + ". " + x.DisplayName + "\n"
					}
				}
			}
			if list != "Users have access:\n" {
				if list2 != "" {
					list2 += "\n"
				}
				client.SendMessage(group, botstate.Fancy(list2+list))
			} else {
				client.SendMessage(group, botstate.Fancy(list2))
			}
		} else {
			list := ""
			for n, from := range targets {
				if utils.InArrayString(botstate.DEVELOPER, from) {
					list += "User have access exist in Developer."
				} else if botstate.UserBot.GetCreator(from) {
					list += "User have access exist in Creator list."
				} else if botstate.UserBot.GetMaker(from) {
					list += "User have access exist in Maker list."
				} else if botstate.UserBot.GetSeller(from) {
					list += "User have access exist in seller list."
				} else if botstate.UserBot.GetBuyer(from) {
					list += "User have access exist in buyer list."
				} else if botstate.UserBot.GetOwner(from) {
					list += "User have access exist in owner list."
				} else if botstate.UserBot.GetMaster(from) {
					list += "User have access exist in master list."
				} else if botstate.UserBot.GetAdmin(from) {
					list += "User have access exist in admin list."
				} else if utils.InArrayString(Room.Gowner, from) {
					list += "User have access exist in gowner list."
				} else if utils.InArrayString(Room.Gadmin, from) {
					list += "User have access exist in gadmin list."
				} else if botstate.UserBot.GetBot(from) {
					list += "User have access exist in bot list."
				} else if botstate.Banned.GetFuck(from) {
					list += "User have access exist in fuck list."
				} else if botstate.Banned.GetBan(from) {
					list += "User have access exist in ban list."
				} else if utils.InArrayString(Room.Gban, from) {
					list += "User have access exist in gban list."
				} else if utils.InArrayString(botstate.Squadlist, from) {
					list += "User have access exist in squad list."
				} else if botstate.Banned.GetMute(from) {
					list += "User have access exist in mute list."
				} else if MemUser(group, from) && !MemBan2(group, from) {
					if pl == 1 {
						list += "Promoted as Buyer:\n"
					} else if pl == 2 {
						list += "Promoted as Owner:\n"
					} else if pl == 3 {
						list += "Promoted as Master:\n"
					} else if pl == 4 {
						list += "Promoted as Admin:\n"
					} else if pl == 5 {
						list += "Promoted as Bot:\n"
					} else if pl == 6 {
						list += "Promoted as Gowner:\n"
					} else if pl == 7 {
						list += "Promoted as Gadmin:\n"
					} else if pl == 8 {
						list += "Added to banlist:\n"
					} else if pl == 9 {
						list += "Added to fucklist:\n"
					} else if pl == 10 {
						list += "Added to gbanlist:\n"
					} else if pl == 11 {
						list += "Added to mutelist:\n"
					} else if pl == 13 {
						list += "Added to Makerlist:\n"
					} else if pl == 15 {
						list += "Added to Creatorlist:\n"
					} else if pl == 17 {
						list += "Added to Sellerlist:\n"
					} else if pl == 18 {
						list += "Added to Friendlist:\n"
					}
					rengs := strconv.Itoa(n + 1)
					x, botstate.Err := client.GetContact(from)if botstate.Err != nil {
						list += "\n   " + rengs + ". Closed Account"
					} else {
						list += "\n   " + rengs + ". " + x.DisplayName
						if pl == 1 {
							botstate.UserBot.AddBuyer(from)
						} else if pl == 2 {
							botstate.UserBot.AddOwner(from)
						} else if pl == 3 {
							botstate.UserBot.AddMaster(from)
						} else if pl == 4 {
							botstate.UserBot.AddAdmin(from)
						} else if pl == 5 {
							botstate.UserBot.AddBot(from)
						} else if pl == 6 {
							Room.Gowner = append(Room.Gowner, from)
						} else if pl == 13 {
							botstate.UserBot.AddMaker(from)
						} else if pl == 15 {
							botstate.UserBot.AddCreator(from)
						} else if pl == 7 {
							Room.Gadmin = append(Room.Gadmin, from)
						} else if pl == 8 {
							Autokickban(client, group, from)
							botstate.Banned.AddBan(from)
						} else if pl == 9 {
							botstate.Banned.AddFuck(from)
						} else if pl == 10 {
							Addgban(from, group)
						} else if pl == 11 {
							botstate.Banned.AddMute(from)
						} else if pl == 17 {
							botstate.UserBot.AddSeller(from)
						}
					}
					if pl == 2 {
						LogAccess(client, group, sinder, "owner", []string{from}, 2)
					} else if pl == 3 {
						LogAccess(client, group, sinder, "master", []string{from}, 2)
					} else if pl == 4 {
						LogAccess(client, group, sinder, "admin", []string{from}, 2)
					} else if pl == 5 {
						LogAccess(client, group, sinder, "bot", []string{from}, 2)
					} else if pl == 6 {
						LogAccess(client, group, sinder, "gowner", []string{from}, 2)
					} else if pl == 7 {
						LogAccess(client, group, sinder, "gadmin", []string{from}, 2)
					} else if pl == 8 {
						LogAccess(client, group, sinder, "ban", []string{from}, 2)
					} else if pl == 9 {
						LogAccess(client, group, sinder, "fuck", []string{from}, 2)
					} else if pl == 10 {
						LogAccess(client, group, sinder, "gban", []string{from}, 2)
					} else if pl == 11 {
						LogAccess(client, group, sinder, "mute", []string{from}, 2)
					}
				}

			}
			client.SendMessage(group, botstate.Fancy(list))
		}
	}
}

func AllBanList(self *linetcr.Account) string {
	listadm := "𝗔𝗹𝗹 𝗯𝗮𝗻𝗹𝗶𝘀𝘁𝘀:\n"
	if len(botstate.Banned.Banlist) != 0 {
		//listadm += "\n\n ☠️ 𝗕𝗮𝗻𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range botstate.Banned.Banlist {
			num++
			rengs := strconv.Itoa(num)
			x, botstate.Err := self.GetContact(xd)if botstate.Err != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	if len(botstate.Banned.Fucklist) != 0 {
		listadm += "\n\n ☠️ 𝗙𝘂𝗰𝗸𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range botstate.Banned.Fucklist {
			num++
			rengs := strconv.Itoa(num)
			x, botstate.Err := self.GetContact(xd)if botstate.Err != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	if len(botstate.Banned.Mutelist) != 0 {
		listadm += "\n\n ☠️ 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁 ☠️ "
		for num, xd := range botstate.Banned.Mutelist {
			num++
			rengs := strconv.Itoa(num)
			x, botstate.Err := self.GetContact(xd)if botstate.Err != nil {
				listadm += "\n " + rengs + ". Closed Account"
			} else {
				listadm += "\n " + rengs + ". " + x.DisplayName
			}
		}
	}
	return listadm
}

func Cekbanwhois(client *linetcr.Account, to string, targets []string) {
	room := linetcr.GetRoom(to)
	list := ""
	if len(targets) > 1 {
		ban := []string{}
		fuck := []string{}
		mute := []string{}
		Gban := []string{}
		for _, from := range targets {
			if botstate.Banned.GetFuck(from) {
				fuck = append(fuck, from)
			} else if botstate.Banned.GetBan(from) {
				ban = append(ban, from)
			} else if botstate.Banned.GetMute(from) {
				mute = append(mute, from)
			} else if utils.InArrayString(room.Gban, from) {
				Gban = append(Gban, from)
			}
		}
		if len(ban) != 0 {
			list += "𝗘𝘅𝗶𝘀𝘁.𝗶𝗻 𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
			for n, xx := range ban {
				rengs := strconv.Itoa(n + 1)
				x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
					list += rengs + ". Closed Account \n"
				} else {
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(fuck) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗳𝘂𝗰𝗸𝗹𝗶𝘀𝘁:\n"
			for n, xx := range fuck {
				rengs := strconv.Itoa(n + 1)
				x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
					list += rengs + ". Closed Account \n"
				} else {
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(Gban) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗴𝗯𝗮𝗻𝗹𝗶𝘀𝘁:\n"
			for n, xx := range Gban {
				rengs := strconv.Itoa(n + 1)
				x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
					list += rengs + ". Closed Account \n"
				} else {
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
		if len(mute) != 0 {
			list += "\n𝗘𝘅𝗶𝘀𝘁 𝗶𝗻 𝗠𝘂𝘁𝗲𝗹𝗶𝘀𝘁:\n"
			for n, xx := range mute {
				rengs := strconv.Itoa(n + 1)
				x, botstate.Err := client.GetContact(xx)if botstate.Err != nil {
					list += rengs + ". Closed Account \n"
				} else {
					list += rengs + ". " + x.DisplayName + "\n"
				}
			}
		}
	} else {
		for _, from := range targets {
			if botstate.Banned.GetFuck(from) {
				list += "User have access exist in fuck list."
			} else if botstate.Banned.GetBan(from) {
				list += "User have access exist in ban list."
			} else if utils.InArrayString(room.Gban, from) {
				list += "User have access exist in gban list."
			} else if botstate.Banned.GetMute(from) {
				list += "User have access exist in mute list."
			}

		}
	}
	if list != "" {
		client.SendMessage(to, botstate.Fancy(list))
	}
}

func PerCheckList() string {
	list := ""
	var test1 string
	if botstate.SetHelper.Rngcmd != nil {
		list += "✠ 𝗟𝗶𝘀𝘁 𝗽𝗲𝗿𝗺 :\n\n"
		for i := range botstate.SetHelper.Rngcmd {
			if botstate.SetHelper.Rngcmd[i] == 0 {
				test1 = "Dev"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 1 {
				test1 = "Creator"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 2 {
				test1 = "Maker"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 3 {
				test1 = "Seller"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 4 {
				test1 = "Buyer"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 5 {
				test1 = "Owner"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 6 {
				test1 = "Master"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 7 {
				test1 = "Admin"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 8 {
				test1 = "Gowner"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			} else if botstate.SetHelper.Rngcmd[i] == 9 {
				test1 = "Gadmin"
				list += fmt.Sprintf("%v : %v\n", i, test1)
			}
		}
	}
	return list
}


func GenerateTimeLog(client *linetcr.Account, to string) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	a := time.Now().In(loc)
	yyyy := strconv.Itoa(a.Year())
	MM := a.Month().String()
	dd := strconv.Itoa(a.Day())
	hh := a.Hour()
	mm := a.Minute()
	ss := a.Second()
	var hhconv string
	var mmconv string
	var ssconv string
	if hh < 10 {
		hhconv = "0" + strconv.Itoa(hh)
	} else {
		hhconv = strconv.Itoa(hh)
	}
	if mm < 10 {
		mmconv = "0" + strconv.Itoa(mm)
	} else {
		mmconv = strconv.Itoa(mm)
	}
	if ss < 10 {
		ssconv = "0" + strconv.Itoa(ss)
	} else {
		ssconv = strconv.Itoa(ss)
	}
	times := "Date : " + dd + "-" + MM + "-" + yyyy + "\nTime : " + hhconv + ":" + mmconv + ":" + ssconv
	client.SendMessage(to, botstate.Fancy(times))
}

func GetComs(gr int, data string) int {
	defer linetcr.PanicOnly()
	_, value := botstate.SetHelper.Rngcmd[data]
	if value == false {
		botstate.SetHelper.Rngcmd[data] = gr
	}
	xx := botstate.SetHelper.Rngcmd[data]
	return xx
}

func CheckAccount(user string) *linetcr.Account {
	for _, cl := range botstate.ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}

func CheckUser(client *linetcr.Account, group string) ([]*linetcr.Account, []string) {
	list := []string{}
	botstate.Err, _, memlist := client.GetGroupMembers(group)
	if botstate.Err != nil {
		return nil, list
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := CheckAccount(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		} else if MemUser(group, mid) {
			list = append(list, mid)
		}
	}
	if len(exe) != 0 && len(list) != 0 {
		return exe, list
	}
	return nil, list
}

func AllBotList(user string) bool {
	if utils.InArrayString(botstate.Squadlist, user) {
		return true
	} else if botstate.UserBot.GetBot(user) {
		return true
	}
	return false
}

func DelJoin(user string) {
	for _, us := range botstate.Opjoin {
		if us == user {
		}
	}
}

func CekKick(optime int64) bool {
	for _, tar := range botstate.Opkick {
		if tar == optime {
			return false
		}
	}
	return true
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

func CekJoin(optime string) bool {
	defer linetcr.PanicOnly()
	for _, tar := range botstate.Opjoin {
		if tar == optime {
			return false
		}
	}
	return true
}

func GetBot(client *linetcr.Account, to string) []*linetcr.Account {
	_, memlist := client.GetGroupMember(to)
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(botstate.Squadlist, mid) {
			cl := GetKorban(mid)
			if cl.Limited == false {
				exe = append(exe, cl)
			}
		}
	}
	sort.Slice(exe, func(i, j int) bool {
		return exe[i].KickPoint < exe[j].KickPoint
	})
	linetcr.GetRoom(to).HaveClient = exe
	return exe
}