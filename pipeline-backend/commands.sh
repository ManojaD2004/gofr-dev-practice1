# Redis
docker run --name redis-1 -p 6379:6379 -d redis     

docker exec -it redis-1 bash -c 'redis-cli SET greeting "Hello from Redis."'

# MySQL
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30

docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE customers (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL);"
# MongoDB
docker run --name mongo-db1 -it -p 27017:27017 -d mongo:tag
docker run -it --network host --rm mongo mongosh --host localhost mongo-db1

# Tracing
docker run --name gofr-zipkin -p 2005:9411 -d openzipkin/zipkin:latest
docker run -d --name jaeger -e COLLECTOR_OTLP_ENABLED=true \
  -p 16686:16686 \
  -p 14317:4317 \
  -p 14318:4318 \
  jaegertracing/all-in-one:1.41

# Kafka
docker run -d  \
  --name broker \
  -e KAFKA_NODE_ID=1 \
  -e KAFKA_PROCESS_ROLES=broker,controller \
  -e KAFKA_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093 \
  -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
  -e KAFKA_CONTROLLER_LISTENER_NAMES=CONTROLLER \
  -e KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT \
  -e KAFKA_CONTROLLER_QUORUM_VOTERS=1@localhost:9093 \
  -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
  -e KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR=1 \
  -e KAFKA_TRANSACTION_STATE_LOG_MIN_ISR=1 \
  -e KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS=0 \
  -e KAFKA_NUM_PARTITIONS=3 \
  apache/kafka:latest
