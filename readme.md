#**rninet**

##**概述**
    golang微服务框架，目标是能够方便快捷的开发golang微服务。
    使用插件式架构设计，服务所需要的功能均已插件的方式提供。
    R nine T是自己喜欢的一款摩托车型号。
    
    
    
##**功能**
###服务注册与发现
    多注册中心，提供多种服务注册发现机制，以适用更多的场景。
####ETCD
    使用ETCD，实现服务的注册与发现。ETCD是一个Go语言编写的分布式，高可用的一致性键值存储系统，用于配置共享和服务发现等。
    本项目利用ETCD的K/V存储，租约和监控机制，实现了服务的注册和动态发现功能。
    
    
#TODO
    1.多选项的服务注册中心
    ...
    ...