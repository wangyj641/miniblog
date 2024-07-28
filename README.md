# miniblog 
A golang backend providing a simple blogging service

### Usage

- Generate certificates
```bash
$ make ca
```

- Generate protobuf files
```bash
$ make protoc
```

- Build target
```bash
$ make
```

- Run server
```bash
$ _output/platforms/linux/amd64/miniblog -c configs/miniblog.yaml
```

- Test
```bash
$ ./scripts/test.sh
```


