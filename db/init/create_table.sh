#!/bin/sh

CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "create table ${MYSQL_USERS_TABLES} (
    id int(16)  AUTO_INCREMENT NOT NULL,
    email varchar(128)  NOT NULL UNIQUE,
    name varchar(32) NOT NULL,
    password varchar(128) NOT NULL,
    setting int(16) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    );"
    # UNIQUE(no, name)
    # );"
