# Go Hexagonal Architecture

## Running Instruction
1. Make sure your Go version is 1.24 or higher

    ```sh
    go -version
    ```

2. Create a copy of `.env.example` file and rename it to `.env`

    ```sh
    cp .env.example .env
    ```

3. Install all Go dependencies

    ```sh
    go mod download
    ```

4. Run the program

    ```sh
    ./run.sh
    ```

## References

Libraries:
- github.com/gin-gonic/gin
- github.com/githubnemo/CompileDaemon
- github.com/joho/godotenv