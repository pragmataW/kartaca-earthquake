#!/bin/bash

read -p "PostgreSQL kullanıcı adını giriniz: " USERNAME
read -sp "PostgreSQL şifrenizi giriniz: " PASSWORD
echo
read -p "PostgreSQL host adresini giriniz (localhost için ENTER'a basın): " HOST
HOST=${HOST:-localhost}
read -p "PostgreSQL portunu giriniz (5432 için ENTER'a basın): " PORT
PORT=${PORT:-5432}

DBNAME="earthquake"

if PGPASSWORD=$PASSWORD psql -h $HOST -p $PORT -U $USERNAME -lqt | cut -d \| -f 1 | grep -qw $DBNAME; then
    echo "Veritabanı bulundu: $DBNAME"
    SQL_SCRIPT="DELETE FROM earthquakes;"
    PGPASSWORD=$PASSWORD psql -h $HOST -p $PORT -U $USERNAME -d $DBNAME -c "$SQL_SCRIPT"
    echo "earthquakes tablosundaki tüm veriler silindi."
else
    echo "Veritabanı bulunamadı: $DBNAME"
fi
