class Sia:
    def __init__(self):
        self.index = 0
        self.content = bytearray()

    def seek(self, index):
        self.index = index
        return self

    def set_content(self, content):
        self.content = bytearray(content)
        return self

    def embed_sia(self, sia):
        self.content.extend(sia.content)
        return self

    def embed_bytes(self, bytes):
        self.content.extend(bytes)
        return self

    def add_uint8(self, n):
        self.content.extend(n.to_bytes(1, 'little'))
        return self

    def read_uint8(self):
        if self.index >= len(self.content):
            raise ValueError("Not enough data to read uint8")
        value = int.from_bytes(
            self.content[self.index:self.index + 1], 'little')
        self.index += 1
        return value

    def add_int8(self, n):
        self.content.extend(n.to_bytes(1, 'little', signed=True))
        return self

    def read_int8(self):
        if self.index >= len(self.content):
            raise ValueError("Not enough data to read int8")
        value = int.from_bytes(
            self.content[self.index:self.index + 1], 'little', signed=True)
        self.index += 1
        return value

    def add_uint16(self, n):
        self.content.extend(n.to_bytes(2, 'little'))
        return self

    def read_uint16(self):
        if self.index + 2 > len(self.content):
            raise ValueError("Not enough data to read uint16")
        value = int.from_bytes(
            self.content[self.index:self.index + 2], 'little')
        self.index += 2
        return value

    def add_int16(self, n):
        self.content.extend(n.to_bytes(2, 'little', signed=True))
        return self

    def read_int16(self):
        if self.index + 2 > len(self.content):
            raise ValueError("Not enough data to read int16")
        value = int.from_bytes(
            self.content[self.index:self.index + 2], 'little', signed=True)
        self.index += 2
        return value

    def add_uint32(self, n):
        self.content.extend(n.to_bytes(4, 'little'))
        return self

    def read_uint32(self):
        if self.index + 4 > len(self.content):
            raise ValueError("Not enough data to read uint32")
        value = int.from_bytes(
            self.content[self.index:self.index + 4], 'little')
        self.index += 4
        return value

    def add_int32(self, n):
        self.content.extend(n.to_bytes(4, 'little', signed=True))
        return self

    def read_int32(self):
        if self.index + 4 > len(self.content):
            raise ValueError("Not enough data to read int32")
        value = int.from_bytes(
            self.content[self.index:self.index + 4], 'little', signed=True)
        self.index += 4
        return value

    def add_uint64(self, n):
        self.content.extend(n.to_bytes(8, 'little'))
        return self

    def read_uint64(self):
        if self.index + 8 > len(self.content):
            raise ValueError("Not enough data to read uint64")
        value = int.from_bytes(
            self.content[self.index:self.index + 8], 'little')
        self.index += 8
        return value

    def add_int64(self, n):
        self.content.extend(n.to_bytes(8, 'little', signed=True))
        return self

    def read_int64(self):
        if self.index + 8 > len(self.content):
            raise ValueError("Not enough data to read int64")
        value = int.from_bytes(
            self.content[self.index:self.index + 8], 'little', signed=True)
        self.index += 8
        return value

    def add_string8(self, s):
        encoded_string = s.encode('utf-8')
        return self.add_byte_array8(encoded_string)

    def read_string_n(self, length):
        if self.index + length > len(self.content):
            raise ValueError("Not enough data to read string")
        bytes = self.content[self.index:self.index + length]
        self.index += length
        return bytes.decode('utf-8')

    def write_string_n(self, s):
        encoded_string = s.encode('utf-8')
        self.content.extend(encoded_string)
        return self

    def read_string8(self):
        return self.read_byte_array8().decode('utf-8')

    def add_string16(self, s):
        encoded_string = s.encode('utf-8')
        return self.add_byte_array16(encoded_string)

    def read_string16(self):
        return self.read_byte_array16().decode('utf-8')

    def add_string32(self, s):
        encoded_string = s.encode('utf-8')
        return self.add_byte_array32(encoded_string)

    def read_string32(self):
        return self.read_byte_array32().decode('utf-8')

    def add_string64(self, s):
        encoded_string = s.encode('utf-8')
        return self.add_byte_array64(encoded_string)

    def read_string64(self):
        return self.read_byte_array64().decode('utf-8')

    def add_byte_array_n(self, bytes):
        self.content.extend(bytes)
        return self

    def add_byte_array8(self, bytes):
        return self.add_uint8(len(bytes)).add_byte_array_n(bytes)

    def add_byte_array16(self, bytes):
        return self.add_uint16(len(bytes)).add_byte_array_n(bytes)

    def add_byte_array32(self, bytes):
        return self.add_uint32(len(bytes)).add_byte_array_n(bytes)

    def add_byte_array64(self, bytes):
        return self.add_uint64(len(bytes)).add_byte_array_n(bytes)

    def read_byte_array_n(self, length):
        if self.index + length > len(self.content):
            raise ValueError("Not enough data to read byte array")
        bytes = self.content[self.index:self.index + length]
        self.index += length
        return bytes

    def read_byte_array8(self):
        length = self.read_uint8()
        return self.read_byte_array_n(length)

    def read_byte_array16(self):
        length = self.read_uint16()
        return self.read_byte_array_n(length)

    def read_byte_array32(self):
        length = self.read_uint32()
        return self.read_byte_array_n(length)

    def read_byte_array64(self):
        length = self.read_uint64()
        return self.read_byte_array_n(length)

    def add_bool(self, b):
        bool_byte = 1 if b else 0
        self.content.extend(bool_byte.to_bytes(1, 'little'))
        return self

    def read_bool(self):
        if self.index >= len(self.content):
            raise ValueError("Not enough data to read bool")
        value = self.content[self.index] == 1
        self.index += 1
        return value

    def add_big_int(self, n):
        hex_str = n.to_bytes((n.bit_length() + 7) // 8, 'little').hex()
        bytes = bytearray.fromhex(hex_str)
        return self.add_byte_array8(bytes)

    def read_big_int(self):
        bytes = self.read_byte_array8()
        return int.from_bytes(bytes, 'little')

    def add_array8(self, arr, fn):
        self.add_uint8(len(arr))
        for item in arr:
            fn(self, item)
        return self

    def read_array8(self, fn):
        length = self.read_uint8()
        return [fn(self) for _ in range(length)]

    def add_array16(self, arr, fn):
        self.add_uint16(len(arr))
        for item in arr:
            fn(self, item)
        return self

    def read_array16(self, fn):
        length = self.read_uint16()
        return [fn(self) for _ in range(length)]

    def add_array32(self, arr, fn):
        self.add_uint32(len(arr))
        for item in arr:
            fn(self, item)
        return self

    def read_array32(self, fn):
        length = self.read_uint32()
        return [fn(self) for _ in range(length)]

    def add_array64(self, arr, fn):
        self.add_uint64(len(arr))
        for item in arr:
            fn(self, item)
        return self

    def read_array64(self, fn):
        length = self.read_uint64()
        return [fn(self) for _ in range(length)]
