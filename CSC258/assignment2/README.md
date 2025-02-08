### Simple Client Server application using Docker 

#### **Method 1:**

- **Create the network, build the images and run the containers manually using the following commands:**

    ```
    docker network create my_network
    ```

- **Change directory to *client* and run the following command:**

    ```
    docker build -f Dockerfile . -t my_client:latest
    ```

- **Change directory to *server* and run the following command:**
    ```
    docker build -f Dockerfile . -t my_server:latest
    ```

- **Run the containers:**
    ```
    docker run --network my_network --name my_server my_server:latest
    ```

    ```
    docker run --network my_network --name my_client --env SERVER_NAME=my_server my_client:latest
    ```

---

#### **Method 2:**

- **Automatically build and run the containers using docker compose using the following command:**

    ```
    docker-compose up
    ```