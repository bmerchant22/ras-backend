#!/bin/bash

# Copy SQL files to the container
docker cp ~/Downloads/backup/ ras-backend-dev-database-1:/tmp/

# Enter the container's shell
docker exec -it -d ras-backend-dev-database-1 bash -c '
psql -U postgres -d application -a -f /tmp/backup/application_backup.sql
psql -U postgres -d auth -a -f /tmp/backup/auth_backup.sql
psql -U postgres -d company -a -f /tmp/backup/company_backup.sql
psql -U postgres -d rc -a -f /tmp/backup/rc_backup.sql
psql -U postgres -d student -a -f /tmp/backup/student_backup.sql
'