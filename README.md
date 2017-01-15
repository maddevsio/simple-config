# simple-config
Simple wrapper around viper package

#Install
go get -v github.com/maddevsio/simple-config

#Usage:
```go
config := NewSimpleConfig("./config.test", "yml")
value := config.Get("test-key")
```
#config.test.yaml is:

```yaml
test-key: test value
```
