import requests
host='http://127.0.0.1:8080/{}'
baseinfo ='info'
headers={'Content-Type':'application/json'}
args=['name','code']
fill_str = '|'.join(args)
data={'sign':'123456','args':fill_str}
resp = requests.post(
    url=host.format(baseinfo),
    data=data
)
print(resp.json())