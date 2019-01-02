from websocket import create_connection

ws = create_connection("ws://127.0.0.1:9999/")
print(ws.recv())

while True:
    msg = input("")

    if msg == "":
        break

    ws.send(msg)
    response = ws.recv()
    print(response)

ws.close()