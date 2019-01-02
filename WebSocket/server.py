from websocket_server import WebsocketServer

def connect_client(client, server):
    message = "[*] Connected -> {}:{}".format(client["address"][0], client["address"][1])
    print(message)
    server.send_message_to_all(message)

def leave_client(client, server):
    message = "[*] Left -> {}:{}".format(client["address"][0], client["address"][1])
    print(message)
    server.send_message_to_all(message)

def receive_message(client, server, _message):
    message = "[*] Chat {}:{} < {}".format(
        client["address"][0],
        client["address"][1],
        _message
    )
    print(message)

    server.send_message_to_all(message)

def main():
    server = WebsocketServer(host='0.0.0.0', port=9999)
    server.set_fn_new_client(connect_client)
    server.set_fn_client_left(leave_client)
    server.set_fn_message_received(receive_message)
    server.run_forever()

if __name__ == '__main__':
    main()