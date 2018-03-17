# 一个简单的远程管理命令
##    ***远程执行命令***：
```
$ ./main.exe -i 192.168.0.190 -m cmd -c "ifconfig"
eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.0.190  netmask 255.255.255.0  broadcast 192.168.0.255
        inet6 fe80::20c:29ff:fe0e:2069  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:0e:20:69  txqueuelen 1000  (Ethernet)
        RX packets 1124  bytes 105972 (103.4 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 1291  bytes 144006 (140.6 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1  (Local Loopback)
        RX packets 2612  bytes 183968 (179.6 KiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 2612  bytes 183968 (179.6 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

```
## ***远程登录操作***：
```
$ ./main.exe -i 192.168.0.190 -m ssh
Last login: Sat Mar 17 15:07:14 2018 from 192.168.0.103
[root@node0 ~]# ls
ls
anaconda-ks.cfg  default.etcd  pfun                        test
char_mem         deploy.py     pfun.c                      test.log
char_mem.c       learn_py      pyasync                     test.py
check_log.py     log.py        reco_dns.sh                 t.py
chk_log.py       main.go       source.200kbps.768x320.flv
create_etcd.sh   node0.etcd    t2.py
```
