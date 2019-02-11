import requests

def main():
    size = int(input("size >> "))
    data = input("data >> ")

    url = "https://chart.googleapis.com/chart"
    params = {
        "cht": "qr",
        "chs": "{}x{}".format(size, size),
        "chl": data
    }

    qr_image = requests.get(url, params=params).content

    with open("./qrcode.png", "wb") as f:
        f.write(qr_image)

if __name__ == '__main__':
    main()