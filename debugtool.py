#!/usr/bin python
#coding:utf8
###############################################
# File Name: debugtool.py
# Author: 郭 璞
# mail: marksinoberg@gmail.com
# Created Time: 五  4/16 19:32:52 2021
# Description: 代理调试脚本
###############################################
import gzip
import requests

url = "http://localhost:8080/proxy?position=fishingtower&price=0dollar/day&triggerword=triggerget1"
payload = {"name": "泰戈尔", "cuid": "", "address": "bj", "triggerword": "triggerpost1"}
headers = {
#        "cuid": "fromheader",
        "cookie": "_ga=GA1.2.1122705097.1588154719; cuid=fromcookie; __gads=ID=30b2fc5ad1ef9e12:T=1588154718:S=ALNI_MY9jL89Mvo3LuOEVCEVRUQBiRQfmA; Hm_lvt_1222b65eebb60dfc52490a24f1446876=1600338479; gr_user_id=982217ab-bfff-4b4a-98bf-c6e4028c52e2; grwng_uid=c4d23aa2-7712-4663-9bee-9bdb44d1a9d7; yjs_id=51321aba1e03877d137c73682d849b9b; Hm_lvt_31deec06ac5374dcfa09b9686205ade7=1602837640; UM_distinctid=176e136030239c-0986ffde0063c9-16386153-1fa400-176e13603037fa; CNZZDATA5032359=cnzz_eid%3D409741483-1598269483-https%253A%252F%252Frecomm.cnblogs.com%252F%26ntime%3D1610092590; Hm_lvt_39b794a97f47c65b6b2e4e1741dcba38=1617543416,1617547763; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8L-rpLgFVEJMgssCVvNUAjvXmXzX_C_Hzk2kdqYnwWToo9lWQuwy9ZCyzkIbX8lcbLtBwFIOCrH0yt_8pEz4lOSWVgcFW2sJlsjM77LQuxtqyKwvwBxDFrg5VaTit_Aob7-bLHGh1bYtHeDYiTm1ka4; _gid=GA1.2.957957947.1618542016; _gat=1"
        }
response = requests.post(url, data=payload, headers=headers)

content_type = response.headers.get("Content-Type")
print("content_type=", content_type)
content = ""
if content_type is not None and "gzip" in content_type:
    content = gzip.decompress(response.content)
    print("解码效果")
else:
    content = response.text
print(content)
#print(response.json())
#print(response.text)
