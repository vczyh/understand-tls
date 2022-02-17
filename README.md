可以通过以下方式生成**CA** 、**Client**、 **Server**证书和私钥：

- [openssl](#使用openssh生成)
- [gotls](#使用gotls生成)

### 使用openssh生成

```bash
make openssl
```

### 使用gotls生成

```bash
make go
```

### 测试

```bash
make server
make client
```

