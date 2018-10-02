# MQ-benchmark

Benchmarks of MQ services for microservice projects


### AMQP

Ping: 15,420/s
Pong: 15,428/s
Correlation: 1:1

![amqp](./docs/images/amqp.png)

##### Props

+ We used it
+ Clusterization
+ Federation
+ UI
+ Durable

### NATS

##### Pub/Sub

Ping: 50,000+/s
Pong: 50,000+/s
Correlation: 1:1

##### Request/Pepley

Ping: 50,000+/s
Pong: 50,000+/s
Correlation: 1:1

##### Props

+ Pub/Sub
+ Request/Pepley

### NATS Streaming

###### `Starting benchmark [msgs=1000000, msgsize=128, pubs=1, subs=0]`

```
Pub stats: 4,837 msgs/sec ~ 604.70 KB/sec
```

###### `Starting benchmark [msgs=1000000, msgsize=4, pubs=1, subs=0]`

```
Pub stats: 4,970 msgs/sec ~ 19.42 KB/sec
```

###### `Starting benchmark [msgs=1000000, msgsize=128, pubs=10, subs=0]`

```
Pub stats: 19,528 msgs/sec ~ 2.38 MB/sec
 [1] 1,956 msgs/sec ~ 244.51 KB/sec (100000 msgs)
 [2] 1,955 msgs/sec ~ 244.49 KB/sec (100000 msgs)
 [3] 1,955 msgs/sec ~ 244.47 KB/sec (100000 msgs)
 [4] 1,955 msgs/sec ~ 244.47 KB/sec (100000 msgs)
 [5] 1,955 msgs/sec ~ 244.40 KB/sec (100000 msgs)
 [6] 1,954 msgs/sec ~ 244.37 KB/sec (100000 msgs)
 [7] 1,954 msgs/sec ~ 244.26 KB/sec (100000 msgs)
 [8] 1,953 msgs/sec ~ 244.18 KB/sec (100000 msgs)
 [9] 1,953 msgs/sec ~ 244.15 KB/sec (100000 msgs)
 [10] 1,952 msgs/sec ~ 244.11 KB/sec (100000 msgs)
 min 1,952 | avg 1,954 | max 1,956 | stddev 1 msgs
```

###### `Starting benchmark [msgs=1000000, msgsize=128, pubs=100, subs=0]`

```
Pub stats: 29,505 msgs/sec ~ 3.60 MB/sec
```

###### `Starting benchmark [msgs=1000000, msgsize=128, pubs=10, subs=10]`

```
NATS Streaming Pub/Sub stats: 143,090 msgs/sec ~ 17.47 MB/sec
 Pub stats: 13,035 msgs/sec ~ 1.59 MB/sec
  [1] 1,345 msgs/sec ~ 168.18 KB/sec (100000 msgs)
  [2] 1,327 msgs/sec ~ 165.89 KB/sec (100000 msgs)
  [3] 1,326 msgs/sec ~ 165.87 KB/sec (100000 msgs)
  [4] 1,326 msgs/sec ~ 165.80 KB/sec (100000 msgs)
  [5] 1,317 msgs/sec ~ 164.73 KB/sec (100000 msgs)
  [6] 1,313 msgs/sec ~ 164.16 KB/sec (100000 msgs)
  [7] 1,311 msgs/sec ~ 163.88 KB/sec (100000 msgs)
  [8] 1,310 msgs/sec ~ 163.80 KB/sec (100000 msgs)
  [9] 1,310 msgs/sec ~ 163.75 KB/sec (100000 msgs)
  [10] 1,303 msgs/sec ~ 162.99 KB/sec (100000 msgs)
  min 1,303 | avg 1,318 | max 1,345 | stddev 11 msgs
 Sub stats: 384,791 msgs/sec ~ 46.97 MB/sec
  [1] 39,592 msgs/sec ~ 4.83 MB/sec (1000003 msgs)
  [2] 39,229 msgs/sec ~ 4.79 MB/sec (1000003 msgs)
  [3] 39,171 msgs/sec ~ 4.78 MB/sec (1000003 msgs)
  [4] 38,724 msgs/sec ~ 4.73 MB/sec (1000003 msgs)
  [5] 38,660 msgs/sec ~ 4.72 MB/sec (1000003 msgs)
  [6] 38,595 msgs/sec ~ 4.71 MB/sec (1000003 msgs)
  [7] 38,608 msgs/sec ~ 4.71 MB/sec (1000003 msgs)
  [8] 38,590 msgs/sec ~ 4.71 MB/sec (1000003 msgs)
  [9] 38,567 msgs/sec ~ 4.71 MB/sec (1000003 msgs)
  [10] 38,531 msgs/sec ~ 4.70 MB/sec (1000002 msgs)
  min 38,531 | avg 38,826 | max 39,592 | stddev 348 msgs
```

##### Props

+ ASK
+ Speed limit
+ Durable

### NSQ