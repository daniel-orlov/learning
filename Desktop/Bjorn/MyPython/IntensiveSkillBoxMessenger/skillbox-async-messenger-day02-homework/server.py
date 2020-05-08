"""
Серверное приложение для соединений
"""
import asyncio
from typing import Optional
from asyncio import transports


class ClientProtocol(asyncio.Protocol):
    login: str
    server: 'Server'
    transport: transports.Transport

    def __init__(self, server: 'Server'):
        self.server = server
        self.login = None

    def data_received(self, data: bytes):
        """Decodes incoming messages, accepts login details, checks login for uniqueness,
        sends last 10 messages after login from chat history
        :param data: bytes
        """
        decoded = data.decode()
        print(decoded)

        #Accepting credentials, checking uniqueness
        if self.login is None:
            if decoded.startswith("login:"):
                new_client_login = decoded.replace("login:", "").replace("\r\n", "")
                for online_client in self.server.clients:
                    if online_client.login == new_client_login:
                        self.transport.write(
                            f"Sorry, login '{new_client_login}' is taken. Please, try a different one.".encode()
                        )
                        self.transport.abort()
                #Greets and sends last 10 messages from chat history
                else:
                    self.login = new_client_login
                    self.transport.write(
                        f"Hello, {self.login}!\nWe've just been talking about:".encode()
                    )
                    if len(self.server.chat_history) >= 10: #TODO: implement sending last messages if < 10
                        self.server.send_history(self.transport)
                    else:
                        self.transport.write(
                            f"Well, nothing, to be totally honest. We don't talk behind your back.".encode()
                        )

        #After login
        else:
            self.send_message(decoded)

    def send_message(self, message):
        """Accepts message, appends to chat history, sends to all online clients
        :param message: str
        """
        format_string = f"<{self.login}> {message}"
        encoded = format_string.encode()
        self.server.chat_history.append(format_string)

        for client in self.server.clients:
            if client.login != self.login:
                client.transport.write(encoded)

    #Overrides BaseProtocol Methods
    def connection_made(self, transport: transports.Transport):
        self.transport = transport
        self.server.clients.append(self)
        print('Connection established')

    # Overrides BaseProtocol Methods
    def connection_lost(self, exc):
        self.server.clients.remove(self)
        print('Connection lost')


class Server:
    clients: list

    def __init__(self):
        self.clients = []
        self.chat_history = []

    def send_history(self, transport: transports.Transport):
        """Sends last ten messages from the server chat history in chronological order.
        NB - Currently only gets invoked when len(self.chat_history) >= 10
        :param message: str
        """
        last_ten_messages = self.chat_history[-10:]
        for message in last_ten_messages:
            transport.write(f"{message}\n".encode())


    def create_protocol(self):
        return ClientProtocol(self)

    async def start(self):
        loop = asyncio.get_running_loop()

        coroutine = await loop.create_server(
            self.create_protocol,
            "127.0.0.1",
            8888
        )

        print("Server loaded...")

        await coroutine.serve_forever()


process = Server()
try:
    asyncio.run(process.start())
except KeyboardInterrupt:
    print("Server closed manually")
