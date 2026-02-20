package botstate

import (
	"../library/linetcr"
)

func GetComs(gr int, data string) int {
	defer linetcr.PanicOnly()
	_, value := SetHelper.Rngcmd[data]
	if value == false {
		SetHelper.Rngcmd[data] = gr
	}
	xx := SetHelper.Rngcmd[data]
	return xx
}

func CheckAccount(user string) *linetcr.Account {
	for _, cl := range ClientBot {
		if cl.MID == user {
			return cl
		}
	}
	return nil
}

func CheckUser(client *linetcr.Account, group string) ([]*linetcr.Account, []string) {
	list := []string{}
	err, _, memlist := client.GetGroupMembers(group)
	if err != nil {
		return nil, nil
	}
	var accs []*linetcr.Account
	for _, mem := range memlist {
		cl := CheckAccount(mem)
		if cl != nil {
			accs = append(accs, cl)
			list = append(list, mem)
		}
	}
	return accs, list
}
