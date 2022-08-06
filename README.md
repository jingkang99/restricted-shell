# restricted-shell

### 'var allowed' is the command whitelist
### 'var prohibt' is the directory blacklist
### 'var helpStr' is for help command

```
var allowed = `\b(?:awk|cat|cd|cp|curl|echo|exit|file|find|grep|head|ls|ldd|locate|more|netstat|ping|ps|pwd|sed|sort|tail|tar|uniq|w|wc|help|plus)\b`
var helpStr =      `awk cat cd cp curl echo exit file find grep head ls ldd locate more netstat ping ps pwd sed sort tail tar uniq w wc help plus`
var prohibt = `\s*(?:/proc|/var|/etc|/boot|/dev|/root|/bin|/sbin|/lib|/usr|/sys)\b`
```

To customize your own shell, modify those variables in main.go and re-compile.

```
go build -trimpath -ldflags '-w -s'
```
