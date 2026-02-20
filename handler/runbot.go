package handler

import (
    "fmt"
"sync"
"strings"
"time"
"runtime"
"../botstate"
"../config"
"../library/linetcr"
"../utils"
"../library/SyncService"
)

func RunBot(client *linetcr.Account, ch chan int) {
	defer utils.PanicHandle("RunBot")
	runtime.GOMAXPROCS(botstate.Cpu)
	client.Revision = -1
	for {
		multiFunc, botstate.Err := client.SyncLoad(100)
		if botstate.Err != nil || len(multiFunc) == 0 {
			continue
		}
		go func(fetch []*SyncService.Operation) {
			for _, op := range multiFunc {
				if op.Type == 124 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Optime := op.CreatedTime
					rngcmd := botstate.GetComs(7, "invitebot")
					Group, user := op.Param1, op.Param2
					invited := strings.Split(op.Param3, "\x1e")
					Room := linetcr.GetRoom(Group)
					if utils.InArrayString(invited, client.MID) {
						if linetcr.IoGOBot(Group, client) {
							if utils.InArrayString(client.Squads, user) {
								go func(client *linetcr.Account, Group string){
									go AccKickBan(client, Group)
								}(client, Group)
							} else if botstate.UserBot.GetBot(user) {
								var wg sync.WaitGroup
								wg.Add(1)
								go func(client *linetcr.Account, Group string){
									defer wg.Done()
									go AccGroup(client, Group)
								}(client, Group)
								wg.Wait()
								if botstate.AutoPurge {
									go JoinLlinetcrBan(client, Group)
								}
								if client.Limited == false {if !utils.InArrayInt64(botstate.CekGo, Optime) {botstate.CekGo = append(botstate.CekGo, Optime)
									AcceptJoin(client, Group)}
								}
							} else if GetCodeprem(rngcmd, user, Group) {
								var wg sync.WaitGroup
								wg.Add(1)
								go func(client *linetcr.Account, Group string){
									defer wg.Done()
									go AccGroup(client, Group)
								}(client, Group)
								wg.Wait()
								if botstate.AutoPurge {
									go JoinLlinetcrBan(client, Group)
								}
								if client.Limited == false {if !utils.InArrayInt64(botstate.CekGo, Optime) {botstate.CekGo = append(botstate.CekGo, Optime)
									AcceptJoin(client, Group)}
								}
							} else {
								grs, _ := client.GetGroupIdsJoined()
								if utils.InArrayString(grs, Group) {
									client.LeaveGroup(Group)
									fl, _ := client.GetAllContactIds()
									if utils.InArrayString(fl, user) {
										client.UnFriend(user)
									}
								}
							}
						}
					} else {
						Optime := op.CreatedTime
						if Room.ProInvite {
							if MemUser(Group, user) {
								go func(client *linetcr.Account, Group string, user string) {
									if botstate.FilterWar.Cek(user) {
										KickCancelProtect(client, user, invited[0], Group)
										botstate.Banned.AddBan(user)
										botstate.FilterWar.Del(user)
									}
								}(client, Group, user)
								if botstate.AutoPurge {
									go func(client *linetcr.Account, Group string, user string) {
										if botstate.FilterWar.Ceki(user) {
											botstate.Banned.AddBan(user)
											kickProtect(client, Group, user)
											botstate.FilterWar.Deli(user)
										}
									}(client, Group, user)
									if botstate.FilterWar.Ceko(Optime) {
										Room.ListInvited = invited
										BanAll(invited, Group)
										go cancelBanInv(client, invited, Group)
									}
								}
							} else {
								if botstate.FilterWar.Ceko(Optime) {
									go cancelallcek(client, invited, Group)
								}
							}
						} else {
							if MemBan(Group, user) {
								go func() {
									if botstate.FilterWar.Ceki(user) {
										kickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
									}
								}()
								if botstate.FilterWar.Ceko(Optime) {
									BanAll(invited, Group)
									go cancelBanInv(client, invited, Group)
									go func() {NodeBans(client, Group, invited)}()
								}
							} else {
								if MemUser(Group, user) {
									go func() {
										if botstate.FilterWar.Ceki(user) {
											for _, vo := range invited {
												if MemBan(Group, vo) {
													botstate.Banned.AddBan(user)
													kickPelaku(client, Group, user)
													break
												}
											}
											botstate.FilterWar.Deli(user)
										}
									}()
									if botstate.FilterWar.Ceko(Optime) {
										go cancelallcek(client, invited, Group)
									}
								}
							}
						}
					}
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
					if botstate.LogMode && MemAccsess(op.Param2) {NotifBot(client, Group, "invite", user)}
				}
				if op.Type == 133 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := linetcr.GetRoom(Group)
					if client.MID == Invited {
						linetcr.Gones(Group, client)
						if MemUser(Group, user) {
							botstate.Banned.AddBan(user)
						}
					} else if !utils.InArrayString(Room.GoMid, client.MID) {
						if utils.InArrayString(client.Squads, Invited) {
							if MemUser(Group, user) {
								if linetcr.IoGOBot(Group, client) {
									botstate.Banned.AddBan(user)
									go func() {
										if botstate.FilterWar.Cek(user) {
											groupBackupKick(client, Group, user, true)
											botstate.FilterWar.Del(user)
										}
									}()
									if botstate.FilterWar.Cek(Invited) {
										groupBackupInv(client, Group, Optime, Invited)
										botstate.FilterWar.Del(Invited)
									}
								}
							}
						} else {
							if !MemUserN(Group, Invited) {
								if Checkkickuser(Group, user, Invited) {
									back(Group, Invited)
									if botstate.FilterWar.Ceki(user) {
										kickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
										if MemUser(Group, user) {
											botstate.Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProKick {
									if MemUser(Group, user) {
										if Room.Backup {
											back(Group, Invited)
										}
										if _, ok := botstate.Nkick.Get(user); !ok {
											botstate.Nkick.Set(user, 1)
											kickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)
										}
                                                                   ct, _ := client.GetContact(Invited)
                                                                   if ct != nil {if ct.CapableBuddy {kickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)}
										}
									}
								}
							}
						}
					} else {
						if MemUser(Group, Invited) {
							if MemUser(Group, user) {
								back(Group, Invited)
								botstate.Banned.AddBan(user)
							}
						}
						if MemUser(Group, user) {
							if botstate.FilterWar.Ceki(user) {
								GhostEnd(client, Group, Optime, user, true)
								botstate.FilterWar.Deli(user)
							}
						}
					}
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
					if botstate.LogMode && MemAccsess(op.Param2) {NotifBot(client, Group, "kick", user)}
				}
				if op.Type == 129 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Group, user := op.Param1, op.Param2
					if botstate.PowerMode == true {
					if MemBan(Group, user) {
						go func() {
							if botstate.FilterWar.Ceki(user) {
								AccKickBan(client, Group)
								botstate.FilterWar.Deli(user)
							}
						}()
					}}
				}
				if op.Type == 130 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Group, user := op.Param1, op.Param2
					Room := linetcr.GetRoom(Group)
					if linetcr.IoGOBot(Group, client) {
						if Room.ProJoin || botstate.AutoproN == true {
							if MemUser(Group, user) {
								if botstate.FilterWar.Ceki(user) {
										botstate.Banned.AddBan(user)
									kickPelaku(client, Group, user)
									botstate.FilterWar.Deli(user)
								}
							}
						} else {
							if MemBan(Group, user) {
								if MemUser(Group, user) {
									if botstate.FilterWar.Ceki(user) {
										botstate.Banned.AddBan(user)
										kickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
									}
								}
							} else {
								if utils.InArrayString(Room.ListInvited, user) {
									if MemUser(Group, user) {
										if cekjoin(user) {
											kickPelaku(client, Group, user)
											deljoin(user)
											Room.ListInvited = Remove(Room.ListInvited, user)
										}
									} else {
										Room.ListInvited = Remove(Room.ListInvited, user)
									}
								} else {
									if Room.Welcome {
										if _, ok := botstate.Cewel.Get(user); !ok {
											botstate.Cewel.Set(user, 1)
											if cekjoin(user) {
												if !utils.InArrayString(botstate.Squadlist, user) {
													Room.WelsomeSet(client, Group, user)
												}
											}
										}
									} else {
										if botstate.LockMode == true {
											if MemUser(Group, user) {
												if botstate.FilterWar.Ceki(user) {
													botstate.Banned.AddBan2(user)
												}
											}
										}
									}
								}
							}
						}
					}
					Optime := op.CreatedTime
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				}
				if op.Type == 122 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Group, user, invited := op.Param1, op.Param2, op.Param3
					Optime := op.CreatedTime
					Room := linetcr.GetRoom(Group)
					if client.Limited == false && linetcr.IoGOBot(Group, client) {
						if MemUser(Group, user) {
							if Room.ProQr || botstate.AutoproN == true {
								if invited == "4" {
									if cekOp2(Optime) {
										go func() {
											cans := linetcr.Actor(Group)
											for _, cl := range cans {
												botstate.Err := cl.UpdateChatQrV2(Group, true)
												if botstate.Err == nil {
													break
												}
											}
										}()
										if botstate.FilterWar.Ceki(user) {
											kickPelaku(client, Group, user)
											botstate.FilterWar.Deli(user)
											botstate.Banned.AddBan(user)
										}
									}
								}
							} else if botstate.KickBanQr == true {
								if invited == "4" {
									if cekOp2(Optime) {go func() {cans := linetcr.Actor(Group);for _, cl := range cans {botstate.Err := cl.UpdateChatQrV2(Group, true);if botstate.Err == nil {break}}}()
										if botstate.FilterWar.Ceki(user) {kickPelaku(client, Group, user);botstate.FilterWar.Deli(user);botstate.Banned.AddBan(user)}
									}
								}
							} else if Room.ProPicture || botstate.AutoproN == true {
								if invited == "4" {
									if cekOp2(Optime) {
										if botstate.FilterWar.Ceki(user) {
											kickPelaku(client, Group, user)
											botstate.FilterWar.Deli(user)
										}
									}
								}
							} else if Room.ProName || botstate.AutoproN == true {
								if invited == "4" {
									if cekOp2(Optime) {
										go func() {
											cans := linetcr.Actor(Group)
											for _, cl := range cans {
												botstate.Err := cl.UpdateChatName(Group, Room.Name)
												if botstate.Err == nil {
													break
												}
											}
										}()
										if botstate.FilterWar.Ceki(user) {
											kickPelaku(client, Group, user)
											botstate.FilterWar.Deli(user)
										}
									}
								}
							} else {
								if MemBan(Group, user) {
									if invited == "4" {
										if cekOp2(Optime) {
											go func() {
												cans := linetcr.Actor(Group)
												for _, cl := range cans {
													botstate.Err := cl.UpdateChatQrV2(Group, true)
													if botstate.Err == nil {
														break
													}
												}
											}()
											if botstate.FilterWar.Ceki(user) {
												kickPelaku(client, Group, user)
												botstate.FilterWar.Deli(user)
												botstate.Banned.AddBan(user)
											}
										}
									} else if invited == "4" {
										if cekOp2(Optime) {
											go func() {
												cans := linetcr.Actor(Group)
												for _, cl := range cans {
													botstate.Err := cl.UpdateChatName(Group, Room.Name)
													if botstate.Err == nil {
														break
													}
												}
											}()
											if botstate.FilterWar.Ceki(user) {
												kickPelaku(client, Group, user)
												botstate.FilterWar.Deli(user)
											}
										}
									}
								}
							}

						}
					}
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				}
				if op.Type == 126 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := linetcr.GetRoom(Group)
					if client.MID == Invited {
						linetcr.Gones(Group, client)
						if MemUser(Group, user) {
							botstate.Banned.AddBan(user)
						}
					} else if !utils.InArrayString(Room.GoMid, client.MID) {
						if utils.InArrayString(client.Squads, Invited) {
							if MemUser(Group, user) {
								if linetcr.IoGOBot(Group, client) {
									botstate.Banned.AddBan(user)
									go func() {
										if botstate.FilterWar.Cek(user) {
											groupBackupCans(client, Group, user, true)
											botstate.FilterWar.Del(user)
										}
									}()
									if botstate.FilterWar.Cek(Invited) {
										groupBackupInv(client, Group, Optime, Invited)
										botstate.FilterWar.Del(Invited)
									}
								}
							}
						} else {
							if !MemUserN(Group, Invited) {
								if Checkkickuser(Group, user, Invited) {
									back(Group, Invited)
									if botstate.FilterWar.Ceki(user) {
										kickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
										if MemUser(Group, user) {
											botstate.Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProCancel {
									if MemUser(Group, user) {
										if Room.Backup {
											back(Group, Invited)
										}
										if _, ok := botstate.Nkick.Get(user); !ok {
											botstate.Nkick.Set(user, 1)
											kickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)
										}
                                                                   ct, _ := client.GetContact(Invited)
                                                                   if ct != nil {if ct.CapableBuddy {kickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)}
										}
									}
								}
							}
						}
					} else {
						if MemUser(Group, Invited) {
							if MemUser(Group, user) {
								back(Group, Invited)
								botstate.Banned.AddBan(user)
							}
						}
						if MemUser(Group, user) {
							if botstate.FilterWar.Ceki(user) {
								GhostEnd(client, Group, Optime, user, true)
								botstate.FilterWar.Deli(user)
							}
						}
					}
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
					if botstate.LogMode && MemAccsess(op.Param2) {NotifBot(client, Group, "cancel", user)}
				} else if op.Type == 33 {
					Group := op.Param1
					if botstate.LogMode && !utils.InArrayString(client.Squads, Group) {logAccess(client, client.Namebot, Group, "deleteaccount", []string{}, 2)}
				} else if op.Type == 5 {
					Group := op.Param1
					if botstate.LogMode && !utils.InArrayString(client.Squads, Group) {logAccess(client, client.Namebot, Group, "addfrind", []string{}, 2)}
				} else if op.Type == 50 {
					Group := op.Param1
					if botstate.LogMode && !utils.InArrayString(client.Squads, Group) {logAccess(client, client.Namebot, Group, "callme", []string{}, 2)}
				} else if op.Type == 55 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					if client.Limited == false && linetcr.IoGOBot(Group, client) {
						if cekOp(Optime) {
							if MemBan(Group, user) {
								kickPelaku(client, Group, user)
							} else {
								Room := linetcr.GetRoom(Group)
								if Room.Lurk && !utils.InArrayString(botstate.CheckHaid, user) {
									Room.CheckLurk(client, Group, user)
								}
							}
						}
					}
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 26 {
					msg := op.Message
					Optime := op.CreatedTime
					if msg.ContentType != 18 {
						if _, ok := botstate.Command.Get(Optime); !ok {
							botstate.Command.Set(Optime, client)
							if _, ok := botstate.Filterop.Get(Optime); !ok {
								botstate.Filterop.Set(Optime, 1)
								Bot(op, client, ch)
							}
						}
					}
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 128 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					if !utils.InArrayString(botstate.Squadlist, user) {
						Room := linetcr.GetRoom(Group)
						if Room.Backleave {
							jangan := true
							tm, ok := botstate.Botleave.Get(user)
							if ok {
								if time.Now().Sub(tm.(time.Time)) < 5*time.Second {
									jangan = false
								}
							}
							if jangan {
								if botstate.FilterWar.Ceki(user) {
									if !MemBan(Group, user) && !utils.InArrayString(botstate.Squadlist, user) && !botstate.UserBot.GetBot(user) && !utils.InArrayString(Room.GoMid, user) {
										Hstg(Group, user)
										Room.Leave = time.Now()
									}
								}
							}
						} else {
							if Room.Leavebool {
								if _, ok := botstate.Cleave.Get(user); !ok {
									botstate.Cleave.Set(user, 1)
									if !MemBan(Group, user) && !utils.InArrayString(botstate.Squadlist, user) && !botstate.UserBot.GetBot(user) && !utils.InArrayString(Room.GoMid, user) {
										Room.LeaveSet(client, Group, user)
									}
								}
							}
						}
					}
					if _, ok := botstate.Filtermsg.Get(Optime); !ok {
						botstate.Filtermsg.Set(Optime, client)
						LogOp(op, client)
						LogGet(op)
					}
				} else if op.Type == 30 {
					Group := op.Param1
					Room := linetcr.GetRoom(Group)
					if Room.Announce && linetcr.IoGOBot(Group, client) {
						Optime := op.CreatedTime
						if cekOp(Optime) {
							Room.CheckAnnounce(client, Group)
						}
					}
				} else if op.Type == 123 {
					client.CInvite()
				} else if op.Type == 132 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Group, user := op.Param1, op.Param2
					if botstate.PowerMode == true {
					if MemBan(Group, user) {
						go func() {
							if botstate.FilterWar.Ceki(user) {
								KickBan132(client, Group, user)
								botstate.FilterWar.Deli(user)
							}
						}()
					}}
					client.CountKick()
				} else if op.Type == 125 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					invited := strings.Split(op.Param3, "\x1e")
					Room := linetcr.GetRoom(Group)
					if botstate.PowerMode == true {
					if MemBan(Group, user) {
						go func() {
							if botstate.FilterWar.Ceko(Optime) {
								BanAll(invited, Group)
								go CancelBan125(client, Group, invited)
							}
						}()
					}}
					if Room.ProInvite {
						if MemUser(Group, user) {
							go func(client *linetcr.Account, Group string, user string) {
								if botstate.FilterWar.Cek(user) {
									KickCancelProtect(client, user, invited[0], Group)
									botstate.Banned.AddBan(user)
									botstate.FilterWar.Del(user)
								}
							}(client, Group, user)
						}
					}
					client.CCancel()
				}
			}
		}(multiFunc)
		for _, ops := range multiFunc {
			client.SetSyncRevision(ops.Revision)
		}
	}
}
