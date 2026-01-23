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
def old_main_menu(picture, name, text=".", text2=".", text3=".", text4=".", text5=".", text6=".", text7=".", text8=".", text9=".", text10="."):
    dataProfile = {
  "type": "bubble",
  "size": "kilo",
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "image",
        "url": "https://s20.directupload.net/images/240222/oehabpyg.jpg", 
        "aspectRatio": "1.02:1.31",
        "size": "full",
        "aspectMode": "cover"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "image",
            "url": "https://i.ibb.co/0yzgJys/PUTAR-KANAN.png",
            "animated": True,
            "size": "sm",
            "aspectRatio": "1:1",
            "aspectMode": "cover"
          }
        ],
        "width": "55.5px",
        "height": "55.5px",
        "borderWidth": "0px",
        "borderColor": "#808080",
        "cornerRadius": "100px",
        "offsetBottom": "265px",
        "offsetStart": "12px",
        "position": "absolute"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "image",
            "url": picture,
            "aspectRatio": "1:1",
            "aspectMode": "cover",
            "size": "57px",
            "action": {
              "type": "uri",
              "uri": picture
            }
          }
        ],
        "position": "absolute",
        "cornerRadius": "100px",
        "offsetBottom": "270px",
        "offsetStart": "17px",
        "width": "45px",
        "height": "45px",
        "borderColor": "#000000",
        "borderWidth": "1px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": "𝗛𝗘𝗟𝗣 𝗠𝗘𝗡𝗨",
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold",
            "align": "center"
          },
          {
            "type": "separator",
            "margin": "sm",
            "color": "#e2323e"
          }
        ],
        "position": "absolute",
        "offsetStart": "95px",
        "offsetTop": "27px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetStart": "20px",
        "offsetTop": "100px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text2,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetTop": "130px",
        "offsetStart": "20px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text3,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetTop": "160px",
        "offsetStart": "20px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text4,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetTop": "190px",
        "offsetStart": "20px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text5,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetTop": "220px",
        "offsetStart": "20px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text6,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetStart": "170px",
        "offsetTop": "100px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text7,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetStart": "170px",
        "offsetTop": "130px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text8,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetStart": "170px",
        "offsetTop": "160px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text9,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetStart": "170px",
        "offsetTop": "190px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "text",
            "text": text10,
            "color": "#e2e2e3",
            "size": "sm",
            "weight": "bold"
          }
        ],
        "position": "absolute",
        "offsetStart": "170px",
        "offsetTop": "220px"
      },
      {
        "type": "box",
        "layout": "vertical",
        "contents": [
          {
            "type": "box",
            "layout": "vertical",
            "contents": [
              {
                "type": "text",
                "text": "𝗦𝗘𝗟𝗙𝗕𝗢𝗧 𝗣𝗥𝗘𝗠𝗜𝗨𝗠",
                "size": "xs",
                "color": "#e2e2e3",
                "weight": "bold",
                "align": "center"
              }
            ],
            "action": {
              "type": "uri",
              "label": "action",
              "uri": "https://line.me/ti/p/~code-bot"
            }
          }
        ],
        "position": "absolute",
        "offsetBottom": "0px",
        "offsetStart": "0px",
        "offsetEnd": "0px",
        "backgroundColor": "#00000066",
        "paddingAll": "20px",
        "paddingTop": "18px"
      }
    ],
    "paddingAll": "0px"
  }
}
    return dataProfile

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
