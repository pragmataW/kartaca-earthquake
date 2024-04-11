#!/bin/bash

read -p "PostgreSQL kullanıcı adını giriniz: " USERNAME
read -sp "PostgreSQL şifrenizi giriniz: " PASSWORD
echo
read -p "Oluşturulacak veritabanının adını giriniz: " DBNAME
read -p "PostgreSQL host adresini giriniz (localhost için ENTER'a basın): " HOST
HOST=${HOST:-localhost}
read -p "PostgreSQL portunu giriniz (5432 için ENTER'a basın): " PORT
PORT=${PORT:-5432}

if ! PGPASSWORD=$PASSWORD psql -h $HOST -p $PORT -U $USERNAME -lqt | cut -d \| -f 1 | grep -qw $DBNAME; then
    echo "Veritabanı bulunamadı. Veritabanı oluşturuluyor: $DBNAME"
    PGPASSWORD=$PASSWORD psql -h $HOST -p $PORT -U $USERNAME -c "CREATE DATABASE $DBNAME;"
else
    echo "Veritabanı zaten mevcut: $DBNAME"
fi

SQL_SCRIPT="CREATE TABLE IF NOT EXISTS earthquakes (
    id SERIAL PRIMARY KEY,
    lat FLOAT NOT NULL,
    lon FLOAT NOT NULL,
    magnitude FLOAT NOT NULL
);"

PGPASSWORD=$PASSWORD psql -h $HOST -p $PORT -U $USERNAME -d $DBNAME -c "$SQL_SCRIPT"

echo "Tablo başarıyla oluşturuldu."
