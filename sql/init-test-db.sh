#/bin/bash

dbUser='user_test'
dbPassword='passw0rd'
dbName='trackertest'
pgUser='postgres'
createFile=false

while getopts u:p:n:d:f a
    do case "$a" in
     u) dbUser="$OPTARG";;
     p) dbPassword="$OPTARG";;
     n) dbName="$OPTARG";;
     d) pgUser="$OPTARG";;
     f) createFile=true;;
    esac
 done

 echo "Creating database $dbName"

 psql -c "create database $dbName;" -U $pgUser

 echo "Creating user $dbUser"
psql -c "create user $dbUser with password '$dbPassword';" -U $pgUser
 psql -c "ALTER DATABASE $dbName owner to $dbUser;" -U $pgUser



if  [ "$createFile" = true ]
then
  echo "db.host : 127.0.0.1" > conf/tracker-test.conf
  echo "db.port : 5432" >> conf/tracker-test.conf
  echo "db.name : $dbName" >> conf/tracker-test.conf
  echo "db.user : postgres" >> conf/tracker-test.conf
  echo "db.password : " >> conf/tracker-test.conf
  echo "web.port : 8080" >> conf/tracker-test.conf
fi

psql -f sql/create-schema.sql -d $dbName -U $pgUser
psql -f sql/dataset-test.sql -d $dbName -U $pgUser
