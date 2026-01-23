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
def Respon(text):
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
            "contents": [],
            "width": "2px",
            "backgroundColor": "#00b900"
          },
          {
            "type": "box",
            "layout": "horizontal",
            "contents": [
              {
                "type": "box",
                "layout": "vertical",
                "contents": [
                  {
                    "type": "image",
                    "url": "https://i.ibb.co/4KWWqhf/line-color.png",
                    "align": "center",
                    "size": "full",
                    "aspectRatio": "1:1",
                    "aspectMode": "cover",
                    "animated": True,
                    "position": "absolute"
                  },
                  {
                    "type": "box",
                    "layout": "vertical",
                    "contents": [
                      {
                        "type": "text",
                        "text": "!",
                        "color": "#00b900",
                        "weight": "bold",
                        "size": "11px"
                      }
                    ],
                    "width": "15px",
                    "height": "15px",
                    "backgroundColor": "#ffffff",
                    "justifyContent": "center",
                    "alignItems": "center",
                    "cornerRadius": "600px"
                  }
                ],
                "width": "18px",
                "height": "18px",
                "cornerRadius": "600px",
                "justifyContent": "center",
                "alignItems": "center"
              },
              {
                "type": "box",
                "layout": "vertical",
                "contents": [
                  {
                    "type": "text",
                    "text": text,
                    "size": "9px",
                    "color": "#999999",
                    "margin": "4px",
                    "wrap": True
                  }
                ]
              }
            ],
            "spacing": "7px",
            "alignItems": "center"
          }
        ],
        "spacing": "6px"
      }
    ],
    "paddingAll": "8px",
    "backgroundColor": "#ffffff",
    "height": "100%",
    "width": "100%"
  }
}
    return data

#================FLEX_SEND==============
def sendflexON(to, text):
    data = Respon(text)
    sendTemplate(to,{"type": "flex", "altText": "Golang_SELFTCR™", "contents": data})

sendflexON(to, text)
#=====================================
