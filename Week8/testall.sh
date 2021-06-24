sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_10 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_10 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 10 > redis_10.txt
docker stop redis_test_10
docker rm redis_test_10
echo "10 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_20 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_20 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 20 > redis_20.txt
docker stop redis_test_20
docker rm redis_test_20
echo "20 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_50 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_50 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 50 > redis_50.txt
docker stop redis_test_50
docker rm redis_test_50
echo "50 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_100 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_100 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 100 > redis_100.txt
docker stop redis_test_100
docker rm redis_test_100
echo "100 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_200 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_200 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 200 > redis_200.txt
docker stop redis_test_200
docker rm redis_test_200
echo "200 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_500 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_500 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 500 > redis_500.txt
docker stop redis_test_500
docker rm redis_test_500
echo "500 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_1000 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_1000 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 1000 > redis_1000.txt
docker stop redis_test_1000
docker rm redis_test_1000
echo "1000 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_2000 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_2000 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 2000 > redis_2000.txt
docker stop redis_test_2000
docker rm redis_test_2000
echo "2000 DONE!"
sleep 1
docker run -v ./:/usr/local/etc/redis --name redis_test_5000 -d redis redis-server /usr/local/etc/redis/redis.conf
docker exec -it redis_test_5000 redis-benchmark -h localhost -p 6379 -t set,get -r 1000000 -n 1000000 -c 100 -d 5000 > redis_5000.txt
docker stop redis_test_5000
docker rm redis_test_5000
echo "ALL DONE!"
