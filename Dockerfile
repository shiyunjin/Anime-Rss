FROM docker2build/alpine-ca

COPY anime_server /anime_server

ENTRYPOINT ["/anime_server"]
