package main
//BOT_GO_MULTY_WAR_PROTECTION
//RECODE_BY: SELFTCR™
//ID_LINE: code-bot
//NEW__All_UPDATE: 21-04-2024

import (
	"./botstate"
	"./config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"github.com/shirou/gopsutil/mem"
	"unicode/utf8"
	"regexp"
	"log"
	"github.com/kardianos/osext"
	"github.com/tidwall/gjson"
	"net"
	"os/signal"
	"syscall"
	"strconv"
	"runtime"
	"net/http"
	"runtime/debug"
	"sort"
	"strings"
	"time"
	"./library/linetcr"
	handler "./handler"
	"./utils"
	"./library/unistyle"
	"github.com/panjf2000/ants"
	"./library/hashmap"
	"./library/SyncService"
	mod "./library/modcompact"
	call "./library/libcall/call"
	namegenerator "./library/RandomName"
	talkservice "./library/linethrift"
	valid "github.com/asaskevich/govalidator"
	"github.com/shirou/gopsutil/host"
	"github.com/opalmer/check-go-version/api"
)
//kick_sticker
//respon_sticker
//stayall_sticker
//leave_sticker
//kickall_sticker
//bypass_sticker
//invite_sticker
//clearban_sticker
//cancelall_sticker
func main() {
	if len(os.Args) < 1 {
		fmt.Println("do not forget arguement")
		os.Exit(1)
	}
	defer ants.Release()
	defer linetcr.PanicOnly()
	debug.SetGCPercent(500)
	jsonFile, botstate.Err := os.Open(botstate.DATABASE)
	if botstate.Err != nil {
		fmt.Println(botstate.Err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &botstate.Data)
	fmt.Println("\n[ START RUN GOLANG ]")
	mod.LoginBase()
	mod.TMoreCompact()
	go handler.GracefulShutdown()
	for no, tok := range botstate.Data.Authoken {
		time.Sleep(250 * time.Millisecond)
		sort := rand.Intn(9999-1000) + 1000
		app := fmt.Sprintf("ANDROID\t14.10.0\tAndroid OS\t13.0.%v", sort)
		mids := strings.Split(tok, ":")
		mid := mids[0]
		var ua = fmt.Sprintf("Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)")
		cl, botstate.Err := linetcr.CreateNewLogin(tok, no, mid, app, ua, botstate.HostName[0])
		if botstate.Err == nil {
			fmt.Println("\n\n  ❏ Name : " + cl.Namebot + "\n  ❏ Mid : " + cl.MID + "\n  ❏ Location : " + cl.Locale + "\n  ❏ Bots No: " + fmt.Sprintf("%v", no+1))
			if linetcr.IsFriends(cl, botstate.DEVELOPER[0]) == false {cl.FindAndAddContactsByMidV2(botstate.DEVELOPER[0]);time.Sleep(1 * time.Second)}
			cl.LoadPrimaryE2EEKeys()
			botstate.ClientMid[cl.MID] = cl
			err2 := cl.DisableE2ee()
			if err2 != nil {
				fmt.Println(err2)
			}
			aaa, err3 := cl.GetE2EEPublicKeys()
			if err3 == nil {
				for _, aa := range aaa {
					cl.RemoveE2EEPublicKey(aa)
				}
			} else {
				fmt.Println(err3)
			}
			r, _ := cl.GetHomeProfile(cl.MID)
			if linetcr.GetBannedChat(r) == 1 {
				linetcr.BanChatAdd(cl)
				fmt.Println("  ❏ Status : botstate.Banned")
			} else {
				fmt.Println("  ❏ Status : Normal")
			}
		} else {
			rs := botstate.Err.Error()
			if strings.Contains(rs, "INTERNAL_ERROR") || strings.Contains(rs, "AUTHENTICATION_FAILED") {
				fmt.Println("\n  ❏ Status : Limited" + "\n  ❏ Mid : " + mid + "\n  ❏ Bots No: " + fmt.Sprintf("%v", no+1))
				cl.MID = mid
				cl.Limited = true
			} else {
				logs := fmt.Sprintf("\n\n  ❏ No: %v ERROR: %s", no+1, botstate.Err)
				fmt.Println(logs)
			}
		}
	}
	for m := range linetcr.HashToMap(linetcr.GetBlock) {
		if !utils.InArrayString(botstate.Squadlist, m) {
			linetcr.GetBlock.Del(m)
		}
	}
	ch := make(chan int, len(botstate.ClientBot))
	if len(botstate.ClientBot) != 0 {
		acl := len(botstate.ClientBot)
		for x := 0; x < acl; x++ {
			cc := x
			cla := botstate.ClientBot[cc]
			runtime.Gosched()
			go RunBot(cla, ch)
		}
		list := append([]*linetcr.Account{}, botstate.ClientBot...)
		sort.Slice(list, func(i, j int) bool {
			return list[i].KickCount < list[j].KickCount
		})
		for i, cl := range list {
			kk := i * 30
			cl.KickPoint = kk
			ko := i * 10
			cl.CustomPoint = ko
		}
		handler.Resprem()
		for i := range botstate.ClientBot {
			for _, x := range botstate.Squadlist {
				if !utils.InArrayString(botstate.ClientBot[i].Squads, x) && x != botstate.ClientBot[i].MID {
					botstate.ClientBot[i].Squads = append(botstate.ClientBot[i].Squads, x)
				}
			}
		}
		botstate.ClientBot[0].SendMessage(botstate.DEVELOPER[0], botstate.ClientBot[0].Namebot)
		if botstate.Data.RestartBack != "" {
			_, memlist, _ := botstate.ClientBot[0].GetChatList(botstate.Data.RestartBack)
			if len(memlist) != 0 {
				for _, mid := range memlist {
					if utils.InArrayString(botstate.Squadlist, mid) {
						cl := handler.GetKorban(mid)
						if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
							cl.SendMessage(botstate.Data.RestartBack, "Done Reboot ♪♪\n")
							botstate.Data.RestartBack = ""
							break
						}
					}
				}
			}
			botstate.Data.RestartBack = ""
		}
		go func() {
			for {
				handler.AutoSet()
				time.Sleep(3 * time.Second)
			}
		}()
		for v := range ch {
			if v == 69 {
				break
			}
		}
		fmt.Println("HAVE FUN")
	}
}


//DONE






//DONE


//DONE

//DONE



//NEED FIX

//DONE




//DONE

//GechatLlinetcr




























//Llinetcr_Allbanlist

















var letters = []rune("123456789")
























































































//Backup_Qr




//Backup_124




//Llinetcr_Banlist
//Backup_133


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
					rngcmd := handler.GetComs(7, "invitebot")
					Group, user := op.Param1, op.Param2
					invited := strings.Split(op.Param3, "\x1e")
					Room := linetcr.GetRoom(Group)
					if utils.InArrayString(invited, client.MID) {
						if linetcr.IoGOBot(Group, client) {
							if utils.InArrayString(client.Squads, user) {
								go func(client *linetcr.Account, Group string){
									go handler.AccKickBan(client, Group)
								}(client, Group)
							} else if botstate.UserBot.GetBot(user) {
								var wg sync.WaitGroup
								wg.Add(1)
								go func(client *linetcr.Account, Group string){
									defer wg.Done()
									go handler.AccGroup(client, Group)
								}(client, Group)
								wg.Wait()
								if botstate.AutoPurge {
									go handler.JoinLlinetcrBan(client, Group)
								}
								if client.Limited == false {if !handler.InArrayInt64(botstate.CekGo, Optime) {botstate.CekGo = append(botstate.CekGo, Optime)
									handler.AcceptJoin(client, Group)}
								}
							} else if handler.CheckPermission(rngcmd, user, Group) {
								var wg sync.WaitGroup
								wg.Add(1)
								go func(client *linetcr.Account, Group string){
									defer wg.Done()
									go handler.AccGroup(client, Group)
								}(client, Group)
								wg.Wait()
								if botstate.AutoPurge {
									go handler.JoinLlinetcrBan(client, Group)
								}
								if client.Limited == false {if !handler.InArrayInt64(botstate.CekGo, Optime) {botstate.CekGo = append(botstate.CekGo, Optime)
									handler.AcceptJoin(client, Group)}
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
							if handler.MemUser(Group, user) {
								go func(client *linetcr.Account, Group string, user string) {
									if botstate.FilterWar.Cek(user) {
										handler.KickCancelProtect(client, user, invited[0], Group)
										botstate.Banned.AddBan(user)
										botstate.FilterWar.Del(user)
									}
								}(client, Group, user)
								if botstate.AutoPurge {
									go func(client *linetcr.Account, Group string, user string) {
										if botstate.FilterWar.Ceki(user) {
											botstate.Banned.AddBan(user)
											handler.KickProtect(client, Group, user)
											botstate.FilterWar.Deli(user)
										}
									}(client, Group, user)
									if botstate.FilterWar.Ceko(Optime) {
										Room.ListInvited = invited
										handler.BanAll(invited, Group)
										go handler.CancelBanInv(client, invited, Group)
									}
								}
							} else {
								if botstate.FilterWar.Ceko(Optime) {
									go handler.CancelAllCek(client, invited, Group)
								}
							}
						} else {
							if handler.MemBan(Group, user) {
								go func() {
									if botstate.FilterWar.Ceki(user) {
										handler.KickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
									}
								}()
								if botstate.FilterWar.Ceko(Optime) {
									handler.BanAll(invited, Group)
									go handler.CancelBanInv(client, invited, Group)
									go func() {handler.NodeBans(client, Group, invited)}()
								}
							} else {
								if handler.MemUser(Group, user) {
									go func() {
										if botstate.FilterWar.Ceki(user) {
											for _, vo := range invited {
												if handler.MemBan(Group, vo) {
													botstate.Banned.AddBan(user)
													handler.KickPelaku(client, Group, user)
													break
												}
											}
											botstate.FilterWar.Deli(user)
										}
									}()
									if botstate.FilterWar.Ceko(Optime) {
										go handler.CancelAllCek(client, invited, Group)
									}
								}
							}
						}
					}
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
					}
					if botstate.LogMode && handler.MemAccsess(op.Param2) {handler.NotifBot(client, Group, "invite", user)}
				}
				if op.Type == 133 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := linetcr.GetRoom(Group)
					if client.MID == Invited {
						linetcr.Gones(Group, client)
						if handler.MemUser(Group, user) {
							botstate.Banned.AddBan(user)
						}
					} else if !utils.InArrayString(Room.GoMid, client.MID) {
						if utils.InArrayString(client.Squads, Invited) {
							if handler.MemUser(Group, user) {
								if linetcr.IoGOBot(Group, client) {
									botstate.Banned.AddBan(user)
									go func() {
										if botstate.FilterWar.Cek(user) {
											handler.GroupBackupKick(client, Group, user, true)
											botstate.FilterWar.Del(user)
										}
									}()
									if botstate.FilterWar.Cek(Invited) {
										handler.GroupBackupInv(client, Group, Optime, Invited)
										botstate.FilterWar.Del(Invited)
									}
								}
							}
						} else {
							if !handler.MemUserN(Group, Invited) {
								if handler.CheckKickUser(Group, user, Invited) {
									handler.Back(Group, Invited)
									if botstate.FilterWar.Ceki(user) {
										handler.KickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
										if handler.MemUser(Group, user) {
											botstate.Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProKick {
									if handler.MemUser(Group, user) {
										if Room.Backup {
											handler.Back(Group, Invited)
										}
										if _, ok := botstate.Nkick.Get(user); !ok {
											botstate.Nkick.Set(user, 1)
											handler.KickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)
										}
                                                                   ct, _ := client.GetContact(Invited)
                                                                   if ct != nil {if ct.CapableBuddy {handler.KickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)}
										}
									}
								}
							}
						}
					} else {
						if handler.MemUser(Group, Invited) {
							if handler.MemUser(Group, user) {
								handler.Back(Group, Invited)
								botstate.Banned.AddBan(user)
							}
						}
						if handler.MemUser(Group, user) {
							if botstate.FilterWar.Ceki(user) {
								handler.GhostEnd(client, Group, Optime, user, true)
								botstate.FilterWar.Deli(user)
							}
						}
					}
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
					}
					if botstate.LogMode && handler.MemAccsess(op.Param2) {handler.NotifBot(client, Group, "kick", user)}
				}
				if op.Type == 129 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Group, user := op.Param1, op.Param2
					if botstate.PowerMode == true {
					if handler.MemBan(Group, user) {
						go func() {
							if botstate.FilterWar.Ceki(user) {
								handler.AccKickBan(client, Group)
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
							if handler.MemUser(Group, user) {
								if botstate.FilterWar.Ceki(user) {
										botstate.Banned.AddBan(user)
									handler.KickPelaku(client, Group, user)
									botstate.FilterWar.Deli(user)
								}
							}
						} else {
							if handler.MemBan(Group, user) {
								if handler.MemUser(Group, user) {
									if botstate.FilterWar.Ceki(user) {
										botstate.Banned.AddBan(user)
										handler.KickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
									}
								}
							} else {
								if utils.InArrayString(Room.ListInvited, user) {
									if handler.MemUser(Group, user) {
										if handler.CekJoin(user) {
											handler.KickPelaku(client, Group, user)
											handler.DelJoin(user)
											Room.ListInvited = utils.RemoveString(Room.ListInvited, user)
										}
									} else {
										Room.ListInvited = utils.RemoveString(Room.ListInvited, user)
									}
								} else {
									if Room.Welcome {
										if _, ok := botstate.Cewel.Get(user); !ok {
											botstate.Cewel.Set(user, 1)
											if handler.CekJoin(user) {
												if !utils.InArrayString(botstate.Squadlist, user) {
													Room.WelsomeSet(client, Group, user)
												}
											}
										}
									} else {
										if botstate.LockMode == true {
											if handler.MemUser(Group, user) {
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
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
					}
				}
				if op.Type == 122 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Group, user, invited := op.Param1, op.Param2, op.Param3
					Optime := op.CreatedTime
					Room := linetcr.GetRoom(Group)
					if client.Limited == false && linetcr.IoGOBot(Group, client) {
						if handler.MemUser(Group, user) {
							if Room.ProQr || botstate.AutoproN == true {
								if invited == "4" {
									if handler.CekOp2(Optime) {
										go func() {
											cans := linetcr.Actor(Group)
											for _, cl := range cans {
												if botstate.Err == nil {
													break
												}
											}
										}()
										if botstate.FilterWar.Ceki(user) {
											handler.KickPelaku(client, Group, user)
											botstate.FilterWar.Deli(user)
											botstate.Banned.AddBan(user)
										}
									}
								}
							} else if botstate.KickBanQr == true {
								if invited == "4" {
									if handler.CekOp2(Optime) {go func() {cans := linetcr.Actor(Group);for _, cl := range cans {botstate.Err := cl.UpdateChatQrV2(Group, true);if botstate.Err == nil {break}}}()
										if botstate.FilterWar.Ceki(user) {handler.KickPelaku(client, Group, user);botstate.FilterWar.Deli(user);botstate.Banned.AddBan(user)}
									}
								}
							} else if Room.ProPicture || botstate.AutoproN == true {
								if invited == "4" {
									if handler.CekOp2(Optime) {
										if botstate.FilterWar.Ceki(user) {
											handler.KickPelaku(client, Group, user)
											botstate.FilterWar.Deli(user)
										}
									}
								}
							} else if Room.ProName || botstate.AutoproN == true {
								if invited == "4" {
									if handler.CekOp2(Optime) {
										go func() {
											cans := linetcr.Actor(Group)
											for _, cl := range cans {
												if botstate.Err == nil {
													break
												}
											}
										}()
										if botstate.FilterWar.Ceki(user) {
											handler.KickPelaku(client, Group, user)
											botstate.FilterWar.Deli(user)
										}
									}
								}
							} else {
								if handler.MemBan(Group, user) {
									if invited == "4" {
										if handler.CekOp2(Optime) {
											go func() {
												cans := linetcr.Actor(Group)
												for _, cl := range cans {
													if botstate.Err == nil {
														break
													}
												}
											}()
											if botstate.FilterWar.Ceki(user) {
												handler.KickPelaku(client, Group, user)
												botstate.FilterWar.Deli(user)
												botstate.Banned.AddBan(user)
											}
										}
									} else if invited == "4" {
										if handler.CekOp2(Optime) {
											go func() {
												cans := linetcr.Actor(Group)
												for _, cl := range cans {
													if botstate.Err == nil {
														break
													}
												}
											}()
											if botstate.FilterWar.Ceki(user) {
												handler.KickPelaku(client, Group, user)
												botstate.FilterWar.Deli(user)
											}
										}
									}
								}
							}

						}
					}
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
					}
				}
				if op.Type == 126 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Optime := op.CreatedTime
					Group, user, Invited := op.Param1, op.Param2, op.Param3
					Room := linetcr.GetRoom(Group)
					if client.MID == Invited {
						linetcr.Gones(Group, client)
						if handler.MemUser(Group, user) {
							botstate.Banned.AddBan(user)
						}
					} else if !utils.InArrayString(Room.GoMid, client.MID) {
						if utils.InArrayString(client.Squads, Invited) {
							if handler.MemUser(Group, user) {
								if linetcr.IoGOBot(Group, client) {
									botstate.Banned.AddBan(user)
									go func() {
										if botstate.FilterWar.Cek(user) {
											handler.GroupBackupCans(client, Group, user, true)
											botstate.FilterWar.Del(user)
										}
									}()
									if botstate.FilterWar.Cek(Invited) {
										handler.GroupBackupInv(client, Group, Optime, Invited)
										botstate.FilterWar.Del(Invited)
									}
								}
							}
						} else {
							if !handler.MemUserN(Group, Invited) {
								if handler.CheckKickUser(Group, user, Invited) {
									handler.Back(Group, Invited)
									if botstate.FilterWar.Ceki(user) {
										handler.KickPelaku(client, Group, user)
										botstate.FilterWar.Deli(user)
										if handler.MemUser(Group, user) {
											botstate.Banned.AddBan(user)
										}
									}
								}
							} else {
								if Room.ProCancel {
									if handler.MemUser(Group, user) {
										if Room.Backup {
											handler.Back(Group, Invited)
										}
										if _, ok := botstate.Nkick.Get(user); !ok {
											botstate.Nkick.Set(user, 1)
											handler.KickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)
										}
                                                                   ct, _ := client.GetContact(Invited)
                                                                   if ct != nil {if ct.CapableBuddy {handler.KickPelaku(client, Group, user)
											botstate.Banned.AddBan(user)}
										}
									}
								}
							}
						}
					} else {
						if handler.MemUser(Group, Invited) {
							if handler.MemUser(Group, user) {
								handler.Back(Group, Invited)
								botstate.Banned.AddBan(user)
							}
						}
						if handler.MemUser(Group, user) {
							if botstate.FilterWar.Ceki(user) {
								handler.GhostEnd(client, Group, Optime, user, true)
								botstate.FilterWar.Deli(user)
							}
						}
					}
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
					}
					if botstate.LogMode && handler.MemAccsess(op.Param2) {handler.NotifBot(client, Group, "cancel", user)}
				} else if op.Type == 33 {
					Group := op.Param1
					if botstate.LogMode && !utils.InArrayString(client.Squads, Group) {handler.LogAccess(client, client.Namebot, Group, "deleteaccount", []string{}, 2)}
				} else if op.Type == 5 {
					Group := op.Param1
					if botstate.LogMode && !utils.InArrayString(client.Squads, Group) {handler.LogAccess(client, client.Namebot, Group, "addfrind", []string{}, 2)}
				} else if op.Type == 50 {
					Group := op.Param1
					if botstate.LogMode && !utils.InArrayString(client.Squads, Group) {handler.LogAccess(client, client.Namebot, Group, "callme", []string{}, 2)}
				} else if op.Type == 55 {
					Optime := op.CreatedTime
					Group, user := op.Param1, op.Param2
					if client.Limited == false && linetcr.IoGOBot(Group, client) {
						if handler.CekOp(Optime) {
							if handler.MemBan(Group, user) {
								handler.KickPelaku(client, Group, user)
							} else {
								Room := linetcr.GetRoom(Group)
								if Room.Lurk && !utils.InArrayString(botstate.CheckHaid, user) {
									Room.CheckLurk(client, Group, user)
								}
							}
						}
					}
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
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
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
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
									if !handler.MemBan(Group, user) && !utils.InArrayString(botstate.Squadlist, user) && !botstate.UserBot.GetBot(user) && !utils.InArrayString(Room.GoMid, user) {
										handler.Hstg(Group, user)
										Room.Leave = time.Now()
									}
								}
							}
						} else {
							if Room.Leavebool {
								if _, ok := botstate.Cleave.Get(user); !ok {
									botstate.Cleave.Set(user, 1)
									if !handler.MemBan(Group, user) && !utils.InArrayString(botstate.Squadlist, user) && !botstate.UserBot.GetBot(user) && !utils.InArrayString(Room.GoMid, user) {
										Room.LeaveSet(client, Group, user)
									}
								}
							}
						}
					}
					if _, ok := botstate.FilterMsg.Get(Optime); !ok {
						botstate.FilterMsg.Set(Optime, client)
						handler.LogOp(op, client)
						handler.LogGet(op)
					}
				} else if op.Type == 30 {
					Group := op.Param1
					Room := linetcr.GetRoom(Group)
					if Room.Announce && linetcr.IoGOBot(Group, client) {
						Optime := op.CreatedTime
						if handler.CekOp(Optime) {
							Room.CheckAnnounce(client, Group)
						}
					}
				} else if op.Type == 123 {
					client.CInvite()
				} else if op.Type == 132 {
					runtime.GOMAXPROCS(botstate.Cpu)
					Group, user := op.Param1, op.Param2
					if botstate.PowerMode == true {
					if handler.MemBan(Group, user) {
						go func() {
							if botstate.FilterWar.Ceki(user) {
								handler.KickBan132(client, Group, user)
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
					if handler.MemBan(Group, user) {
						go func() {
							if botstate.FilterWar.Ceko(Optime) {
								handler.BanAll(invited, Group)
								go handler.CancelBan125(client, Group, invited)
							}
						}()
					}}
					if Room.ProInvite {
						if handler.MemUser(Group, user) {
							go func(client *linetcr.Account, Group string, user string) {
								if botstate.FilterWar.Cek(user) {
									handler.KickCancelProtect(client, user, invited[0], Group)
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






////NEW















































// here func notif SELFTCR™





func Bot(op *SyncService.Operation, client *linetcr.Account, ch chan int) {
	defer utils.PanicHandle("Bot")
	op.Message = client.DecryptE2EEMessage(op.Message)
	msg := op.Message
	if msg.ToType != 2 {
		return
	}
	if _, ok := botstate.Commandss.Get(op.CreatedTime); ok {
		return
	} else {
		botstate.Commandss.Set(op.CreatedTime, client)
	}
	if time.Now().Sub(botstate.Timeabort) >= 60*time.Second {
		handler.Abort()
	}
       if botstate.AutoBackBot {
		time.AfterFunc(time.Duration(botstate.Timebk)*time.Second, func() {  
			//handler.GroupBackupInv2(client, op.Message.To)
			all := []string{}
			handler.GetSquad(client, op.Message.To)
			room := linetcr.GetRoom(op.Message.To)
			cuk := room.Client
			alls := []*linetcr.Account{}
			for _, x := range botstate.ClientBot {
				if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
					alls = append(alls, x)
				}
			}
			for _, x := range botstate.ClientBot {
			    if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
					all = append(all, x.MID)
				} else {
					break
				}
			}
			ClAct := len(alls)
			if ClAct != 0 {
				sort.Slice(alls, func(i, j int) bool {
					return alls[i].KickPoint < alls[j].KickPoint
				})
				for _, a := range all {
					alls[0].InviteIntoGroupNormal(op.Message.To, []string{a})
	            			time.Sleep(100 * time.Millisecond)
					handler.GetSquad(client, op.Message.To)
				}
			}
		})
	}
	Rname := botstate.MsRname
	Sname := botstate.MsSname
	sender := op.Message.From_
	text := op.Message.Text
	receiver := op.Message.To
	var pesan = strings.ToLower(text)
	var to string
	mentions := mentions{}
	if op.Message.ToType == 0 {
		to = sender
	} else {
		to = receiver
	}
       if botstate.DetectCall && msg.ToType == 2 {
		Room := linetcr.GetRoom(to)
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
	      Date := dd + "-" + MM + "-" + yyyy
	      Time := hhconv + ":" + mmconv + ":" + ssconv
		pr, _ := client.GetContact(sender)
              if msg.ContentMetadata["GC_MEDIA_TYPE"] == "AUDIO" || msg.ContentMetadata["GC_EVT_TYPE"] == "S" {
                   a := "「 DETECT 」"
                   a += "\n › Type: Callgroup"
                   a += "\n    • Type: AUDIO"
                   a += "\n    • Date: "+Date
                   a += "\n    • Time: "+Time
                   a += fmt.Sprintf("\n    • Group: %s", Room.Name)
		      a += "\n    • Host: "+pr.DisplayName
                   client.SendMessage(to, botstate.Fancy(a))
		} else {
                   if msg.ContentMetadata["GC_MEDIA_TYPE"] == "VIDEO" || msg.ContentMetadata["GC_EVT_TYPE"] == "S" {
                        a := "「 DETECT 」"
                        a += "\n › Type: Callgroup"
                        a += "\n    • Type: VIDEO"
                        a += "\n    • Date: "+Date
                        a += "\n    • Time: "+Time
                        a += fmt.Sprintf("\n    • Group: %s", Room.Name)
		           a += "\n    • Host: "+pr.DisplayName
                        client.SendMessage(to, botstate.Fancy(a))
                  }
             }
       }
	if len(botstate.Sinderremote) != 0 {
		if utils.InArrayString(botstate.Sinderremote, sender) {
			if botstate.Remotegrupid != "" {
				to = botstate.Remotegrupid
			}
		}
	}
	mentionlist := []string{}
	json.Unmarshal([]byte(op.Message.ContentMetadata["MENTION"]), &mentions)
	for _, mention := range mentions.MENTIONEES {
		if !utils.InArrayString(mentionlist, mention.Mid) {
			mentionlist = append(mentionlist, mention.Mid)
		}
	}
	var Rplay = ""
	var room *linetcr.LineRoom
	var bks = []*linetcr.Account{}
	room = linetcr.GetRoom(to)
	bks = room.Client
	if len(bks) == 0 {
		handler.GetSquad(client, to)
		room = linetcr.GetRoom(to)
		bks = room.Client
	}
	sort.Slice(room.Ava, func(i, j int) bool {
		return room.Ava[i].Client.KickPoint < room.Ava[j].Client.KickPoint
	})
	bk := []*linetcr.Account{}
	bk2 := []*linetcr.Account{}
	for _, n := range bks {
		bk = append(bk, n)
		if !n.Limited {
			bk2 = append(bk2, n)
		}
	}
	clen := len(bk2)
	if clen != 0 {
		client = bk2[0]
		room.Exe = bk2[0]
		room.Limit = false
	} else {
		room.Limit = true
	}
	if room.AntiTag && handler.MemUser(to, msg.From_) && len(mentionlist) != 0 && !room.Automute {
		if room.Limit {
			client.SendMessage(to, botstate.Fancy("All bot in here banned, please try invite another bot"))
			return
		}
		if client.Limited == false {
			client.DeleteOtherFromChats(to, []string{msg.From_})
		} else {
			for _, bot := range bk {
				if bot.Limited == false {
					bot.DeleteOtherFromChats(to, []string{msg.From_})
					break
				}
			}
		}
	}
	if op.Message.RelatedMessageId != "" && len(mentionlist) == 0 {
		asu, _ := client.GetRecentMessagesV2(op.Message.To)
		for _, xx := range asu {
			if xx.ID == op.Message.RelatedMessageId {
				Rplay = xx.From_
				break
			}
		}
	}
	emots := emots{}
	json.Unmarshal([]byte(op.Message.ContentMetadata["REPLACE"]), &emots)
	for _, stiker := range emots.STICON.RESOURCES {
		if handler.CheckPermission(5, sender, to) {
		if !handler.MemUser(to, msg.From_) {
			if botstate.GetStickerKick == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("kick by emote updated"))
				}
			} else if botstate.GetStickerRespon == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("respon by emote updated"))
				}
			} else if botstate.GetStickerStayall == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("stayall by emote updated"))
				}
			} else if botstate.GetStickerLeave == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("leave by emote updated"))
				}
			} else if botstate.GetStickerKickall == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("kickall by emote updated"))
				}
			} else if botstate.GetStickerBypass == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("bypass by emote updated"))
				}
			} else if botstate.GetStickerInvite == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("invite by emote updated"))
				}
			} else if botstate.GetStickerClearban == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("clearban by emote updated"))
				}
			} else if botstate.GetStickerCancelall == 1 {
				if !handler.MemUser(to, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("cancelall by emote updated"))
				}
			} else if stiker.PRODUCTID == botstate.Stkid && stiker.STICONID == botstate.Stkpkgid {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if handler.MemUser(to, lists[i]) && !utils.InArrayString(listuser, lists[i]) {
							if botstate.AutoBan {
										botstate.Banned.AddBan(lists[i])
							}
							listuser = append(listuser, lists[i])
						}
					}
				}
				fmt.Println(listuser)
				if len(listuser) != 0 {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, listuser)
						handler.LogAccess(client, to, sender, "kick", listuser, msg.ToType)
					}
				}
			} else if stiker.PRODUCTID == botstate.Stkid2 && stiker.STICONID == botstate.Stkpkgid2 {
				for _, p := range bk {
					go p.SendMessage(to, botstate.Fancy(botstate.MsgRespon))
				}
			} else if stiker.PRODUCTID == botstate.Stkid3 && stiker.STICONID == botstate.Stkpkgid3 {
				if room.Limit {
					client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
					return
				}
				numb := len(botstate.ClientBot)
				if numb > 0 && numb <= len(botstate.ClientBot) {
					handler.GetSquad(client, to)
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
						handler.GetSquad(client, to)
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
							wi := handler.GetSquad(client, to)
							for i := 0; i < len(all); i++ {
								if i == g {
									break
								}
								l := all[i]
								if l != client && !linetcr.InArrayCl(wi, l) {
									if !l.Limited {
										wg.Add(1)
										go func() {
											l.AcceptTicket(to, ti)
											wg.Done()
										}()
									}
								}
							}
							wg.Wait()
							client.UpdateChatQrV2(to, true)
							handler.GetSquad(client, to)
							handler.LogAccess(client, to, sender, "bringbot", []string{}, 2)
							handler.SaveBackup()
							aa := len(room.Client)
							var name string
							name = fmt.Sprintf("Ready %v bots here", aa)
							client.SendMessage(to, botstate.Fancy(name))
							//newsend += name + "\n"
						}
					}
				}
			} else if stiker.PRODUCTID == botstate.Stkid4 && stiker.STICONID == botstate.Stkpkgid4 {
				_, mem := client.GetGroupInvitation(to)
				anu := []string{}
				for m := range mem {
					if utils.InArrayString(botstate.Squadlist, m) {
						anu = append(anu, m)
					}
				}
				if len(anu) != 0 {
					for _, mid := range anu {
						cl := handler.GetKorban(mid)
						cl.AcceptGroupInvitationNormal(to)
					}
				}
				handler.GetSquad(client, to)
				room := linetcr.GetRoom(to)
				bk = room.Client
				for _, cl := range bk {
					go cl.LeaveGroup(to)
				}
				if botstate.LogGroup == to {
				}
				linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
				handler.SaveBackup()
				handler.LogAccess(client, to, sender, "leave", []string{}, msg.ToType)
			} else if stiker.PRODUCTID == botstate.Stkid5 && stiker.STICONID == botstate.Stkpkgid5 {
				_, memlist, _ := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist {
					if utils.InArrayString(botstate.Squadlist, mid) {
						cl := handler.GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				max := len(exe) * 100
				lkick := []string{}
				for n, v := range memlist {
					if handler.MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0
				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					go func(to string, i int) {
						target := lkick[i]
						cl := nom[i]
						cl.DeleteOtherFromChats(to, []string{target})
					}(to, i)
				}
				handler.LogAccess(client, to, sender, "kickall", lkick, msg.ToType)
			} else if stiker.PRODUCTID == botstate.Stkid7 && stiker.STICONID == botstate.Stkpkgid7 {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if !utils.InArrayString(listuser, lists[i]) {
							listuser = append(listuser, lists[i])
						}
					}
				}
				if len(listuser) != 0 {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						lists := handler.Setinvitetomsg(exe, to, listuser)
						if len(lists) != 0 {
									handler.Cekbanwhois(client, to, lists)
						}
						handler.LogAccess(client, to, sender, "invite", listuser, msg.ToType)
					} else {
						client.SendMessage(to, botstate.Fancy("Please add another bot that has a ban invite."))
					}
				}
			} else if stiker.PRODUCTID == botstate.Stkid8 && stiker.STICONID == botstate.Stkpkgid8 {
				if len(botstate.Banned.Banlist) != 0 {
					msgcbn := fmt.Sprintf(botstate.MsgBan, len(botstate.Banned.Banlist))
					handler.LogAccess(client, to, sender, "clearban", botstate.Banned.Banlist, msg.ToType)
					client.SendMessage(to, botstate.Fancy(msgcbn))
					botstate.Banned.Banlist = []string{}
					botstate.Banned.Exlist = []string{}
				} else {
					client.SendMessage(to, botstate.Fancy("Ban list is empty."))
				}
			} else if stiker.PRODUCTID == botstate.Stkid9 && stiker.STICONID == botstate.Stkpkgid9 {
				_, memlist2, memlist := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist2 {
					if utils.InArrayString(botstate.Squadlist, mid) {
						cl := handler.GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				lkick := []string{}
				max := len(exe) * 10
				for n, v := range memlist {
					if handler.MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0

				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					target := lkick[i]
					cl := nom[i]
					ants.Submit(func() { cl.CancelChatInvitations(to, []string{target}) })
				}
				handler.LogAccess(client, to, sender, "cancelall", lkick, msg.ToType)
			}
		}
	}}
	if botstate.AutoTranslate {
		if handler.MemUser(to, sender) {
			if strings.Contains(pesan, pesan) {
				client.TranslateMe("ID", pesan)
				filepath := fmt.Sprintf("trMe.txt")
				b, botstate.Err := ioutil.ReadFile(filepath)
				if botstate.Err != nil {
					fmt.Print(botstate.Err)
				}
				code := string(b)
				list := code
				client.SendMessage(to, botstate.Fancy(list))
			}
		}
	}
	if botstate.AutoTranslate {
		if !handler.MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				Type := botstate.TypeTrans
				client.TranslateYou(Type, pesan)
				filepath := fmt.Sprintf("trYou.txt")
				b, botstate.Err := ioutil.ReadFile(filepath)
				if botstate.Err != nil {
					fmt.Print(botstate.Err)
				}
				code := string(b)
				list := code
				client.SendMessage(to, botstate.Fancy(list))
			}
		}
	}
	if botstate.RemoteOwner {
		if !handler.MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := handler.StripOut(pesan)
				if botstate.Err != nil {
					client.SendMessage(to, botstate.Fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, botstate.MidRemote)
					listuser := []string{}
					x := 2
					if len(simpan) != 0 {
						for i := range simpan {
							if !utils.InArrayString(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
						handler.CheckListAccess(client, to, listuser, x, sender)
					}
				}
			}
		}
	}
	if botstate.RemoteMaster {
		if !handler.MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := handler.StripOut(pesan)
				if botstate.Err != nil {
					client.SendMessage(to, botstate.Fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, botstate.MidRemote)
					listuser := []string{}
					x := 3
					if len(simpan) != 0 {
						for i := range simpan {
							if !utils.InArrayString(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
						handler.CheckListAccess(client, to, listuser, x, sender)
					}
				}
			}
		}
	}
	if botstate.RemoteAdmin {
		if !handler.MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := handler.StripOut(pesan)
				if botstate.Err != nil {
					client.SendMessage(to, botstate.Fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, botstate.MidRemote)
					listuser := []string{}
					x := 4
					if len(simpan) != 0 {
						for i := range simpan {
							if !utils.InArrayString(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
						handler.CheckListAccess(client, to, listuser, x, sender)
					}
				}
			}
		}
	}
	if botstate.RemoteContact {
		if !handler.MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := handler.StripOut(pesan)
				if botstate.Err != nil {
					client.SendMessage(to, botstate.Fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, botstate.MidRemote)
					listuser := []string{}
					if len(simpan) != 0 {
						for i := range simpan {
							if !utils.InArrayString(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
					}
					if len(listuser) != 0 {
						for _, mek := range listuser {
							client.FindAndAddContactsByMidV2(mek)
							go client.SendContact(to, mek)
						}
					}
				}
			}
		}
	}
	if botstate.RemoteBan {
		if !handler.MemUserN(to, sender) {
			if strings.Contains(pesan, pesan) {
				st := handler.StripOut(pesan)
				if botstate.Err != nil {
					client.SendMessage(to, botstate.Fancy("Please put a number"))
					return
				} else {
					simpan := linetcr.Archimed(st, botstate.MidRemote)
					listuser := []string{}
					if len(simpan) != 0 {
						for i := range simpan {
							if !utils.InArrayString(listuser, simpan[i]) {
								listuser = append(listuser, simpan[i])
							}
						}
					}
					client.SendPollMention(to, "Added banlist member:", listuser)
					if len(listuser) != 0 {
						for _, mek := range listuser {
							botstate.Banned.AddBan(mek)
						}
					}
				}
			}
		}
	}
//botstate.AutoJointicket
	if botstate.AutoJointicket {
		if strings.Contains(pesan, "/ti/g") {
			regex, _ := regexp.Compile(`(?:line\:\/|line\.me\/R)\/ti\/g\/([a-zA-Z0-9_-]+)?`)
			links := regex.FindAllString(msg.Text, -1)
			tickets := []string{}
			for _, link := range links {
				if !utils.InArrayString(tickets, link) {
					tickets = append(tickets, link)
				}
			}
			for _, tick := range tickets {
				tuk := strings.Split(tick, "/")
				ntk := len(tuk) - 1
				ti := tuk[ntk]
				fmt.Println(ti)
				tkt := client.FindChatByTicket(ti)
				client.AcceptTicket(tkt.Chat.ChatMid, ti)
				exe := []*linetcr.Account{}
				for _, p := range bk {
					if p.Limited == false {
						if botstate.Err == nil {
							exe = append(exe, p)
						}
					}
				}
				newsend := ""
				if len(exe) != 0 {
					if botstate.TypeJoin == "normal" {
						newsend += "Succes Accept Group Ticket"
					} else if botstate.TypeJoin == "nuke" {
						go handler.NukJoin(exe[0], op.CreatedTime, tkt.Chat.ChatMid)
						newsend += "Succes Accept Ticket Nuke Group"
					}
				}
			}
		}
	}
//BOMLIKE
       if botstate.BomLike && !handler.MemUser(to, sender) && op.Message.ContentType == 16 {
       	posturl := msg.ContentMetadata["postEndUrl"]
       	if msg.ContentMetadata["serviceType"] == "MH" {
             		mids := strings.Replace(posturl, "https://line.me/R/home/post?userMid=", "", 1)
                    post := strings.Split(mids, "&postId=")
                    client.Bomlike(to, mids, post[1])
		}
	}
//TIMELINE
       if botstate.AutoLike && op.Message.ContentType == 16 {
       	posturl := msg.ContentMetadata["postEndUrl"]
       	if msg.ContentMetadata["serviceType"] == "MH" {
             		mids := strings.Replace(posturl, "https://line.me/R/home/post?userMid=", "", 1)
                    post := strings.Split(mids, "&postId=")
                    client.Timeline(to, mids, post[1],"likepost",)
		}
	}
//MEDIA_DL
	if botstate.MediaDl {
		if strings.Contains(pesan, "tiktok.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"tiktok")
			}
		}
		if strings.Contains(pesan, "instagram.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"instagram")
			}
		}
		if strings.Contains(pesan, "smule.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"smule")
			}
		}
		if strings.Contains(pesan, "sck.io/p/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"snackvideo")
			}
		}
		if strings.Contains(pesan, "facebook.com/") {
			regex, _ := regexp.Compile(`(https?://\S+)`)
			links := regex.FindAllString(msg.Text, -1)
			for _, link := range links {
				client.Sendmediadl(to, link,"facebook")
			}
		}
	}
//BATAS
	if room.ProLink && handler.MemUser(to, msg.From_) && strings.Contains(pesan, "https://") {
		client.DeleteOtherFromChats(to, []string{msg.From_})
		botstate.Banned.AddBan(msg.From_)
	}
	if room.AntiTag && handler.MemUser(to, msg.From_) && strings.Contains(pesan, "@All") {
		client.DeleteOtherFromChats(to, []string{msg.From_})
		botstate.Banned.AddBan(msg.From_)
	}
	if op.Message.ContentType == 22 {
	      if room.ProFlex && handler.MemUser(to, msg.From_) {
		       client.DeleteOtherFromChats(to, []string{msg.From_})
		       botstate.Banned.AddBan(msg.From_)
                    if msg.ContentMetadata["FLEX_VER"] == "u1" || msg.ContentMetadata["ORGCONTP"] == "FLEX" {
		              client.DeleteOtherFromChats(to, []string{msg.From_})
		              botstate.Banned.AddBan(msg.From_)
	             }
	       }
	}
	if op.Message.ContentType == 1 {
		if room.ProImage && handler.MemUser(to, sender) {
		       client.DeleteOtherFromChats(to, []string{sender})
		       botstate.Banned.AddBan(sender)
	       }
	}
	if op.Message.ContentType == 2 {
		if room.ProVideo && handler.MemUser(to, sender) {
		       client.DeleteOtherFromChats(to, []string{sender})
		       botstate.Banned.AddBan(sender)
	       }
	}
//Broadcast_VIDEO_IMAGE
	if botstate.BcImage {
	}
	if botstate.GBcImage {
	}
	if botstate.FBcImage {
	}
	if botstate.SAVEBcImage {
	}
	if op.Message.ContentType == 1 {
		if botstate.StartBc {
			if !handler.MemUser(to, sender) {
				if botstate.BcImage {
					nCount := 0
					typec := "lmid"
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					handler.AddConSingle(lists)
					if len(lists) != 0 {
						for _, i := range lists {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendFoto(i, path)
						}
					}
					client.SendMessage(to, botstate.Fancy("Success broadcast image to mid"))
				}
			}
		}
		if botstate.GStartBc {
			if !handler.MemUserN(to, sender) {
				if botstate.GBcImage {
					gr, _ := client.GetGroupIdsJoined()
					for _, gi := range gr {
						path, _ := client.DownloadObjectMsg(msg.ID)
						client.SendFoto(gi, path)
					}
					client.SendMessage(to, botstate.Fancy("Success all groupcast image"))
				}
			}
		}
		if botstate.FStartBc {
			if !handler.MemUserN(to, sender) {
				if botstate.FBcImage {
					for _, cl := range botstate.ClientBot {
						fl, _ := cl.GetAllContactIds()
						for _, x := range fl {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendFoto(x, path)
							time.Sleep(3 * time.Second)
						}	
					}
					client.SendMessage(to, botstate.Fancy("Success all friendcast image"))
				}
			}
		}
		if botstate.StartSaveBc {
			if !handler.MemUserN(to, sender) {
				if botstate.SAVEBcImage {
					client.Downloadbc(msg.ID)
					client.SendMessage(to, botstate.Fancy("Success save image broadcast"))
                           	time.Sleep(3 * time.Second)
					client.SendMessage(to, botstate.Fancy("Broadcast all group runs every 5 minutes"))
				}
			}
		}
	}
	if botstate.BcVideo {
	}
	if botstate.GBcVideo {
	}
	if botstate.FBcVideo {
	}
	if op.Message.ContentType == 2 {
		if botstate.StartBcV {
			if !handler.MemUser(to, sender) {
				if botstate.BcVideo {
					nCount := 0
					typec := "lmid"
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					handler.AddConSingle(lists)
					if len(lists) != 0 {
						for _, i := range lists {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendVid(i, path)
						}
					}
					client.SendMessage(to, botstate.Fancy("Success broadcast video to mid"))
				}
			}
		}
		if botstate.GStartBcV {
			if !handler.MemUserN(to, sender) {
				if botstate.GBcVideo {
					gr, _ := client.GetGroupIdsJoined()
					for _, gi := range gr {
						path, _ := client.DownloadObjectMsg(msg.ID)
						client.SendVid(gi, path)
					}
					client.SendMessage(to, botstate.Fancy("Success all groupcast video"))
				}
			}
		}
		if botstate.FStartBcV {
			if !handler.MemUserN(to, sender) {
				if botstate.FBcVideo {
					for _, cl := range botstate.ClientBot {
						fl, _ := cl.GetAllContactIds()
						for _, x := range fl {
							path, _ := client.DownloadObjectMsg(msg.ID)
							client.SendVid(x, path)
							time.Sleep(3 * time.Second)
						}	
					}
					client.SendMessage(to, botstate.Fancy("Success all friendcast video"))
				}
			}
		}
	}
//Changepicture
	if botstate.ChangPict && !botstate.AllCheng && !botstate.StartChangeImg {
		if len(mentionlist) != 0 {
			for _, ym := range mentionlist {
				if utils.InArrayString(botstate.Squadlist, ym) {
					cl := handler.GetKorban(ym)
					if !linetcr.Checkarri(botstate.Changepic, cl) {
					}
				}
			}
			if len(botstate.Changepic) != 0 {
				client.SendMessage(to, botstate.Fancy("Please Send Your Image !!!"))
			}
		}
	} else if botstate.ChangCover && !botstate.AllCheng && !botstate.StartChangeImg {
		if len(mentionlist) != 0 {
			for _, ym := range mentionlist {
				if utils.InArrayString(botstate.Squadlist, ym) {
					cl := handler.GetKorban(ym)
					if !linetcr.Checkarri(botstate.Changepic, cl) {
					}
				}
			}
			if len(botstate.Changepic) != 0 {
				client.SendMessage(to, botstate.Fancy("Please Send Your Image !!!"))
			}
		}
	} else if botstate.ChangVpict && !botstate.AllCheng && !botstate.StartChangeImg {
		if len(mentionlist) != 0 {
			for _, ym := range mentionlist {
				if utils.InArrayString(botstate.Squadlist, ym) {
					cl := handler.GetKorban(ym)
					if !linetcr.Checkarri(botstate.Changepic, cl) {
					}
				}
			}
			if len(botstate.Changepic) != 0 {
				client.SendMessage(to, botstate.Fancy("Please Send Your Video !!!"))
			}
		}
	} else if botstate.ChangVcover && !botstate.AllCheng && !botstate.StartChangeImg {
		if len(mentionlist) != 0 {
			for _, ym := range mentionlist {
				if utils.InArrayString(botstate.Squadlist, ym) {
					cl := handler.GetKorban(ym)
					if !linetcr.Checkarri(botstate.Changepic, cl) {
					}
				}
			}
			if len(botstate.Changepic) != 0 {
				client.SendMessage(to, botstate.Fancy("Please Send Your Video !!!"))
			}
		}
	} else if botstate.ChangName {
		if len(mentionlist) != 0 {
			for _, ym := range mentionlist {
				if utils.InArrayString(botstate.Squadlist, ym) {
					cl := handler.GetKorban(ym)
					if !linetcr.Checkarri(botstate.Changepic, cl) {
					}
				}
			}
			if len(botstate.Changepic) != 0 {
				if botstate.MsgName != "" {
					for i := range botstate.Changepic {
						if handler.TimeDown(i) {
							star := botstate.MsgName
							botstate.Changepic[i].UpdateProfileName(star)
							botstate.Changepic[i].SendMessage(to, botstate.Fancy("Profile name success updated."))
						}
					}
				} else {
					client.SendMessage(to, botstate.Fancy("Add name first."))
				}
			}
		}
	} else if botstate.ChangeBio {
		if len(mentionlist) != 0 {
			for _, ym := range mentionlist {
				if utils.InArrayString(botstate.Squadlist, ym) {
					cl := handler.GetKorban(ym)
					if !linetcr.Checkarri(botstate.Changepic, cl) {
					}
				}
			}
			if len(botstate.Changepic) != 0 {
				if botstate.MsgBio != "" {
					for i := range botstate.Changepic {
						if handler.TimeDown(i) {
							star := botstate.MsgBio
							botstate.Changepic[i].UpdateProfileBio(star)
							botstate.Changepic[i].SendMessage(to, botstate.Fancy("Profile status success updated."))
						}
					}
				} else {
					client.SendMessage(to, botstate.Fancy("Add Status first."))
				}
			}
		}
	}
	if op.Message.ContentType == 1 {
		if botstate.StartChangeImg && len(botstate.Changepic) != 0 {
			if !handler.MemUser(to, sender) {
				if botstate.ChangPict {
					path, botstate.Err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(botstate.Changepic))
						for n, p := range botstate.Changepic {
							if handler.TimeDown(n) {
								go func(p *linetcr.Account) {
									if botstate.StartChangevImg2 {
										if botstate.Err != nil {
											fmt.Println(botstate.Err)
											p.SendMessage(to, botstate.Fancy("Update dual profile failure."))
										} else {
											p.SendMessage(to, botstate.Fancy("Update video picture done."))
										}
									} else {
										if botstate.Err != nil {
											fmt.Println(botstate.Err)
											p.SendMessage(to, botstate.Fancy("Update picture profile failure."))
										} else {
											p.SendMessage(to, botstate.Fancy("Update Image picture done."))
										}
									}
									wg.Done()
								}(p)
							}
						}
						wg.Wait()
						os.utils.RemoveString(path)
					} else {
						fmt.Println(botstate.Err)
						if botstate.StartChangevImg2 {
							client.SendMessage(to, botstate.Fancy("Download video picture Failure."))
						} else {
							client.SendMessage(to, botstate.Fancy("Download Image picture Failure."))
						}
					}
				} else if botstate.ChangCover {
					path, botstate.Err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(botstate.Changepic))
						for n, p := range botstate.Changepic {
							if handler.TimeDown(n) {
								go func(p *linetcr.Account) {
									if botstate.StartChangevImg2 {
										if botstate.Err != nil {
											fmt.Println(botstate.Err)
											p.SendMessage(to, botstate.Fancy("Update video cover failure."))
										} else {
											p.SendMessage(to, botstate.Fancy("Update video cover done."))
											time.Sleep(2 * time.Second)
										}
									} else {
										if botstate.Err != nil {
											fmt.Println(botstate.Err)
											p.SendMessage(to, botstate.Fancy("Update picture cover failure."))
										} else {
											p.SendMessage(to, botstate.Fancy("Update Image cover done."))
											time.Sleep(2 * time.Second)
										}
									}
									wg.Done()
								}(p)
							}
						}
						wg.Wait()
						os.utils.RemoveString(path)
					} else {
						fmt.Println(botstate.Err)
						if botstate.StartChangevImg2 {
							client.SendMessage(to, botstate.Fancy("Download video cover Failure."))
						} else {
							client.SendMessage(to, botstate.Fancy("Download Image cover Failure."))
						}
					}
				}
			}
		}
	} else if op.Message.ContentType == 16 {
		if msg.ContentMetadata["serviceType"] == "GB" && handler.MemUser(msg.To, msg.From_) {
			if room.ProNote {
				exe, _ := handler.SelectBot(client, to)
				if exe != nil {
					handler.Setkickto(exe, to, []string{msg.From_})
					botstate.Banned.AddBan(msg.From_)
				}
			}
		}
		if msg.ContentMetadata["serviceType"] == "MH" {
			if handler.MemUser(msg.To, msg.From_) {
				if room.ProPost {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, []string{msg.From_})
						botstate.Banned.AddBan(msg.From_)
					}
				}
			}
		}
		if msg.ContentMetadata["locKey"] == "BA" || msg.ContentMetadata["locKey"] == "BT" {
			if handler.MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, []string{msg.From_})
						botstate.Banned.AddBan(msg.From_)
					}
				}
			}
		}
	} else if op.Message.ContentType == 18 {
		if msg.ContentMetadata["LOC_KEY"] == "BD" {
			if handler.MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, []string{msg.From_})
						botstate.Banned.AddBan(msg.From_)
					}
				}
			}
		} else if msg.ContentMetadata["LOC_KEY"] == "BB" {
			if handler.MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, []string{msg.From_})
						botstate.Banned.AddBan(msg.From_)
					}
				}
			}
		} else if msg.ContentMetadata["LOC_KEY"] == "BO" {
			if handler.MemUser(msg.To, msg.From_) {
				if room.ProAlbum {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, []string{msg.From_})
						botstate.Banned.AddBan(msg.From_)
					}
				}
			}
		}
	} else if op.Message.ContentType == 18 { //NEW PROCJECT
    	runtime.GOMAXPROCS(botstate.Cpu)
	    if msg.ContentMetadata["serviceType"] == "AB"  {
	   	   if room.ProAlbum && handler.MemUser(to, msg.From_){
	   	   	  botstate.Banned.AddBan(msg.From_)
	   	  	  client.DeleteOtherFromChats(to, []string{msg.From_})
	   	  }
	   }
	} else if op.Message.ContentType == 6 {
		if msg.ContentMetadata["GC_EVT_TYPE"] == "S" {
			if handler.MemUser(msg.To, msg.From_) {
				if room.ProCall {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, []string{msg.From_})
						botstate.Banned.AddBan(msg.From_)
					}
				}
			}
		}
              if msg.ToType == 0 {
                    if op.Type != 50 && room.ProSpam {
		             if msg.ContentMetadata["GC_EVT_TYPE"] == "I" {
			            if handler.MemUser(msg.To, msg.From_) {
					     exe, _ := handler.SelectBot(client, to)
					     if exe != nil {
						     handler.Setkickto(exe, to, []string{msg.From_})
						     botstate.Banned.AddBan(msg.From_)
					      }
					}
				}
			}
		}
	} else if op.Message.ContentType == 14 {
		if handler.MemUser(msg.To, msg.From_) {
			if room.ProFile {
				exe, _ := handler.SelectBot(client, to)
				if exe != nil {
					handler.Setkickto(exe, to, []string{msg.From_})
					botstate.Banned.AddBan(msg.From_)
				}
			}
		}
	} else if op.Message.ContentType == 7 {
		if handler.MemUser(msg.To, msg.From_) {
			if room.ProSticker {
				exe, _ := handler.SelectBot(client, to)
				if exe != nil {
					handler.Setkickto(exe, to, []string{msg.From_})
					botstate.Banned.AddBan(msg.From_)
				}
			}
		}//COMMAND_STICKER
		if handler.CheckPermission(5, sender, to) {
		if !handler.MemUser(msg.To, msg.From_) {
			if botstate.GetStickerKick == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("kick by sticker updated"))
				}
			} else if botstate.GetStickerRespon == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("respon by sticker updated"))
				}
			} else if botstate.GetStickerStayall == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("stayall by sticker updated"))
				}
			} else if botstate.GetStickerLeave == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("leave by sticker updated"))
				}
			} else if botstate.GetStickerKickall == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("kickall by sticker updated"))
				}
			} else if botstate.GetStickerBypass == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("bypass by sticker updated"))
				}
			} else if botstate.GetStickerInvite == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("invite by sticker updated"))
				}
			} else if botstate.GetStickerClearban == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("clearban by sticker updated"))
				}
			} else if botstate.GetStickerCancelall == 1 {
				if !handler.MemUser(msg.To, msg.From_) {
					handler.SaveBackup()
					client.SendMessage(to, botstate.Fancy("cancelall by sticker updated"))
				}
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if handler.MemUser(to, lists[i]) && !utils.InArrayString(listuser, lists[i]) {
							if botstate.AutoBan {
										botstate.Banned.AddBan(lists[i])
							}
							listuser = append(listuser, lists[i])
						}
					}
				}
				fmt.Println(listuser)
				if len(listuser) != 0 {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						handler.Setkickto(exe, to, listuser)
						handler.LogAccess(client, to, sender, "kick", listuser, msg.ToType)
					}
				}
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid2 && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid2 {
				for _, p := range bk {
					go p.SendMessage(to, botstate.Fancy(botstate.MsgRespon))
				}
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid3 && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid3 {
				if room.Limit {
					client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
					return
				}
				numb := len(botstate.ClientBot)
				if numb > 0 && numb <= len(botstate.ClientBot) {
					handler.GetSquad(client, to)
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
						handler.GetSquad(client, to)
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
							wi := handler.GetSquad(client, to)
							for i := 0; i < len(all); i++ {
								if i == g {
									break
								}
								l := all[i]
								if l != client && !linetcr.InArrayCl(wi, l) {
									if !l.Limited {
										wg.Add(1)
										go func() {
											l.AcceptTicket(to, ti)
											wg.Done()
										}()
									}
								}
							}
							wg.Wait()
							client.UpdateChatQrV2(to, true)
							handler.GetSquad(client, to)
							handler.LogAccess(client, to, sender, "bringbot", []string{}, 2)
							handler.SaveBackup()
							aa := len(room.Client)
							var name string
							name = fmt.Sprintf("Ready %v bots here", aa)
							client.SendMessage(to, botstate.Fancy(name))
							//newsend += name + "\n"
						}
					}
				}
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid4 && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid4 {
				_, mem := client.GetGroupInvitation(to)
				anu := []string{}
				for m := range mem {
					if utils.InArrayString(botstate.Squadlist, m) {
						anu = append(anu, m)
					}
				}
				if len(anu) != 0 {
					for _, mid := range anu {
						cl := handler.GetKorban(mid)
						cl.AcceptGroupInvitationNormal(to)
					}
				}
				handler.GetSquad(client, to)
				room := linetcr.GetRoom(to)
				bk = room.Client
				for _, cl := range bk {
					go cl.LeaveGroup(to)
				}
				if botstate.LogGroup == to {
				}
				linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
				handler.SaveBackup()
				handler.LogAccess(client, to, sender, "leave", []string{}, msg.ToType)
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid5 && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid5 {
				_, memlist, _ := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist {
					if utils.InArrayString(botstate.Squadlist, mid) {
						cl := handler.GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				max := len(exe) * 100
				lkick := []string{}
				for n, v := range memlist {
					if handler.MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0
				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					go func(to string, i int) {
						target := lkick[i]
						cl := nom[i]
						cl.DeleteOtherFromChats(to, []string{target})
					}(to, i)
				}
				handler.LogAccess(client, to, sender, "kickall", lkick, msg.ToType)
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid7 && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid7 {
				listuser := []string{}
				nCount := 0
				typec := "rplay"
				if nCount == 0 {
					nCount = 1
				}
				lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
				if len(lists) != 0 {
					for i := range lists {
						if !utils.InArrayString(listuser, lists[i]) {
							listuser = append(listuser, lists[i])
						}
					}
				}
				if len(listuser) != 0 {
					exe, _ := handler.SelectBot(client, to)
					if exe != nil {
						lists := handler.Setinvitetomsg(exe, to, listuser)
						if len(lists) != 0 {
									handler.Cekbanwhois(client, to, lists)
						}
						handler.LogAccess(client, to, sender, "invite", listuser, msg.ToType)
					} else {
						client.SendMessage(to, botstate.Fancy("Please add another bot that has a ban invite."))
					}
				}
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid8 && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid8 {
				if len(botstate.Banned.Banlist) != 0 {
					msgcbn := fmt.Sprintf(botstate.MsgBan, len(botstate.Banned.Banlist))
					handler.LogAccess(client, to, sender, "clearban", botstate.Banned.Banlist, msg.ToType)
					client.SendMessage(to, botstate.Fancy(msgcbn))
					botstate.Banned.Banlist = []string{}
					botstate.Banned.Exlist = []string{}
				} else {
					client.SendMessage(to, botstate.Fancy("Ban list is empty."))
				}
			} else if op.Message.ContentMetadata["STKID"] == botstate.Stkid9 && op.Message.ContentMetadata["STKPKGID"] == botstate.Stkpkgid9 {
				_, memlist2, memlist := client.GetChatList(to)
				exe := []*linetcr.Account{}
				oke := []string{}
				for _, mid := range memlist2 {
					if utils.InArrayString(botstate.Squadlist, mid) {
						cl := handler.GetKorban(mid)
						if cl.Limited == false {
							exe = append(exe, cl)
						}
						oke = append(oke, mid)
					}
				}
				lkick := []string{}
				max := len(exe) * 10
				for n, v := range memlist {
					if handler.MemUser(to, v) {
						lkick = append(lkick, v)
					}
					if n > max {
						break
					}
				}
				nom := []*linetcr.Account{}
				ilen := len(lkick)
				xx := 0

				for i := 0; i < ilen; i++ {
					if xx < len(exe) {
						nom = append(nom, exe[xx])
						xx += 1
					} else {
						xx = 0
						nom = append(nom, exe[xx])
					}
				}
				for i := 0; i < ilen; i++ {
					target := lkick[i]
					cl := nom[i]
					ants.Submit(func() { cl.CancelChatInvitations(to, []string{target}) })
				}
				handler.LogAccess(client, to, sender, "cancelall", lkick, msg.ToType)
			}
		}}
	} else if op.Message.ContentType == 13 {
		if handler.MemUser(msg.To, msg.From_) {
			if room.ProContact {
				exe, _ := handler.SelectBot(client, to)
				if exe != nil {
					handler.Setkickto(exe, to, []string{msg.From_})
					botstate.Banned.AddBan(msg.From_)
				}
			}
		}
	} else if op.Message.ContentType == 2 {
		if botstate.StartChangevImg && len(botstate.Changepic) != 0 {
			if !handler.MemUser(to, sender) {
				if botstate.ChangVpict {
					path, botstate.Err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(botstate.Changepic))
						for _, p := range botstate.Changepic {
							go func(p *linetcr.Account) {
								if botstate.Err != nil {
									fmt.Println(botstate.Err)
									p.SendMessage(to, botstate.Fancy("Update video profile failure."))
								}
								wg.Done()
							}(p)
						}
						wg.Wait()
						client.SendMessage(to, botstate.Fancy("Upload video done, now Please Send Your Image !!!"))
						os.utils.RemoveString(path)
					} else {
						fmt.Println(botstate.Err)
						client.SendMessage(to, botstate.Fancy("Download Image Failure."))
					}
				} else if botstate.ChangVcover {
					path, botstate.Err := client.DownloadObjectMsg(msg.ID)
					if path != "" {
						var wg sync.WaitGroup
						wg.Add(len(botstate.Changepic))
						for _, p := range botstate.Changepic {
							go func(p *linetcr.Account) {
								p.UpdateCoverVideo(path)
								wg.Done()
							}(p)
						}
						wg.Wait()
						client.SendMessage(to, botstate.Fancy("Upload video done, now Please Send Your Image !!!"))
						os.utils.RemoveString(path)
					} else {
						fmt.Println(botstate.Err)
						client.SendMessage(to, botstate.Fancy("Download Image Failure."))
					}
				}
			}
		}
	} else if msg.ContentType == 0 && msg.Text != "" {
		if botstate.FixedToken && botstate.From_Token == msg.From_ && botstate.Group_Token == to {
			anu := handler.StripOut(pesan)
			if anu == "done reboot" {
				botstate.Data.RestartBack = fmt.Sprintf("%v", to)
				handler.SaveBackup()
				client.SendMessage(to, fmt.Sprintf("Done Editing, You Can Use After %v Seconds.", len(botstate.ClientBot)))
				handler.ReloginProgram()
			}
		}
		if room.Automute && handler.MemUser(to, msg.From_) {
			if client.Limited == false {
				client.DeleteOtherFromChats(to, []string{msg.From_})
			} else {
				for _, bot := range bk {
					if bot.Limited == false {
						bot.DeleteOtherFromChats(to, []string{msg.From_})
						break
					}
				}
			}
		} else {
			if handler.MemBan2(to, msg.From_) && handler.MemUser(to, msg.From_) {
				if client.Limited == false {
					client.DeleteOtherFromChats(to, []string{msg.From_})
				} else {
					for _, bot := range bk {
						if bot.Limited == false {
							bot.DeleteOtherFromChats(to, []string{msg.From_})
							break
						}
					}
				}
			}
		}
		cmds := handler.GetTxt(sender, client, pesan, Rname, Sname, client.MID, mentionlist, to)
		text := op.Message.Text
		newsend := ""
		var pesan = strings.ToLower(text)
		for _, cmd := range strings.Split(cmds, ",") {
			if strings.HasPrefix(cmd, "creator") && cmd != "creator" {
				if handler.CheckPermission(0, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 15
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "maker") && cmd != "makers" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 13
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "creators" {
				rngcmd := handler.GetComs(0, "creators")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Creator) != 0 {
							list := "Creator List:\n"
							for num, xd := range botstate.UserBot.Creator {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Creator list is empty.\n"
						}
					}
				}
			} else if cmd == "makers" {
				rngcmd := handler.GetComs(1, "makers")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Maker) != 0 {
							list := "Maker List:\n"
							for num, xd := range botstate.UserBot.Maker {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Maker list is empty.\n"
						}
					}
				}
			} else if cmd == "clearcreator" {
				if handler.CheckPermission(0, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Creator) != 0 {
							newsend += fmt.Sprintf("Cleared %v Makerlist\n", len(botstate.UserBot.Creator))
							botstate.UserBot.ClearCreator()
						} else {
							newsend += "Creator list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmaker" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Maker) != 0 {
							newsend += fmt.Sprintf("Cleared %v Makerlist\n", len(botstate.UserBot.Maker))
							botstate.UserBot.ClearMaker()
						} else {
							newsend += "Maker list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "uncreator") {
				if handler.CheckPermission(0, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 10
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "uncreator"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Creator)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmaker") {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 9
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unmaker"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Maker)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "seller") && cmd != "seller" {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 17
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "sellers" {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Seller) != 0 {
							list := " ✠ 𝗦𝗲𝗹𝗹𝗲𝗿𝘀 ✠ \n"
							for num, xd := range botstate.UserBot.Seller {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Seller list is empty.\n"
						}
					}
				}
			} else if cmd == "clearseller" {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Seller) != 0 {
							newsend += fmt.Sprintf("Cleared %v sellerlist\n", len(botstate.UserBot.Seller))
							botstate.UserBot.ClearSeller()
						} else {
							newsend += "Seller list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unseller") {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 17
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unseller"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Seller)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "listcmd" {
				rngcmd := handler.GetComs(5, "listcmd")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						list := handler.CmdListCheck()
						client.SendMessage(to, botstate.Fancy(list))
					}
				}
			} else if strings.HasPrefix(cmd, "expel") {
				if handler.CheckPermission(8, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 8
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "access" || cmd == botstate.Commands.Access && botstate.Commands.Access != "" {
				rngcmd := handler.GetComs(7, "ess")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						allmanagers := []string{}
						listadm := "𝗔𝗰𝗰𝗲𝘀𝘀 𝗹𝗶𝘀𝘁:"
						if len(botstate.DEVELOPER) != 0 {
							listadm += "\n\n >Developer:"
							for num, xd := range botstate.DEVELOPER {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(botstate.UserBot.Creator) != 0 {
							listadm += "\n\n >Creator:"
							for num, xd := range botstate.UserBot.Creator {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(botstate.UserBot.Maker) != 0 {
							listadm += "\n\n >Maker:"
							for num, xd := range botstate.UserBot.Maker {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(botstate.UserBot.Seller) != 0 {
							listadm += "\n\n >Seller:"
							for num, xd := range botstate.UserBot.Seller {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(botstate.UserBot.Buyer) != 0 {
							listadm += "\n\n >Buyer:"
							for num, xd := range botstate.UserBot.Buyer {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(botstate.UserBot.Owner) != 0 {
							listadm += "\n\n >Owner:"
							for num, xd := range botstate.UserBot.Owner {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(botstate.UserBot.Master) != 0 {
							listadm += "\n\n >Master:"
							for num, xd := range botstate.UserBot.Master {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(botstate.UserBot.Admin) != 0 {
							listadm += "\n\n >Admin:"
							for num, xd := range botstate.UserBot.Admin {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(allmanagers) != 0 {
							newsend += listadm + "\n"
						} else {
							newsend += "𝗔ccess is empty.\n"
						}
					}
				}
			} else if cmd == "allbanlist" || cmd == botstate.Commands.Allbanlist && botstate.Commands.Allbanlist != "" {
				rngcmd := handler.GetComs(5, "allbanlist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listadm := handler.AllBanList(client)
						if listadm != "All Banlist:" {
							newsend += listadm + "\n"
						} else {
							newsend += "Banlist is empty.\n"
						}
					}
				}
			} else if cmd == "gaccess" || cmd == botstate.Commands.Gaccess && botstate.Commands.Gaccess != "" {
				rngcmd := handler.GetComs(9, "access")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						allmanagers := []string{}
						listadm := "𝗚𝗮𝗰𝗰𝗲𝘀𝘀 𝗹𝗶𝘀𝘁:"
						if len(room.Gowner) != 0 {
							listadm += "\n\n >Gowner:"
							for num, xd := range room.Gowner {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(room.Gadmin) != 0 {
							listadm += "\n\n >Gadmin:"
							for num, xd := range room.Gadmin {
								num++
								rengs := strconv.Itoa(num)
								allmanagers = append(allmanagers, xd)
								name := handler.GetContactName(client, xd)
									listadm += "\n " + rengs + ". " + name
							}
						}
						if len(allmanagers) != 0 {
							newsend += listadm + "\n"
						} else {
							newsend += "Gaccess is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "buyer") && cmd != "buyers" {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setdate ") {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						ha := strings.Split((cmd), "setdate ")
						haj := ha[1]
						haj = handler.StripOut(haj)
						has := strings.Split(haj, "-")
						if len(has) == 3 {
							yy, _ := strconv.Atoi(has[0])
							mm, _ := strconv.Atoi(has[1])
							dd, _ := strconv.Atoi(has[2])
							var time2 = time.Date(yy, time.Month(mm), dd, 00, 00, 0, 0, time.UTC)
							times := time2.Format(time.RFC3339)
							botstate.Data.Dalltime = times
							str := fmt.Sprintf("⚙️ Date:\n %v-%v-%v", yy, mm, dd)
							ta := time2.Sub(time.Now())
							str += fmt.Sprintf("\n⚙️ Remaining:\n  %v", handler.BotDuration(ta))
							newsend += str + "\n"
						}
					}
				}
			} else if cmd == "addweek" {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", botstate.Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 7 * mont
						t := batas.Add(mont)
						botstate.Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", handler.BotDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addday" {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", botstate.Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						t := batas.Add(mont)
						botstate.Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", handler.BotDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "addmonth" {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", botstate.Data.Dalltime)
						has := strings.Split(d, "-")
						has2 := strings.Split(has[2], "T")
						yy, _ := strconv.Atoi(has[0])
						mm, _ := strconv.Atoi(has[1])
						timeup, _ := strconv.Atoi(has2[0])
						batas := time.Date(yy, time.Month(mm), timeup, 00, 00, 0, 0, time.UTC)
						mont := 24 * time.Hour
						mont = 30 * mont
						t := batas.Add(mont)
						botstate.Data.Dalltime = t.Format(time.RFC3339)
						ta := t.Sub(time.Now())
						str := fmt.Sprintf("⚙️ Remaining:\n\n  %v", handler.BotDuration(ta))
						newsend += str + "\n"
					}
				}
			} else if cmd == "reboot" {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.SaveBackup()
						client.SendMessage(to, botstate.Fancy("Waiting....."))
						handler.ReloginProgram()
						client.SendMessage(to, botstate.Fancy("Done"))
					}
				}
			} else if cmd == "runall" {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.SaveBackup()
						client.SendMessage(to, botstate.Fancy("Waiting....."))
						handler.ReloginProgram()
						client.SendMessage(to, botstate.Fancy("Done"))
					}
				}
			} else if strings.HasPrefix(cmd, "unbuyer") {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unbuyer"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Buyer)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "checkram" || cmd == botstate.Commands.Checkram && botstate.Commands.Checkram != "" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						v, _ := mem.VirtualMemory()
						r := fmt.Sprintf("  ≻ Cpu : %v core\n  ≻ Ram : %v mb\n  ≻ Free : %v mb\n  ≻ Cache : %v mb\n  ≻ UsedPercent : %f %%", botstate.Cpu, handler.BToMb(v.Used+v.Free+v.Buffers+v.Cached), handler.BToMb(v.Free), handler.BToMb(v.Buffers+v.Cached), v.UsedPercent)
						newsend += r + "\n"
					}
				}
			} else if cmd == "memory" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						v, _ := mem.VirtualMemory()
						r := fmt.Sprintf("  ≻ Cpu : %v core\n  ≻ Ram : %v mb\n  ≻ Free : %v mb\n  ≻ Cache : %v mb\n  ≻ UsedPercent : %f %%", botstate.Cpu, handler.BToMb(v.Used+v.Free+v.Buffers+v.Cached), handler.BToMb(v.Free), handler.BToMb(v.Buffers+v.Cached), v.UsedPercent)
						newsend += r + "\n"
					}
				}
			} else if cmd == "clearbuyer" {
				if handler.CheckPermission(3, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Buyer) != 0 {
							newsend += fmt.Sprintf("Cleared %v buyerlist\n", len(botstate.UserBot.Buyer))
							botstate.UserBot.ClearBuyer()
						} else {
							newsend += "Buyer list is empty.\n"
						}
					}
				}
			} else if cmd == "upimage" || cmd == botstate.Commands.Upimage && botstate.Commands.Upimage != "" {
				rngcmd := handler.GetComs(4, "upimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Which bot's you want to update Pict.\n"
					}
				}
			} else if cmd == "upcover" || cmd == botstate.Commands.Upcover && botstate.Commands.Upcover != "" {
				rngcmd := handler.GetComs(4, "upcover")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Which bot's you want to update Cover ?.\n"
					}
				}
			} else if cmd == "upvimage" || cmd == botstate.Commands.Upvimage && botstate.Commands.Upvimage != "" {
				rngcmd := handler.GetComs(4, "upvimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Which bot's you want to update Pict ?.\n"
					}
				}
			} else if cmd == "upvcover" || cmd == botstate.Commands.Upvcover && botstate.Commands.Upvcover != "" {
				rngcmd := handler.GetComs(4, "upvcover")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Which bot's you want to update Cover ?.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "unsend ") {
				rngcmd := handler.GetComs(5, "unsend")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, "𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿")
								return
							} else {
								if result2 > 0 {
									Nganu, _ := client.GetRecentMessagesV2(op.Message.To)
									Mid := []string{}
									unsed := []string{}
									for _, chat := range Nganu {
										if utils.InArrayString(botstate.Squadlist, chat.From_) {
											Mid = append(Mid, chat.ID)
										}
									}
									for i := 0; i < len(Mid); i++ {
										if i < result2 {
											unsed = append(unsed, Mid[i])
										}
									}
									if len(unsed) != 0 {
										exess, _ := handler.SelectallBot(client, to)
										if exess != nil {
											for i := range exess {
												Nganu2, _ := exess[i].GetRecentMessagesV2(op.Message.To)
												for _, chat := range Nganu2 {
													if chat.From_ == exess[i].MID {
														if utils.InArrayString(unsed, chat.ID) {
															exess[i].UnsendChatnume(to, chat.ID)
														}
													}
												}
											}
										}
									}
								} else {
									client.SendMessage(to, botstate.Fancy("out of range."))
								}
							}
						} else {
							client.SendMessage(to, botstate.Fancy("Msg not fund number"))
						}
					}
				}
			} else if cmd == "purgeall" || cmd == botstate.Commands.Purgeall && botstate.Commands.Purgeall != "" {
				rngcmd := handler.GetComs(5, "purgeall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						gr, _ := client.GetGroupIdsJoined()
						all := []string{}
						for _, aa := range gr {
							_, memlist, _ := client.GetChatList(aa)
							lkicks := []string{}
							for _, v := range memlist {
								if handler.MemUser(aa, v) {
									lkicks = append(lkicks, v)
								}
							}
	                                        c, _ := client.GetChats([]string{aa})
	                                        zxc := c.Chats[0].Extra.GroupExtra.InviteeMids
	                                        for k, _ := range zxc {
		                                        if handler.IsBlacklist(client, k) == true {
			                                        go func(k string) {
				                                        go client.CancelChatInvitations(aa, []string{k})
			                                        }(k)
		                                        }
	                                        }
							lkick := []string{}
							for _, ban := range lkicks {
								if handler.MemBan(aa, ban) {
									lkick = append(lkick, ban)
									all = append(all, ban)
								}
							}
							nom := []*linetcr.Account{}
							ilen := len(lkick)
							xx := 0
							exe := []*linetcr.Account{}
							for _, c := range linetcr.GetRoom(aa).Client {
								if !c.Limited {
									exe = append(exe, c)
								}
							}
							if len(exe) != 0 {
								for i := 0; i < ilen; i++ {
									if xx < len(exe) {
										nom = append(nom, exe[xx])
										xx += 1
									} else {
										xx = 0
										nom = append(nom, exe[xx])
									}
								}
								for i := 0; i < ilen; i++ {
									target := lkick[i]
									cl := nom[i]
									go cl.DeleteOtherFromChats(aa, []string{target})
								}
								time.Sleep(1 * time.Second)
                                               }
						}
						newsend += fmt.Sprintf("Success purgeall blacklist")
						handler.LogAccess(client, to, sender, "purgeall", all, msg.ToType)
					}
				}
			} else if cmd == "2purgeall"  {
				rngcmd := handler.GetComs(5, "2purgeall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						gr, _ := client.GetGroupIdsJoined()
						for _, aa := range gr {
							go handler.KickAllBan(client, aa)
							time.Sleep(1 * time.Second)
						}
						newsend += fmt.Sprintf("Success purgeall blacklist")
					}
				}
			} else if strings.HasPrefix(cmd, "gleave") {
				rngcmd := handler.GetComs(5, "gleave")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									su := "gleave"
									str := ""
									if strings.HasPrefix(text, Rname+" ") {
										str = strings.Replace(text, Rname+" "+su+" ", "", 1)
										str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname+" ") {
										str = strings.Replace(text, Sname+" "+su+" ", "", 1)
										str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Rname) {
										str = strings.Replace(text, Rname+su+" ", "", 1)
										str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname) {
										str = strings.Replace(text, Sname+su+" ", "", 1)
										str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
									}
									st := handler.StripOut(str)
									hapuss := linetcr.Archimed(st, botstate.Tempgroup)
									if len(hapuss) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									names := []string{}
									for _, gid := range hapuss {
										name, mem := client.GetGroupInvitation(gid)
										names = append(names, name)
										anu := []string{}
										for m := range mem {
											if utils.InArrayString(botstate.Squadlist, m) {
												anu = append(anu, m)
											}
										}
										if len(anu) != 0 {
											for _, mid := range anu {
												cl := handler.GetKorban(mid)
												cl.AcceptGroupInvitationNormal(gid)
												linetcr.GetRoom(gid).ConvertGo(cl)
											}
										}
										handler.GetSquad(client, gid)
										room := linetcr.GetRoom(gid)
										bk = room.Client
										for _, cl := range bk {
											go cl.LeaveGroup(gid)
										}
										if botstate.LogGroup == gid {
										}
										linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
									}
									strs := strings.Join(names, ", ")
									client.SendMessage(to, botstate.Fancy("Bot's leave from group: \n\n"+strs))
									handler.SaveBackup()
								}
							}
						} else {
							newsend += "Group not found"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "invme ") {
				rngcmd := handler.GetComs(5, "invme")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										handler.GetSquad(client, gid)
										room := linetcr.GetRoom(gid)
										bk := room.Client
										name, mem, inv := client.GetChatList(gid)
										if utils.InArrayString(mem, msg.From_) {
											client.SendMessage(to, botstate.Fancy("You was on group "+name))
											return
										} else {
											if utils.InArrayString(inv, msg.From_) {
												bk[0].CancelChatInvitations(gid, []string{msg.From_})
											}
											for _, cl := range bk {
												if !cl.Limited && !cl.Limitadd {
													handler.AddContact2(cl, msg.From_)
													fl, _ := cl.GetAllContactIds()
													if utils.InArrayString(fl, msg.From_) {
														if botstate.Err != nil {
															code := linetcr.GetCode(botstate.Err)
															if code != 35 && code != 10 {
																client.SendMessage(to, botstate.Fancy("You has invited to group "+name))
																return
															}
														} else {
															client.SendMessage(to, botstate.Fancy("You has invited to group "+name))
															return
														}
													}
												}
											}
											newsend += "Sorry, all bot has invite banned"
										}
									} else {
										newsend += "out of range."
									}
								}
							}
						} else {
							newsend += "Group not found"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "decline ") {
				rngcmd := handler.GetComs(2, "decline")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									num, _ := strconv.Atoi(result[1])
									gr := []string{}
									for i := range botstate.ClientBot {
										grs, _ := botstate.ClientBot[i].GetGroupsInvited()
										if len(grs) != 0 {
											for _, a := range grs {
												if !utils.InArrayString(gr, a) {
													gr = append(gr, a)
												}
											}
										}
									}
									grup, _ := client.GetGroups(gr)
									for _, gi := range grup {
										if !utils.InArrayString(botstate.Tempgroup, gi.ChatMid) {
										}
									}
									if num > 0 && num <= len(botstate.Tempgroup) {
										exe := []*linetcr.Account{}
										gen := botstate.Tempgroup[num-1]
										names, _, _ := client.GetChatList(botstate.Tempgroup[num-1])
										for i := range botstate.ClientBot {
											if botstate.ClientMid[botstate.ClientBot[i].MID].Limited == false {
												grs, _ := botstate.ClientBot[i].GetGroupsInvited()
												if utils.InArrayString(grs, gen) {
													exe = append(exe, botstate.ClientBot[i])
												}
											}
										}
										if len(exe) != 0 {
											for i := range exe {
												exe[i].RejectChatInvitation(gen)
											}
											newsend += fmt.Sprintf("Successfully declined invitation for: %v\n", names)
										}
									} else {
										newsend += "out of range pendinglist.\n"
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "accept") && cmd != "acceptall" {
				rngcmd := handler.GetComs(5, "accept")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if result[1] != "0" {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									su := "accept"
									str := ""
									if strings.HasPrefix(text, Rname+" ") {
										str = strings.Replace(text, Rname+" "+su+" ", "", 1)
										str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname+" ") {
										str = strings.Replace(text, Sname+" "+su+" ", "", 1)
										str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Rname) {
										str = strings.Replace(text, Rname+su+" ", "", 1)
										str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
									} else if strings.HasPrefix(text, Sname) {
										str = strings.Replace(text, Sname+su+" ", "", 1)
										str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
									}
									st := handler.StripOut(str)
									hapuss := linetcr.Archimed(st, botstate.Tempginv)
									if len(hapuss) == 0 {
										newsend += "Please input the right number\nSee group number with command groups"
									} else {
										names := []string{}
										for _, gid := range hapuss {
											name, mem := client.GetGroupInvitation(gid)
											names = append(names, name)
											anu := []string{}
											for m := range mem {
												if utils.InArrayString(botstate.Squadlist, m) {
													anu = append(anu, m)
												}
											}
											if len(anu) != 0 {
												for _, mid := range anu {
													cl := handler.GetKorban(mid)
													cl.AcceptGroupInvitationNormal(gid)
													linetcr.GetRoom(gid).ConvertGo(cl)
												}
											}
										}
										str := strings.Join(names, ", ")
										newsend += "Bot's join to group \n\n" + str
									}
								}
							}
						}
					}
				}
			} else if cmd == "abort" {
				rngcmd := handler.GetComs(5, "abort")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if botstate.Remotegrupidto != "" {
							client.SendMessage(botstate.Remotegrupidto, botstate.Fancy("Done Have abort."))
						} else {
							newsend += "Done Have abort." + "\n"
						}
						handler.Abort()
					}
				}
			} else if cmd == "declineall" {
				rngcmd := handler.GetComs(2, "declineall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for i := range botstate.ClientBot {
							grs, _ := botstate.ClientBot[i].GetGroupsInvited()
							if len(grs) != 0 {
								grup, _ := client.GetGroups(grs)
								for _, gi := range grup {
									if !utils.InArrayString(botstate.Tempgroup, gi.ChatMid) {
									}
									botstate.ClientBot[i].RejectChatInvitation(gi.ChatMid)
								}
								time.Sleep(1 * time.Second)
							}

						}
						if len(botstate.Tempgroup) != 0 {
							newsend += fmt.Sprintf("Successfully declined invitations: (%v)\n", len(botstate.Tempgroup))
						} else {
							newsend += "pending list is empty.\n"
						}
					}
				}
			} else if cmd == "acceptall" {
				rngcmd := handler.GetComs(4, "acceptall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for i := range botstate.ClientBot {
							grs, _ := botstate.ClientBot[i].GetGroupsInvited()
							if len(grs) != 0 {
								grup, _ := client.GetGroups(grs)
								for _, gi := range grup {
									if !utils.InArrayString(botstate.Tempgroup, gi.ChatMid) {
									}
									botstate.ClientBot[i].AcceptGroupInvitationNormal(gi.ChatMid)
									linetcr.GetRoom(gi.ChatMid).ConvertGo(botstate.ClientBot[i])
									time.Sleep(1 * time.Second)
								}
							}
						}
						if len(botstate.Tempgroup) != 0 {
							newsend += fmt.Sprintf("Success accept bot %v Group\n", len(botstate.Tempgroup))
						} else {
							newsend += "pending list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "upstatus") {
				rngcmd := handler.GetComs(4, "upstatus")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "upstatus"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						newsend += fmt.Sprintf("Which bot's should be Status %v", str)
					}
				}
			} else if strings.HasPrefix(cmd, "upname") {
				rngcmd := handler.GetComs(4, "upname")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "upname"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						aa := utf8.RuneCountInString(str)
						if aa != 0 && aa <= 20 {
							newsend += fmt.Sprintf("Which bot's should be Name %v", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "tiktokdl") {
				rngcmd := handler.GetComs(4, "tiktokdl")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "tiktokdl"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := "https://api.minzteam.xyz/tiktokdl?url=" + str + "&apikey="+botstate.Apikey
						client.SendVideoWithURL(to, data)
					}
				}
			} else if strings.HasPrefix(cmd, "youtubedl") {
				rngcmd := handler.GetComs(4, "youtubedl")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "youtubedl"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						client.Sendmediadl(to, str,"youtube")
					}
				}
			} else if cmd == "buyers" {
				rngcmd := handler.GetComs(4, "buyers")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Buyer) != 0 {
							list := " ✠ 𝗯𝘂𝘆𝗲𝗿𝘀 ✠ \n"
							for num, xd := range botstate.UserBot.Buyer {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Buyer list is empty.\n"
						}
					}
				}
			} else if cmd == "history" {
				rngcmd := handler.GetComs(5, "history")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.SaveProHistory()
						list := fmt.Sprintf("History Bot:\n\n Kick: %v \n Cancel: %v \n Invited: %v", botstate.Data.Kikhistory, botstate.Data.Canclhistory, botstate.Data.Invhistory)
						client.SendMessage(to, botstate.Fancy(list))
					}
				}
			} else if cmd == "2history" {
				rngcmd := handler.GetComs(5, "2history")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						countK := 0
						countinv := 0
						countcancel := 0
						for i := range botstate.ClientBot {
							countK = countK + botstate.ClientBot[i].Ckick
							countinv = countinv + botstate.ClientBot[i].Cinvite
							countcancel = countcancel + botstate.ClientBot[i].Ccancel
						}
						list := fmt.Sprintf("History: \n\n Kick: %v \n Cancel: %v \n Invited: %v", countK, countcancel, countinv)
						client.SendMessage(to, botstate.Fancy(list))
					}
				}
			} else if cmd == "clearhide" {
				rngcmd := handler.GetComs(5, "clearhide")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.CheckHaid) != 0 {
							handler.LogAccess(client, to, sender, "clearhid", botstate.CheckHaid, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v Hidelist\n", len(botstate.CheckHaid))
						} else {
							newsend += "Hide list is empty.\n"
						}
					}
				}
			} else if cmd == "hidelist" {
				rngcmd := handler.GetComs(5, "hidelist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.CheckHaid) != 0 {
							list := " ✠ Hide List ✠ \n"
							for num, xd := range botstate.CheckHaid {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Hide list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unhide") || strings.HasPrefix(cmd, "delhide") {
				rngcmd := handler.GetComs(5, "unhide")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						list := ""
						listuser := []string{}
						nCount1 := 0
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							test1 := "User removed from hidelist:\n\n"
							test2 := "User not exist in hidelist:\n\n"
							for n, xx := range listuser {
								if utils.InArrayString(botstate.CheckHaid, xx) {
									nCount1 = nCount1 + 1
								}
								rengs := strconv.Itoa(n + 1)
								name := handler.GetContactName(client, xx)
								list += fmt.Sprintf("%s. %s\n", rengs, name)
							}
							if nCount1 != 0 {
								client.SendMessage(to, botstate.Fancy(test1+list))
							} else {
								client.SendMessage(to, botstate.Fancy(test2+list))
							}
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unhide"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.CheckHaid)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											for _, i := range hapuss {
												if utils.InArrayString(botstate.CheckHaid, i) {
													listuser = append(listuser, i)
												}
											}
											if len(listuser) != 0 {
												list += "User removed from hidelist:\n\n"
												for n, xx := range listuser {
													rengs := strconv.Itoa(n + 1)
													name := handler.GetContactName(client, xx)
													list += fmt.Sprintf("%s. %s\n", rengs, name)
												}
												client.SendMessage(to, botstate.Fancy(list))
											}
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "hide") && cmd != "hidelist" {
				rngcmd := handler.GetComs(5, "hide")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						nCount1 := 0
						list := ""
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							test1 := "User added to hidelist:\n\n"
							test2 := "User already exist in hidelist:\n\n"
							for n, xx := range listuser {
								if !utils.InArrayString(botstate.CheckHaid, xx) {
									nCount1 = nCount1 + 1
								}
								rengs := strconv.Itoa(n + 1)
								name := handler.GetContactName(client, xx)
								list += fmt.Sprintf("%s. %s\n", rengs, name)
							}
							if nCount1 != 0 {
								client.SendMessage(to, botstate.Fancy(test1+list))
							} else {
								client.SendMessage(to, botstate.Fancy(test2+list))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "owner") && cmd != "owners" {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unowner") {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unowner"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Owner)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "clearowner" {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Owner) != 0 {
							handler.LogAccess(client, to, sender, "clearowner", botstate.UserBot.Owner, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v ownerlist\n", len(botstate.UserBot.Owner))
							botstate.UserBot.ClearOwner()
						} else {
							newsend += "Owner list is empty.\n"
						}
					}
				}
			} else if cmd == "notification on" {
				rngcmd := handler.GetComs(3, "notification")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if botstate.LogGroup == to {
							newsend += "Already enabled.\n"
						} else {
							newsend += "Notification is enabled.\n"
						}
					}
				}
			} else if cmd == "notification off" {
				rngcmd := handler.GetComs(3, "notification")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if botstate.LogGroup == to {
							newsend += "Notification is disabled.\n"
						} else {
							newsend += "Already disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setsname ") {
				rngcmd := handler.GetComs(4, "setsname")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "setsname ", "", 1)
						if result == "," || result == "" {
						} else {
						}
						newsend += "Sname set to: " + Sname + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setrname ") {
				rngcmd := handler.GetComs(4, "setrname")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "setrname ", "", 1)
						if result == "," || result == "" {
						} else {
						}
						newsend += "Succes update Rname to " + Rname + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgrespon") {
				rngcmd := handler.GetComs(4, "msgrespon")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "msgrespon"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						newsend += "Message respon set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "upgname") {
				rngcmd := handler.GetComs(4, "upgname")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "upgname"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						client.UpdateChatName(to, str)
						newsend += "group name has been changed to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setlogo") {
				rngcmd := handler.GetComs(4, "setlogo")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "setlogo"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						botstate.Data.Logobot = str
						newsend += "Menu logo set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgwelcome") {
				rngcmd := handler.GetComs(7, "msgwelcome")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "msgwelcome"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.WelcomeMsg = str
						newsend += "Message Welcome set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msgleave") {
				rngcmd := handler.GetComs(7, "msgleave")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "msgleave"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.MsgLeave = str
						newsend += "Message Leave set to: " + str + "\n"
					}
				}

			} else if strings.HasPrefix(cmd, "msgclearban ") {
				rngcmd := handler.GetComs(4, "msgclearban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgclearban ", "", 1)
						newsend += "Message unban set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msglurk") {
				rngcmd := handler.GetComs(7, "msglurk")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "msglurk"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						room.MsgLurk = str
						newsend += "Message sider set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "2addtoken") {
				rngcmd := handler.GetComs(2, "2addtoken")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "2addtoken"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						botstate.Data.Authoken = append(botstate.Data.Authoken,str)
						newsend += "Add token: "+str+"\n\nCmd to runall"
						time.Sleep(2 * time.Millisecond)
						handler.SaveBackup()
					}
				}
			} else if strings.HasPrefix(cmd, "addtoken") {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "addtoken"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						tokenlist := strings.Split(str, "\n")
						for _, token := range tokenlist {
							token_mid := token[:33]
							ct, botstate.Err := client.GetContact(token_mid)
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("Token Error : ")+token)
							} else {
								if !utils.InArrayString(botstate.Data.Authoken, token) {
									botstate.Data.Authoken = append(botstate.Data.Authoken, token)
									client.SendMessage(to, ct.DisplayName + botstate.Fancy(" >> Token Login"))
								} else {
									client.SendMessage(to, ct.DisplayName + botstate.Fancy(" >> Token Alerdy Login"))
								}
							}

						}

						handler.SaveBackup()
						handler.ReloginProgram()
					}
				}
			} else if strings.HasPrefix(cmd, "list token") {
				rngcmd := handler.GetComs(2, "list token")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len (botstate.Data.Authoken) != 0 {
							list :="List Token: ↓"
							for num, xd := range botstate.Data.Authoken {
								num++
								rengs := strconv.Itoa(num)
								mids := strings.Split(xd, ":")
								var prof *talkservice.Contact
								prof, _ = client.GetContact(mids[0])
								name := prof.DisplayName
								list += "\n\n  "+rengs+". "+xd+"\n Name: "+name
							}
							list += "\n\nToken limited: ↓"
							for n, cl := range linetcr.KickBans {
								m := cl.MID
								no := n + 1
								pr, _ := client.GetContact(m)
								cl.Namebot = pr.DisplayName
								list += fmt.Sprintf("\n\n%v. %s\n Name: %v", no, cl.AuthToken, cl.Namebot)
							}
							newsend += list+"\n"
						} else {
							newsend += "Notings"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "untoken") {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "untoken"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						st := handler.StripOut(str)
						hapuss := linetcr.Archimed(st, botstate.Data.Authoken)

						if len(hapuss) == 0 {
							newsend += "empty list\n"
						} else {
							can_reboot := false
							for _, token := range hapuss {
								if utils.InArrayString(botstate.Data.Authoken, token) {
									can_reboot = true
									botstate.Data.Authoken = utils.RemoveString(botstate.Data.Authoken, token)
									token_mid := token[:33]
									name := handler.GetContactName(client, token_mid)

									client.SendMessage(to, name + botstate.Fancy(" >> Token Delete !!!"))
								}
							}
							if can_reboot {
								handler.SaveBackup()
								handler.ReloginProgram()
							} else {
								client.SendMessage(to, botstate.Fancy("No user deleted"))
							}

						}
					}
				}
			} else if cmd == "status token" {
				rngcmd := handler.GetComs(2, "status token")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						kb := "Status Token:"
						for no, tok := range botstate.Data.Authoken {
							mids := strings.Split(tok, ":")
							mid := mids[0]
							if utils.InArrayString(botstate.Squadlist, mid) {
								xxx := handler.GetKorban(mid)
								if linetcr.InArrayCl(linetcr.KickBanChat, xxx) {
									if xxx.Namebot == "" {
										pr, _ := client.GetContact(mid)
										xxx.Namebot = pr.DisplayName
									}
									kb += fmt.Sprintf("\n\n%v. %s\nMid: %s\nStatus: botstate.Banned", no+1, xxx.Namebot, xxx.MID)
								} else {
									kb += fmt.Sprintf("\n\n%v. %s\nmid: %s\nstatus: Normal", no+1, xxx.Namebot, xxx.MID)
								}
							}
						}
						newsend += kb
					}
				}
			} else if cmd == "remtokenbans" {
				rngcmd := handler.GetComs(2, "remtokenbans")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if botstate.FixedToken && botstate.From_Token == sender {
							kb := "Deleted botstate.Banned:"
							for no, tok := range botstate.Data.Authoken {
								mids := strings.Split(tok, ":")
								mid := mids[0]
								Bots := []*linetcr.Account{}
								if utils.InArrayString(botstate.Squadlist, mid) {
									xxx := handler.GetKorban(mid)
									if linetcr.InArrayCl(linetcr.KickBanChat, xxx) || xxx.Frez {
										if xxx.Namebot == "" {
											pr, _ := client.GetContact(mid)
											xxx.Namebot = pr.DisplayName
										}
										botstate.Data.Authoken = utils.RemoveString(botstate.Data.Authoken, tok)
										Bots = append(Bots, xxx)
										if linetcr.InArrayCl(botstate.ClientBot, xxx) {
										}
										kb += fmt.Sprintf("\n\n%v. %s\nMid: %s", no+1, xxx.Namebot, xxx.MID)
									}
								}
								if len(Bots) != 0 {
									linetcr.RemoveBot(Bots)
								}
								for i := range botstate.ClientBot {
									for _, x := range botstate.ClientBot[i].Squads {
										if !utils.InArrayString(botstate.Squadlist, x) {
											botstate.ClientBot[i].Squads = utils.RemoveString(botstate.ClientBot[i].Squads, x)
										}
									}
								}
							}
							handler.SaveData()
							newsend += kb
						}
					}
				}
			} else if cmd == "removelimits" {
				rngcmd := handler.GetComs(2, "removelimits")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if botstate.FixedToken && botstate.From_Token == sender {
							kb := "Deleted Limit:"
							for no, tok := range botstate.Data.Authoken {
								mids := strings.Split(tok, ":")
								mid := mids[0]
								Bots := []*linetcr.Account{}
								if utils.InArrayString(botstate.Squadlist, mid) {
									xxx := handler.GetKorban(mid)
									if linetcr.InArrayCl(linetcr.KickBans, xxx) || xxx.Frez {
										if xxx.Namebot == "" {
											pr, _ := client.GetContact(mid)
											xxx.Namebot = pr.DisplayName
										}
										botstate.Data.Authoken = utils.RemoveString(botstate.Data.Authoken, tok)
										Bots = append(Bots, xxx)
										if linetcr.InArrayCl(botstate.ClientBot, xxx) {
										}
										kb += fmt.Sprintf("\n\n%v. %s\nMid: %s", no+1, xxx.Namebot, xxx.MID)
									}
								}
								if len(Bots) != 0 {
									linetcr.RemoveBot(Bots)
								}
								for i := range botstate.ClientBot {
									for _, x := range botstate.ClientBot[i].Squads {
										if !utils.InArrayString(botstate.Squadlist, x) {
											botstate.ClientBot[i].Squads = utils.RemoveString(botstate.ClientBot[i].Squads, x)
										}
									}
								}
							}
							handler.SaveData()
							newsend += kb
						}
					}
				}
			} else if strings.HasPrefix(cmd, "cosquad") {
				rngcmd := handler.GetComs(2, "cosquad")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, memlist := client.GetGroupMember(to)
						yus := []string{}
						for mek, _ := range memlist {
							if utils.InArrayString(botstate.Squadlist, mek) {
								yus = append(yus, mek)
							}
						}
						if len(yus) != 0 {
							for _, mek := range yus {
								go client.SendContact(to, mek)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "msgfresh ") {
				rngcmd := handler.GetComs(4, "msgfresh")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msgfresh ", "", 1)
						newsend += "Message fresh set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "msglimit ") {
				rngcmd := handler.GetComs(4, "msglimit")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Replace(cmd, "msglimit ", "", 1)
						newsend += "Message limit set to: " + result + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setkick ") {
				rngcmd := handler.GetComs(4, "setkick")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						anjay := strings.Split((cmd), " ")
						num, botstate.Err := strconv.Atoi(anjay[1])
						if botstate.Err != nil {
							newsend += "Please use number!\n"
						} else {
							newsend += "Limiter kick set to " + anjay[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setinvite ") {
				rngcmd := handler.GetComs(4, "setinvite")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						anjay := strings.Split((cmd), " ")
						num, botstate.Err := strconv.Atoi(anjay[1])
						if botstate.Err != nil {
							newsend += "Please use number!\n"
						} else {
							newsend += "Limiter invite set to " + anjay[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setcancel ") {
				rngcmd := handler.GetComs(4, "setcancel")
				if handler.CheckPermission(rngcmd, sender, to) {
					anjay := strings.Split((cmd), " ")
					num, botstate.Err := strconv.Atoi(anjay[1])
					if botstate.Err != nil {
						newsend += "Please use number!\n"
					} else {
						newsend += "Limiter cancel set to " + anjay[1] + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "setlimiter ") {
				rngcmd := handler.GetComs(4, "setlimiter")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						no, botstate.Err := strconv.Atoi(result[1])
						if botstate.Err != nil {
							newsend += "Please use number!\n"
						} else {
							newsend += "Limiter successs set to " + result[1] + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setpend ") {
				rngcmd := handler.GetComs(4, "setpend")
				if handler.CheckPermission(rngcmd, sender, to) {
					anjay := strings.Split((cmd), " ")
					num, botstate.Err := strconv.Atoi(anjay[1])
					if botstate.Err != nil {
						newsend += "Please use number!\n"
					} else {
						newsend += "Cancel pending set to " + anjay[1] + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "fuck") && cmd != "fucklist" {
				rngcmd := handler.GetComs(5, "fuck")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 9
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "mute") && cmd != "mutelist" {
				rngcmd := handler.GetComs(5, "mute")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 11
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "whois") {
				rngcmd := handler.GetComs(5, "whois")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 12
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "name") {
				rngcmd := handler.GetComs(5, "name")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 16
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "mid") {
				rngcmd := handler.GetComs(5, "mid")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 14
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmute") {
				rngcmd := handler.GetComs(5, "unmute")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.handler.CheckUnbanBots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unmute"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.handler.CheckUnbanBots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "owners" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Owner) != 0 {
							list := "𝗼𝘄𝗻𝗲𝗿𝘀:\n"
							for num, xd := range botstate.UserBot.Owner {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Owner list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unmaster") {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unmaster"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Master)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungowner") {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 6
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "ungowner"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, room.Gowner)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "gojoin" {
				rngcmd := handler.GetComs(8, "join")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, mem := client.GetGroupInvitation(to)
						anu := []string{}
						for m := range mem {
							if utils.InArrayString(botstate.Squadlist, m) {
								anu = append(anu, m)
							}
						}
						if len(anu) != 0 {
							for _, mid := range anu {
								cl := handler.GetKorban(mid)
								cl.AcceptGroupInvitationNormal(to)
							}
						}
						handler.GetSquad(client, to)
					}
				}
			} else if strings.HasPrefix(cmd, "master") && cmd != "masters" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gowner") && cmd != "gowners" {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 6
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "setcmd ") {
				rngcmd := handler.GetComs(5, "setcmd")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						txt := strings.ReplaceAll(cmd, "setcmd ", "")
						texts := strings.Split(txt, " ")
						if len(texts) > 1 {
							new := handler.Upsetcmd(texts[0], texts[1])
							if new != "" {
								newsend += new
							} else {
								newsend += "Cmd not found.\n"
							}
						} else {
							newsend += "Cmd not found.\n"
						}
					}
				}
			} else if cmd == "restartcmd" {
				rngcmd := handler.GetComs(5, "restartcmd")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						botstate.Commands.Botname = ""
						botstate.Commands.Upallimage = ""
						botstate.Commands.Upallcover = ""
						botstate.Commands.Unsend = ""
						botstate.Commands.Upvallimage = ""
						botstate.Commands.Upvallcover = ""
						botstate.Commands.Appname = ""
						botstate.Commands.Useragent = ""
						botstate.Commands.Hostname = ""
						botstate.Commands.Friends = ""
						botstate.Commands.Adds = ""
						botstate.Commands.Limits = ""
						botstate.Commands.Addallbots = ""
						botstate.Commands.Addallsquads = ""
						botstate.Commands.Leave = ""
						botstate.Commands.Respon = ""
						botstate.Commands.Ping = ""
						botstate.Commands.Count = ""
						botstate.Commands.Limitout = ""
						botstate.Commands.Access = ""
						botstate.Commands.Allbanlist = ""
						botstate.Commands.Allgaccess = ""
						botstate.Commands.Gaccess = ""
						botstate.Commands.Checkram = ""
						botstate.Commands.Backups = ""
						botstate.Commands.Upimage = ""
						botstate.Commands.Upcover = ""
						botstate.Commands.Upvimage = ""
						botstate.Commands.Upvcover = ""
						botstate.Commands.Bringall = ""
						botstate.Commands.Purgeall = ""
						botstate.Commands.Banlist = ""
						botstate.Commands.Clearban = ""
						botstate.Commands.Stayall = ""
						botstate.Commands.Clearchat = ""
						botstate.Commands.Here = ""
						botstate.Commands.Speed = ""
						botstate.Commands.Status = ""
						botstate.Commands.Tagall = ""
						botstate.Commands.Kick = ""
						botstate.Commands.Max = ""
						botstate.Commands.None = ""
						botstate.Commands.Kickall = ""
						botstate.Commands.Cancelall = ""
						newsend += "Done restart all Cmd.\n"
					}
				}
			} else if cmd == "cleargowner" {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(room.Gowner) != 0 {
							handler.LogAccess(client, to, sender, "cleargowner", room.Gowner, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gownerlist\n", len(room.Gowner))
							room.Gowner = []string{}
						} else {
							newsend += "Gowner list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmaster" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Master) != 0 {
							newsend += fmt.Sprintf("Cleared %v masterlist\n", len(botstate.UserBot.Master))
							handler.LogAccess(client, to, sender, "clearmaster", botstate.UserBot.Master, msg.ToType)
							botstate.UserBot.ClearMaster()
						} else {
							newsend += "Master list is empty.\n"
						}
					}
				}
			} else if cmd == "clearfuck" {
				rngcmd := handler.GetComs(5, "clearfuck")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Fucklist) != 0 {
							handler.LogAccess(client, to, sender, "clearfuck", botstate.Banned.Fucklist, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v fucklist\n", len(botstate.Banned.Fucklist))
							botstate.Banned.Fucklist = []string{}
						} else {
							newsend += "Fuck list is empty.\n"
						}
					}
				}
			} else if cmd == "clearmute" {
				rngcmd := handler.GetComs(5, "clearmute")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Mutelist) != 0 {
							handler.LogAccess(client, to, sender, "clearmute", botstate.Banned.Mutelist, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v fucklist\n", len(botstate.Banned.Mutelist))
							botstate.Banned.Mutelist = []string{}
						} else {
							newsend += "Mute list is empty.\n"
						}
					}
				}
			} else if cmd == "clearallprotect" {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						linetcr.ClearProtect()
						newsend += "Cleared allprotected.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "perm ") {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						ditha := strings.ReplaceAll(cmd, "perm ", "")
						cmdLil := strings.Split(ditha, " ")
						handler.Addpermcmd(client, to, cmdLil[0], cmdLil[1])
					}
				}
			} else if cmd == "permlist" {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						list := handler.PerCheckList()
						if list != "" {
							newsend += list
						} else {
							newsend += "Not have perm in list.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ajsjoin") {
				rngcmd := handler.GetComs(5, "ajsjoin")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "ajsjoin"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "inv" {
							if botstate.Ajsjoin != "inv" {
								newsend += "botstate.Ajsjoin Invite enabled.\n"
							} else {
								newsend += "botstate.Ajsjoin Already Invite.\n"
							}
						} else if str == "qr" {
							if botstate.Ajsjoin != "qr" {
								newsend += "botstate.Ajsjoin qr enabled.\n"
							} else {
								newsend += "botstate.Ajsjoin Already qr.\n"
							}
						} else if str == "off" {
							if botstate.Ajsjoin != "off" {
								newsend += fmt.Sprintf("botstate.Ajsjoin %s disabled.\n", botstate.Ajsjoin)
							} else {
								newsend += "botstate.Ajsjoin Already disabled.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autojoin") {
				rngcmd := handler.GetComs(5, "autojoin")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "autojoin"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "inv" {
							if botstate.Autojoin != "inv" {
								newsend += "botstate.Autojoin Invite enabled.\n"
							} else {
								newsend += "botstate.Autojoin Already Invite.\n"
							}
						} else if str == "qr" {
							if botstate.Autojoin != "qr" {
								newsend += "botstate.Autojoin qr enabled.\n"
							} else {
								newsend += "botstate.Autojoin Already qr.\n"
							}
						} else if str == "off" {
							if botstate.Autojoin != "off" {
								newsend += fmt.Sprintf("botstate.Autojoin %s disabled.\n", botstate.Autojoin)
							} else {
								newsend += "botstate.Autojoin Already disabled.\n"
							}
						}
					}
				}
			} else if cmd == "mutelist" {
				rngcmd := handler.GetComs(5, "mutelist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Mutelist) != 0 {
							list := "Mutelist:"
							client.SendPollMention(to, list, botstate.Banned.Mutelist)
						} else {
							newsend += "Mute list is empty.\n"
						}
					}
				}
			} else if cmd == "fucklist" {
				rngcmd := handler.GetComs(5, "fucklist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Fucklist) != 0 {
							list := "Fucklist:"
							client.SendPollMention(to, list, botstate.Banned.Fucklist)
						} else {
							newsend += "Fuck list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "groupcast") {
				rngcmd := handler.GetComs(5, "upcast")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "groupcast"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							gr, _ := client.GetGroupIdsJoined()
							for _, gi := range gr {
								client.SendMessage(gi, botstate.Fancy(str))
								go client.SendContact(gi, sender)
							}
							newsend += "Success broadcast to " + strconv.Itoa(len(gr)) + " group\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "friendcast") {
				rngcmd := handler.GetComs(5, "ndcast")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "friendcast"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							n := 0
							for _, cl := range botstate.ClientBot {
								fl, _ := cl.GetAllContactIds()
								for _, x := range fl {
									client.SendMessage(x, botstate.Fancy(str))
									go client.SendContact(x, sender)
									time.Sleep(3 * time.Second)
								}
							}
							n += 1
							newsend += "Success broadcast to all friends\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bcimage") {
				rngcmd := handler.GetComs(5, "bcimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "bcimage"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							handler.AddConSingle(result)
							for _, gi := range result {
								list := " "+result[2]
								client.SendMessage(gi, botstate.Fancy(list))
								image := "https://"+result[1]
								
								client.SendFoto(gi, image)
								if !utils.InArrayString(botstate.MidBc, gi) {
								}
							}
							newsend += "Success broadcast image to friends\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bcast") {
				rngcmd := handler.GetComs(5, "bcast")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "bcast"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							handler.AddConSingle(result)
							for _, gi := range result {
								list := " "+result[1]
								client.SendMessage(gi, botstate.Fancy(list))
								go client.SendContact(gi, sender)
								if !utils.InArrayString(botstate.MidBc, gi) {
								}
							}
							newsend += "Success broadcast to friends\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "list broadcast") {
				rngcmd := handler.GetComs(5, "list broadcast")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.MidBc) != 0 {
							list :="List Broadcast:"
							cuh, _ := client.GetContacts(botstate.MidBc)
							for num, prs := range cuh {
								num++
								rengs := strconv.Itoa(num)
								name := prs.DisplayName
								list += fmt.Sprintf("\n ."+rengs+" %v", name)
							}
							newsend += list+"\n"
						} else {
							newsend += "Notings"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "accgroupid") {
				rngcmd := handler.GetComs(5, "accgroupid")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "accgroupid"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							client.AcceptGroupInvitationNormal(str)
							exe := []*linetcr.Account{}
							for _, p := range bk {
								if p.Limited == false {
									if botstate.Err == nil {
										exe = append(exe, p)
									}
								}
							}
							if len(exe) != 0 {
								newsend += "Succes Accept Group ID"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unicode") {
				rngcmd := handler.GetComs(5, "unicode")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "unicode"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							handler.AddConSingle(result)
							no, _ := strconv.Atoi(result[1])
							for _, gi := range result {
								for i := 1; i <= no; i++ {
									time.Sleep(time.Second * 3)
									for _, p := range bk {
										filepath := fmt.Sprintf("unicode.txt")
										b, botstate.Err := ioutil.ReadFile(filepath)
										if botstate.Err != nil {
											fmt.Print(botstate.Err)
										}
										code := string(b)
										list := code
										p.SendMessage(gi, botstate.Fancy(list))
									}
								}
							}
							newsend += "Success send unicode\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gcreate") {
				rngcmd := handler.GetComs(5, "gcreate")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "gcreate"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, ":")
							handler.AddConSingle(result)
							no, _ := strconv.Atoi(result[1])
							for i := 1; i <= no; i++ {
								for _, p := range bk {
									call.CreateGroup("test", result, string(p.AuthToken))
								}
							}
							newsend += "Success create new group\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gourl ") {
				rngcmd := handler.GetComs(5, "gourl")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						num, botstate.Err := strconv.Atoi(result[1])
						if botstate.Err != nil {
							newsend += "invalid number.\n"
						} else {
							gr := []string{}
							for i := range botstate.ClientBot {
								grs, _ := botstate.ClientBot[i].GetGroupIdsJoined()
								for _, a := range grs {
									if !utils.InArrayString(gr, a) {
										gr = append(gr, a)
									}
								}
							}
							groups, _ := client.GetGroups(gr)
							for _, gi := range groups {
							}
							if num > 0 && num <= len(botstate.Tempgroup) {
								gid := botstate.Tempgroup[num-1]
								tick, botstate.Err := client.ReissueChatTicket(gid)
								if botstate.Err == nil {
									mes := make(chan bool)
									go func() {
										if botstate.Err != nil {
											mes <- false
										} else {
											mes <- true
										}
									}()

									newsend += "https://line.me/R/ti/g/" + tick + "\n"
								}
							} else {
								newsend += "out of range.\n"
							}
						}
					}
				}
			} else if cmd == "allgroups" {
				rngcmd := handler.GetComs(4, "allgroups")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							nm := []string{}
							gr, _ := p.GetGroupIdsJoined()
							for c, a := range gr {
								name, _ := p.GetGroupMember(a)
								c += 1
								name = fmt.Sprintf("%v. %s", c, name)
								nm = append(nm, name)
								handler.GetSquad(p, a)
							}
							
							nm1 := []string{}
							gr1, _ := p.GetGroupIdsInvited()
							for c1, a1 := range gr1 {
								name1, _ := p.GetGroupMember(a1)
								c1 += 1
								name1 = fmt.Sprintf("%v. %s", c1, name1)
								nm1 = append(nm1, name1)
								handler.GetSquad(p, a1)
							}
							stf := "Group list:\n\n"
							str := strings.Join(nm, "\n\n")
							stf1 := "\n\nPending list:\n\n"
							str1 := strings.Join(nm1, "\n")
							p.SendText(to, stf+str+stf1+str1)
						}
					}
				}
			} else if cmd == "groups" {
				rngcmd := handler.GetComs(5, "groups")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						gr := []string{}
						for _, p := range botstate.ClientBot {
							if !p.Frez {
								grs, _ := p.GetGroupIdsJoined()
								for _, a := range grs {
									if !utils.InArrayString(gr, a) {
										gr = append(gr, a)
									}
								}
							}
						}
						nm := []string{}
						grup, _ := client.GetGroups(gr)
						ci := []string{}
						for _, g := range grup {
							ci = append(ci, strings.ToLower(g.ChatName))
						}
						sort.Strings(ci)
						groups := []*talkservice.Chat{}
						for _, x := range ci {
							for _, gi := range grup {
								if strings.ToLower(gi.ChatName) == x {
									if !handler.InArrayChat(groups, gi) {
										groups = append(groups, gi)
									}
								}
							}
						}
						for c, a := range groups {
							name, mem := a.ChatName, a.Extra.GroupExtra.MemberMids
							c += 1
							jm := 0
							for mid := range mem {
								if utils.InArrayString(botstate.Squadlist, mid) {
									jm++
								}
							}
							name = fmt.Sprintf("%v. %s (%v/%v)", c, name, jm, len(mem))
							nm = append(nm, name)
							handler.GetSquad(client, a.ChatMid)
						}
						stf := "All Group List:\n\n"
						str := strings.Join(nm, "\n")
						anu := []string{}
						for _, p := range botstate.ClientBot {
							if !p.Frez {
								grs, _ := p.GetGroupIdsInvited()
								for _, a := range grs {
									if !utils.InArrayString(gr, a) && !utils.InArrayString(anu, a) {
										anu = append(anu, a)
									}
								}
							}
						}
						grup, _ = client.GetGroups(anu)
						ci = []string{}
						for _, g := range grup {
							ci = append(ci, strings.ToLower(g.ChatName))
						}
						sort.Strings(ci)
						groups = []*talkservice.Chat{}
						for _, x := range ci {
							for _, gi := range grup {
								if strings.ToLower(gi.ChatName) == x {
									if !handler.InArrayChat(groups, gi) {
										groups = append(groups, gi)
									}
								}
							}
						}
						nm = []string{}
						nn := 1
						for _, a := range groups {
							name, mem, inv := a.ChatName, a.Extra.GroupExtra.MemberMids, a.Extra.GroupExtra.InviteeMids
							if name != "" {
								jm := 0
								for mid := range inv {
									if utils.InArrayString(botstate.Squadlist, mid) {
										jm++
									}
								}
								if jm != 0 {
									name = fmt.Sprintf("%v. %s (invited) (%v/%v)", nn, name, jm, len(mem))
									nm = append(nm, name)
									handler.GetSquad(client, a.ChatMid)
									nn++
								} else {
								}
							} else {
							}
						}
						var strs, strsa = "", ""
						if len(nm) != 0 {
							strs = "\n\nAll Group Invitation:\n\n"
							strsa = strings.Join(nm, "\n")
						}
						newsend += stf + str + strs + strsa
					}
				}
			} else if strings.HasPrefix(cmd, "nukejoin ") {
				rngcmd := handler.GetComs(5, "nukejoin")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "nukejoin ", "", 1)
						if spl == "on" {
							newsend += "Nukejoin is enabled.\n"
						} else if spl == "off" {
							newsend += "Nukejoin is disabled.\n"
						}
					}
				}

			} else if strings.HasPrefix(cmd, "botstate.Canceljoin ") {
				rngcmd := handler.GetComs(5, "botstate.Canceljoin")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "botstate.Canceljoin ", "", 1)
						if spl == "on" {
							newsend += "botstate.Canceljoin is enabled.\n"
						} else if spl == "off" {
							newsend += "botstate.Canceljoin is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autopro ") {
				rngcmd := handler.GetComs(5, "autopro")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autopro ", "", 1)
						if spl == "on" {
							newsend += "Autopro is enabled.\n"
						} else if spl == "off" {
							newsend += "Autopro is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "detectcall ") {
				rngcmd := handler.GetComs(5, "detectcall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "detectcall ", "", 1)
						if spl == "on" {
							newsend += "Detectgroupcall is enabled.\n"
						} else if spl == "off" {
							newsend += "Detectgroupcall is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "kickbanqr ") {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "kickbanqr ", "", 1)
						if spl == "on" {
							newsend += "KickbanQr is enabled.\n"
						} else if spl == "off" {
							newsend += "KickbanQr is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autobackbot ") {
				rngcmd := handler.GetComs(5, "autobackbot")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autobackbot ", "", 1)
						if spl == "on" {
							newsend += "Autobackbot is enabled.\n"
						} else if spl == "off" {
							newsend += "Autobackbot is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autopurge ") {
				rngcmd := handler.GetComs(5, "autopurge")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autopurge ", "", 1)
						if spl == "on" {
							newsend += "Autopurge is enabled.\n"
						} else if spl == "off" {
							newsend += "Autopurge is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autoticket") {
				rngcmd := handler.GetComs(5, "autoticket")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "autoticket"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "normal" {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("botstate.AutoJointicket type normal"))
						} else if str == "nuke" {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("botstate.AutoJointicket type nuke"))
						} else if str == "off" {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("botstate.AutoJointicket is disabled"))
						}
						if count != 0 {
							//newsend += fmt.Sprintf("Autobroadcast type %s", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autotrans") {
				rngcmd := handler.GetComs(5, "autotrans")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "autotrans"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "off" {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("botstate.AutoTranslate is disabled"))
						} else if len(str) != 0 {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("botstate.AutoTranslate type "+str))
						}
						if count != 0 {
							//newsend += fmt.Sprintf("botstate.AutoTranslate type %s", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "flex ") {
				rngcmd := handler.GetComs(5, "flex")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "flex ", "", 1)
						if spl == "1" {
							linetcr.FlexMode = true
							linetcr.FooterMode = false
							linetcr.FlexMode2 = false
							newsend += "Flexmode 1 is enabled.\n"
						} else if spl == "2" {
							linetcr.FlexMode2 = true
							linetcr.FooterMode = false
							linetcr.FlexMode = false
							newsend += "Flexmode 2 is enabled.\n"
						} else if spl == "footer" {
							linetcr.FooterMode = true
							linetcr.FlexMode2 = false
							linetcr.FlexMode = false
							newsend += "Footermode is enabled.\n"
						} else if spl == "off" {
							linetcr.FooterMode = false
							linetcr.FlexMode = false
							linetcr.FlexMode2 = false
							newsend += "Flexmode is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "liff ") {
				rngcmd := handler.GetComs(5, "liff")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "liff ", "", 1)
						if spl == "1" {
							linetcr.Liffid = "1655425084-3OQ8Mn9J"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "2" {
							linetcr.Liffid = "1656652460-LGKR2XXJ"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "3" {
							linetcr.Liffid = "1655623470-81eDd9kM"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "4" {
							linetcr.Liffid = "1653779160-yw2l2v9d"
							newsend += "Liff update:\n  "+linetcr.Liffid
						} else if spl == "5" {
							linetcr.Liffid = "1657707255-WVxqmM35"
							newsend += "Liff update:\n  "+linetcr.Liffid
						}
					}
				}
			} else if strings.HasPrefix(cmd, "lockajs ") {
				rngcmd := handler.GetComs(5, "lockajs")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "lockajs ", "", 1)
						if spl == "on" {
							newsend += "Lockajs is enabled.\n"
						} else if spl == "off" {
							newsend += "Lockajs is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "modepro ") {
				rngcmd := handler.GetComs(5, "modepro")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "modepro ", "", 1)
						if spl == "on" {
							newsend += "Mode protect is enabled.\n"
						} else if spl == "off" {
							newsend += "Mode protect is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "autoban ") {
				rngcmd := handler.GetComs(5, "autoban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "autoban ", "", 1)
						if spl == "on" {
							newsend += "Autoban is enabled.\n"
						} else if spl == "off" {
							newsend += "Autoban is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "lockmode ") {
				rngcmd := handler.GetComs(5, "lockmode")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "lockmode ", "", 1)
						if spl == "on" {
							newsend += "Lockmode is enabled.\n"
						} else if spl == "off" {
							newsend += "Lockmode is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "modewar ") {
				rngcmd := handler.GetComs(5, "modewar")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "modewar ", "", 1)
						if spl == "on" {
							newsend += "Mode war is enabled.\n"
						} else if spl == "off" {
							newsend += "Mode war is disabled.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "banpurge ") {
				rngcmd := handler.GetComs(5, "banpurge")
				if handler.CheckPermission(rngcmd, sender, to) {
					spl := strings.Replace(cmd, "banpurge ", "", 1)
					if spl == "on" {
						newsend += "Banpurge is enabled.\n"
					} else if spl == "off" {
						newsend += "Banpurge is disabled.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "groupinfo ") {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										list := handler.InfoGroup(client, gid)
										client.SendMessage(to, botstate.Fancy(list))
									} else {
										newsend += "out of range.\n"
									}
								} else {
									newsend += "invalid range.\n"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remotegroup ") {
				rngcmd := handler.GetComs(4, "remotegroup")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										if !utils.InArrayString(botstate.Sinderremote, sender) {
										}
										names, _, _ := client.GetChatList(gid)
										handler.GetSquad(client, gid)
										ret := fmt.Sprintf("Group: %v\n\n Send your command.\n", names)
										newsend += ret
									} else {
										newsend += "out of range.\n"
									}
								} else {
									newsend += "invalid range.\n"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remoteowner ") {
				rngcmd := handler.GetComs(4, "remoteowner")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !utils.InArrayString(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !utils.InArrayString(botstate.MidRemote, mid) {
												}
											}
											client.SendMessage(to, botstate.Fancy(list))
										}
										client.SendMessage(to, botstate.Fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remotemaster ") {
				rngcmd := handler.GetComs(4, "remotemaster")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !utils.InArrayString(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !utils.InArrayString(botstate.MidRemote, mid) {
												}
											}
											client.SendMessage(to, botstate.Fancy(list))
										}
										client.SendMessage(to, botstate.Fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remoteadmin ") {
				rngcmd := handler.GetComs(4, "remoteadmin")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !utils.InArrayString(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !utils.InArrayString(botstate.MidRemote, mid) {
												}
											}
											client.SendMessage(to, botstate.Fancy(list))
										}
										client.SendMessage(to, botstate.Fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remotecontact ") {
				rngcmd := handler.GetComs(4, "remotecontact")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !utils.InArrayString(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !utils.InArrayString(botstate.MidRemote, mid) {
												}
											}
											client.SendMessage(to, botstate.Fancy(list))
										}
										client.SendMessage(to, botstate.Fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "remoteban ") {
				rngcmd := handler.GetComs(4, "remoteban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							result2, botstate.Err := strconv.Atoi(result[1])
							if botstate.Err != nil {
								client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
								return
							} else {
								if result2 > 0 {
									if len(botstate.Tempgroup) == 0 {
										client.SendMessage(to, botstate.Fancy("Please input the right number\nSee group number with command groups"))
										return
									}
									nim, _ := strconv.Atoi(result[1])
									nim = nim - 1
									if result2 > 0 && result2 < len(botstate.Tempgroup)+1 {
										gid := botstate.Tempgroup[nim]
										_, mem, _ := client.GetChatList(gid)
										Glist := []string{}
										for _, from := range mem {
											if !utils.InArrayString(Glist, from) {
												Glist = append(Glist, from)
											}
										}
										if len(Glist) != 0 {
											list := "Member: \n"
											cuh, _ := client.GetContacts(Glist)
											for num, prs := range cuh {
												num++
												rengs := strconv.Itoa(num)
												name := prs.DisplayName
												list += fmt.Sprintf("\n   %s. %s", rengs, name)
												mid := prs.Mid
												if !utils.InArrayString(botstate.MidRemote, mid) {
												}
											}
											client.SendMessage(to, botstate.Fancy(list))
										}
										client.SendMessage(to, botstate.Fancy("Please send number member targets"))
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unfuck") {
				rngcmd := handler.GetComs(4, "unfuck")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 2
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.handler.CheckUnbanBots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unfuck"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.handler.CheckUnbanBots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "admin") && cmd != "admins" {
				if handler.CheckPermission(6, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "gowners" {
				if handler.CheckPermission(8, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(room.Gowner) != 0 {
							list := "𝗴𝗼𝘄𝗻𝗲𝗿𝘀:\n"
							for num, xd := range room.Gowner {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Gowner list is empty.\n"
						}
					}
				}
			} else if cmd == "masters" {
				if handler.CheckPermission(6, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Master) != 0 {
							list := " 𝗺𝗮𝘀𝘁𝗲𝗿𝘀:\n"
							for num, xd := range botstate.UserBot.Master {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Master list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "unadmin") {
				if handler.CheckPermission(6, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 4
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unadmin"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Admin)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungadmin") {
				if handler.CheckPermission(8, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 7
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("??𝗹𝗲????𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "ungadmin"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, room.Gadmin)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "squadmid" {
				rngcmd := handler.GetComs(5, "squadmid")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						e, _ := client.GetProfile()
						list := "All bot's mid\n\n"
						list += "1." + e.DisplayName + "\n\n"
						list += client.MID
						for b, a := range client.Squads {
							b++
							x, _ := client.GetContact(a)
							list += fmt.Sprintf("\n\n%v. %s ", b+1, x.DisplayName)
							list += "\n\n" + a
						}
						newsend += list + "\n"
					}
				}
			} else if cmd == "check all" {
				rngcmd := handler.GetComs(5, "check all")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						e, _ := client.GetProfile()
						list := "All bot's: ↓\n\n"
						list += "1." + e.DisplayName + "\nMid "
						list += client.MID
						for b, a := range client.Squads {
							b++
							x, _ := client.GetContact(a)
							list += fmt.Sprintf("\n\n%v. %s ", b+1, x.DisplayName)
							list += "\nMid " + a
						}
						list += "\n\nBot's limited: ↓"
						for n, cl := range linetcr.KickBans {
							m := cl.MID
							no := n + 1
							pr, _ := client.GetContact(m)
							cl.Namebot = pr.DisplayName
							list += fmt.Sprintf("\n\n%v. %s\nMid %v", no, cl.Namebot, cl.MID)
						}
						newsend += list + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "checkban") && cmd != "checkbans" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							newlist := []string{}
							for _, user := range lists {
								if !utils.InArrayString(newlist, user) {
									newlist = append(newlist, user)
								}
							}
							numr := 0
							tex := "Status Account:\n"
							for _, user := range newlist {
								name := client.GetName(user)
								if name == "" {
									name = "Deleted Account"
								}
								numr ++
								tex += fmt.Sprintf("\n%v. %v", numr, name)
								time.Sleep(100 * time.Millisecond)
								r, _ := client.GetHomeProfile(user)
								if linetcr.GetBannedChat(r) == 1 {
									tex += fmt.Sprintf("\n   Status: botstate.Banned\n")
								} else {
									tex += fmt.Sprintf("\n   Status: Normal\n")
								}
							}
							newsend += tex
						}
					}
				}
			} else if strings.HasPrefix(cmd, "wordbanadd") {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "wordbanadd"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if !utils.InArrayString(botstate.Data.WordbanBack, str) {
							botstate.Data.WordbanBack = append(botstate.Data.WordbanBack, str)
							handler.SaveBackup()
							newsend += "Wordban added : " + str + "\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "wordbandel") {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "wordbandel"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if utils.InArrayString(botstate.Data.WordbanBack, str) {
							botstate.Data.WordbanBack = utils.RemoveString(botstate.Data.WordbanBack, str)

							client.SendMessage(to, botstate.Fancy("Wordban delete : "+str))
						}
						handler.SaveBackup()
					} else {
						client.SendMessage(to, botstate.Fancy("No wordban deleted"))
					}
				}
			} else if strings.HasPrefix(cmd, "wordbanlist") {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len (botstate.Data.WordbanBack) != 0 {
							list :="Wordban List: ↓"
							for num, xd := range botstate.Data.WordbanBack {
								num++
								rengs := strconv.Itoa(num)
								list += "\n\n  "+rengs+". "+xd
							}
							newsend += list+"\n"
						} else {
							newsend += "Notings"
						}
					}
				}
			} else if cmd == "wordbanclear" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Data.WordbanBack) != 0 {
							newsend += fmt.Sprintf("Cleared %v wordbanlist", len(botstate.Data.WordbanBack)) + "\n"
							botstate.Data.WordbanBack = []string{}
							handler.SaveBackup()
						} else {
							newsend += "Wordban is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gadmin") && cmd != "gadmins" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 7
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "admins" {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Admin) != 0 {
							list := "𝗮𝗱𝗺𝗶𝗻𝘀:\n"
							for num, xd := range botstate.UserBot.Admin {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Admin list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "antitag ") {
				rngcmd := handler.GetComs(5, "antitag")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						spl := strings.Replace(cmd, "antitag ", "", 1)
						if spl == "on" {
							room.AntiTag = true
							newsend += "antitag enabled.\n"
						} else if spl == "off" {
							room.AntiTag = false
							newsend += "antitag disabled.\n"
						}
					}
				}
			} else if cmd == "banlist" || cmd == botstate.Commands.Banlist && botstate.Commands.Banlist != "" {
				rngcmd := handler.GetComs(7, "banlist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Banlist) != 0 {
							listbl := "banlist:"
							client.SendPollMention(to, listbl, botstate.Banned.Banlist)
						} else {
							newsend += "Ban list is empty.\n"
						}
					}
				}
			} else if cmd == "locklist" {
				rngcmd := handler.GetComs(5, "locklist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Locklist) != 0 {
							listbl := "locklist:"
							client.SendPollMention(to, listbl, botstate.Banned.Locklist)
						} else {
							newsend += "Lock list is empty.\n"
						}
					}
				}
			} else if cmd == "respon" || cmd == botstate.Commands.Respon && botstate.Commands.Respon != "" {
				rngcmd := handler.GetComs(5, "respon")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							go p.SendMessage(to, botstate.Fancy(botstate.MsgRespon))
						}
					}
				}
			} else if cmd == "rollcall" || cmd == botstate.Commands.Botname && botstate.Commands.Botname != "" {
				rngcmd := handler.GetComs(5, "rollcall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							asss := fmt.Sprintf("%v", p.Namebot)
							go p.SendMessage(to, botstate.Fancy(asss))
						}
					}
				}
			} else if cmd == "upallimage" || cmd == botstate.Commands.Upallimage && botstate.Commands.Upallimage != "" {
				rngcmd := handler.GetComs(4, "upallimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							if !linetcr.Checkarri(botstate.Changepic, p) {
							}
						}
						client.SendMessage(to, botstate.Fancy("Send image."))
					}
				}
			} else if cmd == "upallcover" || cmd == botstate.Commands.Upallcover && botstate.Commands.Upallcover != "" {
				rngcmd := handler.GetComs(4, "upallcover")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							if !linetcr.Checkarri(botstate.Changepic, p) {
							}
						}
						client.SendMessage(to, botstate.Fancy("Send image."))
					}
				}
			} else if cmd == "unsend" || cmd == botstate.Commands.Unsend && botstate.Commands.Unsend != "" {
				rngcmd := handler.GetComs(7, "unsend")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.UnsendChat(to)
						}
					}
				}
			} else if cmd == "upvallimage" || cmd == botstate.Commands.Upvallimage && botstate.Commands.Upvallimage != "" {
				rngcmd := handler.GetComs(4, "upvallimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							if !linetcr.Checkarri(botstate.Changepic, p) {
							}
						}
						client.SendMessage(to, botstate.Fancy("Send video."))
					}
				}
			} else if cmd == "upvallcover" || cmd == botstate.Commands.Upvallcover && botstate.Commands.Upvallcover != "" {
				rngcmd := handler.GetComs(4, "upvallcover")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							if !linetcr.Checkarri(botstate.Changepic, p) {
							}
						}
						client.SendMessage(to, botstate.Fancy("Send video."))
					}
				}
			} else if cmd == "appname" || cmd == botstate.Commands.Appname && botstate.Commands.Appname != "" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, botstate.Fancy(string(p.AppName)))
						}
					}
				}
			} else if cmd == "cektoken" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, botstate.Fancy(string(p.AuthToken)))
						}
					}
				}
			} else if cmd == "useragent" || cmd == botstate.Commands.Useragent && botstate.Commands.Useragent != "" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, botstate.Fancy(string(p.UserAgent)))
						}
					}
				}
			} else if cmd == "hostname" || cmd == botstate.Commands.Hostname && botstate.Commands.Hostname != "" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, botstate.Fancy(string(p.Host)))
						}
					}
				}
			} else if cmd == "location" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, botstate.Fancy(string(p.Locale)))
						}
					}
				}
			} else if cmd == "friends" || cmd == botstate.Commands.Friends && botstate.Commands.Friends != "" {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						exe2 := []*linetcr.Account{}
						for _, mid := range mentionlist {
							if utils.InArrayString(botstate.Squadlist, mid) {
								cl := handler.GetKorban(mid)
								exe2 = append(exe2, cl)
							}
						}
						if len(exe2) != 0 {
							for _, p := range exe2 {
								friends, _ := p.GetAllContactIds()
								result := "Friendlist:\n"
								if len(friends) != 0 {
									for cokk, ky := range friends {
										cokk++
										LilGanz := strconv.Itoa(cokk)
										haniku, _ := p.GetContact(ky)
										result += "\n" + LilGanz + ". " + haniku.DisplayName
									}
									client.SendMessage(to, botstate.Fancy(result))
								} else {
									client.SendMessage(to, botstate.Fancy("Friend is empty."))
								}
							}
						} else {
							client.SendMessage(to, botstate.Fancy("Mention Bot First."))
						}
					}
				}
			} else if cmd == "adds" || cmd == botstate.Commands.Adds && botstate.Commands.Adds != "" {
				if handler.CheckPermission(1, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						addb := len(linetcr.Waitadd)
						kb := ""
						if addb != 0 {
							kb += fmt.Sprintf("%v/%v bot's got add/friend banned.", addb, len(botstate.Squadlist))
							for n, cl := range linetcr.Waitadd {
								m := cl.MID
								no := n + 1
								go client.SendContact(to, m)
								var ta time.Duration
								if _, ok := linetcr.BlockAdd.Get(cl.MID); ok {
									t := cl.Timeadd.Add(24 * time.Hour)
									ta = t.Sub(time.Now())
								} else {
									t := cl.Timeadd.Add(1*time.Hour )
									ta = t.Sub(time.Now())
								}
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\nRemaining %v", no, cl.Namebot, handler.FmtDurations(ta))
							}
						}
						if addb == 0 {
							newsend += "All fixed."
						} else {
							newsend += kb
						}
					}
				}
			} else if cmd == "cek" || cmd == botstate.Commands.Limits && botstate.Commands.Limits != "" {
				rngcmd := handler.GetComs(5, "limits")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							var asss string
							if p.Limited == true {
								asss += botstate.MsLimit
							} else {
								asss += botstate.MsFresh
							}
							p.SendMessage(to, botstate.Fancy(asss))
						}
					}
				}
			} else if cmd == "addallsquads" || cmd == botstate.Commands.Addallsquads && botstate.Commands.Addallsquads != "" {
				rngcmd := handler.GetComs(2, "addallsquads")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.SendMessage(to, "Waiting for Add Squad")
						handler.AddConSqV2(botstate.Squadlist)
						friends, _ := client.GetAllContactIds()
						result := "Success added friends:\n"
						if len(friends) != 0 {
							for cokk, ky := range friends {
								cokk++
								LilGanz := strconv.Itoa(cokk)
								haniku, _ := client.GetContact(ky)
								result += "\n" + LilGanz + ". " + haniku.DisplayName
							}
							client.SendMessage(to, botstate.Fancy(result))
						}
					}
				}
			} else if cmd == "unfriendbans" {
				if handler.CheckPermission(2, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.RemBanFriends(client, to)
					}
				}
			} else if cmd == "clearfriends" {
				rngcmd := handler.GetComs(2, "clear friends")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.SendMessage(to, "Waiting for Clear Friend")
						handler.ClearCon()
						asss := "Success clear allfriends."
						client.SendMessage(to, botstate.Fancy(asss))
					}
				}
			} else if cmd == "leave" || cmd == botstate.Commands.Leave && botstate.Commands.Leave != "" {
				rngcmd := handler.GetComs(7, "leave")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, mem := client.GetGroupInvitation(to)
						anu := []string{}
						for m := range mem {
							if utils.InArrayString(botstate.Squadlist, m) {
								anu = append(anu, m)
							}
						}
						if len(anu) != 0 {
							for _, mid := range anu {
								cl := handler.GetKorban(mid)
								cl.AcceptGroupInvitationNormal(to)
							}
						}
						handler.GetSquad(client, to)
						room := linetcr.GetRoom(to)
						bk = room.Client
						for _, cl := range bk {
							go cl.LeaveGroup(to)
						}
						if botstate.LogGroup == to {
						}
						linetcr.SquadRoom = linetcr.RemoveRoom(linetcr.SquadRoom, room)
						handler.SaveBackup()
						handler.LogAccess(client, to, sender, "leave", []string{}, msg.ToType)
					}
				}
			} else if cmd == "ping" || cmd == botstate.Commands.Ping && botstate.Commands.Ping != "" {
				rngcmd := handler.GetComs(5, "ping")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							p.SendMessage(to, botstate.Fancy("pong"))
						}
					}
				}
			} else if cmd == "count" || cmd == botstate.Commands.Count && botstate.Commands.Count != "" {
				rngcmd := handler.GetComs(5, "count")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for n, p := range bk {
							p.SendMessage(to, botstate.Fancy(fmt.Sprintf("%v", n+1)))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "sayall") {
				rngcmd := handler.GetComs(5, "sayall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						str := ""
						var su = "sayall"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						for _, p := range bk {
							p.SendMessage(to, botstate.Fancy(str))
						}
					}
				}
			} else if cmd == "limitout" || cmd == botstate.Commands.Limitout && botstate.Commands.Limitout != "" {
				rngcmd := handler.GetComs(8, "out")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							if p.Limited == true {
								p.LeaveGroup(to)
							}
						}
						handler.GetSquad(client, to)
					}
				}
			} else if strings.HasPrefix(cmd, "upallstatus") {
				rngcmd := handler.GetComs(4, "upallstatus")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							str := ""
							var su = "upallstatus"
							if strings.HasPrefix(text, Rname+" ") {
								str = strings.Replace(text, Rname+" "+su+" ", "", 1)
								str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname+" ") {
								str = strings.Replace(text, Sname+" "+su+" ", "", 1)
								str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Rname) {
								str = strings.Replace(text, Rname+su+" ", "", 1)
								str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname) {
								str = strings.Replace(text, Sname+su+" ", "", 1)
								str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
							}
							for num, p := range bk {
								num++
								rengs := strconv.Itoa(num)
								name := str+" "+rengs
								if handler.TimeDown(num) {
									p.UpdateProfileBio(name)
									p.SendMessage(to, botstate.Fancy("Profile Bio updated."))
								}
							}
						} else {
							client.SendMessage(to, botstate.Fancy("Add Bio first."))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "upallname") {
				rngcmd := handler.GetComs(4, "upallname")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						result := strings.Split((cmd), " ")
						if len(result) > 1 {
							var str string
							var su = "upallname"
							if strings.HasPrefix(text, Rname+" ") {
								str = strings.Replace(text, Rname+" "+su+" ", "", 1)
								str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname+" ") {
								str = strings.Replace(text, Sname+" "+su+" ", "", 1)
								str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Rname) {
								str = strings.Replace(text, Rname+su+" ", "", 1)
								str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
							} else if strings.HasPrefix(text, Sname) {
								str = strings.Replace(text, Sname+su+" ", "", 1)
								str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
							}
							aa := utf8.RuneCountInString(str)
							if aa != 0 && aa <= 20 {
								for num, p := range bk {
									num++
									rengs := strconv.Itoa(num)
									name := str+" "+rengs
									if handler.TimeDown(num) {
										p.UpdateProfileName(name)
										p.SendMessage(to, botstate.Fancy("Profile name success updated."))
									}
								}
							}
						} else {
							client.SendMessage(to, botstate.Fancy("Add name first."))
						}
					}
				}
			} else if cmd == "autoname" {
				rngcmd := handler.GetComs(4, "autoname")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						seed := time.Now().UTC().UnixNano()
						nameGenerator := namegenerator.NewNameGenerator(seed)
						for i,x := range botstate.Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								time.Sleep(1 * time.Second)
								if botstate.Err != nil {
									fmt.Println(botstate.Err)
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update profile name failure."))
								} else {
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update profile name done."))
								}
							}
						}
					}
				}
			} else if cmd == "randomprofile" {
				rngcmd := handler.GetComs(4, "randomprofile")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						seed := time.Now().UTC().UnixNano()
						nameGenerator := namegenerator.NewNameGenerator(seed)
						for i,x := range botstate.Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								botstate.ClientBot[i].UpdateProfileName(nameGenerator.Generate())
								link := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								link1 := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								botstate.ClientBot[i].UpdateCoverWithURL(link1)
								if botstate.Err != nil {
									fmt.Println(botstate.Err)
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update random profile failure."))
								} else {
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update random profile done."))
								}
							}
						}
					}
				}
			} else if cmd == "1autoimage" {
				rngcmd := handler.GetComs(4, "1autoimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for i,x := range botstate.Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								link := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								if botstate.Err != nil {
									fmt.Println(botstate.Err)
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update picture failure."))
								} else {
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update picture done."))
								}
							}
						}
					}
				}
			} else if cmd == "autocover" {
				rngcmd := handler.GetComs(4, "autocover")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for i,x := range botstate.Squadlist {
							if linetcr.IsMembers(client, to, x) == true {
								link := "https://api.zahwazein.xyz/randomimage/cecan?apikey=zenzkey_91cc3a9222"
								if botstate.Err != nil {
									fmt.Println(botstate.Err)
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update cover failure."))
								} else {
									botstate.ClientBot[i].SendMessage(to, botstate.Fancy("Update cover done."))
								}
							}
						}
					}
				}
			} else if cmd == "1autoname" {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for n, p := range bk {
							if handler.TimeDown(n) {
								p.AutoupdateName(p.AuthToken)
								p.SendMessage(to, botstate.Fancy("Success update name"))
							}
						}
					}
				}
			} else if cmd == "autoimage" {
				if handler.CheckPermission(4, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for n, p := range bk {
							if handler.TimeDown(n) {
								p.AutoupdatePict(p.AuthToken)
								p.SendMessage(to, botstate.Fancy("Success update pict"))
							}
						}
					}
				}
			} else if cmd == "clearadmin" {
				if handler.CheckPermission(6, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Admin) != 0 {
							handler.LogAccess(client, to, sender, "clearadmin", botstate.UserBot.Admin, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v adminlist\n", len(botstate.UserBot.Admin))
							botstate.UserBot.ClearAdmin()
						} else {
							newsend += "Admin list is empty.\n"
						}
					}
				}
			} else if cmd == "clearban" || cmd == botstate.Commands.Clearban && botstate.Commands.Clearban != "" {
				rngcmd := handler.GetComs(7, "rban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Banlist) != 0 {
							msgcbn := fmt.Sprintf(botstate.MsgBan, len(botstate.Banned.Banlist))
							handler.LogAccess(client, to, sender, "clearban", botstate.Banned.Banlist, msg.ToType)
							newsend += msgcbn + "\n"
							botstate.Banned.Banlist = []string{}
							botstate.Banned.Exlist = []string{}
						} else {
							newsend += "Ban list is empty.\n"
						}
					}
				}
			} else if cmd == "clearlock" {
				rngcmd := handler.GetComs(5, "clearlock")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.Banned.Locklist) != 0 {
							msgcbn := fmt.Sprintf(botstate.MsgLock, len(botstate.Banned.Locklist))
							handler.LogAccess(client, to, sender, "clearlock", botstate.Banned.Locklist, msg.ToType)
							newsend += msgcbn + "\n"
							botstate.Banned.Locklist = []string{}
							botstate.Banned.Exlist = []string{}
						} else {
							newsend += "Lock list is empty.\n"
						}
					}
				}
			} else if cmd == "cleargadmin" {
				if handler.CheckPermission(8, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(room.Gadmin) != 0 {
							handler.LogAccess(client, to, sender, "cleargadmin", room.Gadmin, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gadminlist\n", len(room.Gadmin))
							room.Gadmin = []string{}
						} else {
							newsend += "Gadmin list is empty.\n"
						}
					}
				}
			} else if cmd == "/list protect" {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						res := linetcr.ListProtect()
						client.SendHelp(to, botstate.Fancy(res + "\n"))
					}
				}
			} else if cmd == "list protect" {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						ret := "✠ List Protect:"
						fukk, _ := client.GetGroupIdsJoined()
						for num, group := range fukk {
							num++
							rengs := strconv.Itoa(num)
							Room := linetcr.GetRoom(group)
							ret += fmt.Sprintf("\n\n%v. %s\n", rengs, Room.Name)
							if Room.ProQr {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ Qʀ"
							} else {
								ret += "\n ⚪ » Protect QR"
							}
							if Room.AntiTag {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							}
							if Room.ProKick {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							}
							if Room.ProInvite {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							}
							if Room.ProCancel {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							}
							if Room.ProJoin {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							}
							if Room.ProName {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							}
							if Room.ProPicture {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							}
							if Room.ProNote {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							}
							if Room.ProAlbum {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							}
							if Room.ProLink {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							}
							if Room.ProFlex {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							}
							if Room.ProImage {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							}
							if Room.ProVideo {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							}
							if Room.ProCall {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							}
							if Room.ProSpam {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							}
							if Room.ProSticker {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							}
							if Room.ProContact {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							}
							if Room.ProPost {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							}
							if Room.ProFile {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							}
							if len(Room.GoMid) > 0 {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							}
						}
						client.SendHelp(to, botstate.Fancy(ret + "\n"))
					}
				}
			} else if cmd == "bringall" || cmd == botstate.Commands.Bringall && botstate.Commands.Bringall != "" {
				rngcmd := handler.GetComs(5, "bringall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if op.Message.ToType != 2 {
							return
						}
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						exe, _ := handler.SelectBot(client, to)
						if exe != nil {
							handler.SetInviteTo(exe, to, exe.Squads)
							handler.LogAccess(client, to, sender, "invite", exe.Squads, msg.ToType)
							time.Sleep(1 * time.Second)
							handler.GetSquad(exe, to)
						} else {
							newsend += "Invite banned try with another bot.\n"
						}
					}
				}
			} else if cmd == "stayall" || cmd == botstate.Commands.Stayall && botstate.Commands.Stayall != "" {
				rngcmd := handler.GetComs(7, "stayall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						//_, mem := client.GetGroupInvitation(to)
						//anu := []string{}
						//for m := range mem {if utils.InArrayString(botstate.Squadlist, m) {anu = append(anu, m)}
						//}
						//if len(anu) != 0 {for _, mid := range anu {cl := handler.GetKorban(mid);cl.AcceptGroupInvitationNormal(to)}
						//}
						//handler.GetSquad(client, to)
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						numb := len(botstate.ClientBot)
						if numb > 0 && numb <= len(botstate.ClientBot) {
							handler.GetSquad(client, to)
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
								handler.GetSquad(client, to)
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
									wi := handler.GetSquad(client, to)
									for i := 0; i < len(all); i++ {
										if i == g {
											break
										}
										l := all[i]
										if l != client && !linetcr.InArrayCl(wi, l) {
											if !l.Limited {
												wg.Add(1)
												go func() {
													l.AcceptTicket(to, ti)
													wg.Done()
												}()
											}
										}
									}
									wg.Wait()
									client.UpdateChatQrV2(to, true)
									handler.GetSquad(client, to)
									handler.LogAccess(client, to, sender, "bringbot", []string{}, 2)
									handler.SaveBackup()
									aa := len(room.Client)
									var name string
									name = fmt.Sprintf("Ready %v bots here", aa)
									newsend += name + "\n"
								}
							}
						}
						if botstate.LockAjs {
							str := botstate.CountAjs
							numb, _ := strconv.Atoi(str)
							if numb == 0 {
								list := append([]*linetcr.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for n, o := range list {
									if n < 2 {
										o.LeaveGroup(to)
										linetcr.GetRoom(to).RevertGo(o)

									} else {
										break
									}
								}
								room := linetcr.GetRoom(to)
								cls := room.Client
								for _, cl := range cls {
									if !cl.Limited {
										for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
										}
										break
									}
								}
							} else {
								list := append([]*linetcr.Account{}, room.Client...)
								sort.Slice(list, func(i, j int) bool {
									return list[i].KickPoint > list[j].KickPoint
								})
								for n, o := range list {
									if n < numb {
										o.LeaveGroup(to)
										linetcr.GetRoom(to).RevertGo(o)
									} else {
										break
									}
								}
								room := linetcr.GetRoom(to)
								cls := room.Cans()
								for _, cl := range cls {
									if !cl.Limited {
										for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
										}
										break
									}
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "go") && cmd != "gojoin" {
				rngcmd := handler.GetComs(8, "o")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						str := strings.Replace(cmd, "go ", "", 1)
						numb, _ := strconv.Atoi(str)
						if numb == 0 {
							list := append([]*linetcr.Account{}, room.Client...)
							sort.Slice(list, func(i, j int) bool {
								return list[i].KickPoint > list[j].KickPoint
							})
							for n, o := range list {
								if n < 2 {
									o.LeaveGroup(to)
									linetcr.GetRoom(to).RevertGo(o)

								} else {
									break
								}
							}
							room := linetcr.GetRoom(to)
							cls := room.Client
							for _, cl := range cls {
								if !cl.Limited {
									for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
									}
									break
								}
							}
						} else {
							list := append([]*linetcr.Account{}, room.Client...)
							sort.Slice(list, func(i, j int) bool {
								return list[i].KickPoint > list[j].KickPoint
							})
							for n, o := range list {
								if n < numb {
									o.LeaveGroup(to)
									linetcr.GetRoom(to).RevertGo(o)
								} else {
									break
								}
							}
							room := linetcr.GetRoom(to)
							cls := room.Cans()
							for _, cl := range cls {
								if !cl.Limited {
									for i:= range room.GoMid {cl.InviteIntoGroupNormal(to, []string{room.GoMid[i]})
									}
									break
								}
							}
						}
					}
				}
			} else if cmd == "leaveall" {
				rngcmd := handler.GetComs(4, "leaveall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, p := range bk {
							gr, _ := p.GetGroupIdsJoined()
							for _, g := range gr {
								if g != msg.To {
									p.LeaveGroup(g)
									time.Sleep(1 * time.Second)
								}
							}
						}
						newsend += "Leave done"
						linetcr.RoomClear(room)
						handler.SaveBackup()
					}
				}
			} else if strings.HasPrefix(cmd, "bring") {
				rngcmd := handler.GetComs(5, "bring")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "bring ", "", 1)
						numb, _ := strconv.Atoi(str)
						if numb > 0 && numb <= len(botstate.ClientBot) {
							handler.GetSquad(client, to)
							all := []string{}
							room := linetcr.GetRoom(to)
							cuk := room.Client
							alls := []*linetcr.Account{}
							for _, x := range botstate.ClientBot {
								if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
									alls = append(alls, x)
								}
							}
							sort.Slice(all, func(i, j int) bool {
								return alls[i].KickCount < alls[j].KickCount
							})
							for _, x := range botstate.ClientBot {
								if len(all) < numb {
									if !linetcr.InArrayCl(cuk, x) && !linetcr.InArrayCl(linetcr.KickBans, x) && !linetcr.InArrayCl(room.GoClient, x) {
										all = append(all, x.MID)
									}
								} else {
									break
								}
							}
							cl := linetcr.GetRoom(to).Choose(client)
							if cl.Limited {
								cl.InviteIntoGroupNormal(to, all)
								time.Sleep(1 * time.Second)
								handler.GetSquad(client, to)
							} else {
								client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							}
						} else {
							newsend += "out of range.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "stay ") {
				rngcmd := handler.GetComs(7, "stay")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "stay ", "", 1)
						numb, _ := strconv.Atoi(str)
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						if numb > 0 && numb <= len(botstate.ClientBot) {
							handler.GetSquad(client, to)
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
								handler.GetSquad(client, to)
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
									wi := handler.GetSquad(client, to)
									for i := 0; i < len(all); i++ {
										if i == g {
											break
										}
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
									handler.GetSquad(client, to)
									handler.LogAccess(client, to, sender, "bringbot", []string{}, 2)
									handler.SaveBackup()
									aa := len(room.Client)
									var name string
									name = fmt.Sprintf("Ready %v bots here", aa)
									newsend += name + "\n"
								}
							}
						} else {
							newsend += "out of range.\n"
						}
					}
				}
			} else if cmd == "lastset" {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						a := "Lastset botstate.Command: "
						a += "\n\n    lkick"
						a += "\n    lcancel"
						a += "\n    Lqr"
						a += "\n    linvite"
						a += "\n    ljoin"
						a += "\n    lleave"
						a += "\n    lcon"
						a += "\n    ltag"
						a += "\n    lmid"
						a += "\n    lmessage"
						a += "\n    lbanlist"
						a += "\n    @me"
						a += "\n    @all"
						a += "\n    @oa"
						a += "\n    pend"
						a += "\n    pendingall"
						a += "\n    numpend"
						newsend += a + "\n"
					}
				}
			} else if cmd == "stickercmd" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						a := "✠ Sticker botstate.Command:"
						a += "\n\n  stickerkick"
						a += "\n  stickerinvite"
						a += "\n  stickerkickall"
						a += "\n  stickercancel"
						a += "\n  stickerbypass"
						a += "\n  stickerstayall"
						a += "\n  stickerleave"
						a += "\n  stickerrespon"
						a += "\n  stickerclearban"
						a += "\n  stickerclear"
						newsend += a + "\n"
					}
				}
			} else if cmd == "stickerclear" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.SaveBackup()
						a := "Clear sticker command"
						newsend += a + "\n"
					}
				}
			} else if cmd == "stickerkick" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerrespon" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerstayall" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerleave" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerinvite" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerkickall" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerbypass" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickerclearban" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if cmd == "stickercancel" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Please send sticker.\n"
					}
				}
			} else if pesan == "sname" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.SendMessage(to, botstate.Fancy(Sname))
					}
				}
			} else if pesan == "prefix" {
				if handler.CheckPermission(7, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.SendMessage(to, botstate.Fancy("Rname: "+Rname+"\nSname: "+Sname))
					}
				}
			} else if pesan == "rname" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.SendMessage(to, botstate.Fancy(Rname))
					}
				}
			} else if pesan == Sname {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.SendMessage(to, botstate.Fancy(botstate.MsgRespon))
					}
				}
			} else if pesan == Rname {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.SendMessage(to, botstate.Fancy(botstate.MsgRespon))
					}
				}
			} else if cmd == "gadmins" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(room.Gadmin) != 0 {
							list := "𝗴𝗮𝗱𝗺𝗶𝗻𝘀:\n"
							for num, xd := range room.Gadmin {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						} else {
							newsend += "Gadmin list is empty.\n"
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bot") && cmd != "botlist" {
				rngcmd := handler.GetComs(5, "bot")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 5
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "gban") && cmd != "gbanlist" {
				rngcmd := handler.GetComs(8, "gban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 10
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if cmd == "nukeban" {
				rngcmd := handler.GetComs(7, "nukeban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						go func() {
						      handler.KIckbansPurges(client, to)
						}()
					}
				}
			} else if cmd == "nukebot" {
				rngcmd := handler.GetComs(7, "nukebot")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						pelaku := op.Param2
						defer utils.PanicHandle("groupBackup")
						Room := linetcr.GetRoom(to)
						memlist, _ := client.GetChatListMap(to)
						ban := []string{}
						exe := []*linetcr.Account{}
						for mid, _ := range memlist {
							if utils.InArrayString(botstate.Squadlist, mid) {
								cl := handler.GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
							} else if handler.MemBan(to, mid) {
								ban = append(ban, mid)
							}
						}
						if len(exe) != 0 {
							sort.Slice(exe, func(i, j int) bool {
								return exe[i].KickPoint < exe[j].KickPoint
							})
							Room.HaveClient = exe
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
								if valid.Abs(ct) <= 10 {
									if handler.MemUser(to, mid) {
										botstate.Banned.AddBan(mid)
										ban = append(ban, mid)
									}
								}
							}
							no := 0
							ah := 0
							for _, target := range ban {
								go func(target string, no int) {
									exe[no].DeleteOtherFromChats(to, []string{target})
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
						}
					}
                           	client.SendMessage(to, botstate.Fancy("Not have bot enemy"))
				}
			} else if strings.HasPrefix(cmd, "spamcallto") {
				rngcmd := handler.GetComs(7, "spamcallto")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						targets := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(targets, lists[i]) {
									targets = append(targets, lists[i])
								}
							}
						}
                    		             if msg.ToType == 2 {
						     nu, _ := strconv.Atoi(botstate.CountSpam)
						     for i := 1; i <= nu; i++ {
								call.AcquireGroupCallRoute(to, client.AuthToken)
								call.InviteIntoGroupCall(to, client.AuthToken, targets)
							}
						}
                           		time.Sleep(3 * time.Second)
                           		client.SendMessage(to, botstate.Fancy("Successful spam invite call group"))
					}
				}
			} else if strings.HasPrefix(cmd, "spamcall") {
				rngcmd := handler.GetComs(7, "spamcall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "spamcall"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						result := strings.Split(str,"")
                    		             if msg.ToType == 2 {
						     nu, _ := strconv.Atoi(result[0])
							_, target, _ := client.GetChatList(to)
							targets := []string{}
							for i := range target {
								if !utils.InArrayString(botstate.CheckHaid, target[i]) {
									targets = append(targets, target[i])
								}
							 }
							for i := 1; i <= nu; i++ {
								call.AcquireGroupCallRoute(to, client.AuthToken)
								call.InviteIntoGroupCall(to, client.AuthToken, targets)
						       }
						}
                           		client.SendMessage(to, botstate.Fancy("Successful spam invite call group"))
					}
				}
			} else if cmd == "fixed" {
				rngcmd := handler.GetComs(7, "fixed")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.SaveBackup()
						newsend += "done.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "autobc") {
				rngcmd := handler.GetComs(7, "autobc")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "autobc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "msg" {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("Autobroadcast type message"))
						} else if str == "img" {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("Autobroadcast type image\n Please Send Image"))
						} else if str == "off" {
							count = count + 1
							client.SendMessage(to, botstate.Fancy("Autobroadcast is disabled"))
						}
						if count != 0 {
							//newsend += fmt.Sprintf("Autobroadcast type %s", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "msgbc") {
				rngcmd := handler.GetComs(7, "msgbc")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "msgbc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						handler.SaveBackup()
						newsend += "Message broadcast set to: " + str + "\n"
					}
				}
			} else if strings.HasPrefix(cmd, "groupbc") {
				rngcmd := handler.GetComs(7, "groupbc")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "groupbc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "image" {
							client.SendMessage(to, botstate.Fancy("Please send image"))
						} else if str == "video" {
							client.SendMessage(to, botstate.Fancy("Please send video"))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "friendbc") {
				rngcmd := handler.GetComs(7, "friendbc")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "friendbc"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "image" {
							client.SendMessage(to, botstate.Fancy("Please send image"))
						} else if str == "video" {
							client.SendMessage(to, botstate.Fancy("Please send video"))
						}
					}
				}
			} else if strings.HasPrefix(cmd, "startbcmid") {
				rngcmd := handler.GetComs(7, "startbcmid")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "startbcmid"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "image" {
							client.SendMessage(to, botstate.Fancy("Please send image"))
						} else if str == "video" {
							client.SendMessage(to, botstate.Fancy("Please send video"))
						}
					}
				}
			} else if cmd == "hentai" {
				rngcmd := handler.GetComs(7, "hentai")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.Sendmedias(to, "hentai","hentai")
					}
				}
			} else if cmd == "pornstart" {
				rngcmd := handler.GetComs(7, "pornstart")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.Sendmedias(to, "pornstart","pornstart")
					}
				}
			} else if strings.HasPrefix(cmd, "videoporn") {
				rngcmd := handler.GetComs(7, "videoporn")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "videoporn"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"pornvideo")
					}
				}
			} else if strings.HasPrefix(cmd, "tiktok") {
				rngcmd := handler.GetComs(7, "tiktok")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "tiktok"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"tiktok")
					}
				}
			} else if strings.HasPrefix(cmd, "smule") {
				rngcmd := handler.GetComs(7, "smule")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "smule"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"smule")
					}
				}
			} else if strings.HasPrefix(cmd, "joox") {
				rngcmd := handler.GetComs(7, "joox")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "joox"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"joox")
					}
				}
			} else if strings.HasPrefix(cmd, "youtube") {
				rngcmd := handler.GetComs(7, "youtube")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "youtube"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"youtube")
					}
				}
			} else if strings.HasPrefix(cmd, "instagram") {
				rngcmd := handler.GetComs(7, "instagram")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "instagram"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"instagram")
					}
				}
			} else if strings.HasPrefix(cmd, "textimage") {
				rngcmd := handler.GetComs(7, "textimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "textimage"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"image_text")
					}
				}
			} else if strings.HasPrefix(cmd, "calculator") {
				rngcmd := handler.GetComs(7, "calculator")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "calculator"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"calculator")
					}
				}
			} else if strings.HasPrefix(cmd, "cuaca") {
				rngcmd := handler.GetComs(7, "cuaca")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "cuaca"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"cuaca")
					}
				}
			} else if strings.HasPrefix(cmd, "simi") {
				rngcmd := handler.GetComs(7, "simi")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "simi"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"simi")
					}
				}
			} else if strings.HasPrefix(cmd, "artinama") {
				rngcmd := handler.GetComs(7, "artinama")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "artinama"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"artinama")
					}
				}
			} else if strings.HasPrefix(cmd, "gimage") {
				rngcmd := handler.GetComs(7, "gimage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "gimage"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"google")
					}
				}
			} else if strings.HasPrefix(cmd, "pinterest") {
				rngcmd := handler.GetComs(7, "pinterest")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "pinterest"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						data := str
						client.Sendmedias(to, data,"pinterest")
					}
				}
			} else if cmd == "#gtotal" {
				rngcmd := handler.GetComs(7, "gtotal")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						list := handler.Gtotal(client, to)
						client.SendMessage(to, botstate.Fancy(list))
					}
				}
			} else if cmd == "tes" {
				rngcmd := handler.GetComs(7, "tes")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						for _, tok := range botstate.Data.Authoken {
							handler.SendBigImage(tok, to, "https://thumbs.gfycat.com/SmartDenseBuck.webp")
					       }
					}
				}
			} else if cmd == "getcall" {
				rngcmd := handler.GetComs(7, "getcall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
                    		             if msg.ToType == 2 {
                        		            gcall, _ := call.GetGroupCall(to, client.AuthToken)
                                             Room := linetcr.GetRoom(to)
                        		            res := "Get Call Group:"
                        		            if gcall.MediaType == 1 {
                                                    res += "\n  • Type: Audio Call"
                        		                   res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                    loc, _ := time.LoadLocation("Asia/Jakarta")
	                                             a := time.Now().In(loc)
	                                             yyyy := strconv.Itoa(a.Year())
	                                             MM := a.Month().String()
	                                             dd := strconv.Itoa(a.Day())
	                                             Date := dd + "-" + MM + "-" + yyyy
                                                    cok := gcall.Started / 1000
			                                i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                                tm := time.Unix(i, 0)
			                                ss := time.Since(tm)
			                                sp := handler.FmtDurations(ss)
                                                    res += "\n  • Date: "+Date
                                                    res += "\n  • Started: "+sp
                        		                   res += "\n  • Members:"
                                                    mmk := gcall.MemberMids
						            if len(mmk) != 0 {
							            for num, xd := range mmk {
								            num++
								            rengs := strconv.Itoa(num)
								            x, _ := client.GetContact(xd)
								            res += "\n      " + rengs + ". "+x.DisplayName
							              }
					                     }
                                                       client.SendMessage(to, botstate.Fancy(res))
                                             }
                        		            if gcall.MediaType == 2 {
                                                    res += "\n  • Type: Video Call"
                        		                   res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                    loc, _ := time.LoadLocation("Asia/Jakarta")
	                                             a := time.Now().In(loc)
	                                             yyyy := strconv.Itoa(a.Year())
	                                             MM := a.Month().String()
	                                             dd := strconv.Itoa(a.Day())
	                                             Date := dd + "-" + MM + "-" + yyyy
                                                    cok := gcall.Started / 1000
			                                i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                                tm := time.Unix(i, 0)
			                                ss := time.Since(tm)
			                                sp := handler.FmtDurations(ss)
                                                    res += "\n  • Date: "+Date
                                                    res += "\n  • Started: "+sp
                        		                   res += "\n  • Members:"
                                                    mmk := gcall.MemberMids
						            if len(mmk) != 0 {
							            for num, xd := range mmk {
								            num++
								            rengs := strconv.Itoa(num)
								            x, _ := client.GetContact(xd)
								            res += "\n      " + rengs + ". "+x.DisplayName
							              }
					                     }
                                                       client.SendMessage(to, botstate.Fancy(res))
                                               }
					      }
				      }
				}
			} else if cmd == "bans" {
				rngcmd := handler.GetComs(7, "bans")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if handler.SendMybuyer(sender) {
							handler.CheckChatBan()
						}
						toy := len(linetcr.KickBans)
						banchat := len(linetcr.KickBanChat)
						addb := len(linetcr.Waitadd)
						kb := ""
						if toy != 0 {
							kb += fmt.Sprintf("%v/%v bot's kick/inv banned.", toy, len(botstate.Squadlist))
							for n, cl := range linetcr.KickBans {
								m := cl.MID
								no := n + 1
								var ta time.Duration
								if _, ok := linetcr.GetBlock.Get(cl.MID); ok {
									t := cl.TimeBan.Add(24 * time.Hour)
									ta = t.Sub(time.Now())
								} else {
									t := cl.TimeBan.Add(1*time.Hour )
									ta = t.Sub(time.Now())
								}
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\n%s\nRemaining %v", no, cl.Namebot, cl.MID, handler.FmtDurations(ta))
							}
						}
						fris := []*linetcr.Account{}
						for _, cl := range botstate.ClientBot {
							if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
								if cl.Frez {
									fris = append(fris, cl)
								}
							}
						}
						if len(fris) != 0 {
							no := 1
							mm := kb
							kb += fmt.Sprintf("\n\n%v/%v bot's freeze.", len(fris), len(botstate.Squadlist))
							for _, cl := range fris {
								t := cl.TimeBan.Add(1*time.Hour )
								ta := t.Sub(time.Now())
								if ta > 1*time.Second {
									kb += fmt.Sprintf("\n\n%v. %s\n%s\nRemaining %v", no, cl.Namebot, cl.MID, handler.FmtDurations(ta))
									no++
								} else {
									if _, ok := linetcr.GetBlock.Get(cl.MID); !ok {
										linetcr.KickBans = linetcr.RemoveCl(linetcr.KickBans, cl)
										cl.Limited = false
									}
									cl.Frez = false
								}
							}
							if no == 1 {
								kb = mm
							}
						}
						if addb != 0 {
							kb += fmt.Sprintf("\n\n%v/%v bot's add/friend banned.", addb, len(botstate.Squadlist))
							for n, cl := range linetcr.Waitadd {
								if !linetcr.InArrayCl(linetcr.KickBanChat, cl) {
									m := cl.MID
									no := n + 1
									var ta time.Duration
									if _, ok := linetcr.BlockAdd.Get(cl.MID); ok {
										t := cl.Timeadd.Add(24 * time.Hour)
										ta = t.Sub(time.Now())
									} else {
										t := cl.Timeadd.Add(1 * time.Hour)
										ta = t.Sub(time.Now())
									}
									if cl.Namebot == "" {
										pr, _ := client.GetContact(m)
										cl.Namebot = pr.DisplayName
									}
									kb += fmt.Sprintf("\n\n%v. %s\nRemaining %v", no, cl.Namebot, handler.FmtDurations(ta))
								}
							}
						}
						if banchat != 0 {
							kb += fmt.Sprintf("\n\n%v/%v bot's banchat.", banchat, len(botstate.Squadlist))
							for n, cl := range linetcr.KickBanChat {
								m := cl.MID
								if cl.Namebot == "" {
									pr, _ := client.GetContact(m)
									cl.Namebot = pr.DisplayName
								}
								kb += fmt.Sprintf("\n\n%v. %s\nmid: %s", n+1, cl.Namebot, m)
							}
						}
						if len(fris) == 0 && toy == 0 && addb == 0 && banchat == 0 {
							newsend += "All fixed."
						} else {
							newsend += kb
						}
					}

				}
			} else if cmd == "botlist" {
				rngcmd := handler.GetComs(5, "botlist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Bot) != 0 {
							list := "Botlist:\n"
							targets := []string{}
							for _, i := range botstate.UserBot.Bot {
								targets = append(targets, i)
							}
							client.SendPollMention(to, list, targets)
						} else {
							newsend += "Botlist is empty.\n"
						}
					}
				}
			} else if cmd == "clearbot" {
				rngcmd := handler.GetComs(5, "clearbot")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(botstate.UserBot.Bot) != 0 {
							newsend += fmt.Sprintf("Cleared %v botlist\n", len(botstate.UserBot.Bot))
							handler.LogAccess(client, to, sender, "clearbot", botstate.UserBot.Bot, msg.ToType)
							botstate.UserBot.ClearBot()
						} else {
							newsend += "Bot is empty.\n"
						}
					}
				}
			} else if cmd == "cleargban" {
				rngcmd := handler.GetComs(8, "cleargban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(room.Gban) != 0 {
							handler.LogAccess(client, to, sender, "cleargban", room.Gban, msg.ToType)
							newsend += fmt.Sprintf("Cleared %v gbanlist", len(room.Gban)) + "\n"
							room.Gban = []string{}
						} else {
							newsend += "Gban is empty.\n"
						}
					}
				}
			} else if cmd == "clears" || cmd == botstate.Commands.Clearchat && botstate.Commands.Clearchat != "" {
				rngcmd := handler.GetComs(5, "clears")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						for i := range botstate.ClientBot {
							if utils.InArrayString(memb, botstate.ClientBot[i].MID) {
								botstate.ClientBot[i].RemoveAllMessage(string(op.Param2))
							}
						}
						newsend += "Cleared all message.\n"
					}
				}
			} else if cmd == "clearall" {
				rngcmd := handler.GetComs(5, "clearall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						for i := range botstate.ClientBot {
							if utils.InArrayString(memb, botstate.ClientBot[i].MID) {
								botstate.ClientBot[i].RemoveAllMessage(string(op.Param2))
							}
						}
						exec.Command("bash", "-c", "sudo systemd-resolve --flush-caches").Output()
						exec.Command("bash", "-c", "echo 3 > /proc/sys/vm/drop_caches&&swapoff -a&&swapon -a").Output()
						newsend += "Cleared all message and cache.\n"
					}
				}
			} else if cmd == "clearcache" {
				rngcmd := handler.GetComs(5, "clearcache")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						exec.Command("bash", "-c", "sudo systemd-resolve --flush-caches").Output()
						exec.Command("bash", "-c", "echo 3 > /proc/sys/vm/drop_caches&&swapoff -a&&swapon -a").Output()
						//exec.Command("bash", "-c", "sudo apt update").Output()
						newsend += "Cleared all cache.\n"
					}
				}
			} else if cmd == "enablee2ee" {
				rngcmd := handler.GetComs(1, "enablee2ee")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.EnableE2ee()
						newsend += "Done Succes Enable E2EE.\n"
					}
				}
			} else if cmd == "disablee2ee" {
				rngcmd := handler.GetComs(1, "disablee2ee")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						client.DisableE2ee()
						newsend += "Done Succes Disable E2EE.\n"
					}
				}
			} else if cmd == "gbanlist" {
				rngcmd := handler.GetComs(8, "gbanlist")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(room.Gban) != 0 {
							list := "Gbanlist:"
							client.SendPollMention(to, list, room.Gban)
						} else {
							newsend += "Gban list is empty.\n"
						}
					}
				}
			} else if cmd == "infogo" {
				rngcmd := handler.GetComs(8, "infogo")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						o, _ := host.Info()
						r, _ := api.GetRunningVersion()
						Gplat := fmt.Sprintf("%v", r.Platform)
						//Gover := fmt.Sprintf("%v", r.Version)
						Garch := fmt.Sprintf("%v", r.Architecture)
						OS := fmt.Sprintf("Ubuntu %v ", o.PlatformVersion)
                   				a := "✠ INFORMATION:"
						a += "\n\n ⚙️ Platfrom : " + Gplat
						a += "\n ⚙️ OS : " + OS
						a += "\n ⚙️ Executed : Go 1.22.5"
						a += "\n ⚙️ Architecture : " + Garch
                   				a += "\n ⚙️ AppName : ANDROID"
                   				a += "\n ⚙️ UserAgent : Line/14.10.0"
                   				a += "\n ⚙️ Host : legy-jp-addr-long"
                   				a += "\n ⚙️ X-lal : "+client.Locale
						a += "\n ⚙️ Update : 10-07-2024"
						a += "\n ⚙️ Version : Sync5"
						a += "\n ⚙️ Condition : Good"
						a += "\n ⚙️ Team : SELFTCR™"
						a += handler.InfoCreator(client)
						newsend += a
					}
				}
			} else if cmd == "here" || cmd == botstate.Commands.Here && botstate.Commands.Here != "" {
				rngcmd := handler.GetComs(6, "here")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GetSquad(client, to)
						aa := len(room.Client)
						//bb := len(room.Client)
						cc := len(room.GoMid)
						var name string
						name = fmt.Sprintf("%v bots here - %v squad", aa, len(botstate.ClientBot))
						if cc != 0 {
							name += fmt.Sprintf("\n%v bots on stay.", cc)
						}
						toy := len(linetcr.KickBans)
						if toy != 0 {
							name += fmt.Sprintf("\n%v bots limited", toy)
						}
						newsend += name + "\n"
					}
				}
			} else if cmd == "ourl" {
				rngcmd := handler.GetComs(6, "ourl")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						tick, botstate.Err := client.ReissueChatTicket(to)
						if botstate.Err == nil {
							mes := make(chan bool)
							go func() {
								if botstate.Err != nil {
									mes <- false
								} else {
									mes <- true
								}
							}()
							newsend += "https://line.me/R/ti/g/" + tick + "\n"
						}
					}
				}
			} else if cmd == "curl" {
				rngcmd := handler.GetComs(6, "curl")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, "Sorry, all bot Here banned Try Invite Anther Bot")
							return
						}
						mes := make(chan bool)
						go func() {
							if botstate.Err != nil {
								mes <- true
							} else {
								mes <- false
							}
						}()
					}
				}
			} else if strings.HasPrefix(cmd, "say ") {
				rngcmd := handler.GetComs(6, "say")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						str := strings.Replace(cmd, "say ", "", 1)
						client.SendMessage(to, botstate.Fancy(str))
					}
				}
			} else if cmd == "timeleft" {
				rngcmd := handler.GetComs(6, "timeleft")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						d := fmt.Sprintf("%v", botstate.Data.Dalltime)
						has := strings.Split(d, "-")
						hass := strings.Split(has[2], "T")
						if len(has) == 3 {
							yy, _ := strconv.Atoi(has[0])
							mm, _ := strconv.Atoi(has[1])
							dd, _ := strconv.Atoi(hass[0])
							var time2 = time.Date(yy, time.Month(mm), dd, 00, 00, 0, 0, time.UTC)
							str := fmt.Sprintf("Date:\n %v-%v-%v", yy, mm, dd)
							ta := time2.Sub(time.Now())
							str += fmt.Sprintf("\nRemaining:\n  %v", handler.BotDuration(ta))
							newsend += str + "\n"
						}
					}
				}
			} else if cmd == "timenow" {
				rngcmd := handler.GetComs(7, "timenow")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GenerateTimeLog(client, to)
					}
				}
			} else if cmd == "runtime" {
				rngcmd := handler.GetComs(7, "runtime")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						elapsed := time.Since(botstate.BotStart)
						newsend += "Running Time:\n\n" + handler.BotDuration(elapsed) + "\n"
					}
				}
			} else if cmd == "setbot" {
				rngcmd := handler.GetComs(5, "setbot")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						ret := "✠ 𝗦𝗲𝘁 𝗠𝗼𝗱𝗲 𝗕𝗼𝘁𝘀:"
						ret += "\n"
						if botstate.AutoPro {ret += "\n ⚫ » ᴀᴜᴛᴏ ᴘʀᴏ"} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴘʀᴏ"}
						if botstate.AutoPurge {ret += "\n ⚫ » ᴀᴜᴛᴏ ᴘᴜʀɢᴇ"} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴘᴜʀɢᴇ"}
						if botstate.AutoJointicket {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ᴛɪᴄᴋᴇᴛ : %s", botstate.TypeJoin)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴛɪᴄᴋᴇᴛ"}
						if botstate.AutoTranslate {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ᴛʀᴀɴꜱʟᴀᴛᴇ : %s", botstate.TypeTrans)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴛʀᴀɴꜱʟᴀᴛᴇ"}
						if botstate.AutoLike {ret += "\n ⚫ » ᴀᴜᴛᴏ ʟɪᴋᴇ"} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ʟɪᴋᴇ"}
						if botstate.AutoBc {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ʙʀᴏᴀᴅᴄᴀꜱᴛ : %s", botstate.Typebc)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ʙʀᴏᴀᴅᴄᴀꜱᴛ"}
						if botstate.Autojoin != "off" {ret += fmt.Sprintf("\n ⚫ » ᴀᴜᴛᴏ ᴊᴏɪɴ :  %s", botstate.Autojoin)} else {ret += "\n ⚪ » ᴀᴜᴛᴏ ᴊᴏɪɴ"}
						if botstate.Ajsjoin != "off" {ret += fmt.Sprintf("\n ⚫ » ᴀᴊꜱ ᴊᴏɪɴ :  %s", botstate.Ajsjoin)} else {ret += "\n ⚪ » ᴀᴊꜱ ᴊᴏɪɴ"}
						if botstate.Canceljoin {ret += "\n ⚫ » ᴄᴀɴᴄᴇʟ ᴊᴏɪɴ"} else {ret += "\n ⚪ » ᴄᴀɴᴄᴇʟ ᴊᴏɪɴ"}
						if botstate.NukeJoin {ret += "\n ⚫ » ɴᴜᴋᴇ ᴊᴏɪɴ"} else {ret += "\n ⚪ » ɴᴜᴋᴇ ᴊᴏɪɴ"}
						if botstate.PowerMode {ret += "\n ⚫ » ᴍᴏᴅᴇ ᴡᴀʀ"} else {ret += "\n ⚪ » ᴍᴏᴅᴇ ᴡᴀʀ"}
						if botstate.ProtectMode {ret += "\n ⚫ » ᴍᴏᴅᴇ ᴘʀᴏᴛᴇᴄᴛ"} else {ret += "\n ⚪ » ᴍᴏᴅᴇ ᴘʀᴏᴛᴇᴄᴛ"}
						if botstate.AutoBackBot {ret += "\n ⚫ » ᴍᴏᴅᴇ ʙᴀᴄᴋ"} else {ret += "\n ⚪ » ᴍᴏᴅᴇ ʙᴀᴄᴋ"}
						if botstate.KickBanQr {ret += "\n ⚫ » ᴋɪᴄᴋ ʙᴀɴQʀ"} else {ret += "\n ⚪ » ᴋɪᴄᴋ ʙᴀɴQʀ"}
						if botstate.ModeBackup != "" {ret += fmt.Sprintf("\n ⚫ » ʙᴀᴄᴋᴜᴘ : %s", botstate.ModeBackup)} else {ret += "\n ⚪ » ʙᴀᴄᴋᴜᴘ"}
						if linetcr.FlexMode {ret += "\n ⚫ » ꜰʟᴇx ᴍᴏᴅᴇ1"} else if linetcr.FlexMode2 {ret += "\n ⚫ » ꜰʟᴇx ᴍᴏᴅᴇ2"} else if linetcr.FooterMode {ret += "\n ⚫ » ꜰᴏᴏᴛᴇʀ ᴍᴏᴅᴇ"} else {ret += "\n ⚪ » ꜰʟᴇx / ꜰᴏᴏᴛᴇʀ"}
						if botstate.MediaDl {ret += "\n ⚫ » ᴍᴇᴅɪᴀ ᴅᴏᴡɴʟᴏᴀᴅ"} else {ret += "\n ⚪ » ᴍᴇᴅɪᴀ ᴅᴏᴡɴʟᴏᴀᴅ"}
						//ret += "\n\n▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎▪︎"
						ret += "\n"
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ᴋɪᴄᴋ : %v", botstate.Data.Logobot, botstate.MaxKick)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ɪɴᴠɪᴛᴇ : %v", botstate.Data.Logobot, botstate.MaxInvite)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ᴄᴀɴꜱ : %v", botstate.Data.Logobot, botstate.MaxCancel)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ᴘᴇɴᴅ : %v", botstate.Data.Logobot, botstate.CancelPend)
						ret += fmt.Sprintf("\n %v ʟɪᴍɪᴛᴇʀ ꜱᴘᴀᴍ : %v", botstate.Data.Logobot, botstate.CountSpam)
						if botstate.LockAjs {ret += fmt.Sprintf("\n %v ʟᴏᴄᴋ ᴀɴᴛɪᴊꜱ : %v (true)", botstate.Data.Logobot, botstate.CountAjs)} else {ret += fmt.Sprintf("\n %v ʟᴏᴄᴋ ᴀɴᴛɪᴊꜱ : %v (false)", botstate.Data.Logobot, botstate.CountAjs)}
						ret += fmt.Sprintf("\n %v ᴍᴏᴅᴇᴛᴇxᴛ : %v", botstate.Data.Logobot, botstate.Fancy)
						//ret += fmt.Sprintf("\n %v ᴍᴏᴅᴇᴛᴇxᴛ : "+botstate.Fancy, botstate.Data.Logobot)
						rng1 := handler.GetComs(7, "invitebot")
						rng12 := handler.GetComs(4, "remote")
						xx := handler.GetGrade(rng1)
						yy := handler.GetGrade(rng12)
						ret += fmt.Sprintf("\n %v ᴘᴇʀᴍ ɪɴᴠɪᴛᴇʙᴏᴛ : %v ", botstate.Data.Logobot, xx)
						ret += fmt.Sprintf("\n %v ᴘᴇʀᴍ ʀᴇᴍᴏᴛᴇ : %v ", botstate.Data.Logobot, yy)
						ret += fmt.Sprintf("\n %v ᴄʀᴇᴀᴛɪᴏɴ ᴛᴇᴀᴍ : ꜱᴇʟꜰᴛᴄʀ™", botstate.Data.Logobot)
						client.SendHelp(to, botstate.Fancy(ret))
					}
				}
			} else if cmd == "setgroup" {
				rngcmd := handler.GetComs(9, "set")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GetSquad(client, to)
						aa := len(room.Client)
						cc := len(room.GoMid)
						ret := "✠ 𝗣𝗿𝗼𝘁𝗲𝗰𝘁 𝗦𝗲𝘁𝘁𝗶𝗻𝗴:\n"
						if op.Message.ToType == 2 {
							if room.ProQr {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ Qʀ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ Qʀ"
							}
							if room.ProKick {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴋɪᴄᴋ"
							}
							if room.ProInvite {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪɴᴠɪᴛᴇ"
							}
							if room.ProCancel {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀɴᴄᴇʟ"
							}
							if room.ProJoin {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴊᴏɪɴ"
							}
							if room.ProName {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴀᴍᴇ"
							}
							if room.AntiTag {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴛᴀɢ"
							}
							if room.ProPicture {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘɪᴄᴛᴜʀᴇ"
							}
							if room.ProNote {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɴᴏᴛᴇ"
							}
							if room.ProAlbum {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀʟʙᴜᴍ"
							}
							if room.ProLink {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ʟɪɴᴋ"
							}
							if room.ProFlex {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰʟᴇx"
							}
							if room.ProImage {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ɪᴍᴀɢᴇ"
							}
							if room.ProVideo {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴠɪᴅᴇᴏ"
							}
							if room.ProCall {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴀʟʟ"
							}
							if room.ProSpam {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴘᴀᴍ"
							}
							if room.ProSticker {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜱᴛɪᴄᴋᴇʀ"
							}
							if room.ProContact {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴄᴏɴᴛᴀᴄᴛ"
							}
							if room.ProPost {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴘᴏꜱᴛ"
							}
							if room.ProFile {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ꜰɪʟᴇ"
							}
							if len(room.GoMid) > 0 {
								ret += "\n ⚫ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							} else {
								ret += "\n ⚪ » ᴘʀᴏᴛᴇᴄᴛ ᴀᴊꜱ"
							}
							ret += "\n"
							ret += "\n✠ 𝗚𝗿𝗼𝘂𝗽 𝗦𝗲𝘁𝘁𝗶𝗻𝗴:\n"
							if room.Lurk {
								ret += fmt.Sprintf("\n ⚫ » ʟᴜʀᴋɪɴɢ %s", room.NameLurk)
							} else {
								ret += "\n ⚪ » ʟᴜʀᴋɪɴɢ"
							}
							if room.Automute {
								ret += "\n ⚫ » ᴀᴜᴛᴏᴍᴜᴛᴇ"
							} else {
								ret += "\n ⚪ » ᴀᴜᴛᴏᴍᴜᴛᴇ"
							}
							if room.Welcome {
								ret += "\n ⚫ » ᴡᴇʟᴄᴏᴍᴇ"
							} else {
								ret += "\n ⚪ » ᴡᴇʟᴄᴏᴍᴇ"
							}
							if room.Leavebool {
								ret += "\n ⚫ » ʟᴇᴀᴠᴇ"
							} else {
								ret += "\n ⚪ » ʟᴇᴀᴠᴇ"
							}
							if room.Announce {
								ret += "\n ⚫ » ᴀɴɴᴏᴜɴᴄᴇ"
							} else {
								ret += "\n ⚪ » ᴀɴɴᴏᴜɴᴄᴇ"
							}
							if room.Backleave {
								ret += "\n ⚫ » ʜᴏꜱᴛᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ʜᴏꜱᴛᴀɢᴇ"
							}
							if botstate.LogGroup == to {
								ret += "\n ⚫ » ɴᴏᴛɪꜰɪᴄᴀᴛɪᴏɴ"
							} else {
								ret += "\n ⚪ » ɴᴏᴛɪꜰɪᴄᴀᴛɪᴏɴ"
							}
							if room.ImageLurk {
								ret += "\n ⚫ » ʟᴜʀᴋ ɪᴍᴀɢᴇ"
							} else {
								ret += "\n ⚪ » ʟᴜʀᴋ ɪᴍᴀɢᴇ"
							}
							if room.Backup {
								ret += "\n ⚫ » ʙᴀᴄᴋᴜᴘ ᴜꜱᴇʀ"
							} else {
								ret += "\n ⚪ » ʙᴀᴄᴋᴜᴘ ᴜꜱᴇʀ"
							}
							if botstate.DetectCall {
								ret += "\n ⚫ » ᴅᴇᴛᴇᴄᴛ ɢʀᴏᴜᴘᴄᴀʟʟ"
							} else {
								ret += "\n ⚪ » ᴅᴇᴛᴇᴄᴛ ɢʀᴏᴜᴘᴄᴀʟʟ"
							}
						}
						ret += fmt.Sprintf("\n\n %v/%v ʙᴏᴛꜱ ʜᴇʀᴇ.", aa, len(botstate.ClientBot))
						if cc != 0 {
							ret += fmt.Sprintf("\n %v ʙᴏᴛꜱ ᴏɴ ꜱᴛᴀʏ.", cc)
						}
						client.SendHelp(to, botstate.Fancy(ret))
					}
				}
			} else if cmd == "lurk image" {
				//rngcmd := handler.GetComs(9, "image")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						room.Lurk = true
						room.ImageLurk = true
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk name" {
				//rngcmd := handler.GetComs(8, "name")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk mention" {
				//rngcmd := handler.GetComs(8, "mention")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "mention"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "lurk on" {
				//rngcmd := handler.GetComs(8, "on")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						room.Lurk = true
						room.NameLurk = "name"
						room.Userlurk = []string{}
						newsend += "Lurking enabled.\n"
					}
				}
			} else if cmd == "mediadl on" {
				rngcmd := handler.GetComs(9, "on")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Mediad download enabled.\n"
					}
				}
			} else if cmd == "mediadl off" {
				rngcmd := handler.GetComs(9, "off")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Media download enabled.\n"
					}
				}
			} else if cmd == "autolike on" {
				rngcmd := handler.GetComs(9, "on")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Autolike enabled.\n"
					}
				}
			} else if cmd == "autolike off" {
				rngcmd := handler.GetComs(9, "off")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Autolike enabled.\n"
					}
				}
			} else if cmd == "bomlike 10" {
				rngcmd := handler.GetComs(9, "10")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Share you post\n"
					}
				}
			} else if cmd == "bomlike off" {
				rngcmd := handler.GetComs(9, "off")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						newsend += "Bomlike enabled.\n"
					}
				}
			} else if strings.HasPrefix(cmd, "countspam") {
				rngcmd := handler.GetComs(5, "countspam")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "countspam"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
                        			client.SendMessage(to, botstate.Fancy("Success update countspam: "+botstate.CountSpam+""))
					}
				}
			} else if strings.HasPrefix(cmd, "countajs") {
				rngcmd := handler.GetComs(5, "countajs")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "countajs"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
                        			client.SendMessage(to, botstate.Fancy("Success update countajs: "+botstate.CountAjs+""))
					}
				}
			} else if strings.HasPrefix(cmd, "killmode") {
				rngcmd := handler.GetComs(5, "killmode")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "killmode"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "kill" {
							count = count + 1
						} else if str == "purge" {
							count = count + 1
						} else if str == "range" {
							count = count + 1
						} else if str == "random" {
							count = count + 1
						} else if str == "off" {
							count = count + 1
						}
						if count != 0 {
							newsend += fmt.Sprintf("botstate.Killmode state : %s\nTurn on", str)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "textmode") {
				rngcmd := handler.GetComs(5, "textmode")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						var su = "textmode"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "1" {
						} else if str == "2" {
						} else if str == "3" {
						} else if str == "4" {
						} else if str == "5" {
						} else if str == "6" {
						} else if str == "7" {
						} else if str == "8" {
						} else if str == "9" {
						} else if str == "10" {
						} else if str == "normal" {
						}
						newsend += "Update mode text: "+str
					}
				}
			} else if strings.HasPrefix(cmd, "backup") {
				rngcmd := handler.GetComs(5, "backup")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						var str string
						count := 0
						var su = "backup"
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if str == "inv" {
							count = count + 1
						} else if str == "qr" {
							count = count + 1
						}
						if count != 0 {
							newsend += fmt.Sprintf("Mode Backup state : %s\nTurn on", str)
						}
					}
				}
			} else if cmd == "lurk" {
				rngcmd := handler.GetComs(6, "lurk")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						room.Lurk = true
						room.Userlurk = []string{}
						room.NameLurk = "hide"
						newsend += "Lurking...\n"
					}
				}
			} else if cmd == "lurks" {
				rngcmd := handler.GetComs(6, "lurk")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(room.Userlurk) != 0 {
							list := "✠ Lurkers:\n"
							for num, xd := range room.Userlurk {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"

						} else {
							newsend += "Lurk list empty enable first.\n"
						}
					}
				}
			} else if cmd == "lurk off" {
				rngcmd := handler.GetComs(9, "off")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						room.Lurk = false
						room.ImageLurk = false
						if len(room.Userlurk) != 0 {
							list := " ✠ Lurkers ✠ \n"
							for num, xd := range room.Userlurk {
								num++
								rengs := strconv.Itoa(num)
								name := handler.GetContactName(client, xd)
									list += "\n   " + rengs + ". " + name
							}
							newsend += list + "\n"
						}
						room.Userlurk = []string{}
					}
				}
			} else if cmd == "/status all" {
				rngcmd := handler.GetComs(6, "/statusall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						ret := "✠ Status Allbot:"
						ret += "\n"
						for i := range botstate.ClientBot {
							botstate.ClientBot[i].DeleteOtherFromChats(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							if botstate.ClientBot[i].Limited == true {
								ret += fmt.Sprintf("\n%v. "+botstate.ClientBot[i].Namebot+": %s", i+1, botstate.Data.Limit)
							} else {
								ret += fmt.Sprintf("\n%v. "+botstate.ClientBot[i].Namebot+": %s", i+1, botstate.Data.Fresh)
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "status add" {
				rngcmd := handler.GetComs(6, "status add")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						ret := "✠ Status Add:"
						ret += "\n"
						for i := range botstate.ClientBot {
							ve := "uc52554b082eca0360da013d33df023e0"
							botstate.Err, _ := botstate.ClientBot[i].FindAndAddContactsByMidV2(ve)
							fff := fmt.Sprintf("%v", botstate.Err)
							er := strings.Contains(fff, "request blocked")
							if er == true {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, botstate.Data.Limit)
							} else {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, botstate.Data.Fresh)
							}
						}
						ret += "\n"
						client.SendHelp(to, botstate.Fancy(ret))
					}
				}
			} else if cmd == "/status" || cmd == botstate.Commands.Status && botstate.Commands.Status != "" {
				rngcmd := handler.GetComs(6, "/status")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, memb, _ := client.GetChatList(to)
						var a = 0
						ret := "✠ Status Bot:"
						ret += "\n"
						for i := range botstate.ClientBot {
							if utils.InArrayString(memb, botstate.ClientBot[i].MID) {
								a = a + 1
								botstate.ClientBot[i].DeleteOtherFromChats(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
								if botstate.ClientBot[i].Limited == true {
									ret += fmt.Sprintf("\n%v. "+botstate.ClientBot[i].Namebot+": %s", a, botstate.Data.Limit)
								} else {
									ret += fmt.Sprintf("\n%v. "+botstate.ClientBot[i].Namebot+": %s", a, botstate.Data.Fresh)
								}
							}
						}
						ret += "\n"
						newsend += ret
					}
				}
			} else if cmd == "sp" {
				rngcmd := handler.GetComs(8, "sp")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GetSquad(client, to)
						for _, p := range bk {
							start := time.Now()
							p.GetContact(p.MID)
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret := fmt.Sprintf("%v ms", sp)
							p.SendMessage(to, botstate.Fancy(ret))
						}
					}
				}
			} else if cmd == "speed" || cmd == "speed" || cmd == botstate.Commands.Speed && botstate.Commands.Speed != "" {
				rngcmd := handler.GetComs(8, "speed")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GetSquad(client, to)
						var a = 0
						ret := "✠ Speed Bot:"
						ret += "\n"
						for _, p := range bk {
							a = a + 1
							start := time.Now()
							p.GetContact(p.MID)
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", a, sp)
						}
						ret += "\n"
						client.SendHelp(to, botstate.Fancy(ret))
					}
				}
			} else if cmd == "speed all" {
				rngcmd := handler.GetComs(8, "speed all")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GetSquad(client, to)
						var a = 0
						ret := "✠ Speed Profile:"
						ret += "\n"
						for _, p := range bk {
							a = a + 1
							start := time.Now()
							p.GetProfile()
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", a, sp)
						}
						var b = 0
						ret += "\n\n✠ Speed Contact:"
						ret += "\n"
						for _, p := range bk {
							b = b + 1
							start := time.Now()
							p.GetContact(p.MID)
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", b, sp)
						}
						var c = 0
						ret += "\n\n✠ Speed Message:"
						ret += "\n"
						for _, p := range bk {
							c = c + 1
							start := time.Now()
							p.SendMessage("u27623a2c021c18746b7aa34e3d2b2220", "sp")
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", c, sp)
						}
						var d = 0
						ret += "\n\n✠ Speed Kick:"
						ret += "\n"
						for _, p := range bk {
							d = d + 1
							start := time.Now()
							p.DeleteOtherFromChats(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", d, sp)
						}
						var e = 0
						ret += "\n\n✠ Speed Invite:"
						ret += "\n"
						for _, p := range bk {
							e = e + 1
							start := time.Now()
							p.InviteIntoGroupNormal(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", e, sp)
						}
						var f = 0
						ret += "\n\n✠ Speed Cancel:"
						ret += "\n"
						for _, p := range bk {
							f = f + 1
							start := time.Now()
							p.CancelChatInvitations(to, []string{"u27623a2c021c18746b7aa34e3d2b2220"})
							elapsed := time.Since(start)
							sp := fmt.Sprintf("%v", elapsed)
							sp = sp[:3]
							ret += fmt.Sprintf("\n Bots%v: %v ms", f, sp)
						}
						ret += "\n"
						client.SendHelp(to, botstate.Fancy(ret))
					}
				}
			} else if cmd == "status" || cmd == botstate.Commands.Status && botstate.Commands.Status != "" {
				rngcmd := handler.GetComs(8, "tus")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GetSquad(client, to)
						var a = 0
						ret := "✠ Status Bot:"
						ret += "\n"
						for _, p := range bk {
							a = a + 1
							if p.Limited == true {
								ret += fmt.Sprintf("\n Bots%v: %s", a, botstate.MsLimit)
							} else {
								ret += fmt.Sprintf("\n Bots%v: %s", a, botstate.MsFresh)
							}
						}
						ret += "\n"
						client.SendHelp(to, botstate.Fancy(ret))
					}
				}
			} else if cmd == "status all" {
				rngcmd := handler.GetComs(5, "statusall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						ret := "✠ Status Allbot:"
						ret += "\n"
						for i := range botstate.ClientBot {
							if botstate.ClientBot[i].Limited == true {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, botstate.MsLimit)
							} else {
								ret += fmt.Sprintf("\n Bots%v: %s", i+1, botstate.MsFresh)
							}
						}
						ret += "\n"
						client.SendHelp(to, botstate.Fancy(ret))
					}
				}
			} else if strings.HasPrefix(cmd, "help ") && cmd != "help" {
				if !handler.MemUser(to, sender) {
					if handler.handler.CheckExprd(client, to, sender) {
						txt := strings.ReplaceAll(cmd, "help ", "")
						texts := strings.Split(txt, " ")
						if len(texts) != 0 {
							kata := texts[0]
							if kata == "all" {
								res := "✠ 𝗣𝗿𝗼𝘁𝗲𝗰𝘁𝗶𝗼𝗻 𝗠𝗲𝗻𝘂:"
								res += "\n"
								for _, x := range botstate.Helppro {
									res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
								}
								if utils.InArrayString(botstate.DEVELOPER,sender) {
									if handler.CheckPermission(0, sender, to) {
										res += "\n"
										res += "\n✠ 𝗗𝗲𝘃𝗲𝗹𝗼𝗽𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpdeveloper {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
									}
								}
								if handler.SendMycreator(sender) {
									if handler.CheckPermission(1, sender, to) {
										res += "\n"
										res += "\n✠ 𝗖𝗿𝗲𝗮𝘁𝗼𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpcreator {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
									}
								}
								if handler.SendMymaker(sender) {
									if handler.CheckPermission(2, sender, to) {
										res += "\n"
										res += "\n✠ 𝗠𝗮𝗸𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpmaker {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
									}
								}
								if handler.SendMyseller(sender) {
									if handler.CheckPermission(3, sender, to) {
										res += "\n"
										res += "\n✠ 𝗦𝗲𝗹𝗹𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpseller {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
									}
								}
								if handler.SendMybuyer(sender) {
									if handler.CheckPermission(4, sender, to) {
										res += "\n"
										res += "\n✠ 𝗕𝘂𝘆𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpbuyer {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
									}
								}
								if handler.SendMyowner(sender) {
									if handler.CheckPermission(5, sender, to) {
										res += "\n"
										res += "\n✠ 𝗢𝘄𝗻𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpowner {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
									}
								}
								if handler.SendMymaster(sender) {
									if handler.CheckPermission(6, sender, to) {
										res += "\n"
										res += "\n✠ 𝗠𝗮𝘀𝘁𝗲𝗿 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpmaster {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
									}
								}
								if handler.SendMyadmin(sender) {
									if handler.CheckPermission(7, sender, to) {
										res += "\n"
										res += "\n✠ 𝗔𝗱𝗺𝗶𝗻 𝗠𝗲𝗻𝘂:"
										res += "\n"
										for _, x := range botstate.Helpadmin {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, res + "\n")//botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "dev" {
								if handler.CheckPermission(0, sender, to) {
									if utils.InArrayString(botstate.DEVELOPER,sender) {
										res := "✠ 𝗗𝗲𝘃𝗲𝗹𝗼𝗽𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpdeveloper {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "creator" {
								if handler.CheckPermission(1, sender, to) {
									if handler.SendMycreator(sender) {
										res := "✠ 𝗖𝗿𝗲𝗮𝘁𝗼𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpcreator {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "maker" {
								if handler.CheckPermission(1, sender, to) {
									if handler.SendMymaker(sender) {
										res := "✠ 𝗠𝗮𝗸𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpmaker {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "seller" {
								if handler.CheckPermission(2, sender, to) {
									if handler.SendMyseller(sender) {
										res := "✠ 𝗦𝗲𝗹𝗹𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpseller {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "buyer" {
								if handler.CheckPermission(3, sender, to) {
									if handler.SendMybuyer(sender) {
										res := "✠ 𝗕𝘂𝘆𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpbuyer {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "owner" {
								if handler.CheckPermission(4, sender, to) {
									if handler.SendMyowner(sender) {
										res := "✠ 𝗢𝘄𝗻𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpowner {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "master" {
								if handler.CheckPermission(5, sender, to) {
									if handler.SendMymaster(sender) {
										res := "✠ 𝗠𝗮𝘀𝘁𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpmaster {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "admin" {
								if handler.CheckPermission(6, sender, to) {
									if handler.SendMyadmin(sender) {
										res := "✠ 𝗔𝗱𝗺𝗶𝗻 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpadmin {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "gowner" {
								if handler.CheckPermission(8, sender, to) {
									if handler.SendMygowner(to, sender) {
										res := "✠ 𝗚𝗼𝘄𝗻𝗲𝗿 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpgowner {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "gadmin" {
								if handler.CheckPermission(9, sender, to) {
									if handler.SendMygadmin(to, sender) {
										res := "✠ 𝗚𝗮𝗱𝗺𝗶𝗻 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
										res += "\n"
										for _, x := range botstate.Helpgadmin {
											res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
										}
										client.SendHelp(to, botstate.Fancy(res + "\n"))
									}
								}
							} else if kata == "protect" {
								res := "✠ 𝗣𝗿𝗼𝘁𝗲𝗰𝘁 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
								res += "\n"
								for _, x := range botstate.Helppro {
									res += fmt.Sprintf("\n  %v %s", botstate.Data.Logobot, x)
								}
								client.SendHelp(to, res)
							} else {
								k := handler.GetKey(kata)
								det, anu := botstate.Details[k]
								tt := fmt.Sprintf(det, botstate.Used, k)
								if anu {
									newsend += tt
								} else {
									newsend += "Not found any command's that's have."
								}
							}
						}
					}
				}
			} else if cmd == "help" {
				if handler.CheckPermission(9, sender, to) {
					res := "✠ 𝗠𝗲𝗻𝘂 𝗠𝗲𝘀𝘀𝗮𝗴𝗲:"
					res += fmt.Sprintf("\n %v ᴀᴄᴄᴇꜱꜱ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ɢᴀᴄᴄᴇꜱꜱ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ (ᴛʏᴘᴇ ᴘᴜʙʟɪᴄ)",botstate.Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗔𝗹𝗹:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴀʟʟ",botstate.Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗣𝗿𝗼𝘁𝗲𝗰𝘁:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴘʀᴏᴛᴇᴄᴛ",botstate.Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗔𝗰𝗰𝗲𝘀𝘀:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴅᴇᴠ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴄʀᴇᴀᴛᴏʀ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴍᴀᴋᴇʀ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ꜱᴇʟʟᴇʀ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ʙᴜʏᴇʀ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴏᴡɴᴇʀ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴍᴀꜱᴛᴇʀ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ᴀᴅᴍɪɴ",botstate.Data.Logobot)
					res += "\n\n✠ 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀 𝗚𝗔𝗰𝗰𝗲𝘀𝘀:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ɢᴏᴡɴᴇʀ",botstate.Data.Logobot)
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ ɢᴀᴅᴍɪɴ",botstate.Data.Logobot)
					res += "\n\n✠ 𝗘𝘅𝗮𝗺𝗽𝗹𝗲 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
					res += fmt.Sprintf("\n %v ʜᴇʟᴘ (ᴄᴏᴍᴍᴀɴᴅ)",botstate.Data.Logobot)
					res += "\n\n✠ 𝗣𝗿𝗲𝗳𝗶𝘅:"
					res += fmt.Sprintf("\n %v ꜱɴᴀᴍᴇ : %v",botstate.Data.Logobot, Sname)
					res += fmt.Sprintf("\n %v ʀɴᴀᴍᴇ : %v",botstate.Data.Logobot, Rname)
					client.SendHelp(to, botstate.Fancy(res))
				}
			} else if cmd == "about" {
				rngcmd := handler.GetComs(5, "about")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						handler.GetSquad(client, to)
						var a = 0
						ret := "✠ Set Account:"
						for _, p := range bk {
							a = a + 1
							cokk, _ := p.GetSettings()
							ret += fmt.Sprintf("\n\nBot%v:\n", a)
							if cokk.PrivacyReceiveMessagesFromNotFriend == true {
								ret += "   ✓   Filter\n"
							} else {
								ret += "   ✘   Filter\n"
							}
							if cokk.EmailConfirmationStatus == 3 {
								ret += "   ✓   Email\n"
							} else {
								ret += "   ✘   Email\n"
							}
							if cokk.E2eeEnable == true {
								ret += "   ✓   Lsealing\n"
							} else {
								ret += "   ✘   Lsealing\n"
							}
							if cokk.PrivacyAllowSecondaryDeviceLogin == true {
								ret += "   ✓   Secondary\n"
							} else {
								ret += "   ✘   Secondary\n"
							}
						}
						client.SendMessage(to, botstate.Fancy(ret+"\n\nDevelolper: \nhttps://line.me/ti/p/~code-bot"))
					}
				}
			} else if cmd == "tagall" || cmd == botstate.Commands.Tagall && botstate.Commands.Tagall != "" {
				//rngcmd := handler.GetComs(8, "all")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, target, _ := client.GetChatList(to)
						targets := []string{}
						for i := range target {
							if !utils.InArrayString(botstate.CheckHaid, target[i]) {
								targets = append(targets, target[i])
							}
						}
						client.SendPollMention(to, "Mentions member:\n", targets)
					}
				}
			} else if cmd == "ftagall" {
				rngcmd := handler.GetComs(8, "tagall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						_, target, _ := client.GetChatList(to)
						targets := []string{}
						for i := range target {
							if !utils.InArrayString(botstate.CheckHaid, target[i]) {
								targets = append(targets, target[i])
							}
						}
						client.FakeMention(to, targets)
					}
				}
			} else if strings.HasPrefix(cmd, "unbot") {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 5
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.Checklistexpel(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unbot"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.UserBot.Bot)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.Checklistexpel(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "ungban") {
				rngcmd := handler.GetComs(8, "ungban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 3
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.handler.CheckUnbanBots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "ungban"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, room.Gban)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.handler.CheckUnbanBots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "cancel") && cmd != "cancelall" {
				rngcmd := handler.GetComs(7, "cancel")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := handler.SelectBot(client, to)
							if exe != nil {
								handler.Setcancelto(exe, to, listuser)
								handler.LogAccess(client, to, sender, "cancel", listuser, msg.ToType)
							} else {
								client.SendMessage(to, botstate.Fancy("Please add another bot that has a ban cancel."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "invite") {
				rngcmd := handler.GetComs(7, "invite")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := handler.SelectBot(client, to)
							if exe != nil {
								lists := handler.Setinvitetomsg(exe, to, listuser)
								if len(lists) != 0 {
									handler.Cekbanwhois(client, to, lists)
								}
								handler.LogAccess(client, to, sender, "invite", listuser, msg.ToType)
							} else {
								client.SendMessage(to, botstate.Fancy("Please add another bot that has a ban invite."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "addfriends") {
				rngcmd := handler.GetComs(2, "addfriends")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						x := 18
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := handler.SelectBot(client, to)
							if exe != nil {
								handler.AddCon(listuser)
								handler.CheckListAccess(client, to, listuser, x, sender)
							} else {
								client.SendMessage(to, botstate.Fancy("Please add another bot that has a ban addfriends."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "vkick") && cmd != "kickall" || strings.HasPrefix(cmd, botstate.Commands.Kick) && botstate.Commands.Kick != "" && cmd != "kickall" {
				rngcmd := handler.GetComs(7, "vkick")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if handler.MemUser(to, lists[i]) && !utils.InArrayString(listuser, lists[i]) {
									if botstate.AutoBan {
										botstate.Banned.AddBan(lists[i])
									}
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := handler.SelectBot(client, to)
							if exe != nil {
								handler.Setkickto(exe, to, listuser)
								handler.Setinvitetomsg(exe, to, listuser)
								handler.Setcancelto(exe, to, listuser)
								//botstate.AutoproN = true
								handler.LogAccess(client, to, sender, "vkick", listuser, msg.ToType)
							} else {
								client.SendMessage(to, botstate.Fancy("Please add another bot that has a ban kick."))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "kick") && cmd != "kickall" || strings.HasPrefix(cmd, botstate.Commands.Kick) && botstate.Commands.Kick != "" && cmd != "kickall" {
				rngcmd := handler.GetComs(7, "kick")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						listuser := []string{}
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if handler.MemUser(to, lists[i]) && !utils.InArrayString(listuser, lists[i]) {
									if botstate.AutoBan {
										botstate.Banned.AddBan(lists[i])
									}
									listuser = append(listuser, lists[i])
								}
							}
						}
						if len(listuser) != 0 {
							exe, _ := handler.SelectBot(client, to)
							if exe != nil {
								handler.Setkickto(exe, to, listuser)
								//botstate.AutoproN = true
								handler.LogAccess(client, to, sender, "kick", listuser, msg.ToType)
							} else {
								client.SendMessage(to, botstate.Fancy("Please add another bot that has a ban kick."))
							}
						}
					}
				}
			
			} else if strings.HasPrefix(cmd, "ban") && cmd != "bans" {
				rngcmd := handler.GetComs(7, "ban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 8
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.CheckListAccess(client, to, listuser, x, sender)
						}
					}
				}
			} else if strings.HasPrefix(cmd, "contact") {
				rngcmd := handler.GetComs(5, "contact")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								client.SendContact(to, i)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "bio") {
				rngcmd := handler.GetComs(7, "bio")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendMessage(to, botstate.Fancy(x.StatusMessage))
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "tag") {
				rngcmd := handler.GetComs(7, "tag")
				if handler.CheckPermission(rngcmd, sender, to) {
					listuser := []string{}
					nCount := 0
					fl := strings.Split(cmd, " ")
					typec := strings.Replace(cmd, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for i := range lists {
							if !utils.InArrayString(listuser, lists[i]) {
								listuser = append(listuser, lists[i])
							}
						}
						client.SendPollMention(to, "Tag Users:", listuser)
					}
				}
			} else if strings.HasPrefix(cmd, "story") {
				rngcmd := handler.GetComs(7, "story")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.Timeline(to, x.Mid,"getstory","getstory")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "cover") {
				rngcmd := handler.GetComs(7, "cover")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.Timeline(to, x.Mid,"cover","cover")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "extracover") {
				rngcmd := handler.GetComs(7, "extracover")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.Timeline(to, x.Mid,"extracover","extracover")
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "image") {
				rngcmd := handler.GetComs(7, "image")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendFoto(to, "https://profile.line-scdn.net/"+x.PictureStatus)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "getvideo") {
				rngcmd := handler.GetComs(7, "getvideo")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								x, _ := client.GetContact(i)
								client.SendVideoWithURL(to, "https://profile.line-scdn.net/"+x.PictureStatus)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "zoom") {
				rngcmd := handler.GetComs(7, "zoom")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						nCount := 0
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for _, i := range lists {
								zom := "https://api.minzteam.xyz/zoom?url="+i+"&apikey="+botstate.Apikey
								client.SendImageWithURL(to, zom)
							}
						}
					}
				}
			} else if strings.HasPrefix(text, Sname+"nunban ") {
				rngcmd := handler.GetComs(7, "nunban ")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						str := strings.Replace(text, Sname+"nunban ", "", 1)
						nm := []string{}
						target := []string{}
						count := strings.Split(str, ",")
						for c, nmr := range count {
							num, _ := strconv.Atoi(nmr)
							if num > 0 && num <= len(botstate.Banned.Banlist) {
								target = append(target, botstate.Banned.Banlist[num-1])
								pr, _ := client.GetContact(botstate.Banned.Banlist[num-1])
								name := pr.DisplayName
								c += 1
								name = fmt.Sprintf(". %s", name)
								nm = append(nm, name)
							}
						}
						if len(target) == 0 {
							newsend += "User not found.\n"
						} else {
	                                       for _, from := range target {
		                                       if botstate.Banned.GetBan(from) {
			                                       botstate.Banned.DelBan(from)
						            }
						      }
						}
						stx := strings.Join(nm, "\n")
						newsend += "Unban:\n\n"+stx
					}
				}
			} else if strings.HasPrefix(cmd, "unban") {
				rngcmd := handler.GetComs(7, "unban")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						listuser := []string{}
						nCount := 0
						x := 1
						fl := strings.Split(cmd, " ")
						typec := strings.Replace(cmd, fl[0]+" ", "", 1)
						re := regexp.MustCompile("([a-z]+)([0-9]+)")
						matches := re.FindStringSubmatch(typec)
						if len(matches) == 3 {
							typec = matches[1]
							nCount, _ = strconv.Atoi(matches[2])
						}
						if nCount == 0 {
							nCount = 1
						}
						lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
						if len(lists) != 0 {
							for i := range lists {
								if !utils.InArrayString(listuser, lists[i]) {
									listuser = append(listuser, lists[i])
								}
							}
							handler.handler.CheckUnbanBots(client, to, listuser, x, sender)
						} else {
							result := strings.Split((cmd), " ")
							if len(result) > 1 {
								result2, botstate.Err := strconv.Atoi(result[1])
								if botstate.Err != nil {
									client.SendMessage(to, botstate.Fancy("𝗣𝗹𝗲𝗮𝘀𝗲 𝗽𝘂𝘁 𝗮 𝗻𝘂𝗺𝗯𝗲𝗿"))
									return
								} else {
									if result2 > 0 {
										su := "unban"
										str := ""
										if strings.HasPrefix(text, Rname+" ") {
											str = strings.Replace(text, Rname+" "+su+" ", "", 1)
											str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname+" ") {
											str = strings.Replace(text, Sname+" "+su+" ", "", 1)
											str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Rname) {
											str = strings.Replace(text, Rname+su+" ", "", 1)
											str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
										} else if strings.HasPrefix(text, Sname) {
											str = strings.Replace(text, Sname+su+" ", "", 1)
											str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
										}
										st := handler.StripOut(str)
										hapuss := linetcr.Archimed(st, botstate.Banned.Banlist)
										if len(hapuss) == 0 {
											newsend += "User not found.\n"
										} else {
											handler.handler.CheckUnbanBots(client, to, hapuss, x, sender)
										}
									}
								}
							} else {
								newsend += "User not found.\n"
							}
						}
					}
				}
			} else if cmd == "pronote on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProNote {
							newsend += "Already enabled.\n"
						} else {
							room.ProNote = true
							newsend += "Protect Note Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "pronote off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProNote {
							newsend += "Already disabled.\n"
						} else {
							room.ProNote = false
							newsend += "Protect Note Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "progpict on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProPicture {
							newsend += "Already enabled.\n"
						} else {
							room.ProPicture = true
							newsend += "Protect Picture Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "progpict off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProPicture {
							newsend += "Already disabled.\n"
						} else {
							room.ProPicture = false
							newsend += "Protect Picture Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proalbum on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProAlbum {
							newsend += "Already enabled.\n"
						} else {
							room.ProAlbum = true
							newsend += "Protect Album Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proalbum off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProAlbum {
							newsend += "Already disabled.\n"
						} else {
							room.ProAlbum = false
							newsend += "Protect Album Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prolink on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProLink {
							newsend += "Already enabled.\n"
						} else {
							room.ProLink = true
							newsend += "Protect Link Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prolink off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProLink {
							newsend += "Already disabled.\n"
						} else {
							room.ProLink = false
							newsend += "Protect Link Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proflex on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProFlex {
							newsend += "Already enabled.\n"
						} else {
							room.ProFlex = true
							newsend += "Protect Flex Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proflex off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProFlex {
							newsend += "Already disabled.\n"
						} else {
							room.ProFlex = false
							newsend += "Protect Flex Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proimage on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProImage {
							newsend += "Already enabled.\n"
						} else {
							room.ProImage = true
							newsend += "Protect Image Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proimage off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProImage {
							newsend += "Already disabled.\n"
						} else {
							room.ProImage = false
							newsend += "Protect Image Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "provideo on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProVideo {
							newsend += "Already enabled.\n"
						} else {
							room.ProVideo = true
							newsend += "Protect Video Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "provideo off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProVideo {
							newsend += "Already disabled.\n"
						} else {
							room.ProVideo = false
							newsend += "Protect Video Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "procall on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProCall {
							newsend += "Already enabled.\n"
						} else {
							room.ProCall = true
							newsend += "Protect Call Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "procall off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProCall {
							newsend += "Already disabled.\n"
						} else {
							room.ProCall = false
							newsend += "Protect Call Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prospam on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProSpam {
							newsend += "Already enabled.\n"
						} else {
							room.ProSpam = true
							newsend += "Protect Spamcall Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prospam off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProSpam {
							newsend += "Already disabled.\n"
						} else {
							room.ProSpam = false
							newsend += "Protect Spamcall Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prosticker on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProSticker {
							newsend += "Already enabled.\n"
						} else {
							room.ProSticker = true
							newsend += "Protect Sticker Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prosticker off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProSticker {
							newsend += "Already disabled.\n"
						} else {
							room.ProSticker = false
							newsend += "Protect Sticker Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "procontact on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProContact {
							newsend += "Already enabled.\n"
						} else {
							room.ProContact = true
							newsend += "Protect Contact Turn on\n"
					       }
						handler.SaveBackup()
					}
				}
			} else if cmd == "procontact off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProContact {
							newsend += "Already disabled.\n"
						} else {
							room.ProContact = false
							newsend += "Protect Contact Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "propost on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProPost {
							newsend += "Already enabled.\n"
						} else {
							room.ProPost = true
							newsend += "Protect Post Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "propost off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProPost {
							newsend += "Already disabled.\n"
						} else {
							room.ProPost = false
							newsend += "Protect Post Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "profile on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProFile {
							newsend += "Already enabled.\n"
						} else {
							room.ProFile = true
							newsend += "Protect File Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "profile off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProFile {
							newsend += "Already disabled.\n"
						} else {
							room.ProFile = false
							newsend += "Protect File Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prokick on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProKick {
							newsend += "Already enabled.\n"

						} else {
							room.ProKick = true
							newsend += "Protect Kick Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "prokick off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProKick {
							newsend += "Already disabled.\n"
						} else {
							room.ProKick = false
							newsend += "Protect Kick Turn off\n"
						}
						handler.SaveBackup()
					}

				}
			} else if cmd == "announce on" {
				//rngcmd := handler.GetComs(8, "on")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Announce {
							newsend += "Already enabled.\n"
						} else {
							room.Announce = true
							newsend += "Announcement is enabled.\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "announce off" {
				//rngcmd := handler.GetComs(8, "off")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Announce {
							room.Announce = false
							newsend += "Announcement is disabled.\n"
						} else {
							newsend += "Already disabled.\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proqr on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProQr {
							newsend += "Already enabled.\n"
						} else {
							room.ProQr = true
							newsend += "Protect Qr Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proqr off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProQr {
							newsend += "Already disabled.\n"
						} else {
							room.ProQr = false
							newsend += "Protect Qr Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proinvite on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProInvite {
							newsend += "Already enabled.\n"
						} else {
							room.ProInvite = true
							newsend += "Protect Invite Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "proinvite off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProInvite {
							newsend += "Already disabled.\n"
						} else {
							room.ProInvite = false
							newsend += "Protect Invite Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "automute on" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Automute {
							newsend += "Already enabled.\n"
						} else {
							room.Automute = true
							newsend += "Automute enabled.\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "automute off" {
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.Automute {
							newsend += "Already disabled.\n"
						} else {
							room.Automute = false
							newsend += "Automute disabled.\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "procancel on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProCancel {
							newsend += "Already enabled.\n"
						} else {
							room.ProCancel = true
							newsend += "Protect Cancel Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "procancel off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProCancel {
							newsend += "Already disabled.\n"
						} else {
							room.ProCancel = false
							newsend += "Protect Cancel Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "projoin on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProJoin {
							newsend += "Already enabled.\n"
						} else {
							room.ProJoin = true
							newsend += "Protect Join Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "projoin off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProJoin {
							newsend += "Already disabled.\n"
						} else {
							room.ProJoin = false
							newsend += "Protect Join Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "progname on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProName {
							newsend += "Already enabled.\n"
						} else {
							room.ProName = true
							newsend += "Protect Group Name Turn on\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "progname off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProName {
							newsend += "Already disabled.\n"
						} else {
							room.ProName = false
							newsend += "Protect Group Name Turn off\n"
						}
						handler.SaveBackup()
					}
				}
			} else if cmd == "leave on" {
				//rngcmd := handler.GetComs(8, "on")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Leavebool {
							newsend += "Already enabled.\n"
						} else {
							room.Leavebool = true
							newsend += "Leave Message Turn on\n"
						}
					}
				}
			} else if cmd == "sendimage on" {
				//rngcmd := handler.GetComs(4, "sendimage")
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ImageLurk {
							newsend += "Already enabled.\n"
						} else {
							room.ImageLurk = true
							newsend += "Sendimage set enabled.\n"
						}
					}
				}
			} else if cmd == "sendimage off" {
				//rngcmd := handler.GetComs(4, "sendimage")
				if handler.CheckPermission(5, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ImageLurk {
							newsend += "Already disabled.\n"
						} else {
							room.ImageLurk = false
							newsend += "Sendimage set disabled.\n"
						}
					}
				}
			} else if cmd == "leave off" {
				//rngcmd := handler.GetComs(8, "off")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.Leavebool {
							newsend += "Already disabled.\n"
						} else {
							room.Leavebool = false
							newsend += "Leave Message Turn off\n"
						}
					}
				}
			} else if cmd == "welcome on" {
				//rngcmd := handler.GetComs(8, "on")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Welcome {
							newsend += "Already enabled.\n"
						} else {
							room.Welcome = true
							newsend += "Welcome Message Turn on\n"
						}
					}
				}
			} else if cmd == "welcome off" {
				//rngcmd := handler.GetComs(8, "off")
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.Welcome {
							newsend += "Already disabled.\n"
						} else {
							room.Welcome = false
							newsend += "Welcome Message Turn off\n"
						}
					}
				}
			} else if cmd == "backuser on" {
				rngcmd := handler.GetComs(5, "backuser")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Backup {
							newsend += "Already enabled.\n"
						} else {
							room.Backup = true
							newsend += "backup user set enabled.\n"
						}
					}
				}
			} else if cmd == "backuser off" {
				rngcmd := handler.GetComs(5, "backuser")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.Backup {
							newsend += "Already disabled.\n"
						} else {
							room.Backup = false
							newsend += "backup user set disabled.\n"
						}
					}
				}
			} else if cmd == "hostage on" {
				rngcmd := handler.GetComs(5, "hostage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Backleave {
							newsend += "Already enabled.\n"
						} else {
							room.Backleave = true
							newsend += "hostage set enabled.\n"
						}
					}
				}
			} else if cmd == "hostage off" {
				rngcmd := handler.GetComs(5, "hostage")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.Backleave {
							newsend += "Already disabled.\n"
						} else {
							room.Backleave = false
							newsend += "hostage set disabled.\n"
						}
					}
				}
			} else if cmd == "allprotect off" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProName && !room.ProCancel && !room.ProInvite && !room.ProKick && !room.ProQr && !room.ProJoin && !room.ProPicture && !room.ProNote && !room.ProAlbum && !room.ProLink && !room.ProFlex && !room.ProImage && !room.ProVideo && !room.ProCall && !room.ProSpam && !room.ProSticker && !room.ProContact && !room.ProPost && !room.ProFile && !room.AntiTag {
							newsend += "All protection is Already disabled.\n"
						} else {
						      room.ProCancel = false
						      room.ProInvite = false
						      room.ProKick = false
						      room.ProQr = false
						      room.ProName = false
						      room.ProJoin = false
						      room.ProPicture = false
						      room.ProNote = false
						      room.ProAlbum = false
						      room.ProLink = false
						      room.ProFlex = false
						      room.ProImage = false
						      room.ProVideo = false
						      room.ProCall = false
						      room.ProSpam = false
						      room.ProSticker = false
						      room.ProContact = false
						      room.ProPost = false
						      room.ProFile = false
						      room.AntiTag = false
						      handler.SaveBackup()
						      newsend += "All Protect Turn off\n"
						}
					}
				}
			} else if cmd == "allprotect on" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProName && room.ProCancel && room.ProInvite && room.ProKick && room.ProQr && room.ProJoin && room.ProPicture && room.ProNote && room.ProAlbum && room.ProLink && room.ProFlex && room.ProImage && room.ProVideo && room.ProCall && room.ProSpam && room.ProSticker && room.ProContact && room.ProPost && room.ProFile && room.AntiTag {
							newsend += "All protection is Already enabled.\n"
						} else {
						      room.ProCancel = true
						      room.ProInvite = true
						      room.ProKick = true
						      room.ProQr = true
						      room.ProName = true
						      room.ProJoin = true
						      room.ProPicture = true
						      room.ProNote = true
						      room.ProAlbum = true
						      room.ProLink = true
						      room.ProFlex = true
						      room.ProImage = true
						      room.ProVideo = true
						      room.ProCall = true
						      room.ProSpam = true
						      room.ProSticker = true
						      room.ProContact = true
						      room.ProPost = true
						      room.ProFile = true
						      room.AntiTag = true
						      handler.SaveBackup()
						      newsend += "All Protect Turn on\n"
						}
					}
				}
			} else if cmd == "protectmax on" || cmd == botstate.Commands.Max && botstate.Commands.Max != "" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.ProName && room.ProCancel && room.ProInvite && room.ProKick && room.ProQr {
							newsend += "Max protection is Already enabled.\n"
						} else {
							room.ProName = true
							room.ProCancel = true
							room.ProInvite = true
							room.ProKick = true
							room.ProQr = true
						      handler.SaveBackup()
							newsend += "Max Protect Turn on\n"
						}
					}
				}
			} else if cmd == "protectmax off" || cmd == botstate.Commands.None && botstate.Commands.None != "" {
				if handler.CheckPermission(9, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if !room.ProName && !room.ProCancel && !room.ProInvite && !room.ProKick && !room.ProQr {
							newsend += "Max protection is Already disabled.\n"
						} else {
							room.ProName = false
							room.ProCancel = false
							room.ProInvite = false
							room.ProKick = false
							room.ProQr = false
						      handler.SaveBackup()
							newsend += "Max Protect Turn off\n"
						}
					}
				}
			} else if cmd == "restartperm" {
				handler.Resprem()
				list := handler.PerCheckList()
				newsend += list
			} else if cmd == "nukejs" {
				rngcmd := handler.GetComs(5, "nukejs")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						_, memlist, _ := client.GetChatList(to)
						exe := []*linetcr.Account{}
						oke := []string{}
						for _, mid := range memlist {
							if utils.InArrayString(botstate.Squadlist, mid) {
								cl := handler.GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						max := len(exe) * 100
						lkick := []string{}
						for n, v := range memlist {
							if handler.MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*linetcr.Account{}
						ilen := len(lkick)
						xx := 0
						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							go func(to string, i int) {
								target := lkick[i]
								cl := nom[i]
								cl.Nodejs(to, target)
							}(to, i)
						}
						handler.LogAccess(client, to, sender, "kickall", lkick, msg.ToType)
					}
				}
			} else if cmd == "kickall" || cmd == botstate.Commands.Kickall && botstate.Commands.Kickall != "" {
				rngcmd := handler.GetComs(5, "kickall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						_, memlist, _ := client.GetChatList(to)
						exe := []*linetcr.Account{}
						oke := []string{}
						for _, mid := range memlist {
							if utils.InArrayString(botstate.Squadlist, mid) {
								cl := handler.GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						max := len(exe) * 100
						lkick := []string{}
						for n, v := range memlist {
							if handler.MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*linetcr.Account{}
						ilen := len(lkick)
						xx := 0
						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							go func(to string, i int) {
								target := lkick[i]
								cl := nom[i]
								cl.DeleteOtherFromChats(to, []string{target})
							}(to, i)
						}
						handler.LogAccess(client, to, sender, "kickall", lkick, msg.ToType)
					}
				}
			} else if cmd == "cancelall" || cmd == botstate.Commands.Cancelall && botstate.Commands.Cancelall != "" {
				rngcmd := handler.GetComs(5, "cancelall")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if room.Limit {
							client.SendMessage(to, botstate.Fancy("Sorry, all bot Here banned Try Invite Anther Bot"))
							return
						}
						_, memlist2, memlist := client.GetChatList(to)
						exe := []*linetcr.Account{}
						oke := []string{}
						for _, mid := range memlist2 {
							if utils.InArrayString(botstate.Squadlist, mid) {
								cl := handler.GetKorban(mid)
								if cl.Limited == false {
									exe = append(exe, cl)
								}
								oke = append(oke, mid)
							}
						}
						lkick := []string{}
						max := len(exe) * 10
						for n, v := range memlist {
							if handler.MemUser(to, v) {
								lkick = append(lkick, v)
							}
							if n > max {
								break
							}
						}
						nom := []*linetcr.Account{}
						ilen := len(lkick)
						xx := 0

						for i := 0; i < ilen; i++ {
							if xx < len(exe) {
								nom = append(nom, exe[xx])
								xx += 1
							} else {
								xx = 0
								nom = append(nom, exe[xx])
							}
						}
						for i := 0; i < ilen; i++ {
							target := lkick[i]
							cl := nom[i]
							ants.Submit(func() { cl.CancelChatInvitations(to, []string{target}) })
						}
						handler.LogAccess(client, to, sender, "cancelall", lkick, msg.ToType)
					}
				}
			} else if strings.HasPrefix(cmd, "joinqr") {
				rngcmd := handler.GetComs(5, "joinqr")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "joinqr"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, "/")
							tkt := client.FindChatByTicket(result[4])
							client.AcceptTicket(tkt.Chat.ChatMid, result[4])
							exe := []*linetcr.Account{}
							for _, p := range bk {
								if p.Limited == false {
									if botstate.Err == nil {
										exe = append(exe, p)
									}
								}
							}
							if len(exe) != 0 {
								newsend += "Succes Accept Group Ticket"
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "nukeqr") {
				rngcmd := handler.GetComs(5, "nukeqr")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						su := "nukeqr"
						str := ""
						if strings.HasPrefix(text, Rname+" ") {
							str = strings.Replace(text, Rname+" "+su+" ", "", 1)
							str = strings.Replace(str, Rname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname+" ") {
							str = strings.Replace(text, Sname+" "+su+" ", "", 1)
							str = strings.Replace(str, Sname+" "+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Rname) {
							str = strings.Replace(text, Rname+su+" ", "", 1)
							str = strings.Replace(str, Rname+strings.Title(su)+" ", "", 1)
						} else if strings.HasPrefix(text, Sname) {
							str = strings.Replace(text, Sname+su+" ", "", 1)
							str = strings.Replace(str, Sname+strings.Title(su)+" ", "", 1)
						}
						if len(str) != 0 {
							result := strings.Split(str, "/")
							tkt := client.FindChatByTicket(result[4])
							client.AcceptTicket(tkt.Chat.ChatMid, result[4])
							exe := []*linetcr.Account{}
							for _, p := range bk {
								if p.Limited == false {
									if botstate.Err == nil {
										exe = append(exe, p)
									}
								}
							}
							if len(exe) != 0 {
								go handler.NukJoin(exe[0], op.CreatedTime, tkt.Chat.ChatMid)
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "qrjoin") {
				rngcmd := handler.GetComs(5, "qrjoin")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if strings.Contains(pesan, "/ti/g") {
							regex, _ := regexp.Compile(`(?:line\:\/|line\.me\/R)\/ti\/g\/([a-zA-Z0-9_-]+)?`)
							links := regex.FindAllString(msg.Text, -1)
							tickets := []string{}
							for _, link := range links {
								if !utils.InArrayString(tickets, link) {
									tickets = append(tickets, link)
								}
							}
							for _, tick := range tickets {
								tuk := strings.Split(tick, "/")
								ntk := len(tuk) - 1
								ti := tuk[ntk]
								tkt := client.FindChatByTicket(ti)
								client.AcceptTicket(tkt.Chat.ChatMid, ti)
								exe := []*linetcr.Account{}
								for _, p := range bk {
									if p.Limited == false {
										if botstate.Err == nil {
											exe = append(exe, p)
										}
									}
								}
								if len(exe) != 0 {
									newsend += "Succes Accept Group Link"
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "qrjoinkick") {
				rngcmd := handler.GetComs(5, "qrjoinkick")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if strings.Contains(pesan, "/ti/g") {
							regex, _ := regexp.Compile(`(?:line\:\/|line\.me\/R)\/ti\/g\/([a-zA-Z0-9_-]+)?`)
							links := regex.FindAllString(msg.Text, -1)
							tickets := []string{}
							for _, link := range links {
								if !utils.InArrayString(tickets, link) {
									tickets = append(tickets, link)
								}
							}
							for _, tick := range tickets {
								tuk := strings.Split(tick, "/")
								ntk := len(tuk) - 1
								ti := tuk[ntk]
								tkt := client.FindChatByTicket(ti)
								client.AcceptTicket(tkt.Chat.ChatMid, ti)
								exe := []*linetcr.Account{}
								for _, p := range bk {
									if p.Limited == false {
										if botstate.Err == nil {
											exe = append(exe, p)
										}
									}
								}
								if len(exe) != 0 {
									go handler.NukJoin(exe[0], op.CreatedTime, to)
								}
							}
						}
					}
				}
			} else if strings.HasPrefix(cmd, "clone") {
				rngcmd := handler.GetComs(5, "clone")
				if handler.CheckPermission(rngcmd, sender, to) {
					if handler.handler.CheckExprd(client, to, sender) {
						if len(mentionlist) == 1 {
							cok := strings.Split((cmd), " ")
							if len(cok) > 1 {
								targets := ""
								var pp, vp, co, cv, name, stats string
								cok := strings.Split((cmd), " ")
								if len(cok) > 1 {
									ann := cok[1]
									var prof *talkservice.Contact
									if ann == "@me" {
										prof, _ = client.GetContact(msg.From_)
										targets = msg.From_
									}
									if prof != nil {
										name = prof.DisplayName
										stats = prof.StatusMessage
										if prof.VideoProfile != "" {
											ps, botstate.Err := client.Downloads("http://dl.profile.line-cdn.net"+prof.PicturePath+"/vp", "mp4")
											if botstate.Err != nil {
												client.SendMessage(to, botstate.Fancy("Download video profile error."))
											} else {
												vp = ps
											}
										}
										if prof.PicturePath != "" {
											ps, botstate.Err := client.Downloads("http://dl.profile.line.naver.jp"+prof.PicturePath, "jpg")
											if botstate.Err != nil {
												client.SendMessage(to, botstate.Fancy("Download picture profile error."))
											} else {
												pp = ps
											}
										}
										profs := client.GetProfileDetail(msg.From_)
										pss, botstate.Err := client.Downloads("https://obs.line-scdn.net/r/myhome/c/"+gjson.Get(profs, "result.objectId").String(), "jpg")
										if botstate.Err == nil {
											co = pss
										}
										pss, botstate.Err = client.Downloads("https://obs.line-scdn.net/r/myhome/vc/"+gjson.Get(profs, "result.objectId").String(), "mp4")
										if botstate.Err == nil {
											cv = pss
										}
										if len(mentionlist) != 0 {
											clon := false
											for _, target := range mentionlist {
												if target != targets && utils.InArrayString(botstate.Squadlist, target) {
													idx := handler.GetKorban(target)
													handler.Clone(idx, pp, vp, co, cv, name, stats)
													idx.SendMention(to, "Cloning @! profile done.", []string{targets})
													clon = true
												}
											}
											if !clon {
												if pp != "" {
													os.utils.RemoveString(pp)
												}
												if vp != "" {
													os.utils.RemoveString(vp)
												}
												if co != "" {
													os.utils.RemoveString(co)
												}
												if cv != "" {
													os.utils.RemoveString(cv)
												}
											}
										} else {
											if pp != "" {
												os.utils.RemoveString(pp)
											}
											if vp != "" {
												os.utils.RemoveString(vp)
											}
											if co != "" {
												os.utils.RemoveString(co)
											}
											if cv != "" {
												os.utils.RemoveString(cv)
											}
										}
									}
								}
							}
						}
					}
				}
			} else if botstate.PublicMode {
				 if pesan == "help" {
					res := "✠ 𝗣𝘂𝗯𝗹𝗶𝗰 𝗖𝗼𝗺𝗺𝗮𝗻𝗱𝘀:"
					res += "\n"
					for _, x := range botstate.Helppublic {
						res += fmt.Sprintf("\n %v %s", botstate.Data.Logobot, x)
					}
					res += "\n\n✠ 𝗧𝘆𝗽𝗲 𝗥𝗲𝗽𝗹𝘆:"
					res += "\n"
					for _, x := range botstate.Helpreply {
						res += fmt.Sprintf("\n %v %s", botstate.Data.Logobot, x)
					}
					client.SendHelp(to, res)
				} else if pesan == "mentionall" {
					_, target, _ := client.GetChatList(to)
					targets := []string{}
					for i := range target {
						if !utils.InArrayString(botstate.CheckHaid, target[i]) {
							targets = append(targets, target[i])
						}
					}
					client.SendPollMention(to, "Mentions member:\n", targets)
				} else if pesan == "groupinfo" {
					list := handler.InfoGroup(client, to)
					client.SendMessage(to, botstate.Fancy(list))
				} else if strings.HasPrefix(pesan, "detectcall ") {
					spl := strings.Replace(pesan, "detectcall ", "", 1)
					if spl == "on" {
						newsend += "Detectgroupcall is enabled.\n"
					} else if spl == "off" {
						newsend += "Detectgroupcall is disabled.\n"
					}
				} else if pesan == "welcome on" {
					if room.Welcome {
						newsend += "Already enabled.\n"
					} else {
						room.Welcome = true
						newsend += "Welcome Message Turn on\n"
					}
				} else if pesan == "welcome off" {
					if !room.Welcome {
						newsend += "Already disabled.\n"
					} else {
						room.Welcome = false
						newsend += "Welcome Message Turn off\n"
					}
				} else if pesan == "lurk on" {
					room.Lurk = true
					room.NameLurk = "mention"
					room.Userlurk = []string{}
					newsend += "Lurking enabled.\n"
				} else if pesan == "lurk off" {
					room.Lurk = false
					room.ImageLurk = false
					if len(room.Userlurk) != 0 {
						list := " ✠ Lurkers ✠ \n"
						for num, xd := range room.Userlurk {
							num++
							rengs := strconv.Itoa(num)
							name := handler.GetContactName(client, xd)
								list += "\n   " + rengs + ". " + name
						}
						newsend += list + "\n"
					}
					room.Userlurk = []string{}
				} else if strings.HasPrefix(pesan, "getpict") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.SendFoto(to, "https://profile.line-scdn.net/"+x.PictureStatus)
						}
					}
				} else if strings.HasPrefix(pesan, "getcontact") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							client.SendContact(to, i)
						}
					}
				} else if strings.HasPrefix(pesan, "getbio") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.SendMessage(to, botstate.Fancy(x.StatusMessage))
						}
					}
				} else if strings.HasPrefix(pesan, "getname") {
					listuser := []string{}
					nCount := 0
					x := 16
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for i := range lists {
							if !utils.InArrayString(listuser, lists[i]) {
								listuser = append(listuser, lists[i])
							}
						}
						handler.CheckListAccess(client, to, listuser, x, sender)
					}
				} else if strings.HasPrefix(pesan, "getmid") {
					listuser := []string{}
					nCount := 0
					x := 14
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for i := range lists {
							if !utils.InArrayString(listuser, lists[i]) {
								listuser = append(listuser, lists[i])
							}
						}
						handler.CheckListAccess(client, to, listuser, x, sender)
					}
				} else if strings.HasPrefix(pesan, "getcover") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.Timeline(to, x.Mid,"cover","cover")
						}
					}
				} else if strings.HasPrefix(pesan, "gextracover") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.Timeline(to, x.Mid,"extracover","extracover")
						}
					}
				} else if strings.HasPrefix(pesan, "getstory") {
					nCount := 0
					fl := strings.Split(pesan, " ")
					typec := strings.Replace(pesan, fl[0]+" ", "", 1)
					re := regexp.MustCompile("([a-z]+)([0-9]+)")
					matches := re.FindStringSubmatch(typec)
					if len(matches) == 3 {
						typec = matches[1]
						nCount, _ = strconv.Atoi(matches[2])
					}
					if nCount == 0 {
						nCount = 1
					}
					lists := handler.LlistCheck(client, to, typec, nCount, sender, Rplay, mentionlist)
					if len(lists) != 0 {
						for _, i := range lists {
							x, _ := client.GetContact(i)
							client.Timeline(to, x.Mid,"getstory","getstory")
						}
					}
				} else if pesan == "getcall" {
                    		         if msg.ToType == 2 {
                        		        gcall, _ := call.GetGroupCall(to, client.AuthToken)
                                         Room := linetcr.GetRoom(to)
                        		        res := "Get Call Group:"
                        		        if gcall.MediaType == 1 {
                                                res += "\n  • Type: Audio Call"
                        		               res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                loc, _ := time.LoadLocation("Asia/Jakarta")
	                                         a := time.Now().In(loc)
	                                         yyyy := strconv.Itoa(a.Year())
	                                         MM := a.Month().String()
	                                         dd := strconv.Itoa(a.Day())
	                                         Date := dd + "-" + MM + "-" + yyyy
                                                cok := gcall.Started / 1000
			                            i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                            tm := time.Unix(i, 0)
			                            ss := time.Since(tm)
			                            sp := handler.FmtDurations(ss)
                                                res += "\n  • Date: "+Date
                                                res += "\n  • Started: "+sp
                        		               res += "\n  • Members:"
                                                mmk := gcall.MemberMids
						        if len(mmk) != 0 {
							        for num, xd := range mmk {
								        num++
								        rengs := strconv.Itoa(num)
								        x, _ := client.GetContact(xd)
								        res += "\n      " + rengs + ". "+x.DisplayName
							          }
					                 }
                                                       client.SendMessage(to, botstate.Fancy(res))
                                         }
                        		        if gcall.MediaType == 2 {
                                                res += "\n  • Type: Video Call"
                        		               res += fmt.Sprintf("\n  • Group: %s", Room.Name)
                                                loc, _ := time.LoadLocation("Asia/Jakarta")
	                                         a := time.Now().In(loc)
	                                         yyyy := strconv.Itoa(a.Year())
	                                         MM := a.Month().String()
	                                         dd := strconv.Itoa(a.Day())
	                                         Date := dd + "-" + MM + "-" + yyyy
                                                cok := gcall.Started / 1000
			                            i, _ := strconv.ParseInt(fmt.Sprintf("%v", cok), 10, 64)
			                            tm := time.Unix(i, 0)
			                            ss := time.Since(tm)
			                            sp := handler.FmtDurations(ss)
                                                res += "\n  • Date: "+Date
                                                res += "\n  • Started: "+sp
                        		               res += "\n  • Members:"
                                                mmk := gcall.MemberMids
						        if len(mmk) != 0 {
							        for num, xd := range mmk {
								        num++
								        rengs := strconv.Itoa(num)
								        x, _ := client.GetContact(xd)
								        res += "\n      " + rengs + ". "+x.DisplayName
							          }
					                 }
                                                   client.SendMessage(to, botstate.Fancy(res))
                                           }
					  }
				} else if len(botstate.Data.WordbanBack) != 0 {
					for _, selftcr := range botstate.Data.WordbanBack {
						if pesan == selftcr {
							if handler.MemUser(msg.To, msg.From_) {
								exe, _ := handler.SelectBot(client, to)
								if exe != nil {
									handler.Setkickto(exe, to, []string{msg.From_})
									botstate.Banned.AddBan(msg.From_)
								}
							}
						}					
					}
				}
			}
		}
		if newsend != "" {
			client.SendMessage(to, botstate.Fancy(newsend))
		}
	}
}

