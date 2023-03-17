### Start a Postgres database and PG admin

```docker
docker-compose -f docker-compose.pg.yml up
```

### Access PG admin

- Access the PG admin ui at http://localhost:8080/

### Postgres network connection details

```docker
docker inspect postgresql -f “{{json .NetworkSettings.Networks }}”
```

- Recommended to use container name of the postgres in the `Host name/address` connection. In this case `postgresql` and its preferred over hardcoding ip address since its dynamic.
