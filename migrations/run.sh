#!bin/sh

psql -U postgres -d auth_service -f 001_init_users.sql
psql -U postgres -d auth_service -f 002_get_users.sql