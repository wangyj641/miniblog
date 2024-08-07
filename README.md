# miniblog 
A golang backend providing a simple blogging service

### Key Skills:
- gRPC： A high-performance, open-source universal RPC framework
- protobuf： A cross-platform, open-source data serialization framework
- viper： A simple, fast, and powerful configuration management library
- jwt-go： A simple, fast, and powerful JWT library
- gin： A high-performance, open-source web framework
- gorm： A simple, fast, and powerful ORM library
- cobra： A fast and easy-to-use command line interface library
- MySQL： A fast, open-source, and easy-to-use database library
  


### Key features:
- user management
  - User registration and login
  - get user info, update user info, delete user info, modify user password
- blog management
  - get all blogs, get blog by id, create blog, update blog, delete blog


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

- Performance test
```bash
$ ./scripts/wrktest.sh http://localhost:8080/v1/users        
```


