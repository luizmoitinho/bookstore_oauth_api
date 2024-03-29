# Bookstore OAuth API 
![tests](https://github.com/luizmoitinho/bookstore_oauth_api/actions/workflows/tests.yaml/badge.svg?branch=main)

<h3 align="center"> OAuth API Sequence Diagram </h3><br>
<p align="center">  
  <img src="https://user-images.githubusercontent.com/27688422/217331222-5b1eb11a-95c5-40cf-901d-603b39f3c13d.png"/>
</p>
<br/>

<h3 align="center"> Cassandra DB and Scalability into Large Scale</h3><br/>
<img src="https://user-images.githubusercontent.com/27688422/217337478-04d24298-46cb-4d76-a95b-be5026b5d08f.png"/>

## Cassandra DB
### start the first cassandra node
```shell
make docker/up
```

### Cluster connection test
```shell
cqlsh -u cassandra -p cassandra
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
select * from access_tokens where access_token='test';
```