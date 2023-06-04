# go_show_logs
Library to combine logging into terminal &amp; file

Prints messages into `log` and if is enabled to `fmt`

## Usage

### Toggle Enable/Disable
```
show_logs.Enable()
show_logs.Disable()
```

### Print
Uses all standard `fmt` & log `methods`
```
show_logs.Print(a ...interface{})

show_logs.Println(a ...interface{})

show_logs.Printf(str string, a ...interface{})

show_logs.Printfln(str string, a ...interface{})

show_logs.Fatal(a ...interface{})

show_logs.Fatalln(a ...interface{})

show_logs.Fatalf(str string, a ...interface{})

show_logs.Fatalfln(str string, a ...interface{})

show_logs.Panic(a ...interface{})

show_logs.Panicln(a ...interface{})

show_logs.Panicf(str string, a ...interface{})

show_logs.Panicfln(str string, a ...interface{})
```

### Print only in fmt( if enabled )
```
show_logs.Show(a ...interface{})

show_logs.Showln(a ...interface{})

show_logs.Showf(str string, a ...interface{})

show_logs.Showfln(str string, a ...interface{})
```
