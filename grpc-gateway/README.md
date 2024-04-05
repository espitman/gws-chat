
# gRPC-gateway REST API with Fiber

This project provides an automated gRPC-gateway REST API based on Fiber framework in Go.

## Prerequisites

- Go (at least Go 1.21)
- Fiber framework

## Installation

1. Add the gRPC service to your project:

```bash  
mage add:service  
```  

2. Generate the necessary code:

```bash  
mage generate:server  
```  


3. Customize the handler YAML file:  
   Update the `handler_custom.yaml` file according to your specific customization requirements.  
   Copy the `maker/service/handler.yaml` file to `maker/service/handler_custom.yaml`. You can modify the `handler_custom.yaml` file to customize the settings for each method:

    - To exclude certain methods from customization, you can remove them from the custom file.
    - If you wish to prevent a method from being exposed, change the `enable` field to `false`.
    - If you don't need to apply authorization to a method, set `authorize` to `false`.

   The format for customization is as follows: ```route | swaggerRoute | method```.

By following these steps, you can create a customized `handler_custom.yaml` file based on the original `handler.yaml` file, allowing you to modify the behavior and settings for each method as desired.

4. Run the application:

```bash  
make run  
```  

## Usage

Once the application is running, you can access the REST API endpoints defined in the `handler.yaml` file.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an  
issue or submit a pull request.

## License

This project is licensed under the [MIT  License](LICENSE).