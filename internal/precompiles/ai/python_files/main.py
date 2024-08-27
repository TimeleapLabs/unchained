#!/usr/bin/env python

import asyncio
from websockets.server import serve
import gen
import translate
from sia import Sia
import warnings

# Ignore all warnings
warnings.filterwarnings("ignore")


async def echo(websocket):
    async for message in websocket:
        sia = Sia().set_content(bytearray(message))
        opcode = sia.read_uint16()
        if opcode == 0:
            response = gen.request_handler(sia)
            await websocket.send(bytes(response))
        elif opcode == 1:
            response = translate.request_handler(sia)
            await websocket.send(bytes(response))
        else:
            await websocket.send(b"\0")


async def main():
    async with serve(echo, "127.0.0.1", 8765):
        await asyncio.Future()  # run forever

if __name__ == "__main__":
    print("Server started on ws://localhost:8765")
    asyncio.run(main())
