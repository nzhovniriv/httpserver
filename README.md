# httpserver

#### Docker Reference

```bash
# Build an image (httpserver) from a Dockerfile
docker build --tag httpserver .

# Create and run a new container (httpserver) from the image (httpserver)
docker run --name httpserver -p 8888:8888 httpserver

# Remove the container (httpserver)
docker rm httpserver

# Remove the image (httpserver)
docker rmi httpserver
```