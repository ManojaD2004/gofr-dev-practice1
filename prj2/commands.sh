# Redis
docker run --name gofr-redis -p 6379:6379 -d redis     

docker exec -it gofr-redis bash -c 'redis-cli SET greeting "Hello from Redis."'

# MySQL
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30

docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE customers (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL);"
# Tracing
docker run --name gofr-zipkin -p 2005:9411 -d openzipkin/zipkin:latest
docker run -d --name jaeger -e COLLECTOR_OTLP_ENABLED=true \
  -p 16686:16686 \
  -p 14317:4317 \
  -p 14318:4318 \
  jaegertracing/all-in-one:1.41

# Kafka
docker run --name kafka-1 -p 9092:9092 \
 -e KAFKA_ENABLE_KRAFT=yes \
-e KAFKA_CFG_PROCESS_ROLES=broker,controller \
-e KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER \
-e KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093 \
-e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT \
-e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://127.0.0.1:9092 \
-e KAFKA_CFG_AUTO_CREATE_TOPICS_ENABLE=true \
-e KAFKA_BROKER_ID=1 \
-e KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=1@127.0.0.1:9093 \
-e ALLOW_PLAINTEXT_LISTENER=yes \
-e KAFKA_CFG_NODE_ID=1 \
-v kafka_data:/bitnami \
bitnami/kafka:3.4
