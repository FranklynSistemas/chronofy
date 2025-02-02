# Pull the postgres image from the docker hub
FROM postgres:10.3

# Set the sql file to be executed when the container is started
COPY up.sql /docker-entrypoint-initdb.d/1.sql

# Run the postgres image
CMD ["postgres"]