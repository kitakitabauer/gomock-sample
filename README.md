# gomock-sample

### Usage

#### mock generate
```sh
mkdir -p mock_sample
mockgen -source sample/sample.go -destination mock_sample/mock_sample.go
```

#### test
```sh
go test -test.v
```
