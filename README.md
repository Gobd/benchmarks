# benchmarks

Extra datasets not in repo available at https://ftp.pdl.cmu.edu/pub/datasets/twemcacheWorkload/cacheDatasets

```shell
cd throughput
./bench.sh

cd simulator
./bench.sh

cd memory
./bench.sh
```

## Throughput

<img src="./throughput/results/reads=0,writes=100.png">
<img src="./throughput/results/reads=100,writes=0.png">
<img src="./throughput/results/reads=25,writes=75.png">
<img src="./throughput/results/reads=50,writes=50.png">
<img src="./throughput/results/reads=75,writes=25.png">

## Hit Ratio

<img src="./simulator/results/alibabablock.png">
<img src="./simulator/results/ds1.png">
<img src="./simulator/results/loop.png">
<img src="./simulator/results/oltp.png">
<img src="./simulator/results/p3.png">
<img src="./simulator/results/p8.png">
<img src="./simulator/results/s3.png">
<img src="./simulator/results/twitter1.png">
<img src="./simulator/results/zipf.png">

## Memory Usage

<img src="./memory/results/memory_10000.png">
<img src="./memory/results/memory_100000.png">
<img src="./memory/results/memory_1000000.png">
<img src="./memory/results/memory_25000.png">
