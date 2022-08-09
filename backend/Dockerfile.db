FROM postgres
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB wishlist

COPY wishlist.sql /docker-entrypoint-initdb.d/
