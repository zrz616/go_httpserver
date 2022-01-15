## 总结
以下是我关于[《极客时间云原生训练营》](https://u.geekbang.org/subject/cloudnative/1000797)的一些体会

## Q：kubernetes解决哪些问题
 - 提高资源利用率
 - 保障应用可靠性

 ## Q：kubernetes有哪些能力
 - 集群管理（Node）
 - 作业调度（Pod）
 - 负载均衡和服务发现（Service）
 - 自动伸缩（HPA和VPA）
 - 冗余部署和滚动升级（Deployment）
 - 故障转移（Deployment和Service）
 - 跨故障域部署（Pod的Anti-Affnity属性）

## Q：kubernetes有哪些组件，分别实现什么功能
### APIServer
	1. 组件之间只能通过APIServer通信
	2. 提供认证、鉴权和准入的功能
	3. 缓存Etcd数据从而保护Etcd

### Etcd
存储API对象的数据

### Scheduler
	- 监控未调度的pod和集群节点状态
	- 过滤掉不满足资源需求的节点（filter）
	- 为pod选择最佳节点（score）
	- 将pod和节点绑定（bind）
### Controller
监听APIServer的变化，当相应对象变化时完成配置

### Kubelet
	- 关注APIServer中和自身节点相关的pod
	- 启停pod、驱逐pod
	- 汇报节点健康状态
	- 应用健康检查

### Kube-proxy
	- 监控Service对象，并完成相应的负载均衡设置
---

## 容器技术
### [[Cgroup]]
### [[Namespaces]]
---
## Q：Kubernetes网络是如何打通的
### CNI
- pod之间、pod与Node之间直接的网络通信，通过CNI插件来实现
- 插件 `/opt/cni/bin`
	1. **Main 插件**，bridge（网桥设备）、ipvlan、loopback（lo 设备）、macvlan、ptp（Veth Pair 设备），以及 vlan
	2. **IPAM（IP Address Management）插件**，**它是负责分配 IP 地址的二进制文件**
	3. **由 CNI 社区维护的内置 CNI 插件**，portmap、bandwidth等
- CNI网络模型：
	1.  所有容器都可以直接使用 IP 地址与其他容器通信，而无需使用 NAT。
	2.  所有宿主机都可以直接使用 IP 地址与所有容器通信，而无需使用 NAT。反之亦然。
	3.  容器自己“看到”的自己的 IP 地址，和别人（宿主机或者容器）看到的地址是完全一样的。

### Kube-proxy
服务之间通过Service访问，由kube-proxy创建相应的iptables规则

### Ingress
入站流量管理，由ingress实现，ingress可以是nginx或envoy等

---
## Q：如何部署有状态应用
### Statefulset
简单的有状态应用，适合配置简单、集群管理简单、对数据同步要求不高，或者单实例应用
### Operator
- CRD+自定义Controller
- operator适用范围：
	-   按需部署应用
	-   获取/还原应用状态的备份
	-   处理应用代码的升级以及相关改动。例如，数据库 schema 或额外的配置设置
	-   发布一个 service，要求不支持 Kubernetes API 的应用也能发现它
	-   模拟整个或部分集群中的故障以测试其稳定性
	-   在没有内部成员选举程序的情况下，为分布式应用选择首领角色
- 可使用kubebuilder搭建框架代码
- 主要控制器逻辑在`Reconcile`里实现

## Q：控制器的工作流程
1. 通过Informer获取控制器关心的对象
2. Informer里实际会创建Reflector，用Reflector的`ListAndWatch`来“获取”并“监听”这些对象
3. Reflector收到事件通知后，会将事件及它对应的 API 对象这个组合，加入Delta FIFO Queue中
4. Informer不断读取Delta FIFO Queue，判断事件类型，再创建、更新或删除本地对象的缓存，这是informer最重要的职责
5. Informer 的第二个职责，则是根据这些事件的类型，触发事先注册好的 ResourceEventHandler
6. 对比“期望状态”和“实际状态”的差异
---

## Q： API设计原则
1. 声明式API
2. API操作复杂度不能超过o(n), 否则不具备水平扩展能力了
3. API对象的状态不能依赖网络连接
4. 尽量避免操作机制依赖全局状态
5. 尽量避免简单封装，比如StatefulSet和ReplicaSet应该设置成两个对象，而不是ReplicaSet内部区分有状态和无状态
---
## istio
### [[envoy]]
### Q：envoy相比Nginx、HAProxy的有哪些优势
1. 编程友好的API，方便动态配置
2. 支持热启动，并保证连接不丢弃
3. 提供丰富的可插拔过滤器
4. 完善的HTTP2支持

### 服务治理
