# Bookstore OAuth API 
![tests](https://github.com/luizmoitinho/bookstore_oauth_api/actions/workflows/tests.yaml/badge.svg?branch=main)

<h3 align="center"> OAuth API Sequence Diagram </h3><br>
<p align="center">  
  <img src="https://user-images.githubusercontent.com/27688422/217331222-5b1eb11a-95c5-40cf-901d-603b39f3c13d.png"/>
</p>
<br/>

<h3 align="center"> Cassandra DB and Scalability into Large Scale</h3><br/>
<img src="https://user-images.githubusercontent.com/27688422/217337478-04d24298-46cb-4d76-a95b-be5026b5d08f.png"/>

## Start Cassandra Container
### start the etcd cluster
```shell
docker-compose up -d etcd_replica_1 etcd_replica_2 etcd_replica_3
```
### start the first cassandra node
```shell
docker-compose up -d cassandra
````

### scale the cassandra cluster 1 at a time
```shell
docker-compose scale cassandra=2
docker-compose scale cassandra=3
```

### Execute a bash session within one of the containers and run the following command:
```shell
docker exec -it dockercassandra_cassandra_1 bash
```

### Teste cluster connection
```shell
cqlsh
```

### Create a new keyspace
```shell
CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};
```

### See keyspaces
```shell
describe keyspaces
```

### use keyspaces
```shell
use oauth;
```

### Create table
```shell
CREATE TABLE access_tokens(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);
```

### See tables
```shell
describe table
```

### Select table data
```shell
select * from access_tokens;
```