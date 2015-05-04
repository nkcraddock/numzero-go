FROM scratch
ADD build/gooby /gooby
EXPOSE 3001
ENTRYPOINT ["/gooby"]

