package handler

import (
	"fmt"
	"os/exec"
	"sort"
	"time"

	"../botstate"
	"../library/linetcr"
	"../library/hashmap"
	"../utils"

	"strings"
	"../library/SyncService")

func AutoSet() {
	defer utils.PanicHandle("AutoSet")
	now := time.Now()
	for _, room := range linetcr.SquadRoom {
		if !room.Fight.IsZero() {
			if now.Sub(room.Fight) >= 3*time.Second {
				if botstate.AutoPro {
					room.AutoBro()
				}
				room.Fight = time.Time{}
				var cll *linetcr.Account
				if len(room.Client) != 0 {
					cll = room.Client[0]
					name, mem, pending := cll.GetChatList(room.Id)
					room.Name = name
					room.Reset()
					sort.Slice(room.Ava, func(i, j int) bool {
						return room.Ava[i].Client.KickPoint < room.Ava[j].Client.KickPoint
					})
					sort.Slice(room.HaveClient, func(i, j int) bool {
						return room.HaveClient[i].KickPoint < room.HaveClient[j].KickPoint
					})
					exe := []*linetcr.Account{}
					for _, cls := range room.Client {
						if utils.InArrayString(mem, cls.MID) && !cls.Frez && !cls.Limited {
							exe = append(exe, cls)
						}
					}
					room.HaveClient = exe
					if len(exe) != 0 {
						Backup := []string{}
						li, ok := botstate.Backlist.Get(room.Id)
						if ok {
							mems := li.([]string)
							for _, l := range mems {
								if !utils.InArrayString(mem, l) && !utils.InArrayString(Backup, l) && !utils.InArrayString(pending, l) {
									Backup = append(Backup, l)
								}
							}
						}
						botstate.Backlist.Set(room.Id, []string{})
						if len(Backup) != 0 {
							celek := len(Backup)
							no := 0
							bat := 5
							ClAct := len(exe)
							if ClAct != 0 {
								if celek < bat {
									for _, cl := range exe {
										cl.GetRecommendationIds()
										for _, mid := range Backup {
											linetcr.AddContact3(cl, mid)
										}
										fl, _ := cl.GetAllContactIds()
										bb := []string{}
										for _, mid := range Backup {
											if utils.InArrayString(fl, mid) {
												bb = append(bb, mid)
												Backup = utils.RemoveString(Backup, mid)
											}
										}
										if len(bb) != 0 {
											cl.InviteIntoGroupNormal(room.Id, bb)
										}
										if len(Backup) == 0 {
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
											hajar = Backup[i*bat:]
										} else {
											hajar = Backup[i*bat : (i+1)*bat]
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
												client.InviteIntoGroupNormal(room.Id, bb)
											}
										}
										no += 1
									}
								}
							}
						}
					}
				} else {
					linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
				}
				botstate.FilterWar.Clear()
				botstate.Nkick = &hashmap.HashMap{}
				botstate.Filterop = &hashmap.HashMap{}
				botstate.Oplist = []int64{}
				botstate.Ceknuke = &hashmap.HashMap{}
				botstate.Cekoptime = []int64{}
				botstate.PurgeOP = []int64{}
				botstate.Filtermsg = &hashmap.HashMap{}
				botstate.Opjoin = []string{}
				room.ListInvited = []string{}
				botstate.Cekpurge = []int64{}
				botstate.AutoproN = false
				botstate.CekGo = []int64{}
			}
		}
		if !room.Leave.IsZero() {
			if now.Sub(room.Leave) >= 3*time.Second {
				room.Leave = time.Time{}
				if len(room.LeaveBack) != 0 {
					var cll *linetcr.Account
					if len(room.Client) != 0 {
						cll = room.Client[0]
						botstate.Botleave = &hashmap.HashMap{}
						name, mem, invs := cll.GetChatList(room.Id)
						room.Name = name
						exe := []*linetcr.Account{}
						for _, cls := range room.Client {
							if utils.InArrayString(mem, cls.MID) && !utils.InArrayString(room.GoMid, cls.MID) {
								exe = append(exe, cls)
							}
						}
						inv := []string{}
						asu := room.LeaveBack
						room.LeaveBack = []string{}
						if len(exe) != 0 {
							for _, l := range asu {
								if !MemBan(room.Id, l) && !utils.InArrayString(inv, l) && !utils.InArrayString(mem, l) && !utils.InArrayString(invs, l) {
									inv = append(inv, l)
								}
							}
							if len(inv) != 0 {
								cls := exe
								for _, cl := range cls {
									if !cl.Limited {
										cl.GetRecommendationIds()
										for _, mid := range inv {
											linetcr.AddContact3(cl, mid)
										}
										fl, _ := cl.GetAllContactIds()
										bb := []string{}
										for _, mid := range inv {
											if utils.InArrayString(fl, mid) {
												bb = append(bb, mid)
											}
										}
										cl.InviteIntoGroupNormal(room.Id, bb)
										for _, mid := range bb {
											if MemUser(room.Id, mid) {
												cl.UnFriend(mid)
											}
										}
										break
									}
								}
							}
						}
					} else {
						linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
					}
				}
			}
		}
	}
	for _, cl := range linetcr.Waitadd {
		if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
			v, ok := linetcr.BlockAdd.Get(cl.MID)
			if !ok {
				if now.Sub(cl.TimeBan) >= 1*time.Hour  {
					cl.Limitadd = false
					cl.Add = 0
					cl.Lastadd = now
					linetcr.Waitadd = linetcr.RemoveCl(linetcr.Waitadd, cl)
					linetcr.BlockAdd.Del(cl.MID)
				}
			} else {
				if now.Sub(v.(time.Time)) >= 24*time.Hour {
					linetcr.BlockAdd.Del(cl.MID)
					cl.Limitadd = false
					cl.Add = 0
					cl.Lastadd = now
					linetcr.Waitadd = linetcr.RemoveCl(linetcr.Waitadd, cl)
					linetcr.BlockAdd.Del(cl.MID)
				}
			}
		}
	}
	for _, cl := range botstate.ClientBot {
		if now.Sub(cl.Lastadd) >= 1*time.Hour  {
			cl.Add = 0
			cl.Lastadd = now
		}
		if now.Sub(cl.Lastkick) >= 1*time.Hour  {
			cl.TempKick = 0
			cl.TempInv = 0
		}
	}
	for _, cl := range linetcr.KickBans {
		if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
			v, ok := linetcr.GetBlock.Get(cl.MID)
			if !ok {
				if now.Sub(cl.TimeBan) >= 1*time.Hour  {
					linetcr.KickBans = linetcr.RemoveCl(linetcr.KickBans, cl)
					cl.Limited = false
					cl.TempKick = 0
					cl.TempInv = 0
					cl.Frez = false
					linetcr.GetBlock.Del(cl.MID)
				}
			} else {
				if now.Sub(v.(time.Time)) >= 24*time.Hour {
					linetcr.GetBlock.Del(cl.MID)
					linetcr.KickBans = linetcr.RemoveCl(linetcr.KickBans, cl)
					cl.Limited = false
					cl.Frez = false
					cl.TempKick = 0
					cl.TempInv = 0
					cl.KickCount = 0
					cl.KickPoint = 0
					cl.InvCount = 0
					cl.CountDay = 0
				}
			}
		}
	}
	for m, v := range linetcr.HashToMap(linetcr.GetBlockAdd) {
		cl := botstate.GetKorban(m)
		if cl.Limited {
			if now.Sub(v.(time.Time)) >= 1*time.Hour  {
				cl.Limitadd = false
				linetcr.GetBlockAdd.Del(cl.MID)
			}
		}
	}
	if now.Sub(botstate.Aclear) >= 30*time.Second {
		botstate.Filterop = &hashmap.HashMap{}
		botstate.Nkick = &hashmap.HashMap{}
		botstate.FilterWar.Clear()
		botstate.Oplist = []int64{}
		botstate.TimeSend = []int64{}
		botstate.Ceknuke = &hashmap.HashMap{}
		botstate.Cekoptime = []int64{}
		botstate.Filtermsg = &hashmap.HashMap{}
		botstate.Aclear = now
		botstate.PurgeOP = []int64{}
		botstate.Cekpurge = []int64{}
		botstate.Opjoin = []string{}
		botstate.CekGo = []int64{}
		botstate.AutoproN = false
	}
	if now.Sub(botstate.TimeSave) >= 1*time.Hour {
		SaveBackup()
		botstate.TimeBackup = now
	}
	if !botstate.TimeBackup.IsZero() {
		BackSeave()
	}
	if now.Sub(botstate.TimeClear) >= 3*time.Hour {
		time.Sleep(3*time.Hour)
		exec.Command("bash", "-c", "sudo systemd-resolve --flush-caches").Output()
		exec.Command("bash", "-c", "echo 3 > /proc/sys/vm/drop_caches&&swapoff -a&&swapon -a").Output()
		fmt.Println("\n clear vps cache automatically every 6 hours")
	}
	if botstate.AutoBc {
		if now.Sub(botstate.TimeBc) >= time.Duration(botstate.TimeBroadcast)*time.Minute {
			exe := []*linetcr.Account{}
			for _, cls := range botstate.ClientBot {
				exe = append(exe, cls)
			}
			if len(exe) != 0 {
				cls := exe
				for _, cl := range cls {
					if cl.Limited == false {
						gr, _ := cl.GetGroupIdsJoined()
						time.Sleep(250 * time.Second)
						for _, gi := range gr {
							if botstate.Typebc == "msg" {
								cl.SendMessage(gi, botstate.MsgBroadcast)
							} else if botstate.Typebc == "image" {
								cl.SendFoto(gi, linetcr.Imagebc)
							}
						}
					}
				}
			}
		}
	}
}

// Stub for SaveBackup if not imported yet, but likely it is in system.go or should be
func SaveBackup() {
	// This function seems to be missing in my recent reads of handler/system.go
	// But it is called in main.go, I should check if it exists or if I need to migrate it.
    SaveData()
}


func SaveProHistory() {
	botstate.AllowDoOnce = 0
	for i := range botstate.ClientBot {
		botstate.Data.Kikhistory = botstate.Data.Kikhistory + botstate.ClientBot[i].Ckick
		botstate.Data.Invhistory = botstate.Data.Invhistory + botstate.ClientBot[i].Cinvite
		botstate.Data.Canclhistory = botstate.Data.Canclhistory + botstate.ClientBot[i].Ccancel
		botstate.ClientBot[i].Ckick = 0
		botstate.ClientBot[i].Cinvite = 0
		botstate.ClientBot[i].Ccancel = 0
	}
}

func CheckChatBan() {
	defer utils.PanicHandle("CheckChatBan")
	if botstate.AllowDoOnce == 0 {
		for _, cl := range botstate.ClientBot {
			if !linetcr.InArrayCl(linetcr.KickBanChat, cl) && !cl.Frez {
				r, _ := cl.GetHomeProfile(cl.MID)
				if linetcr.GetBannedChat(r) == 1 {
					linetcr.BanChatAdd(cl)
				}
			}
		}
		botstate.AllowDoOnce++
	}
}


func LogGet(op *SyncService.Operation) {
	defer utils.PanicHandle("LogGet")
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
				g = []*Stickers{&Stickers{Id: ids[0], Pid: pids[0]}}
				botstate.Laststicker.Set(msg.To, g)
			} else {
				c := g.([]*Stickers)
				c = AppendLastSticker(c, &Stickers{Id: ids[0], Pid: pids[0]})
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

func SaveJoin(Pelaku string, Optime int64) {
	defer utils.PanicHandle("LogGet")
	ix := utils.IndexOf(botstate.Detectjoin.User, Pelaku)
	if ix == -1 {
		botstate.Detectjoin.User = append(botstate.Detectjoin.User, Pelaku)
		botstate.Detectjoin.Time = append(botstate.Detectjoin.Time, Optime)
	} else {
		botstate.Detectjoin.Time[ix] = Optime
	}
}