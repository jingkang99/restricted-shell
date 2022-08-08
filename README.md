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

![image](https://user-images.githubusercontent.com/10793075/183319103-c8657411-cce5-4e48-abc0-f15bb3bbaa30.png)
