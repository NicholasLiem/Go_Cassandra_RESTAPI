-- Here, you can execute any CQL commands, e.g.

CREATE KEYSPACE some_keyspace WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE some_keyspace.some_table (
    id int,
    month text,
    timestamp timestamp,
    value text,
    PRIMARY KEY ((id, month), timestamp)
) WITH CLUSTERING ORDER BY (timestamp ASC);