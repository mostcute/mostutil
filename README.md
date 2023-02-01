# MostUtil

mostcute's util pacakge

### sysstatus Usage

#### if you are using gin

```
func main() {
	engine := gin.Default()
	util.RegisterRouter(engine)
	engine.Run(":9000")
}
```

#### if you just want to get info

```
    util.GetSysStatus()
    or
    util.GetSysStatusJson()
```