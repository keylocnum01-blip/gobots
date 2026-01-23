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
profile = sys.argv[5]

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
def Reader(text, pict):
    data = {
  "type": "bubble",
  "size": "micro",
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "box",
        "layout": "horizontal",
        "contents": [
          {
            "type": "box",
            "layout": "vertical",
            "contents": [
              {
                "type": "box",
                "layout": "vertical",
                "contents": [
                  {
                    "type": "image",
                    "url": pict,
                    "aspectMode": "cover",
                    "aspectRatio": "1:1",
                    "size": "full",
                    "align": "center"
                  }
                ],
                "width": "43px",
                "height": "43px",
                "cornerRadius": "600px",
                "backgroundColor": "#00FFFF",
                "borderWidth": "2px",
                "borderColor": "#00FFFF"
              }
            ],
            "width": "46px",
            "height": "46px",
            "cornerRadius": "600px",
            "backgroundColor": "#00b900",
            "justifyContent": "center",
            "alignItems": "center"
          },
          {
            "type": "box",
            "layout": "vertical",
            "contents": [
              {
                "type": "text",
                "align": "center",
                "contents": [
                  {
                    "type": "span",
                    "text": text,
                    "size": "9px",
                    "weight": "bold"
                  }
                ]
              }
            ],
            "justifyContent": "center"
          }
        ],
        "spacing": "5px",
        "paddingAll": "10px"
      }
    ],
    "paddingAll": "0px"
  }
}
    return data

#================FLEX_SEND==============
def sendflexON(to, text):
    pict = profile
    data = Reader(text, pict)
    sendTemplate(to,{"type": "flex", "altText": "Golang_SELFTCR™", "contents": data})

sendflexON(to, text)
#=====================================
