# -*- coding: utf-8 -*-
# GOLANG TEMPLATE
# TYPE: TEMPLATE
# CREATE BY: SELFTCR™
# HELPER: 2023
# ID LINE: code-bot
# WHATSAPP: +6282278984821
#==========IMPORT==========#
from MySplit import MySplit
from linepy import *
from liff.ttypes import LiffChatContext, LiffContext, LiffSquareChatContext, LiffNoneContext, LiffViewRequest
import sys
import re, ast, httpx, random, html5lib, requests, os, json

#==========LOGIN_MAIN==========#
TOKEN = sys.argv[1]
to = sys.argv[2]
text = sys.argv[3]
id = sys.argv[4]

Tcr = LINE(TOKEN, appName="IOS\t12.6.1\tiOS\t15.5")
#================MAIN_FLEX==============
def allowLiff():
    liffId = id[0:10]
    url = 'https://access.line.me/dialog/api/permissions'
    data = {'on': ['P', 'CM'], 'off': []}
    headers = {'X-Line-Access': Tcr.authToken, 'X-Line-Application': 'IOS\t12.6.1\tiOS\t15.5', 'X-Line-ChannelId': '{}'.format(liffId), 'Content-Type': 'application/json'}
    requests.post(url, json=data, headers=headers)

def sendTemplate(to, data):
    allowLiff()
    xyz = LiffChatContext(to)
    xyzz = LiffContext(chat=xyz)
    view = LiffViewRequest(id, xyzz)
    token = Tcr.liff.issueLiffView(view)
    url = 'https://api.line.me/message/v3/share'
    headers = {'Content-Type': 'application/json','Authorization': 'Bearer %s' % token.accessToken}
    data = {"messages":[data]}
    requests.post(url, headers=headers, data=json.dumps(data))
#================FLEX_MENU==============
def flex(pict, logoname, text):
    data = {"type": "bubble", "size": "micro", "header": {"type": "box", "layout": "vertical", "contents": [{"type": "image", "url": "https://s20.directupload.net/images/230723/8eap5t7w.jpg", "aspectRatio": "750:1700", "aspectMode": "cover", "position": "absolute", "size": "full"}, {"type": "image", "url": "https://s20.directupload.net/images/230723/8eap5t7w.jpg", "aspectRatio": "750:1700", "aspectMode": "cover", "size": "full", "position": "absolute", "offsetTop": "347.5px"}, {"type": "box", "layout": "vertical", "contents": [{"type": "image", "url": "https://i.ibb.co/dDKypRk/top-msg-robots.png", "aspectMode": "cover", "aspectRatio": "215:67.5", "size": "full", "animated": True}, {"type": "box", "layout": "vertical", "contents": [{"type": "image", "url": "https://i.ibb.co/QX3yxbH/hlth-bluefire.png", "size": "full", "aspectMode": "cover", "aspectRatio": "1:1", "animated": True}, {"type": "box", "layout": "vertical", "contents": [{"type": "image", "url": pict, "aspectMode": "cover"}], "position": "absolute", "width": "38px", "height": "38px", "cornerRadius": "100px", "offsetTop": "1.1px", "offsetStart": "1.1px"}], "width": "40px", "height": "40px", "position": "absolute", "cornerRadius": "100px", "offsetStart": "10px", "offsetTop": "6.3px"}, {"type": "box", "layout": "vertical", "contents": [{"type": "text", "text": logoname, "size": "xxs", "color": "#FFFFFF"}], "width": "70px", "height": "20px", "position": "absolute", "offsetStart": "75px", "offsetTop": "17.5px", "justifyContent": "center", "alignItems": "center"}], "paddingAll": "0px"}, {"type": "box", "layout": "vertical", "contents": [{"type": "box", "layout": "vertical", "contents": [{"type": "text", "text": text, "color": "#FFFFFF", "size": "10px", "wrap": True,"action": {"type": "uri","label": "action","uri": "https://line.me/ti/p/~code-bot"}}], "paddingAll": "4px", "backgroundColor": "#50505075", "cornerRadius": "3px"}], "paddingAll": "7px"}, {"type": "box", "layout": "vertical", "contents": [{"type": "image", "url": "https://i.ibb.co/SKHfcvR/bottom-robots-msg.png", "aspectMode": "cover", "aspectRatio": "215:50", "size": "full", "animated": True}], "paddingAll": "0px"}], "paddingAll": "0px"} }
    return data

#================FLEX_SEND==============

def sendflexON(to, text):
    logoname = "GO_FLEX"
    pict = "https://s20.directupload.net/images/230926/t77p8smw.jpg"
    data = flex(pict, logoname, text)
    sendTemplate(to,{"type": "flex", "altText": "Golang_SELFTCR™", "contents": data})

sendflexON(to, text)
#=====================================
