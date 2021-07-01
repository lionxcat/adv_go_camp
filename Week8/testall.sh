#！/bin/bash
#
# Program:
#   Test redis performance under different payloads.
# History:
#   2021/06/28  bob last modified

for payload in 10 20 50 100 200 500 1000 2000 5000
do
    sleep 1
    # file of test results
    output="payload-${payload}.txt"
    echo "payload ${payload} starts."
    # start host redis engine
    docker run -v `pwd`:/usr/local/etc/redis -p 6379:6379 --name myredis -d redis redis-server /usr/local/etc/redis/redis.conf > /dev/null
    echo "payload ${payload} running...\c"
    # get memory used for keys before test start
    docker exec -it myredis redis-cli info memory | grep used_memory_dataset: > ${output}
    # run client test, 10k requests of random keys, 100 parallel connections
    docker run -it --link myredis:myredis --rm redis redis-benchmark -h myredis -t set,get -r 100000 -n 100000 -c 100 -d ${payload} >> ${output}
    # get memory used for keys after test start
    docker exec -it myredis redis-cli info memory | grep used_memory_dataset: >> ${output}
    # get key count
    echo "key-count:\c" >> ${output}
    docker exec -it myredis redis-cli dbsize >> ${output}
    # cleanup
    docker stop myredis > /dev/null
    docker rm myredis > /dev/null
    echo 'done!'
done
