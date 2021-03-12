# This script isn't meant to be used with takeover.
# I made this script just to help me update the providers.json file.

import requests
import json
from bs4 import BeautifulSoup

URL = "https://github.com/EdOverflow/can-i-take-over-xyz/blob/master/README.md"

r = requests.get(URL)
soup = BeautifulSoup(r.content, "html.parser")

table = soup.find("table")
table_body = soup.find_all("tbody")
rows = soup.find_all("tr")
data = []
for row in rows:
    website = {}
    cols = row.find_all("td")
    if cols and cols[1].text.strip() == "Vulnerable":
        website["name"] = cols[0].text.strip()
        website["cname"] = [" "]
        website["respomse"] = [cols[2].text.strip()]
        data.append(website)
json_data = json.dumps(data)
print(json_data)
