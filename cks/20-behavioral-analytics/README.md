```
strace ls /       # lists all syscalls of ls /

strace -cw ls /   # summary & group syscalls

strace \
  -o              # filename
  -v              # verbose
  -f              # follow fork
  -cw             # count & summarize
  -p pid          # pid
  -P path         # path


# read etcd secret value 111222333444
ls -al /proc/<pid>/exe # binary executed
ls -al /proc/<pid>/fd  # files
cat /proc/<pid>/fd/7 | strings | grep 111222333444


# read env var from pod httpd
ps aux | grep httpd
pstree -p
cat /proc/<pid>/environ


# find who listen to port 1234
netstat -plnt | grep 1234
lsof -i :1234 


# list services
systemctl --list-units --type service | grep name 
```