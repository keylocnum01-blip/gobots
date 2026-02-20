package botstate

import (
	"../library/linetcr"
	"../utils"
)

func GetKorban(user string) *linetcr.Account {
	for _, cl := range ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}

func SquadMention(mlist []string) (m *linetcr.Account, b bool) {
	for _, l := range mlist {
		if utils.InArrayString(Squadlist, l) {
			cl := GetKorban(l)
			return cl, true
		}
	}
	return nil, false
}

func SelectBot(client *linetcr.Account, to string) (*linetcr.Account, error) {
	Err, _, memlist := client.GetGroupMembers(to)
	if Err != nil {
		return nil, Err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], Err
	}
	return nil, Err
}

func CheckBot(client *linetcr.Account, to string) (*linetcr.Account, error) {
	Err, _, memlist := client.GetGroupMembers(to)
	if Err != nil {
		return nil, Err
	}
	exe := []*linetcr.Account{}
	for mid, _ := range memlist {
		if utils.InArrayString(Squadlist, mid) {
			cl := GetKorban(mid)
			if !cl.Limited {
				exe = append(exe, cl)
			}
		}
	}
	if len(exe) != 0 {
		return exe[0], Err
	}
	return nil, Err
}
