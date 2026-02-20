package handler

import (
	"fmt"
	"time"
	"../botstate"
	"../library/linetcr"
	"../utils"
)

func AddCon(cons []string) {
	n := 0
	for _, con := range cons {
		for _, cl := range botstate.ClientBot {
			fl, _ := cl.GetAllContactIds()
			if !utils.InArrayString(fl, con) && con != cl.MID && !cl.Limitadd {
				if linetcr.IsFriends(cl, con) == false {cl.FindAndAddContactsByMidV2(con);time.Sleep(3 * time.Second)}
			}
		}
		n += 1
	}
}

func AddConSq(to string, cons []string) {
	n := 0
	for _, con := range cons {
		for _, cl := range botstate.ClientBot {
			fl, _ := cl.GetAllContactIds()
			if !utils.InArrayString(fl, con) && con != cl.MID && !cl.Limitadd {
				if linetcr.IsFriends(cl, con) == false {time.Sleep(5 * time.Second);cl.FindAndAddContactsByMidV2(con);time.Sleep(5 * time.Second)}
			}
			cl.SendMessage(to, botstate.Fancy("Added friends"))
		}
		n += 1
	}
}


func AddConSqV2(cons []string) {
	for _, cl := range botstate.ClientBot {
		for _, con := range cons {
			if linetcr.IsFriends(cl, con) == false && con != cl.MID {
				time.Sleep(5 * time.Second)
				cl.FindAndAddContactsByMidV5(con)
				time.Sleep(250 * time.Millisecond)
			}
		}
	}
}

func AddConSingle(cons []string) {
	n := 0
	for _, con := range cons {
		for _, cl := range botstate.ClientBot {
			fl, _ := cl.GetAllContactIds()
			if !utils.InArrayString(fl, con) && con != cl.MID && !cl.Limitadd {
				if linetcr.IsFriends(cl, con) == false {cl.FindAndAddContactsByMidV2(con);time.Sleep(3 * time.Second)}
			}
		}
		n += 1
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

func DataMention(to string, mtxt string, targetlist []string) {
	defer utils.PanicHandle("DataMention")
	Room := linetcr.GetRoom(to)
	if Room.Id != "" {
		Cans := Room.Client
		if len(Cans) != 0 {
			memlist := []string{}
			for _, mid := range targetlist {
				if !utils.InArrayString(botstate.Squadlist, mid) {
					memlist = append(memlist, mid)
				}
			}
			if len(memlist) != 0 {
				if len(memlist) <= 20 || len(Cans) == 1 {
					cl := Cans[0]
					cl.SendPollMention(to, mtxt, memlist)
					time.Sleep(1 * time.Second)
				} else {
					tx := ""
					nob := 0
					ta := false
					tag := []string{}
					z := len(memlist) / 20
					y := z + 1
					for i := 0; i < y; i++ {
						if !ta {
							tx += fmt.Sprintf("%s\n", mtxt)
							ta = true
						}
						if i == z {
							tag = memlist[i*20:]
							no := i * 20
							no += 1
							for i := 0; i < len(tag); i++ {
								iki := no + i
								if iki < 10 {
									tx += fmt.Sprintf("0%v. @!\n", iki)
								} else {
									tx += fmt.Sprintf("%v. @!\n", iki)
								}
							} 
						} else {
							tag = memlist[i*20 : (i+1)*20]
							no := i * 20
							no += 1
							for i := 0; i < len(tag); i++ {
								iki := no + i
								if iki < 10 {
									tx += fmt.Sprintf("0%v. @!\n", iki)
								} else {
									tx += fmt.Sprintf("%v. @!\n", iki)
								}
							}
						}
						if len(tag) != 0 {
							if nob >= len(Cans) {
								nob = 0
							}
							Cans[nob].SendMention(to, tx, tag)
							nob++
						}
						tx = ""
					}
				}
			}
		}
	}
}


func CheckQr() {
	botstate.Qrwar = true
	time.Sleep(1 * time.Second)
	botstate.Qrwar = false
}

func CmdListCheck() string {
	list2 := "𝗟𝗶𝘀𝘁 𝗖𝗺𝗱:\n\n"
	list := ""
	if botstate.Commands.Botname != "" {
		list += fmt.Sprintf(" - Botname: %s\n", botstate.Commands.Botname)
	}
	if botstate.Commands.Upallimage != "" {
		list += fmt.Sprintf(" - Upallimage: %s\n", botstate.Commands.Upallimage)
	}
	if botstate.Commands.Upallcover != "" {
		list += fmt.Sprintf(" - Upallcover: %s\n", botstate.Commands.Upallcover)
	}
	if botstate.Commands.Unsend != "" {
		list += fmt.Sprintf(" - Unsend: %s\n", botstate.Commands.Unsend)
	}
	if botstate.Commands.Upvallimage != "" {
		list += fmt.Sprintf(" - Upvallimage: %s\n", botstate.Commands.Upvallimage)
	}
	if botstate.Commands.Upvallcover != "" {
		list += fmt.Sprintf(" - Upvallcover: %s\n", botstate.Commands.Upvallcover)
	}
	if botstate.Commands.Appname != "" {
		list += fmt.Sprintf(" - Appname: %s\n", botstate.Commands.Appname)
	}
	if botstate.Commands.Useragent != "" {
		list += fmt.Sprintf(" - Useragent: %s\n", botstate.Commands.Useragent)
	}
	if botstate.Commands.Hostname != "" {
		list += fmt.Sprintf(" - Hostname: %s\n", botstate.Commands.Hostname)
	}
	if botstate.Commands.Friends != "" {
		list += fmt.Sprintf(" - Friends: %s\n", botstate.Commands.Friends)
	}
	if botstate.Commands.Adds != "" {
		list += fmt.Sprintf(" - Adds: %s\n", botstate.Commands.Adds)
	}
	if botstate.Commands.Limits != "" {
		list += fmt.Sprintf(" - Limits: %s\n", botstate.Commands.Limits)
	}
	if botstate.Commands.Addallbots != "" {
		list += fmt.Sprintf(" - Addallbots: %s\n", botstate.Commands.Addallbots)
	}
	if botstate.Commands.Addallsquads != "" {
		list += fmt.Sprintf(" - Addallsquads: %s\n", botstate.Commands.Addallsquads)
	}
	if botstate.Commands.Leave != "" {
		list += fmt.Sprintf(" - Leave: %s\n", botstate.Commands.Leave)
	}
	if botstate.Commands.Respon != "" {
		list += fmt.Sprintf(" - Respon: %s\n", botstate.Commands.Respon)
	}
	if botstate.Commands.Ping != "" {
		list += fmt.Sprintf(" - Ping: %s\n", botstate.Commands.Ping)
	}
	if botstate.Commands.Count != "" {
		list += fmt.Sprintf(" - Count: %s\n", botstate.Commands.Count)
	}
	if botstate.Commands.Limitout != "" {
		list += fmt.Sprintf(" - 1111111: %s\n", botstate.Commands.Limitout)
	}
	if botstate.Commands.Access != "" {
		list += fmt.Sprintf(" - Access: %s\n", botstate.Commands.Access)
	}
	if botstate.Commands.Allbanlist != "" {
		list += fmt.Sprintf(" - Allbanlist: %s\n", botstate.Commands.Allbanlist)
	}
	if botstate.Commands.Allgaccess != "" {
		list += fmt.Sprintf(" - Allgaccess: %s\n", botstate.Commands.Allgaccess)
	}
	if botstate.Commands.Gaccess != "" {
		list += fmt.Sprintf(" - Gaccess: %s\n", botstate.Commands.Gaccess)
	}
	if botstate.Commands.Checkram != "" {
		list += fmt.Sprintf(" - Checkram: %s\n", botstate.Commands.Checkram)
	}
	if botstate.Commands.Backups != "" {
		list += fmt.Sprintf(" - Backups: %s\n", botstate.Commands.Backups)
	}
	if botstate.Commands.Upimage != "" {
		list += fmt.Sprintf(" - Upimage: %s\n", botstate.Commands.Upimage)
	}
	if botstate.Commands.Upcover != "" {
		list += fmt.Sprintf(" - Upcover: %s\n", botstate.Commands.Upcover)
	}
	if botstate.Commands.Upvimage != "" {
		list += fmt.Sprintf(" - Upvimage: %s\n", botstate.Commands.Upvimage)
	}
	if botstate.Commands.Upvcover != "" {
		list += fmt.Sprintf(" - Upvcover: %s\n", botstate.Commands.Upvcover)
	}
	if botstate.Commands.Bringall != "" {
		list += fmt.Sprintf(" - Bringall: %s\n", botstate.Commands.Bringall)
	}
	if botstate.Commands.Purgeall != "" {
		list += fmt.Sprintf(" - Purgeall: %s\n", botstate.Commands.Purgeall)
	}
	if botstate.Commands.Banlist != "" {
		list += fmt.Sprintf(" - Banlist: %s\n", botstate.Commands.Banlist)
	}
	if botstate.Commands.Clearban != "" {
		list += fmt.Sprintf(" - Clearban: %s\n", botstate.Commands.Clearban)
	}
	if botstate.Commands.Stayall != "" {
		list += fmt.Sprintf(" - Stayall: %s\n", botstate.Commands.Stayall)
	}
	if botstate.Commands.Clearchat != "" {
		list += fmt.Sprintf(" - Clearchat: %s\n", botstate.Commands.Clearchat)
	}
	if botstate.Commands.Here != "" {
		list += fmt.Sprintf(" - Here: %s\n", botstate.Commands.Here)
	}
	if botstate.Commands.Speed != "" {
		list += fmt.Sprintf(" - Speed: %s\n", botstate.Commands.Speed)
	}
	if botstate.Commands.Status != "" {
		list += fmt.Sprintf(" - Status: %s\n", botstate.Commands.Status)
	}
	if botstate.Commands.Tagall != "" {
		list += fmt.Sprintf(" - Tagall: %s\n", botstate.Commands.Tagall)
	}
	if botstate.Commands.Kick != "" {
		list += fmt.Sprintf(" - Kick: %s\n", botstate.Commands.Kick)
	}
	if botstate.Commands.Max != "" {
		list += fmt.Sprintf(" - Protect Max: %s\n", botstate.Commands.Max)
	}
	if botstate.Commands.None != "" {
		list += fmt.Sprintf(" - Protect None: %s\n", botstate.Commands.None)
	}
	if botstate.Commands.Kickall != "" {
		list += fmt.Sprintf(" - Kickall: %s\n", botstate.Commands.Kickall)
	}
	if botstate.Commands.Cancelall != "" {
		list += fmt.Sprintf(" - Cancelall: %s\n", botstate.Commands.Cancelall)
	}
	if list != "" {
		return list2 + list

	} else {
		return "Not found set Cmd.\n"
	}
}

func AddContact2(cl *linetcr.Account, con string) int {
	fl, _ := cl.GetAllContactIds()
	if !utils.InArrayString(fl, con) {
		if con != cl.MID && !cl.Limitadd {
			_, botstate.Err := cl.FindAndAddContactsByMidV2(con)
			if botstate.Err != nil {
				println(fmt.Sprintf("%v", botstate.Err.Error()))
				return 0
			}
			return 1
		} else {
			return 0
		}
	}
	return 1
}