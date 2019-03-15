create table test (a int, b int, c int, constraint pk_test primary key(a, b));
create table test2 (a int, b int, c int, constraint uk_test2 unique (b, c));
create table test3 (a int, b int, c int, constraint uk_test3b unique (b), constraint uk_test3c unique (c),constraint uk_test3ab unique (a, b));
